package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// CosmosKeys holds placeholder Cosmos SDK store keys.
// Real keys will be added once modules are wired.
type CosmosKeys struct {
	MainKey      *storetypes.KVStoreKey
	TransientKey *storetypes.TransientStoreKey
	MemoryKey    *storetypes.MemoryStoreKey
}

// NewCosmosKeys returns an empty CosmosKeys placeholder.
// In next steps, each Cosmos module will receive its dedicated key(s).
func NewCosmosKeys() CosmosKeys {
	return CosmosKeys{
		MainKey:      storetypes.NewKVStoreKey("main"),
		TransientKey: storetypes.NewTransientStoreKey("transient"),
		MemoryKey:    storetypes.NewMemoryStoreKey("memory"),
	}
}
