package app

// App is a placeholder for the NOORCHAIN Cosmos/Ethermint application.
// Later we will integrate Cosmos SDK modules and Ethermint here.
type App struct {
	// TODO: add Cosmos SDK baseapp and keepers
}

// NewNoorchainApp will later initialize the full NOORCHAIN app.
// For now it just returns an empty placeholder App.
func NewNoorchainApp() *App {
	return &App{}
}
