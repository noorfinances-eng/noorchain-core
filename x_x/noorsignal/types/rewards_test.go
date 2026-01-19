package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// helper pour construire des Params simples pour les tests
func newTestParams(
	enabled bool,
	baseReward sdk.Coin,
	maxRewardPerDay sdk.Coin,
	weightMicro uint32,
	weightParticipation uint32,
	weightContent uint32,
	weightCCN uint32,
) Params {
	return Params{
		PoSSEnabled:               enabled,
		MaxSignalsPerDay:          20,
		MaxSignalsPerCuratorPerDay: 100,
		MaxRewardPerDay:           maxRewardPerDay,
		BaseReward:                baseReward,
		WeightMicroDonation:       weightMicro,
		WeightParticipation:       weightParticipation,
		WeightContent:             weightContent,
		WeightCCN:                 weightCCN,
		PoSSReserveDenom:          "unur",
		HalvingPeriodBlocks:       0, // pas de halving pour ce premier test
	}
}

// Test 1 : si PoSSEnabled = false → rewards = 0 / 0 (mais bon denom)
func TestComputeSignalReward_PoSSDisabled_ReturnsZero(t *testing.T) {
	baseReward := sdk.NewInt64Coin("unur", 100)
	maxReward := sdk.NewInt64Coin("unur", 1000)

	params := newTestParams(
		false, // PoSSEnabled
		baseReward,
		maxReward,
		5, 2, 3, 1, // poids quelconques
	)

	// On prend un height basique (1) et un type "micro_donation"
	rewardParticipant, rewardCurator, err := ComputeSignalReward(
		params,
		SignalTypeMicroDonation,
		1,
	)

	if err != nil {
		t.Fatalf("expected no error when PoSSEnabled=false, got: %v", err)
	}

	if !rewardParticipant.IsZero() {
		t.Fatalf("expected participant reward to be zero when PoSSEnabled=false, got %s", rewardParticipant.String())
	}
	if !rewardCurator.IsZero() {
		t.Fatalf("expected curator reward to be zero when PoSSEnabled=false, got %s", rewardCurator.String())
	}

	if rewardParticipant.Denom != "unur" {
		t.Fatalf("expected denom 'unur' for participant reward, got %s", rewardParticipant.Denom)
	}
	if rewardCurator.Denom != "unur" {
		t.Fatalf("expected denom 'unur' for curator reward, got %s", rewardCurator.Denom)
	}
}

// Test 2 : si PoSSEnabled = true + BaseReward * weight → 70/30 correct
func TestComputeSignalReward_PoSSEnabled_MicroDonation_WeightAndSplit(t *testing.T) {
	// BaseReward = 100 unur, weight micro_donation = 5
	// total théorique = 100 * 5 = 500 unur
	baseReward := sdk.NewInt64Coin("unur", 100)
	maxReward := sdk.NewInt64Coin("unur", 10000)

	params := newTestParams(
		true, // PoSSEnabled
		baseReward,
		maxReward,
		5, 2, 3, 1, // micro=5, participation=2, content=3, ccn=1
	)

	height := int64(1) // pas de halving (HalvingPeriodBlocks=0)

	participantReward, curatorReward, err := ComputeSignalReward(
		params,
		SignalTypeMicroDonation,
		height,
	)

	if err != nil {
		t.Fatalf("expected no error when PoSSEnabled=true, got: %v", err)
	}

	total := participantReward.Amount.Add(curatorReward.Amount)
	expectedTotal := sdk.NewInt(500) // 100 * 5

	if !total.Equal(expectedTotal) {
		t.Fatalf("expected total reward 500unur, got %sunur", total.String())
	}

	// 70% pour participant, 30% pour curator
	expectedParticipant := sdk.NewInt(350) // 70% de 500
	expectedCurator := sdk.NewInt(150)     // 30% de 500

	if !participantReward.Amount.Equal(expectedParticipant) {
		t.Fatalf("expected participant reward 350unur, got %sunur", participantReward.Amount.String())
	}
	if !curatorReward.Amount.Equal(expectedCurator) {
		t.Fatalf("expected curator reward 150unur, got %sunur", curatorReward.Amount.String())
	}

	if participantReward.Denom != "unur" || curatorReward.Denom != "unur" {
		t.Fatalf("expected denom 'unur' for both rewards, got %s / %s",
			participantReward.Denom, curatorReward.Denom)
	}
}
