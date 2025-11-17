package app

// This file defines the list of modules that NOORCHAIN will use.
// For now it only provides simple string-based placeholders.
//
// Later, when wiring the full Cosmos SDK application, these names
// will be used to configure:
// - the ModuleManager
// - BeginBlockers / EndBlockers
// - InitGenesis / ExportGenesis
// - Ordering of routes and queries
//
// Keeping this centralised makes it easier to reason about the app layout.

const (
	// Cosmos SDK core modules
	ModuleAuth     = "auth"
	ModuleBank     = "bank"
	ModuleStaking  = "staking"
	ModuleMint     = "mint"
	ModuleSlashing = "slashing"
	ModuleGov      = "gov"
	ModuleParams   = "params"
	ModuleCrisis   = "crisis"
	ModuleUpgrade  = "upgrade"

	// IBC related modules (future)
	ModuleIBC      = "ibc"
	ModuleTransfer = "transfer"

	// Ethermint / EVM related modules (future)
	ModuleEvm       = "evm"
	ModuleFeeMarket = "feemarket"

	// NOORCHAIN custom module (PoSS)
	ModuleNoorSignal = "noorsignal"
)

// ModuleList groups modules for easier reasoning.
// At this stage, this is only used as documentation in code.
// Later it can be used to build the ModuleManager.
var (
	// AllModules lists all modules that NOORCHAIN is expected to use in the long run.
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

	// BasicModules contains the minimal set of modules required for a basic chain
	// before enabling IBC, EVM and PoSS.
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
