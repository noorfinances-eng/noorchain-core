package app

import (
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
)

// AppKeepers regroupe tous les keepers principaux utilisés
// par l'application NOORCHAIN.
//
// À ce stade, ils ne sont pas encore initialisés.
// Les initialisations réelles arrivent dans les étapes suivantes.
type AppKeepers struct {
	// Cosmos SDK core
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper

	// NOOR PoSS module
	NoorSignalKeeper noorsignalkeeper.Keeper

	// ⚠️ Les keepers EVM / FeeMarket (Ethermint) seront ajoutés plus tard
	// quand on aura choisi une version Ethermint compatible Cosmos SDK v0.50+.
	// EvmKeeper       *evmkeeper.Keeper
	// FeeMarketKeeper *feemarketkeeper.Keeper
}
