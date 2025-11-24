package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
)

// CosmosBaseApp is an alias to the Cosmos SDK BaseApp.
// In Phase 2 this is our entry point for wiring a real Cosmos application.
type CosmosBaseApp = baseapp.BaseApp
