package app

import (
	"encoding/json"
)

// ----------------------------------------------------------------------------
// Genesis State Types (Phase 2 placeholder)
// ----------------------------------------------------------------------------

// GenesisState defines the NOORCHAIN genesis structure.
// Phase 2: empty objects for each module.
type GenesisState map[string]json.RawMessage

// DefaultGenesis returns a minimal genesis with empty module states.
func DefaultGenesis() GenesisState {
	return GenesisState{
		"auth":       MustMarshalEmpty(),
		"bank":       MustMarshalEmpty(),
		"staking":    MustMarshalEmpty(),
		"gov":        MustMarshalEmpty(),
		"evm":        MustMarshalEmpty(),
		"feemarket":  MustMarshalEmpty(),
	}
}

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------

// MustMarshalEmpty returns an empty JSON object `{}` in RawMessage form.
func MustMarshalEmpty() json.RawMessage {
	b, err := json.Marshal(struct{}{})
	if err != nil {
		panic(err)
	}
	return b
}

// ----------------------------------------------------------------------------
// Validation (Phase 2 placeholder)
// ----------------------------------------------------------------------------

// ValidateGenesis performs a basic validation of the genesis state.
// In Phase 2, everything is considered valid.
func ValidateGenesis(data GenesisState) error {
	// No validation in Phase 2
	_ = data
	return nil
}

// ----------------------------------------------------------------------------
// InitGenesis (Phase 2 placeholder)
// ----------------------------------------------------------------------------

// InitGenesis initializes NOORCHAIN state from a given genesis state.
// In Phase 2, no state is actually written.
func (app *NOORChainApp) InitGenesis(ctx Context, data GenesisState) {
	_ = ctx
	_ = data
	// No state initialization in Phase 2
}
