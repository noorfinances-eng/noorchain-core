package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
)

// CosmosModuleManager is a placeholder wrapping the Cosmos SDK module manager.
// In Phase 2, this structure remains empty until modules are wired.
type CosmosModuleManager struct {
	Manager *module.Manager
}

// NewCosmosModuleManager returns an empty CosmosModuleManager placeholder.
func NewCosmosModuleManager() CosmosModuleManager {
	return CosmosModuleManager{
		Manager: module.NewManager(), // empty for now
	}
}
