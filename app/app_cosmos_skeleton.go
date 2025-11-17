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
// At this stage, this function only:
// - receives the classical Cosmos SDK constructor parameters
// - creates the encoding config
// - returns a minimal App with BaseApp still nil
//
// In the next phases, we will:
// - create baseapp.NewBaseApp()
// - wire encoding, keepers, module manager
// - integrate Ethermint (EVM)
// - integrate PoSS
func NewNoorchainAppWithCosmos(
	logger sdk.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts interface{},
) *App {

	// --- 1) Create encoding configuration (currently empty skeleton)
	encCfg := MakeEncodingConfig()
	_ = encCfg // placeholder usage to avoid unused variable

	// --- 2) Placeholder: in future we will create BaseApp like:
	// base := baseapp.NewBaseApp(
	//     AppName,
	//     logger,
	//     db,
	//     encCfg.TxConfig.TxDecoder(),
	//     baseapp.SetChainID(ChainID),
	// )
	//
	// And assign: BaseApp: base,

	return &App{
		BaseApp: nil,                 // will be replaced later
		Name:    "NOORCHAIN",
		Version: "0.0.1-dev",
	}
}
