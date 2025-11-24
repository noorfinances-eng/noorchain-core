package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"cosmossdk.io/log"
)

// NOORChainApp is the root structure for the real Cosmos SDK application.
// In Phase 2, this is only a placeholder base containing BaseApp + Keys.
type NOORChainApp struct {
	*baseapp.BaseApp

	Keys CosmosKeys
}

// NewNOORChainApp constructs the skeleton of a Cosmos SDK application.
// In Phase 2, it initializes only BaseApp and basic store keys.
func NewNOORChainApp() *NOORChainApp {
	// Minimal logger for Phase 2 (stdout)
	logger := log.NewNopLogger()

	// Placeholder BaseApp with logger
	bApp := baseapp.NewBaseApp(
		AppName,
		logger,
		nil, // DB (added later)
		nil, // TxDecoder (added later with encoding)
	)

	keys := NewCosmosKeys()

	return &NOORChainApp{
		BaseApp: bApp,
		Keys:    keys,
	}
}
