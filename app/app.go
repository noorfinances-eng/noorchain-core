package app

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
)

// App est le type principal de l'application NOORCHAIN.
//
// Il encapsule :
// - BaseApp : le cœur Cosmos SDK
// - Name / Version : métadonnées simples
// - Keepers : structure des gestionnaires de modules (remplie plus tard)
// - Encoding : configuration d'encodage (codecs, TxConfig, etc.)
type App struct {
	*baseapp.BaseApp

	Name    string
	Version string

	Keepers  AppKeepers
	Encoding EncodingConfig
}

// NewNoorchainApp crée une instance très simple de l'application NOORCHAIN.
//
// Cette version est encore "minimaliste" :
// - BaseApp = nil
// - Keepers = structure vide
// - Encoding = MakeEncodingConfig() (codec standard Cosmos)
func NewNoorchainApp() *App {
	encCfg := MakeEncodingConfig()

	return &App{
		BaseApp: nil,
		Name:    "NOORCHAIN",
		Version: "0.0.1-dev",
		Keepers: AppKeepers{},
		Encoding: encCfg,
	}
}

// Start est une méthode placeholder qui, pour l'instant,
// se contente d'afficher un message.
//
// Plus tard, cette méthode pourra démarrer le node complet
// (ABCI, CometBFT, services, etc.).
func (a *App) Start() error {
	fmt.Printf("%s node (version %s) starting...\n", a.Name, a.Version)
	fmt.Println("Cosmos SDK + Ethermint wiring will be added in the next technical phases.")
	return nil
}
