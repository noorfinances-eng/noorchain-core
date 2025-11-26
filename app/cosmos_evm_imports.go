package app

// Phase 2: Import placeholders for Ethermint EVM and FeeMarket modules.
// We do NOT wire anything yet. This only prepares structure.

import (
	evmmodule "github.com/evmos/ethermint/x/evm/module"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	feemarketmodule "github.com/evmos/ethermint/x/feemarket/module"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

// Placeholder references to avoid unused import errors in Phase 2.
var (
	_ = evmmodule.AppModuleBasic{}
	_ = evmtypes.MsgEthereumTx{}

	_ = feemarketmodule.AppModuleBasic{}
	_ = feemarkettypes.MsgUpdateParams{}
)
