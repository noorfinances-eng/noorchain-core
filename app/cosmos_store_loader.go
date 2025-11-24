package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// CosmosStoreLoader is a placeholder responsible for setting up the store keys
// inside the Cosmos BaseApp. In Phase 2 it does not perform any initialization.
type CosmosStoreLoader struct{}

// NewCosmosStoreLoader returns an empty placeholder loader.
func NewCosmosStoreLoader() CosmosStoreLoader {
	return CosmosStoreLoader{}
}

// LoadStores attaches placeholder store keys to the BaseApp.
// In Phase 2, this function does not register real stores yet.
func LoadStores(base *CosmosBaseApp, keys CosmosKeys) {
	// Placeholder: real store mounting will happen in later steps.
	_ = base
	_ = keys
}
