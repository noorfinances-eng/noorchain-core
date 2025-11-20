package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// StoreKeys regroupe les clés de store (KV + transient) pour les
// modules principaux de NOORCHAIN.
//
// À ce stade, ces clés NE SONT PAS encore montées dans le multi-store
// de BaseApp. Elles servent de base pour l'instanciation future
// des keepers Cosmos SDK, du module PoSS et des modules EVM / FeeMarket.
type StoreKeys struct {
	// Stores KV principaux
	AuthKey       *storetypes.KVStoreKey
	BankKey       *storetypes.KVStoreKey
	StakingKey    *storetypes.KVStoreKey
	GovKey        *storetypes.KVStoreKey
	ParamsKey     *storetypes.KVStoreKey
	NoorSignalKey *storetypes.KVStoreKey

	// Stores pour EVM / FeeMarket (Ethermint)
	EvmKey       *storetypes.KVStoreKey
	FeeMarketKey *storetypes.KVStoreKey

	// Stores transients (temporaire, typiquement pour Params)
	ParamsTransientKey *storetypes.TransientStoreKey
}

// NewStoreKeys crée les clés de store pour les modules principaux.
//
// Remarque :
// - Les noms utilisés pour les KVStoreKey sont alignés avec les constantes
//   de modules.go (ModuleAuth, ModuleBank, etc.).
// - Pour Noorsignal, on réutilise ModuleNoorSignal.
// - Pour EVM et FeeMarket, on utilise ModuleEvm et ModuleFeeMarket.
func NewStoreKeys() StoreKeys {
	return StoreKeys{
		AuthKey:       storetypes.NewKVStoreKey(ModuleAuth),
		BankKey:       storetypes.NewKVStoreKey(ModuleBank),
		StakingKey:    storetypes.NewKVStoreKey(ModuleStaking),
		GovKey:        storetypes.NewKVStoreKey(ModuleGov),
		ParamsKey:     storetypes.NewKVStoreKey(ModuleParams),
		NoorSignalKey: storetypes.NewKVStoreKey(ModuleNoorSignal),

		EvmKey:       storetypes.NewKVStoreKey(ModuleEvm),
		FeeMarketKey: storetypes.NewKVStoreKey(ModuleFeeMarket),

		ParamsTransientKey: storetypes.NewTransientStoreKey("transient_params"),
	}
}
