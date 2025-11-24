package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// CosmosStoreLoader is responsible for mounting the store keys
// inside the Cosmos BaseApp.
type CosmosStoreLoader struct{}

// NewCosmosStoreLoader returns an empty loader for Phase 2.
func NewCosmosStoreLoader() CosmosStoreLoader {
	return CosmosStoreLoader{}
}

// LoadStores mounts the Cosmos store keys into the BaseApp.
// In Phase 2, only the basic keys are mounted.
func LoadStores(base *CosmosBaseApp, keys CosmosKeys) {
	if base == nil {
		return
	}

	// Mounting the basic store keys (Phase 2)
	base.MountStore(keys.MainKey, storetypes.StoreTypeIAVL)
	base.MountStore(keys.TransientKey, storetypes.StoreTypeTransient)
	base.MountStore(keys.MemoryKey, storetypes.StoreTypeMemory)
}
