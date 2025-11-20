package app

import (
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
)

// AppKeepers regroupe tous les keepers principaux utilisés par
// l'application NOORCHAIN.
//
// Pour l'instant, on y trouve :
// - AccountKeeper    : gestion des comptes (x/auth)
// - BankKeeper       : gestion des soldes et transferts (x/bank)
// - ParamsKeeper     : paramètres globaux (x/params)
// - NoorSignalKeeper : logique PoSS (module x/noorsignal)
//
// Et nous préparons déjà l'intégration "EVM-ready" :
// - EvmKeeper       : compatibilité EVM (smart contracts Solidity)
// - FeeMarketKeeper : gestion avancée du gas / base fee (x/feemarket)
//
// D'autres keepers (staking, gov, ibc, etc.) pourront être ajoutés
// progressivement dans des étapes futures.
type AppKeepers struct {
	// Cosmos core
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper

	// PoSS module (NOORCHAIN social signals)
	NoorSignalKeeper noorsignalkeeper.Keeper

	// EVM / Ethermint integration
	EvmKeeper       evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper
}
