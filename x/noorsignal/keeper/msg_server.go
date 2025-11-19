package keeper

import (
	"context"

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

	// 6) Créer et stocker le signal via le Keeper.
	_ = s.Keeper.CreateSignal(ctx, signal)

	// 7) Incrémenter le compteur quotidien si une limite est active.
	if cfg.MaxSignalsPerDay > 0 {
		s.incrementDailySignalCount(ctx, participantAddr, dayBucket)
	}

	return &sdk.Result{}, nil
}

// ValidateSignal gère la réception d'un MsgValidateSignal
// (validation d'un signal existant par un curator).
//
// Étapes :
// - convertir l'adresse du curator
// - récupérer le signal
// - vérifier qu'il existe et n'est pas déjà validé
// - calculer les récompenses PoSS (total / 70% / 30%)
// - enregistrer les récompenses dans le signal
// - associer le curator et sauvegarder.
//
// AUCUN transfert de NUR n'est encore réalisé ici. Les montants sont
// simplement stockés comme "récompenses théoriques PoSS".
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

	// 2) Récupérer le signal à valider.
	signal, found := s.Keeper.GetSignal(ctx, msg.SignalId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "signal not found")
	}

	// 3) Vérifier qu'il n'est pas déjà validé.
	if signal.Curator != nil && len(signal.Curator) > 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "signal already validated")
	}

	// 4) Calculer les récompenses PoSS pour ce signal.
	total, part, cur, ok := s.Keeper.ComputeSignalRewardsCurrentEra(ctx, signal.Weight)
	if !ok {
		// Si aucune config n'est présente, on peut soit :
		// - refuser la validation
		// - soit considérer que la récompense est nulle.
		// Ici, on choisit de considérer total = 0.
		total = 0
		part = 0
		cur = 0
	}

	// 5) Associer le curator et enregistrer les récompenses.
	signal.Curator = curatorAddr
	signal.TotalReward = total
	signal.RewardParticipant = part
	signal.RewardCurator = cur

	// 6) Mettre à jour le signal dans le store.
	s.Keeper.SetSignal(ctx, signal)

	// TODO (plus tard) :
	// - vérifier que le curator est autorisé (stockage Curator)
	// - utiliser BankKeeper pour distribuer réellement les récompenses
	// - générer des events pour les explorateurs.

	return &sdk.Result{}, nil
}
