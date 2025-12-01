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
	return json.RawMessage(`{}`)
}

// ValidateGenesis performs genesis state validation.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	txCfg client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	// Later we will call noorsignaltypes.ValidateGenesis when we
	// wire a full JSON genesis codec without gogo/proto.
	return nil
}

// RegisterGRPCGatewayRoutes registers gRPC-Gateway routes.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// RegisterInterfaces registers the module's interface types.
//
// NOTE: on purpose, we don't register any proto Msg here to avoid
// gogo/proto dependencies for now.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {}

// GetTxCmd returns the root tx command for the module.
// For now, PoSS has no CLI tx commands.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns the root query command for the module.
// For now, PoSS has no CLI query commands.
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

// Name returns the module name (for safety, we re-expose it here).
func (am AppModule) Name() string {
	return noorsignaltypes.ModuleName
}

// RegisterServices registers module services (Msg/Query servers).
// We stay on legacy Route/Handler disabled for now (no proto Msg service).
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis initializes the module genesis state.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	// PoSS genesis wiring (using GenesisState) will be added later.
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the module genesis state.
func (am AppModule) ExportGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
) json.RawMessage {
	// For now, we export an empty JSON object.
	return json.RawMessage(`{}`)
}

// BeginBlock is called at the beginning of every block.
//
// NOTE: daily resets / halving / PoSS logic will be added later.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// no-op for now
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

// Legacy routing: for now, we don't expose any tx route.
// This keeps the module structurally valid without wiring MsgCreateSignal
// (which would require proto.Message).
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
