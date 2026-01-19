GENESIS_Overview_1.1_EN.md (version correcte, sans blocks)
NOORCHAIN — Genesis Overview

Version 1.1 — Official Document
Last Updated: 2025-XX-XX
Language: EN

1. Abstract

This document provides the official overview of the NOORCHAIN 1.0 genesis.
It defines the immutable supply, the economic allocation model, the governance constraints, the PoSS structural rules, and the role of the five genesis wallets.

It is part of the Genesis Pack 1.1, intended for developers, auditors, partners, institutions, and investors.

2. Genesis Philosophy

The NOORCHAIN genesis follows three guiding principles:

1. Immutable supply
The token supply is fixed forever at 299,792,458 NUR (symbolic reference to the speed of light).

2. Swiss ethical blockchain model
The chain operates under Legal Light CH principles: transparency, non-speculation, no custody, and no promised financial returns.

3. Human-centric consensus
Consensus incentives come from PoSS (Proof of Signal Social), which distributes part of a pre-allocated reserve based on validated social actions.

3. Total Supply (Immutable)

Total supply: 299,792,458 NUR
Decimals: 18
Inflation: 0% permanently
Future minting: permanently disabled
Supply modification via governance: forbidden

The supply cannot change at any moment during the life of the blockchain.

4. Economic Allocation (5 / 5 / 5 / 5 / 80)

5% — NOOR Foundation (public governance, transparency, operations)

5% — Noor Dev Sàrl (development entity, 5% liquid + 10% vested outside genesis)

5% — PoSS Stimulus Pool (early adoption incentives)

5% — Optional Pre-sale (with vesting, Swiss rules)

80% — PoSS Mintable Reserve (social mining, halving every 8 years)

This model is fixed and cannot be altered after genesis.

5. Structural PoSS Rules Included in Genesis

The genesis contains the structural PoSS parameters that cannot be changed:

Reward split: 70% participant / 30% curator

Halving cycle: every 8 years

Source of rewards: pre-allocated PoSS Reserve only

No inflation allowed

Daily limits active

Curator tiers (Bronze / Silver / Gold)

Signals weights (Micro-donation, Participation, Content, CCN)

These rules remain constant throughout the chain’s lifetime.

6. Governance Foundations (Immutable Aspects)

The genesis defines the following constraints:

Multi-sig: 3/5 threshold (Foundation)

No token custody by Foundation or Sàrl

No PSP internal to the blockchain

Supply and allocations cannot be changed by governance

PoSS structural rules cannot be changed

Only parameters (not structure) may evolve through governance proposals

This ensures legal compliance and long-term stability.

7. Genesis Wallets (Placeholders Until Phase 7)

Five genesis wallets must be inserted during Phase 7:

Foundation Wallet

Noor Dev Sàrl Wallet

PoSS Stimulus Wallet

Optional Pre-sale Wallet

PoSS Reserve Wallet

The addresses will be integrated into:

genesis.json (testnet & mainnet)

x/noorsignal/types/addresses.go

Governance & Legal documents

Genesis Pack PDF

8. Genesis Security Model

The genesis ensures:

deterministic validator set creation

deterministic state root

PoSS disabled initially on mainnet (security-first)

no risk of inflation

no external dependency for supply or economic rules

It is designed for Swiss Legal Light compliance and for long-term sustainability.

9. Purpose of This Document

This file serves as:

the high-level reference for understanding the genesis

the economic and governance foundation

the link between legal documentation and technical genesis.json

the base for the Genesis Pack PDF

10. Versioning

Genesis Overview — Version 1.1
Updates require governance approval.
Structural elements cannot be modified.
