package main

import (
	"fmt"
	"os"

	"github.com/tendermint/tendermint/libs/log"

	dbm "github.com/tendermint/tm-db"

	noorapp "github.com/noorfinances-eng/noorchain-core/app"
)

func main() {
	// Logger: simple stdout logger
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// Database: in-memory for Phase 2
	db := dbm.NewMemDB()

	// Create NOORCHAIN App
	app := noorapp.New(logger, db)

	// Placeholder for Phase 2
	fmt.Println("NOORCHAIN node placeholder (Phase 2 - basic App created)")
	fmt.Println("App loaded:", app.BaseApp.Name())

	// NOTE:
	// In Phase 3+ we will add:
	// - CLI commands
	// - Tendermint server startup
	// - Genesis loading
	// - PoSS module wiring
}
