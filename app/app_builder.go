package app

import (
	"io"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// --- Cosmos SDK keepers ---
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	// --- NOORCHAIN PoSS ---
	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
)

// ------------------------------------------------------------
// AppBuilder : construit progressivement NOORCHAIN
// ------------------------------------------------------------
type AppBuilder struct {
	logger     sdk.Logger
	db         dbm.DB
	traceStore io.Writer
	loadLatest bool
	appOpts    interface{}

	encCfg    EncodingConfig
	storeKeys StoreKeys
}

// ------------------------------------------------------------
// NewAppBuilder : constructeur
// ------------------------------------------------------------
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

// ------------------------------------------------------------
// EncodingConfig getter
// ------------------------------------------------------------
func (b *AppBuilder) EncodingConfig() EncodingConfig {
	return b.encCfg
}

// ------------------------------------------------------------
// StoreKeys getter
// ------------------------------------------------------------
func (b *AppBuilder) StoreKeys() StoreKeys {
	return b.storeKeys
}

// ------------------------------------------------------------
// BuildBaseApp : crée BaseApp + monte les stores KV
// ------------------------------------------------------------
func (b *AppBuilder) BuildBaseApp() *baseapp.BaseApp {
	sk := b.storeKeys
	txDecoder := b.encCfg.TxConfig.TxDecoder()

	// 1) BaseApp
	base := baseapp.NewBaseApp(
		AppName,
		b.logger,
		b.db,
		txDecoder,
		baseapp.SetChainID(ChainID),
	)

	// 2) Mount KVStores
	if base != nil {
		// --- Stores Cosmos classiques
		base.MountKVStore(sk.AuthKey)
		base.MountKVStore(sk.BankKey)
		base.MountKVStore(sk.StakingKey)
		base.MountKVStore(sk.GovKey)
		base.MountKVStore(sk.ParamsKey)

		// --- Store PoSS
		base.MountKVStore(sk.NoorSignalKey)

		// ⚠️ AUCUN store EVM / FeeMarket pour l’instant (version light)

		// --- Transient store
		base.MountTransientStore(sk.ParamsTransientKey)
	}

	// 3) Load version
	if b.loadLatest && base != nil {
		if err := base.LoadLatestVersion(); err != nil {
			panic(err)
		}
	}

	return base
}

// ------------------------------------------------------------
// BuildKeepers : crée les keepers principaux (sans EVM pour l’instant)
// ------------------------------------------------------------
func (b *AppBuilder) BuildKeepers() AppKeepers {
	sk := b.storeKeys
	enc := b.encCfg

	// 1) ParamsKeeper
	paramsKeeper := paramskeeper.NewKeeper(
		enc.Marshaler,
		enc.Amino,
		sk.ParamsKey,
		sk.ParamsTransientKey,
	)

	// 2) AccountKeeper
	maccPerms := map[string][]string{}
	accountKeeper := authkeeper.NewAccountKeeper(
		enc.Marshaler,
		sk.AuthKey,
		maccPerms,
		Bech32MainPrefix,
		authtypes.ProtoBaseAccount,
	)

	// 3) BankKeeper
	blockedAddrs := map[string]bool{}
	bankKeeper := bankkeeper.NewBaseKeeper(
		enc.Marshaler,
		sk.BankKey,
		accountKeeper,
		blockedAddrs,
		"",
	)

	// 4) NoorSignalKeeper (PoSS)
	noorSignalKeeper := noorsignalkeeper.NewKeeper(
		enc.Marshaler,
		sk.NoorSignalKey,
	)

	// 5) Retourner les keepers (version light, sans EVM)
	return AppKeepers{
		AccountKeeper:    accountKeeper,
		BankKeeper:       bankKeeper,
		ParamsKeeper:     paramsKeeper,
		NoorSignalKeeper: noorSignalKeeper,
	}
}
