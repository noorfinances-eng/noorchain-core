package app

// RuntimeIntegration is a placeholder combining the ExtendedRuntime
// and RuntimeServices into a unified structure for Phase 2.
type RuntimeIntegration struct {
	Extended ExtendedRuntime
	Services RuntimeServices
}

// NewRuntimeIntegration assembles the placeholder parts of the runtime.
func NewRuntimeIntegration() RuntimeIntegration {
	return RuntimeIntegration{
		Extended: NewExtendedRuntime(),
		Services: NewRuntimeServices(),
	}
}
