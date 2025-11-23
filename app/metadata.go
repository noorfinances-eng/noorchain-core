package app

// Metadata contains high-level token metadata for NOORCHAIN.
// In Phase 2 this is a simple placeholder.
type Metadata struct {
	Name     string
	Symbol   string
	Decimals int
}

// DefaultMetadata returns the default NOORCHAIN token metadata.
func DefaultMetadata() Metadata {
	return Metadata{
		Name:     DefaultDisplayDenom,
		Symbol:   DefaultDisplayDenom,
		Decimals: DefaultDecimals,
	}
}
