# NOORCHAIN — Testnet Toolbox

Ce dossier contient les fichiers et outils liés au **testnet** de NOORCHAIN.
Pour l’instant, il s’agit surtout de squelettes documentaires.  
Le vrai genesis final et les distributions seront définis plus tard, quand on aura les **adresses réelles** (fondation, fondateur, réserve PoSS, validateurs, etc.).

---

## 1. Fichiers présents

### `genesis_skeleton.json`

- **But :** servir de *modèle documentaire* pour le futur `genesis.json`.
- **Statut :** non utilisé directement pour démarrer un nœud, uniquement une base de travail.
- **Contenu :**
  - `chain_id = "noorchain-testnet-1"`
  - Modules principaux documentés :
    - `auth` (params de base, comptes vides)
    - `bank` (balances et supply vides)
    - `staking` (bond_denom = `unur`, params génériques)
    - `evm` (params Ethermint de base, chain_config London activable plus tard)
    - `feemarket` (EIP-1559, à ajuster avant le vrai testnet)
    - `noorsignal` (module PoSS, pour l’instant vide)

---

## 2. Fichiers prévus (prochaines étapes)

Ces fichiers **n’existent pas encore**, ils seront créés plus tard dans la Phase Testnet :

### `genesis.json`

- Le **genesis réel** utilisé par les nœuds de NOORCHAIN Testnet.
- Sera généré à partir :
  - du modèle `genesis_skeleton.json`,
  - des adresses bech32 réelles (fondation, fondateur, réserve PoSS, etc.),
  - de la distribution initiale (5 % / 5 % / 5 % / 5 % / 80 % en `unur`).

### `genesis_distribution.json`

- Fichier de travail pour décrire **qui reçoit quoi** au lancement :
  - Fondation NOORCHAIN
  - Noor Dev (fondateur)
  - Réserve PoSS Stimulus
  - Pré-vente optionnelle (si utilisée)
  - Réserve PoSS mintable (80 %)
- Permet de garder une **trace claire, diff-able et audit-able** de la distribution économique.

---

## 3. Règles officielles à respecter pour le Testnet

Rappel des contraintes de NOORCHAIN 1.0 :

- **Supply fixe :** `299 792 458 NUR` (unur au niveau technique).
- **Modèle économique :**  
  - 5 % Fondation NOOR  
  - 5 % Noor Dev Sàrl (ou wallet fondateur temporaire)  
  - 5 % PoSS Stimulus  
  - 5 % Pré-vente optionnelle  
  - 80 % PoSS mintables (Proof of Signal Social)
- **Halving :** tous les 8 ans.
- **Legal Light CH :**
  - Aucun yield garanti.
  - Aucune promesse de performance.
  - Transparence totale sur la distribution et les règles PoSS.

Ces règles devront être **strictement respectées** au moment où nous remplirons `genesis.json` et `genesis_distribution.json`.

---

## 4. Prochaine étape (technique)

La prochaine phase Testnet consistera à :

1. Générer les **5 adresses bech32 réelles** (fondation, fondateur, PoSS Stimulus, pré-vente, réserve PoSS).
2. Synchroniser ces adresses dans :
   - `testnet/genesis.json`
   - `testnet/genesis_distribution.json`
   - `x/noorsignal/types/addresses.go`
3. Vérifier la cohérence entre :
   - le code Go (PoSS, BankKeeper, StakingKeeper, etc.),
   - et les fichiers `genesis*.json`.

Ce README sert simplement de **boussole** pour ne pas se perdre entre les phases.

---
