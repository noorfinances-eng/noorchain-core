package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	authmodule "github.com/cosmos/cosmos-sdk/x/auth/module"
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank/module"
	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking/module"
)

// ModuleBasics defines the basic module elements for NOORCHAIN.
// Phase 2: includes AUTH + BANK + STAKING modules.
var ModuleBasics = module.BasicManager(
	authmodule.AppModuleBasic{},
	bankmodule.AppModuleBasic{},
	stakingmodule.AppModuleBasic{},
)
