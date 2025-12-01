package types

// Module name and basic routing keys for the PoSS (x/noorsignal) module.
const (
	ModuleName   = "noorsignal"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

// Keys used in the KVStore to manage PoSS state.
//
// - KeyLastResetDay: (reserved) last day (Unix days) when daily counters were reset.
// - KeyGenesisState: JSON-encoded GenesisState for the module.
// - KeyPrefixParticipantDailyCount / KeyPrefixCuratorDailyCount:
//   reserved prefixes for future per-address daily counters (if needed).
var (
	KeyLastResetDay                = []byte{0x01}
	KeyGenesisState                = []byte{0x02}
	KeyPrefixParticipantDailyCount = []byte{0x10}
	KeyPrefixCuratorDailyCount     = []byte{0x11}
)
