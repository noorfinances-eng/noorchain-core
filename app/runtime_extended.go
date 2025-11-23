package app

import (
	cfg "github.com/noorfinances-eng/noorchain-core/config"
)

// ExtendedRuntime is a placeholder for NOORCHAIN runtime with basic config loading.
// In Phase 2 this only assembles placeholders into a single structure.
type ExtendedRuntime struct {
	Config cfg.Config
	Paths  cfg.Paths
}

// NewExtendedRuntime returns an ExtendedRuntime with default config and paths.
func NewExtendedRuntime() ExtendedRuntime {
	return ExtendedRuntime{
		Config: cfg.Load(),
		Paths:  cfg.DefaultPaths(),
	}
}
