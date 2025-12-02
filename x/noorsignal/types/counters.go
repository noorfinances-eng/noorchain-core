package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -----------------------------------------------------------------------------
// DAILY COUNTERS (per address, per day)
// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------
// GLOBAL PoSS STATE KEYS (stored once for the whole chain)
// -----------------------------------------------------------------------------

var (
	// Total number of PoSS signals processed since genesis.
	KeyTotalSignals = []byte("total_signals")

	// Total amount of NUR minted via PoSS ("unur" smallest unit).
	KeyTotalMinted = []byte("total_minted")

	// Pending mint (not yet distributed) accumulated by ProcessSignalInternal.
	KeyPendingMint = []byte("pending_mint")
)

// -----------------------------------------------------------------------------
// GENERIC HELPERS FOR GLOBAL PoSS STATE (uint64 / string)
// -----------------------------------------------------------------------------

// GetUint64 reads a uint64 from KVStore (big-endian).
func GetUint64(store sdk.KVStore, key []byte) uint64 {
	b := store.Get(key)
	if len(b) == 0 {
		return 0
	}
	return sdk.BigEndianToUint64(b)
}

// SetUint64 stores a uint64 under the given key (big-endian).
func SetUint64(store sdk.KVStore, key []byte, value uint64) {
	store.Set(key, sdk.Uint64ToBigEndian(value))
}

// GetString reads a raw UTF-8 string from the store.
func GetString(store sdk.KVStore, key []byte) string {
	b := store.Get(key)
	if len(b) == 0 {
		return ""
	}
	return string(b)
}

// SetString stores a raw UTF-8 string in the store.
func SetString(store sdk.KVStore, key []byte, value string) {
	store.Set(key, []byte(value))
}
