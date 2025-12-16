package app

import (
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

// RegisterAPIRoutes is required by the Cosmos SDK server interface.
// Phase 8.A: minimal implementation (no REST, no gRPC, no routes).
func (app *NoorchainApp) RegisterAPIRoutes(
	_ *api.Server,
	_ serverconfig.APIConfig,
) {
	// no-op
}
