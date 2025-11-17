package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewNoorchainAppWithCosmos defines the future real Cosmos SDK application
// constructor for NOORCHAIN.
func NewNoorchainAppWithCosmos(
	logger sdk.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts interface{},
) *App {

	// --- 1) Create encoding configuration (still empty skeleton)
	encCfg := MakeEncodingConfig()
	_ = encCfg

	// --- 2) Prepare BaseApp variable (not initialized yet)
	// In future steps:
	// base = baseapp.NewBaseApp(...)
	var base *baseapp.BaseApp = nil

	// --- 3) Return a partially prepared App
	return &App{
		BaseApp: base,   // nil for now
		Name:    "NOORCHAIN",
		Version: "0.0.1-dev",
	}
}
