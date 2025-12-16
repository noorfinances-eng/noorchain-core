package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/api"
)

// RegisterAPIRoutes is required by the Cosmos SDK server interface.
// For Phase 8.A (public testnet proof), we keep it minimal: no REST API wiring yet.
func (app *NoorchainApp) RegisterAPIRoutes(_ *api.Server, _ client.Context) {
	// no-op (minimal start)
}
