package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	noorsignal "github.com/noorfinances-eng/noorchain-core/x/noorsignal"
)

// AppModules regroupe les éléments liés aux modules Cosmos de NOORCHAIN.
//
// Pour l’instant, on conserve uniquement un module.Manager.
// Plus tard, on pourra étendre cette structure si on ajoute des helpers
// ou des gestionnaires spécifiques.
type AppModules struct {
	Manager *module.Manager
}

// NewAppModuleManager crée un module.Manager pour NOORCHAIN.
//
// Étape actuelle :
// - on crée un module.Manager avec le module PoSS (x/noorsignal)
//   déjà câblé avec son Keeper.
func NewAppModuleManager(keepers AppKeepers, _ EncodingConfig) AppModules {
	// 1) Construire le module PoSS (noorsignal) à partir de son Keeper.
	noorSignalModule := noorsignal.NewAppModule(keepers.NoorSignalKeeper)

	// 2) Créer le module.Manager avec le module PoSS.
	mm := module.NewManager(
		noorSignalModule,
		// ⚠️ Les autres modules (auth, bank, staking, gov, evm, feemarket…)
		// seront ajoutés ici plus tard.
	)

	return AppModules{
		Manager: mm,
	}
}

// ConfigureModuleManagerOrder applique au module.Manager l'ordre
// d'exécution défini dans modules_layout.go pour :
//
// - BeginBlocker
// - EndBlocker
// - InitGenesis
func ConfigureModuleManagerOrder(mm *module.Manager) {
	if mm == nil {
		return
	}

	mm.SetOrderBeginBlockers(BeginBlockerOrder...)
	mm.SetOrderEndBlockers(EndBlockerOrder...)
	mm.SetOrderInitGenesis(InitGenesisOrder...)
}
