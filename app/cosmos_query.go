package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
)

// RegisterQueryRouter registers a minimal query router for NOORCHAIN.
// Phase 2: router is empty but required by Cosmos SDK / Ethermint.
func (app *NOORChainApp) RegisterQueryRouter() {
	queryRouter := baseapp.NewQueryRouter()

	// No real routes yet (Phase 2)
	// Example later: queryRouter.AddRoute("bank", someHandler)

	app.BaseApp.SetQueryRouter(queryRouter)
}
