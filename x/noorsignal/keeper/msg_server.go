package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// MsgServer est le point d'entrée pour les transactions (Msg)
// du module PoSS (noorsignal).
//
// À ce stade, il implémente une logique simple pour :
// - l'émission de signaux (SubmitSignal)
// - la validation de signaux (ValidateSignal, sans récompenses pour l'instant).
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
	// 1) Récupérer le sdk.Context à partir du context gRPC.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 2) Valider le poids du signal.
	if msg.Weight == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "weight must be >= 1")
	}
	if msg.Weight > 100 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "weight must be <= 100")
	}

	// 3) Convertir l'adresse du participant (Bech32 -> sdk.AccAddress).
	participantAddr, err := msg.GetParticipantAddress()
	if err != nil {
		return nil, err
	}

	// 4) Construire un Signal de base (sans curator pour l'instant).
	signal := noorsignaltypes.Signal{
		Participant: participantAddr,
		Curator:     nil,
		Weight:      msg.Weight,
		Time:        ctx.BlockTime(),
		Metadata:    msg.Metadata,
	}

	// 5) Créer et stocker le signal via le Keeper.
	_ = s.Keeper.CreateSignal(ctx, signal)

	// 6) Retourner un sdk.Result simple (sans events pour l'instant).
	return &sdk.Result{}, nil
}

// ValidateSignal gère la réception d'un MsgValidateSignal
// (validation d'un signal existant par un curator).
//
// Étapes actuelles :
// - récupérer le contexte
// - convertir l'adresse du curator
// - lire le signal via son ID
// - vérifier qu'il existe
// - vérifier qu'il n'est pas déjà validé
// - enregistrer l'adresse du curator sur le signal
//
// TODO (plus tard) :
// - vérifier que l'adresse est bien un curator autorisé
// - attribuer les récompenses 70% / 30% via BankKeeper
//   en utilisant ComputeSignalRewardsCurrentEra
// - générer des événements pour l'exploration / indexation.
func (s MsgServer) ValidateSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgValidateSignal,
) (*sdk.Result, error) {
	// 1) Récupérer le sdk.Context.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 2) Convertir l'adresse du curator (Bech32 -> sdk.AccAddress).
	curatorAddr, err := msg.GetCuratorAddress()
	if err != nil {
		return nil, err
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

	// 5) Associer le curator au signal.
	signal.Curator = curatorAddr

	// 6) Mettre à jour le signal dans le store.
	s.Keeper.SetSignal(ctx, signal)

	// 7) Retourner un sdk.Result simple (sans events ni récompenses pour l'instant).
	return &sdk.Result{}, nil
}
