package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
)

// AppBuilder est un helper qui va progressivement construire
// l'application Cosmos SDK complète de NOORCHAIN.
//
// Il centralise :
// - logger
// - base de données
// - trace store
// - flag loadLatest
// - options d'application
// - configuration d'encodage
// - clés de store (StoreKeys) pour les modules Cosmos / PoSS.
type AppBuilder struct {
	logger     sdk.Logger
	db         dbm.DB
	traceStore io.Writer
	loadLatest bool
	appOpts    interface{}

	encCfg    EncodingConfig
	storeKeys StoreKeys
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
	storeKeys := NewStoreKeys()

	return &AppBuilder{
		logger:     logger,
		db:         db,
		traceStore: traceStore,
		loadLatest: loadLatest,
		appOpts:    appOpts,
		encCfg:     encCfg,
		storeKeys:  storeKeys,
	}
}

// EncodingConfig retourne la configuration d'encodage utilisée
// par ce builder (codec, TxConfig, etc.).
func (b *AppBuilder) EncodingConfig() EncodingConfig {
	return b.encCfg
}

// StoreKeys retourne les clés de store (KV + transient) associées
// aux modules principaux de NOORCHAIN.
func (b *AppBuilder) StoreKeys() StoreKeys {
	return b.storeKeys
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

// BuildKeepers crée et retourne la structure AppKeepers.
//
// Étape actuelle :
// - instanciation réelle du ParamsKeeper
// - les autres keepers restent à leur valeur zéro et seront ajoutés
//   dans des étapes futures.
func (b *AppBuilder) BuildKeepers() AppKeepers {
	sk := b.storeKeys
	enc := b.encCfg

	// 1) Créer un ParamsKeeper réel.
	//
	// NewKeeper(cdc, legacyAmino, key, tkey) retourne un paramskeeper.Keeper.
	paramsKeeper := paramskeeper.NewKeeper(
		enc.Marshaler,     // codec binaire (Protobuf)
		enc.Amino,         // codec legacy Amino pour JSON
		sk.ParamsKey,      // KVStoreKey pour les params
		sk.ParamsTransientKey, // TransientStoreKey pour les params temporaires
	)

	// 2) Construire la structure AppKeepers.
	//
	// Pour le moment, seul ParamsKeeper est réellement instancié.
	// Les autres seront remplis plus tard (AccountKeeper, BankKeeper, etc.).
	return AppKeepers{
		ParamsKeeper: paramsKeeper,
	}
}
