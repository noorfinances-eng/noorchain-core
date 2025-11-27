package app

import (
    "github.com/cosmos/cosmos-sdk/types/module"

    // Modules
    auth "github.com/cosmos/cosmos-sdk/x/auth"
    authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"

    bank "github.com/cosmos/cosmos-sdk/x/bank"
    bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"

    staking "github.com/cosmos/cosmos-sdk/x/staking"
    stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"

    mint "github.com/cosmos/cosmos-sdk/x/mint"
    mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"

    genutil "github.com/cosmos/cosmos-sdk/x/genutil"
    genutilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

    params "github.com/cosmos/cosmos-sdk/x/params"
    paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ModuleBasics defines all modules installed at the base level.
var ModuleBasics = module.NewBasicManager(
    auth.AppModuleBasic{},
    bank.AppModuleBasic{},
    staking.AppModuleBasic{},
    mint.AppModuleBasic{},
    genutil.AppModuleBasic{},
    params.AppModuleBasic{},
)

// Fake references prevent unused imports
var (
    _ = authTypes.ModuleName
    _ = bankTypes.ModuleName
    _ = stakingTypes.ModuleName
    _ = mintTypes.ModuleName
    _ = genutilTypes.ModuleName
    _ = paramsTypes.ModuleName
)

// Expose externally if needed
func GetModuleBasics() module.BasicManager {
    return ModuleBasics
}
