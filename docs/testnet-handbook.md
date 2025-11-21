# NOORCHAIN — Testnet 1.0 Handbook

Version : 1.0  
Statut : Actif  
Scope : Cosmos SDK + PoSS + Économie NUR « en sommeil »

---

# 1. Introduction

Ce document décrit **toute la structure technique** du Testnet 1.0 de NOORCHAIN.

Il couvre :

- la configuration du testnet
- les modules actifs
- le genesis économique (5/5/5/5/80)
- le wiring Cosmos SDK
- le fonctionnement PoSS côté testnet
- la manière correcte de mettre à jour (adresses, modules, genesis)

Ce Handbook sert de guide pour :

- recréer un testnet propre à tout moment  
- préparer le passage futur au mainnet  
- aligner les développements avec le plan NOORCHAIN 1.0

---

# 2. Architecture du Testnet

Le testnet NOORCHAIN 1.0 repose sur :

- **Cosmos SDK (BaseApp)**
- **AppBuilder NOORCHAIN**
- **AppKeepers** : auth, bank, params, noorsignal
- **ModuleManager + Configurator**
- **Module PoSS (x/noorsignal)** — version avancée :
  - SubmitSignal
  - ValidateSignal
  - Curators (3 niveaux)
  - Limites quotidiennes
  - Barème PoSS
  - Rewards 70/30 (calculées mais pas distribuées)

EVM & IBC sont désactivés dans Testnet 1.0 (préparés mais non câblés).

---

# 3. Genesis du Testnet 1.0

Fichier : `testnet/genesis.json`

Champs principaux :

- `chain_id = "noorchain-testnet-1"`
- `genesis_time = "2025-01-01T00:00:00Z"`
- `app_state` :
  - **noorsignal** (config PoSS)
  - **bank** (économie ajoutée automatiquement par ApplyEconomicGenesis)

Aucun signal ni curator au genesis → normal.

---

# 4. Économie Testnet (après InitChain)

`ApplyEconomicGenesis` ajoute automatiquement les 5 accounts économiques :

| Compte               | Pourcentage | Montant en NUR | Montant en unur |
|----------------------|-------------|----------------|-----------------|
| Fondation            | 5%          | 14 989 622     | x1e6            |
| Dev Wallet (fondateur) | 5%       | 14 989 622     | x1e6            |
| Stimulus PoSS        | 5%          | 14 989 622     | x1e6            |
| Pré-vente (optionnelle) | 5%       | 14 989 622     | x1e6            |
| Réserve PoSS         | 80%         | 239 833 966    | x1e6            |

Ces adresses sont **placeholder** en testnet.  
Les vraies adresses seront injectées plus tard avant Mainnet 1.0.

---

# 5. Cycle de vie du Testnet

1. BaseApp est construite
2. Les stores sont montés
3. Les Keepers sont instanciés
4. ModuleManager + Configurator sont attachés
5. InitChainer :
   - lit `genesis.json`
   - applique l’économie NUR
   - initialise les modules
6. Le chain-node démarre
7. RPC / gRPC / Tendermint deviennent accessibles

---

# 6. Commandes CLI disponibles

### Query :

