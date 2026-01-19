# NOORCHAIN 1.0 — Phase 4 — EVM & Core Cosmos (v0.46.11) — Version 1.1

**Objectif**  
Documenter l’état **réel** du noyau technique NOORCHAIN après l’intégration :

- du Cosmos SDK (auth / bank / staking / gov / params),
- d’Ethermint (EVM + FeeMarket v0.22.0),
- des stores et keepers,
- sans encore implémenter l’ante handler EVM complet ni PoSS.

Ce document décrit **ce qui existe déjà dans le code** et ce qui viendra plus tard.

---

## 1. Versions et stack technique

- **Langage** : Go 1.22
- **Cosmos SDK** : `v0.46.11`
- **Tendermint** : remplacé par **CometBFT** via `replace`  
  - `github.com/tendermint/tendermint => github.com/cometbft/cometbft v0.34.27`
- **Base de données** : `github.com/tendermint/tm-db v0.6.7`
- **Ethermint** : `v0.22.0`
  - Module EVM : `github.com/evmos/ethermint/x/evm`
  - Module FeeMarket (EIP-1559) : `github.com/evmos/ethermint/x/feemarket`

**Fichier principal** : `app/app.go`  
**Binaire** : `cmd/noord/main.go`

---

## 2. Structure de `NoorchainApp`

```go
type NoorchainApp struct {
    *baseapp.BaseApp

    appCodec          codec.Codec
    interfaceRegistry codectypes.InterfaceRegistry

    // KV stores
    keys  map[string]*storetypes.KVStoreKey
    // Transient stores (Params + FeeMarket)
    tkeys map[string]*storetypes.TransientStoreKey

    // Params
    ParamsKeeper paramskeeper.Keeper

    // Cosmos SDK keepers
    AccountKeeper authkeeper.AccountKeeper
    BankKeeper    bankkeeper.BaseKeeper
    StakingKeeper stakingkeeper.Keeper
    // GovKeeper viendra plus complet plus tard

    // Ethermint
    FeeMarketKeeper feemarketkeeper.Keeper
    EvmKeeper       *evmkeeper.Keeper // ajouté dans les blocs EVM

    mm *module.Manager
}
Points importants :

BaseApp est initialisée avec :

un TxConfig basé sur authtx.NewTxConfig,

un InterfaceRegistry propre,

un TxDecoder correct pour Cosmos.

On utilise un EncodingConfig minimal :

ProtoCodec + LegacyAmino + TxConfig avec DefaultSignModes.

3. Stores montés (KV + transient)
3.1. KV stores
Dans NewNoorchainApp, on crée et on monte les clés KV :

auth : authtypes.StoreKey

bank : banktypes.StoreKey

staking : stakingtypes.StoreKey

gov : govtypes.StoreKey (préparation pour un Gov complet)

params : paramstypes.StoreKey (Paramètres globaux)

evm : evmtypes.StoreKey (état EVM : storage, code, logs…)

feemarket : feemarkettypes.StoreKey (état EIP-1559)

Tous ces stores sont montés en :

go
Copier le code
app.MountStore(key, storetypes.StoreTypeIAVL)
3.2. Transient stores
On a également :

paramstypes.TStoreKey : transient store pour Params

feemarkettypes.StoreKey : transient store pour FeeMarket (gas wanted, etc.)

Montés en :

go
Copier le code
app.MountStore(tkey, storetypes.StoreTypeTransient)
4. ParamsKeeper et Subspaces
On utilise un ParamsKeeper réel :

go
Copier le code
app.ParamsKeeper = paramskeeper.NewKeeper(
    app.appCodec,
    encCfg.Amino,
    app.keys[paramstypes.StoreKey],
    app.tkeys[paramstypes.TStoreKey],
)
Puis on crée des subspaces par module :

authSubspace := app.ParamsKeeper.Subspace(authtypes.ModuleName)

bankSubspace := app.ParamsKeeper.Subspace(banktypes.ModuleName)

stakingSubspace := app.ParamsKeeper.Subspace(stakingtypes.ModuleName)

govSubspace := app.ParamsKeeper.Subspace(govtypes.ModuleName) (Gov)

feemarketSubspace := app.ParamsKeeper.Subspace(feemarkettypes.ModuleName)

evmSubspace := app.ParamsKeeper.Subspace(evmtypes.ModuleName) (préparé, utilisé plus tard)

Ces subspaces permettent à chaque module de gérer ses paramètres (feemarket, EVM, etc.) proprement.

5. Keepers Cosmos SDK
5.1. AccountKeeper
go
Copier le code
app.AccountKeeper = authkeeper.NewAccountKeeper(
    app.appCodec,
    app.keys[authtypes.StoreKey],
    authSubspace,
    authtypes.ProtoBaseAccount,
    map[string][]string{},
    "noorchain",
)
Utilise ProtoBaseAccount

ParamSubspace dédié

Bech32 prefix logique "noorchain" (sera aligné plus tard avec noor / noorvaloper… via config.go / params.go).

5.2. BankKeeper
go
Copier le code
app.BankKeeper = bankkeeper.NewBaseKeeper(
    app.appCodec,
    app.keys[banktypes.StoreKey],
    app.AccountKeeper,
    bankSubspace,
    map[string]bool{},
)
Restreint par une map sendEnabled (vide pour l’instant → tout activé).

5.3. StakingKeeper
go
Copier le code
app.StakingKeeper = stakingkeeper.NewKeeper(
    app.appCodec,
    app.keys[stakingtypes.StoreKey],
    app.AccountKeeper,
    app.BankKeeper,
    stakingSubspace,
)
Gère les validateurs, délégations, etc.

À terme, il sera alimenté dans le genesis (Phase Testnet).

6. FeeMarketKeeper (EIP-1559)
On instancie un FeeMarketKeeper compatbile avec Ethermint 0.22 :

Store :

KV : feemarkettypes.StoreKey

Transient : feemarkettypes.StoreKey (même nom, usage différent)

Autorité :

feeAuthority := authtypes.NewModuleAddress("gov")
→ l’adresse du module gov est l’autorité qui pourra mettre à jour les params (plus tard, via propositions on-chain).

go
Copier le code
app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
    app.appCodec,
    feeAuthority,
    app.keys[feemarkettypes.StoreKey],
    app.tkeys[feemarkettypes.StoreKey],
    feemarketSubspace,
)
Ce keeper gère notamment :

BlockGasWanted (via KV store),

TransientGasWanted (transient store),

les params comme MinGasPrice, MinGasMultiplier,

la base fee EIP-1559 utilisée par l’EVM.

7. EvmKeeper (squelette intégré)
Nous avons intégré le Keeper EVM d’Ethermint (x/evm/keeper) avec :

storeKey pour l’état EVM,

transientKey pour les données temporaires par bloc,

accès à :

AccountKeeper

BankKeeper

StakingKeeper

FeeMarketKeeper

authority = module gov (pour les MsgUpdateParams)

champ eip155ChainID (déduit de ctx.ChainID() à l’exécution)

customPrecompiles + evmConstructor (pour des precompiles plus tard)

gestion du tracer pour débogage EVM.

Les principales responsabilités du Keeper EVM dans la version intégrée :

Gestion de la bloom (logs EVM → Web3).

Index de transaction EVM dans le bloc (transient).

Taille des logs.

Stockage des comptes (nonce, codehash) via StateDB.

Lecture de la balance en utilisant le denom EVM défini dans les params (ex. unur).

Récupération de la base fee via FeeMarketKeeper.

Pour l’instant :

Le Keeper est en place.

Mais l’ante handler EVM (pipeline de vérification + exécution) n’est pas encore branché.

Aucun MsgEthereumTx n’est encore enregistré dans le ModuleManager (pas de routes, pas de CLI).

C’est volontaire : on avance étape par étape pour ne jamais casser go build.

8. ModuleManager actuel
Dans NewNoorchainApp, on initialise un module.Manager minimal :

go
Copier le code
app.mm = module.NewManager(
    auth.NewAppModule(app.appCodec, app.AccountKeeper, nil),
    bank.NewAppModule(app.appCodec, app.BankKeeper, app.AccountKeeper),
    staking.NewAppModule(app.appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
    // gov.NewAppModule(...) viendra plus tard
    // evm.NewAppModule(...) viendra plus tard
    // feemarket.NewAppModule(...) viendra plus tard
)
Ordre d’init Genesis :

go
Copier le code
app.mm.SetOrderInitGenesis(
    authtypes.ModuleName,
    banktypes.ModuleName,
    stakingtypes.ModuleName,
    // govtypes.ModuleName (plus tard)
    // evmtypes.ModuleName (plus tard)
    // feemarkettypes.ModuleName (plus tard)
)
Pour l’instant :

On n’a pas encore branché Gov/EVM/FeeMarket dans le ModuleManager.

On garde un `squelette minimal** qui compile et qui nous laisse la liberté de compléter l’ordre InitGenesis, BeginBlocker, EndBlocker, etc. dans un bloc ultérieur.

9. Limites actuelles (volontaires)
À ce stade de la Phase 4 :

✅ Le noyau Cosmos/EVM compile sans erreur.

✅ Stores et keepers sont en place (auth, bank, staking, gov store, params, evm, feemarket).

❌ Pas encore :

de GovKeeper branché dans le ModuleManager,

de Evm / Feemarket AppModules,

de routes et handlers pour les messages EVM,

d’ante handler spécifique pour EVM,

de logique PoSS (x/noorsignal) dans l’app.

C’est un point de sauvegarde stable de la Phase 4 :

“Cosmos core + EVM/FeeMarket prêts, mais non activés.”

10. Prochaines étapes recommandées (Phase 4, suite)
EVM Ante handler (EVM-ANTE 1–3)

Créer un ante handler qui :

vérifie signatures EVM,

vérifie les frais / gas selon FeeMarket,

prépare l’exécution EVM.

AppModules EVM + FeeMarket (EVM-MOD 1–2)

Ajouter evm.NewAppModule(...) et feemarket.NewAppModule(...) au ModuleManager.

Ajouter evm et feemarket à SetOrderInitGenesis, BeginBlockers / EndBlockers si nécessaire.

Gov complet (GOV 3–6)

GovKeeper finalisé avec router, MsgServiceRouter, params.

gov.NewAppModule(...) dans le ModuleManager.

Ordre InitGenesis incluant govtypes.ModuleName.

Squelette PoSS — module x/noorsignal

Créer le module de base (types, keeper, module.go, genesis.go).

Le brancher dans NoorchainApp sans la logique complète de récompenses (plus tard).

