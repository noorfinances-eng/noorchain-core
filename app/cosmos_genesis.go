package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState is a placeholder type for NOORCHAIN genesis state.
// In Phase 2 this remains empty and will be expanded in later phases.
type GenesisState struct{}

// DefaultGenesis returns a minimal default genesis state for NOORCHAIN.
// In Phase 2 this is an empty struct.
func DefaultGenesis() GenesisState {
	return GenesisState{}
}

// ValidateGenesis performs basic validation of the NOORCHAIN genesis state.
// In Phase 2 this always returns nil (no error).
func ValidateGenesis(_ GenesisState) error {
	return nil
}

// InitGenesis is a placeholder helper that will initialize the chain state
// from the given genesis state. For Phase 2 it does nothing.
func (app *NOORChainApp) InitGenesis(ctx sdk.Context, state GenesisState) {
	_ = ctx
	_ = state
}
