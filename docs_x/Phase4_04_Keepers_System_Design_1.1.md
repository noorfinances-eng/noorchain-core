# NOORCHAIN — Phase 4A  
## Keepers System Design (Cosmos SDK + Ethermint + PoSS)  
### Version 1.1 — Architecture Only (No Code)  
### Last Updated: 2025-12-03  

---

# 1. Purpose of This Document

This document defines the complete **Keeper System Architecture** of  
NOORCHAIN 1.0 as validated at the end of **Phase 4**.

It specifies:

- all required keepers  
- their responsibilities  
- their dependencies  
- correct instantiation order  
- inter-keeper interactions  
- store mappings  
- connections to the ModuleManager  
- integration of the PoSS Keeper  
- requirements for Testnet and Mainnet  

This is the **canonical reference** for Phase 4B (PoSS Logic), Phase 4C  
(Testnet 1.0 coding), Phase 5 (Governance), and Phase 7 (Mainnet).

---

# 2. What Is a Keeper?

In the Cosmos SDK:

**A Keeper is the state manager of a module.**

It is responsible for:

- reading/writing KVStore  
- validating state transitions  
- executing module logic  
- calling other keepers  
- exposing MsgServer & QueryServer services  

A keeper is **consensus-critical**: wrong state logic → chain halt.

---

# 3. Required Keepers in NOORCHAIN

NOORCHAIN uses three keeper categories:

---

## 3.1 Cosmos SDK Keepers (Base Layer)

### **AccountKeeper**
- accounts, sequences, public keys  
- foundation for all other modules  

### **BankKeeper**
- balances, transfers, supply  
- denomination `unur` (main NOORCHAIN token)  

### **StakingKeeper**
- validators, delegations, power updates  
- communicates with consensus engine  

### **GovKeeper**
- proposals, votes, governance parameters  
- relies on staking voting power  

---

## 3.2 Ethermint Keepers (EVM Layer)

### **EVMKeeper**
- EVM execution  
- gas accounting  
- stateDB compatibility with Ethereum tools  

### **FeeMarketKeeper**
- EIP-1559-style dynamic base fee  
- depends entirely on EVMKeeper  

---

## 3.3 NOORCHAIN Custom Keeper (PoSS Layer)

### **PoSSKeeper** (`x/noorsignal`)
Handles:

- PoSS signal ingestion  
- reward computation  
- halving logic (every 8 years)  
- 70/30 distribution  
- anti-abuse counters  
- PoSS parameters (via ParamsKeeper)  
- event emission  
- block hooks (BeginBlock / EndBlock)  
- interaction with bank & staking  

PoSS does **not** depend on EVM.

---

# 4. Keeper Dependencies (Final & Verified)

## 4.1 Cosmos Keepers

AccountKeeper → BankKeeper → StakingKeeper → GovKeeper

markdown
Copier le code

- AccountKeeper has no dependencies (root keeper)  
- BankKeeper needs AccountKeeper  
- StakingKeeper needs AccountKeeper + BankKeeper  
- GovKeeper needs StakingKeeper (voting power)  

---

## 4.2 Ethermint Keepers

### **EVMKeeper**
Depends on:

- AccountKeeper  
- BankKeeper  
- StakingKeeper  

Reason:
- balances needed for gas charges  
- signature/nonce validation relies on Account  
- validator power influences EVM block-context  

### **FeeMarketKeeper**
Depends on:

- EVMKeeper  

---

## 4.3 PoSS Keeper

### **PoSSKeeper**
Depends on:

- AccountKeeper  
- BankKeeper  
- StakingKeeper  
- ParamsKeeper (PoSS subspace)  

Also requires:

- BeginBlock hook  
- chain timestamp  
- block height  
- access to validator power (PoSS v2)  

It does **not** depend on EVM modules.

---

# 5. Keeper Instantiation Order (Strict)

This is **critical**.  
A wrong order = panic at startup.

NOORCHAIN official order:

1. **AccountKeeper**  
2. **BankKeeper**  
3. **StakingKeeper**  
4. **GovKeeper**  
5. **ParamsKeeper + all subspaces**  
6. **EVMKeeper**  
7. **FeeMarketKeeper**  
8. **PoSSKeeper**

This order ensures:

- bank has accounts  
- staking has accounts + bank  
- governance has staking  
- EVM has account/bank/staking  
- feemarket has EVM  
- PoSS has bank + staking  

---

# 6. Keeper Interaction Graph (Final)

markdown
Copier le code
      ┌───────────────┐
      │ AccountKeeper │
      └───────┬───────┘
              │
      ┌───────▼───────┐
      │  BankKeeper   │
      └───────┬───────┘
              │
      ┌───────▼────────┐
      │ StakingKeeper  │
      └───────┬────────┘
      ┌───────▼────────┐
      │   GovKeeper    │
      └─────────────────┘

      ┌───────────────────────────────┐
      │            EVMKeeper          │
      └───────────────────────────────┘
             ▲     ▲        ▲
             │     │        │
   AccountKeeper BankKeeper StakingKeeper

      ┌────────────────────┐
      │ FeeMarketKeeper   │
      └───────────▲────────┘
                  │
              EVMKeeper

      ┌────────────────────┐
      │    PoSSKeeper     │
      └────────────────────┘
           ▲    ▲    ▲    ▲
AccountKeeper BankKeeper StakingKeeper ParamsKeeper

yaml
Copier le code

---

# 7. Store Keys for Each Keeper

| Keeper           | Store Key   |
|------------------|-------------|
| AccountKeeper    | auth        |
| BankKeeper       | bank        |
| StakingKeeper    | staking     |
| GovKeeper        | gov         |
| EVMKeeper        | evm         |
| FeeMarketKeeper  | feemarket   |
| PoSSKeeper       | noorsignal  |
| ParamsKeeper     | params      |

All stores must be mounted in `app.go`.

---

# 8. Keeper → ModuleManager Connections

Each keeper must expose:

- **MsgServer** (transaction messages)  
- **QueryServer** (gRPC queries)  
- **InitGenesis** + **ExportGenesis**  
- **BeginBlocker** and **EndBlocker** (when required)  
- **Invariants** (for bank, staking, and PoSS optional)  

The ModuleManager will then:

- set BeginBlock order  
- set EndBlock order  
- register services  
- load genesis state  

PoSS module connects via:

- BeginBlock hooks  
- parameters (ParamsKeeper subspace)  

---

# 9. App Constructor Integration (Blueprint)

`NewNoorchainApp()` must:

1. create **KVStoreKeys** & **TransientStoreKeys**  
2. instantiate keepers in correct order  
3. wire dependencies  
4. attach hooks  
5. initialize ModuleManager  
6. register services  
7. set `InitGenesis`, `ExportGenesis`  
8. set `BeginBlocker` and `EndBlocker`  
9. load state  

This ensures deterministic, mainnet-safe initialization.

---

# 10. Pre-Test Checklist (Before Phase 4C)

All below MUST be true before Testnet 1.0:

- all keepers defined  
- all store keys mounted  
- all dependencies mapped  
- instantiation order validated  
- PoSS subspace created  
- no circular dependencies  
- BeginBlock hooks connected  
- PoSSKeeper reading correct params  

---

# 11. Summary Table

| Keeper     | Depends On            | Store       | Purpose |
|------------|------------------------|-------------|---------|
| Account    | none                   | auth        | accounts, signatures |
| Bank       | account                | bank        | balances, transfers |
| Staking    | account, bank          | staking     | validators, delegation |
| Gov        | staking                | gov         | proposals, voting |
| EVM        | account, bank, staking | evm         | EVM execution |
| FeeMarket  | evm                    | feemarket   | base fee (EIP-1559) |
| PoSS       | account, bank, staking, params | noorsignal | social consensus |

---

# 12. Final Statement

This document is considered **final** and is now the reference for all  
future development, governance, auditing, and mainnet preparation work.
