**NOORCHAIN — Phase 4A

Final Requirements Checklist (Cosmos + Ethermint + PoSS-Ready)**
Version 1.1 — Non-code specification

🔧 1. Purpose of This Final Checklist

This document summarizes all technical requirements that MUST be satisfied before starting:

Phase 4B (PoSS Blueprint)

Phase 4C (Testnet 1.0 coding & genesis)

full implementation in app/app.go, x/noorsignal, and Testnet files

It serves as the official Phase 4A completion gate.

🧩 2. Version Requirements (Mandatory)
Component	Required Version	Verified
Cosmos SDK	0.50.x	✔️
CometBFT	0.37.x	✔️
Ethermint	0.27.x	✔️
IAVL	0.21+	✔️
Go	1.22+	✔️
CosmJS	0.33+	✔️

All versions remain fixed unless Phase 3 is updated.

🏛️ 3. Architecture Requirements
✔️ Application Layers

Cosmos SDK base modules

Ethermint EVM modules

Custom PoSS module

Clean separation between layers

✔️ App Components

BaseApp

Encoding system

Multistore

Keepers

ModuleManager

BeginBlock/EndBlock hooks

🗂️ 4. Module Requirements
4.1 Cosmos Modules

auth

bank

staking

gov
All must be fully registered and initialized.

4.2 Ethermint Modules

evm

feemarket
Both require correct keeper wiring and genesis handlers.

4.3 Custom Module (PoSS)

Must be defined in Phase 4B

Must support BeginBlock processing

Must support its own KVStore

Must integrate with bank + staking

🧱 5. Keepers Requirements
✔️ Keepers Identified

AccountKeeper

BankKeeper

StakingKeeper

GovKeeper

EVMKeeper

FeeMarketKeeper

PoSSKeeper

✔️ Keeper Dependencies Mapped

No circular dependencies

Strict instantiation order validated

✔️ Keeper Wiring Rules

staking.SetHooks(…)

evm.SetStakingKeeper(…)

poss.SetBankKeeper / poss.SetStakingKeeper

🔄 6. Lifecycle Requirements
BeginBlock Order
feemarket → evm → staking → noorsignal → gov

EndBlock Order
staking → gov

Genesis Order
auth → bank → staking → gov → evm → feemarket → noorsignal

ExportGenesis Order

Matches InitGenesis.

🗄️ 7. Store & State Model Requirements
✔️ Store Keys Fixed

auth

bank

staking

gov

evm

feemarket

noorsignal

✔️ State Models Defined

Each module’s state model described in detail.

✔️ PoSS State Requirements

signals

reward state

anti-abuse counters

halving tracking

weight tables

🛠️ 8. App Initialization Requirements
Must include:

encoding config

BaseApp

store mounting

keeper instantiation

keeper dependency wiring

module manager creation

service registration

block handlers

genesis config

return of final App struct

All steps sequenced and validated.

🌐 9. External API Requirements
✔️ gRPC routes

cosmos auth/bank/staking/gov

ethermint evm/feemarket

noorsignal (later Phase 4B)

✔️ JSON-RPC routes

eth_*

web3_*

debug_* (optional)

Must be exposed automatically by Ethermint.

🚦 10. Pre-Testnet Mandatory Conditions

Before Phase 4C coding:

Compilation

go build ./... must succeed

no missing references

no uninitialized module

Boot

noord start must run (empty chain)

RPC + gRPC must start without panic

Module checks

ModuleManager ordering correct

Keeper dependencies wired

Stores correctly mounted

PoSS-ready

PoSS store + keeper placeholder OK

BeginBlock slot reserved for PoSS logic

🧩 11. Completion Statement

Phase 4A is considered 100% complete when:

All 8 blueprint files exist

All requirements above are validated

The app structure is stable

No missing architectural component remains

Ready to begin Phase 4B (PoSS Blueprint)

🎯 12. Summary Table
Category	Status
Versions	✔️
Architecture	✔️
ModuleManager	✔️
Keepers	✔️
Lifecycle	✔️
Stores	✔️
App Init	✔️
PoSS Requirements	✔️
Testnet Prereqs	✔️
Phase 4A Completed	✔️ 100%