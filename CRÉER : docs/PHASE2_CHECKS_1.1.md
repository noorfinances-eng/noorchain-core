# NOORCHAIN — Phase 2 Checks (1.1)

> Liste des vérifications manuelles à faire avant de considérer la Phase 2 comme “terminée”.

---

## 1. Vérifications Go

- [ ] Lancer `go mod tidy` à la racine du projet.
- [ ] Lancer `go build ./...` à la racine du projet.
- [ ] Noter toutes les erreurs éventuelles dans un fichier séparé (ex: `docs/PHASE2_ERRORS_1.1.md`).

---

## 2. Structure du projet

- [ ] Vérifier la présence des dossiers :
  - `cmd/`
  - `app/`
  - `x/`
  - `config/`
  - `scripts/`
  - `docs/`
- [ ] Vérifier que les fichiers de base existent :
  - `go.mod`
  - `README.md`
  - `Makefile` (si utilisé)

---

## 3. Application Cosmos (app/)

- [ ] `NOORChainApp` défini dans `cosmos_app.go`.
- [ ] `NewNOORChainApp()` :
  - [ ] logger minimal (log.NewNopLogger).
  - [ ] DB mémoire (memdb.NewDB).
  - [ ] BaseApp avec `CosmosTxDecoder`.
  - [ ] `LoadStores` appelé.
  - [ ] Encoding construit (CosmosEncodingConfig).
  - [ ] QueryRouter et MsgServiceRouter enregistrés.
  - [ ] AnteHandler enregistré.
  - [ ] BeginBlocker, EndBlocker, InitChainer enregistrés.

---

## 4. ABCI & Genesis

- [ ] `BeginBlocker`, `EndBlocker`, `InitChainer` présents dans `cosmos_block.go`.
- [ ] `GenesisState`, `DefaultGenesis`, `ValidateGenesis`, `InitGenesis` présents dans `cosmos_genesis.go`.
- [ ] `InitChainer` appelle `DefaultGenesis()` puis `InitGenesis()`.

---

## 5. ModuleManager & ModuleBasics

- [ ] `CosmosModuleManager` présent et utilisé dans `BuildCosmosApp()`.
- [ ] `SetOrderInitGenesis`, `SetOrderBeginBlockers`, `SetOrderEndBlockers` appelés (même sans modules).
- [ ] `ModuleBasics` défini (`module.BasicManager{}`) dans `cosmos_module_basics.go`.

---

## 6. Routers

- [ ] QueryRouter minimal défini et enregistré (`RegisterQueryRouter`).
- [ ] MsgServiceRouter minimal défini et enregistré (`RegisterMsgServiceRouter`).

---

## 7. Documentation Phase 2

- [ ] `docs/PHASE2_STATUS_1.1.md` présent et à jour.
- [ ] `docs/PHASE2_TODO_1.1.md` présent et à jour.
- [ ] `docs/PHASE2_CHECKS_1.1.md` (ce fichier) présent.

---

## 8. Avant de passer à la suite

- [ ] Confirmer que **PoSS n’apparaît dans aucun fichier** (Phase 4 uniquement).
- [ ] Confirmer que la Phase 2 se limite à :
  - structure Cosmos,
  - placeholders,
  - aucun module custom logique.
