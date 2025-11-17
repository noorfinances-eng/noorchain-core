package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewNoorchainAppWithCosmos defines the *future* real Cosmos SDK application
// constructor for NOORCHAIN.
//
// IMPORTANT:
// - This is a **skeleton only** at this stage.
// - It is NOT yet used by cmd/noord/main.go.
// - It exists to define the structure and parameters we will use later.
//
// In the next technical phases, this function will be expanded to:
// - create a real BaseApp
// - configure encoding
// - register all modules and keepers
// - wire Ethermint (EVM) and the PoSS module.
func NewNoorchainAppWithCosmos(
	logger sdk.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts interface{},
) *App {
	// Placeholder: in the future we will build a real BaseApp here using:
	//
	// baseapp.NewBaseApp(
	//     AppName,
	//     logger,
	//     db,
	//     txDecoder,
	//     baseapp.SetChainID(ChainID),
	//     ...
	// )
	//
	// For now, we simply create the same placeholder App as in NewNoorchainApp(),
	// but we keep this separate so that we can progressively migrate the code
	// without breaking the minimal node skeleton.

	return &App{
		BaseApp: nil, // will be set to a real BaseApp later
		Name:    "NOORCHAIN",
		Version: "0.0.1-dev",
	}
}
