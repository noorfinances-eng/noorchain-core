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
//
// For now, it only prepares the structure and returns a nil BaseApp.
// In later phases, BuildBaseApp will construct a real baseapp.BaseApp
// with all modules wired.
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

// BuildBaseApp will, in future phases, create and configure a real
// baseapp.BaseApp instance using Cosmos SDK and Ethermint.
//
// For now, it only returns nil as a placeholder so that the structure
// compiles and can be filled step by step.
func (b *AppBuilder) BuildBaseApp() *baseapp.BaseApp {
	return nil
}
