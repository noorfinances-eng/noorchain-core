package noorsignal

import (
	"context"
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// ------------------------------------------------------------
// AppModuleBasic : partie "statique" du module (sans Keeper)
// ------------------------------------------------------------

type AppModuleBasic struct{}

var _ module.AppModuleBasic = AppModuleBasic{}

func (AppModuleBasic) Name() string { return noorsignaltypes.ModuleName }

func (AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

// Pour l’instant, on ne déclare pas encore d’interfaces spécifiques.
func (AppModuleBasic) RegisterInterfaces(_ cdctypes.InterfaceRegistry) {}

// DefaultGenesis renvoie l'état de genèse par défaut du module.
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	// Ici on ne dépend PAS de types.DefaultGenesis (qui n'existe pas encore).
	gs := noorsignaltypes.GenesisState{
		Config:   noorsignaltypes.DefaultPossConfig(),
		Curators: []noorsignaltypes.Curator{},
	}

	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// ValidateGenesis valide l'état de genèse (version minimale).
func (AppModuleBasic) ValidateGenesis(
	_ codec.JSONCodec,
	_ client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	if len(bz) == 0 {
		return nil
	}

	var gs noorsignaltypes.GenesisState
	if err := json.Unmarshal(bz, &gs); err != nil {
		return err
	}

	// Version minimale : on ne fait pas encore de validation complexe.
	return nil
}

// RegisterGRPCGatewayRoutes : pas encore de routes spécifiques.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(
	_ client.Context,
	_ *runtime.ServeMux,
) {
}

// GetTxCmd : pas encore de commandes CLI Tx spécifiques.
func (AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

// GetQueryCmd : pas encore de commandes CLI Query spécifiques.
func (AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

// ------------------------------------------------------------
// AppModule : partie "avec Keeper" (logique métier)
// ------------------------------------------------------------

type AppModule struct {
	AppModuleBasic
	keeper noorsignalkeeper.Keeper
}

// On s'assure qu'on implémente bien module.AppModule
var _ module.AppModule = AppModule{}

// IsAppModule est une méthode "marqueur" demandée par certains SDK récents.
func (AppModule) IsAppModule() {}

// NewAppModule construit le module PoSS avec son Keeper.
func NewAppModule(k noorsignalkeeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

// RegisterServices : pour l’instant, on ne branche PAS encore Msg/Query gRPC
// car les fichiers proto / pb.go ne sont pas finalisés.
func (am AppModule) RegisterServices(_ module.Configurator) {
	// On ajoutera plus tard :
	// - RegisterMsgServer
	// - RegisterQueryServer
}

// InitGenesis initialise l'état de genèse du module.
func (am AppModule) InitGenesis(
	ctx context.Context,
	_ codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	var gs noorsignaltypes.GenesisState
	if len(data) == 0 {
		gs = noorsignaltypes.GenesisState{
			Config:   noorsignaltypes.DefaultPossConfig(),
			Curators: []noorsignaltypes.Curator{},
		}
	} else {
		if err := json.Unmarshal(data, &gs); err != nil {
			panic(err)
		}
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return InitGenesis(sdkCtx, am.keeper, gs)
}

// ExportGenesis exporte l'état de genèse courant.
func (am AppModule) ExportGenesis(
	ctx context.Context,
	_ codec.JSONCodec,
) json.RawMessage {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gs := ExportGenesis(sdkCtx, am.keeper)

	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// ConsensusVersion permet d’indiquer une version du module (pour migrations futures).
func (am AppModule) ConsensusVersion() uint64 { return 1 }
