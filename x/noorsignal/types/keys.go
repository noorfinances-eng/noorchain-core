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
// - KeyLastResetDay: stores the last day (Unix days) when daily counters were reset.
// - KeyPrefixParticipantDailyCount: prefix for participant daily counters.
// - KeyPrefixCuratorDailyCount: prefix for curator daily counters.
// - KeyTotalSignals: global counter of all PoSS signals processed.
// - KeyTotalMinted: global amount of NUR minted via PoSS (in unur, as uint64).
var (
	KeyLastResetDay                = []byte{0x01}
	KeyPrefixParticipantDailyCount = []byte{0x10}
	KeyPrefixCuratorDailyCount     = []byte{0x11}

	KeyTotalSignals = []byte{0x20}
	KeyTotalMinted  = []byte{0x21}
)
