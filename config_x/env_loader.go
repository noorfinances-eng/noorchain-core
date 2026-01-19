package config

import "os"

// LoadFromEnv is a placeholder that loads configuration values
// from environment variables. In Phase 2 only defaults are used.
func LoadFromEnv() Config {
	chainID := os.Getenv("NOORCHAIN_CHAIN_ID")
	env := os.Getenv("NOORCHAIN_ENV")

	if chainID == "" {
		chainID = "noorchain_9000-1"
	}
	if env == "" {
		env = "phase2"
	}

	return Config{
		ChainID: chainID,
		Env:     env,
	}
}
