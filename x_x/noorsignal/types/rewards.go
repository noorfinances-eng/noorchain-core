package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SignalType represents the four PoSS signal families.
// These values will be reused later in msgs / proto.
type SignalType string

const (
	SignalTypeMicroDonation SignalType = "micro_donation"
	SignalTypeParticipation SignalType = "participation"
	SignalTypeContent       SignalType = "content"
	SignalTypeCCN           SignalType = "ccn"
)

// Structural 70/30 split for PoSS rewards.
// MUST stay aligned with genesis (GenesisState.ParticipantShare / CuratorShare).
const (
	ParticipantSharePercent uint32 = 70
	CuratorSharePercent     uint32 = 30
)

// WeightForSignalType returns the param weight corresponding
// to the given PoSS signal type.
//
// It uses the weights defined in Params (params.go).
func (p Params) WeightForSignalType(signalType SignalType) uint32 {
	switch signalType {
	case SignalTypeMicroDonation:
		return p.WeightMicroDonation
	case SignalTypeParticipation:
		return p.WeightParticipation
	case SignalTypeContent:
		return p.WeightContent
	case SignalTypeCCN:
		return p.WeightCCN
	default:
		// Defensive default: neutral weight 1 if type is unknown.
		return 1
	}
}

// ComputeBaseReward returns the raw PoSS reward for a single signal,
// BEFORE halving, BEFORE daily caps, BEFORE the 70/30 split.
//
// Formula: BaseReward * weight(signal_type)
func ComputeBaseReward(p Params, signalType SignalType) sdk.Coin {
	weight := p.WeightForSignalType(signalType)
	if weight == 0 {
		// In case of misconfigured params, fall back to 1.
		weight = 1
	}

	amount := p.BaseReward.Amount.MulRaw(int64(weight))
	return sdk.NewCoin(p.BaseReward.Denom, amount)
}

// ApplyHalving applies the PoSS halving schedule on a given reward,
// based on the current block height and HalvingPeriodBlocks.
//
// If HalvingPeriodBlocks == 0, halving is considered "not configured"
// and the reward is returned unchanged.
func ApplyHalving(p Params, height int64, reward sdk.Coin) sdk.Coin {
	if p.HalvingPeriodBlocks == 0 || height <= 0 {
		return reward
	}

	periods := height / int64(p.HalvingPeriodBlocks)
	if periods <= 0 {
		return reward
	}

	denom := reward.Denom
	amount := reward.Amount

	// Divide by 2^periods, but never go below zero.
	for i := int64(0); i < periods && amount.IsPositive(); i++ {
		amount = amount.QuoRaw(2)
	}

	return sdk.NewCoin(denom, amount)
}

// SplitReward70_30 splits a total reward into participant and curator
// shares, using the structural 70/30 rule.
//
// The sum of the two coins is always equal to `total` (we compute the
// curator share as "total - participant" to avoid rounding drift).
func SplitReward70_30(total sdk.Coin) (participant sdk.Coin, curator sdk.Coin, err error) {
	if ParticipantSharePercent+CuratorSharePercent != 100 {
		return sdk.Coin{}, sdk.Coin{}, fmt.Errorf(
			"invalid PoSS share config: %d + %d != 100",
			ParticipantSharePercent, CuratorSharePercent,
		)
	}

	if total.Amount.IsNegative() {
		return sdk.Coin{}, sdk.Coin{}, fmt.Errorf("total reward cannot be negative")
	}

	totalAmt := total.Amount

	// participant = total * 70 / 100
	pAmt := totalAmt.MulRaw(int64(ParticipantSharePercent)).QuoRaw(100)
	// curator = total - participant (to guarantee exact sum)
	cAmt := totalAmt.Sub(pAmt)

	participant = sdk.NewCoin(total.Denom, pAmt)
	curator = sdk.NewCoin(total.Denom, cAmt)

	return participant, curator, nil
}

// ComputeSignalReward is the high-level helper that combines:
//
//   1) Base reward (BaseReward * weight),
//   2) Halving logic,
//   3) 70/30 split.
//
// It DOES NOT:
//   - check daily limits,
//   - check PoSSReserve balances,
//   - update any state.
//
// These checks will live in the Keeper (PoSS Logic later).
func ComputeSignalReward(
	p Params,
	signalType SignalType,
	height int64,
) (participant sdk.Coin, curator sdk.Coin, err error) {
	// If PoSS is disabled, we return 0/0 but still with the correct denom.
	if !p.PoSSEnabled {
		zero := sdk.NewCoin(p.BaseReward.Denom, sdk.ZeroInt())
		return zero, zero, nil
	}

	base := ComputeBaseReward(p, signalType)
	halved := ApplyHalving(p, height, base)

	return SplitReward70_30(halved)
}
