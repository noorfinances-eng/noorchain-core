package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper is the minimal keeper for the x/noorsignal (PoSS) module.
// In PoSS Logic 7 we add the daily reset logic and internal helpers.
type Keeper struct {
	// Codec used to (de)serialize state (kept for future use).
	cdc codec.Codec

	// storeKey gives access to the module KVStore.
	storeKey storetypes.StoreKey
}

// NewKeeper creates a new minimal PoSS keeper.
//
// Later we will extend it with params, hooks and links to Bank/Staking.
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
// Internal helpers
// -----------------------------------------------------------------------------

// getStore returns the KVStore for the PoSS module.
func (k Keeper) getStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

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

// -----------------------------------------------------------------------------
// Daily reset logic (PoSS Logic 7)
// -----------------------------------------------------------------------------

// ResetDailyCountersIfNeeded resets the daily participant/curator counters
// if a new day has started since the last reset.
//
// This is called from BeginBlock of the module and is fully automatic:
// - no cron
// - no external script
// - pure on-chain logic.
func (k Keeper) ResetDailyCountersIfNeeded(ctx sdk.Context) {
	currentDay := k.getCurrentDay(ctx)
	lastResetDay := k.GetLastResetDay(ctx)

	// Same day → nothing to do.
	if lastResetDay == currentDay {
		return
	}

	// New day → clear all daily counters.
	store := k.getStore(ctx)

	// Participant daily counters
	pStore := prefix.NewStore(store, noorsignaltypes.KeyPrefixParticipantDailyCount)
	pIter := pStore.Iterator(nil, nil)
	defer pIter.Close()

	for ; pIter.Valid(); pIter.Next() {
		pStore.Delete(pIter.Key())
	}

	// Curator daily counters
	cStore := prefix.NewStore(store, noorsignaltypes.KeyPrefixCuratorDailyCount)
	cIter := cStore.Iterator(nil, nil)
	defer cIter.Close()

	for ; cIter.Valid(); cIter.Next() {
		cStore.Delete(cIter.Key())
	}

	// Update last reset day
	k.setLastResetDay(ctx, currentDay)
}
