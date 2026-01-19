package app

// Chain constants for NOORCHAIN.
// (Used later when wiring CometBFT and genesis configuration.)

const (
    // Global network identifiers (Phase 2 minimal)
    AppName    = "NOORCHAIN"
    ChainID    = "noorchain-1"

    // Coin metadata (will evolve in Phase 3+)
    BondDenom  = "unur" // base minimal denom
)
