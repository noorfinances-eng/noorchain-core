# NOORCHAIN — Phase 3  
## Fichier 20 — Architecture officielle du projet — v1.1

### 🎯 Objectif du document

Ce document définit l’architecture **officielle, stable et complète** du projet `noorchain-core`  
pour les Phases 3 → 9 (Testnet puis Mainnet).

Il sert de référence unique pour :
- organiser les fichiers,
- éviter le chaos,
- s’assurer de la cohérence de toute la suite du développement (PoSS, genesis, RPC, CLI, EVM…).

---

# 1. Racine du projet
noorchain-core/
│
├── app/ → cœur de l’App Cosmos (BaseApp, keepers, modules, config)
├── cmd/ → commandes CLI (noord)
├── x/ → modules personnalisés (ex : noorsignal pour le PoSS)
├── proto/ → définitions Protobuf
├── testnet/ → fichiers genesis + scripts testnet
├── docs/ → documentation officielle
└── go.mod / go.sum → dépendances

---

# 2. Détail par dossier

## 2.1 `app/` — Le cœur de NOORCHAIN

Contiendra à partir de la Phase 4 :
app/
│
├── app.go → Définition principale de l’application
├── app_builder.go → Construction de l'app (extension future)
├── config.go → Config NOORCHAIN (Bech32, denom, version)
├── encoding.go → Encodage (Amino + Protobuf)
├── keepers.go → Déclaration globale des keepers
├── module_manager.go → Wiring complet des modules
├── params.go → Paramètres globaux
└── types.go → Types internes

📌 **Remarque** :  
Tout reste minimal en Phase 3 — on écrit seulement la structure, pas le contenu Cosmos.

---

## 2.2 `cmd/noord/`

Le binaire principal de la blockchain :
cmd/noord/
│
└── main.go → Point d’entrée du node

Plus tard (Phase 5+) :
- commandes CLI (init, collect-gentxs, start, unsafe-reset-all)

---

## 2.3 `x/` — Modules personnalisés

Lieu officiel des modules NOORCHAIN.
x/
│
└── noorsignal/ → Module PoSS (Proof of Signal Social)
├── keeper/
├── types/
├── module.go
├── genesis.go
├── msgs.go
└── events.go

En Phase 4 :
- création du **squelette vide** `x/noorsignal`

---

## 2.4 `proto/`

Les définitions `.proto` utilisées par :
- App
- modules
- PoSS

Exemple futur :
proto/noorsignal/
├── tx.proto
├── query.proto
├── genesis.proto
└── types.proto

---

## 2.5 `testnet/`

En Phase 6 :
testnet/
│
├── genesis.json
├── genesis_distribution.json
├── config.toml
└── addrbook.json

---

## 2.6 `docs/`

Contient **toute la documentation officielle**, versionnée en *1.1* :
docs/
│
├── NOORCHAIN_Phase3_01_VersionsBase_1.1.md
├── NOORCHAIN_Phase3_02_ArchitectureProjet_1.1.md ← ce fichier
├── NOORCHAIN_Phase3_03_KeepersPlan_1.1.md (à venir)
├── NOORCHAIN_Phase3_04_ModuleManagerPlan_1.1.md (à venir)
└── ...

---

# 3. Règles architecturales officielles

### 3.1 Pas de fichiers Cosmos non utilisés
Toute importation Cosmos SDK doit venir **uniquement** lorsque nécessaire en Phase 4.

### 3.2 Le module PoSS (`x/noorsignal`) est **isolé**
Aucun code PoSS ne sera placé dans :
- `app/`
- `cmd/`
- `runtime/`
- `store/`

### 3.3 Chaque fichier `.go` doit rester simple en Phase 3
Les keepers, modules et BaseApp seront remplis **uniquement** en Phase 4.

### 3.4 La documentation (`docs/`) reste la vérité officielle
Si un fichier n’est pas décrit dans les docs :  
❌ il n’est pas créé.

### 3.5 Pas de code généré sans décision explicite
- pas de `buf`
- pas de compilation proto
- pas de `makefile`
… tant que je ne te le demande pas.

---

# 4. Résumé exécutif

Le projet `noorchain-core` est organisé en **6 dossiers principaux** :

1. `app/` — cœur de l’application Cosmos  
2. `cmd/noord/` — binaire node  
3. `x/` — modules personnalisés (dont PoSS)  
4. `proto/` — définitions Protobuf  
5. `testnet/` — configuration réseau  
6. `docs/` — documentation officielle  

Cette architecture est désormais **officielle**, stable et versionnée.

---

# 5. Statut

**Décision validée** :  
Cette architecture est considérée comme le socle officiel de la Phase 3.

