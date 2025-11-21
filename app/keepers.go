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
// Et on prépare déjà le terrain pour :
// - EvmKeeper       : futur keeper du module EVM (Ethermint)
// - FeeMarketKeeper : futur keeper du module feemarket (Ethermint)
//
// Pour l’instant, EvmKeeper et FeeMarketKeeper sont des placeholders (interface{}).
// Ils seront typés correctement quand on intégrera réellement Ethermint.
type AppKeepers struct {
	// Cosmos SDK core
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper

	// NOORCHAIN custom module (PoSS)
	NoorSignalKeeper noorsignalkeeper.Keeper

	// Ethermint / EVM (squelette, à remplir plus tard)
	EvmKeeper       interface{}
	FeeMarketKeeper interface{}
}
