package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// -----------------------------------------------------------------------------
// GLOBAL COUNTERS (PoSS Logic 27)
// TotalSignals / TotalMinted updated progressively by the Keeper
// -----------------------------------------------------------------------------

// GetTotalSignals returns the global PoSS signal counter stored in KV.
func (k Keeper) GetTotalSignals(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(noorsignaltypes.KeyTotalSignals)
	if len(bz) == 0 {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

// IncreaseTotalSignals increments TotalSignals by n.
func (k Keeper) IncreaseTotalSignals(ctx sdk.Context, n uint64) {
	total := k.GetTotalSignals(ctx)
	newTotal := total + n

	store := ctx.KVStore(k.storeKey)
	store.Set(noorsignaltypes.KeyTotalSignals, sdk.Uint64ToBigEndian(newTotal))
}

// -----------------------------------------------------------------------------
// TOTAL MINTED (string amount in "unur")
// -----------------------------------------------------------------------------

// GetTotalMinted returns the total minted supply stored in genesis style.
// Value is stored as a STRING inside KV, not Coin.
func (k Keeper) GetTotalMinted(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(noorsignaltypes.KeyTotalMinted)
	if len(bz) == 0 {
		return "0"
	}
	return string(bz)
}

// IncreaseTotalMinted adds a reward Coin to the global supply counter.
//
// Example: TotalMinted = "351"
//          reward = 12unur
//          => new TotalMinted = "363"
func (k Keeper) IncreaseTotalMinted(ctx sdk.Context, coin sdk.Coin) {
	if !coin.Amount.IsPositive() {
		return
	}

	current := k.GetTotalMinted(ctx)
	curInt, ok := sdk.NewIntFromString(current)
	if !ok {
		// fallback: reset to 0 if corrupted
		curInt = sdk.ZeroInt()
	}

	newInt := curInt.Add(coin.Amount)

	store := ctx.KVStore(k.storeKey)
	store.Set(noorsignaltypes.KeyTotalMinted, []byte(newInt.String()))
}
