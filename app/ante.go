package app

import (
	ethermintante "github.com/evmos/ethermint/app/ante"
)

// SetupAnteHandler configure le vrai AnteHandler Ethermint (EVM + Cosmos)
// pour NOORCHAIN, en utilisant nos keepers et la TxConfig de l'app.
func (app *NoorchainApp) SetupAnteHandler() {
	if app == nil {
		return
	}

	// Construction des options Ethermint.
	// validate() exige seulement :
	// - AccountKeeper (non nil)
	// - BankKeeper (non nil)
	// - SignModeHandler (non nil)
	// - FeeMarketKeeper (non nil)
	// - EvmKeeper (non nil)
	options := ethermintante.HandlerOptions{
		AccountKeeper:   app.AccountKeeper,
		BankKeeper:      app.BankKeeper,
		IBCKeeper:       nil, // IBC pas encore câblé
		FeeMarketKeeper: app.FeeMarketKeeper,
		EvmKeeper:       app.EvmKeeper,
		FeegrantKeeper:  nil, // Feegrant non utilisé pour l’instant

		SignModeHandler: app.TxConfig.SignModeHandler(),

		// Gas signatures : on utilise le consommateur par défaut d’Ethermint
		SigGasConsumer: ethermintante.DefaultSigVerificationGasConsumer,

		// Pas de limite MaxTxGasWanted spécifique au niveau app pour l’instant
		MaxTxGasWanted: 0,

		// Pas de règles spéciales pour authz / extension options / fee checker
		ExtensionOptionChecker: nil,
		TxFeeChecker:           nil,
		DisabledAuthzMsgs:      []string{},
	}

	anteHandler, err := ethermintante.NewAnteHandler(options)
	if err != nil {
		// Si la config est invalide, on préfère planter fort et net
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
}
