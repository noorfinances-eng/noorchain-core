package app

// Runtime is a placeholder for NOORCHAIN runtime wiring.
// In Phase 2 it remains minimal and does not initialize any real services.
type Runtime struct{}

// NewRuntime returns an empty Runtime placeholder.
func NewRuntime() Runtime {
	return Runtime{}
}
