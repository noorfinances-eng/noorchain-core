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

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// -----------------------------------------------------------------------------
// AppModuleBasic
// -----------------------------------------------------------------------------

// AppModuleBasic implements the basic, stateless methods of the PoSS module.
type AppModuleBasic struct{}

// Name returns the module name.
func (AppModuleBasic) Name() string {
	return noorsignaltypes.ModuleName
}

// RegisterLegacyAminoCodec registers module types on the Amino codec.
// We don't use Amino for PoSS, so this is intentionally empty.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// DefaultGenesis returns the default genesis state as raw JSON.
//
// IMPORTANT: we use encoding/json directly instead of codec.JSONCodec
// to avoid requiring proto.Message on GenesisState.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	gs := noorsignaltypes.DefaultGenesis()
	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// ValidateGenesis performs basic validation of the genesis state.
func (AppModuleBasic) ValidateGenesis(
	_ codec.JSONCodec,
	_ client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	if len(bz) == 0 {
		// Empty genesis is allowed and treated as default.
		return nil
	}

	var gs noorsignaltypes.GenesisState
	if err := json.Unmarshal(bz, &gs); err != nil {
		return err
	}

	// For now, no validation rules. When we add real PoSS fields,
	// we can call noorsignaltypes.ValidateGenesis(&gs).
	return nil
}

// RegisterGRPCGatewayRoutes registers gRPC-Gateway routes.
// No public PoSS gRPC endpoints yet.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// RegisterInterfaces registers module interface types.
// We will add Msg / Query interfaces later with real PoSS logic.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {}

// GetTxCmd returns the root tx command for the module.
// No CLI commands yet, so we return nil.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns the root query command for the module.
// No CLI queries yet, so we return nil.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// -----------------------------------------------------------------------------
// AppModule
// -----------------------------------------------------------------------------

// AppModule is the full module type for x/noorsignal.
// For now it only wraps AppModuleBasic and the codec.
// The Keeper will be wired here when PoSS logic is implemented.
type AppModule struct {
	AppModuleBasic
	cdc codec.Codec
}

// NewAppModule creates a new AppModule instance for x/noorsignal.
func NewAppModule(cdc codec.Codec) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		cdc:            cdc,
	}
}

// ConsensusVersion implements AppModule. We start at version 1.
func (am AppModule) ConsensusVersion() uint64 {
	return 1
}

// RegisterInvariants registers module invariants. None for PoSS yet.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// Route is deprecated in v0.46; PoSS does not use legacy routing.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute is deprecated in v0.46; PoSS does not use legacy queries.
func (am AppModule) QuerierRoute() string {
	return ""
}

// LegacyQuerierHandler is deprecated; we return nil (no legacy querier).
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers gRPC Msg and Query servers.
// For now, PoSS has no services yet.
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis initializes the module genesis state.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	_ codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	var gs noorsignaltypes.GenesisState

	if len(data) == 0 {
		def := noorsignaltypes.DefaultGenesis()
		if def != nil {
			gs = *def
		} else {
			gs = noorsignaltypes.GenesisState{}
		}
	} else {
		if err := json.Unmarshal(data, &gs); err != nil {
			panic(err)
		}
	}

	// Later: use the Keeper to set PoSS state from gs.
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the module genesis state.
func (am AppModule) ExportGenesis(
	ctx sdk.Context,
	_ codec.JSONCodec,
) json.RawMessage {
	// Later: build the real PoSS genesis from keeper state.
	gs := noorsignaltypes.DefaultGenesis()
	if gs == nil {
		gs = &noorsignaltypes.GenesisState{}
	}

	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// BeginBlock is called at the beginning of every block.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// PoSS per-block logic will be added here later
	// (daily limits, halving checks, etc.).
}

// EndBlock is called at the end of every block.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	// PoSS does not change validator set.
	return []abci.ValidatorUpdate{}
}
