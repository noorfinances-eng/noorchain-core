package app

// CosmosEncodingBuilder is a placeholder builder for the CosmosEncodingConfig.
// In Phase 2 this simply delegates to MakeCosmosEncodingConfig.
type CosmosEncodingBuilder struct{}

// NewCosmosEncodingBuilder creates a new encoding builder instance.
func NewCosmosEncodingBuilder() CosmosEncodingBuilder {
	return CosmosEncodingBuilder{}
}

// Build constructs the default CosmosEncodingConfig.
func (b CosmosEncodingBuilder) Build() CosmosEncodingConfig {
	return MakeCosmosEncodingConfig()
}
