# NOORCHAIN — Phase 2 (Infrastructure & Outillage Testnet)

> Version 1.1 — Skeleton Cosmos SDK + Ethermint (sans PoSS)

---

## 1. Objectif de la Phase 2

- Créer une structure propre de blockchain Cosmos SDK + Ethermint.
- Avoir une application NOORCHAIN qui **compile** avec `go build ./...`.
- Préparer toute l’infrastructure de base **avant** d’ajouter PoSS (Phase 4).

---

## 2. État actuel (structure technique)

- Projet Go initialisé (`go.mod`, `cmd/`, `app/`, `x/`, `config/`, `scripts/`, `docs/`).
- CLI `noord` minimal :
  - `main`, `Execute`, `Dispatch`, `Start`, `Run`, `Version`, `Help`.
  - `Start` → `StartNOORChain()`.

- Application Cosmos `NOORChainApp` :
  - `BaseApp` avec :
    - logger minimal,
    - base de données en mémoire (`MemDB`),
    - `CosmosTxDecoder` (placeholder).
  - Stores montés via `LoadStores`.
  - Encoding Cosmos (`CosmosEncodingConfig` + builder).
  - Keepers, ModuleManager, StoreLoader (placeholders).
  - Query router + Msg service router (placeholders).
  - AnteHandler (placeholder, accepte tout).
  - ABCI hooks :
    - `BeginBlocker`,
    - `EndBlocker`,
    - `InitChainer` + `InitGenesis` + `DefaultGenesis`.

---

## 3. ModuleManager

- `CosmosModuleManager` créé.
- `SetOrderInitGenesis`, `SetOrderBeginBlockers`, `SetOrderEndBlockers` appelés (listes encore vides).
- Aucun module réel encore branché (auth, bank, staking, gov, evm, feemarket).

---

## 4. Génesis

- Type `GenesisState` (vide).
- `DefaultGenesis()` + `ValidateGenesis()` placeholders.
- `InitGenesis()` appelée depuis `InitChainer()` avec `DefaultGenesis()`.

---

## 5. Scripts & Config

- Dossier `config/` :
  - config, loader, paths, files, validator, env_loader.
- Dossier `scripts/` :
  - init, env, testnet, reset, status, keys, logs, validators, network,
    upgrade, health, tools, debug, security (placeholders).

---

## 6. À faire pour clôturer la Phase 2

- Wiring minimal des **modules Cosmos réels** (auth, bank, staking, gov).
- Intégration Ethermint (evm, feemarket) dans l’app et le ModuleManager.
- Ajout d’un `init` / `start` Cosmos plus standard dans le CLI.
- Génération d’un genesis minimal complet (avec modules déclarés).
- Vérification finale : `go build ./...` + notes des erreurs éventuelles.
