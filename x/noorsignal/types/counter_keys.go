package types

// DailyCounterPrefix is the prefix used in the KVStore to store
// per-address, per-day PoSS counters.
//
// Concretely, keys will look like:
//   "daily_counter:noor1xyz...:2026-03-17"
const DailyCounterPrefix = "daily_counter:"

// DailyCounterKey builds the store key for a given (address, date) pair.
//
// - address: NOOR bech32 address (noor1...)
// - date: ISO day string "YYYY-MM-DD"
//
// Example resulting key:
//   []byte("daily_counter:noor1xyz...:2026-03-17")
func DailyCounterKey(address, date string) []byte {
	return []byte(DailyCounterPrefix + address + ":" + date)
}
