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
//
// We bypass the Cosmos JSONCodec here and use encoding/json directly
// to avoid any dependency on gogo/proto.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	gs := noorsignaltypes.DefaultGenesis()
	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// ValidateGenesis performs genesis state validation using the PoSS
// GenesisState (TotalSignals, TotalMinted, MaxSignalsPerDay, 70/30 split).
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	txCfg client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	if len(bz) == 0 {
		// treat empty as default
		return nil
	}

	var gs noorsignaltypes.GenesisState
	if err := json.Unmarshal(bz, &gs); err != nil {
		return err
	}

	return noorsignaltypes.ValidateGenesis(&gs)
}

// RegisterGRPCGatewayRoutes registers gRPC-Gateway routes.
// No public gRPC endpoints yet for PoSS.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// RegisterInterfaces registers the module's interface types.
// We will plug Msg/Query types here later (PoSS Logic + proto).
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
// For now, PoSS has no Msg or Query service registered at the app level.
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis initializes the module genesis state.
//
// We decode the raw JSON into the PoSS GenesisState using encoding/json,
// validate it with ValidateGenesis, and will later use it to initialize
// the on-chain PoSS state.
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

	if err := noorsignaltypes.ValidateGenesis(&gs); err != nil {
		panic(err)
	}

	// PoSS Logic: later we will use gs to initialize counters, totals, etc.
	// For PoSS Logic 18, we keep InitGenesis structurally correct but neutral.
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the module genesis state.
//
// For now, we simply export the default genesis. Later, we will read the
// actual PoSS state (totals, params, etc.) from the keeper.
func (am AppModule) ExportGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
) json.RawMessage {
	gs := noorsignaltypes.DefaultGenesis()
	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// BeginBlock is called at the beginning of every block.
//
// In the future, this is where we will:
// - enforce daily PoSS limits,
// - apply halving-based adjustments,
// - update PoSS counters progressively.
//
// For now, it is intentionally empty to keep the chain logic neutral.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// PoSS daily logic will be added later (PoSS Logic runtime).
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
