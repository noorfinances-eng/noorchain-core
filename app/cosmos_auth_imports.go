package app

// Phase 2: Import placeholders for the Cosmos SDK auth module.
// We do NOT wire anything yet. We only prepare the structure,
// required before adding real keepers and module manager entries.

import (
	authmodule "github.com/cosmos/cosmos-sdk/x/auth/module"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// Placeholder references to avoid unused import errors in Phase 2.
var (
	_ = authmodule.AppModuleBasic{}
	_ = auth.BaseAccount{}
)
