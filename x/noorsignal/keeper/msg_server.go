package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// MsgServer est le point d'entrée pour les transactions (Msg)
// du module PoSS (noorsignal).
//
// À ce stade, il commence à implémenter une logique simple pour
// l'émission de signaux (SubmitSignal). La validation, les limites
// et les récompenses seront ajoutées progressivement.
type MsgServer struct {
	Keeper
}

// NewMsgServer construit un MsgServer à partir d'un Keeper PoSS.
func NewMsgServer(k Keeper) MsgServer {
	return MsgServer{Keeper: k}
}

// SubmitSignal gère la réception d'un MsgSubmitSignal
// (émission d'un nouveau signal social PoSS).
//
// Étapes actuelles :
// - conversion de l'adresse du participant
// - création d'une struct Signal (sans curator)
// - assignation d'un ID auto-incrémenté via CreateSignal
// - stockage du signal dans le KVStore
//
// TODO (plus tard) :
// - valider le poids (Weight)
// - appliquer les limites quotidiennes (MaxSignalsPerDay)
// - éventuellement déclencher le calcul des récompenses PoSS
//   et les transferts de NUR via BankKeeper.
func (s MsgServer) SubmitSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSubmitSignal,
) (*sdk.Result, error) {
	// 1) Récupérer le sdk.Context à partir du context gRPC.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 2) Convertir l'adresse du participant (Bech32 -> sdk.AccAddress).
	participantAddr, err := msg.GetParticipantAddress()
	if err != nil {
		return nil, err
	}

	// 3) Construire un Signal de base (sans curator pour l'instant).
	signal := noorsignaltypes.Signal{
		// Id sera rempli par CreateSignal.
		Participant: participantAddr,
		// Curator vide à ce stade (sera renseigné lors de la validation).
		Curator:  nil,
		Weight:   msg.Weight,
		Time:     ctx.BlockTime(),
		Metadata: msg.Metadata,
	}

	// 4) Créer et stocker le signal via le Keeper.
	// CreateSignal attribue un Id et l'enregistre.
	_ = s.Keeper.CreateSignal(ctx, signal)

	// 5) Retourner un sdk.Result simple (sans events pour l'instant).
	// Plus tard, on pourra ajouter des événements (events) pour
	// faciliter l'indexation et l'exploration.
	return &sdk.Result{}, nil
}

// ValidateSignal gérera plus tard la réception d'un MsgValidateSignal
// (validation d'un signal existant par un curator).
//
// Pour l'instant, la fonction reste un squelette.
func (s MsgServer) ValidateSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgValidateSignal,
) (*sdk.Result, error) {
	// TODO: implémenter la logique de validation de signal :
	// - vérifier le curator
	// - vérifier l'existence du signal
	// - marquer le signal comme validé
	// - attribuer les récompenses 70% / 30%
	// - générer des événements
	return &sdk.Result{}, nil
}
