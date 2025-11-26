package app

// Phase 2: Import placeholders for the Cosmos SDK bank module.
// We do NOT wire anything yet. We only prepare the structure,
// required before adding real keepers and module manager entries.

import (
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank/module"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// Placeholder references to avoid unused import errors in Phase 2.
var (
	_ = bankmodule.AppModuleBasic{}
	_ = bank.MsgSend{}
)
