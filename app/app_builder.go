package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AppBuilder is a helper struct that will progressively build the full
// NOORCHAIN Cosmos SDK application.
//
// It centralises:
// - logger
// - database handle
// - tracing writer
// - "load latest" flag
// - app options (config)
// - encoding configuration
type AppBuilder struct {
	logger     sdk.Logger
	db         dbm.DB
	traceStore io.Writer
	loadLatest bool
	appOpts    interface{}
	encCfg     EncodingConfig
}

// NewAppBuilder creates a new AppBuilder instance using the given
// Cosmos-style constructor parameters.
func NewAppBuilder(
	logger sdk.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts interface{},
) *AppBuilder {
	encCfg := MakeEncodingConfig()

	return &AppBuilder{
		logger:     logger,
		db:         db,
		traceStore: traceStore,
		loadLatest: loadLatest,
		appOpts:    appOpts,
		encCfg:     encCfg,
	}
}

// BuildBaseApp creates a minimal baseapp.BaseApp instance.
//
// For now, this is still a very early version:
// - it uses the encoding config skeleton (TxConfig may be nil)
// - it sets the chain ID
// - it optionally loads the latest version from the DB
//
// Later, this method will be extended to:
// - use a real TxDecoder from encCfg.TxConfig
// - set all necessary BaseApp options
// - integrate Ethermint (EVM) and other modules.
func (b *AppBuilder) BuildBaseApp() *baseapp.BaseApp {
	// Derive the transaction decoder from the encoding config if available.
	var txDecoder sdk.TxDecoder
	if b.encCfg.TxConfig != nil {
		txDecoder = b.encCfg.TxConfig.TxDecoder()
	} else {
		// For now, we allow a nil decoder; this will be replaced later when
		// the encoding configuration is fully wired.
		txDecoder = nil
	}

	// Create a minimal BaseApp instance.
	base := baseapp.NewBaseApp(
		AppName,
		b.logger,
		b.db,
		txDecoder,
		baseapp.SetChainID(ChainID),
	)

	// Optionally load the latest version from the DB.
	// NOTE:
	// - In a future iteration, we will return an error instead of panicking.
	if b.loadLatest && base != nil {
		if err := base.LoadLatestVersion(); err != nil {
			panic(err)
		}
	}

	return base
}
