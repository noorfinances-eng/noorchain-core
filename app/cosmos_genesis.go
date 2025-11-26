package app

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState defines the NOORCHAIN genesis structure.
// Phase 2: empty objects for each module.
type GenesisState map[string]json.RawMessage

// DefaultGenesis returns a minimal genesis with empty module states.
func DefaultGenesis() GenesisState {
	return GenesisState{
		"auth":      MustMarshalEmpty(),
		"bank":      MustMarshalEmpty(),
		"staking":   MustMarshalEmpty(),
		"gov":       MustMarshalEmpty(),
		"evm":       MustMarshalEmpty(),
		"feemarket": MustMarshalEmpty(),
	}
}

// MustMarshalEmpty returns an empty JSON object `{}` in RawMessage form.
func MustMarshalEmpty() json.RawMessage {
	b, err := json.Marshal(struct{}{})
	if err != nil {
		panic(err)
	}
	return b
}

// ValidateGenesis performs a basic validation of the genesis state.
// In Phase 2, everything is considered valid.
func ValidateGenesis(state GenesisState) error {
	_ = state
	return nil
}

// InitGenesis initializes NOORCHAIN state from a given genesis state.
// In Phase 2, no state is actually written.
func (app *NOORChainApp) InitGenesis(ctx sdk.Context, state GenesisState) {
	_ = ctx
	_ = state
}
