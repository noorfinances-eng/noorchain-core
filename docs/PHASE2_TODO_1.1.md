# NOORCHAIN — Phase 2 TODO (Wiring minimal Cosmos/Ethermint)

> Version 1.1 — Liste des tâches restantes pour clôturer la Phase 2

---

## 1. Modules Cosmos de base

- [ ] Importer et déclarer les modules :
  - auth
  - bank
  - staking
  - gov
- [ ] Étendre `ModuleBasics` avec les Basic modules réels.
- [ ] Enregistrer les codecs de chaque module dans l’encoding (si nécessaire).

---

## 2. ModuleManager (réel, minimal)

- [ ] Ajouter les modules réels au `module.Manager`.
- [ ] Définir les ordres :
  - [ ] `SetOrderInitGenesis` avec les modules de base.
  - [ ] `SetOrderBeginBlockers` (staking, gov, etc. plus tard).
  - [ ] `SetOrderEndBlockers` (idem).

---

## 3. Keepers (squelette réel)

- [ ] Créer des keepers minimalistes pour :
  - [ ] auth
  - [ ] bank
  - [ ] staking
  - [ ] gov
- [ ] Injecter les keepers dans `NOORChainApp` (struct `CosmosKeepers`).
- [ ] Connecter les keepers au ModuleManager (Msg services + Queries).

---

## 4. Ethermint (EVM + FeeMarket)

- [ ] Importer Ethermint (modules evm + feemarket).
- [ ] Ajouter evm + feemarket à `ModuleBasics`.
- [ ] Ajouter evm + feemarket au `ModuleManager`.
- [ ] Créer des keepers minimalistes pour EVM et FeeMarket.
- [ ] Configurer les params de base Ethermint.

---

## 5. Génesis minimal

- [ ] Étendre `GenesisState` pour :
  - [ ] auth
  - [ ] bank
  - [ ] staking
  - [ ] gov
  - [ ] evm
  - [ ] feemarket
- [ ] Ajouter `DefaultGenesis` réel pour chaque module.
- [ ] Appeler les `InitGenesis` de chaque module depuis `InitGenesis` principal.
- [ ] Garder la logique PoSS totalement absente (Phase 4 uniquement).

---

## 6. CLI minimal Cosmos

- [ ] Ajouter une commande `init` (ou `unsafe-init`) pour générer un genesis minimal.
- [ ] Ajuster `start` pour lancer la NOORCHAIN avec ce genesis.
- [ ] Laisser toutes les futures commandes PoSS pour la Phase 4.

---

## 7. Vérifications finales

- [ ] `go mod tidy`
- [ ] `go build ./...` (sans erreur)
- [ ] Noter dans la doc les éventuelles limitations connues de la Phase 2 :
  - pas de PoSS,
  - pas de staking réel,
  - pas de gouvernance réelle,
  - pas de transactions utiles encore.
