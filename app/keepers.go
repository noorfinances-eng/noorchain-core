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
// D'autres keepers (staking, gov, evm, etc.) pourront être ajoutés
// progressivement, dans des étapes futures.
type AppKeepers struct {
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper

	NoorSignalKeeper noorsignalkeeper.Keeper
}
