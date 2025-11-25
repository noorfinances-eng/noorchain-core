package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"cosmossdk.io/log"
	"github.com/cometbft/cometdb/memdb"
)

// NOORChainApp is the root structure for the real Cosmos SDK application.
// In Phase 2, this is only a placeholder base containing BaseApp + structures.
type NOORChainApp struct {
	*baseapp.BaseApp

	Keys          CosmosKeys
	Keepers       CosmosKeepers
	ModuleManager CosmosModuleManager
	StoreLoader   CosmosStoreLoader
	Encoding      CosmosEncodingConfig
}

// NewNOORChainApp constructs the skeleton of a Cosmos SDK application.
// In Phase 2, it initializes BaseApp, logger, MemDB, TxDecoder placeholder,
// mounts the basic store keys, and prepares the encoding config.
func NewNOORChainApp() *NOORChainApp {
	// Minimal logger for Phase 2
	logger := log.NewNopLogger()

	// In-memory database for Phase 2 (no persistent storage yet)
	db := memdb.NewDB()

	// Placeholder BaseApp with logger, MemDB, and TxDecoder
	bApp := baseapp.NewBaseApp(
		AppName,
		logger,
		db,
		CosmosTxDecoder,
	)

	keys := NewCosmosKeys()

	// Mount basic stores into the BaseApp
	LoadStores(bApp, keys)

	// Build encoding configuration
	encodingBuilder := NewCosmosEncodingBuilder()
	encoding := encodingBuilder.Build()

	return &NOORChainApp{
		BaseApp:       bApp,
		Keys:          keys,
		Keepers:       NewCosmosKeepers(),
		ModuleManager: NewCosmosModuleManager(),
		StoreLoader:   NewCosmosStoreLoader(),
		Encoding:      encoding,
	}
}
