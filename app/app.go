package app

import "fmt"

// App is a placeholder for the NOORCHAIN application.
//
// In the next steps, this struct will be extended to embed the Cosmos SDK
// base application and the Ethermint (EVM) modules.
type App struct {
    Name    string
    Version string
}

// NewNoorchainApp creates a new placeholder instance of the NOORCHAIN app.
//
// Later, this function will take many parameters (logger, database, encoding,
// app options, etc.) but for now we keep it minimal so that the project
// structure is clear and buildable.
func NewNoorchainApp() *App {
    return &App{
        Name:    "NOORCHAIN",
        Version: "0.0.1-dev",
    }
}

// Start is a placeholder method that will later start the full node logic.
//
// For now, it just prints a message. In future steps, this will be replaced
// by proper Cosmos SDK + Ethermint wiring and ABCI server startup.
func (a *App) Start() error {
    fmt.Printf("%s node (version %s) starting...\n", a.Name, a.Version)
    fmt.Println("Cosmos SDK + Ethermint wiring will be added in the next technical phases.")
    return nil
}
