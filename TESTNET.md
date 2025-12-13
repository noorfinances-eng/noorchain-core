# NOORCHAIN — Testnet (Public)

NOORCHAIN is a Swiss social blockchain based on the **Proof of Signal Social (PoSS)** model.

This repository contains the **core implementation of the NOORCHAIN blockchain**, built with **Cosmos SDK + Ethermint (EVM)**.

The Testnet is provided for:
- technical validation,
- educational pilots,
- protocol review,
- early ecosystem testing.

NOORCHAIN follows a **Legal Light CH** framework:
- non-custodial,
- non-financial,
- no promised returns,
- utility-only token model.

---

## Status

- Core blockchain: ✅ implemented
- PoSS module: ✅ implemented (disabled by default)
- Testnet: ✅ operational
- Mainnet: ⏳ not launched

This Testnet is **not intended for production use**.

---

## Architecture Overview

- **Consensus base:** Cosmos SDK
- **EVM compatibility:** Ethermint
- **Social layer:** PoSS (Proof of Signal Social)
- **Token supply:** fixed (299,792,458 NUR)
- **Inflation:** none (PoSS uses a capped reserve)
- **Governance:** on-chain + multi-sig (3/5)

---

## Proof of Signal Social (PoSS)

PoSS is a **social validation mechanism**, not a financial consensus.

Key principles:
- Signals represent **verified social actions**
- No staking, no APR, no yield
- Rewards are symbolic and capped
- Structural split: **70% participant / 30% curator**
- Halving every 8 years
- PoSS can be enabled or disabled via governance

PoSS is implemented as a native Cosmos module (`x/noorsignal`).

---

## Testnet Information

### Network
- Chain ID: `noorchain-testnet`
- Token denom: `unur` (display: NUR)
- EVM enabled: yes
- PoSS enabled: no (default)

### Access
At this stage, the testnet is:
- launched internally,
- available for controlled pilots,
- documented for transparency.

Public RPC endpoints and explorers will be published progressively.

---

## Running a Local Node (Developer Preview)

> For developers and auditors only.

Requirements:
- Go ≥ 1.22
- Make

```bash
git clone https://github.com/noorfinances-eng/noorchain-core.git
cd noorchain-core
go build ./...
