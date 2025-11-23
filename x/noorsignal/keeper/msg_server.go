package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// MsgServer est le point d'entrée pour les transactions (Msg)
// du module PoSS (noorsignal).
type MsgServer struct {
	Keeper Keeper
}

// NewMsgServer construit un MsgServer basé uniquement sur le Keeper PoSS.
//
// NOTE : Dans cette V1, on ne branche PAS encore de BankKeeper.
//        Les transferts réels de NUR seront ajoutés plus tard.
func NewMsgServer(k Keeper) MsgServer {
	return MsgServer{
		Keeper: k,
	}
}

// ---------------------------------------------------------
// SubmitSignal : un participant émet un signal PoSS
// ---------------------------------------------------------
func (s MsgServer) SubmitSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSubmitSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Validation basique (poids, adresse…)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	participantAddr, err := msg.GetParticipantAddress()
	if err != nil {
		return nil, err
	}

	// 2) Récupérer la config PoSS actuelle (ou défaut)
	cfg, found := s.Keeper.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	// 3) Vérifier la limite quotidienne si activée
	var dayBucket uint64
	if cfg.MaxSignalsPerDay > 0 {
		ts := ctx.BlockTime().Unix()
		if ts < 0 {
			ts = 0
		}
		dayBucket = uint64(ts) / 86400

		current := s.Keeper.GetDailySignalCount(ctx, participantAddr, dayBucket)
		if current >= cfg.MaxSignalsPerDay {
			return nil, errors.New("daily signal limit reached")
		}
	}

	// 4) Construire le signal (sans rewards, sans curator)
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

	// 5) Enregistrer le signal (ID auto-incrémenté)
	signal = s.Keeper.CreateSignal(ctx, signal)

	// 6) Incrémenter le compteur quotidien
	if cfg.MaxSignalsPerDay > 0 {
		s.Keeper.IncrementDailySignalCount(ctx, participantAddr, dayBucket)
	}

	// 7) Émettre un event poss.signal_submitted
	ctx.EventManager().EmitEvent(
		noorsignaltypes.NewSignalSubmittedEvent(signal, ctx.BlockHeight()),
	)

	return &sdk.Result{}, nil
}

// ---------------------------------------------------------
// ValidateSignal : un curator valide un signal
// (V1 : pas encore de BankKeeper.SendCoins)
// ---------------------------------------------------------
func (s MsgServer) ValidateSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgValidateSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Validation basique
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	curatorAddr, err := msg.GetCuratorAddress()
	if err != nil {
		return nil, err
	}

	// 2) Vérifier que le curator est actif
	if !s.Keeper.IsActiveCurator(ctx, curatorAddr) {
		return nil, errors.New("curator not active or not authorized")
	}

	// 3) Charger le signal
	signal, found := s.Keeper.GetSignal(ctx, msg.SignalId)
	if !found {
		return nil, errors.New("signal not found")
	}

	// 4) Vérifier qu'il n'est pas déjà validé
	if signal.Curator != nil && len(signal.Curator) > 0 {
		return nil, errors.New("signal already validated")
	}

	// 5) Calcul des rewards PoSS (local, sans transfert réel)
	total, part, cur, ok := s.Keeper.ComputeSignalRewardsCurrentEra(ctx, signal.Weight)
	if !ok {
		total = 0
		part = 0
		cur = 0
	}

	// 6) Mettre à jour le signal
	signal.Curator = curatorAddr
	signal.TotalReward = total
	signal.RewardParticipant = part
	signal.RewardCurator = cur

	s.Keeper.SetSignal(ctx, signal)

	// 7) Incrémenter le compteur de validations du curator
	s.Keeper.IncrementCuratorValidatedCount(ctx, curatorAddr)

	// 8) Émettre un event poss.signal_validated
	ctx.EventManager().EmitEvent(
		noorsignaltypes.NewSignalValidatedEvent(signal, ctx.BlockHeight()),
	)

	// 9) FUTUR : ici on utilisera BankKeeper.SendCoins pour
	// distribuer les rewards depuis la réserve PoSS.

	return &sdk.Result{}, nil
}
