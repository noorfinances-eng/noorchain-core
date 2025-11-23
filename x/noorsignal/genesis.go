package noorsignal

import (
	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// InitGenesis initialise le module PoSS (noorsignal) à partir du GenesisState.
func InitGenesis(
	ctx sdk.Context,
	k noorsignalkeeper.Keeper,
	genState noorsignaltypes.GenesisState,
) []abci.ValidatorUpdate {
	// 1) Config PoSS
	if genState.Config != nil {
		k.SetConfig(ctx, *genState.Config)
	} else {
		k.InitDefaultConfig(ctx)
	}

	// 2) Curators
	for _, c := range genState.Curators {
		k.SetCurator(ctx, c)
	}

	// 3) (Optionnel) Signaux initiaux, compteurs, etc.
	// Pour l’instant on ne charge rien d’autre.

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exporte l'état courant du module PoSS vers un GenesisState.
func ExportGenesis(
	ctx sdk.Context,
	k noorsignalkeeper.Keeper,
) noorsignaltypes.GenesisState {
	cfg, found := k.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	// Pour l’instant, on n'exporte pas les signaux ni les compteurs journaliers.
	// On retournera une liste vide de curators tant qu’on n’a pas besoin de plus.
	return noorsignaltypes.GenesisState{
		Config:   &cfg,
		Curators: []noorsignaltypes.Curator{},
	}
}
