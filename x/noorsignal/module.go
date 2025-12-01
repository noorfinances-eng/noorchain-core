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

// DefaultGenesis returns default genesis state as raw JSON,
// using the pure Go GenesisState (no proto).
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	gs := noorsignaltypes.DefaultGenesis()
	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// ValidateGenesis performs genesis state validation using pure Go structs.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	txCfg client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	if len(bz) == 0 {
		// Empty genesis is treated as DefaultGenesis.
		return nil
	}

	var gs noorsignaltypes.GenesisState
	if err := json.Unmarshal(bz, &gs); err != nil {
		return err
	}

	return noorsignaltypes.ValidateGenesis(&gs)
}

// RegisterGRPCGatewayRoutes registers gRPC-Gateway routes.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// RegisterInterfaces registers the module's interface types.
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
// PoSS Logic: Msg/Query will be wired later (we keep this empty for now).
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis initializes the module genesis state in the KVStore.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	var gs noorsignaltypes.GenesisState

	if len(data) == 0 {
		gs = *noorsignaltypes.DefaultGenesis()
	} else {
		if err := json.Unmarshal(data, &gs); err != nil {
			panic(err)
		}
	}

	// Validate and persist via keeper.
	am.keeper.InitGenesis(ctx, gs)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the module genesis state from the KVStore.
func (am AppModule) ExportGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)

	bz, err := json.Marshal(&gs)
	if err != nil {
		panic(err)
	}

	return bz
}

// BeginBlock is called at the beginning of every block.
// PoSS daily logic (limits, halving, etc.) will be added later.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// No per-block PoSS state changes yet.
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
