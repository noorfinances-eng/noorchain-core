package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper is the minimal keeper for the x/noorsignal (PoSS) module.
//
// For PoSS Logic 1 – Step B, it only manages a single global counter:
// TotalSignals (stored under KeyTotalSignals in the KVStore).
type Keeper struct {
	// cdc is the codec used to (un)marshal state if needed.
	cdc codec.Codec

	// storeKey gives access to the module's KVStore.
	storeKey storetypes.StoreKey
}

// NewKeeper creates a new PoSS Keeper instance.
//
// In future steps, we will extend this keeper with:
// - params
// - references to Bank / Staking / EVM
// - PoSS reward and limit logic.
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// getStore returns the module KVStore for the given context.
func (k Keeper) getStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

// -----------------------------------------------------------------------------
// Global TotalSignals counter
// -----------------------------------------------------------------------------

// GetTotalSignals returns the current global number of validated PoSS signals.
//
// If the value has never been set, it returns 0.
func (k Keeper) GetTotalSignals(ctx sdk.Context) uint64 {
	store := k.getStore(ctx)
	bz := store.Get(noorsignaltypes.KeyTotalSignals)
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

// SetTotalSignals sets the global TotalSignals counter to the given value.
//
// This is a low-level method; in most cases, we will use IncrementTotalSignals.
func (k Keeper) SetTotalSignals(ctx sdk.Context, total uint64) {
	store := k.getStore(ctx)
	bz := sdk.Uint64ToBigEndian(total)
	store.Set(noorsignaltypes.KeyTotalSignals, bz)
}

// IncrementTotalSignals increments the global TotalSignals counter by `delta`
// and returns the new value.
//
// NOTE:
// - This method does not implement any limit or anti-abuse logic yet.
// - All PoSS rules (daily caps, per-curator limits, etc.) will be added in
//   future steps when we introduce real PoSS messages and rewards.
func (k Keeper) IncrementTotalSignals(ctx sdk.Context, delta uint64) uint64 {
	current := k.GetTotalSignals(ctx)
	newTotal := current + delta
	k.SetTotalSignals(ctx, newTotal)
	return newTotal
}
