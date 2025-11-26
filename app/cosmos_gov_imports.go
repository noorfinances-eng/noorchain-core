package app

// Phase 2: Import placeholders for the Cosmos SDK gov module.
// We do NOT wire anything yet. We only prepare the structure,
// required before adding real keepers and module manager entries.

import (
	govmodule "github.com/cosmos/cosmos-sdk/x/gov/module"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Placeholder references to avoid unused import errors in Phase 2.
var (
	_ = govmodule.AppModuleBasic{}
	_ = gov.MsgVote{}
)
