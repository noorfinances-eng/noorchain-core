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
	Keeper
}

// NewMsgServer construit un MsgServer à partir d'un Keeper PoSS.
func NewMsgServer(k Keeper) MsgServer {
	return MsgServer{Keeper: k}
}

// SubmitSignal gère la réception d'un MsgSubmitSignal
// (émission d'un nouveau signal social PoSS).
func (s MsgServer) SubmitSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSubmitSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Valider le poids du signal.
	if msg.Weight == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "weight must be >= 1")
	}
	if msg.Weight > 100 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "weight must be <= 100")
	}

	// 2) Convertir l'adresse du participant.
	participantAddr, err := msg.GetParticipantAddress()
	if err != nil {
		return nil, err
	}

	// 3) Récupérer la configuration PoSS (pour MaxSignalsPerDay).
	cfg, found := s.Keeper.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	// 4) Vérifier la limite quotidienne si active.
	var dayBucket uint64
	if cfg.MaxSignalsPerDay > 0 {
		ts := ctx.BlockTime().Unix()
		if ts < 0 {
			ts = 0
		}
		dayBucket = uint64(ts) / 86400

		current := s.getDailySignalCount(ctx, participantAddr, dayBucket)
		if current >= cfg.MaxSignalsPerDay {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "daily signal limit reached")
		}
	}

	// 5) Construire un Signal de base (sans curator, sans récompense).
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

	// 6) Créer et stocker le signal via le Keeper (on récupère l'ID créé).
	created := s.Keeper.CreateSignal(ctx, signal)

	// 7) Incrémenter le compteur quotidien si une limite est active.
	if cfg.MaxSignalsPerDay > 0 {
		s.incrementDailySignalCount(ctx, participantAddr, dayBucket)
	}

	// 8) Émettre l'event poss.signal_submitted pour les explorateurs / indexers.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			noorsignaltypes.EventTypeSignalSubmitted,
			sdk.NewAttribute(noorsignaltypes.AttrKeySignalID, strconv.FormatUint(created.Id, 10)),
			sdk.NewAttribute(noorsignaltypes.AttrKeyParticipant, participantAddr.String()),
			sdk.NewAttribute(noorsignaltypes.AttrKeyWeight, strconv.FormatUint(uint64(msg.Weight), 10)),
			sdk.NewAttribute(noorsignaltypes.AttrKeyTimestamp, strconv.FormatInt(ctx.BlockTime().Unix(), 10)),
			sdk.NewAttribute(noorsignaltypes.AttrKeyMetadata, msg.Metadata),
		),
	)

	return &sdk.Result{}, nil
}

// ValidateSignal gère la réception d'un MsgValidateSignal
// (validation d'un signal existant par un curator).
func (s MsgServer) ValidateSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgValidateSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Convertir l'adresse du curator.
	curatorAddr, err := msg.GetCuratorAddress()
	if err != nil {
		return nil, err
	}

	// 2) Vérifier que cette adresse correspond à un Curator actif.
	if !s.Keeper.IsActiveCurator(ctx, curatorAddr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "curator not authorized or not active")
	}

	// 3) Récupérer le signal à valider.
	signal, found := s.Keeper.GetSignal(ctx, msg.SignalId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "signal not found")
	}

	// 4) Vérifier qu'il n'est pas déjà validé.
	if signal.Curator != nil && len(signal.Curator) > 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "signal already validated")
	}

	// 5) Calculer les récompenses PoSS pour ce signal.
	total, part, cur, ok := s.Keeper.ComputeSignalRewardsCurrentEra(ctx, signal.Weight)
	if !ok {
		total = 0
		part = 0
		cur = 0
	}

	// 6) Associer le curator et enregistrer les récompenses.
	signal.Curator = curatorAddr
	signal.TotalReward = total
	signal.RewardParticipant = part
	signal.RewardCurator = cur

	// 7) Mettre à jour le signal dans le store.
	s.Keeper.SetSignal(ctx, signal)

	// 8) Incrémenter le compteur de signaux validés pour ce Curator.
	s.Keeper.IncrementCuratorValidatedCount(ctx, curatorAddr)

	// 9) Émettre l'event poss.signal_validated.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			noorsignaltypes.EventTypeSignalValidated,
			sdk.NewAttribute(noorsignaltypes.AttrKeySignalID, strconv.FormatUint(signal.Id, 10)),
			sdk.NewAttribute(noorsignaltypes.AttrKeyCurator, curatorAddr.String()),
			sdk.NewAttribute(noorsignaltypes.AttrKeyParticipant, signal.Participant.String()),
			sdk.NewAttribute(noorsignaltypes.AttrKeyTotalReward, strconv.FormatUint(total, 10)),
			sdk.NewAttribute(noorsignaltypes.AttrKeyRewardParticipant, strconv.FormatUint(part, 10)),
			sdk.NewAttribute(noorsignaltypes.AttrKeyRewardCurator, strconv.FormatUint(cur, 10)),
			sdk.NewAttribute(noorsignaltypes.AttrKeyWeight, strconv.FormatUint(uint64(signal.Weight), 10)),
			sdk.NewAttribute(noorsignaltypes.AttrKeyTimestamp, strconv.FormatInt(ctx.BlockTime().Unix(), 10)),
		),
	)

	// TODO (plus tard) :
	// - utiliser BankKeeper pour distribuer réellement les récompenses.

	return &sdk.Result{}, nil
}
