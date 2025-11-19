package noorsignal

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"

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
}

// RegisterInterfaces enregistre les interfaces Protobuf (Msg, Query, etc.).
func (AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	noorsignaltypes.RegisterInterfaces(reg)
}

// DefaultGenesis retourne un état genesis par défaut pour le module PoSS.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	state := noorsignaltypes.GenesisState{
		Config:   noorsignaltypes.DefaultPossConfig(),
		Signals:  []noorsignaltypes.Signal{},
		Curators: []noorsignaltypes.Curator{},
	}
	return cdc.MustMarshalJSON(&state)
}

// ValidateGenesis vérifie que le genesis fourni est valide pour le module.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	_ module.ClientGenesisValidation,
	bz json.RawMessage,
) error {
	if bz == nil {
		return nil
	}

	var state noorsignaltypes.GenesisState
	if err := cdc.UnmarshalJSON(bz, &state); err != nil {
		return err
	}

	// TODO: validations supplémentaires si nécessaire (ex: parts 70/30 cohérentes).
	_ = state
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
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// Plus tard :
	// noorsignaltypes.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(am.keeper))
	// noorsignaltypes.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(am.keeper))
}

// InitGenesis initialise l'état du module à partir du genesis.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	bz json.RawMessage,
) []abci.ValidatorUpdate {
	if bz == nil {
		// Pas de genesis spécifique : config PoSS par défaut.
		am.keeper.InitDefaultConfig(ctx)
		return []abci.ValidatorUpdate{}
	}

	var state noorsignaltypes.GenesisState
	if err := cdc.UnmarshalJSON(bz, &state); err != nil {
		// En cas d'erreur de decoding, on se replie sur la config par défaut.
		am.keeper.InitDefaultConfig(ctx)
		return []abci.ValidatorUpdate{}
	}

	// 1) Config PoSS.
	am.keeper.SetConfig(ctx, state.Config)

	// 2) Enregistrer les Curators fournis dans le genesis (s'ils existent).
	for _, curator := range state.Curators {
		am.keeper.SetCurator(ctx, curator)
	}

	// 3) Les signaux éventuels fournis dans le genesis sont ignorés pour V1.

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exporte l'état du module vers un genesis JSON.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
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
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// Plus tard : logique PoSS liée au temps (reset quotidiens, etc.).
}

// EndBlock est appelé à la fin de chaque bloc.
func (am AppModule) EndBlock(
	ctx sdk.Context,
	_ abci.RequestEndBlock,
) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
