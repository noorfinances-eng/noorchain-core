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
// On y trouve :
// - AccountKeeper : gestion des comptes (x/auth)
// - BankKeeper    : gestion des soldes et transferts (x/bank)
// - ParamsKeeper  : paramètres globaux (x/params)
// - EvmKeeper     : exécution EVM (smart-contracts Solidity, Ethermint)
// - FeeMarketKeeper : gestion du marché des frais EVM (gas, base fee, etc.)
// - NoorSignalKeeper : logique PoSS (module x/noorsignal)
//
// D'autres keepers (staking, gov, ibc, etc.) pourront être ajoutés
// progressivement, dans des étapes futures.
type AppKeepers struct {
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper

	EvmKeeper       evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper

	NoorSignalKeeper noorsignalkeeper.Keeper
}
