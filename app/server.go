package app

// StartServer is a placeholder for the NOORCHAIN ABCI server launcher.
// It initializes the Cosmos-based NOORChainApp and delegates to BuildServer.
func StartServer() error {
	noorApp := InitApp()
	return BuildServer(noorApp)
}
