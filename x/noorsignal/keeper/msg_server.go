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

// -----------------------------------------------------------------------------
// MsgSubmitSignal : émission d'un nouveau signal social PoSS
// -----------------------------------------------------------------------------

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

	// 6) Créer et stocker le signal via le Keeper (récupère l'ID).
	signal = s.Keeper.CreateSignal(ctx, signal)

	// 7) Incrémenter le compteur quotidien si une limite est active.
	if cfg.MaxSignalsPerDay > 0 {
		s.incrementDailySignalCount(ctx, participantAddr, dayBucket)
	}

	// 8) Émettre un event poss.signal_submitted.
	ctx.EventManager().EmitEvent(
		noorsignaltypes.NewSignalSubmittedEvent(signal, ctx.BlockHeight()),
	)

	return &sdk.Result{}, nil
}

// -----------------------------------------------------------------------------
// MsgValidateSignal : validation + récompenses réelles PoSS
// -----------------------------------------------------------------------------

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

	// -----------------------------------------------------------------
	// 6) Distribuer les récompenses réelles NUR (unur) si total > 0.
	// -----------------------------------------------------------------
	if total > 0 {
		// Source : réserve PoSS (80 %) définie dans types/addresses.go
		reserveAddr := noorsignaltypes.TestPoSSReserveAddr

		denom := "unur" // doit être aligné avec app.CoinDenom

		// 6.a) Récompense pour le participant (70 %)
		if part > 0 {
			coinPart := sdk.NewCoin(denom, sdk.NewIntFromUint64(part))
			coinsPart := sdk.NewCoins(coinPart)

			if err := s.BankKeeper.SendCoins(ctx, reserveAddr, signal.Participant, coinsPart); err != nil {
				return nil, sdkerrors.Wrap(err, "failed to send reward to participant")
			}
		}

		// 6.b) Récompense pour le curator (30 %)
		if cur > 0 {
			coinCur := sdk.NewCoin(denom, sdk.NewIntFromUint64(cur))
			coinsCur := sdk.NewCoins(coinCur)

			if err := s.BankKeeper.SendCoins(ctx, reserveAddr, curatorAddr, coinsCur); err != nil {
				return nil, sdkerrors.Wrap(err, "failed to send reward to curator")
			}
		}
	}

	// 7) Associer le curator et enregistrer les récompenses dans le signal.
	signal.Curator = curatorAddr
	signal.TotalReward = total
	signal.RewardParticipant = part
	signal.RewardCurator = cur

	// 8) Mettre à jour le signal dans le store.
	s.Keeper.SetSignal(ctx, signal)

	// 9) Incrémenter le compteur de signaux validés pour ce Curator.
	s.Keeper.IncrementCuratorValidatedCount(ctx, curatorAddr)

	// 10) Émettre un event poss.signal_validated.
	ctx.EventManager().EmitEvent(
		noorsignaltypes.NewSignalValidatedEvent(signal, ctx.BlockHeight()),
	)

	return &sdk.Result{}, nil
}

// -----------------------------------------------------------------------------
// MsgAddCurator / MsgRemoveCurator / MsgSetConfig (Admin)
// -----------------------------------------------------------------------------

// AddCurator gère la réception d'un MsgAddCurator.
// Il permet à une "authority" d'ajouter ou de réactiver un Curator
// avec un certain niveau (BRONZE / SILVER / GOLD).
func (s MsgServer) AddCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgAddCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Vérifier que l'authority est présente.
	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	// 2) Convertir l'adresse du curator.
	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid curator address")
	}

	// 3) Récupérer un éventuel Curator existant.
	curator, found := s.Keeper.GetCurator(ctx, curatorAddr)
	if !found {
		// Nouveau Curator.
		curator = noorsignaltypes.Curator{
			Address:               curatorAddr,
			Level:                 msg.Level,
			TotalSignalsValidated: 0,
			Active:                true,
		}
	} else {
		// Mise à jour d'un Curator existant.
		curator.Level = msg.Level
		curator.Active = true
	}

	// 4) Enregistrer le Curator dans le store.
	s.Keeper.SetCurator(ctx, curator)

	return &sdk.Result{}, nil
}

// RemoveCurator gère la réception d'un MsgRemoveCurator.
// Il permet à une "authority" de désactiver un Curator existant.
func (s MsgServer) RemoveCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgRemoveCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Vérifier que l'authority est présente.
	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	// 2) Convertir l'adresse du curator.
	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid curator address")
	}

	// 3) Récupérer le Curator.
	curator, found := s.Keeper.GetCurator(ctx, curatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "curator not found")
	}

	// 4) Le désactiver (on garde l'historique).
	curator.Active = false

	// 5) Enregistrer la mise à jour.
	s.Keeper.SetCurator(ctx, curator)

	return &sdk.Result{}, nil
}

// SetConfig gère la réception d'un MsgSetConfig.
// Il permet à une "authority" de mettre à jour la configuration PoSS
// (base_reward, max_signals_per_day, era_index, ratios 70/30, etc.)
// sans toucher au genesis.
func (s MsgServer) SetConfig(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSetConfig,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Vérifier que l'authority est présente.
	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	// 2) Parser la base de récompense depuis la string.
	baseReward, err := strconv.ParseUint(msg.BaseReward, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid base_reward")
	}

	// 3) Vérifier les ratios participant / curator.
	totalRatio := msg.ParticipantRatio + msg.CuratorRatio
	if totalRatio == 0 || totalRatio != 100 {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			"participant_ratio + curator_ratio must be exactly 100",
		)
	}

	// 4) Construire une nouvelle configuration PoSS.
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

	// 5) Enregistrer cette nouvelle config dans le store PoSS.
	s.Keeper.SetConfig(ctx, newCfg)

	return &sdk.Result{}, nil
}
