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

func (AppModuleBasic) Name() string {
    return noorsignaltypes.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
    gs := noorsignaltypes.DefaultGenesis()
    bz, err := json.Marshal(gs)
    if err != nil {
        panic(err)
    }
    return bz
}

func (AppModuleBasic) ValidateGenesis(
    cdc codec.JSONCodec,
    txCfg client.TxEncodingConfig,
    bz json.RawMessage,
) error {
    if len(bz) == 0 {
        return nil
    }

    var gs noorsignaltypes.GenesisState
    if err := json.Unmarshal(bz, &gs); err != nil {
        return err
    }

    return noorsignaltypes.ValidateGenesis(&gs)
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {}

func (AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

func (AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

//
// -----------------------------------------------------------------------------
// AppModule
// -----------------------------------------------------------------------------

type AppModule struct {
    AppModuleBasic
    cdc    codec.Codec
    keeper keeper.Keeper
}

var (
    _ module.AppModule      = AppModule{}
    _ module.AppModuleBasic = AppModuleBasic{}
)

func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
    return AppModule{
        AppModuleBasic: AppModuleBasic{},
        cdc:            cdc,
        keeper:         k,
    }
}

func (am AppModule) Name() string {
    return noorsignaltypes.ModuleName
}

// 🚀 ROUTE LEGACY ENABLED HERE
func (am AppModule) Route() sdk.Route {
    return sdk.NewRoute(noorsignaltypes.ModuleName, NewHandler(am.keeper))
}

func (AppModule) QuerierRoute() string { return "" }

func (AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
    return nil
}

func (am AppModule) RegisterServices(cfg module.Configurator) {}

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

    am.keeper.InitGenesis(ctx, gs)
    return []abci.ValidatorUpdate{}
}

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

func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
    return []abci.ValidatorUpdate{}
}

func (AppModule) ConsensusVersion() uint64 { return 1 }

func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}
