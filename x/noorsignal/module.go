package noorsignal

import (
	"context"
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// -------------------------------------
// AppModuleBasic : partie "statique"
// -------------------------------------

// AppModuleBasic gère les éléments de base du module (nom, codec, genesis).
type AppModuleBasic struct{}

// NewAppModuleBasic retourne une instance d'AppModuleBasic.
func NewAppModuleBasic() AppModuleBasic {
	return AppModuleBasic{}
}

// Name retourne le nom du module.
func (AppModuleBasic) Name() string {
	return noorsignaltypes.ModuleName
}

// RegisterLegacyAminoCodec enregistre les types legacy Amino (si besoin).
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// Pour l'instant, aucun type spécifique Amino.
	// On pourra ajouter l'enregistrement des Msgs si nécessaire.
}

// RegisterInterfaces enregistre les interfaces Protobuf (Msg, Query, etc.).
func (AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	// Plus tard, on enregistrera ici :
	// - les Msg (MsgSubmitSignal, MsgValidateSignal)
	// - les éventuelles interfaces supplémentaires.
	noorsignaltypes.RegisterInterfaces(reg)
}

// DefaultGenesis retourne un état genesis par défaut pour le module PoSS.
//
// Ici, on initialise simplement la configuration PoSS par défaut
// et une liste de signaux/curators vide.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	state := noorsignaltypes.GenesisState{
		Config:   noorsignaltypes.DefaultPossConfig(),
		Signals:  []noorsignaltypes.Signal{},
		Curators: []noorsignaltypes.Curator{},
	}
	return cdc.MustMarshalJSON(&state)
}

// ValidateGenesis vérifie que le genesis fourni est valide pour le module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	if bz == nil {
		// Pas de genesis, rien à valider.
		return nil
	}

	var state noorsignaltypes.GenesisState
	if err := cdc.UnmarshalJSON(bz, &state); err != nil {
		return err
	}

	// On pourrait ajouter ici des validations plus poussées :
	// - cohérence des parts 70/30
	// - limites, etc.
	// Pour l'instant, on accepte tout état bien formé JSON.
	return nil
}

// RegisterGRPCGatewayRoutes enregistre les routes gRPC-Gateway (REST).
// Pour l'instant, aucune route spécifique PoSS.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	// TODO: enregistrer les routes gRPC-Gateway si nécessaire.
}

// GetTxCmd retourne la commande CLI principale de transactions du module.
// Pour l'instant, on ne fournit aucune commande CLI.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd retourne la commande CLI principale de requêtes du module.
// Pour l'instant, on ne fournit aucune commande CLI.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// -------------------------------------
// AppModule : partie "avec Keeper"
// -------------------------------------

// AppModule est la structure principale du module PoSS intégrée
// dans l'application NOORCHAIN. Elle possède un Keeper.
type AppModule struct {
	AppModuleBasic

	keeper noorsignalkeeper.Keeper
}

// NewAppModule construit un AppModule à partir d'un Keeper PoSS.
func NewAppModule(k noorsignalkeeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(),
		keeper:         k,
	}
}

// RegisterServices enregistre les services gRPC (Msg, Query).
//
// Pour l'instant, cette fonction est un squelette. Plus tard, on
// y enregistrera :
// - le MsgServer (SubmitSignal, ValidateSignal)
// - le QueryServer (lecture des signaux, config, etc.).
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// Exemple futur :
	// noorsignaltypes.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(am.keeper))
	// noorsignaltypes.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(am.keeper))
}

// InitGenesis initialise l'état du module à partir du genesis.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) []abci.ValidatorUpdate {
	if bz == nil {
		// Si pas de genesis spécifique, on init la config par défaut.
		am.keeper.InitDefaultConfig(ctx)
		return []abci.ValidatorUpdate{}
	}

	var state noorsignaltypes.GenesisState
	if err := cdc.UnmarshalJSON(bz, &state); err != nil {
		// En cas d'erreur, on initialise simplement la config par défaut.
		am.keeper.InitDefaultConfig(ctx)
		return []abci.ValidatorUpdate{}
	}

	// Initialiser la configuration PoSS depuis le genesis.
	am.keeper.SetConfig(ctx, state.Config)

	// Plus tard, on pourra restaurer :
	// - les signaux existants
	// - les curators
	// etc.

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exporte l'état du module vers un genesis JSON.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	// Pour l'instant, on exporte uniquement la configuration PoSS,
	// avec des listes vides pour signaux et curators.
	cfg, found := am.keeper.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	state := noorsignaltypes.GenesisState{
		Config:   cfg,
		Signals:  []noorsignaltypes.Signal{},
		Curators: []noorsignaltypes.Curator{},
	}

	return cdc.MustMarshalJSON(&state)
}

// BeginBlock est appelé au début de chaque bloc.
// Pour l'instant, aucune logique PoSS spécifique au BeginBlock.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// Plus tard, on pourra :
	// - réinitialiser des compteurs quotidiens
	// - appliquer des règles temps/dates, etc.
}

// EndBlock est appelé à la fin de chaque bloc.
// Pour l'instant, aucune logique PoSS spécifique au EndBlock.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
