package config

// Load loads NOORCHAIN configuration from defaults.
// In Phase 2 this is a simple wrapper around New().
func Load() Config {
	return New()
}
