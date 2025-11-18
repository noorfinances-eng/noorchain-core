package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewNoorchainAppWithCosmos defines the future real Cosmos SDK application
// constructor for NOORCHAIN.
//
// It now uses the AppBuilder helper to prepare the application. The builder
// will progressively be extended to create a real BaseApp and wire all modules.
func NewNoorchainAppWithCosmos(
	logger sdk.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts interface{},
) *App {

	// 1) Create an AppBuilder with all Cosmos-style constructor parameters.
	builder := NewAppBuilder(
		logger,
		db,
		traceStore,
		loadLatest,
		appOpts,
	)

	// 2) Build the BaseApp using the builder.
	//    For now, BuildBaseApp() still returns nil as a placeholder.
	var base *baseapp.BaseApp = builder.BuildBaseApp()

	// 3) Return the NOORCHAIN App instance.
	return &App{
		BaseApp: base, // will be non-nil once BuildBaseApp is fully implemented
		Name:    "NOORCHAIN",
		Version: "0.0.1-dev",
	}
}
