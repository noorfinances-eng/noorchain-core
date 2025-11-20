package noorsignal

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// ------------------------------------------------------------
//  AppModuleBasic : version minimale (codec, nom, genesis default)
// ------------------------------------------------------------

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return noorsignaltypes.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	noorsignaltypes.RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	noorsignaltypes.RegisterInterfaces(reg)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(noorsignaltypes.DefaultGenesis())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data noorsignaltypes.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return err
	}
	return noorsignaltypes.ValidateGenesis(data)
}

// ------------------------------------------------------------
//  AppModule : logique du module (keeper, services)
// ------------------------------------------------------------

type AppModule struct {
	AppModuleBasic
	keeper noorsignalkeeper.Keeper
}

func NewAppModule(k noorsignalkeeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

// RegisterServices : l’étape IMPORTANTISSIME
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// 1) Msg Server (submit + validate + admin)
	noorsignaltypes.RegisterMsgServer(
		cfg.MsgServer(),
		noorsignalkeeper.NewMsgServer(am.keeper),
	)

	// 2) Query Server (GetSignal, ListSignals, etc.)
	noorsignaltypes.RegisterQueryServer(
		cfg.QueryServer(),
		noorsignalkeeper.NewQueryServer(am.keeper),
	)
}

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (AppModule) Route() sdk.Route { return sdk.Route{} }
func (AppModule) QuerierRoute() string { return noorsignaltypes.ModuleName }

func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier { return nil }

// ------------------------------------------------------------
//  Lifecycles : InitGenesis / ExportGenesis
// ------------------------------------------------------------

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var gs noorsignaltypes.GenesisState
	cdc.MustUnmarshalJSON(data, &gs)

	am.keeper.SetConfig(ctx, gs.Config)

	// Charger Curators
	for _, c := range gs.Curators {
		addr, err := sdk.AccAddressFromBech32(c.Address)
		if err == nil {
			am.keeper.SetCurator(ctx, noorsignaltypes.Curator{
				Address:               addr,
				Level:                 c.Level,
				TotalSignalsValidated: c.TotalSignalsValidated,
				Active:                c.Active,
			})
		}
	}

	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	cfg, _ := am.keeper.GetConfig(ctx)

	curators := []noorsignaltypes.Curator{}
	// TODO: Exporter curators si besoin (pas nécessaire pour testnet V1)

	gs := noorsignaltypes.GenesisState{
		Config:   cfg,
		Signals:  []noorsignaltypes.Signal{}, // ignoré en V1
		Curators: curators,
	}

	return cdc.MustMarshalJSON(&gs)
}
