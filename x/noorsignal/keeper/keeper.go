package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper is the minimal keeper for the x/noorsignal (PoSS) module.
//
// At this stage, it handles:
// - codec (for future state encoding/decoding),
// - storeKey (access to the KVStore),
// - simple daily counters for PoSS signals,
// - thin wrappers around the PoSS Params and reward helpers,
// - global PoSS counters (TotalSignals / TotalMinted).
type Keeper struct {
	// Codec used to encode/decode module state (for future use).
	cdc codec.Codec

	// storeKey gives access to the module KVStore.
	storeKey storetypes.StoreKey
}

// NewKeeper creates a new minimal PoSS Keeper.
// We will add more dependencies later (params, hooks, links to Bank/Staking, etc.).
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// -----------------------------------------------------------------------------
// Internal store helpers
// -----------------------------------------------------------------------------

// getStore is a tiny helper to access the PoSS KVStore from a context.
func (k Keeper) getStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

// -----------------------------------------------------------------------------
// Global PoSS counters (TotalSignals / TotalMinted)
// -----------------------------------------------------------------------------
//
// NOTE:
// - We store both as big-endian uint64 in the KVStore.
// - This is enough for NOORCHAIN cap (299,792,458 NUR with reasonable decimals).
// - Genesis wiring will be done later; for now, empty store = zero.

// GetTotalSignals returns the global number of PoSS signals processed so far.
func (k Keeper) GetTotalSignals(ctx sdk.Context) uint64 {
	store := k.getStore(ctx)
	bz := store.Get(noorsignaltypes.KeyTotalSignals)
	if len(bz) == 0 {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

// SetTotalSignals sets the global number of PoSS signals.
func (k Keeper) SetTotalSignals(ctx sdk.Context, value uint64) {
	store := k.getStore(ctx)
	store.Set(noorsignaltypes.KeyTotalSignals, sdk.Uint64ToBigEndian(value))
}

// IncrementTotalSignals increments the global PoSS signal counter by 1
// and returns the new value.
func (k Keeper) IncrementTotalSignals(ctx sdk.Context) uint64 {
	current := k.GetTotalSignals(ctx)
	next := current + 1
	k.SetTotalSignals(ctx, next)
	return next
}

// GetTotalMinted returns the global amount of NUR minted via PoSS (in unur).
func (k Keeper) GetTotalMinted(ctx sdk.Context) uint64 {
	store := k.getStore(ctx)
	bz := store.Get(noorsignaltypes.KeyTotalMinted)
	if len(bz) == 0 {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

// SetTotalMinted sets the global amount of NUR minted via PoSS (in unur).
func (k Keeper) SetTotalMinted(ctx sdk.Context, value uint64) {
	store := k.getStore(ctx)
	store.Set(noorsignaltypes.KeyTotalMinted, sdk.Uint64ToBigEndian(value))
}

// AddToTotalMinted adds `amount` (in unur) to the global PoSS minted total
// and returns the new value.
func (k Keeper) AddToTotalMinted(ctx sdk.Context, amount uint64) uint64 {
	if amount == 0 {
		return k.GetTotalMinted(ctx)
	}

	current := k.GetTotalMinted(ctx)
	next := current + amount

	// Basic overflow protection (should never happen with NOOR cap).
	if next < current {
		// In case of overflow, we keep the current value and do not update.
		return current
	}

	k.SetTotalMinted(ctx, next)
	return next
}

// -----------------------------------------------------------------------------
// Daily counters (per address, per day)
// -----------------------------------------------------------------------------

// GetDailySignalsCount returns how many PoSS signals have already been
// recorded for a given (address, date) pair.
//
// - address: NOOR bech32 address (noor1...)
// - date: ISO day string "YYYY-MM-DD"
func (k Keeper) GetDailySignalsCount(ctx sdk.Context, address, date string) uint32 {
	store := k.getStore(ctx)
	bz := store.Get(noorsignaltypes.DailyCounterKey(address, date))
	if len(bz) == 0 {
		return 0
	}

	// We store counters as big-endian uint64; convert back to uint32.
	return uint32(sdk.BigEndianToUint64(bz))
}

// SetDailySignalsCount sets the PoSS daily counter for (address, date).
func (k Keeper) SetDailySignalsCount(ctx sdk.Context, address, date string, count uint32) {
	store := k.getStore(ctx)
	store.Set(
		noorsignaltypes.DailyCounterKey(address, date),
		sdk.Uint64ToBigEndian(uint64(count)),
	)
}

// IncrementDailySignalsCount increments the daily counter for (address, date)
// and returns the new value.
//
// This function will be used later when processing a PoSS signal:
// - read the current count
// - +1
// - store it back
// - check against MaxSignalsPerDay
func (k Keeper) IncrementDailySignalsCount(ctx sdk.Context, address, date string) uint32 {
	current := k.GetDailySignalsCount(ctx, address, date)
	next := current + 1
	k.SetDailySignalsCount(ctx, address, date, next)
	return next
}

// -----------------------------------------------------------------------------
// Params & reward helpers (PoSS Logic 11)
// -----------------------------------------------------------------------------

// GetParams returns the current PoSS params.
//
// For now, we simply return DefaultParams(), which means:
// - PoSS is effectively configured off-chain in code,
// - later we will plug this into x/params with a subspace
//   and make everything adjustable by governance.
func (k Keeper) GetParams(ctx sdk.Context) noorsignaltypes.Params {
	_ = ctx // context will be useful once we use a ParamSubspace
	return noorsignaltypes.DefaultParams()
}

// ComputeSignalRewardForBlock is a thin wrapper around the pure helpers
// in types/rewards.go. It:
//
//   1) fetches PoSS Params (currently DefaultParams),
//   2) uses the current block height from the context,
//   3) calls ComputeSignalReward (base * weight -> halving -> 70/30 split).
//
// It DOES NOT:
//   - check daily limits,
//   - check balances in the PoSS reserve,
//   - persist anything in the store.
//
// All those checks and state updates will be implemented in later PoSS Logic
// steps inside the Keeper.
func (k Keeper) ComputeSignalRewardForBlock(
	ctx sdk.Context,
	signalType noorsignaltypes.SignalType,
) (participant sdk.Coin, curator sdk.Coin, err error) {
	params := k.GetParams(ctx)
	height := ctx.BlockHeight()

	return noorsignaltypes.ComputeSignalReward(
		params,
		signalType,
		height,
	)
}
