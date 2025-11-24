package app

// BuildCosmosServer constructs a CosmosServer from the NOORChainApp.
// In Phase 2 this is only structural and performs no real server setup.
func BuildCosmosServer(app *NOORChainApp) CosmosServer {
	return NewCosmosServer(app)
}
