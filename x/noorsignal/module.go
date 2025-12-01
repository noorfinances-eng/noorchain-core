package noorsignal

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

//
// -----------------------------------------------------------------------------
// AppModuleBasic
// -----------------------------------------------------------------------------

type AppModuleBasic struct{}

// Name returns the PoSS module name.
func (AppModuleBasic) Name() string {
	return noorsignaltypes.ModuleName
}

// RegisterLegacyAminoCodec registers the module's types on the Amino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// DefaultGenesis returns default genesis state as raw JSON.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	// For now, simple empty JSON. We will wire real genesis later if needed.
	return json.RawMessage(`{}`)
}

// ValidateGenesis performs genesis state validation.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	txCfg client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	// No validation for now (PoSS Logic: skeleton only).
	return nil
}

// RegisterGRPCGatewayRoutes registers gRPC-Gateway routes.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// RegisterInterfaces registers the module's interface types.
// NOTE: we keep this empty for now to avoid any dependency on protobuf
// generated types while PoSS Logic is still in pure-Go mode.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {}

// GetTxCmd returns the root tx command for the module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns the root query command for the module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

//
// -----------------------------------------------------------------------------
// AppModule
// -----------------------------------------------------------------------------

// AppModule is the full module type for x/noorsignal.
type AppModule struct {
	AppModuleBasic
	cdc    codec.Codec
	keeper keeper.Keeper
}

// Compile-time assertions: ensure AppModule satisfies module interfaces.
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// NewAppModule creates a new AppModule instance for x/noorsignal.
func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		cdc:            cdc,
		keeper:         k,
	}
}

// Name returns the module name.
func (am AppModule) Name() string {
	return noorsignaltypes.ModuleName
}

// RegisterServices registers module services (Msg/Query servers).
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis initializes the module genesis state.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	// PoSS logic will be wired later (using types.GenesisState if needed).
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the module genesis state.
func (am AppModule) ExportGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
) json.RawMessage {
	// For now, we just return an empty JSON object.
	return json.RawMessage(`{}`)
}

// BeginBlock is called at the beginning of every block.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// PoSS daily reset hook (compatible with PoSS Logic 7).
	am.keeper.ResetDailyCountersIfNeeded(ctx)
}

// EndBlock is called at the end of every block.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// ConsensusVersion returns the module consensus version.
func (AppModule) ConsensusVersion() uint64 {
	return 1
}

// RegisterInvariants registers module invariants.
func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// Legacy routing (kept empty, but required by interface).
func (AppModule) Route() sdk.Route {
	return sdk.Route{}
}

func (AppModule) QuerierRoute() string {
	return ""
}

// LegacyQuerierHandler is required by module.AppModule (v0.46).
// We don't use legacy queriers, so we simply return nil.
func (AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return nil
}

// CLI — no custom tx/query commands wired at AppModule level yet.
func (AppModule) GetTxCmd() *cobra.Command {
	return nil
}

func (AppModule) GetQueryCmd() *cobra.Command {
	return nil
}
