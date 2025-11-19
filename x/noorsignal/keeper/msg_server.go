package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// MsgServer est le point d'entrée pour les transactions (Msg)
// du module PoSS (noorsignal).
//
// À ce stade, il ne contient que des squelettes de méthodes.
// La logique métier (création de signaux, validation, récompenses)
// sera ajoutée dans des phases ultérieures.
type MsgServer struct {
	Keeper
}

// NewMsgServer construit un MsgServer à partir d'un Keeper PoSS.
func NewMsgServer(k Keeper) MsgServer {
	return MsgServer{Keeper: k}
}

// SubmitSignal gérera plus tard la réception d'un MsgSubmitSignal
// (émission d'un nouveau signal social PoSS).
//
// Pour l'instant, la fonction est un simple squelette.
func (s MsgServer) SubmitSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSubmitSignal,
) (*sdk.Result, error) {
	// TODO: implémenter la logique de création de signal PoSS :
	// - valider les champs
	// - appliquer les limites (MaxSignalsPerDay)
	// - enregistrer le signal dans le store
	// - éventuellement générer des événements
	// - déclencher (ou préparer) les récompenses PoSS
	return &sdk.Result{}, nil
}

// ValidateSignal gérera plus tard la réception d'un MsgValidateSignal
// (validation d'un signal existant par un curator).
//
// Pour l'instant, la fonction est un simple squelette.
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
