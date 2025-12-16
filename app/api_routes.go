package app

import (
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	tmservice "github.com/cosmos/cosmos-sdk/server/grpc/tmservice"
)

// RegisterAPIRoutes is required by the Cosmos SDK server interface.
// Phase 8.A: minimal implementation (no REST, no routes).
func (app *NoorchainApp) RegisterAPIRoutes(
	_ *api.Server,
	_ serverconfig.APIConfig,
) {
	// no-op
}

// RegisterTendermintService is required by Cosmos SDK v0.46.
// Phase 8.A: minimal implementation (no gRPC exposure).
func (app *NoorchainApp) RegisterTendermintService(
	_ tmservice.ServiceServer,
) {
	// no-op
}
