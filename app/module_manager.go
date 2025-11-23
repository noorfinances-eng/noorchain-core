package app

// ModuleManager is a placeholder for the Cosmos SDK module manager.
// It will be wired with real modules (auth, bank, staking, gov, evm, feemarket)
// later in Phase 2.
type ModuleManager struct{}

// NewModuleManager returns an empty ModuleManager placeholder.
func NewModuleManager() ModuleManager {
	return ModuleManager{}
}
