package app

import (
	"io"

	dbm "github.com/tendermint/tm-db"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	auth "github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	bank "github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	// Ethermint EVM & FeeMarket
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

// NoorchainApp is the minimal Cosmos SDK application for NOORCHAIN.
// Phase 4 — Cosmos core + ParamsKeeper + EVM store/keeper + FeeMarket keeper (+ Gov store prep).
type NoorchainApp struct {
	*baseapp.BaseApp

	appCodec          codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry

	// KV stores
	keys map[string]*storetypes.KVStoreKey
	// Transient stores (used by Params + EVM + FeeMarket)
	tkeys map[string]*storetypes.TransientStoreKey

	// Params
	ParamsKeeper paramskeeper.Keeper

	// Cosmos SDK keepers
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.BaseKeeper
	StakingKeeper stakingkeeper.Keeper

	// Ethermint keepers
	EvmKeeper       *evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper

	mm *module.Manager
}

// NewNoorchainApp creates the base app (no ante-handler / begin / end block logic yet).
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
		keys:              make(map[string]*storetypes.KVStoreKey),
		tkeys:             make(map[string]*storetypes.TransientStoreKey),
	}

	// --- Store keys (KV) ---

	// Core Cosmos
	app.keys[authtypes.StoreKey] = storetypes.NewKVStoreKey(authtypes.StoreKey)
	app.keys[banktypes.StoreKey] = storetypes.NewKVStoreKey(banktypes.StoreKey)
	app.keys[stakingtypes.StoreKey] = storetypes.NewKVStoreKey(stakingtypes.StoreKey)

	// Gov KV store (préparation x/gov pour plus tard)
	app.keys[govtypes.StoreKey] = storetypes.NewKVStoreKey(govtypes.StoreKey)

	// Params KV store
	app.keys[paramstypes.StoreKey] = storetypes.NewKVStoreKey(paramstypes.StoreKey)

	// EVM + FeeMarket KV store keys
	app.keys[evmtypes.StoreKey] = storetypes.NewKVStoreKey(evmtypes.StoreKey)
	app.keys[feemarkettypes.StoreKey] = storetypes.NewKVStoreKey(feemarkettypes.StoreKey)

	// --- Transient store keys ---

	// Params transient store
	app.tkeys[paramstypes.TStoreKey] = storetypes.NewTransientStoreKey(paramstypes.TStoreKey)

	// EVM transient store (pour bloom, tx index, logs, gas used)
	app.tkeys[evmtypes.StoreKey] = storetypes.NewTransientStoreKey(evmtypes.StoreKey)

	// FeeMarket transient store (uses same name as module store key)
	app.tkeys[feemarkettypes.StoreKey] = storetypes.NewTransientStoreKey(feemarkettypes.StoreKey)

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

	// Subspaces par module
	authSubspace := app.ParamsKeeper.Subspace(authtypes.ModuleName)
	bankSubspace := app.ParamsKeeper.Subspace(banktypes.ModuleName)
	stakingSubspace := app.ParamsKeeper.Subspace(stakingtypes.ModuleName)
	// Gov subspace (pour plus tard, GovKeeper complet si besoin)
	_ = app.ParamsKeeper.Subspace(govtypes.ModuleName)
	// EVM subspace (utilisé par EVM keeper)
	evmSubspace := app.ParamsKeeper.Subspace(evmtypes.ModuleName)
	// FeeMarket subspace
	feemarketSubspace := app.ParamsKeeper.Subspace(feemarkettypes.ModuleName)

	// --- Base Cosmos keepers ---

	// Accounts
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		app.appCodec,
		app.keys[authtypes.StoreKey],
		authSubspace,
		authtypes.ProtoBaseAccount,
		map[string][]string{},
		"noorchain", // bech32 prefix / name
	)

	// Bank
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		app.appCodec,
		app.keys[banktypes.StoreKey],
		app.AccountKeeper,
		bankSubspace,
		map[string]bool{},
	)

	// Staking
	app.StakingKeeper = stakingkeeper.NewKeeper(
		app.appCodec,
		app.keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		stakingSubspace,
	)

	// --- Ethermint FeeMarket keeper (avec vrai subspace params) ---

	feeAuthority := authtypes.NewModuleAddress("gov")

	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		app.appCodec,
		feeAuthority,
		app.keys[feemarkettypes.StoreKey],
		app.tkeys[feemarkettypes.StoreKey],
		feemarketSubspace,
	)

	// --- Ethermint EVM keeper (version minimale, sans custom precompiles) ---

	// NOTE :
	// - On passe `nil` pour customPrecompiles et evmConstructor (valeurs neutres).
	// - On met un tracer vide "" (pas de debug EVM pour l’instant).
	evmAuthority := authtypes.NewModuleAddress("gov")

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

	// ⚠️ Ici la correction : evmKeeper est déjà un *Keeper
	app.EvmKeeper = evmKeeper

	// --- Module manager (toujours minimal : auth + bank + staking) ---
	// On n’a pas encore branché les AppModules EVM / FeeMarket / Gov.
	app.mm = module.NewManager(
		auth.NewAppModule(app.appCodec, app.AccountKeeper, nil),
		bank.NewAppModule(app.appCodec, app.BankKeeper, app.AccountKeeper),
		staking.NewAppModule(app.appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
	)

	app.mm.SetOrderInitGenesis(
		authtypes.ModuleName,
		banktypes.ModuleName,
		stakingtypes.ModuleName,
	)

	app.mm.RegisterServices(
		module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()),
	)

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
	cdc := codec.NewProtoCodec(interfaceRegistry)
	txCfg := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         cdc,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}
