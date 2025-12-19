> ⚠️ **Status: TECH MIGRATION**  
> This document is being realigned to the **NOORCHAIN 2.0 Technical Baseline**.  
> Reference: `NOORCHAIN_Tech_Baseline_2.0_1.1.md`  
> Branch: `main-3` — Date: 2025-12-18

# NOORCHAIN — Technical Baseline 2.0 (v1.1)

**Status:** ACTIVE — Technical Source of Truth  
**Branch:** main-3  
**Date:** 2025-12-18

---

## 1. Purpose of this document

This document defines the **official technical baseline for NOORCHAIN 2.0**.

It replaces all previous technical references related to the NOORCHAIN 1.0 implementation,
without altering the vision, ethics, or legal positioning of the project.

Any future technical implementation MUST comply with this baseline.

---

## 2. Target technical stack (locked)

NOORCHAIN 2.0 is built on the following stack:

- **Go:** 1.22.x
- **Cosmos SDK:** v0.50.x
- **Consensus:** CometBFT v0.38.x
- **EVM:** Evmos (native EVM, Cosmos-compatible)
- **Interfaces:** gRPC, Protobuf, ABCI++

No downgrade of these versions is permitted after validation of this baseline.

---

## 3. Architectural principles

NOORCHAIN 2.0 follows these core principles:

- **Strict modularity** (loosely coupled modules, no monolithic design)
- **Native upgradability** (upgrade handlers, migrations, future governance)
- **EVM compatibility** without compromising Cosmos architecture
- **Clear separation** between:
  - protocol layer
  - economic logic
  - application / usage layer

No on-chain dApps are developed before Mainnet.

---

## 4. Public testnet principle

The NOORCHAIN 2.0 public testnet is intended for:

- observable technical proof,
- consensus stability validation,
- RPC / gRPC / EVM compatibility testing.

It has **no financial purpose**.
Tokens and economic mechanisms remain conceptual at this stage.

---

## 5. Relationship with NOORCHAIN 1.0

NOORCHAIN 2.0:

- does **not migrate** code from NOORCHAIN 1.0,
- **replaces** the 1.0 technical implementation,
- **fully preserves** the original vision, ethics, and founding principles.

NOORCHAIN 1.0 is considered a learning and experimental phase.

---

## 6. Scope and authority

This document:

- prevails over any previous technical documentation,
- applies to Phases 2 through 10 of the roadmap,
- contains no financial or investment promises.

