package types

// ModuleName defines the module name for x/noorsignal (PoSS).
const (
	ModuleName = "noorsignal"

	// StoreKey is the primary KVStore key for the PoSS module.
	StoreKey = ModuleName

	// RouterKey is used for routing messages to this module.
	RouterKey = ModuleName

	// MemStoreKey is the in-memory store key (not used yet, but kept for future use).
	MemStoreKey = "mem_noorsignal"
)

// -----------------------------------------------------------------------------
// Store keys
// -----------------------------------------------------------------------------

// KeyTotalSignals is the key used to store the global total number of validated
// PoSS signals in the KVStore.
//
// This is a single uint64 value encoded in BigEndian.
// It will be incremented later when PoSS messages are processed.
var (
	KeyTotalSignals = []byte{0x01}
)
