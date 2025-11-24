package app

// StartNOORChain is the high-level Cosmos entrypoint for Phase 2.
// It builds the Cosmos app, wraps it into a CosmosServer, and starts it.
// In Phase 2 this remains structural only.
func StartNOORChain() error {
	// Build the Cosmos-ready application skeleton
	noorApp := BuildCosmosApp()

	// Build the Cosmos server wrapper
	server := BuildCosmosServer(noorApp)

	// Start the placeholder Cosmos server
	return StartCosmosServer(server)
}
