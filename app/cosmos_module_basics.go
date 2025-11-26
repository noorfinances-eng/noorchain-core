package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	authmodule "github.com/cosmos/cosmos-sdk/x/auth/module"
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank/module"
)

// ModuleBasics defines the basic module elements for NOORCHAIN.
// Phase 2: now includes AUTH + BANK modules.
var ModuleBasics = module.BasicManager(
	authmodule.AppModuleBasic{},
	bankmodule.AppModuleBasic{},
)
