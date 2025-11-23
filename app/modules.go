package app

// Nom des modules Cosmos / Ethermint / NOORCHAIN utilisés dans l'application.
//
// Ces constantes sont utilisées par :
// - store_keys.go           (pour créer les KVStoreKey)
// - modules_layout.go       (pour définir l'ordre BeginBlock / EndBlock / InitGenesis)
// - éventuellement d'autres parties de l'app plus tard.
//
// IMPORTANT :
// - Les valeurs doivent correspondre exactement aux ModuleName des modules
//   Cosmos SDK / Ethermint / NOORCHAIN (ex: "auth", "bank", "evm", "noorsignal").
// - Ne pas mettre d'espaces ni de majuscules.

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
