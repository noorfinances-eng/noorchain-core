package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// -----------------------------------------------------------------------------
// GLOBAL PoSS STATS HELPERS
// -----------------------------------------------------------------------------

// getTotalSignals returns the total number of PoSS signals processed
// since genesis, stored under KeyTotalSignals.
func (k Keeper) getTotalSignals(ctx sdk.Context) uint64 {
	store := k.getStore(ctx)
	return noorsignaltypes.GetUint64(store, noorsignaltypes.KeyTotalSignals)
}

// setTotalSignals sets the total number of PoSS signals processed
// since genesis.
func (k Keeper) setTotalSignals(ctx sdk.Context, value uint64) {
	store := k.getStore(ctx)
	noorsignaltypes.SetUint64(store, noorsignaltypes.KeyTotalSignals, value)
}

// incrementTotalSignals adds delta to the global PoSS signal counter.
func (k Keeper) incrementTotalSignals(ctx sdk.Context, delta uint64) {
	current := k.getTotalSignals(ctx)
	k.setTotalSignals(ctx, current+delta)
}

// getTotalMinted returns the total amount of NUR minted via PoSS,
// stored as a raw string (for now) under KeyTotalMinted.
//
// The value is expected to be a decimal integer in "unur" (smallest unit),
// but we do not parse it here yet.
func (k Keeper) getTotalMinted(ctx sdk.Context) string {
	store := k.getStore(ctx)
	return noorsignaltypes.GetString(store, noorsignaltypes.KeyTotalMinted)
}

// setTotalMinted sets the total PoSS-minted NUR as a raw string.
func (k Keeper) setTotalMinted(ctx sdk.Context, value string) {
	store := k.getStore(ctx)
	noorsignaltypes.SetString(store, noorsignaltypes.KeyTotalMinted, value)
}

// getPendingMint returns the current pending PoSS mint (not yet distributed),
// stored as a raw string.
func (k Keeper) getPendingMint(ctx sdk.Context) string {
	store := k.getStore(ctx)
	return noorsignaltypes.GetString(store, noorsignaltypes.KeyPendingMint)
}

// setPendingMint sets the current pending PoSS mint as a raw string.
func (k Keeper) setPendingMint(ctx sdk.Context, value string) {
	store := k.getStore(ctx)
	noorsignaltypes.SetString(store, noorsignaltypes.KeyPendingMint, value)
}
