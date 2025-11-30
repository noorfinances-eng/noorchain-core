# NOORCHAIN — x/noorsignal (PoSS Module)

This folder contains the **PoSS module skeleton** for NOORCHAIN.  
It is introduced during **Phase 4 — Implementation (PoSS Block 1 → Block N)**.

At this stage, the module is intentionally minimal.  
It compiles, it mounts a KV store, it registers a basic AppModule, and it
is ready for future upgrades.

---

## 📌 What is PoSS?

**PoSS = Proof of Signal Social**  
A consensus-level mechanism where participants and curators emit verifiable
“signals” (micro-actions, validated interactions, certified contributions)
rewarded by rules defined in the chain.

The final version (implemented later) will include:

- 70/30 reward split (Participant / Curator)
- Daily limits and anti-abuse
- Halving every 8 years
- Fixed supply: **299,792,458 NUR**
- Validation through PoSS Reserve
- Genesis allocation: 5/5/5/5/80 model
- EVM compatibility (CosmosSDK + Ethermint)

**None of these mechanisms are implemented yet.**  
Right now this module is only a technical placeholder.

---

## 📁 Current State (Phase 4 — PoSS Skeleton)

### Implemented

- `types/`  
  - `keys.go`, `types.go`  
  - module name, store keys, proto basics  
- `keeper/`  
  - minimal keeper using `BinaryCodec`  
  - KV store access  
  - module logger  
- `module.go`  
  - AppModuleBasic  
  - AppModule  
  - default genesis (empty)  
  - empty BeginBlock / EndBlock  
- module registered inside `app/app.go`  
  - KV store mounted  
  - AppModule added to `module.Manager`  
  - Module included in `InitGenesis`, `BeginBlock`, `EndBlock`

### Not implemented yet (future PoSS phases)

- Msg server (emit signal, validate signal, curator logic)
- Query server (daily totals, curator ranking, pending validations)
- State transitions and PoSS rules
- Genesis allocation logic
- Daily counters / anti-abuse
- Events
- Integration with BankKeeper and StakingKeeper
- Full params set
- PoSS Reserve handling
- Block-level reward distribution

---

## 🔧 Why is the module empty now?

Because PoSS is a **consensus-level mechanism**, it must be added in a
controlled, incremental way:

1. Make module compile → **done**
2. Register module → **done**
3. Add keeper structure → **done**
4. Add state machine → *later*
5. Add tx messages → *later*
6. Add reward logic → *later*
7. Add gRPC queries → *later*
8. Integrate with genesis → *later*

NOORCHAIN Phase 4 explicitly requires that we finish:

- Cosmos core
- Ethermint integration
- AnteHandler (EVM + Cosmos)
- Params keeper
- FeeMarket keeper
- Testnet readiness  
**before** coding the PoSS business logic.

---

## 🚀 Next Steps (Future PoSS Blocks)

- Add GenesisState definition (`proto`)  
- Add Msg types (`MsgEmitSignal`, `MsgValidateSignal`, etc.)
- Add Query types (`SignalsOfDay`, `CuratorStats`, etc.)
- Add keeper methods (state writes, reward logic)
- Implement PoSS Reserve + 5/5/5/5/80 distribution
- Full validation logic inside BeginBlock / EndBlock

---

## ℹ️ Notes

- This README is part of **NOORCHAIN Phase 4 documentation**.
- It must always stay **skeleton-only** until PoSS Phase 5–6.
- Do not add tokenomics/predictions/speculation: follow **Legal Light CH**.

