package app

// DefaultConfig holds placeholder configuration values for NOORCHAIN.
// In Phase 2 this is minimal and will be expanded in later phases.
type DefaultConfig struct {
	ChainID string
	Env     string
}

// NewDefaultConfig returns the default NOORCHAIN configuration.
func NewDefaultConfig() DefaultConfig {
	return DefaultConfig{
		ChainID: DefaultChainID,
		Env:     "phase2",
	}
}
