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
	"github.com/cosmos/cosmos-sdk/server"
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

	genutil "github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	"github.com/cosmos/cosmos-sdk/std"

	// Crypto interface registrations
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	evmmodule "github.com/evmos/ethermint/x/evm"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	feemarketmodule "github.com/evmos/ethermint/x/feemarket"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	noorsignal "github.com/noorfinances-eng/noorchain-core/x/noorsignal"
	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"

	"github.com/spf13/cast"
)

const (
	feeMarketTransientKeyName = "feemarket/transient"
	evmTransientKeyName       = "evm/transient"
)

type NoorchainApp struct {
	*baseapp.BaseApp

	appCodec          codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry
	txConfig          client.TxConfig

	keys  map[string]*storetypes.KVStoreKey
	tkeys map[string]*storetypes.TransientStoreKey

	ParamsKeeper paramskeeper.Keeper

	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.BaseKeeper
	StakingKeeper stakingkeeper.Keeper

	FeeMarketKeeper feemarketkeeper.Keeper
	EvmKeeper       *evmkeeper.Keeper

	NoorSignalKeeper noorsignalkeeper.Keeper

	mm *module.Manager

	// --- CRITICAL: keep a DecCoins copy for InitChainer(ctx.WithMinGasPrices)
	minGasPricesStr string
	minGasPricesDec sdk.DecCoins
}

func NewNoorchainApp(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
) *NoorchainApp {
	encCfg := MakeEncodingConfig()

	// -------------------------------------------------------------------
	// CRITICAL (Cosmos SDK v0.46 + Ethermint v0.22):
	// We keep minGasPrices in BOTH string and DecCoins form.
	// The AnteHandler reads from ctx.MinGasPrices() during DeliverGenTxs.
	// -------------------------------------------------------------------
	minGasPrices := ""
	if appOpts != nil {
		// server.FlagMinGasPrices == "minimum-gas-prices"
		minGasPrices = cast.ToString(appOpts.Get(server.FlagMinGasPrices))
	}

	// Fallback must be non-empty.
	// NOTE: you can also set "0unur;0aphoton" later, once you want both denoms.
	if minGasPrices == "" {
		minGasPrices = "0.000000001unur"
	}

	minGasDec, err := sdk.ParseDecCoins(minGasPrices)
	if err != nil {
		panic(err)
	}
	// -------------------------------------------------------------------

	bApp := baseapp.NewBaseApp(
		"noorchain",
		logger,
		db,
		encCfg.TxConfig.TxDecoder(),
		// Keep it: server may not have injected min gas prices into ctx during DeliverGenTxs.
		baseapp.SetMinGasPrices(minGasPrices),
	)

	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetInterfaceRegistry(encCfg.InterfaceRegistry)

	app := &NoorchainApp{
		BaseApp:           bApp,
		appCodec:          encCfg.Marshaler,
		interfaceRegistry: encCfg.InterfaceRegistry,
		txConfig:          encCfg.TxConfig,
		keys:              make(map[string]*storetypes.KVStoreKey),
		tkeys:             make(map[string]*storetypes.TransientStoreKey),

		minGasPricesStr: minGasPrices,
		minGasPricesDec: minGasDec,
	}

	app.keys[authtypes.StoreKey] = storetypes.NewKVStoreKey(authtypes.StoreKey)
	app.keys[banktypes.StoreKey] = storetypes.NewKVStoreKey(banktypes.StoreKey)
	app.keys[stakingtypes.StoreKey] = storetypes.NewKVStoreKey(stakingtypes.StoreKey)
	app.keys[govtypes.StoreKey] = storetypes.NewKVStoreKey(govtypes.StoreKey)
	app.keys[paramstypes.StoreKey] = storetypes.NewKVStoreKey(paramstypes.StoreKey)
	app.keys[evmtypes.StoreKey] = storetypes.NewKVStoreKey(evmtypes.StoreKey)
	app.keys[feemarkettypes.StoreKey] = storetypes.NewKVStoreKey(feemarkettypes.StoreKey)
	app.keys[noorsignaltypes.StoreKey] = storetypes.NewKVStoreKey(noorsignaltypes.StoreKey)

	app.tkeys[paramstypes.TStoreKey] = storetypes.NewTransientStoreKey(paramstypes.TStoreKey)
	app.tkeys[feemarkettypes.StoreKey] = storetypes.NewTransientStoreKey(feeMarketTransientKeyName)
	app.tkeys[evmtypes.StoreKey] = storetypes.NewTransientStoreKey(evmTransientKeyName)

	for _, key := range app.keys {
		app.MountStore(key, storetypes.StoreTypeIAVL)
	}
	for _, tkey := range app.tkeys {
		app.MountStore(tkey, storetypes.StoreTypeTransient)
	}

	app.ParamsKeeper = paramskeeper.NewKeeper(
		app.appCodec,
		encCfg.Amino,
		app.keys[paramstypes.StoreKey],
		app.tkeys[paramstypes.TStoreKey],
	)

	consensusParamsSubspace := app.ParamsKeeper.Subspace(baseapp.Paramspace).
		WithKeyTable(paramstypes.ConsensusParamsKeyTable())
	app.BaseApp.SetParamStore(consensusParamsSubspace)

	authSubspace := app.ParamsKeeper.Subspace(authtypes.ModuleName)
	bankSubspace := app.ParamsKeeper.Subspace(banktypes.ModuleName)
	stakingSubspace := app.ParamsKeeper.Subspace(stakingtypes.ModuleName)
	_ = app.ParamsKeeper.Subspace(govtypes.ModuleName)

	evmSubspace := app.ParamsKeeper.Subspace(evmtypes.ModuleName)
	feemarketSubspace := app.ParamsKeeper.Subspace(feemarkettypes.ModuleName)
	noorsignalSubspace := app.ParamsKeeper.Subspace(noorsignaltypes.ModuleName)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		evmtypes.ModuleName:            nil,
		feemarkettypes.ModuleName:      nil,
		noorsignaltypes.ModuleName:     nil,
		govtypes.ModuleName:            nil,
	}

	blockedAddrs := map[string]bool{}
	for name := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(name).String()] = true
	}

	app.AccountKeeper = authkeeper.NewAccountKeeper(
		app.appCodec,
		app.keys[authtypes.StoreKey],
		authSubspace,
		authtypes.ProtoBaseAccount,
		maccPerms,
		"noorchain",
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		app.appCodec,
		app.keys[banktypes.StoreKey],
		app.AccountKeeper,
		bankSubspace,
		blockedAddrs,
	)

	app.StakingKeeper = stakingkeeper.NewKeeper(
		app.appCodec,
		app.keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		stakingSubspace,
	)

	feeAuthority := authtypes.NewModuleAddress(govtypes.ModuleName)
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		app.appCodec,
		feeAuthority,
		app.keys[feemarkettypes.StoreKey],
		app.tkeys[feemarkettypes.StoreKey],
		feemarketSubspace,
	)

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
		nil,
		nil,
		"",
		evmSubspace,
	)
	app.EvmKeeper = evmKeeper

	noorSignalKeeper := noorsignalkeeper.NewKeeper(
		app.appCodec,
		app.keys[noorsignaltypes.StoreKey],
		noorsignalSubspace,
	)
	app.NoorSignalKeeper = noorSignalKeeper

	evmAppModule := evmmodule.NewAppModule(app.EvmKeeper, app.AccountKeeper, evmSubspace)
	feemarketAppModule := feemarketmodule.NewAppModule(app.FeeMarketKeeper, feemarketSubspace)
	noorsignalAppModule := noorsignal.NewAppModule(app.appCodec, app.NoorSignalKeeper)

	genutilAppModule := genutil.NewAppModule(
		app.AccountKeeper,
		app.StakingKeeper,
		app.BaseApp.DeliverTx,
		app.txConfig,
	)

	app.mm = module.NewManager(
		auth.NewAppModule(app.appCodec, app.AccountKeeper, nil),
		bank.NewAppModule(app.appCodec, app.BankKeeper, app.AccountKeeper),
		staking.NewAppModule(app.appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		genutilAppModule,
		evmAppModule,
		feemarketAppModule,
		noorsignalAppModule,
	)

	app.mm.SetOrderInitGenesis(
		authtypes.ModuleName,
		banktypes.ModuleName,
		stakingtypes.ModuleName,
		feemarkettypes.ModuleName,
		genutiltypes.ModuleName,
		evmtypes.ModuleName,
		noorsignaltypes.ModuleName,
	)

	app.mm.RegisterServices(
		module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()),
	)

	app.Router().AddRoute(
		sdk.NewRoute(
			noorsignaltypes.ModuleName,
			noorsignal.NewHandler(app.NoorSignalKeeper),
		),
	)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	app.SetupAnteHandler()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			panic(err)
		}
	}

	return app
}

func NewApp(logger tmlog.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) *NoorchainApp {
	return NewNoorchainApp(logger, db, traceStore, true, appOpts)
}

func (app *NoorchainApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	// Keep it: useful for module init and any ctx-dependent reads.
	ctx = ctx.WithMinGasPrices(app.minGasPricesDec)

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

func (app *NoorchainApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *NoorchainApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	std.RegisterLegacyAminoCodec(amino)
	std.RegisterInterfaces(interfaceRegistry)

	// Crypto: register PubKey interfaces correctly (required for gentx decoding / start)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations(
		(*cryptotypes.PubKey)(nil),
		&ed25519.PubKey{},
		&secp256k1.PubKey{},
	)

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
