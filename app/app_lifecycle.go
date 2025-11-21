package app

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// InitChainer est appelé une seule fois au lancement de la chaîne (genesis).
//
// Rôle :
// - décoder l'état genesis complet
// - appliquer la configuration économique (NUR)
// - initialiser la configuration PoSS si présente
// - déléguer aux modules via ModuleManager.InitGenesis
func (app *App) InitChainer(
	ctx sdk.Context,
	req abci.RequestInitChain,
) abci.ResponseInitChain {

	if app.Modules.Manager == nil || app.Encoding.Marshaler == nil {
		return abci.ResponseInitChain{}
	}

	// --------------------------------------------------
	// 1) Décoder le genesis JSON -> map[string]json.RawMessage
	// --------------------------------------------------
	var genesisState map[string]json.RawMessage
	if len(req.AppStateBytes) > 0 {
		if err := app.Encoding.Marshaler.UnmarshalJSON(req.AppStateBytes, &genesisState); err != nil {
			panic(err)
		}
	} else {
		genesisState = make(map[string]json.RawMessage)
	}

	// --------------------------------------------------
	// 2) Ajouter les 5 comptes économiques (placeholders)
	// --------------------------------------------------
	ApplyEconomicGenesis(genesisState, app.Encoding.Marshaler)

	// --------------------------------------------------
	// 3) Initialiser PoSS si config présente
	// --------------------------------------------------
	if gs, ok := genesisState[noorsignaltypes.ModuleName]; ok {
		var possGenesis noorsignaltypes.GenesisState
		if err := app.Encoding.Marshaler.UnmarshalJSON(gs, &possGenesis); err != nil {
			panic(err)
		}
		// PoSS : configuration + curators génesis
		app.Keepers.NoorSignalKeeper.InitDefaultConfig(ctx)
		for _, c := range possGenesis.Curators {
			app.Keepers.NoorSignalKeeper.SetCurator(ctx, c)
		}
	}

	// --------------------------------------------------
	// 4) Déléguer au ModuleManager
	// --------------------------------------------------
	validatorUpdates := app.Modules.Manager.InitGenesis(
		ctx,
		app.Encoding.Marshaler,
		genesisState,
	)

	// --------------------------------------------------
	// 5) Finalisation InitChain
	// --------------------------------------------------
	return abci.ResponseInitChain{
		Validators: validatorUpdates,
	}
}
