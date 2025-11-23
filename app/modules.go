package app

// Nom des modules Cosmos / Ethermint / NOORCHAIN utilisés dans l'app.
//
// Ces constantes sont utilisées pour :
// - les clés de KVStore (StoreKeys)
// - l'ordre des BeginBlockers / EndBlockers / InitGenesis
//   (voir modules_layout.go).
//
// Toute la logique AppModules (module.Manager, etc.) est dans
// module_manager.go.

const (
	// Modules Cosmos SDK de base
	ModuleAuth     = "auth"
	ModuleBank     = "bank"
	ModuleStaking  = "staking"
	ModuleMint     = "mint"
	ModuleSlashing = "slashing"
	ModuleGov      = "gov"
	ModuleParams   = "params"
	ModuleCrisis   = "crisis"
	ModuleUpgrade  = "upgrade"
	ModuleIBC      = "ibc"
	ModuleTransfer = "transfer"

	// Modules Ethermint / EVM
	ModuleEvm       = "evm"
	ModuleFeeMarket = "feemarket"

	// Module PoSS NOORCHAIN
	ModuleNoorSignal = "noorsignal"
)
