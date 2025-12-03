# NOORCHAIN — Phase 4A  
## Store & State Model Blueprint  
### Version 1.1 — Architecture Only (No Code)  
### Last Updated: 2025-12-03  

---

# 1. Purpose of This Document

This document defines the **complete, authoritative state model** of  
NOORCHAIN 1.0. It describes:

- store key layout  
- multistore architecture  
- module-specific state trees  
- IAVL structure  
- EVM state DB behaviour  
- PoSS state model  
- inter-store dependencies  
- determinism and persistence rules  

This is the **final reference** for implementing stores during Phase 4C,  
for building Testnet 1.0, and for preparing Mainnet.

---

# 2. Multistore Overview (IAVL)

NOORCHAIN’s state is managed by the Cosmos SDK **Multistore**, composed of:

- **IAVL stores** (persistent key/value state)  
- **Transient stores** (ephemeral during block execution)  
- **Memory stores** (in-memory for parameters and caches)  

The Multistore root holds all module KVStores under separate store keys,  
ensuring deterministic Merkle proofs and consensus safety.

---

# 3. Store Key Map

| Store Key   | Module         | Keeper           |
|-------------|----------------|------------------|
| `auth`      | auth           | AccountKeeper    |
| `bank`      | bank           | BankKeeper       |
| `staking`   | staking        | StakingKeeper    |
| `gov`       | gov            | GovKeeper        |
| `evm`       | evm            | EVMKeeper        |
| `feemarket` | feemarket      | FeeMarketKeeper  |
| `noorsignal`| x/noorsignal   | PoSSKeeper       |
| `params`    | params         | ParamsKeeper     |

All are mounted as **IAVL-backed KVStores**, except:

- params: memory + kv + transient components  
- evm: uses its own EVM stateDB internally  

---

# 4. Module State Models (Detailed)

This section describes the complete internal state structure of each module.

---

## 4.1 auth Module State

Stores:

- base accounts  
- sequences  
- account numbers  
- public keys  

Used for:

- signature verification  
- nonce/sequence checks  
- Ethereum tx sender validation  

---

## 4.2 bank Module State

Stores:

- balances (`address → coins`)  
- supply (per denomination)  
- denom metadata (`unur` main token)  

Used by:

- staking (bonding / unbonding)  
- evm (gas deduction)  
- poss (reward distribution from PoSS reserve)  

---

## 4.3 staking Module State

Stores:

- validators  
- delegations  
- unbonding delegations  
- redelegations  
- validator power  
- staking parameters  
- slashing metadata  

Used by:

- consensus (validator set)  
- governance (voting power)  
- PoSS (validator power relevance for future curation logic)  

---

## 4.4 gov Module State

Stores:

- proposals  
- deposits  
- votes  
- tally results  
- governance parameters  

Governance depends on staking for voting power.

PoSS never writes to or modifies governance state.

---

## 4.5 evm Module State

This module uses its own **EVM stateDB**, embedded in the `evm` store.

Stores:

- contract code  
- contract storage  
- nonces  
- logs  
- receipts  
- bloom filters  
- EVM block metadata  

This state is large and append-heavy.

Only the EVM module writes here.

---

## 4.6 feemarket Module State

Stores:

- base fee (EIP-1559)  
- block gas usage metrics  
- fee market parameters  

Used by:

- evm (block pricing logic)  
- antehandler (gas validation)  

---

## 4.7 noorsignal (PoSS) Module State

The PoSS store is a custom state tree containing:

### 1. Signal Registry

Key: `signal_id` →  
{
sender_address,
curator_address,
signal_type,
timestamp,
weight,
block_height
}

### 2. Reward State

Stores:

- current reward epoch  
- halving index  
- last reward block  
- reward multipliers  
- PoSS base reward  

### 3. Anti-Abuse State

Stores:

- daily counters (per participant)  
- daily curator validations  
- rate-limit data  
- last signal timestamps  

### 4. PoSS Parameters

- daily max signals  
- curator max validations  
- weights per signal  
- halving schedule  
- PoSSEnabled flag  
- immutable 70/30 split  

### 5. Statistics

- total signals  
- total minted from PoSS reserve  
- per-curator stats  

---

# 5. Inter-Store Interactions

The following interactions are consensus-critical:

---

## PoSS ↔ staking

- validator power needed for reward governance (future PoSS v2)  
- block height & timestamps for halving  

---

## PoSS ↔ bank

- reward distribution (when PoSS enabled)  
- PoSS Reserve transfers  
- balance checks  

---

## EVM ↔ bank

- gas deduction  
- payer balance verification  

---

## staking ↔ gov

- governance voting power = validator power  
- proposal finalization depends on staking state  

---

## bank ↔ auth

- balances tied to account objects  

---

# 6. State Consistency Requirements

These rules guarantee deterministic and safe execution:

---

### 6.1 Deterministic Writes

- BeginBlock and EndBlock must write in a fixed order.  
- Only staking and governance may write in EndBlock.  
- PoSS must *never* mutate state in EndBlock.  

---

### 6.2 DeliverTx Restrictions

- PoSS **must not** write heavy state or perform reward distribution in DeliverTx.  
- EVM tx must use atomic commit / revert.  

---

### 6.3 No Circular Dependencies

- EVM must not depend on PoSS  
- PoSS must not depend on EVM  
- staking must not call PoSS during validator updates  

---

### 6.4 PoSS Disabled by Default

At genesis:

PoSSEnabled = false

---

# 7. State Persistence Guarantees

### On every commit:

1. **IAVL Multistore commit**  
   - deterministic Merkle root  
   - sent to CometBFT as next block app hash  

2. **EVM state commit**  
   - flush contract state  
   - store logs  
   - compute bloom filter  

3. **PoSS state commit**  
   - commit counters  
   - commit halving indices  
   - commit reward metadata  

**Commit ordering:**

IAVL commit → EVM commit → PoSS commit
This ensures:

- EVM state is consistent  
- PoSS indexing is consistent  
- no mismatches between app hash and EVM receipts  

---

# 8. Store Architecture Diagram

root (IAVL)
│
├── auth
├── bank
├── staking
├── gov
├── evm
├── feemarket
└── noorsignal

yaml
Copier le code

Each module defines an isolated subtree under the Merkle root.

---

# 9. Summary Table

| Module      | Stores                       | Critical For |
|-------------|------------------------------|--------------|
| auth        | accounts                     | signatures, sequences |
| bank        | balances, supply             | staking, evm, poss |
| staking     | validators, delegations      | consensus, gov, poss |
| gov         | proposals, votes             | governance |
| evm         | stateDB, logs, receipts      | EVM compatibility |
| feemarket   | base fee, gas metrics        | gas pricing |
| noorsignal  | signals, rewards, anti-abuse | PoSS engine |

---

# 10. Final Notes

This blueprint guarantees:

- deterministic state evolution  
- clean integration of EVM + PoSS  
- stable data model for Testnet 1.0  
- correct PoSS reward accounting  
- predictable Mainnet behaviour  

Any change to this architecture MUST be reflected in the Phase 3 documentation and revalidated in audit.
