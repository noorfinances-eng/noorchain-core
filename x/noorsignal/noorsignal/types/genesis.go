package types

// GenesisState defines the initial state of the noorsignal (PoSS) module.
//
// For now it is intentionally minimal. We will add real PoSS fields later
// (PoSS params, pools, counters, etc.).
type GenesisState struct {
	// Reserved for future fields.
}

// NewGenesisState returns a new (empty) PoSS genesis state.
func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

// DefaultGenesis returns the default genesis state for the module.
func DefaultGenesis() *GenesisState {
	return NewGenesisState()
}

// ValidateGenesis validates the PoSS genesis state.
//
// At this stage, there is nothing to validate yet. We will add checks
// when we introduce real PoSS fields (limits, weights, pools...).
func ValidateGenesis(gs *GenesisState) error {
	if gs == nil {
		// Treat nil as an empty/default state.
		return nil
	}

	// No validation rules yet.
	return nil
}
