# NOORCHAIN Core — Technical Architecture (Draft)

This document describes the **technical architecture** of the NOORCHAIN core
repository. It is focused on code structure, modules, and how the Cosmos SDK
and Ethermint (EVM) will be wired in future steps.

> Status: skeleton only — will be updated as we integrate Cosmos SDK, Ethermint and PoSS.

---

## 1. High-level goals

- Independent Swiss blockchain (**NOORCHAIN**)
- Built on **Cosmos SDK** + **CometBFT**
- EVM compatibility via **Ethermint**
- Custom consensus layer logic via a PoS-like validator set,
  plus a **Proof of Signal Social (PoSS)** module that mints rewards
  according to social signals instead of hash power.

Token:

- Native token: `NUR`
- Fixed supply: `299,792,458 NUR`
- Halving every 8 years (time-based, not block-based)
- Genesis distribution: 5% / 5% / 5% / 5% / 80% (Foundation, Dev, PoSS Stimulus, Pre-sale, PoSS)

---

## 2. Repository structure (current)

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
│   └── architecture.md   (this file)
│
├── go.mod
├── Makefile
└── README.md
