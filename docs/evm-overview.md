# NOORCHAIN — EVM & RPC Overview (Draft 1.0)

Ce document décrit comment NOORCHAIN utilisera l’EVM (Ethermint) et
le JSON-RPC pour être compatible avec Metamask et les dApps EVM.

---

## 1. Objectif EVM

- Permettre aux utilisateurs d’utiliser NOORCHAIN comme une L1 EVM :
  - Metamask
  - WalletConnect
  - DApps EVM (DEX, NFT, etc.)
- Garder la logique PoSS et l’économie NUR au centre du système.
- Rester compatible avec le cadre suisse “Legal Light”.

---

## 2. Stack technique cible

- Cosmos SDK + CometBFT
- Module EVM (Ethermint)
- Module Feemarket (EIP-1559 style)
- Module Bank pour les soldes NUR (denom : `unur`)
- RPC standard :
  - Tendermint RPC (26657)
  - gRPC
  - EVM JSON-RPC (8545 / 8546)

---

## 3. Denom EVM

- Denom interne : `unur`
- 1 NUR = 1 000 000 `unur`
- L’EVM utilisera `unur` comme :
  - gas denom
  - unité de base pour les transferts EVM

---

## 4. Configuration EVM (principes)

Paramètres principaux :

- `evm_denom` = `unur`
- `enable_create` = true (déploiement de smart-contracts)
- `enable_call` = true (appel de contrats)
- `chain_id` EVM aligné avec NOORCHAIN (exemple : 7777)
- `gas_price` minimal défini de manière raisonnable

Les valeurs exactes seront décidées juste avant le Mainnet 1.0.

---

## 5. JSON-RPC & Metamask (cible)

Exemple de configuration Metamask (cible, à adapter en prod) :

- Network Name : NOORCHAIN Mainnet
- New RPC URL : `https://rpc.noorchain.com`
- Chain ID : `7777` (exemple)
- Currency Symbol : `NUR`
- Block Explorer URL : (à définir plus tard)

Pour le Testnet :

- Network Name : NOORCHAIN Testnet
- RPC URL : `https://rpc-testnet.noorchain.com` (ou IP/localhost)
- Chain ID : ex `7778`
- Symbol : `tNUR`
- Explorer Testnet : à définir

---

## 6. Intégration avec PoSS

- Le PoSS reste un module Cosmos (x/noorsignal)
- Les récompenses sont en `unur` (NUR)
- Les contrats EVM peuvent interagir avec PoSS plus tard via :
  - précompiles
  - messages spécifiques
  - ou des modules passerelles

La priorité actuelle : avoir un EVM standard compatible Metamask,
puis ajouter des intégrations avancées dans NOORCHAIN 2.0.

---

## 7. Statut actuel (fin 2025)

- Design EVM : défini
- Modules EVM / Feemarket : prévus dans l’architecture
- Genesis template : stub EVM/feemarket prévu
- Aucun code EVM encore intégré → pour éviter les erreurs tant que le
  reste de la chaîne n’est pas totalement stabilisé.

---

## 8. Étapes futures pour activer l’EVM

1. Ajouter les modules EVM / Feemarket (Ethermint) dans l’app Go.
2. Ajouter la configuration EVM complète dans le genesis.
3. Exposer le JSON-RPC EVM.
4. Tester :
   - deployment de contrat
   - transfers EVM
   - compatibilité Metamask
5. Documenter un guide “NOORCHAIN + Metamask”.

Ce document doit être mis à jour au moment exact où l’EVM sera codée.
