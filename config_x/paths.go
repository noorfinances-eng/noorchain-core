package config

// Paths holds placeholder filesystem paths for NOORCHAIN.
// In Phase 2 these are simple placeholders.
type Paths struct {
	DataDir string
	ConfigDir string
}

// DefaultPaths returns minimal default filesystem paths.
func DefaultPaths() Paths {
	return Paths{
		DataDir:   "./data",
		ConfigDir: "./config",
	}
}
