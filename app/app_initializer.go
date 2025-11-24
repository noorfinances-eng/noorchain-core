package app

// InitApp initializes the real Cosmos application structure for NOORCHAIN.
// In Phase 2 this delegates to BuildCosmosApp, which remains a placeholder
// until module wiring is implemented.
func InitApp() *NOORChainApp {
	return BuildCosmosApp()
}
