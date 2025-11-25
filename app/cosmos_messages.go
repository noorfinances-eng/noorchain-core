package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
)

// RegisterMsgServiceRouter registers a minimal msg service router for NOORCHAIN.
// Phase 2: router is empty but required by Cosmos SDK and Ethermint.
func (app *NOORChainApp) RegisterMsgServiceRouter() {
	router := baseapp.NewMsgServiceRouter()

	// No routes yet (Phase 2 placeholder)
	// Example later: bank.RegisterMsgServer(router, bankKeeper)

	app.BaseApp.SetMsgServiceRouter(router)
}
