# NOORCHAIN Core — Technical Architecture (Draft)

This document describes the **technical architecture** of the NOORCHAIN Core
repository. It explains the current structure, the planned modules, and how
Cosmos SDK + Ethermint + PoSS will be integrated.

> Status: draft — will be updated progressively as we integrate Cosmos SDK,
> Ethermint (EVM), and the custom PoSS module.

---

## 1. High-Level Technical Goals

- Build an independent Swiss blockchain: **NOORCHAIN**
- Based on:
  - **Cosmos SDK**
  - **CometBFT (Tendermint)** for consensus
  - **Ethermint** for EVM compatibility
- Add a custom module: **PoSS (Proof of Signal Social)**  
  → social signals validated on-chain  
  → NUR rewards distributed automatically (70/30 split)

### Token Economics (technical reference)
- Native token: **NUR**
- Fixed supply: **299,792,458 NUR**
- Halving: every **8 years** (time-based)
- Genesis distribution: **5% / 5% / 5% / 5% / 80%**

---

## 2. Current Repository Structure

```text
noorchain-core/
├── cmd/
│   └── noord/
│       └── main.go
│
├── app/
│   ├── app.go
│   └── encoding.go
│
├── docs/
│   ├── architecture.md
│   └── setup-dev.md
│
├── go.mod
├── Makefile
└── README.md
