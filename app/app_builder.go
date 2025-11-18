package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AppBuilder est un helper qui va progressivement construire
// l'application Cosmos SDK complète de NOORCHAIN.
type AppBuilder struct {
	logger     sdk.Logger
	db         dbm.DB
	traceStore io.Writer
	loadLatest bool
	appOpts    interface{}
	encCfg     EncodingConfig
}

// NewAppBuilder crée un nouveau AppBuilder en utilisant les paramètres
// de constructeur classiques Cosmos (logger, DB, trace, etc.).
func NewAppBuilder(
	logger sdk.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts interface{},
) *AppBuilder {
	encCfg := MakeEncodingConfig()

	return &AppBuilder{
		logger:     logger,
		db:         db,
		traceStore: traceStore,
		loadLatest: loadLatest,
		appOpts:    appOpts,
		encCfg:     encCfg,
	}
}

// EncodingConfig retourne la configuration d'encodage utilisée
// par ce builder (codec, TxConfig, etc.).
func (b *AppBuilder) EncodingConfig() EncodingConfig {
	return b.encCfg
}

// BuildBaseApp crée une instance minimale de baseapp.BaseApp.
//
// Maintenant que MakeEncodingConfig() est réel, on utilise toujours
// un TxDecoder valide provenant de encCfg.TxConfig.
func (b *AppBuilder) BuildBaseApp() *baseapp.BaseApp {
	// Récupérer le décodeur de transactions depuis la config d'encodage.
	txDecoder := b.encCfg.TxConfig.TxDecoder()

	// Créer une BaseApp minimale.
	base := baseapp.NewBaseApp(
		AppName,
		b.logger,
		b.db,
		txDecoder,
		baseapp.SetChainID(ChainID),
	)

	// Charger la dernière version stockée en DB si demandé.
	if b.loadLatest && base != nil {
		if err := base.LoadLatestVersion(); err != nil {
			// Plus tard, on remplacera ce panic par une gestion propre de l'erreur.
			panic(err)
		}
	}

	return base
}
