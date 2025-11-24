package app

// CosmosServer is a placeholder for the NOORCHAIN Cosmos/CometBFT server.
// In Phase 2 this structure remains minimal until real ABCI wiring is added.
type CosmosServer struct {
	App *NOORChainApp
}

// NewCosmosServer creates a new CosmosServer placeholder from a NOORChainApp instance.
func NewCosmosServer(app *NOORChainApp) CosmosServer {
	return CosmosServer{
		App: app,
	}
}

// StartCosmosServer is a placeholder for starting the Cosmos/CometBFT server.
// In Phase 2 it does nothing and returns nil.
func StartCosmosServer(server CosmosServer) error {
	_ = server
	return nil
}
