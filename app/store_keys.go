package app

import (
	storetypes "cosmossdk.io/store/types"
)

// StoreKeys regroupe toutes les clés de KVStore et TransientStore
// utilisées par les modules NOORCHAIN.
//
// Ces clés sont montées dans BaseApp, puis injectées
// dans les keepers et les modules.
type StoreKeys struct {
	// Cosmos SDK core stores
	AuthKey    *storetypes.KVStoreKey
	BankKey    *storetypes.KVStoreKey
	StakingKey *storetypes.KVStoreKey
	GovKey     *storetypes.KVStoreKey
	ParamsKey  *storetypes.KVStoreKey

	// NOORCHAIN custom module: PoSS (x/noorsignal)
	NoorSignalKey *storetypes.KVStoreKey

	// Ethermint / EVM stores
	EvmKey       *storetypes.KVStoreKey
	FeeMarketKey *storetypes.KVStoreKey

	// Transient stores (params, fee market, etc.)
	ParamsTransientKey *storetypes.TransientStoreKey
}

// NewStoreKeys crée l’ensemble des clés de stores nécessaires
// pour l’application NOORCHAIN.
func NewStoreKeys() StoreKeys {
	return StoreKeys{
		// Cosmos SDK stores
		AuthKey:    storetypes.NewKVStoreKey(ModuleAuth),
		BankKey:    storetypes.NewKVStoreKey(ModuleBank),
		StakingKey: storetypes.NewKVStoreKey(ModuleStaking),
		GovKey:     storetypes.NewKVStoreKey(ModuleGov),
		ParamsKey:  storetypes.NewKVStoreKey(ModuleParams),

		// PoSS module store
		NoorSignalKey: storetypes.NewKVStoreKey(ModuleNoorSignal),

		// EVM & FeeMarket (Ethermint) — à cabler plus tard
		EvmKey:       storetypes.NewKVStoreKey(ModuleEvm),
		FeeMarketKey: storetypes.NewKVStoreKey(ModuleFeeMarket),

		// Transient stores
		ParamsTransientKey: storetypes.NewTransientStoreKey("transient_params"),
	}
}
