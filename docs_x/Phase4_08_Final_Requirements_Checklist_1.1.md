# NOORCHAIN — Phase 4A  
## Final Requirements Checklist (Cosmos + Ethermint + PoSS-Ready)  
### Version 1.1 — Non-code Specification  
### Last Updated: 2025-12-03  

---

# 1. Purpose of This Final Checklist

This document summarizes all **mandatory technical requirements** that MUST be satisfied before starting:

- Phase 4B (PoSS Blueprint)
- Phase 4C (Testnet 1.0 coding & genesis)
- full implementation in `app/app.go`, `x/noorsignal`, and Testnet configuration files

It is the **official Phase 4A completion gate**.

---

# 2. Version Requirements (Mandatory & Verified)

NOORCHAIN uses **fixed and validated versions**:

| Component      | Required Version | Status |
|----------------|------------------|--------|
| Cosmos SDK     | **0.46.11**      | ✔️ |
| Ethermint       | **0.22.0**       | ✔️ |
| CometBFT/TM     | **0.34.27** (via replace) | ✔️ |
| IAVL            | **0.19–0.21** compatible | ✔️ |
| Go              | **1.22+**       | ✔️ |
| CosmJS          | **0.33+**       | ✔️ |

**Important:**  
NOORCHAIN deliberately stays on SDK 0.46 for stability and mainnet maturity.

All versions remain **locked** unless Phase 3 documentation is amended.

---

# 3. Architecture Requirements

### ✔️ Application Layers (All Present)

- Cosmos base modules  
- Ethermint EVM layer  
- NOORCHAIN custom PoSS module  
- Strict separation of responsibilities  

### ✔️ App Components Verified

- BaseApp  
- Encoding system  
- Multistore  
- Keepers  
- ParamsKeeper + Subspaces  
- ModuleManager  
- BeginBlock / EndBlock handlers  

---

# 4. Module Requirements

## 4.1 Cosmos Modules

✔️ `auth`  
✔️ `bank`  
✔️ `staking`  
✔️ `gov`  

All registered, stores mounted, and storages defined.

## 4.2 Ethermint Modules

✔️ `evm`  
✔️ `feemarket`  

Require:

- correct keeper wiring  
- correct EVM genesis parameters  
- EIP-1559 base fee logic active  

## 4.3 Custom Module (PoSS)

Before Phase 4C:

- Keeper defined  
- Store keys created  
- Params included  
- BeginBlock slot reserved  
- No reward logic active (PoSSEnabled = false)  

---

# 5. Keepers Requirements

### ✔️ Keepers Identified

- AccountKeeper  
- BankKeeper  
- StakingKeeper  
- GovKeeper  
- EVMKeeper  
- FeeMarketKeeper  
- PoSSKeeper  
- ParamsKeeper  

### ✔️ Dependency Graph Mapped

No circular dependencies.

### ✔️ Strict Instantiation Order Validated

1. AccountKeeper  
2. BankKeeper  
3. StakingKeeper  
4. GovKeeper  
5. ParamsKeeper (subspaces)  
6. EVMKeeper  
7. FeeMarketKeeper  
8. PoSSKeeper  

### ✔️ Wiring Rules

- `staking.SetHooks(...)`  
- `evm.SetStakingKeeper(...)` or Ethermint equivalent  
- `poss.SetBankKeeper(...)`  
- `poss.SetStakingKeeper(...)`  
- PoSS receives its own params subspace  

---

# 6. Lifecycle Requirements

### **BeginBlock Order**
feemarket → evm → staking → noorsignal → gov


### **DeliverTx Order**


ante → routing → state writes → EVM execution


### **EndBlock Order**


staking → gov


### **Genesis Order**


auth → bank → staking → gov → evm → feemarket → noorsignal


### **ExportGenesis Order**
Matches InitGenesis exactly.

All ordering constraints have been validated.

---

# 7. Store & State Model Requirements

✔️ All store keys fixed and mounted  
✔️ IAVL subtree layout validated  
✔️ PoSS store defined (signals, counters, reward state)  
✔️ anti-abuse state + epoch state included  
✔️ EVM stateDB integrated correctly  
✔️ ParamsKeeper subspaces created for all modules  

---

# 8. App Initialization Requirements

The app constructor **must** include:

- encoding config  
- BaseApp creation  
- KVStore + TransientStore mounting  
- all keeper instantiation (strict order)  
- dependency wiring  
- ModuleManager creation  
- service registration  
- BeginBlock & EndBlock mapping  
- InitGenesis & ExportGenesis wiring  
- state loading  
- returning final App struct  

All steps validated in the App Initialization Blueprint.

---

# 9. External API Requirements

### ✔️ gRPC routes

- cosmos auth / bank / staking / gov  
- ethermint evm / feemarket  
- noorsignal (Phase 4B+)  

### ✔️ JSON-RPC routes (Ethermint)

- `eth_*`  
- `web3_*`  
- `net_*`  
- `debug_*` (optional)  

RPC must run without panic on empty chain.

---

# 10. Pre-Testnet Mandatory Conditions

Before Phase 4C coding begins:

### ✔️ Compilation Check


go build ./...

must succeed.

No missing references, no nil keepers, no unmounted stores.

### ✔️ Boot Check


noord start

must start successfully.

Node must:

- start BaseApp  
- expose RPC & gRPC  
- open WebSocket interface  
- produce empty-block sequence (no panic)

### ✔️ Module Checks
- ModuleManager ordering validated  
- All keepers correctly injected  
- All stores correctly mounted  
- ParamsKeeper functional  

### ✔️ PoSS Ready (but disabled)
- PoSSKeeper initialized  
- PoSS KVStore mounted  
- PoSS params defaulted  
- BeginBlock slot present  
- PoSSEnabled = false  

---

# 11. Completion Statement

Phase 4A is considered **100% complete** when:

- All 8 blueprint documents exist  
- All requirements in this checklist are validated  
- The application architecture is stable  
- No missing keeper, store, module, or lifecycle component remains  

At this moment, NOORCHAIN is officially ready to enter:

## → Phase 4B (PoSS Blueprint)  
## → Phase 4C (Testnet 1.0 Development)

---

# 12. Summary Table

| Category        | Status |
|-----------------|--------|
| Versions        | ✔️ |
| Architecture    | ✔️ |
| ModuleManager   | ✔️ |
| Keepers         | ✔️ |
| Lifecycle       | ✔️ |
| Stores          | ✔️ |
| App Init        | ✔️ |
| PoSS Ready      | ✔️ |
| Testnet Prereqs | ✔️ |
| Phase 4A        | **100% Complete** |
