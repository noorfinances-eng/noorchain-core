package app

import (
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// CosmosAuthImports keeps references to x/auth types to ensure
// the module is pulled by Go modules in Phase 2.
var (
	_ = auth.AppModuleBasic{}
	_ = authtypes.ModuleName
)
