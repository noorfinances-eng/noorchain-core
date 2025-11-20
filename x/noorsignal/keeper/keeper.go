package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper est le gestionnaire principal du module PoSS (noorsignal).
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.Codec

	// 🔥 BankKeeper ajouté ici → transferts réels NUR
	BankKeeper bankkeeper.Keeper
}

// NewKeeper construit un nouveau Keeper PoSS pour NOORCHAIN.
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	bk bankkeeper.Keeper,
) Keeper {
	return Keeper{
		storeKey:    storeKey,
		cdc:         cdc,
		BankKeeper:  bk,
	}
}

// getStore retourne le KVStore du module.
func (k Keeper) getStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

// ---------- Stores préfixés ----------

func (k Keeper) signalStore(ctx sdk.Context) prefix.Store {
	return noorsignaltypes.GetSignalStore(k.getStore(ctx))
}

func (k Keeper) curatorStore(ctx sdk.Context) prefix.Store {
	return noorsignaltypes.GetCuratorStore(k.getStore(ctx))
}

func (k Keeper) configStore(ctx sdk.Context) prefix.Store {
	return noorsignaltypes.GetConfigStore(k.getStore(ctx))
}

func (k Keeper) dailyCounterStore(ctx sdk.Context) prefix.Store {
	return noorsignaltypes.GetDailyCounterStore(k.getStore(ctx))
}

// ---------------------------
// Gestion de la configuration
// ---------------------------

func (k Keeper) SetConfig(ctx sdk.Context, cfg noorsignaltypes.PossConfig) {
	store := k.configStore(ctx)
	bz := k.cdc.MustMarshal(&cfg)
	store.Set([]byte("config"), bz)
}

func (k Keeper) GetConfig(ctx sdk.Context) (noorsignaltypes.PossConfig, bool) {
	store := k.configStore(ctx)

	bz := store.Get([]byte("config"))
	if bz == nil {
		return noorsignaltypes.PossConfig{}, false
	}

	var cfg noorsignaltypes.PossConfig
	k.cdc.MustUnmarshal(bz, &cfg)
	return cfg, true
}

func (k Keeper) InitDefaultConfig(ctx sdk.Context) {
	_, found := k.GetConfig(ctx)
	if found {
		return
	}
	k.SetConfig(ctx, noorsignaltypes.DefaultPossConfig())
}

// -------------------------------------
// Récompenses PoSS (calcul)
// -------------------------------------

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
// Gestion des identifiants de signaux
// -------------------------------------

func (k Keeper) getNextSignalID(ctx sdk.Context) uint64 {
	bz := k.getStore(ctx).Get(noorsignaltypes.KeyNextSignalID)
	if bz == nil {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextSignalID(ctx sdk.Context, nextID uint64) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, nextID)
	k.getStore(ctx).Set(noorsignaltypes.KeyNextSignalID, bz)
}

func (k Keeper) CreateSignal(ctx sdk.Context, sig noorsignaltypes.Signal) noorsignaltypes.Signal {
	nextID := k.getNextSignalID(ctx)
	sig.Id = nextID

	store := k.signalStore(ctx)
	bz := k.cdc.MustMarshal(&sig)
	store.Set(noorsignaltypes.SignalKey(sig.Id), bz)

	k.setNextSignalID(ctx, nextID+1)
	return sig
}

func (k Keeper) SetSignal(ctx sdk.Context, sig noorsignaltypes.Signal) {
	bz := k.cdc.MustMarshal(&sig)
	k.signalStore(ctx).Set(noorsignaltypes.SignalKey(sig.Id), bz)
}

func (k Keeper) GetSignal(ctx sdk.Context, id uint64) (noorsignaltypes.Signal, bool) {
	bz := k.signalStore(ctx).Get(noorsignaltypes.SignalKey(id))
	if bz == nil {
		return noorsignaltypes.Signal{}, false
	}

	var sig noorsignaltypes.Signal
	k.cdc.MustUnmarshal(bz, &sig)
	return sig, true
}

// -------------------------------------
// Limite quotidienne PoSS
// -------------------------------------

func (k Keeper) getDailySignalCount(ctx sdk.Context, addr sdk.AccAddress, dayBucket uint64) uint32 {
	bz := k.dailyCounterStore(ctx).Get(noorsignaltypes.DailyCounterKey(addr, dayBucket))
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint32(bz)
}

func (k Keeper) incrementDailySignalCount(ctx sdk.Context, addr sdk.AccAddress, dayBucket uint64) uint32 {
	current := k.getDailySignalCount(ctx, addr, dayBucket)
	next := current + 1
	bz := make([]byte, 4)
	binary.BigEndian.PutUint32(bz, next)
	k.dailyCounterStore(ctx).Set(noorsignaltypes.DailyCounterKey(addr, dayBucket), bz)
	return next
}

// -------------------------------------
// Gestion des Curators
// -------------------------------------

func (k Keeper) SetCurator(ctx sdk.Context, curator noorsignaltypes.Curator) {
	bz := k.cdc.MustMarshal(&curator)
	k.curatorStore(ctx).Set(curator.Address.Bytes(), bz)
}

func (k Keeper) GetCurator(ctx sdk.Context, addr sdk.AccAddress) (noorsignaltypes.Curator, bool) {
	bz := k.curatorStore(ctx).Get(addr.Bytes())
	if bz == nil {
		return noorsignaltypes.Curator{}, false
	}

	var curator noorsignaltypes.Curator
	k.cdc.MustUnmarshal(bz, &curator)
	return curator, true
}

func (k Keeper) IsActiveCurator(ctx sdk.Context, addr sdk.AccAddress) bool {
	c, ok := k.GetCurator(ctx, addr)
	return ok && c.Active
}

func (k Keeper) IncrementCuratorValidatedCount(ctx sdk.Context, addr sdk.AccAddress) {
	c, ok := k.GetCurator(ctx, addr)
	if !ok {
		return
	}
	c.TotalSignalsValidated++
	k.SetCurator(ctx, c)
}
