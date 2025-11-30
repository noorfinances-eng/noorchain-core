package types

import "fmt"

// GenesisState defines the initial state of the x/noorsignal (PoSS) module.
//
// At this stage, the struct is intentionally empty. We will extend it later with:
// - PoSS params (daily limits, weights, etc.)
// - PoSS reserve balances
// - Curator / participant statistics
// - Any additional state required by the PoSS engine.
type GenesisState struct {
	// TODO: add fields in future PoSS phases.
}

// NewGenesisState creates a new (empty) GenesisState instance.
func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

// DefaultGenesis returns the default genesis state for x/noorsignal.
//
// For now, it returns an empty object. This is enough for Phase 4, as long as
// the module compiles and can be included in the global app genesis.
func DefaultGenesis() *GenesisState {
	return NewGenesisState()
}

// Validate performs basic validation of the genesis state.
//
// At this stage, there are no rules to enforce yet. We just make sure the
// object is non-nil so later checks can safely rely on it.
func (gs *GenesisState) Validate() error {
	if gs == nil {
		return fmt.Errorf("noorsignal genesis state cannot be nil")
	}

	// Future checks will go here (params ranges, non-negative values, etc.).

	return nil
}
