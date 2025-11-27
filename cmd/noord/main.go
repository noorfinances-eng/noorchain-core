package main

import (
	"fmt"

	noorapp "github.com/noorfinances-eng/noorchain-core/app"
)

func main() {
	// Création de l'application minimale NOORCHAIN (Phase 2)
	app := noorapp.NewApp()
	info := app.Info

	// Affichage d'un message placeholder propre
	fmt.Println("NOORCHAIN node placeholder (Phase 2 – skeleton only)")
	fmt.Printf("App: %s v%s\n", info.Name, info.Version)
}
