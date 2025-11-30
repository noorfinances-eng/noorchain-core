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
// On retourne simplement un objet JSON vide.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return json.RawMessage(`{}`)
}

// ValidateGenesis performs genesis state validation.
// Pour l’instant, on se contente de vérifier que le JSON est bien formé.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	txCfg client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	if len(bz) == 0 {
		return nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(bz, &raw); err != nil {
		return err
	}

	return nil
}

// RegisterGRPCGatewayRoutes registers gRPC-Gateway routes.
// Aucun endpoint pour l’instant.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// RegisterInterfaces registers the module's interface types.
// On complétera plus tard (Msgs, queries, etc.).
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {}

// GetTxCmd returns the root tx command for the module.
// On ne définit pas encore de commandes CLI, donc on renvoie nil.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns the root query command for the module.
// On ne définit pas encore de commandes CLI, donc on renvoie nil.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// -----------------------------------------------------------------------------
// AppModule
// -----------------------------------------------------------------------------

// AppModule is the full module type for x/noorsignal.
// Pour l’instant, il ne contient que le codec (pas encore de keeper branché).
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

// RegisterInvariants registers the module invariants.
// PoSS minimal : aucune invariant à enregistrer pour le moment.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the module (legacy SDK).
// PoSS v1 : pas de Msg handler legacy, on renvoie une route vide.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute returns the module's querier route name.
// On ne supporte pas encore le querier legacy, donc chaîne vide.
func (am AppModule) QuerierRoute() string {
	return ""
}

// LegacyQuerierHandler returns the legacy querier handler for the module.
// Non utilisé avec le stack gRPC moderne → nil.
func (am AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return nil
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
	// PoSS minimal : on ignore le contenu pour l'instant.
	// Plus tard, on utilisera un vrai GenesisState et un keeper.
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the module genesis state.
func (am AppModule) ExportGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
) json.RawMessage {
	// Pour l’instant, on exporte simplement "{}".
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

// ConsensusVersion returns the x/noorsignal module consensus version.
// On démarre à 1 (première version du module PoSS).
func (am AppModule) ConsensusVersion() uint64 {
	return 1
}
