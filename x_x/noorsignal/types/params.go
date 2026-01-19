package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ----------------------------------------------------------
// PoSS Params — structure et valeurs par défaut
// ----------------------------------------------------------
//
// Champs stables décidés ensemble (PoSS Logic 4)
//
// - PoSSEnabled (bool): master switch pour activer/désactiver les rewards.
// - MaxSignalsPerDay (int): max de signaux par adresse et par jour.
// - MaxSignalsPerCuratorPerDay (int): max de signaux validés par un curateur.
// - MaxRewardPerDay (Coin): plafond de reward par adresse et par jour.
// - BaseReward (Coin): unité de reward, multipliée par un poids.
// - WeightMicroDonation / Participation / Content / CCN (ints): poids relatifs.
// - PoSSReserveDenom (string): denom de la réserve PoSS (toujours "unur").
// - HalvingPeriodBlocks (uint64): période de halving en blocs.
//
// ⚠️ Ici on ne fixe PAS encore les valeurs finales : ce sont des
// valeurs de base raisonnables, qu’on pourra adapter plus tard.
type Params struct {
	PoSSEnabled bool `json:"poss_enabled" yaml:"poss_enabled"`

	MaxSignalsPerDay           uint64 `json:"max_signals_per_day" yaml:"max_signals_per_day"`
	MaxSignalsPerCuratorPerDay uint64 `json:"max_signals_per_curator_per_day" yaml:"max_signals_per_curator_per_day"`

	MaxRewardPerDay sdk.Coin `json:"max_reward_per_day" yaml:"max_reward_per_day"`
	BaseReward      sdk.Coin `json:"base_reward" yaml:"base_reward"`

	WeightMicroDonation uint32 `json:"weight_micro_donation" yaml:"weight_micro_donation"`
	WeightParticipation uint32 `json:"weight_participation" yaml:"weight_participation"`
	WeightContent       uint32 `json:"weight_content" yaml:"weight_content"`
	WeightCCN           uint32 `json:"weight_ccn" yaml:"weight_ccn"`

	PoSSReserveDenom    string `json:"poss_reserve_denom" yaml:"poss_reserve_denom"`
	HalvingPeriodBlocks uint64 `json:"halving_period_blocks" yaml:"halving_period_blocks"`
}

// Default values (provisoires, modifiables plus tard par gouvernance)
const (
	DefaultPoSSEnabled = false

	DefaultMaxSignalsPerDay           uint64 = 20
	DefaultMaxSignalsPerCuratorPerDay uint64 = 100

	// BaseReward / MaxRewardPerDay seront en "unur".
	DefaultBaseRewardAmount      int64 = 1
	DefaultMaxRewardPerDayAmount int64 = 100

	DefaultWeightMicroDonation uint32 = 5
	DefaultWeightParticipation uint32 = 2
	DefaultWeightContent       uint32 = 3
	DefaultWeightCCN           uint32 = 1

	DefaultPoSSReserveDenom = "unur"

	// Placeholder : on décidera plus tard du vrai nombre de blocs
	// pour 8 ans (en fonction du temps de bloc réel).
	DefaultHalvingPeriodBlocks uint64 = 0
)

// -----------------------------------------------------------------------------
// ParamStore keys
// -----------------------------------------------------------------------------

var (
	KeyPoSSEnabled                = []byte("PoSSEnabled")
	KeyMaxSignalsPerDay           = []byte("MaxSignalsPerDay")
	KeyMaxSignalsPerCuratorPerDay = []byte("MaxSignalsPerCuratorPerDay")
	KeyMaxRewardPerDay            = []byte("MaxRewardPerDay")
	KeyBaseReward                 = []byte("BaseReward")
	KeyWeightMicroDonation        = []byte("WeightMicroDonation")
	KeyWeightParticipation        = []byte("WeightParticipation")
	KeyWeightContent              = []byte("WeightContent")
	KeyWeightCCN                  = []byte("WeightCCN")
	KeyPoSSReserveDenom           = []byte("PoSSReserveDenom")
	KeyHalvingPeriodBlocks        = []byte("HalvingPeriodBlocks")
)

// ParamKeyTable retourne la KeyTable à enregistrer dans le Subspace
// (appelée une fois dans le Keeper).
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs permet à x/params de lire/écrire les champs dans le store.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyPoSSEnabled, &p.PoSSEnabled, validateNoop),
		paramstypes.NewParamSetPair(KeyMaxSignalsPerDay, &p.MaxSignalsPerDay, validateNoop),
		paramstypes.NewParamSetPair(KeyMaxSignalsPerCuratorPerDay, &p.MaxSignalsPerCuratorPerDay, validateNoop),
		paramstypes.NewParamSetPair(KeyMaxRewardPerDay, &p.MaxRewardPerDay, validateNoop),
		paramstypes.NewParamSetPair(KeyBaseReward, &p.BaseReward, validateNoop),
		paramstypes.NewParamSetPair(KeyWeightMicroDonation, &p.WeightMicroDonation, validateNoop),
		paramstypes.NewParamSetPair(KeyWeightParticipation, &p.WeightParticipation, validateNoop),
		paramstypes.NewParamSetPair(KeyWeightContent, &p.WeightContent, validateNoop),
		paramstypes.NewParamSetPair(KeyWeightCCN, &p.WeightCCN, validateNoop),
		paramstypes.NewParamSetPair(KeyPoSSReserveDenom, &p.PoSSReserveDenom, validateNoop),
		paramstypes.NewParamSetPair(KeyHalvingPeriodBlocks, &p.HalvingPeriodBlocks, validateNoop),
	}
}

// validateNoop : on délègue la vraie validation à (p Params).Validate().
// Ici, on ne bloque rien champ par champ.
func validateNoop(i interface{}) error {
	return nil
}

// DefaultParams retourne une config PoSS "safe off":
// - PoSS désactivé,
// - limites raisonnables,
// - rewards définis mais non utilisés tant que PoSS est éteint.
func DefaultParams() Params {
	return Params{
		PoSSEnabled: DefaultPoSSEnabled,

		MaxSignalsPerDay:           DefaultMaxSignalsPerDay,
		MaxSignalsPerCuratorPerDay: DefaultMaxSignalsPerCuratorPerDay,

		MaxRewardPerDay: sdk.NewInt64Coin(DefaultPoSSReserveDenom, DefaultMaxRewardPerDayAmount),
		BaseReward:      sdk.NewInt64Coin(DefaultPoSSReserveDenom, DefaultBaseRewardAmount),

		WeightMicroDonation: DefaultWeightMicroDonation,
		WeightParticipation: DefaultWeightParticipation,
		WeightContent:       DefaultWeightContent,
		WeightCCN:           DefaultWeightCCN,

		PoSSReserveDenom:    DefaultPoSSReserveDenom,
		HalvingPeriodBlocks: DefaultHalvingPeriodBlocks,
	}
}

// Validate fait une validation basique des paramètres.
// Rien de "bloquant économique" ici, juste de la cohérence.
func (p Params) Validate() error {
	if p.PoSSReserveDenom == "" {
		return fmt.Errorf("PoSSReserveDenom cannot be empty")
	}

	if p.MaxSignalsPerDay == 0 {
		return fmt.Errorf("MaxSignalsPerDay must be > 0")
	}

	if p.MaxSignalsPerCuratorPerDay == 0 {
		return fmt.Errorf("MaxSignalsPerCuratorPerDay must be > 0")
	}

	if p.BaseReward.Denom != p.PoSSReserveDenom {
		return fmt.Errorf("BaseReward denom must match PoSSReserveDenom (%s)", p.PoSSReserveDenom)
	}

	if p.MaxRewardPerDay.Denom != p.PoSSReserveDenom {
		return fmt.Errorf("MaxRewardPerDay denom must match PoSSReserveDenom (%s)", p.PoSSReserveDenom)
	}

	if p.BaseReward.Amount.IsNegative() {
		return fmt.Errorf("BaseReward amount cannot be negative")
	}

	if p.MaxRewardPerDay.Amount.IsNegative() {
		return fmt.Errorf("MaxRewardPerDay amount cannot be negative")
	}

	if p.WeightMicroDonation == 0 ||
		p.WeightParticipation == 0 ||
		p.WeightContent == 0 ||
		p.WeightCCN == 0 {
		return fmt.Errorf("all PoSS weights must be > 0")
	}

	// HalvingPeriodBlocks = 0 veut dire "pas encore configuré".
	// Ce n’est pas une erreur bloquante à ce stade.
	return nil
}
