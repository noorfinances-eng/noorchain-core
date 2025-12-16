package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethermintante "github.com/evmos/ethermint/app/ante"
)

// SetupAnteHandler configure le vrai AnteHandler Ethermint (EVM + Cosmos)
// pour NoorchainApp, en utilisant les keepers déjà initialisés dans app.go.
//
// IMPORTANT (v0.46 + Ethermint v0.22):
// Pendant InitGenesis -> DeliverGenTxs, le ctx utilisé par DeliverTx peut ne pas porter
// MinGasPrices correctement. On injecte donc systématiquement app.minGasPricesDec dans ctx
// avant d'exécuter l'AnteHandler Ethermint, pour éviter le panic dans MinGasPriceDecorator.
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

	options := ethermintante.HandlerOptions{
		AccountKeeper:   app.AccountKeeper,
		BankKeeper:      app.BankKeeper,
		IBCKeeper:       nil,
		FeeMarketKeeper: app.FeeMarketKeeper,
		EvmKeeper:       app.EvmKeeper,
		FeegrantKeeper:  nil,

		SignModeHandler: app.txConfig.SignModeHandler(),
		SigGasConsumer:  ethermintante.DefaultSigVerificationGasConsumer,

		MaxTxGasWanted: 0,

		ExtensionOptionChecker: nil,
		TxFeeChecker:           nil,

		DisabledAuthzMsgs: []string{},
	}

	innerAnte, err := ethermintante.NewAnteHandler(options)
	if err != nil {
		panic(err)
	}

	// ✅ Critical wrapper: always inject min gas prices into ctx before running Ethermint ante.
	wrappedAnte := func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
		// Ensure ctx has proper MinGasPrices (also covers DeliverGenTxs at height=0).
		if app.minGasPricesDec != nil && len(app.minGasPricesDec) > 0 {
			ctx = ctx.WithMinGasPrices(app.minGasPricesDec)
		}
		return innerAnte(ctx, tx, simulate)
	}

	app.SetAnteHandler(wrappedAnte)
}
