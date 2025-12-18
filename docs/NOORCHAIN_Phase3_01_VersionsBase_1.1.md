> ⚠️ **Status: TECH MIGRATION**  
> This document is being realigned to the **NOORCHAIN 2.0 Technical Baseline**.  
> Reference: `NOORCHAIN_Tech_Baseline_2.0_1.1.md`  
> Branch: `main-3` — Date: 2025-12-18

# NOORCHAIN — Phase 3  
## Fichier 19 — Versions de base (Cosmos SDK / CometBFT) — v1.1

### 🎯 Objectif du document

Ce document fixe **officiellement** les versions techniques de base utilisées pour construire  
NOORCHAIN 1.0 (Testnet puis Mainnet) avec le stack Cosmos.

L’objectif est d’éviter les changements permanents de versions et de stabiliser le socle technique  
avant de passer à la Phase 4 (wiring réel de l’App Cosmos + PoSS).

---

## 1. Versions retenues (officielles)

- **Langage Go**
  - Version cible : **Go 1.22**
  - Raison : version moderne, compatible avec les dernières libs Cosmos / CometBFT.

- **Cosmos SDK**
  - Version retenue : **v0.50.10**
  - Type : version **stable** et largement utilisée en production.
  - Rôles :
    - Gestion des comptes, balances, transactions.
    - Modules de base (bank, staking, governance, etc. si utilisés).
    - Intégration future du module PoSS (module personnalisé NOORCHAIN).

- **CometBFT**
  - Version retenue : **v0.38.17**
  - Rôle :
    - Moteur de consensus (remplace l’ancien Tendermint).
    - Gestion des blocs, consensus, P2P, state syncing.

---

## 2. Pourquoi ces versions ?

### 2.1 Cosmos SDK v0.50.10

- Version **long-term support** (LTS) et largement testée.
- Compatible avec les patterns modernes Cosmos (abciv1, modules plus propres).
- Suffisamment récente pour :
  - être maintenue,
  - profiter des améliorations,
  - mais **sans être expérimentale**.

### 2.2 CometBFT v0.38.17

- Successeur officiel de Tendermint.
- Version stable utilisée dans plusieurs projets Cosmos récents.
- Compatibilité naturelle avec Cosmos SDK 0.50.x.

---

## 3. Contraintes et règles internes NOORCHAIN

1. **Pas de retour en arrière**  
   - NOORCHAIN 1.0 devra rester sur :
     - Cosmos SDK **0.50.10**
     - CometBFT **0.38.17**
   - Sauf raison **critique** (bug majeur, faille de sécurité).

2. **Évolutions futures**
   - Une montée de version possible (0.51.x ou +) sera traitée :
     - après le lancement du Testnet stable,
     - via une **Phase dédiée “Upgrade”**, documentée à part.

3. **PoSS & Modules custom**
   - Le module PoSS (Proof of Signal Social) sera développé directement pour :
     - Cosmos SDK **0.50.10**
   - Pas de support prévu pour des versions ou forks antérieurs.

---

## 4. Impact sur la suite des phases

- **Phase 3 (actuelle)** :
  - Architecture, design des keepers, modules, layout fichiers.
  - Aucun code Cosmos réel encore ajouté.

- **Phase 4** :
  - Mise à jour de `go.mod` pour intégrer :
    - `github.com/cosmos/cosmos-sdk v0.50.10`
    - `github.com/cometbft/cometbft v0.38.17`
  - Création de l’App Cosmos réelle (BaseApp, ModuleManager, Keepers de base).
  - Préparation de l’intégration du module PoSS.

---

## 5. Résumé exécutif

- NOORCHAIN utilisera officiellement :
  - **Go 1.22**
  - **Cosmos SDK v0.50.10**
  - **CometBFT v0.38.17**

- Ce fichier sert de **référence** pour :
  - tout développeur qui rejoint le projet,
  - toute future décision technique,
  - toute montée de version.

**Décision :**  
✅ Les versions ci-dessus sont considérées comme le **socle officiel** de NOORCHAIN 1.0.
