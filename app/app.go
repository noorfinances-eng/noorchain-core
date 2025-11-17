package app

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
)

// App is the main NOORCHAIN application type.
//
// For now, it is a very simple skeleton that only holds:
// - a reference to a Cosmos SDK BaseApp (currently nil, will be wired later)
// - a name
// - a version
//
// In the next technical phases, this struct will be extended to embed
// all required Cosmos SDK keepers, module managers, and configuration.
type App struct {
	*baseapp.BaseApp

	Name    string
	Version string
}

// NewNoorchainApp creates a new placeholder instance of the NOORCHAIN app.
//
// IMPORTANT:
// - At this stage, BaseApp is still nil.
// - Later, this function will be rewritten to fully initialize a Cosmos SDK
//   application (with logger, DB, encoding, modules, etc.).
func NewNoorchainApp() *App {
	return &App{
		BaseApp: nil,
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
