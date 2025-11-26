package app

import (
	evm "github.com/evmos/ethermint/x/evm"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarket "github.com/evmos/ethermint/x/feemarket"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

// CosmosEVMImports keeps references to Ethermint modules to ensure
// they are pulled by Go modules in Phase 2.
var (
	_ = evm.AppModuleBasic{}
	_ = evmtypes.ModuleName

	_ = feemarket.AppModuleBasic{}
	_ = feemarkettypes.ModuleName
)
