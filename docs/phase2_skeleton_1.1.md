# NOORCHAIN – Phase 2 (Skeleton minimal)

Date : 2025-11-27

## Objectif de la Phase 2

Mettre en place un squelette minimal de l’application NOORCHAIN sans dépendances Cosmos complexes, juste assez pour :

- Avoir une structure `app/` propre.
- Avoir un type d’application simple (`App`).
- Pouvoir compiler un binaire `noord`.
- Afficher les métadonnées de base (nom, version).

## Fichiers Go principaux

### 1. `app/types.go`
- Définit les constantes de base :
  - `AppName`      = `"NOORCHAIN"`
  - `AppVersion`   = `"0.1.0"`
  - `ChainID`      = `"noorchain-1"`
  - `BondDenom`    = `"unur"`
- Sert de point central pour les valeurs globales de la chaîne.

### 2. `app/app_type.go`
- Type `AppType` avec :
  - `Name`
  - `Version`
- Fonction `DefaultAppType()` qui renvoie les métadonnées par défaut de NOORCHAIN.

### 3. `app/config.go`
- Contient la configuration de base (denom / décimales) pour NOORCHAIN.
- Utilisé plus tard pour garder la cohérence entre :
  - le genesis,
  - les modules PoSS,
  - les dApps (explorers, wallets, etc.).

### 4. `app/encoding.go`
- Fournit `EncodingConfig` (Amino, InterfaceRegistry, Marshaler, TxConfig).
- Pour l’instant : configuration minimale, sans modules personnalisés.
- En Phase 4, le module PoSS sera ajouté dans l’`InterfaceRegistry`.

### 5. `app/bech32.go`
- Définit le préfixe Bech32 : `noor`.
- Prépare les préfixes :
  - comptes : `noor1...`
  - validateurs : `noorvaloper1...`
  - consensus : `noorvalcons1...`
- La fonction `ConfigureBech32Prefixes()` sera appelée au moment où on branchera réellement Cosmos SDK.

### 6. `app/app.go`
- Définit une structure minimale `App` contenant :
  - `Info` (métadonnées : nom, version)
- Fonction `NewApp()` qui retourne une instance simple :
  - Sans BaseApp
  - Sans Cosmos SDK
  - Juste assez pour que le binaire ait quelque chose à instancier.

### 7. `cmd/noord/main.go`
- Point d’entrée du binaire `noord`.
- Crée l’application via `noorapp.NewApp()`.
- Affiche dans le terminal :
  - `NOORCHAIN node placeholder (Phase 2 – skeleton only)`
  - `App Name: NOORCHAIN`
  - `App Version: 0.1.0`

## Résultat actuel

- `go build -o noord ./cmd/noord` ✅
- `./noord` ✅ affiche les métadonnées (nom + version).
- Aucun import Cosmos-SDK complexe dans la Phase 2.
- La suite (Phase 3 / Phase 4) pourra :
  - ajouter `BaseApp`,
  - brancher Cosmos SDK + CometBFT proprement,
  - introduire PoSS sans casser ce squelette.

## Prochaine étape

- Phase 3 : définir proprement la stratégie d’intégration Cosmos SDK / CometBFT :
  - choix versions stables,
  - plan de fichiers (keepers, modules, génesis),
  - sans toucher à l’ADN économique (PoSS, 5/5/5/5/80, halving 8 ans).
