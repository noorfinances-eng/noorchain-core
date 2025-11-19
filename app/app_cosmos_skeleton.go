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
// - construire un BaseApp réel
// - construire la structure des keepers
// - récupérer la configuration d'encodage
// (les keepers concrets et les modules seront ajoutés plus tard).
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

	// 3) Construire la structure des keepers (placeholder pour l'instant).
	keepers := builder.BuildKeepers()

	// 4) Récupérer la config d'encodage utilisée par ce builder.
	encCfg := builder.EncodingConfig()

	// 5) Retourner l'instance de l'application NOORCHAIN.
	return &App{
		BaseApp:  base,    // non-nil quand tout sera câblé
		Name:     "NOORCHAIN",
		Version:  "0.0.1-dev",
		Keepers:  keepers, // pour l'instant vide
		Encoding: encCfg,
	}
}
