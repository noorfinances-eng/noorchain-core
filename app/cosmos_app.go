package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// AnteHandler returns a minimal AnteHandler for NOORCHAIN.
// In Phase 2 this is a placeholder that always accepts the transaction.
func (app *NOORChainApp) AnteHandler() sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
		_ = tx
		_ = simulate
		return ctx, nil
	}
}

// NewNOORChainApp constructs the skeleton of a Cosmos SDK application.
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

	// Construct the application
	app := &NOORChainApp{
		BaseApp:       bApp,
		Keys:          keys,
		Keepers:       NewCosmosKeepers(),
		ModuleManager: NewCosmosModuleManager(),
		StoreLoader:   NewCosmosStoreLoader(),
		Encoding:      encoding,
	}

	// Register query router (Phase 2 placeholder)
	app.RegisterQueryRouter()

	// Register AnteHandler (Phase 2 placeholder)
	app.SetAnteHandler(app.AnteHandler())

	// Register ABCI handlers (Phase 2 placeholders)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetInitChainer(app.InitChainer)

	return app
}
