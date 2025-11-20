package noorsignal

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	"github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// InitGenesis initialise l'état du module noorsignal (PoSS) à partir du genesis.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	// 1) Config PoSS
	cfg := data.Config
	if !cfg.Enabled && cfg.BaseReward == 0 && cfg.ParticipantShare == 0 && cfg.CuratorShare == 0 {
		// Si aucune config n'est fournie, utiliser la config par défaut.
		cfg = types.DefaultPossConfig()
	}
	k.SetConfig(ctx, cfg)

	// 2) Curators initiaux
	for _, c := range data.Curators {
		k.SetCurator(ctx, c)
	}

	// 3) Signaux
	// Pour V1, on ignore les signaux au genesis, même si la liste est présente.
	// Ils seront uniquement créés à runtime via MsgSubmitSignal / MsgValidateSignal.

	// 4) Adresses économiques PoSS
	//
	// Pour l'instant, le module PoSS utilise les adresses de test
	// définies dans types/addresses.go (TestPoSSReserveAddr, etc.).
	// Les champs PoSSReserveAddr / FoundationAddr / ... du GenesisState
	// servent surtout pour l'export et la transparence.
	//
	// Lors du passage en mainnet, on pourra :
	// - remplir ces champs dans le genesis.json
	// - aligner types/addresses.go avec ces mêmes adresses.
}

// ExportGenesis exporte l'état courant du module noorsignal dans un GenesisState.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// 1) Config actuelle
	cfg, found := k.GetConfig(ctx)
	if !found {
		cfg = types.DefaultPossConfig()
	}

	// 2) Curators
	// Pour garder les choses simples dans V1, on n'itère pas sur tous les curators.
	// On exporte une liste vide. Plus tard, on pourra ajouter une fonction
	// d'itération dans le keeper pour retourner la liste complète.
	var curators []types.Curator

	// 3) Signals
	// Même logique : pour V1, on exporte une liste vide.
	var signals []types.Signal

	// 4) Adresses économiques (version TEST)
	//
	// On expose dans le GenesisState les adresses de test définies
	// dans types/addresses.go, de façon à ce que l'export genesis
	// reflète le modèle économique PoSS (5 / 5 / 5 / 5 / 80).
	return &types.GenesisState{
		Config:   cfg,
		Signals:  signals,
		Curators: curators,

		PoSSReserveAddr: types.TestPoSSReserveAddr.String(),
		FoundationAddr:  types.TestFoundationAddr.String(),
		DevWalletAddr:   types.TestDevWalletAddr.String(),
		StimulusAddr:    types.TestPoSSStimulusAddr.String(),
		PreSaleAddr:     types.TestPreSaleAddr.String(),
	}
}
