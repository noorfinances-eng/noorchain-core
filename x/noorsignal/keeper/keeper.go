package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper is the minimal keeper for the x/noorsignal (PoSS) module.
//
// At this stage, it handles only:
// - codec (for future state encoding/decoding),
// - storeKey (access to the KVStore),
// - simple daily counters for PoSS signals,
// - a thin wrapper around the PoSS Params and reward helpers,
// - a basic daily reset mechanism (PoSS Logic 7 style, ready for later use).
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
// Daily counters (per address, per day) — current model
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
func (k Keeper) IncrementDailySignalsCount(ctx sdk.Context, address, date string) uint32 {
	current := k.GetDailySignalsCount(ctx, address, date)
	next := current + 1
	k.SetDailySignalsCount(ctx, address, date, next)
	return next
}

// -----------------------------------------------------------------------------
// Legacy-style daily reset helpers (PoSS Logic 7 style)
// -----------------------------------------------------------------------------
//
// NOTE: avec le modèle (address, "YYYY-MM-DD"), on n’a PLUS besoin de vider
// tous les compteurs chaque jour : la clé change de toute façon.
// MAIS on garde ce mécanisme pour être compatible avec notre plan PoSS Logic 7
// et avec les clés KeyLastResetDay / prefixes, au cas où on les utilise plus tard.

// getCurrentDay returns the current day as an integer
// (Unix timestamp in seconds divided by 86400).
func (k Keeper) getCurrentDay(ctx sdk.Context) int64 {
	return ctx.BlockTime().Unix() / 86400
}

// GetLastResetDay returns the last day (Unix days) when the daily counters were reset.
// If it has never been set, it returns 0.
func (k Keeper) GetLastResetDay(ctx sdk.Context) int64 {
	store := k.getStore(ctx)
	bz := store.Get(noorsignaltypes.KeyLastResetDay)
	if len(bz) == 0 {
		return 0
	}

	return int64(sdk.BigEndianToUint64(bz))
}

// setLastResetDay stores the last reset day as a uint64.
func (k Keeper) setLastResetDay(ctx sdk.Context, day int64) {
	store := k.getStore(ctx)
	bz := sdk.Uint64ToBigEndian(uint64(day))
	store.Set(noorsignaltypes.KeyLastResetDay, bz)
}

// ResetDailyCountersIfNeeded resets the legacy-style daily counters
// if a new day has started since the last reset.
//
// Avec notre modèle (address, date), cette fonction est quasi "no-op"
// (car on ne remplit pas encore les prefixes), mais elle est prête
// pour PoSS Logic avancé et est appelée dans BeginBlock.
func (k Keeper) ResetDailyCountersIfNeeded(ctx sdk.Context) {
	currentDay := k.getCurrentDay(ctx)
	lastResetDay := k.GetLastResetDay(ctx)

	// Same day → nothing to do.
	if lastResetDay == currentDay {
		return
	}

	store := k.getStore(ctx)

	// Participant daily counters (legacy prefix)
	pStore := prefix.NewStore(store, noorsignaltypes.KeyPrefixParticipantDailyCount)
	pIter := pStore.Iterator(nil, nil)
	defer pIter.Close()

	for ; pIter.Valid(); pIter.Next() {
		pStore.Delete(pIter.Key())
	}

	// Curator daily counters (legacy prefix)
	cStore := prefix.NewStore(store, noorsignaltypes.KeyPrefixCuratorDailyCount)
	cIter := cStore.Iterator(nil, nil)
	defer cIter.Close()

	for ; cIter.Valid(); cIter.Next() {
		cStore.Delete(cIter.Key())
	}

	// Update last reset day
	k.setLastResetDay(ctx, currentDay)
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
// in types/rewards.go.
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
