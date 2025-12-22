package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	sdklog "cosmossdk.io/log"
	"cosmossdk.io/store"
	pruningtypes "cosmossdk.io/store/pruning/types"
)

const (
	AppName = "noorchain"
)

type NoorchainApp struct {
	*baseapp.BaseApp

	appCodec codec.Codec
	txConfig client.TxConfig
}

func NewNoorchainApp(
	logger sdklog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) *NoorchainApp {
	enc := MakeEncodingConfig()

	// Options BaseApp (pattern v0.53)
	var bappOpts []func(*baseapp.BaseApp)

	if cacheEnabled := cast.ToBool(appOpts.Get(server.FlagInterBlockCache)); cacheEnabled {
		bappOpts = append(bappOpts, baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager()))
	}

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

	return &NoorchainApp{
		BaseApp:  bApp,
		appCodec: enc.Marshaler,
		txConfig: enc.TxConfig,
	}
}

// servertypes.Application required methods (Phase 2 = no-op)
func (app *NoorchainApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {}

func (app *NoorchainApp) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

func (app *NoorchainApp) RegisterTxService(clientCtx client.Context) {}

func (app *NoorchainApp) RegisterTendermintService(clientCtx client.Context) {}

func (app *NoorchainApp) RegisterNodeService(clientCtx client.Context, cfg config.Config) {}

func (app *NoorchainApp) AppCodec() codec.Codec     { return app.appCodec }
func (app *NoorchainApp) TxConfig() client.TxConfig { return app.txConfig }

// Compile-time check
var _ servertypes.Application = (*NoorchainApp)(nil)
