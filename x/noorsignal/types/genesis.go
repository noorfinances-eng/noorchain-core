package types

import "fmt"

// GenesisState defines the initial state of the PoSS (x/noorsignal) module.
//
// For PoSS Logic 1 – Step A, we only track a very simple global counter:
// - TotalSignals: how many social signals have been registered in total.
//   At this stage, this is just a placeholder field; it will be updated
//   later when we wire real PoSS logic (messages, daily limits, rewards, etc.).
type GenesisState struct {
	// TotalSignals is a global counter of all validated PoSS signals
	// (participants + curators) since the genesis of the chain.
	//
	// NOTE:
	// - For now, this value will always be 0 at genesis and will not
	//   be modified by any logic yet. It is a placeholder for future steps.
	TotalSignals uint64 `json:"total_signals" yaml:"total_signals"`
}

// NewGenesisState creates a new GenesisState instance.
//
// For now, it only sets the global TotalSignals counter.
func NewGenesisState(totalSignals uint64) *GenesisState {
	return &GenesisState{
		TotalSignals: totalSignals,
	}
}

// DefaultGenesis returns the default genesis state for the PoSS module.
//
// We start Noorchain with zero PoSS signals counted.
func DefaultGenesis() *GenesisState {
	return NewGenesisState(0)
}

// ValidateGenesis performs basic validation of the PoSS genesis state.
//
// At this stage, we only check that the pointer is not nil and that the
// TotalSignals counter is non-negative (which is guaranteed by uint64).
// More advanced checks (limits, params, reserved supply, etc.) will be added
// later during PoSS Logic 2/3.
func ValidateGenesis(gs *GenesisState) error {
	if gs == nil {
		return fmt.Errorf("noorsignal genesis state cannot be nil")
	}

	// No complex validation yet; this keeps the module very simple
	// while still allowing us to evolve the structure safely later.
	return nil
}
