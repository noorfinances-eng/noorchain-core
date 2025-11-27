package app

// Module identifiers for NOORCHAIN.
// Phase 2: only names. Phase 3+: KVStore keys + keepers.
const (
	// Standard Cosmos SDK modules
	ModuleNameAuth    = "auth"
	ModuleNameBank    = "bank"
	ModuleNameStaking = "staking"
	ModuleNameGov     = "gov"

	// Future custom PoSS module (Phase 4+)
	ModuleNamePoSS = "noorsignal"
)
