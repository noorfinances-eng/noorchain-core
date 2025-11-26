package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	authmodule "github.com/cosmos/cosmos-sdk/x/auth/module"
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank/module"
	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking/module"
	govmodule "github.com/cosmos/cosmos-sdk/x/gov/module"
	evmmodule "github.com/evmos/ethermint/x/evm/module"
	feemarketmodule "github.com/evmos/ethermint/x/feemarket/module"
)

// ModuleBasics defines the basic module elements for NOORCHAIN.
// Phase 2: includes AUTH + BANK + STAKING + GOV + EVM + FEEMARKET modules.
var ModuleBasics = module.BasicManager(
	authmodule.AppModuleBasic{},
	bankmodule.AppModuleBasic{},
	stakingmodule.AppModuleBasic{},
	govmodule.AppModuleBasic{},
	evmmodule.AppModuleBasic{},
	feemarketmodule.AppModuleBasic{},
)
