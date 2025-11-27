package main

import (
	"fmt"

	noorapp "github.com/noorfinances-eng/noorchain-core/app"
)

func main() {
	// Crée une instance minimale de l'application NOORCHAIN
	app := noorapp.NewApp()

	// Récupère les métadonnées (nom, version)
	info := app.Info

	// Placeholder Phase 2
	fmt.Println("NOORCHAIN node placeholder (Phase 2 – skeleton only)")
	fmt.Printf("App Name: %s\n", info.Name)
	fmt.Printf("App Version: %s\n", info.Version)

	// NOTE :
	// - En Phase 3+, on ajoutera :
	//   * Cosmos SDK
	//   * BaseApp + stores
	//   * Tendermint / CometBFT node
	//   * CLI, genesis, PoSS, etc.
}
