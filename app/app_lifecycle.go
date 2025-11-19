package app

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker est appelé au début de chaque bloc par BaseApp.
//
// Il délègue l'appel au ModuleManager, qui lui-même appellera
// BeginBlock sur chaque module enregistré (dont le module PoSS
// noorsignal).
func (app *App) BeginBlocker(
	ctx sdk.Context,
	req abci.RequestBeginBlock,
) abci.ResponseBeginBlock {
	if app.Modules.Manager != nil {
		app.Modules.Manager.BeginBlock(ctx, req)
	}

	// Pour l'instant, aucun ValidatorUpdate n'est renvoyé par les modules.
	return abci.ResponseBeginBlock{}
}

// EndBlocker est appelé à la fin de chaque bloc par BaseApp.
//
// Il délègue l'appel au ModuleManager, puis renvoie les éventuels
// ValidatorUpdates produits par les modules (si un jour NOORCHAIN
// gère un système de validateurs).
func (app *App) EndBlocker(
	ctx sdk.Context,
	req abci.RequestEndBlock,
) abci.ResponseEndBlock {
	var updates []abci.ValidatorUpdate

	if app.Modules.Manager != nil {
		updates = app.Modules.Manager.EndBlock(ctx, req)
	}

	return abci.ResponseEndBlock{
		ValidatorUpdates: updates,
	}
}
