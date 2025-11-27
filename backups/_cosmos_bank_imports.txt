package app

import (
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// CosmosBankImports keeps references to x/bank types to ensure
// the module is pulled by Go modules in Phase 2.
var (
	_ = bank.AppModuleBasic{}
	_ = banktypes.ModuleName
)
