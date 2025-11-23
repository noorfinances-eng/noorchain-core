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
	// 1) Config PoSS : on applique simplement la config fournie
	// (si elle est "zero value", ce sera la responsabilité de DefaultGenesis).
	k.SetConfig(ctx, genState.Config)

	// 2) Curators
	for _, c := range genState.Curators {
		k.SetCurator(ctx, c)
	}

	// Pas de validators spécifiques à renvoyer pour PoSS.
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

	// Pour l’instant, on n'exporte pas les signaux ni les compteurs journaliers
	// et on laisse la liste de curators vide (ou à améliorer plus tard).
	return noorsignaltypes.GenesisState{
		Config:   cfg,
		Curators: []noorsignaltypes.Curator{},
	}
}
