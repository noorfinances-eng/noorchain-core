package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	auth "github.com/cosmos/cosmos-sdk/x/auth"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	staking "github.com/cosmos/cosmos-sdk/x/staking"

	evm "github.com/evmos/ethermint/x/evm"
	feemarket "github.com/evmos/ethermint/x/feemarket"

	noorsignal "github.com/noorfinances-eng/noorchain-core/x/noorsignal"
)

// ModuleBasics defines the module BasicManager used by the CLI (init, gentx, etc.).
// Minimal set for NOORCHAIN Public Testnet.
var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	bank.AppModuleBasic{},
	staking.AppModuleBasic{},

	evm.AppModuleBasic{},
	feemarket.AppModuleBasic{},

	noorsignal.AppModuleBasic{},
)
