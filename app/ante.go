package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethermintante "github.com/evmos/ethermint/app/ante"
)

// SetupAnteHandler configure le vrai AnteHandler Ethermint (EVM + Cosmos)
// pour NoorchainApp, en utilisant les keepers déjà initialisés dans app.go.
//
// Si quelque chose d’important manque (txConfig ou EvmKeeper), on tombe
// proprement sur un ante handler NO-OP pour ne jamais casser la compilation.
func (app *NoorchainApp) SetupAnteHandler() {
	if app == nil {
		return
	}

	// Safety net : si l’app n’est pas complètement initialisée, on garde un ante minimal.
	if app.txConfig == nil || app.EvmKeeper == nil {
		app.SetAnteHandler(func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
			return ctx, nil
		})
		return
	}

	// Options Ethermint 0.22.0
	options := ethermintante.HandlerOptions{
		AccountKeeper:   app.AccountKeeper,
		BankKeeper:      app.BankKeeper,
		IBCKeeper:       nil, // Pas d’IBC pour l’instant
		FeeMarketKeeper: app.FeeMarketKeeper,
		EvmKeeper:       app.EvmKeeper,
		FeegrantKeeper:  nil, // Pas de feegrant au début

		SignModeHandler: app.txConfig.SignModeHandler(),
		SigGasConsumer:  ethermintante.DefaultSigVerificationGasConsumer,

		// 0 = pas de hard cap custom sur le gas demandé
		MaxTxGasWanted: 0,

		// Pas encore de logique spéciale ici
		ExtensionOptionChecker: nil,
		TxFeeChecker:           nil,

		// On ne désactive aucun type de Msg authz pour l’instant
		DisabledAuthzMsgs: []string{},
	}

	anteHandler, err := ethermintante.NewAnteHandler(options)
	if err != nil {
		// Si la config est invalide, on panique au démarrage plutôt
		// que d’avoir un noeud incohérent.
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
}
