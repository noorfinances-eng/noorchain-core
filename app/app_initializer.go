package app

// InitApp creates a basic NOORCHAIN App instance.
// In Phase 2 this only wires placeholder modules and a placeholder module manager.
func InitApp() *App {
	mods := NewAppModules()
	_ = BuildModuleManager(mods)
	return New()
}
