package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	sdklog "cosmossdk.io/log"
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
) *NoorchainApp {
	enc := MakeEncodingConfig()

	bApp := baseapp.NewBaseApp(
		AppName,
		logger,
		db,
		enc.TxConfig.TxDecoder(),
	)

	// Important v0.53 : BaseApp doit connaître l’InterfaceRegistry utilisée
	// pour les Any, les signatures, et la conversion d’adresses.
	bApp.SetInterfaceRegistry(enc.InterfaceRegistry)

	bApp.SetCommitMultiStoreTracer(traceStore)

	return &NoorchainApp{
		BaseApp:   bApp,
		appCodec:  enc.Marshaler,
		txConfig:  enc.TxConfig,
	}
}

// servertypes.Application required methods (Phase 2 = no-op)
func (app *NoorchainApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
}

func (app *NoorchainApp) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
}

func (app *NoorchainApp) RegisterTxService(clientCtx client.Context) {
}

func (app *NoorchainApp) RegisterTendermintService(clientCtx client.Context) {
}

func (app *NoorchainApp) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
}

func (app *NoorchainApp) AppCodec() codec.Codec     { return app.appCodec }
func (app *NoorchainApp) TxConfig() client.TxConfig { return app.txConfig }

// Compile-time check
var _ servertypes.Application = (*NoorchainApp)(nil)
