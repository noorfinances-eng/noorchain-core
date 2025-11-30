package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetupAnteHandler définit pour l'instant un ante handler minimal (no-op).
// Il sera remplacé plus tard par un AnteHandler complet compatible EVM
// (Ethermint v0.22.0 + Cosmos SDK v0.46.11).
//
// Objectif de ante evm 1 :
// - Avoir une méthode propre sur NoorchainApp,
// - ZÉRO logique métier,
// - Compilation garantie.
func (app *NoorchainApp) SetupAnteHandler() {
	if app == nil {
		return
	}

	// AnteHandler minimal : ne fait rien, accepte tout.
	// On le remplacera par la vraie chaîne de decorators (signature checks,
	// fees, EVM-specific checks, etc.) dans les blocs suivants.
	app.SetAnteHandler(func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
		return ctx, nil
	})
}
