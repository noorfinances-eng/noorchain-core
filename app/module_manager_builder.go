package app

// BuildModuleManager constructs a placeholder ModuleManager
// containing all NOORCHAIN modules. No Cosmos SDK wiring yet.
func BuildModuleManager(mods AppModules) ModuleManager {
	return NewModuleManager()
}
