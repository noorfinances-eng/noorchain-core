package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

// RegisterAPIRoutes is required by Cosmos SDK v0.46.
// Phase 8.A: minimal, no REST routes exposed.
func (app *NoorchainApp) RegisterAPIRoutes(
	_ *api.Server,
	_ serverconfig.APIConfig,
) {
	// no-op
}

// RegisterTendermintService is required by Cosmos SDK v0.46.
// Signature uses client.Context.
func (app *NoorchainApp) RegisterTendermintService(
	_ client.Context,
) {
	// no-op
}

// RegisterTxService is required by Cosmos SDK v0.46.
// Phase 8.A: minimal, do not expose tx service.
func (app *NoorchainApp) RegisterTxService(
	_ client.Context,
) {
	// no-op
}
