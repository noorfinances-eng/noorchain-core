package app

import (
	"context"
	"encoding/json"
	"io"

	dbm "github.com/cosmos/cosmos-db"

	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	cmtservice "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	node "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	authmodule "github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	bankmodule "github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	genutilmodule "github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	paramsmodule "github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cast"

	sdklog "cosmossdk.io/log"
	"cosmossdk.io/store"
	pruningtypes "cosmossdk.io/store/pruning/types"
	storetypes "cosmossdk.io/store/types"

	sdkruntime "github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

const (
	AppName = "noorchain"
)

// Module account permissions
var maccPerms = map[string][]string{
	authtypes.FeeCollectorName:     nil,
	stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
	stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
}

// Blocked addresses (none blocked in this minimal profile)
func BlockedModuleAccountAddrs() map[string]bool {
	addrs := map[string]bool{}
	return addrs
}

// In-memory consensus param store to satisfy BaseApp
type NoopConsensusParamStore struct {
	cp *tmproto.ConsensusParams
}

func (ps *NoopConsensusParamStore) Has(ctx context.Context) (bool, error) {
	return ps.cp != nil, nil
}

func (ps *NoopConsensusParamStore) Get(ctx context.Context) (tmproto.ConsensusParams, error) {
	if ps.cp == nil {
		return tmproto.ConsensusParams{}, nil
	}
	return *ps.cp, nil
}

func (ps *NoopConsensusParamStore) Set(ctx context.Context, cp tmproto.ConsensusParams) error {
	ps.cp = &cp
	return nil
}

type NoorchainApp struct {
	*baseapp.BaseApp

	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// Keys
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// Keepers
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.BaseKeeper
	StakingKeeper *stakingkeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper

	// Module manager
	mm     *module.Manager
	config module.Configurator
}

func NewNoorchainApp(
	logger sdklog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) *NoorchainApp {
	enc := MakeEncodingConfig()

	// BaseApp options
	var bappOpts []func(*baseapp.BaseApp)

	if cacheEnabled := cast.ToBool(appOpts.Get(server.FlagInterBlockCache)); cacheEnabled {
		bappOpts = append(bappOpts, baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager()))
	}

	// Pruning (flags/app.toml)
	if pruningStr := cast.ToString(appOpts.Get("pruning")); pruningStr != "" {
		pruningOpts := pruningtypes.NewPruningOptionsFromString(pruningStr)
		bappOpts = append(bappOpts, baseapp.SetPruning(pruningOpts))
	}

	if minGas := cast.ToString(appOpts.Get(server.FlagMinGasPrices)); minGas != "" {
		bappOpts = append(bappOpts, baseapp.SetMinGasPrices(minGas))
	}

	if haltHeight := cast.ToUint64(appOpts.Get(server.FlagHaltHeight)); haltHeight > 0 {
		bappOpts = append(bappOpts, baseapp.SetHaltHeight(haltHeight))
	}
	if haltTime := cast.ToUint64(appOpts.Get(server.FlagHaltTime)); haltTime > 0 {
		bappOpts = append(bappOpts, baseapp.SetHaltTime(haltTime))
	}

	// Propager le chain-id fourni via --chain-id aux options BaseApp
	if chainID := cast.ToString(appOpts.Get(flags.FlagChainID)); chainID != "" {
		bappOpts = append(bappOpts, baseapp.SetChainID(chainID))
	}

	bApp := baseapp.NewBaseApp(
		AppName,
		logger,
		db,
		enc.TxConfig.TxDecoder(),
		bappOpts...,
	)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetInterfaceRegistry(enc.InterfaceRegistry)
	bApp.SetTxEncoder(enc.TxConfig.TxEncoder())
	// Param store consensus in-memory
	bApp.SetParamStore(&NoopConsensusParamStore{})

	app := &NoorchainApp{
		BaseApp:           bApp,
		appCodec:          enc.Marshaler,
		txConfig:          enc.TxConfig,
		interfaceRegistry: enc.InterfaceRegistry,
	}

	// 1) Keys
	app.keys = map[string]*storetypes.KVStoreKey{
		authtypes.StoreKey:    storetypes.NewKVStoreKey(authtypes.StoreKey),
		banktypes.StoreKey:    storetypes.NewKVStoreKey(banktypes.StoreKey),
		stakingtypes.StoreKey: storetypes.NewKVStoreKey(stakingtypes.StoreKey),
		paramstypes.StoreKey:  storetypes.NewKVStoreKey(paramstypes.StoreKey),
	}
	app.tkeys = map[string]*storetypes.TransientStoreKey{
		paramstypes.TStoreKey: storetypes.NewTransientStoreKey(paramstypes.TStoreKey),
	}
	app.memKeys = map[string]*storetypes.MemoryStoreKey{}

	// 2) Keepers
	app.ParamsKeeper = initParamsKeeper(enc.Marshaler, enc.Amino,
		app.keys[paramstypes.StoreKey],
		app.tkeys[paramstypes.TStoreKey],
	)

	app.AccountKeeper = authkeeper.NewAccountKeeper(
		enc.Marshaler,
		sdkruntime.NewKVStoreService(app.keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec("noor"),
		AppName,
		authtypes.NewModuleAddress("gov").String(),
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		enc.Marshaler,
		sdkruntime.NewKVStoreService(app.keys[banktypes.StoreKey]),
		app.AccountKeeper,
		BlockedModuleAccountAddrs(),
		authtypes.NewModuleAddress("gov").String(),
		logger,
	)

	app.StakingKeeper = stakingkeeper.NewKeeper(
		enc.Marshaler,
		sdkruntime.NewKVStoreService(app.keys[stakingtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.NewModuleAddress("gov").String(),
		addresscodec.NewBech32Codec("noorvaloper"),
		addresscodec.NewBech32Codec("noorvalcons"),
	)

	// 3) Mount stores
	app.MountKVStores(app.keys)
	app.MountTransientStores(app.tkeys)
	app.MountMemoryStores(app.memKeys)

	// 4) Module manager
	app.mm = module.NewManager(
		paramsmodule.NewAppModule(app.ParamsKeeper),
		authmodule.NewAppModule(enc.Marshaler, app.AccountKeeper, nil, nil),
		bankmodule.NewAppModule(enc.Marshaler, app.BankKeeper, app.AccountKeeper, nil),
		stakingmodule.NewAppModule(enc.Marshaler, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, nil),
		genutilmodule.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp, enc.TxConfig),
	)

	app.config = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.config)

	// Orderings
	app.mm.SetOrderBeginBlockers(stakingtypes.ModuleName)
	app.mm.SetOrderEndBlockers(stakingtypes.ModuleName)
	app.mm.SetOrderInitGenesis(
		paramstypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		stakingtypes.ModuleName,
		genutiltypes.ModuleName,
	)

	// 5) Lifecycle handlers
	app.SetInitChainer(func(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
		if req.ChainId != "" {
			ctx = ctx.WithChainID(req.ChainId)
		}

		genesisState := make(map[string]json.RawMessage)
		if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
			return nil, err
		}
		return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
	})

	app.SetBeginBlocker(app.mm.BeginBlock)
	app.SetEndBlocker(app.mm.EndBlock)

	// 6) Load latest version
	if err := app.LoadLatestVersion(); err != nil {
		panic(err)
	}

	return app
}

// Params keeper init (pattern v0.53)
func initParamsKeeper(
	appCodec codec.Codec,
	legacyAmino *codec.LegacyAmino,
	key *storetypes.KVStoreKey,
	tkey *storetypes.TransientStoreKey,
) paramskeeper.Keeper {
	pk := paramskeeper.NewKeeper(
		appCodec,
		legacyAmino,
		key,
		tkey,
	)
	pk.Subspace(authtypes.ModuleName)
	pk.Subspace(banktypes.ModuleName)
	pk.Subspace(stakingtypes.ModuleName)
	pk.Subspace(baseapp.Paramspace)
	return pk
}

// servertypes.Application required methods
func (app *NoorchainApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	cmtservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	node.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

func (app *NoorchainApp) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	cmtservice.RegisterGRPCGatewayRoutes(clientCtx, mux)
	node.RegisterGRPCGatewayRoutes(clientCtx, mux)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, mux)
}

func (app *NoorchainApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.GRPCQueryRouter(), clientCtx, app.Simulate, app.interfaceRegistry)
}

func (app *NoorchainApp) RegisterTendermintService(clientCtx client.Context) {
	cmtservice.RegisterTendermintService(clientCtx, app.GRPCQueryRouter(), app.interfaceRegistry, app.Query)
	node.RegisterNodeService(clientCtx, app.GRPCQueryRouter(), config.Config{})
}

func (app *NoorchainApp) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	// handled in RegisterTendermintService
}

func (app *NoorchainApp) AppCodec() codec.Codec     { return app.appCodec }
func (app *NoorchainApp) TxConfig() client.TxConfig { return app.txConfig }

// Compile-time check
var _ servertypes.Application = (*NoorchainApp)(nil)
