# NOORCHAIN — Ethermint / EVM Integration Plan (Option A — Full EVM)

Objectif : intégrer **Ethermint complet** dans NOORCHAIN, pour :

- permettre l'utilisation de MetaMask,
- déployer des smart-contracts Solidity,
- utiliser un explorer type Blockscout,
- garder la compatibilité EVM tout en ayant PoSS comme moteur natif.

Ce document décrit **le plan technique d'intégration**, sans encore écrire
le code Go. Il sert de feuille de route pour la Phase EVM.

---

## 1. Modules Ethermint à intégrer

NOORCHAIN utilisera les modules Ethermint classiques :

- `x/evm`        → exécution EVM (smart-contracts, opcodes, etc.)
- `x/feemarket`  → gestion des frais EVM (base fee, EIP-1559, etc.)

Ils seront intégrés dans :

- `app/keepers.go`       → `EvmKeeper`, `FeeMarketKeeper`
- `app/app_builder.go`   → instanciation réelle des keepers
- `app/module_manager.go` → ajout des AppModules EVM
- `app/modules_layout.go` → ordre BeginBlock/EndBlock/InitGenesis
- `testnet/genesis.json` → section `evm` + `feemarket` complétées

---

## 2. Rappel : état actuel côté EVM (Phase 1)

À ce stade (avant intégration complète) :

- `modules.go` contient déjà :

  ```go
  const (
    ModuleEvm       = "evm"
    ModuleFeeMarket = "feemarket"
  )
