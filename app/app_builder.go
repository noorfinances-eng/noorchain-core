package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
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
// - clés de store (StoreKeys) pour les modules Cosmos / PoSS / EVM.
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

// BuildBaseApp crée une instance minimale de baseapp.BaseApp
// et monte les stores principaux (KV + transient).
//
// Maintenant que MakeEncodingConfig() est réel, on utilise toujours
// un TxDecoder valide provenant de encCfg.TxConfig.
func (b *AppBuilder) BuildBaseApp() *baseapp.BaseApp {
	sk := b.storeKeys

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

	// Monter les KVStores des modules principaux.
	if base != nil {
		base.MountKVStore(sk.AuthKey)
		base.MountKVStore(sk.BankKey)
		base.MountKVStore(sk.StakingKey)
		base.MountKVStore(sk.GovKey)
		base.MountKVStore(sk.ParamsKey)
		base.MountKVStore(sk.NoorSignalKey)

		// Stores pour EVM / FeeMarket (Ethermint).
		base.MountKVStore(sk.EvmKey)
		base.MountKVStore(sk.FeeMarketKey)

		// Store transient pour le module Params.
		base.MountTransientStore(sk.ParamsTransientKey)
	}

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
// - instanciation réelle d'un AccountKeeper
// - instanciation réelle d'un BankKeeper minimal
// - instanciation réelle du NoorSignalKeeper (PoSS)
// Les autres keepers (staking, gov, evm, feemarket, etc.) seront ajoutés
// dans des étapes futures.
func (b *AppBuilder) BuildKeepers() AppKeepers {
	sk := b.storeKeys
	enc := b.encCfg

	// 1) Créer un ParamsKeeper réel.
	paramsKeeper := paramskeeper.NewKeeper(
		enc.Marshaler,         // codec binaire (Protobuf)
		enc.Amino,             // codec legacy Amino pour JSON
		sk.ParamsKey,          // KVStoreKey pour les params
		sk.ParamsTransientKey, // TransientStoreKey pour les params temporaires
	)

	// 2) Préparer les permissions des comptes module (maccPerms).
	//
	// Pour l'instant, on utilise une map vide. Plus tard, on pourra
	// ajouter des comptes module (frais, distribution, PoSS, etc.)
	maccPerms := map[string][]string{}

	// 3) Créer un AccountKeeper minimal.
	accountKeeper := authkeeper.NewAccountKeeper(
		enc.Marshaler,
		sk.AuthKey,
		maccPerms,
		Bech32MainPrefix,
		authtypes.ProtoBaseAccount,
	)

	// 4) Préparer la liste des adresses "bloquées" pour le BankKeeper.
	//
	// Pour l'instant, on laisse la map vide; elle sera complétée plus tard.
	blockedAddrs := map[string]bool{}

	// 5) Créer un BankKeeper minimal.
	bankKeeper := bankkeeper.NewBaseKeeper(
		enc.Marshaler,
		sk.BankKey,
		accountKeeper,
		blockedAddrs,
		"", // authority (sera défini plus clairement plus tard)
	)

	// 6) Créer le NoorSignalKeeper (PoSS).
	//
	// Il utilise :
	// - le codec binaire principal (enc.Marshaler)
	// - la store key dédiée au module PoSS (sk.NoorSignalKey)
	noorSignalKeeper := noorsignalkeeper.NewKeeper(
		enc.Marshaler,
		sk.NoorSignalKey,
	)

	// 7) Construire la structure AppKeepers.
	// (Les keepers EVM / FeeMarket seront instanciés dans une étape dédiée.)
	return AppKeepers{
		AccountKeeper:    accountKeeper,
		BankKeeper:       bankKeeper,
		ParamsKeeper:     paramsKeeper,
		NoorSignalKeeper: noorSignalKeeper,

		EvmKeeper:       nil,
		FeeMarketKeeper: nil,
	}
}
