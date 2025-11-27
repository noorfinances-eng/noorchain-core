**NOORCHAIN — Phase 4A

App Initialization Flow (Cosmos SDK + Ethermint + PoSS)**
Version 1.1 — Architecture Only (No Code)

🔧 1. Purpose of This Document

This document defines the complete initialization flow of the NOORCHAIN application.

It describes:

how the blockchain app is constructed

in which order components must be instantiated

how modules hook into the lifecycle

how the app transitions from “empty” to “running chain”

how PoSS fits into the init process

This is the reference for building the real constructor in app/app.go during Phase 4C.

🏗️ 2. Initialization Flow Overview

Initialization occurs in three phases:

Application Construction

ModuleManager Wiring

Genesis Initialization

Each phase must be executed in the correct order to avoid unstable state, missing stores, or EVM panic.

🧱 3. Phase 1 — Application Construction

The app constructor must perform the following steps:

Step 1 — Build Encoding Config

Construct:

Interface Registry

Amino (legacy) codec

Proto codec

TxConfig

JSON marshaler

Purpose:
→ Enable message processing, signing, decoding, transaction routing.

Step 2 — Create BaseApp

Build BaseApp with:

app name

logger

database

TxDecoder / TxEncoder

interface registry

Purpose:
→ Core ABCI engine.

Step 3 — Define Store Keys

For:

auth

bank

staking

gov

evm

feemarket

noorsignal

Purpose:
→ Prepare KVStores to mount persistent chain state.

Step 4 — Mount Stores

Using MountKVStore and others.

Purpose:
→ State persistence.

Step 5 — Create Keepers

In strict order:

1. AccountKeeper
2. BankKeeper
3. StakingKeeper
4. GovKeeper
5. EVMKeeper
6. FeeMarketKeeper
7. PoSSKeeper


Purpose:
→ Establish all state managers.

Step 6 — Wire Keeper Dependencies

Examples:

staking.SetHooks()

evm.SetStakingKeeper()

poss.SetStakingKeeper()

poss.SetBankKeeper()

Purpose:
→ Proper module interconnection.

Step 7 — Configure ModuleManager

Provides:

begin-block order

end-block order

InitGenesis order

ExportGenesis order

Purpose:
→ Deterministic execution.

Step 8 — Register Services

Message services and query services for:

auth

bank

staking

gov

evm

feemarket

noorsignal

Purpose:
→ Expose RPC & gRPC functionality.

Step 9 — Register BeginBlocker

Order:

feemarket → evm → staking → noorsignal → gov


Purpose:
→ Correct block lifecycle execution.

Step 10 — Register EndBlocker

Order:

staking → gov


Purpose:
→ Finalize validator set & governance.

Step 11 — Register InitGenesis / ExportGenesis

Modules must initialize in this order:

auth → bank → staking → gov → evm → feemarket → noorsignal


Purpose:
→ Deterministic genesis.

Step 12 — Load Latest Version

BaseApp loads the application state from disk (or genesis).

Purpose:
→ Prepare the node to run.

Step 13 — Expose App Structure

Return the final App struct with:

keepers

codec

BaseApp

stores

Mmgr reference

interface registry

router

Purpose:
→ Complete application ready to start.

🔌 4. Phase 2 — Runtime Initialization

After construction:

Node starts via noord start

BaseApp calls InitChain if height 0

Modules execute InitGenesis

Validator set committed

Chain moves to block height 1

Then block lifecycle begins:

BeginBlock → DeliverTx → EndBlock → Commit

🌍 5. Phase 3 — Genesis Initialization (More Detail)
Auth Genesis

Create base accounts (5 genesis wallets)

Apply account numbers

Bank Genesis

Assign initial balances (5/5/5/5/80)

Validate supply

Staking Genesis

Set params

Set initial validator (or delegations later in Testnet 1.0)

Gov Genesis

Set voting params

Set deposit params

EVM Genesis

Set chain ID

Configure base EVM parameters

Deploy fee market params

FeeMarket Genesis

Set initial base fee (usually 0)

PoSS Genesis

Initialize PoSS state store

Set halving epoch 0

Set initial reward indexes

🚦 6. Initialization Timeline Diagram
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

🎯 7. Summary

This initialization flow is now the canonical reference for Phase 4C implementation.

It ensures:

deterministic chain startup

correct module dependency wiring

stable EVM operation

stable PoSS integration

correct block lifecycle

Nothing should be added or removed unless explicitly validated in Phase 3 specifications.