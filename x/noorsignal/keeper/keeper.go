package keeper

import (
	"encoding/binary"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/store/prefix"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper est le gestionnaire principal du module PoSS (noorsignal) pour NOORCHAIN.
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.Codec
}

// NewKeeper construit un nouveau Keeper PoSS pour NOORCHAIN.
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// getStore retourne le KVStore brut du module à partir du contexte.
func (k Keeper) getStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

// ---------- Stores préfixés ----------

func (k Keeper) signalStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetSignalStore(parent)
}

func (k Keeper) curatorStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetCuratorStore(parent)
}

func (k Keeper) configStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetConfigStore(parent)
}

func (k Keeper) dailyCounterStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetDailyCounterStore(parent)
}

// ---------------------------
// Gestion de la configuration
// ---------------------------

func (k Keeper) SetConfig(ctx sdk.Context, cfg noorsignaltypes.PossConfig) {
	store := k.configStore(ctx)

	bz, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	store.Set([]byte("config"), bz)
}

func (k Keeper) GetConfig(ctx sdk.Context) (noorsignaltypes.PossConfig, bool) {
	store := k.configStore(ctx)

	bz := store.Get([]byte("config"))
	if bz == nil {
		return noorsignaltypes.PossConfig{}, false
	}

	var cfg noorsignaltypes.PossConfig
	if err := json.Unmarshal(bz, &cfg); err != nil {
		// Si la config est corrompue, on préfère paniquer en dev.
		panic(err)
	}

	return cfg, true
}

func (k Keeper) InitDefaultConfig(ctx sdk.Context) {
	_, found := k.GetConfig(ctx)
	if found {
		return
	}

	defaultCfg := noorsignaltypes.DefaultPossConfig()
	k.SetConfig(ctx, defaultCfg)
}

// ---------------------------------
// Calcul des récompenses PoSS (aide)
// ---------------------------------

func (k Keeper) ComputeSignalRewardsFromConfig(
	ctx sdk.Context,
	weight uint32,
	era uint32,
) (total uint64, participant uint64, curator uint64, found bool) {
	cfg, ok := k.GetConfig(ctx)
	if !ok {
		return 0, 0, 0, false
	}

	total, participant, curator = noorsignaltypes.ComputeSignalRewards(cfg, weight, era)
	return total, participant, curator, true
}

func (k Keeper) ComputeSignalRewardsCurrentEra(
	ctx sdk.Context,
	weight uint32,
) (total uint64, participant uint64, curator uint64, found bool) {
	cfg, ok := k.GetConfig(ctx)
	if !ok {
		return 0, 0, 0, false
	}

	era := cfg.EraIndex
	total, participant, curator = noorsignaltypes.ComputeSignalRewards(cfg, weight, era)
	return total, participant, curator, true
}

// -------------------------------------
// Gestion des identifiants et des signaux
// -------------------------------------

func (k Keeper) getNextSignalID(ctx sdk.Context) uint64 {
	store := k.getStore(ctx)

	bz := store.Get(noorsignaltypes.KeyNextSignalID)
	if bz == nil {
		return 1
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextSignalID(ctx sdk.Context, nextID uint64) {
	store := k.getStore(ctx)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, nextID)
	store.Set(noorsignaltypes.KeyNextSignalID, bz)
}

func (k Keeper) CreateSignal(ctx sdk.Context, sig noorsignaltypes.Signal) noorsignaltypes.Signal {
	nextID := k.getNextSignalID(ctx)
	sig.Id = nextID

	sstore := k.signalStore(ctx)
	key := noorsignaltypes.SignalKey(sig.Id)

	bz, err := json.Marshal(sig)
	if err != nil {
		panic(err)
	}
	sstore.Set(key, bz)

	k.setNextSignalID(ctx, nextID+1)
	return sig
}

func (k Keeper) SetSignal(ctx sdk.Context, sig noorsignaltypes.Signal) {
	sstore := k.signalStore(ctx)
	key := noorsignaltypes.SignalKey(sig.Id)

	bz, err := json.Marshal(sig)
	if err != nil {
		panic(err)
	}
	sstore.Set(key, bz)
}

func (k Keeper) GetSignal(ctx sdk.Context, id uint64) (noorsignaltypes.Signal, bool) {
	sstore := k.signalStore(ctx)
	key := noorsignaltypes.SignalKey(id)

	bz := sstore.Get(key)
	if bz == nil {
		return noorsignaltypes.Signal{}, false
	}

	var sig noorsignaltypes.Signal
	if err := json.Unmarshal(bz, &sig); err != nil {
		panic(err)
	}
	return sig, true
}

// -------------------------------------
// Compteurs quotidiens PoSS (limites)
// -------------------------------------

func (k Keeper) getDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
) uint32 {
	store := k.dailyCounterStore(ctx)
	key := noorsignaltypes.DailyCounterKey(addr, dayBucket)

	bz := store.Get(key)
	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint32(bz)
}

func (k Keeper) setDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
	count uint32,
) {
	store := k.dailyCounterStore(ctx)
	key := noorsignaltypes.DailyCounterKey(addr, dayBucket)

	bz := make([]byte, 4)
	binary.BigEndian.PutUint32(bz, count)
	store.Set(key, bz)
}

func (k Keeper) incrementDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
) uint32 {
	current := k.getDailySignalCount(ctx, addr, dayBucket)
	next := current + 1
	k.setDailySignalCount(ctx, addr, dayBucket, next)
	return next
}

// Helpers exportés pour le MsgServer (BankKeeper à venir).
func (k Keeper) GetDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
) uint32 {
	return k.getDailySignalCount(ctx, addr, dayBucket)
}

func (k Keeper) IncrementDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
) uint32 {
	return k.incrementDailySignalCount(ctx, addr, dayBucket)
}

// -------------------------------------
// Gestion des Curators PoSS
// -------------------------------------

func (k Keeper) curatorKey(addr sdk.AccAddress) []byte {
	return addr.Bytes()
}

func (k Keeper) SetCurator(ctx sdk.Context, curator noorsignaltypes.Curator) {
	store := k.curatorStore(ctx)
	key := k.curatorKey(curator.Address)

	bz, err := json.Marshal(curator)
	if err != nil {
		panic(err)
	}
	store.Set(key, bz)
}

func (k Keeper) GetCurator(ctx sdk.Context, addr sdk.AccAddress) (noorsignaltypes.Curator, bool) {
	store := k.curatorStore(ctx)
	key := k.curatorKey(addr)

	bz := store.Get(key)
	if bz == nil {
		return noorsignaltypes.Curator{}, false
	}

	var curator noorsignaltypes.Curator
	if err := json.Unmarshal(bz, &curator); err != nil {
		panic(err)
	}
	return curator, true
}

func (k Keeper) IsActiveCurator(ctx sdk.Context, addr sdk.AccAddress) bool {
	curator, found := k.GetCurator(ctx, addr)
	if !found {
		return false
	}
	return curator.Active
}

// IncrementCuratorValidatedCount augmente de 1 le nombre total de signaux
// validés par ce Curator, si le Curator existe.
func (k Keeper) IncrementCuratorValidatedCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
) {
	curator, found := k.GetCurator(ctx, addr)
	if !found {
		return
	}

	curator.TotalSignalsValidated++
	k.SetCurator(ctx, curator)
}
