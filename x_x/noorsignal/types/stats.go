package types

// PoSSStats is a simple, read-friendly snapshot of the PoSS global state.
//
// It is meant to be used by:
// - CLI commands ("noord query poss stats"),
// - dashboards,
// - monitoring tools.
//
// It does NOT change any state: it's a pure view of:
// - Genesis-level counters (TotalSignals, TotalMinted),
// - high-level PoSS params (enabled flag, daily limits, reserve denom).
type PoSSStats struct {
	// Total number of PoSS signals processed since genesis.
	TotalSignals uint64 `json:"total_signals" yaml:"total_signals"`

	// Total amount of NUR planned to be minted via PoSS
	// (sum of all participant + curator rewards), stored as a string in "unur".
	TotalMinted string `json:"total_minted" yaml:"total_minted"`

	// Global on/off switch for PoSS rewards.
	PoSSEnabled bool `json:"poss_enabled" yaml:"poss_enabled"`

	// Maximum number of PoSS signals allowed per participant per day.
	MaxSignalsPerDay uint64 `json:"max_signals_per_day" yaml:"max_signals_per_day"`

	// Maximum number of PoSS signals a curator can validate per day.
	MaxSignalsPerCuratorPerDay uint64 `json:"max_signals_per_curator_per_day" yaml:"max_signals_per_curator_per_day"`

	// Denom of the PoSS reserve (always "unur" for NOORCHAIN).
	PoSSReserveDenom string `json:"poss_reserve_denom" yaml:"poss_reserve_denom"`
}
