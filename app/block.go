package app

import (
	abci "github.com/tendermint/tendermint/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker sera branché au ModuleManager plus tard (Phase 3+).
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) {
	// Pour l'instant : rien. On ajoutera la logique plus tard.
}

// EndBlocker retournera les mises à jour des validateurs plus tard.
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	// Pour l'instant : aucun changement de validateurs.
	return []abci.ValidatorUpdate{}
}

// InitChainer sera utilisé pour initialiser le genesis (state initial).
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	// Pour l’instant : on renvoie une réponse vide.
	return abci.ResponseInitChain{}
}
