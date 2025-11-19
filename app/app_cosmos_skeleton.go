package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewNoorchainAppWithCosmos définit le constructeur "Cosmos" futur
// pour l'application NOORCHAIN.
//
// Il utilise AppBuilder pour :
// - construire un BaseApp réel (stores montés)
// - construire la structure des keepers
// - récupérer la configuration d'encodage
// - construire le ModuleManager (AppModules) et appliquer l'ordre des modules.
// - enregistrer les hooks (BeginBlock, EndBlock, InitGenesis).
func NewNoorchainAppWithCosmos(
	logger sdk.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts interface{},
) *App {

	// 1) Créer un builder avec les paramètres Cosmos standard.
	builder := NewAppBuilder(
		logger,
		db,
		traceStore,
		loadLatest,
		appOpts,
	)

	// 2) Construire la BaseApp via le builder.
	var base *baseapp.BaseApp = builder.BuildBaseApp()

	// 3) Construire la structure des keepers.
	keepers := builder.BuildKeepers()

	// 4) Récupérer la config d'encodage utilisée par ce builder.
	encCfg := builder.EncodingConfig()

	// 5) Créer le ModuleManager (AppModules) à partir des keepers + encCfg.
	modules := NewAppModuleManager(keepers, encCfg)

	// 6) Appliquer l'ordre des modules (BeginBlock, EndBlock, InitGenesis)
	// défini dans modules_layout.go.
	ConfigureModuleManagerOrder(modules.Manager)

	// 7) Construire l'instance de l'application NOORCHAIN.
	app := &App{
		BaseApp:  base,
		Name:     "NOORCHAIN",
		Version:  "0.0.1-dev",
		Keepers:  keepers,
		Encoding: encCfg,
		Modules:  modules,
	}

	// 8) Enregistrer les hooks de cycle de vie (BeginBlock / EndBlock)
	// auprès de BaseApp, pour que le ModuleManager soit appelé à chaque bloc.
	if app.BaseApp != nil {
		app.SetBeginBlocker(app.BeginBlocker)
		app.SetEndBlocker(app.EndBlocker)

		// 9) Enregistrer l'InitChainer pour le genesis.
		app.SetInitChainer(app.InitChainer)
	}

	return app
}
