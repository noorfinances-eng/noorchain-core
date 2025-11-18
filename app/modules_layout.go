package app

// Ce fichier définit le "plan" des modules NOORCHAIN :
// - ordre d'exécution des BeginBlockers / EndBlockers
// - ordre d'initialisation du genesis
//
// À ce stade, ce ne sont que des listes de noms (string) basées sur
// les constantes définies dans modules.go. Plus tard, ces listes
// serviront à configurer le ModuleManager réel.

var (
	// Modules exécutés dans BeginBlocker (au début de chaque bloc).
	BeginBlockerOrder = []string{
		ModuleUpgrade,
		ModuleCrisis,
		ModuleSlashing,
		ModuleStaking,
		ModuleGov,
		ModuleEvm,
		ModuleFeeMarket,
		ModuleNoorSignal,
	}

	// Modules exécutés dans EndBlocker (à la fin de chaque bloc).
	EndBlockerOrder = []string{
		ModuleGov,
		ModuleStaking,
		ModuleEvm,
		ModuleFeeMarket,
		ModuleNoorSignal,
	}

	// Ordre d'initialisation du genesis (InitGenesis).
	InitGenesisOrder = []string{
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
)
