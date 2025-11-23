package app

// EncodingConfig represents the placeholder for all
// codecs and interface registries used by the application.
// It will be fully implemented when Cosmos SDK modules are wired.
type EncodingConfig struct{}

// MakeEncodingConfig returns an empty placeholder config for now.
func MakeEncodingConfig() EncodingConfig {
	return EncodingConfig{}
}
