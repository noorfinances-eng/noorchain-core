package app

// This file defines the list of modules that NOORCHAIN will use.
// For now it contains core Cosmos modules + Ethermint EVM modules + PoSS.
//
// These names are used in:
// - the ModuleManager
// - BeginBlockers / EndBlockers
// - InitGenesis ordering
// - StoreKey creation
// - future routing & gRPC services

const (
	// --- Cosmos SDK core modules ---
	ModuleAuth     = "auth"
	ModuleBank     = "bank"
	ModuleStaking  = "staking"
	ModuleMint     = "mint"
	ModuleSlashing = "slashing"
	ModuleGov      = "gov"
	ModuleParams   = "params"
	ModuleCrisis   = "crisis"
	ModuleUpgrade  = "upgrade"

	// --- IBC ---
	ModuleIBC      = "ibc"
	ModuleTransfer = "transfer"

	// --- Ethermint ---
	ModuleEvm       = "evm"
	ModuleFeeMarket = "feemarket"

	// --- NOORCHAIN custom PoSS module ---
	ModuleNoorSignal = "noorsignal"
)

// Full list for documentation / future tooling.
var (
	AllModules = []string{
		ModuleAuth,
		ModuleBank,
		ModuleStaking,
		ModuleMint,
		ModuleSlashing,
		ModuleGov,
		ModuleParams,
		ModuleCrisis,
		ModuleUpgrade,
		ModuleIBC,
		ModuleTransfer,
		ModuleEvm,
		ModuleFeeMarket,
		ModuleNoorSignal,
	}

	BasicModules = []string{
		ModuleAuth,
		ModuleBank,
		ModuleStaking,
		ModuleMint,
		ModuleSlashing,
		ModuleGov,
		ModuleParams,
		ModuleCrisis,
		ModuleUpgrade,
	}
)
