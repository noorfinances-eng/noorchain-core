*NOORCHAIN — Phase 4A

Keepers System Design (Cosmos SDK + Ethermint + PoSS)**
Version 1.1 — Architecture only, no code

🔧 1. Purpose of This Document

This document defines the entire Keeper System of NOORCHAIN 1.0:

all required keepers

their responsibilities

their dependencies

their interactions

their order of instantiation

how PoSS integrates into the keeper layer

how keepers connect to the ModuleManager

This is the official keeper reference for Phase 4B (PoSS blueprint) and the coding of Phase 4C.

🏛️ 2. What Is a Keeper?

In the Cosmos SDK:

A Keeper is the module’s “state manager”.
It directly reads/writes state, applies business logic, and interacts with other modules.

Each keeper is a Go struct responsible for:

accessing KVStore

validating state transitions

performing module logic

interacting with other keepers

exposing methods for Msg / Query services

🧩 3. Required Keepers in NOORCHAIN
3.1 Cosmos SDK Keepers

AccountKeeper

BankKeeper

StakingKeeper

GovKeeper

These ensure the chain can:

authenticate accounts

manage balances

handle validators & delegation

execute governance

3.2 Ethermint Keepers

EVMKeeper

FeeMarketKeeper

These ensure:

full EVM execution

Ethereum transaction compatibility

dynamic base fee model (EIP-1559)

3.3 Custom Keeper (NOORCHAIN)

PoSSKeeper (x/noorsignal)

Handles:

PoSS signals

reward calculations

halving schedule

distribution 70/30

indexing

anti-abuse rules

block hooks (BeginBlock)

Detailed blueprint in Phase 4B.

🔗 4. Keeper Dependencies
4.1 Cosmos Keepers
AccountKeeper

no dependencies

foundational keeper

BankKeeper

Depends on:

AccountKeeper

StakingKeeper

Depends on:

AccountKeeper
BankKeeper

GovKeeper

Depends on:

StakingKeeper

4.2 Ethermint Keepers
EVMKeeper

Depends on:

AccountKeeper
BankKeeper
StakingKeeper


Because :

gas charge uses Bank

sender validation uses Account

voting power is needed for EVM consensus rules

FeeMarketKeeper

Depends on:

EVMKeeper

4.3 PoSS Keeper

Depends on:

AccountKeeper
BankKeeper
StakingKeeper


It also needs:

BeginBlock hooks

consensus timestamp

validator power

But it does not depend on EVM modules.

🧱 5. Keeper Instantiation Order (strict)

The correct and tested instantiation order:

1. AccountKeeper
2. BankKeeper
3. StakingKeeper
4. GovKeeper
5. EVMKeeper
6. FeeMarketKeeper
7. PoSSKeeper


Why this order?

bank requires account

staking requires bank

gov requires staking

evm requires account/bank/staking

feemarket requires evm

poss requires account/bank/staking

Any deviation = chain panic or wrong execution order.

🧠 6. Keeper Interaction Graph
Graph (text-based)
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
     AccountKeeper  BankKeeper  StakingKeeper

          ┌────────────────────┐
          │ FeeMarketKeeper   │
          └───────────▲────────┘
                      │
                  EVMKeeper

          ┌────────────────────┐
          │    PoSSKeeper     │
          └────────────────────┘
               ▲    ▲    ▲
    AccountKeeper BankKeeper StakingKeeper

📦 7. Store Keys for Each Keeper
Keeper	Store Key
AccountKeeper	auth
BankKeeper	bank
StakingKeeper	staking
GovKeeper	gov
EVMKeeper	evm
FeeMarketKeeper	feemarket
PoSSKeeper	noorsignal

All KVStores must be mounted in app.go.

🔌 8. Keeper → ModuleManager Connections

Each keeper must expose:

Message service (MsgServer)

Query service (QueryServer)

Genesis handlers

BeginBlock / EndBlock handlers

Invariants (optional but recommended)

ModuleManager will call:

SetOrderBeginBlockers

SetOrderEndBlockers

SetOrderInitGenesis

SetOrderExportGenesis

RegisterServices

PoSS module will plug into BeginBlock.

🏗️ 9. App Constructor Integration (Blueprint)

In Phase 4C code, the keepers will be placed inside the app constructor:

Create store keys

Create each keeper (in order defined above)

Wire keeper dependencies

Expose keeper references on App struct

Provide keepers to ModuleManager

Install hooks (staking → poss, etc.)

This document ensures implementation is consistent.

🧪 10. Pre-Test Checklist (Before Phase 4C)

Before writing any code:

All keepers defined

All dependencies mapped

Execution order validated

Storage structure stable

PoSS dependencies final

No missing lifecycle hook

No circular dependency

No module requiring additional keeper

This blueprint must be considered final for coding.

🎯 11. Summary Table
Keeper	Depends on	Store	Purpose
Account	none	auth	accounts, signatures
Bank	account	bank	balances, transfers
Staking	account, bank	staking	validators, delegation
Gov	staking	gov	proposals, voting
EVM	account, bank, staking	evm	EVM execution
FeeMarket	evm	feemarket	base fee (EIP-1559)
PoSS	account, bank, staking	noorsignal	social consensus