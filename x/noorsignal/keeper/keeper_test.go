package keeper

import (
	"testing"

	dbm "github.com/tendermint/tm-db"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// setupKeeperAndContext crée un Keeper PoSS réel + un contexte en mémoire,
// avec un ParamSubspace fonctionnel, pour les tests.
func setupKeeperAndContext(t *testing.T) (Keeper, sdk.Context) {
	t.Helper()

	// --- Stores en mémoire ---
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	// Store du module noorsignal
	storeKey := storetypes.NewKVStoreKey(noorsignaltypes.StoreKey)

	// Stores pour x/params
	paramsKey := storetypes.NewKVStoreKey(paramstypes.StoreKey)
	paramsTKey := storetypes.NewTransientStoreKey(paramstypes.TStoreKey)

	cms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(paramsKey, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(paramsTKey, storetypes.StoreTypeTransient, nil)

	if err := cms.LoadLatestVersion(); err != nil {
		t.Fatalf("failed to load multistore: %v", err)
	}

	// --- Codec minimal (comme dans app.MakeEncodingConfig) ---
	amino := codec.NewLegacyAmino()
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	// --- ParamsKeeper + Subspace PoSS ---
	pk := paramskeeper.NewKeeper(
		cdc,
		amino,
		paramsKey,
		paramsTKey,
	)

	subspace := pk.Subspace(noorsignaltypes.ModuleName)

	// --- Keeper PoSS ---
	k := NewKeeper(
		cdc,
		storeKey,
		subspace,
	)

	// --- Contexte ---
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, tmlog.NewNopLogger())

	return k, ctx
}

// TestGetParams_DefaultsStored vérifie que GetParams() initialise bien
// les DefaultParams dans le ParamSubspace au premier appel.
func TestGetParams_DefaultsStored(t *testing.T) {
	k, ctx := setupKeeperAndContext(t)

	params := k.GetParams(ctx)

	if params.PoSSReserveDenom != noorsignaltypes.DefaultPoSSReserveDenom {
		t.Fatalf("expected PoSSReserveDenom=%s, got %s",
			noorsignaltypes.DefaultPoSSReserveDenom, params.PoSSReserveDenom)
	}

	if params.PoSSEnabled != noorsignaltypes.DefaultPoSSEnabled {
		t.Fatalf("expected PoSSEnabled=%v, got %v",
			noorsignaltypes.DefaultPoSSEnabled, params.PoSSEnabled)
	}

	if params.MaxSignalsPerDay != noorsignaltypes.DefaultMaxSignalsPerDay {
		t.Fatalf("expected MaxSignalsPerDay=%d, got %d",
			noorsignaltypes.DefaultMaxSignalsPerDay, params.MaxSignalsPerDay)
	}

	if params.MaxSignalsPerCuratorPerDay != noorsignaltypes.DefaultMaxSignalsPerCuratorPerDay {
		t.Fatalf("expected MaxSignalsPerCuratorPerDay=%d, got %d",
			noorsignaltypes.DefaultMaxSignalsPerCuratorPerDay, params.MaxSignalsPerCuratorPerDay)
	}

	// Second appel : doit retourner la même chose (params stockés dans le Subspace).
	params2 := k.GetParams(ctx)
	if params2.PoSSReserveDenom != params.PoSSReserveDenom {
		t.Fatalf("expected PoSSReserveDenom to be stable, got %s / %s",
			params.PoSSReserveDenom, params2.PoSSReserveDenom)
	}
}

// TestSetParams_RoundTrip vérifie qu'un SetParams suivi d'un GetParams
// renvoie bien les valeurs modifiées.
func TestSetParams_RoundTrip(t *testing.T) {
	k, ctx := setupKeeperAndContext(t)

	custom := noorsignaltypes.DefaultParams()
	custom.PoSSEnabled = true
	custom.MaxSignalsPerDay = 42
	custom.MaxSignalsPerCuratorPerDay = 200

	k.SetParams(ctx, custom)

	readBack := k.GetParams(ctx)

	if readBack.PoSSEnabled != true {
		t.Fatalf("expected PoSSEnabled=true after SetParams, got %v", readBack.PoSSEnabled)
	}
	if readBack.MaxSignalsPerDay != 42 {
		t.Fatalf("expected MaxSignalsPerDay=42 after SetParams, got %d", readBack.MaxSignalsPerDay)
	}
	if readBack.MaxSignalsPerCuratorPerDay != 200 {
		t.Fatalf("expected MaxSignalsPerCuratorPerDay=200 after SetParams, got %d", readBack.MaxSignalsPerCuratorPerDay)
	}
}

// TestProcessSignalInternal_UpdatesCountersAndGenesis vérifie que
// ProcessSignalInternal :
// - incrémente le compteur journalier,
// - stocke un TotalSignals++,
// - augmente TotalMinted du montant des rewards.
func TestProcessSignalInternal_UpdatesCountersAndGenesis(t *testing.T) {
	k, ctx := setupKeeperAndContext(t)

	// On force des params PoSS simples et activés.
	params := noorsignaltypes.DefaultParams()
	params.PoSSEnabled = true
	// BaseReward = 10 unur, poids micro_donation = 1 → total = 10, split 7/3
	params.BaseReward = sdk.NewInt64Coin(noorsignaltypes.DefaultPoSSReserveDenom, 10)
	params.MaxRewardPerDay = sdk.NewInt64Coin(noorsignaltypes.DefaultPoSSReserveDenom, 1000)
	params.WeightMicroDonation = 1
	params.WeightParticipation = 1
	params.WeightContent = 1
	params.WeightCCN = 1

	k.SetParams(ctx, params)

	participant := "noor1participant"
	curator := "noor1curator"
	date := "2025-01-01"

	// Avant : compteurs et genesis.
	initialCount := k.GetDailySignalsCount(ctx, participant, date)
	if initialCount != 0 {
		t.Fatalf("expected initial daily count = 0, got %d", initialCount)
	}

	gsBefore := k.ExportGenesis(ctx)
	if gsBefore.TotalSignals != 0 {
		t.Fatalf("expected initial TotalSignals=0, got %d", gsBefore.TotalSignals)
	}
	if gsBefore.TotalMinted != "0" {
		t.Fatalf("expected initial TotalMinted=0, got %s", gsBefore.TotalMinted)
	}

	// Exécution du pipeline interne.
	participantReward, curatorReward, err := k.ProcessSignalInternal(
		ctx,
		participant,
		curator,
		noorsignaltypes.SignalTypeMicroDonation,
		date,
	)
	if err != nil {
		t.Fatalf("ProcessSignalInternal returned error: %v", err)
	}

	// On s'assure qu'il y a bien un reward non nul (puisqu'on a activé PoSS).
	if !participantReward.Amount.IsPositive() || !curatorReward.Amount.IsPositive() {
		t.Fatalf("expected positive rewards when PoSSEnabled=true, got %s / %s",
			participantReward.String(), curatorReward.String())
	}

	// Compteur journalier incrémenté.
	newCount := k.GetDailySignalsCount(ctx, participant, date)
	if newCount != initialCount+1 {
		t.Fatalf("expected daily count = %d, got %d", initialCount+1, newCount)
	}

	// Genesis mis à jour.
	gsAfter := k.ExportGenesis(ctx)
	if gsAfter.TotalSignals != gsBefore.TotalSignals+1 {
		t.Fatalf("expected TotalSignals=%d, got %d",
			gsBefore.TotalSignals+1, gsAfter.TotalSignals)
	}

	// TotalMinted doit avoir augmenté du montant des deux rewards.
	beforeInt, ok := sdk.NewIntFromString(gsBefore.TotalMinted)
	if !ok {
		t.Fatalf("invalid TotalMinted before: %s", gsBefore.TotalMinted)
	}
	afterInt, ok := sdk.NewIntFromString(gsAfter.TotalMinted)
	if !ok {
		t.Fatalf("invalid TotalMinted after: %s", gsAfter.TotalMinted)
	}

	expectedDelta := participantReward.Amount.Add(curatorReward.Amount)
	actualDelta := afterInt.Sub(beforeInt)

	if !actualDelta.Equal(expectedDelta) {
		t.Fatalf("expected TotalMinted delta = %s, got %s",
			expectedDelta.String(), actualDelta.String())
	}
}

// TestProcessSignalInternal_RespectsMaxSignalsPerDay vérifie que lorsque
// MaxSignalsPerDay est atteint, le pipeline refuse le signal, n'incrémente
// plus le compteur, et ne modifie plus TotalSignals / TotalMinted.
func TestProcessSignalInternal_RespectsMaxSignalsPerDay(t *testing.T) {
	k, ctx := setupKeeperAndContext(t)

	// Params avec PoSS activé et MaxSignalsPerDay=1
	params := noorsignaltypes.DefaultParams()
	params.PoSSEnabled = true
	params.BaseReward = sdk.NewInt64Coin(noorsignaltypes.DefaultPoSSReserveDenom, 10)
	params.MaxRewardPerDay = sdk.NewInt64Coin(noorsignaltypes.DefaultPoSSReserveDenom, 1000)
	params.MaxSignalsPerDay = 1
	params.MaxSignalsPerCuratorPerDay = 100

	k.SetParams(ctx, params)

	participant := "noor1participantlimit"
	curator := "noor1curatorlimit"
	date := "2025-01-02"

	// Premier signal : doit passer.
	_, _, err := k.ProcessSignalInternal(
		ctx,
		participant,
		curator,
		noorsignaltypes.SignalTypeMicroDonation,
		date,
	)
	if err != nil {
		t.Fatalf("expected first signal to pass, got error: %v", err)
	}

	countAfterFirst := k.GetDailySignalsCount(ctx, participant, date)
	if countAfterFirst != 1 {
		t.Fatalf("expected daily count=1 after first signal, got %d", countAfterFirst)
	}

	gsAfterFirst := k.ExportGenesis(ctx)

	// Deuxième signal le même jour : doit être refusé (limite atteinte).
	_, _, err2 := k.ProcessSignalInternal(
		ctx,
		participant,
		curator,
		noorsignaltypes.SignalTypeMicroDonation,
		date,
	)
	if err2 == nil {
		t.Fatalf("expected error when exceeding MaxSignalsPerDay, got nil")
	}

	// Le compteur ne doit pas avoir augmenté.
	countAfterSecond := k.GetDailySignalsCount(ctx, participant, date)
	if countAfterSecond != countAfterFirst {
		t.Fatalf("expected daily count to remain %d after blocked signal, got %d",
			countAfterFirst, countAfterSecond)
	}

	// Genesis ne doit pas changer non plus.
	gsAfterSecond := k.ExportGenesis(ctx)
	if gsAfterSecond.TotalSignals != gsAfterFirst.TotalSignals {
		t.Fatalf("expected TotalSignals to remain %d after blocked signal, got %d",
			gsAfterFirst.TotalSignals, gsAfterSecond.TotalSignals)
	}

	beforeInt, ok := sdk.NewIntFromString(gsAfterFirst.TotalMinted)
	if !ok {
		t.Fatalf("invalid TotalMinted after first signal: %s", gsAfterFirst.TotalMinted)
	}
	afterInt, ok := sdk.NewIntFromString(gsAfterSecond.TotalMinted)
	if !ok {
		t.Fatalf("invalid TotalMinted after second signal: %s", gsAfterSecond.TotalMinted)
	}

	if !afterInt.Equal(beforeInt) {
		t.Fatalf("expected TotalMinted to remain %s after blocked signal, got %s",
			beforeInt.String(), afterInt.String())
	}
}
