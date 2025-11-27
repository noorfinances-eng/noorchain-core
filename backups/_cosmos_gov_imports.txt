package app

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// CosmosGovImports keeps references to x/gov types to ensure
// the module is pulled by Go modules in Phase 2.
var (
	_ = gov.AppModuleBasic{}
	_ = govtypes.ModuleName
)
