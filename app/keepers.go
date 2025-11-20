package app

import (
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
)

// AppKeepers regroupe tous les keepers principaux utilisés par
// l'application NOORCHAIN.
//
// Pour l'instant, on y trouve :
// - AccountKeeper : gestion des comptes (x/auth)
// - BankKeeper    : gestion des soldes et transferts (x/bank)
// - ParamsKeeper  : paramètres globaux (x/params)
// - NoorSignalKeeper : logique PoSS (module x/noorsignal)
//
// Squelette EVM / FeeMarket :
// - EvmKeeper        : sera branché plus tard sur le module Ethermint x/evm
// - FeeMarketKeeper  : sera branché plus tard sur le module x/feemarket
type AppKeepers struct {
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper

	NoorSignalKeeper noorsignalkeeper.Keeper

	// Squelettes pour l'intégration Ethermint (EVM / FeeMarket).
	// On ne les typpe pas encore avec les keepers réels pour éviter
	// d'ajouter les dépendances Ethermint avant l'étape dédiée.
	EvmKeeper       interface{}
	FeeMarketKeeper interface{}
}
