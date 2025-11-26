package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	authmodule "github.com/cosmos/cosmos-sdk/x/auth/module"
)

// ModuleBasics defines the basic module elements for NOORCHAIN.
// Phase 2: now includes the AUTH module (first real Cosmos module).
var ModuleBasics = module.BasicManager(
	authmodule.AppModuleBasic{},
)
