package app

import (
	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// CosmosStakingImports keeps references to x/staking types to ensure
// the module is pulled by Go modules in Phase 2.
var (
	_ = staking.AppModuleBasic{}
	_ = stakingtypes.ModuleName
)
