package app

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitChainer est appelé une seule fois au lancement de la chaîne (genesis).
//
// Rôle :
// - décoder l'état genesis global (AppStateBytes)
// - déléguer l'initialisation des modules au ModuleManager
// - renvoyer les éventuelles mises à jour de validateurs.
func (app *App) InitChainer(
	ctx sdk.Context,
	req abci.RequestInitChain,
) abci.ResponseInitChain {
	// Sécurité : si on n'a pas de ModuleManager ou de codec, on ne fait rien.
	if app.Modules.Manager == nil || app.Encoding.Marshaler == nil {
		return abci.ResponseInitChain{}
	}

	// 1) Décoder l'état genesis (map "nomModule" -> JSON brut).
	var genesisState map[string]json.RawMessage
	if len(req.AppStateBytes) > 0 {
		if err := app.Encoding.Marshaler.UnmarshalJSON(req.AppStateBytes, &genesisState); err != nil {
			// Pour l'instant, on panic en cas de genesis invalide.
			// Plus tard, on pourra gérer l'erreur plus proprement.
			panic(err)
		}
	} else {
		genesisState = make(map[string]json.RawMessage)
	}

	// 2) Déléguer l'initialisation au ModuleManager.
	validatorUpdates := app.Modules.Manager.InitGenesis(
		ctx,
		app.Encoding.Marshaler,
		genesisState,
	)

	// 3) Retourner la réponse InitChain avec les ValidatorUpdates.
	return abci.ResponseInitChain{
		Validators: validatorUpdates,
	}
}
