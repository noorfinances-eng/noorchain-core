package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
)

// CosmosModuleManager is a wrapper around the Cosmos SDK module manager.
// In Phase 2, it prepares the structure for ordering InitGenesis and block hooks.
type CosmosModuleManager struct {
	Manager *module.Manager
}

// NewCosmosModuleManager returns a CosmosModuleManager with an empty Manager.
func NewCosmosModuleManager() CosmosModuleManager {
	return CosmosModuleManager{
		Manager: module.NewManager(), // empty for now
	}
}

// RegisterModules is a Phase 2 placeholder for registering app modules.
// The real registration (AUTH, BANK, STAKING, GOV, EVM, FEEMARKET) will be added later.
func (mm *CosmosModuleManager) RegisterModules() {
	if mm.Manager == nil {
		return
	}
	// Placeholder: module registration will be implemented in the wiring step.
}

// SetOrderInitGenesis configures the InitGenesis order for modules.
// Phase 2: no modules registered yet, so this calls the underlying method with no args.
func (mm *CosmosModuleManager) SetOrderInitGenesis() {
	if mm.Manager == nil {
		return
	}
	mm.Manager.SetOrderInitGenesis()
}

// SetOrderBeginBlockers configures the BeginBlocker order for modules.
// Phase 2: no modules registered yet, so this calls the underlying method with no args.
func (mm *CosmosModuleManager) SetOrderBeginBlockers() {
	if mm.Manager == nil {
		return
	}
	mm.Manager.SetOrderBeginBlockers()
}

// SetOrderEndBlockers configures the EndBlocker order for modules.
// Phase 2: no modules registered yet, so this calls the underlying method with no args.
func (mm *CosmosModuleManager) SetOrderEndBlockers() {
	if mm.Manager == nil {
		return
	}
	mm.Manager.SetOrderEndBlockers()
}
