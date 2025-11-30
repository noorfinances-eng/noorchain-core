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

//
// -----------------------------------------------------------------------------
// AppModuleBasic
// -----------------------------------------------------------------------------

// AppModuleBasic implements the basic methods for the PoSS module without any
// keeper wiring. C'est le squelette minimal : pas de logique métier ici.
type AppModuleBasic struct{}

// Name returns the PoSS module name.
func (AppModuleBasic) Name() string {
	return noorsignaltypes.ModuleName
}

// RegisterLegacyAminoCodec registers the module's types on the Amino codec.
// Pour l'instant, aucun type spécifique.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// DefaultGenesis returns default genesis state as raw JSON.
// On renvoie un JSON vide pour l'instant : "{}".
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return json.RawMessage(`{}`)
}

// ValidateGenesis performs genesis state validation.
// Pour le squelette, on ne valide rien encore.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	txCfg client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	return nil
}

// RegisterGRPCGatewayRoutes registers gRPC-Gateway routes.
// Aucun endpoint pour l’instant.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// RegisterInterfaces registers the module's interface types.
// On complétera plus tard (Msgs, queries, etc.).
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {}

// GetTxCmd returns the root tx command for the module.
// PoSS 2 : pas encore de CLI, on renvoie nil.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns the root query command for the module.
// PoSS 2 : pas encore de CLI, on renvoie nil.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

//
// -----------------------------------------------------------------------------
// AppModule
// -----------------------------------------------------------------------------

// AppModule is the full module type for x/noorsignal.
// Pour l’instant, il ne contient que le codec (pas encore de keeper métier).
type AppModule struct {
	AppModuleBasic

	cdc codec.Codec
}

// NewAppModule creates a new AppModule instance for x/noorsignal.
// À ce stade, il est volontairement "léger".
func NewAppModule(cdc codec.Codec) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		cdc:            cdc,
	}
}

// RegisterServices registers module services (Msg/Query servers).
// On les ajoutera quand on aura les proto + keepers PoSS.
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis initializes the module genesis state.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	// PoSS 2 : pas de state encore, on ne fait que retourner une liste vide.
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the module genesis state.
func (am AppModule) ExportGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
) json.RawMessage {
	// PoSS 2 : on renvoie un JSON vide.
	return json.RawMessage(`{}`)
}

// BeginBlock is called at the beginning of every block.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// La logique PoSS par bloc viendra plus tard (limites journalières, etc.).
}

// EndBlock is called at the end of every block.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	// Pas de changement sur les validateurs pour PoSS.
	return []abci.ValidatorUpdate{}
}

// RegisterInvariants registers invariants for the module.
// Pour l’instant, aucune invariant métier.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// Route returns the message routing key for the module.
// PoSS 2 : pas encore de Msg handler, on renvoie une route vide.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute returns the module's querier route name.
// PoSS 2 : aucun legacy querier, donc chaîne vide.
func (am AppModule) QuerierRoute() string {
	return ""
}

// LegacyQuerierHandler returns the legacy querier handler for the module.
// PoSS 2 : aucun handler legacy, on renvoie nil.
func (am AppModule) LegacyQuerierHandler(cdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// ConsensusVersion returns the module consensus version.
// On démarre simplement à 1.
func (am AppModule) ConsensusVersion() uint64 {
	return 1
}
