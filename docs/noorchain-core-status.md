# NOORCHAIN — Core Status (November 2025)
Version officielle : 1.0 — Récapitulatif complet

Ce document regroupe l’état exact du projet : technique, économique, PoSS,
EVM, testnet et étapes restantes.  
Il sert de snapshot de référence avant d’attaquer les phases suivantes.

---

# 1. Architecture technique (Cosmos SDK + Ethermint)

NOORCHAIN utilise :

- **Cosmos SDK**
- **CometBFT**
- **BaseApp complète**
- **ModuleManager fonctionnel**
- **StoreKeys complets**
- **Keepers réels** :  
  - `AccountKeeper`  
  - `BankKeeper`  
  - `ParamsKeeper`  
  - `NoorSignalKeeper` (PoSS)
- **PoSS** (Proof of Signal Social) comme module natif (`x/noorsignal`)
- **Future** : Ethermint complet (`x/evm`, `x/feemarket`)

## 1.1 Modules déjà présents dans l’app

- PoSS (x/noorsignal)
- Auth (accounts)
- Bank (balances)
- Params (global params)
- Layout complet des modules BeginBlock / EndBlock / InitGenesis

## 1.2 Fichiers Cosmos déjà finalisés

- `app/app.go`
- `app/app_cosmos_skeleton.go`
- `app/app_builder.go`
- `app/keepers.go`
- `app/module_manager.go`
- `app/modules.go`
- `app/modules_layout.go`
- `app/genesis.go`
- `cmd/noord/main.go`
- `x/noorsignal/*` (keeper, msg_server, types)

L’architecture Cosmos est propre, stable, extensible, et prête pour :
- l’intégration BankKeeper 100%
- l’intégration EVM (Ethermint)
- le testnet

---

# 2. Module PoSS (x/noorsignal)

Le module PoSS est entièrement fonctionnel en mode “logique” :

- SubmitSignal OK
- ValidateSignal OK
- Stores OK (signal, curator, config, counters)
- Rewards PoSS calculés (BaseReward, Era, 70/30)
- Daily limit OK
- Admin Msg (AddCurator, RemoveCurator, SetConfig)
- genesis PoSS OK
- Events PoSS OK
- QueryServer PoSS OK (gRPC + CLI)

En attente :  
- branchement BankKeeper (argent réel)
- débit / crédit du compte réserve PoSS
- synchronisation avec le genesis économique

---

# 3. Économie NOORCHAIN (Genesis économique 5/5/5/5/80)

Distribution officielle (déjà enregistrée dans la mémoire du projet) :

- **5 % Fondation NOOR** (association)
- **5 % Noor Dev (Walid)** ← ton wallet perso Bech32
- **5 % PoSS Stimulus**
- **5 % Pré-vente optionnelle (avec vesting + multisig)**
- **80 % Réserve PoSS (minage social)**

Token natif :
- symbol = **NUR**
- denom interne = **unur**
- supply fixe = **299 792 458 NUR**

Genesis économique est **préparé**, mais pas encore activé dans le code.

---

# 4. Testnet 1.0 — État actuel

Dossier existant :

