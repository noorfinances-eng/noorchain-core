package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	noorsignalmodule "github.com/noorfinances-eng/noorchain-core/x/noorsignal"
)

// AppModules regroupe les éléments liés aux modules Cosmos de NOORCHAIN.
//
// Pour l'instant, il ne contient qu'un module.Manager. Plus tard, on
// pourra y ajouter d'autres structures (par exemple un configurateur
// de routes, des helpers pour InitGenesis, etc.).
type AppModules struct {
	Manager *module.Manager
}

// NewAppModuleManager crée un module.Manager pour NOORCHAIN.
//
// Étape actuelle :
// - crée un module.Manager avec les modules disponibles
//   (pour l'instant : seulement le module PoSS noorsignal).
//
// Dans des phases futures, on ajoutera ici les autres modules Cosmos
// (auth, bank, staking, gov, evm, etc.) en appelant
// module.NewManager(authModule, bankModule, ...).
func NewAppModuleManager(keepers AppKeepers, encCfg EncodingConfig) AppModules {
	// 1) Construire le module PoSS (noorsignal) à partir de son Keeper.
	noorSignalModule := noorsignalmodule.NewAppModule(keepers.NoorSignalKeeper)

	// 2) Créer le module.Manager avec le module PoSS.
	mm := module.NewManager(
		noorSignalModule,
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
//
// Remarque :
// - Cette fonction ne crée pas les modules, elle ne fait que définir
//   l'ordre dans lequel ils s'exécuteront une fois ajoutés au manager.
func ConfigureModuleManagerOrder(mm *module.Manager) {
	if mm == nil {
		return
	}

	// Ordre pour BeginBlockers, EndBlockers et InitGenesis.
	mm.SetOrderBeginBlockers(BeginBlockerOrder...)
	mm.SetOrderEndBlockers(EndBlockerOrder...)
	mm.SetOrderInitGenesis(InitGenesisOrder...)
}
