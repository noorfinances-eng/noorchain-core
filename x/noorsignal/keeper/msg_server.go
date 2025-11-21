package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// MsgServer est le point d'entrée pour les transactions (Msg)
// du module PoSS (noorsignal).
type MsgServer struct {
	Keeper     Keeper
	BankKeeper noorsignaltypes.BankKeeper // <— pré-câblé (future utilisation)
}

// NewMsgServer construit un MsgServer avec BankKeeper (préparation future).
func NewMsgServer(k Keeper, bk noorsignaltypes.BankKeeper) MsgServer {
	return MsgServer{
		Keeper:     k,
		BankKeeper: bk, // <— pas encore utilisé
	}
}

// ---------------------------------------------------------
// SubmitSignal : pas de modification dans cette étape
// ---------------------------------------------------------
func (s MsgServer) SubmitSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSubmitSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Weight == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "weight must be >= 1")
	}
	if msg.Weight > 100 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "weight must be <= 100")
	}

	participantAddr, err := msg.GetParticipantAddress()
	if err != nil {
		return nil, err
	}

	cfg, found := s.Keeper.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	var dayBucket uint64
	if cfg.MaxSignalsPerDay > 0 {
		ts := ctx.BlockTime().Unix()
		if ts < 0 {
			ts = 0
		}
		dayBucket = uint64(ts) / 86400

		current := s.Keeper.GetDailySignalCount(ctx, participantAddr, dayBucket)
		if current >= cfg.MaxSignalsPerDay {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "daily signal limit reached")
		}
	}

	signal := noorsignaltypes.Signal{
		Participant:       participantAddr,
		Curator:           nil,
		Weight:            msg.Weight,
		Time:              ctx.BlockTime(),
		Metadata:          msg.Metadata,
		TotalReward:       0,
		RewardParticipant: 0,
		RewardCurator:     0,
	}

	_ = s.Keeper.CreateSignal(ctx, signal)

	if cfg.MaxSignalsPerDay > 0 {
		s.Keeper.IncrementDailySignalCount(ctx, participantAddr, dayBucket)
	}

	return &sdk.Result{}, nil
}

// ---------------------------------------------------------
// ValidateSignal : toujours sans BankKeeper (phase future)
// ---------------------------------------------------------
func (s MsgServer) ValidateSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgValidateSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	curatorAddr, err := msg.GetCuratorAddress()
	if err != nil {
		return nil, err
	}

	if !s.Keeper.IsActiveCurator(ctx, curatorAddr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "curator not active or not authorized")
	}

	signal, found := s.Keeper.GetSignal(ctx, msg.SignalId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "signal not found")
	}

	if signal.Curator != nil && len(signal.Curator) > 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "signal already validated")
	}

	total, part, cur, ok := s.Keeper.ComputeSignalRewardsCurrentEra(ctx, signal.Weight)
	if !ok {
		total = 0
		part = 0
		cur = 0
	}

	signal.Curator = curatorAddr
	signal.TotalReward = total
	signal.RewardParticipant = part
	signal.RewardCurator = cur

	s.Keeper.SetSignal(ctx, signal)

	s.Keeper.IncrementCuratorValidatedCount(ctx, curatorAddr)

	// ---------------------------------------------------------
	// FUTURE
	// Ici, on utilisera BankKeeper.SendCoins
	// pour créditer :
	//   - participant
	//   - curator
	// depuis la réserve PoSS
	// ---------------------------------------------------------

	return &sdk.Result{}, nil
}

// ---------------------------------------------------------
// AddCurator / RemoveCurator / SetConfig
// (identiques à la version précédente)
// ---------------------------------------------------------

func (s MsgServer) AddCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgAddCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid curator address")
	}

	curator, found := s.Keeper.GetCurator(ctx, curatorAddr)
	if !found {
		curator = noorsignaltypes.Curator{
			Address:               curatorAddr,
			Level:                 msg.Level,
			TotalSignalsValidated: 0,
			Active:                true,
		}
	} else {
		curator.Level = msg.Level
		curator.Active = true
	}

	s.Keeper.SetCurator(ctx, curator)
	return &sdk.Result{}, nil
}

func (s MsgServer) RemoveCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgRemoveCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid curator address")
	}

	curator, found := s.Keeper.GetCurator(ctx, curatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "curator not found")
	}

	curator.Active = false
	s.Keeper.SetCurator(ctx, curator)

	return &sdk.Result{}, nil
}

func (s MsgServer) SetConfig(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSetConfig,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	baseReward, err := strconv.ParseUint(msg.BaseReward, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid base_reward")
	}

	totalRatio := msg.ParticipantRatio + msg.CuratorRatio
	if totalRatio != 100 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "participant_ratio + curator_ratio must = 100")
	}

	if msg.EraIndex > uint64(^uint32(0)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "era_index out of range")
	}

	newCfg := noorsignaltypes.PossConfig{
		BaseReward:       baseReward,
		ParticipantShare: msg.ParticipantRatio,
		CuratorShare:     msg.CuratorRatio,
		MaxSignalsPerDay: msg.MaxSignalsPerDay,
		Enabled:          true,
		EraIndex:         uint32(msg.EraIndex),
	}

	s.Keeper.SetConfig(ctx, newCfg)
	return &sdk.Result{}, nil
}
