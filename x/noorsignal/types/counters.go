package types

// DailyCounter represents a simple per-address counter for a given day.
//
// It will be used by the PoSS keeper later to:
// - track how many signals a participant has emitted today,
// - track how many signals a curator has validated today,
// and enforce the daily limits defined in params/genesis.
type DailyCounter struct {
	// Address is the bech32 NOOR address (noor1...).
	Address string `json:"address" yaml:"address"`

	// Date is stored as an ISO string "YYYY-MM-DD".
	// Example: "2025-12-01".
	Date string `json:"date" yaml:"date"`

	// Signals is the number of PoSS signals for this (address, date).
	Signals uint32 `json:"signals" yaml:"signals"`
}
