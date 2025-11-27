package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/cometbft/cometbft/abci/types"
)

// BeginBlocker is the NOORCHAIN BeginBlock hook.
// In Phase 2 this is a placeholder and does nothing.
func (app *NOORChainApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	_ = ctx
	_ = req
	return abci.ResponseBeginBlock{}
}

// EndBlocker is the NOORCHAIN EndBlock hook.
// In Phase 2 this is a placeholder and does nothing.
func (app *NOORChainApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	_ = ctx
	_ = req
	return abci.ResponseEndBlock{}
}

// InitChainer is the NOORCHAIN initialization hook.
// In Phase 2 this wires the default (empty) genesis state and calls InitGenesis.
func (app *NOORChainApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	_ = req

	// Build a default NOORCHAIN genesis state (placeholder).
	genesisState := DefaultGenesis()

	// Initialize chain state from genesis (placeholder).
	app.InitGenesis(ctx, genesisState)

	return abci.ResponseInitChain{}
}
