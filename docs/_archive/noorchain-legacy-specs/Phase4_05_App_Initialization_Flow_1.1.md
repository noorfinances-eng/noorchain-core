# NOORCHAIN — Phase 4A  
## App Initialization Flow (Cosmos SDK + Ethermint + PoSS)  
### Version 1.1 — Architecture Only (No Code)  
### Last Updated: 2025-12-03  

---

## 1. Purpose of This Document

This document defines the complete **initialization flow** of the NOORCHAIN 1.0 application.

It describes:

- how the blockchain app is constructed  
- in which order components must be instantiated  
- how modules hook into the lifecycle  
- how the app transitions from “empty” to “running chain”  
- how PoSS fits into the init process  

This is the **reference** for building the real constructor in `app/app.go` during Phase 4C (implementation & testnet).

---

## 2. Initialization Flow Overview

Initialization occurs in three main phases:

1. **Application Construction**  
2. **ModuleManager Wiring**  
3. **Genesis Initialization**

Each phase must be executed in the **correct order** to avoid:

- unstable state  
- missing stores  
- EVM or PoSS panics at runtime  

---

## 3. Phase 1 — Application Construction

The app constructor must perform the following steps in order:

---

### Step 1 — Build Encoding Config

Construct:

- Interface Registry  
- Amino (legacy) codec  
- Protobuf codec  
- TxConfig  
- JSON marshaler  

**Purpose:**  
Enable message processing, signing, decoding, and transaction routing.

---

### Step 2 — Create BaseApp

Build `BaseApp` with:

- app name  
- logger  
- database handle  
- TxDecoder / TxEncoder  
- interface registry  

**Purpose:**  
Provide the core ABCI engine for NOORCHAIN.

---

### Step 3 — Define Store Keys

Define `KVStoreKey` / `TransientStoreKey` for:

- `auth`  
- `bank`  
- `staking`  
- `gov`  
- `evm`  
- `feemarket`  
- `noorsignal`  
- `params`  

**Purpose:**  
Prepare all module stores to mount persistent chain state.

---

### Step 4 — Mount Stores

Use `MountKVStore`, `MountTransientStore`, etc. for:

- all module KV stores  
- transient stores (where required, e.g. feemarket transient store)  

**Purpose:**  
Ensure all module states are persisted and accessible by the keepers.

---

### Step 5 — Create Keepers (Strict Order)

Instantiate keepers in this **strict order** (aligned with the Keeper System Design):

1. **AccountKeeper**  
2. **BankKeeper**  
3. **StakingKeeper**  
4. **GovKeeper**  
5. **ParamsKeeper** (with subspaces for all modules)  
6. **EVMKeeper**  
7. **FeeMarketKeeper**  
8. **PoSSKeeper** (`x/noorsignal`)  

**Purpose:**  
Establish all state managers with correct dependency order.

---

### Step 6 — Wire Keeper Dependencies

Connect keepers and hooks, for example:

- `staking.SetHooks(...)`  
- `evm.SetStakingKeeper(...)` (or equivalent Ethermint wiring)  
- `poss.SetStakingKeeper(...)`  
- `poss.SetBankKeeper(...)`  
- PoSS keeper gets its params subspace via `ParamsKeeper`  

**Purpose:**  
Ensure proper inter-module communication and PoSS access to balances, staking, and parameters.

---

### Step 7 — Configure ModuleManager

Create and configure the `ModuleManager`:

- register all modules (auth, bank, staking, gov, evm, feemarket, noorsignal)  
- set **BeginBlock** order  
- set **EndBlock** order  
- set **InitGenesis** order  
- set **ExportGenesis** order  

**Purpose:**  
Guarantee deterministic and predictable execution.

---

### Step 8 — Register Services

Register:

- **Msg services** (transaction handlers)  
- **Query services** (gRPC queries)  

for:

- `auth`  
- `bank`  
- `staking`  
- `gov`  
- `evm`  
- `feemarket`  
- `noorsignal`  

**Purpose:**  
Expose the chain functionality over RPC / gRPC.

---

### Step 9 — Register BeginBlocker

Define **BeginBlocker** order, for example:

1. `feemarket`  
2. `evm`  
3. `staking`  
4. `noorsignal` (PoSS)  
5. `gov`  

**Purpose:**  

- update base fee (feemarket)  
- prepare EVM block context  
- apply staking updates  
- process PoSS signals & rewards (when enabled)  
- update governance state  

---

### Step 10 — Register EndBlocker

Define **EndBlocker** order, for example:

1. `staking`  
2. `gov`  

**Purpose:**  

- finalize validator set updates  
- finalize governance calculations & proposal results  

---

### Step 11 — Register InitGenesis / ExportGenesis

Define **InitGenesis** order (example):

1. `auth`  
2. `bank`  
3. `staking`  
4. `gov`  
5. `evm`  
6. `feemarket`  
7. `noorsignal`  

Each module:

- reads its own genesis state  
- validates it  
- initializes internal store state  

**Purpose:**  
Ensure deterministic, reproducible genesis state creation.

---

### Step 12 — Load Latest Version

`BaseApp` loads the **latest application state**:

- from disk (if restarting an existing chain)  
- from genesis (if `height == 0`)  

**Purpose:**  
Prepare the node to run from the correct state.

---

### Step 13 — Expose App Structure

The constructor returns the final `App` struct containing:

- `BaseApp`  
- keepers  
- codecs and interface registry  
- store keys  
- ModuleManager  
- router / service infrastructure  

**Purpose:**  
Provide a complete, ready-to-run NOORCHAIN application instance.

---

## 4. Phase 2 — Runtime Initialization

After the application has been constructed:

1. The node starts via `noord start`.  
2. If chain height is `0`, `InitChain` is called:  
   - `InitGenesis` is executed in the order defined above.  
   - initial validator set is created.  
3. The chain moves to **block height 1**.  

Then the standard block lifecycle starts and repeats:

- **BeginBlock**  
- **DeliverTx**  
- **EndBlock**  
- **Commit**  

PoSS integration (when enabled) happens mainly in:

- **BeginBlock** (processing signals & rewards)  
- transaction handlers (PoSS messages in `DeliverTx`).  

---

## 5. Phase 3 — Genesis Initialization (Detailed)

This section describes the roles of the main modules at genesis time.

---

### 5.1 Auth Genesis

- Create base accounts (e.g. 5 genesis wallets for: Foundation, Founder, Stimulus, Pre-sale, PoSS Reserve).  
- Assign account numbers and sequences.  

---

### 5.2 Bank Genesis

- Assign initial balances according to the **5 / 5 / 5 / 5 / 80** economic model (in `unur`).  
- Validate that total supply is consistent.  

---

### 5.3 Staking Genesis

- Set staking parameters.  
- Optionally configure initial validator(s) or leave for later testnet configuration.  

---

### 5.4 Gov Genesis

- Set governance parameters:  
  - voting period  
  - deposit requirements  
  - quorum, threshold, veto rules  

---

### 5.5 EVM Genesis

- Set EVM chain ID.  
- Configure base EVM parameters (gas, block config, etc.).  
- Initialize EVM state (if any pre-deployed contracts).  

---

### 5.6 FeeMarket Genesis

- Set initial parameters for the fee market module.  
- Initial base fee is generally `0` at genesis.  

---

### 5.7 PoSS Genesis (`x/noorsignal`)

- Initialize PoSS state store.  
- Set `PoSSEnabled = false` at genesis (PoSS disabled by default).  
- Set initial PoSS parameters:  
  - daily limits  
  - base reward  
  - weights  
  - halving settings  
- Ensure `HalvingEpoch = 0` at chain start.  
- Set counters and totals to zero (`TotalSignals = 0`, `TotalMinted = 0`).  

---

## 6. Initialization Timeline Diagram

```text
┌────────────────────────────────┐
│      1. Encoding Config        │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│          2. BaseApp            │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│       3. Store Keys            │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│      4. Mount Stores           │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│        5. Keepers              │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│     6. Keeper Wiring           │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│     7. ModuleManager           │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│     8. Register Services       │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│ 9./10. Begin/End Block Hooks   │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│11. Init/Export Genesis Orders  │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│    12. Load Latest Version     │
└───────────────┬────────────────┘
                ▼
┌────────────────────────────────┐
│   13. App Ready to Start       │
└────────────────────────────────┘
7. Summary
This App Initialization Flow is now the canonical reference for:

Phase 4C (implementation in app/app.go)

Phase 5 (legal & governance validation)

Phase 7 (Mainnet 1.0 preparation)

It guarantees:

deterministic chain startup

correct module dependency wiring

stable EVM operation

safe PoSS integration (initially disabled)

a clean, auditable initialization path for auditors and partners

Nothing should be added or removed without updating this document
and validating changes against Phase 3 specifications.
