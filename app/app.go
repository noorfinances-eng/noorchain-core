package app

import (
    "github.com/cosmos/cosmos-sdk/baseapp"
    sdk "github.com/cosmos/cosmos-sdk/types"
    storetypes "github.com/cosmos/cosmos-sdk/store/types"
    "github.com/cosmos/cosmos-sdk/types/module"
)

// App is the root application for NOORCHAIN.
// For Phase 2: minimal scaffolding, no modules wired yet.
type App struct {
    BaseApp        *baseapp.BaseApp
    ModuleBasics   module.BasicManager
    EncodingConfig EncodingConfig
}

// New returns a new NOORCHAIN App instance.
func New(logger baseapp.Logger, db storetypes.CommitMultiStore) *App {
    // 1) Configure Bech32 (noor1...)
    ConfigureBech32Prefixes()

    // 2) Encoding (Proto codecs + Amino)
    enc := MakeEncodingConfig()

    // 3) Create BaseApp
    bApp := baseapp.NewBaseApp(
        "noorchain",
        logger,
        db,
        enc.TxConfig.TxDecoder(),
    )

    // 4) Create the full App struct
    app := &App{
        BaseApp:        bApp,
        ModuleBasics:   ModuleBasics,
        EncodingConfig: enc,
    }

    // NOTE: In Phase 3+ we will add:
    // - Keepers
    // - ModuleManager
    // - Routes
    // - PoSS Module
    // - AnteHandler
    // - Upgrade handlers

    return app
}
