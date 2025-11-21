# NOORCHAIN — Pre-Launch Checklist (Testnet 1.0)

Ce document liste les éléments **obligatoires** avant de lancer
le premier vrai testnet NOORCHAIN.  
Il sert de guide pour éviter les oublis et garantir un lancement propre.

---

# 1. Génération & synchronisation des 5 adresses bech32

Les cinq adresses maîtres doivent être générées puis synchronisées partout :

### 1.1 Adresses à générer
- Fondation NOOR (association)
- Fondateur / Dev (Walid)
- PoSS Stimulus Pool
- Pré-vente optionnelle
- Réserve PoSS (80 %)

### 1.2 À mettre à jour dans :
- `docs/testnet-addresses-master.json`
- `testnet/genesis.json`
- `testnet/genesis_distribution.json`
- `x/noorsignal/types/addresses.go`

---

# 2. Vérification du Genesis économique

### 2.1 Supply totale
- 299 792 458 NUR  
- denom = `unur` (1 NUR = 1,000,000 unur)

### 2.2 Distribution 5 / 5 / 5 / 5 / 80
- Fondation : 14 989 622.9 NUR  
- Walid (Founding wallet) : 14 989 622.9 NUR  
- Stimulus : 14 989 622.9 NUR  
- Pré-vente optionnelle : 14 989 622.9 NUR  
- Réserve PoSS : 239 833 966.4 NUR

(Les chiffres exacts en `unur` seront insérés dans `genesis_distribution.json`.)

### 2.3 Bank balances
- Tous les comptes initiaux doivent être inscrits dans `bank.balances`.

---

# 3. Configuration PoSS

### 3.1 Vérifier que le genesis contient :
- `base_reward`
- `participant_share = 70`
- `curator_share = 30`
- `max_signals_per_day`
- `era_index`

### 3.2 Vérifier que :
- Le Keeper charge la config au démarrage  
- Les curators initiaux (optionnels) sont bien pris en compte

---

# 4. Intégration EVM (Phase 2 nécessaire avant testnet)

Avant un testnet complet, Ethermint doit être câblé :

### 4.1 Stores EVM montés (OK)
### 4.2 Keepers EVM déclarés (OK)
### 4.3 Restent à faire :
- Instanciation réelle `EvmKeeper`
- Instanciation réelle `FeeMarketKeeper`
- ModuleManager : ajouter `evm` et `feemarket`
- Genesis EVM :
  - chainConfig
  - baseFee
  - denom = `unur`
- RPC EVM activé

---

# 5. BankKeeper “récompenses PoSS réelles”

### À implémenter avant testnet interactif :
- Débiter PoSS Reserve
- Créditer participant / curator via `SendCoins`
- Vérifier `blockedAddrs` + comptes module
- Ajouter compte module `PossReserve` (module account)

---

# 6. Build & Execution

### 6.1 Compilation binaire
