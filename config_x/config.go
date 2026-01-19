package config

// Config is a placeholder for external configuration handling.
// In Phase 2 this remains minimal and will be expanded later.
type Config struct {
	ChainID string
	Env     string
}

// New returns a minimal Config placeholder.
func New() Config {
	return Config{
		ChainID: "noorchain_9000-1",
		Env:     "phase2",
	}
}
