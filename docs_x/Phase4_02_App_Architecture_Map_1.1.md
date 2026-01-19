# NOORCHAIN — Phase 4A  
## App Architecture Map (Cosmos SDK + Ethermint + PoSS)  
### Version 1.1 — Final Structural Overview  
### Last Updated: 2025-12-03  

This document defines the **final, authoritative architecture map** of the  
NOORCHAIN Core Application as implemented at the end of **Phase 4**.

It is fully aligned with:

- Cosmos SDK v0.46.11  
- Ethermint v0.22.0  
- CometBFT/Tendermint v0.34.27 (with replace directive)  
- NOORCHAIN custom module `x/noorsignal` (PoSS)  
- The final Keeper structure validated in Phase 4  
- The Testnet 1.0 application architecture  

No code is included.

---

# 1. Purpose of This Document

This file acts as the **structural reference** for:

- Phase 4B (PoSS Logic Integration)  
- Phase 4C (Testnet 1.0 assembly)  
- Phase 5 (Governance & Legal Architecture)  
- Phase 6 (Genesis Pack, Website)  
- Phase 7 (Mainnet Preparation)  

It defines:

- all application layers  
- keeper relationships  
- store layout  
- module dependencies  
- block lifecycle  
- genesis lifecycle  
- complete app constructor structure  

---

# 2. Application Layer Overview

NOORCHAIN is composed of **three major layers**:

---

## 2.1 Cosmos SDK Layer (Base Layer)

This layer provides the essential blockchain primitives.

### Modules included:
- **auth** — accounts, signatures, sequence numbers  
- **bank** — balances, transfers, supply  
- **staking** — validators, delegations, power updates  
- **gov** — governance proposals, votes, tallying  
- **params** — parameter subspaces (implicit in SDK v0.46.x)

### Purpose:
> Maintain chain security, state integrity, validator set, governance, and fundamental on-chain logic.

---

## 2.2 Ethermint Layer (EVM Layer)

Provides full Ethereum EVM compatibility on top of Cosmos.

### Modules included:
- **evm** — EVM execution, stateDB, gas rules  
- **feemarket** — dynamic base fee (EIP-1559 style)

### Purpose:
> Allow EVM smart contracts, Ethereum tooling, RPC compatibility, and wallets (e.g. MetaMask).

---

## 2.3 Custom Layer — NOORCHAIN PoSS (x/noorsignal)

A custom module implementing the Proof of Signal Social system.

### Responsibilities:
- signal ingestion  
- reward computation  
- 70/30 split  
- halving logic  
- counters and anti-abuse  
- event emission  
- PoSS parameters management  
- integration with BeginBlock / EndBlock  

### Purpose:
> Provide NOORCHAIN’s unique **Social Consensus + Reward Layer**.

---

# 3. App Composition Structure

The NOORCHAIN application is composed of five core components:

---

## 3.1 BaseApp

Responsible for:

- ABCI interface  
- transaction execution  
- message routing  
- block lifecycle  
- mempool logic  
- state commitments  

---

## 3.2 Encoding System

The app uses:

- Interface Registry (for type URLs)  
- Protobuf Codec (primary codec)  
- Amino (legacy support for signatures when required)  
- Tx signing modes  
- gRPC types  

---

## 3.3 Store System

Each module receives its own:

- **KVStore**  
- **Transient store** (if required)  
- **Memory store** (mostly params)

### Store keys mounted in BaseApp:
auth
bank
staking
gov
evm
feemarket
noorsignal
params

---

## 3.4 Keepers

Keepers are responsible for:

- maintaining module state  
- reading/writing to KV stores  
- verifying invariants  
- interacting with other modules  

NOORCHAIN employs the following keepers:

- **AccountKeeper**  
- **BankKeeper**  
- **StakingKeeper**  
- **GovKeeper**  
- **EVMKeeper**  
- **FeeMarketKeeper**  
- **ParamsKeeper** (with subspaces)  
- **PoSSKeeper** (`x/noorsignal`)

---

## 3.5 ModuleManager

The ModuleManager:

- registers all modules  
- configures BeginBlock / EndBlock order  
- defines InitGenesis & ExportGenesis order  
- binds services  
- exposes gRPC and REST endpoints  

Modules included:

- `auth`, `bank`, `staking`, `gov`, `evm`, `feemarket`, `noorsignal`  

---

# 4. Keeper Dependency Graph (Final)

This graph represents how keepers depend on each other.

---

## 4.1 Cosmos Keepers (base layer)
AccountKeeper → BankKeeper → StakingKeeper → GovKeeper
The FeeMarketKeeper requires access to EVM’s internal gas accounting.

---

## 4.3 PoSS Keeper (Custom)

PoSSKeeper → AccountKeeper
→ BankKeeper
→ StakingKeeper
→ ParamsKeeper (PoSS subspace)
→ Hooks (BeginBlock + EndBlock)

yaml
Copier le code

The PoSS module does *not* depend on EVM, but needs:

- bank for reward movement (future activation)  
- staking for validator power (future curation logic)  
- params for PoSS settings  
- hooks to perform daily resets and halving evaluation  

---

# 5. Store Layout Map

A final snapshot of all stores mounted at app init:

RootStore (IAVL Merkle Tree)
│
├── auth (KVStore)
├── bank (KVStore)
├── staking (KVStore)
├── gov (KVStore)
├── evm (KVStore)
├── feemarket (KVStore)
├── noorsignal (KVStore)
└── params (KV + Transient)

markdown
Copier le code

All mounted stores use deterministic prefixes and follow Cosmos SDK conventions.

---

# 6. Block Lifecycle (Final Phase 4 Model)

### 6.1 BeginBlock
Order is carefully chosen:

1. FeeMarket → updates EIP-1559 base fee  
2. EVM → prepares block context  
3. Staking → validator updates (consensus-critical)  
4. PoSS → future: signal counters, halving epoch tracking  
5. Gov → proposal tallies and time updates  

### 6.2 DeliverTx sequence
1. AnteHandler (signature verification, fee deduction)  
2. Msg routing  
3. Module logic  
4. Gas metering  
5. State transitions  

### 6.3 EndBlock
1. Staking validator set updates  
2. Gov proposal status updates  
3. Future: PoSS daily aggregation hooks  

### 6.4 Commit
State root committed (IAVL).  
Next block begins.

---

# 7. Genesis Lifecycle

At genesis initialization:

- accounts created  
- balances applied  
- staking params initialized  
- gov params set  
- EVM chain configuration loaded  
- feemarket params initialized  
- PoSS params loaded (PoSSEnabled = false by default)  
- PoSS counters set to zero  
- initial validator set computed  
- chain starts at height 1  

---

# 8. Inter-Module Interaction Summary

### auth ↔ bank  
- accounts  
- balances  
- signature semantics  

### bank ↔ staking  
- delegation shares  
- locked funds  
- reward pools  

### staking ↔ gov  
- voting power  
- validator-based governance  

### evm ↔ feemarket  
- gas price calculation  
- EIP-1559 base fee  

### noorsignal ↔ bank / staking  
- future PoSS reward distribution  
- signal validation economics  
- validator-power-dependent curation (Phase 5+)  

---

# 9. App Constructor — Final Map (as implemented in Phase 4)

The `NewNoorchainApp()` constructor performs:

1. Build app encoding config  
2. Instantiate BaseApp  
3. Define KVStoreKeys and TransientStoreKeys  
4. Instantiate:
   - AccountKeeper  
   - BankKeeper  
   - StakingKeeper  
   - GovKeeper  
   - ParamsKeeper (with subspaces for each module)  
   - EVMKeeper  
   - FeeMarketKeeper  
   - PoSSKeeper  
5. Configure keeper dependencies  
6. Create ModuleManager  
7. Register:
   - services  
   - `InitGenesis`  
   - `ExportGenesis`  
8. Register `BeginBlocker` and `EndBlocker`  
9. Load latest application state  

The full constructor guarantees deterministic app assembly.

---

# 10. Summary Table (Final)

| Layer        | Component                     | Purpose |
|--------------|-------------------------------|---------|
| Cosmos Base  | auth/bank/staking/gov/params  | Core blockchain logic |
| EVM Layer    | evm/feemarket                 | Ethereum compatibility + gas economy |
| Custom Layer | noorsignal                    | PoSS social consensus engine |
| App System   | BaseApp, Keepers, Stores      | Runtime infrastructure |
| Runtime      | BeginBlock/EndBlock           | Consensus lifecycle |
