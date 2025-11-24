package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"cosmossdk.io/log"
	"github.com/cometbft/cometdb/memdb"
)

// NOORChainApp is the root structure for the real Cosmos SDK application.
// In Phase 2, this is only a placeholder base containing BaseApp + Keys.
type NOORChainApp struct {
	*baseapp.BaseApp

	Keys CosmosKeys
}

// NewNOORChainApp constructs the skeleton of a Cosmos SDK application.
// In Phase 2, it initializes only BaseApp, a minimal logger, and a MemDB.
func NewNOORChainApp() *NOORChainApp {
	// Minimal logger for Phase 2
	logger := log.NewNopLogger()

	// In-memory database for Phase 2 (no persistent storage yet)
	db := memdb.NewDB()

	// Placeholder BaseApp with logger and MemDB
	bApp := baseapp.NewBaseApp(
		AppName,
		logger,
		db,  // <-- Phase 2 DB
		nil, // TxDecoder (added later with encoding)
	)

	keys := NewCosmosKeys()

	return &NOORChainApp{
		BaseApp: bApp,
		Keys:    keys,
	}
}
