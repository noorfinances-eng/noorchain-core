package app

// Chain constants for NOORCHAIN.
// (Used when wiring CometBFT, CLI and genesis configuration.)

const (
	// Global network identifiers
	AppName = "NOORCHAIN"
	ChainID = "noorchain-1"

	// Default node home directory (used by CLI: init, start, etc.)
	// Kept explicit to avoid any ambiguity during testnet runs.
	DefaultNodeHome = ".noor-public-testnet"

	// Coin metadata
	BondDenom = "unur" // base minimal denom
)
