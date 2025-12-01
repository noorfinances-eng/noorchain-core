package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
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
// - a PendingMint queue (PoSS Logic 21) that records planned rewards,
// - and simple GenesisState load/store helpers (PoSS Logic 22).
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
// GenesisState helpers (PoSS Logic 22)
// -----------------------------------------------------------------------------

// InitGenesis stores the initial PoSS genesis state in the KVStore.
//
// For now, this is a simple JSON blob under KeyGenesisState.
// Later, we can split it into separate keys if needed.
func (k Keeper) InitGenesis(ctx sdk.Context, gs noorsignaltypes.GenesisState) {
	if err := noorsignaltypes.ValidateGenesis(&gs); err != nil {
		panic(fmt.Errorf("invalid PoSS genesis: %w", err))
	}

	store := k.getStore(ctx)
	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}

	store.Set(noorsignaltypes.KeyGenesisState, bz)
}

// ExportGenesis reads the PoSS genesis-equivalent state from the KVStore.
//
// If nothing was stored yet, it falls back to DefaultGenesis().
func (k Keeper) ExportGenesis(ctx sdk.Context) noorsignaltypes.GenesisState {
	store := k.getStore(ctx)
	bz := store.Get(noorsignaltypes.KeyGenesisState)
	if len(bz) == 0 {
		// Fresh chain or no PoSS state stored yet.
		return *noorsignaltypes.DefaultGenesis()
	}

	var gs noorsignaltypes.GenesisState
	if err := json.Unmarshal(bz, &gs); err != nil {
		panic(err)
	}

	return gs
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

// -----------------------------------------------------------------------------
// Pending mint queue (PoSS Logic 21 — planning only, no real mint)
// -----------------------------------------------------------------------------

// RecordPendingMint stores a PendingMint entry in the PoSS KVStore.
//
// IMPORTANT:
// - This does NOT mint any coins.
// - This does NOT move any real balances.
// - It is purely an internal "to be minted later" record.
//
// Key format (simple, debug-friendly):
//   "pending_mint:<height>:<participant>:<timestamp_nano>"
func (k Keeper) RecordPendingMint(
	ctx sdk.Context,
	participantAddr string,
	curatorAddr string,
	signalType noorsignaltypes.SignalType,
	participantReward sdk.Coin,
	curatorReward sdk.Coin,
) error {
	store := k.getStore(ctx)

	pending := noorsignaltypes.PendingMint{
		BlockHeight:       ctx.BlockHeight(),
		Timestamp:         ctx.BlockTime(),
		Participant:       participantAddr,
		Curator:           curatorAddr,
		SignalType:        signalType,
		ParticipantReward: participantReward,
		CuratorReward:     curatorReward,
	}

	bz, err := json.Marshal(pending)
	if err != nil {
		return err
	}

	key := []byte(fmt.Sprintf(
		"pending_mint:%d:%s:%d",
		ctx.BlockHeight(),
		participantAddr,
		ctx.BlockTime().UnixNano(),
	))

	store.Set(key, bz)
	return nil
}

// -----------------------------------------------------------------------------
// Internal signal pipeline (PoSS Logic 19 + 21 — without minting)
// -----------------------------------------------------------------------------

// ProcessSignalInternal is the first internal pipeline for a PoSS signal.
//
// It is intentionally LIMITED:
// - it computes the raw reward (participant / curator),
// - it increments the participant's daily counter,
// - it records a PendingMint entry for later processing,
// - it returns the rewards to the caller,
// - it does NOT:
//   * enforce daily limits yet (MaxSignalsPerDay, MaxSignalsPerCuratorPerDay),
//   * update TotalSignals / TotalMinted in genesis,
//   * move any real coins in Bank.
//
// Parameters:
// - participantAddr: bech32 NOOR account receiving the 70 % part later.
// - curatorAddr:     bech32 NOOR curator account receiving the 30 % part later.
// - signalType:      type of signal (micro-donation, participation, content, CCN...).
// - date:            ISO date string ("YYYY-MM-DD") for the daily counter.
//
// This function is the safe “sandbox step” before we wire actual minting
// and state changes in future PoSS Logic blocks.
func (k Keeper) ProcessSignalInternal(
	ctx sdk.Context,
	participantAddr string,
	curatorAddr string,
	signalType noorsignaltypes.SignalType,
	date string,
) (sdk.Coin, sdk.Coin, error) {
	// 1) Compute the raw PoSS rewards for this signal at this block height.
	participantReward, curatorReward, err := k.ComputeSignalRewardForBlock(ctx, signalType)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}

	// 2) Increment participant daily counter for this date.
	//
	// NOTE:
	// - We do NOT enforce MaxSignalsPerDay yet.
	// - We do NOT touch curator counters yet.
	k.IncrementDailySignalsCount(ctx, participantAddr, date)

	// 3) Record a PendingMint entry (planning only).
	if err := k.RecordPendingMint(
		ctx,
		participantAddr,
		curatorAddr,
		signalType,
		participantReward,
		curatorReward,
	); err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}

	// 4) Return rewards to the caller.
	// Later, the caller will:
	// - check limits,
	// - check PoSS reserve,
	// - actually credit participant & curator balances.
	return participantReward, curatorReward, nil
}
