package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper is the minimal keeper for the x/noorsignal (PoSS) module.
//
// At this stage, it handles:
// - codec (for future state encoding/decoding),
// - storeKey (access to the KVStore),
// - PoSS Params via x/params Subspace,
// - simple daily counters for PoSS signals,
// - daily reward tracking for MaxRewardPerDay,
// - a thin wrapper around the PoSS Params and reward helpers,
// - a PendingMint queue (planning only, no real mint),
// - GenesisState load/store helpers,
// - an internal "signal pipeline" without real minting,
// - and a high-level PoSSStats view for dashboards/CLI.
type Keeper struct {
	// Codec used to encode/decode module state.
	cdc codec.Codec

	// storeKey gives access to the module KVStore.
	storeKey storetypes.StoreKey

	// paramSpace allows PoSS Params to be stored and modified via x/params.
	paramSpace paramstypes.Subspace
}

// NewKeeper creates a new minimal PoSS Keeper.
// We will add more dependencies later (hooks, links to Bank/Staking, etc.).
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	paramSpace paramstypes.Subspace,
) Keeper {
	// Ensure the Subspace has the PoSS KeyTable registered.
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(noorsignaltypes.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
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
// GenesisState helpers (PoSS Logic 22 + 23)
// -----------------------------------------------------------------------------

// getGenesisState returns the current PoSS GenesisState from the store,
// or DefaultGenesis() if nothing is stored yet.
func (k Keeper) getGenesisState(ctx sdk.Context) noorsignaltypes.GenesisState {
	store := k.getStore(ctx)
	bz := store.Get(noorsignaltypes.KeyGenesisState)
	if len(bz) == 0 {
		return *noorsignaltypes.DefaultGenesis()
	}

	var gs noorsignaltypes.GenesisState
	if err := json.Unmarshal(bz, &gs); err != nil {
		panic(err)
	}

	return gs
}

// setGenesisState validates and stores the PoSS GenesisState as JSON.
func (k Keeper) setGenesisState(ctx sdk.Context, gs noorsignaltypes.GenesisState) {
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

// InitGenesis stores the initial PoSS genesis state in the KVStore.
func (k Keeper) InitGenesis(ctx sdk.Context, gs noorsignaltypes.GenesisState) {
	k.setGenesisState(ctx, gs)
}

// ExportGenesis reads the PoSS genesis-equivalent state from the KVStore.
func (k Keeper) ExportGenesis(ctx sdk.Context) noorsignaltypes.GenesisState {
	return k.getGenesisState(ctx)
}

// -----------------------------------------------------------------------------
// High-level PoSS stats view (PoSS Logic 24)
// -----------------------------------------------------------------------------

// GetGlobalStats returns a consolidated, read-only view of the PoSS
// global state, meant for CLI and dashboards.
//
// It combines:
// - GenesisState counters (TotalSignals, TotalMinted),
// - PoSS Params (PoSSEnabled, daily limits, reserve denom).
func (k Keeper) GetGlobalStats(ctx sdk.Context) noorsignaltypes.PoSSStats {
	gs := k.getGenesisState(ctx)
	params := k.GetParams(ctx)

	return noorsignaltypes.PoSSStats{
		TotalSignals:               gs.TotalSignals,
		TotalMinted:                gs.TotalMinted,
		PoSSEnabled:                params.PoSSEnabled,
		MaxSignalsPerDay:           params.MaxSignalsPerDay,
		MaxSignalsPerCuratorPerDay: params.MaxSignalsPerCuratorPerDay,
		PoSSReserveDenom:           params.PoSSReserveDenom,
	}
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
func (k Keeper) IncrementDailySignalsCount(ctx sdk.Context, address, date string) uint32 {
	current := k.GetDailySignalsCount(ctx, address, date)
	next := current + 1
	k.SetDailySignalsCount(ctx, address, date, next)
	return next
}

// -----------------------------------------------------------------------------
// Daily reward tracking (per participant, per day) — MaxRewardPerDay
// -----------------------------------------------------------------------------

// dailyRewardKey builds the store key used to track how much reward
// a participant has already received for a given date.
func dailyRewardKey(address, date string) []byte {
	// Simple debug-friendly format:
	// "daily_reward:<address>:<date>"
	return []byte("daily_reward:" + address + ":" + date)
}

// getDailyRewardAmount returns the total PoSS reward (participant side)
// accumulated so far for (address, date), as sdk.Int.
//
// Internally we store it as a big-endian uint64, which is largement
// suffisant pour les caps PoSS v1 (BaseReward et MaxRewardPerDay petits).
func (k Keeper) getDailyRewardAmount(ctx sdk.Context, address, date string) sdk.Int {
	store := k.getStore(ctx)
	bz := store.Get(dailyRewardKey(address, date))
	if len(bz) == 0 {
		return sdk.ZeroInt()
	}

	amt := sdk.BigEndianToUint64(bz)
	return sdk.NewIntFromUint64(amt)
}

// setDailyRewardAmount writes the total participant reward for (address, date).
func (k Keeper) setDailyRewardAmount(ctx sdk.Context, address, date string, amount sdk.Int) {
	if amount.IsNegative() {
		amount = sdk.ZeroInt()
	}
	store := k.getStore(ctx)
	store.Set(dailyRewardKey(address, date), sdk.Uint64ToBigEndian(amount.Uint64()))
}

// -----------------------------------------------------------------------------
// Params & reward helpers (PoSS Logic 11 + ParamSubspace)
// -----------------------------------------------------------------------------

// SetParams enregistre des Params PoSS dans le ParamSubspace.
//
// Les Params sont validés avant enregistrement.
func (k Keeper) SetParams(ctx sdk.Context, params noorsignaltypes.Params) {
	if err := params.Validate(); err != nil {
		panic(fmt.Errorf("invalid PoSS params: %w", err))
	}
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetParams retourne les Params PoSS stockés dans le ParamSubspace.
//
// Comportement de sécurité :
// - si aucun param n'est encore stocké (Subspace vide), on écrit DefaultParams()
//   dans le store, puis on les retourne,
// - si les Params stockés sont invalides, on retourne DefaultParams()
//   (sans panic, pour éviter de bloquer la chaîne).
func (k Keeper) GetParams(ctx sdk.Context) noorsignaltypes.Params {
	var params noorsignaltypes.Params

	// Premier lancement / aucun param en store :
	// on teste la présence d'au moins une clé (PoSSReserveDenom).
	if !k.paramSpace.Has(ctx, noorsignaltypes.KeyPoSSReserveDenom) {
		params = noorsignaltypes.DefaultParams()
		k.SetParams(ctx, params)
		return params
	}

	// Ici on sait qu'il y a quelque chose dans le store,
	// on peut donc appeler GetParamSet sans risque de panic "empty bytes".
	k.paramSpace.GetParamSet(ctx, &params)

	if err := params.Validate(); err != nil {
		// En cas de problème, on revient aux defaults "safe off".
		return noorsignaltypes.DefaultParams()
	}

	return params
}

// ComputeSignalRewardForBlock is a thin wrapper around the pure helpers
// in types/rewards.go. It:
//
//   1) fetches PoSS Params (currently DefaultParams on first run),
//   2) uses the current block height from the context,
//   3) calls ComputeSignalReward (base * weight -> halving -> 70/30 split).
//
// It DOES NOT:
//   - check daily limits,
//   - check balances in the PoSS reserve,
//   - persist anything in the store.
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
// Internal signal pipeline (PoSS Logic 19 + 21 + 23 — without minting)
// -----------------------------------------------------------------------------

// ProcessSignalInternal is the first internal pipeline for a PoSS signal.
//
// It is intentionally LIMITED:
// - it computes the raw reward (participant / curator),
// - it enforces MaxRewardPerDay for the participant,
// - it increments the participant's daily counter,
// - it records a PendingMint entry for later processing,
// - it increments the global Genesis totals (TotalSignals / TotalMinted),
// - it returns the rewards to the caller,
// - it does NOT:
//   * enforce signals-per-day limits (cela viendra plus tard),
//   * move any real coins in Bank.
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

	// 2) Apply MaxRewardPerDay for the participant (if configured).
	//
	// Règle choisie (PoSS v1):
	// - si MaxRewardPerDay.Amount == 0 : pas de limite de volume (désactivée),
	// - sinon :
	//   * tant que la somme des rewards journaliers du participant <= MaxRewardPerDay,
	//     on crédite normalement,
	//   * dès qu'un signal ferait dépasser le plafond, ce signal donne 0/0
	//     (participantReward=0, curatorReward=0), on compte quand même le signal.
	params := k.GetParams(ctx)
	maxDaily := params.MaxRewardPerDay.Amount

	if !participantReward.Amount.IsZero() && !maxDaily.IsZero() {
		current := k.getDailyRewardAmount(ctx, participantAddr, date)

		// Si déjà au plafond ou au-dessus → ce signal ne donne plus rien.
		if current.GTE(maxDaily) {
			participantReward = sdk.NewCoin(participantReward.Denom, sdk.ZeroInt())
			curatorReward = sdk.NewCoin(curatorReward.Denom, sdk.ZeroInt())
		} else {
			// Reward potentiel pour ce signal.
			newTotal := current.Add(participantReward.Amount)

			// Si ce signal ferait dépasser le plafond, on coupe tout pour ce signal.
			if newTotal.GT(maxDaily) {
				participantReward = sdk.NewCoin(participantReward.Denom, sdk.ZeroInt())
				curatorReward = sdk.NewCoin(curatorReward.Denom, sdk.ZeroInt())
			} else {
				// Sinon, on enregistre le nouveau total pour le participant.
				k.setDailyRewardAmount(ctx, participantAddr, date, newTotal)
			}
		}
	}

	// 3) Increment participant daily counter for this date.
	k.IncrementDailySignalsCount(ctx, participantAddr, date)

	// 4) Record a PendingMint entry (planning only).
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

	// 5) Update global PoSS totals in GenesisState.
	gs := k.getGenesisState(ctx)
	gs.TotalSignals++

	// TotalMinted is stored as a string; we convert it to Int, add rewards, and store back.
	totalInt, ok := sdk.NewIntFromString(gs.TotalMinted)
	if !ok {
		return sdk.Coin{}, sdk.Coin{}, fmt.Errorf("invalid TotalMinted in PoSS genesis: %s", gs.TotalMinted)
	}

	// Sum both rewards (they can both be zero if PoSS is disabled or capped).
	sum := participantReward.Amount.Add(curatorReward.Amount)
	totalInt = totalInt.Add(sum)

	gs.TotalMinted = totalInt.String()
	k.setGenesisState(ctx, gs)

	// 6) Return rewards to the caller.
	return participantReward, curatorReward, nil
}
