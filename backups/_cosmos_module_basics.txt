package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	auth "github.com/cosmos/cosmos-sdk/x/auth"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	gov "github.com/cosmos/cosmos-sdk/x/gov"
	staking "github.com/cosmos/cosmos-sdk/x/staking"

	evm "github.com/evmos/ethermint/x/evm"
	feemarket "github.com/evmos/ethermint/x/feemarket"
)

// ModuleBasics defines the basic modules (AppModuleBasic) enabled in NOORCHAIN Phase 2.
var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	bank.AppModuleBasic{},
	staking.AppModuleBasic{},
	gov.AppModuleBasic{},
	evm.AppModuleBasic{},
	feemarket.AppModuleBasic{},
)
