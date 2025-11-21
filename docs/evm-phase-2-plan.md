# NOORCHAIN — EVM Phase 2 (Ethermint Wiring Plan)

Objectif : brancher complètement Ethermint (EVM + FeeMarket) dans NOORCHAIN,
sans casser la chaîne, et en restant compatible avec la logique PoSS.

Ce document sert de plan détaillé pour les prochaines étapes EVM.

---

## 1. État actuel (fin Phase EVM-1)

Déjà en place :

- `go.mod` contient :
  - `github.com/evmos/ethermint`
- `app/modules.go` :
  - constantes `ModuleEvm` et `ModuleFeeMarket`
- `app/store_keys.go` :
  - `EvmKey` et `FeeMarketKey`
- `app/app_builder.go` :
  - les stores EVM sont montés dans BaseApp :
    - `MountKVStore(sk.EvmKey)`
    - `MountKVStore(sk.FeeMarketKey)`
- `app/keepers.go` :
  - `EvmKeeper *evmkeeper.Keeper`
  - `FeeMarketKeeper *feemarketkeeper.Keeper`
- `app/module_manager.go` :
  - ModuleManager configuré avec le module noorsignal (PoSS)

Les keepers EVM sont encore `nil`, volontairement :
ils seront initialisés dans cette Phase 2.

---

## 2. Étapes prévues (techniques) pour EVM Phase 2

### 2.1 Instanciation réelle des keepers EVM

Dans `app/app_builder.go` → `BuildKeepers()` :

- Créer un `EvmKeeper` via `evmkeeper.NewKeeper(...)` :
  - dépendances typiques :
    - codec
    - `EvmKey`
    - `AccountKeeper`
    - `BankKeeper`
    - `ParamsKeeper`
    - config EVM (chaîne, denom, etc.)
- Créer un `FeeMarketKeeper` via `feemarketkeeper.NewKeeper(...)` :
  - dépendances :
    - codec
    - `FeeMarketKey`
    - `ParamsKeeper`
    - config fees

> ⚠️ Remarque : les signatures exactes dépendront de la version Ethermint
> choisie (v1.0.x, etc.).  
> Le wiring sera fait plus tard à partir de la doc officielle Ethermint.

---

### 2.2 Ajout des AppModules EVM

Dans `app/module_manager.go` → `NewAppModuleManager` :

- Construire :
  - un module EVM :
    - `evmModule := evmmodule.NewAppModule(*keepers.EvmKeeper, ...)`
  - un module FeeMarket :
    - `feeMarketModule := feemarketmodule.NewAppModule(*keepers.FeeMarketKeeper, ...)`
- Ajouter au `module.NewManager(...)` :
  - `noorSignalModule` (déjà là)
  - `evmModule`
  - `feeMarketModule`

Puis, garder `ConfigureModuleManagerOrder(mm)` pour appliquer :

- `BeginBlockerOrder`
- `EndBlockerOrder`
- `InitGenesisOrder`

---

### 2.3 Genesis EVM & FeeMarket

Dans `testnet/genesis.json` :

- Vérifier / compléter les sections :
  - `"evm": { ... }`
  - `"feemarket": { ... }`
- Définir :
  - `evm.params.evm_denom = "unur"`
  - `evm.params.chain_config` :
    - chainId EVM (spécifique à NOORCHAIN)
    - support EIPs (155, 1559, etc.)
  - `feemarket.params` :
    - base fee
    - min gas price
    - params EIP-1559

---

### 2.4 Configuration du chainId EVM

Créer ou compléter un fichier (par exemple) :

- `app/evm_config.go`

Avec :

- un `const EvmChainID` (ex: `"1001"`, ou un autre ID unique)
- une fonction qui retourne la `chainConfig` pour Ethermint :
  - fork blocks (London, Byzantium, etc.)
  - chain parameters propres à NOORCHAIN

Ce chainId EVM doit être cohérent avec :

- le `ChainID` Cosmos (`noorchain-testnet-1`, `noorchain-1`)
- la configuration utilisée dans MetaMask

---

### 2.5 JSON-RPC & compatibilité MetaMask

Plus tard (hors scope direct de ce repo core), préparer :

- un process pour exposer l’API JSON-RPC (EVM) :
  - port 8545
  - derrière un reverse proxy si nécessaire
- les paramètres MetaMask :
  - RPC URL
  - chainId
  - symbol `NUR`
  - explorer URL (à venir)

---

## 3. Intégration PoSS + EVM

La logique PoSS (x/noorsignal) reste indépendante de l’EVM,
mais l’EVM pourra être utilisé pour :

- des smart-contracts NOOR dApps (curators, Proof of Light, etc.)
- des NFT / badges PoSS
- des dApps sociales / éducatives

PoSS continue d’utiliser :

- BankKeeper pour les récompenses en `unur`
- le module PoSS pour les signaux et curators

EVM ne remplace pas PoSS, il le complète.

---

## 4. Ordre recommandé d’implémentation

1. Finaliser Testnet côté **Genesis + Bank + PoSS** (récompenses réelles)  
2. Ensuite, ajouter :
   - `EvmKeeper` et `FeeMarketKeeper` réels
   - AppModules EVM
   - Genesis EVM propre
3. Compiler `noord`
4. Lancer un node local
5. Tester :
   - PoSS (submit / validate)
   - EVM (MetaMask + contrat “Hello NOORCHAIN”)

---

## 5. Rôle de ce document

Ce fichier **ne contient PAS de code Go**, uniquement un plan.  
Il sert à :

- garder la vision claire de l’intégration Ethermint,
- éviter les changements précipités dans le code,
- pouvoir reprendre facilement la Phase EVM-2 plus tard,
- servir de référence pour tout dev / auditeur technique.

