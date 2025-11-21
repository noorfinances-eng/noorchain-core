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

	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
)

// AppBuilder construit progressivement l'application Cosmos SDK complète.
type AppBuilder struct {
	logger     sdk.Logger
	db         dbm.DB
	traceStore io.Writer
	loadLatest bool
	appOpts    interface{}

	encCfg    EncodingConfig
	storeKeys StoreKeys
}

// NewAppBuilder crée un nouveau builder.
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

func (b *AppBuilder) EncodingConfig() EncodingConfig { return b.encCfg }
func (b *AppBuilder) StoreKeys() StoreKeys           { return b.storeKeys }

// BuildBaseApp crée une instance BaseApp complète (stores montés).
func (b *AppBuilder) BuildBaseApp() *baseapp.BaseApp {
	sk := b.storeKeys
	txDecoder := b.encCfg.TxConfig.TxDecoder()

	// 1) BaseApp minimale
	base := baseapp.NewBaseApp(
		AppName,
		b.logger,
		b.db,
		txDecoder,
		baseapp.SetChainID(ChainID),
	)

	// 2) Monter les KVStores
	if base != nil {
		// --- Stores Cosmos classiques ---
		base.MountKVStore(sk.AuthKey)
		base.MountKVStore(sk.BankKey)
		base.MountKVStore(sk.StakingKey)
		base.MountKVStore(sk.GovKey)
		base.MountKVStore(sk.ParamsKey)

		// --- Store PoSS ---
		base.MountKVStore(sk.NoorSignalKey)

		// --- EVM Stores (Ethermint) ---
		base.MountKVStore(sk.EvmKey)
		base.MountKVStore(sk.FeeMarketKey)

		// --- Transient store ---
		base.MountTransientStore(sk.ParamsTransientKey)
	}

	// 3) Charger la version si demandé
	if b.loadLatest && base != nil {
		if err := base.LoadLatestVersion(); err != nil {
			panic(err)
		}
	}

	return base
}

// BuildKeepers crée les keepers (version minimaliste pour l’instant).
func (b *AppBuilder) BuildKeepers() AppKeepers {
	sk := b.storeKeys
	enc := b.encCfg

	paramsKeeper := paramskeeper.NewKeeper(
		enc.Marshaler,
		enc.Amino,
		sk.ParamsKey,
		sk.ParamsTransientKey,
	)

	maccPerms := map[string][]string{}
	accountKeeper := authkeeper.NewAccountKeeper(
		enc.Marshaler,
		sk.AuthKey,
		maccPerms,
		Bech32MainPrefix,
		authtypes.ProtoBaseAccount,
	)

	blockedAddrs := map[string]bool{}
	bankKeeper := bankkeeper.NewBaseKeeper(
		enc.Marshaler,
		sk.BankKey,
		accountKeeper,
		blockedAddrs,
		"",
	)

	noorSignalKeeper := noorsignalkeeper.NewKeeper(
		enc.Marshaler,
		sk.NoorSignalKey,
	)

	// Keepers EVM non encore initialisés
	var evmKeeper evmkeeper.Keeper
	var feeMarketKeeper feemarketkeeper.Keeper

	return AppKeepers{
		AccountKeeper:    accountKeeper,
		BankKeeper:       bankKeeper,
		ParamsKeeper:     paramsKeeper,
		EvmKeeper:        evmKeeper,
		FeeMarketKeeper:  feeMarketKeeper,
		NoorSignalKeeper: noorSignalKeeper,
	}
}
