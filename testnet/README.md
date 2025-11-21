# NOORCHAIN — Testnet 1.0 (Planning & Structure)

Ce dossier contient tout ce qui est nécessaire pour préparer et lancer
le **premier testnet NOORCHAIN** (Testnet 1.0), basé sur :

- Cosmos SDK
- Ethermint (EVM complet, à intégrer en Phase 2)
- Module PoSS (`x/noorsignal`)
- Token natif `NUR` (denom `unur`)
- Distribution économique 5 / 5 / 5 / 5 / 80

> ⚠️ IMPORTANT  
> Ce dossier sert de **base de travail**.  
> Le testnet n’est pas encore destiné à être lancé en production.  
> Il manque encore les **vraies adresses bech32**, le câblage complet
> du BankKeeper “argent réel” et l’intégration finale Ethermint.

---

## 1. Fichiers principaux du testnet

### 1.1 `testnet/genesis.json`

Fichier genesis principal du testnet.  
Il contient :

- `chain_id = "noorchain-testnet-1"`
- les modules Cosmos classiques (auth, bank, params, etc.)
- la section PoSS (`noorsignal`) avec :
  - `config` (BaseReward, 70/30, max_signals_per_day, eraIndex)
  - `curators` initiaux (optionnel)
- des sections `evm` / `feemarket` pour la future intégration Ethermint

Certaines adresses sont encore des **placeholders** du type :

```json
"address": "TO_FILL_FOUNDATION_BECH32_NOOR1..."
