package app

import (
	"encoding/json"
	"io"

	dbm "github.com/tendermint/tm-db"
	abci "github.com/tendermint/tendermint/abci/types"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	auth "github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	bank "github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	// Ethermint EVM
	evmmodule "github.com/evmos/ethermint/x/evm"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	// Ethermint FeeMarket
	feemarketmodule "github.com/evmos/ethermint/x/feemarket"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	// NOORSIGNAL (PoSS)
	noorsignal "github.com/noorfinances-eng/noorchain-core/x/noorsignal"
	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignalmodule "github.com/noorfinances-eng/noorchain-core/x/noorsignal"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

const (
	// Keep transient store names unique (cannot match KV store names).
	feeMarketTransientKeyName = "feemarket/transient"
	evmTransientKeyName       = "evm/transient"
)

// NoorchainApp is the minimal Cosmos SDK application for NOORCHAIN.
// Phase 4 — Cosmos core + ParamsKeeper + FeeMarket keeper + EVM keeper + PoSS module.
type NoorchainApp struct {
	*baseapp.BaseApp

	appCodec          codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry
	txConfig          client.TxConfig

	// KV stores
	keys map[string]*storetypes.KVStoreKey
	// Transient stores (used by Params + FeeMarket + EVM)
	tkeys map[string]*storetypes.TransientStoreKey

	// Params
	ParamsKeeper paramskeeper.Keeper

	// Cosmos SDK keepers
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.BaseKeeper
	StakingKeeper stakingkeeper.Keeper

	// Ethermint keepers
	FeeMarketKeeper feemarketkeeper.Keeper
	EvmKeeper       *evmkeeper.Keeper

	// PoSS keeper
	NoorSignalKeeper noorsignalkeeper.Keeper

	mm *module.Manager
}

// NewNoorchainApp creates the base app (no PoSS economic logic yet).
func NewNoorchainApp(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
) *NoorchainApp {
	encCfg := MakeEncodingConfig()

	bApp := baseapp.NewBaseApp("noorchain", logger, db, encCfg.TxConfig.TxDecoder())
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetInterfaceRegistry(encCfg.InterfaceRegistry)

	app := &NoorchainApp{
		BaseApp:           bApp,
		appCodec:          encCfg.Marshaler,
		interfaceRegistry: encCfg.InterfaceRegistry,
		txConfig:          encCfg.TxConfig,
		keys:              make(map[string]*storetypes.KVStoreKey),
		tkeys:             make(map[string]*storetypes.TransientStoreKey),
	}

	// --- Store keys (KV) ---
	app.keys[authtypes.StoreKey] = storetypes.NewKVStoreKey(authtypes.StoreKey)
	app.keys[banktypes.StoreKey] = storetypes.NewKVStoreKey(banktypes.StoreKey)
	app.keys[stakingtypes.StoreKey] = storetypes.NewKVStoreKey(stakingtypes.StoreKey)

	// Gov KV store (préparation GovKeeper / x-gov pour plus tard)
	app.keys[govtypes.StoreKey] = storetypes.NewKVStoreKey(govtypes.StoreKey)

	// Params KV store
	app.keys[paramstypes.StoreKey] = storetypes.NewKVStoreKey(paramstypes.StoreKey)

	// EVM + FeeMarket KV store keys
	app.keys[evmtypes.StoreKey] = storetypes.NewKVStoreKey(evmtypes.StoreKey)
	app.keys[feemarkettypes.StoreKey] = storetypes.NewKVStoreKey(feemarkettypes.StoreKey)

	// PoSS KV store key
	app.keys[noorsignaltypes.StoreKey] = storetypes.NewKVStoreKey(noorsignaltypes.StoreKey)

	// --- Transient store keys ---
	// Params transient store
	app.tkeys[paramstypes.TStoreKey] = storetypes.NewTransientStoreKey(paramstypes.TStoreKey)

	// FeeMarket transient store (MUST have unique name, not "feemarket")
	app.tkeys[feemarkettypes.StoreKey] = storetypes.NewTransientStoreKey(feeMarketTransientKeyName)

	// EVM transient store (MUST have unique name, not "evm")
	app.tkeys[evmtypes.StoreKey] = storetypes.NewTransientStoreKey(evmTransientKeyName)

	// Mount KV stores
	for _, key := range app.keys {
		app.MountStore(key, storetypes.StoreTypeIAVL)
	}
	// Mount transient stores
	for _, tkey := range app.tkeys {
		app.MountStore(tkey, storetypes.StoreTypeTransient)
	}

	// --- ParamsKeeper réel ---
	app.ParamsKeeper = paramskeeper.NewKeeper(
		app.appCodec,
		encCfg.Amino,
		app.keys[paramstypes.StoreKey],
		app.tkeys[paramstypes.TStoreKey],
	)

	// ------------------------------------------------------
	// CRITICAL (Cosmos SDK v0.46):
	// BaseApp needs a ParamStore configured for consensus params
	// ------------------------------------------------------
	consensusParamsSubspace := app.ParamsKeeper.Subspace(baseapp.Paramspace).
		WithKeyTable(paramstypes.ConsensusParamsKeyTable())
	app.BaseApp.SetParamStore(consensusParamsSubspace)

	// Subspaces par module
	authSubspace := app.ParamsKeeper.Subspace(authtypes.ModuleName)
	bankSubspace := app.ParamsKeeper.Subspace(banktypes.ModuleName)
	stakingSubspace := app.ParamsKeeper.Subspace(stakingtypes.ModuleName)

	// Gov subspace (préparé pour le futur GovKeeper)
	govSubspace := app.ParamsKeeper.Subspace(govtypes.ModuleName)
	_ = govSubspace

	// EVM subspace pour le keeper & module EVM
	evmSubspace := app.ParamsKeeper.Subspace(evmtypes.ModuleName)

	// FeeMarket subspace
	feemarketSubspace := app.ParamsKeeper.Subspace(feemarkettypes.ModuleName)

	// PoSS / NoorSignal subspace (pour les Params PoSS, plus tard)
	noorsignalSubspace := app.ParamsKeeper.Subspace(noorsignaltypes.ModuleName)

	// ------------------------------------------------------
	// Module accounts permissions (CRITICAL for staking pools)
	// ------------------------------------------------------
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,

		// Staking pools (required by x/staking)
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},

		// Ethermint modules (safe minimal)
		evmtypes.ModuleName:       nil,
		feemarkettypes.ModuleName: nil,

		// Local PoSS module
		noorsignaltypes.ModuleName: nil,

		// Placeholder gov module address used as authority in this phase
		govtypes.ModuleName: nil,
	}

	// Block external sends to module accounts (safe default).
	blockedAddrs := map[string]bool{}
	for name := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(name).String()] = true
	}

	// --- Base Cosmos keepers ---

	// Accounts
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		app.appCodec,
		app.keys[authtypes.StoreKey],
		authSubspace,
		authtypes.ProtoBaseAccount,
		maccPerms,
		"noorchain", // bech32 prefix / name
	)

	// Bank
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		app.appCodec,
		app.keys[banktypes.StoreKey],
		app.AccountKeeper,
		bankSubspace,
		blockedAddrs,
	)

	// Staking
	app.StakingKeeper = stakingkeeper.NewKeeper(
		app.appCodec,
		app.keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		stakingSubspace,
	)

	// --- Ethermint FeeMarket keeper ---
	feeAuthority := authtypes.NewModuleAddress(govtypes.ModuleName)

	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		app.appCodec,
		feeAuthority,
		app.keys[feemarkettypes.StoreKey],
		app.tkeys[feemarkettypes.StoreKey],
		feemarketSubspace,
	)

	// --- EVM Keeper ---
	evmAuthority := authtypes.NewModuleAddress(evmtypes.ModuleName)

	evmKeeper := evmkeeper.NewKeeper(
		app.appCodec,
		app.keys[evmtypes.StoreKey],
		app.tkeys[evmtypes.StoreKey],
		evmAuthority,
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.FeeMarketKeeper,
		nil, // customPrecompiles
		nil, // evmConstructor
		"",  // tracer
		evmSubspace,
	)

	app.EvmKeeper = evmKeeper

	// --- PoSS / NoorSignal Keeper ---
	noorSignalKeeper := noorsignalkeeper.NewKeeper(
		app.appCodec,
		app.keys[noorsignaltypes.StoreKey],
		noorsignalSubspace,
	)
	app.NoorSignalKeeper = noorSignalKeeper

	// --- AppModules EVM + FeeMarket + PoSS ---
	evmAppModule := evmmodule.NewAppModule(
		app.EvmKeeper,
		app.AccountKeeper,
		evmSubspace,
	)

	feemarketAppModule := feemarketmodule.NewAppModule(
		app.FeeMarketKeeper,
		feemarketSubspace,
	)

	noorsignalAppModule := noorsignalmodule.NewAppModule(
		app.appCodec,
		app.NoorSignalKeeper,
	)

	// --- Module manager ---
	app.mm = module.NewManager(
		auth.NewAppModule(app.appCodec, app.AccountKeeper, nil),
		bank.NewAppModule(app.appCodec, app.BankKeeper, app.AccountKeeper),
		staking.NewAppModule(app.appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		evmAppModule,
		feemarketAppModule,
		noorsignalAppModule,
	)

	app.mm.SetOrderInitGenesis(
		authtypes.ModuleName,
		banktypes.ModuleName,
		stakingtypes.ModuleName,
		evmtypes.ModuleName,
		feemarkettypes.ModuleName,
		noorsignaltypes.ModuleName,
	)

	// Register Msg services + gRPC
	app.mm.RegisterServices(
		module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()),
	)

	// Legacy router — enable PoSS MsgCreateSignal
	app.Router().
		AddRoute(
			sdk.NewRoute(
				noorsignaltypes.ModuleName,
				noorsignal.NewHandler(app.NoorSignalKeeper),
			),
		)

	// ABCI handlers
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// AnteHandler (EVM/cosmos)
	app.SetupAnteHandler()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			panic(err)
		}
	}

	return app
}

// NewApp is a thin wrapper used by cmd/noord.
func NewApp(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) *NoorchainApp {
	return NewNoorchainApp(logger, db, traceStore, true, appOpts)
}

// InitChainer is called on chain initialization.
func (app *NoorchainApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState map[string]json.RawMessage

	if len(req.AppStateBytes) > 0 {
		if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
			panic(err)
		}
	} else {
		genesisState = make(map[string]json.RawMessage)
	}

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// BeginBlocker is called at the beginning of every block.
func (app *NoorchainApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker is called at the end of every block.
func (app *NoorchainApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// --- Encoding config minimal ---

type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	// IMPORTANT:
	// Register all module interfaces BEFORE the app registers services.
	// This prevents: "type_url /cosmos.bank.v1beta1.MsgSend has not been registered yet"
	ModuleBasics.RegisterInterfaces(interfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(amino)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	txCfg := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         cdc,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}
