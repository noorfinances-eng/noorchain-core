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

	// Ethermint store types (Bloc A: only keys, no keepers yet)
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

// NoorchainApp is the minimal Cosmos SDK application for NOORCHAIN.
// Phase 4 (impl1 + Bloc A EVM): auth / bank / staking + EVM/feemarket store keys.
// Pas encore de Gov, pas encore de keepers EVM/feemarket.
type NoorchainApp struct {
	*baseapp.BaseApp

	appCodec          codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry

	keys map[string]*storetypes.KVStoreKey

	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.BaseKeeper
	StakingKeeper stakingkeeper.Keeper

	mm *module.Manager
}

// NewNoorchainApp creates the base Cosmos SDK app (no EVM logic yet).
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
	}

	// --- Store keys (KV) ---
	app.keys[authtypes.StoreKey] = storetypes.NewKVStoreKey(authtypes.StoreKey)
	app.keys[banktypes.StoreKey] = storetypes.NewKVStoreKey(banktypes.StoreKey)
	app.keys[stakingtypes.StoreKey] = storetypes.NewKVStoreKey(stakingtypes.StoreKey)

	// Bloc A: ajouter les stores EVM + FeeMarket (KV uniquement pour l'instant)
	app.keys[evmtypes.StoreKey] = storetypes.NewKVStoreKey(evmtypes.StoreKey)
	app.keys[feemarkettypes.StoreKey] = storetypes.NewKVStoreKey(feemarkettypes.StoreKey)

	for _, key := range app.keys {
		app.MountStore(key, storetypes.StoreTypeIAVL)
	}

	// --- Fake / empty params subspace (pour satisfaire la signature) ---
	// On branchera le vrai params keeper plus tard.
	emptySubspace := paramstypes.Subspace{}

	// --- Keepers de base ---

	// Accounts
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		app.appCodec,
		app.keys[authtypes.StoreKey],
		emptySubspace,
		authtypes.ProtoBaseAccount,
		map[string][]string{},
		"noorchain", // bech32 prefix / name
	)

	// Bank
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		app.appCodec,
		app.keys[banktypes.StoreKey],
		app.AccountKeeper,
		emptySubspace,
		map[string]bool{},
	)

	// Staking
	app.StakingKeeper = stakingkeeper.NewKeeper(
		app.appCodec,
		app.keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		emptySubspace,
	)

	// --- Module manager minimal (auth + bank + staking) ---
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
