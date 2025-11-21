package app

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitChainer est appelé une seule fois au lancement de la chaîne (genesis).
//
// Rôle :
// - décoder l'état genesis complet
// - appliquer la configuration économique (NUR)
// - déléguer aux modules via ModuleManager.InitGenesis
func (app *App) InitChainer(
	ctx sdk.Context,
	req abci.RequestInitChain,
) abci.ResponseInitChain {

	if app.Modules.Manager == nil || app.Encoding.Marshaler == nil {
		return abci.ResponseInitChain{}
	}

	// 1) Décoder le genesis JSON -> map[string]json.RawMessage
	var genesisState map[string]json.RawMessage
	if len(req.AppStateBytes) > 0 {
		if err := app.Encoding.Marshaler.UnmarshalJSON(req.AppStateBytes, &genesisState); err != nil {
			panic(err)
		}
	} else {
		genesisState = make(map[string]json.RawMessage)
	}

	// 2) Appliquer l'économie NUR (5/5/5/5/80) via Bank (balances + supply)
	ApplyEconomicGenesis(genesisState, app.Encoding.Marshaler)

	// 3) Déléguer l'initialisation à tous les modules via ModuleManager
	validatorUpdates := app.Modules.Manager.InitGenesis(
		ctx,
		app.Encoding.Marshaler,
		genesisState,
	)

	// 4) Retourner la réponse InitChain
	return abci.ResponseInitChain{
		Validators: validatorUpdates,
	}
}
