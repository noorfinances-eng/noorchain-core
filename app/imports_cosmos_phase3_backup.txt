package app

// This file exists only to force Go modules to keep required Cosmos SDK
// packages in the project, avoiding "unused import" cleanup.
// It does NOT execute any logic.

import (
	_ "github.com/cosmos/cosmos-sdk/baseapp"
	_ "github.com/cosmos/cosmos-sdk/client"
	_ "github.com/cosmos/cosmos-sdk/codec"
	_ "github.com/cosmos/cosmos-sdk/server"
	_ "github.com/cosmos/cosmos-sdk/store"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/module"

	// Basic Cosmos modules
	_ "github.com/cosmos/cosmos-sdk/x/auth"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx"

	_ "github.com/cosmos/cosmos-sdk/x/bank"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	_ "github.com/cosmos/cosmos-sdk/x/staking"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	_ "github.com/cosmos/cosmos-sdk/x/mint"
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	_ "github.com/cosmos/cosmos-sdk/x/genutil"
	genutilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	// Helpers
	_ "github.com/cosmos/cosmos-sdk/x/params"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	// Tendermint (CometBFT)
	_ "github.com/cometbft/cometbft/abci/types"
	_ "github.com/cometbft/cometbft/types"
)

// Avoid unused warnings for imported types (no functional usage).
var (
	_ = authTypes.ModuleName
	_ = bankTypes.ModuleName
	_ = stakingTypes.ModuleName
	_ = mintTypes.ModuleName
	_ = genutilTypes.ModuleName
	_ = paramsTypes.ModuleName
)
