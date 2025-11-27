**NOORCHAIN — Phase 4A

App Architecture Map (Cosmos SDK + Ethermint + PoSS)**
Version 1.1 — Structural Overview (no code)

🔧 1. Purpose of This Document

This document provides a clean and visual architecture map of the NOORCHAIN Core Application.

It defines:

application layers

module composition

keeper structure

store layout

dependency graph

app lifecycle overview

No code is included.
This file is the structural reference for Phase 4B (PoSS blueprint) and Phase 4C (Testnet 1.0).

🏛️ 2. Application Layers Overview

NOORCHAIN is composed of three major layers:

2.1 Cosmos SDK Layer (Base Layer)

This layer provides the fundamental blockchain mechanics.

Modules used:

auth – accounts & signatures

bank – balances & token transfers

staking – validators, delegation system

gov – governance & proposals

params – configuration subsystem (implicit in SDK 0.50.x)

Purpose:
→ Handle state integrity, consensus-critical logic, and chain governance.

2.2 Ethermint Layer (EVM Layer)

Provides full Ethereum compatibility inside Cosmos.

Modules used:

evm – EVM execution, state DB, logs, gas rules

feemarket – dynamic base fee (EIP-1559-like model)

Purpose:
→ Allow smart contracts, dApps, wallets like MetaMask, Ethereum RPC, etc.

2.3 Custom Layer (NOORCHAIN-Specific)
x/noorsignal (PoSS module)

Defined in Phase 4B :

PoSS signal ingestion

reward calculation engine

halving schedule (every 8 years)

70/30 participant/curator split

hooks into block lifecycle

PoSS state store (KVStore)

Purpose:
→ Implement NOORCHAIN’s unique Social Consensus & Reward Layer.

🗂️ 3. App Composition Structure

The full app is composed around five core components:

3.1 BaseApp (Cosmos)

The engine executing:

transactions

messages

ABCI calls

block lifecycle (BeginBlock / EndBlock)

3.2 Encoding System

Includes:

interface registry

amino codec (legacy)

protobuf codec (primary)

signing modes

3.3 Store System

Each module gets a:

KVStore

Transient store (if needed)

Memory store (for params)

Store keys include:
auth, bank, staking, gov, evm, feemarket, noorsignal

3.4 Keepers

Go structs responsible for:

reading/writing state

executing module logic

interacting with other keepers

3.5 ModuleManager

Registers:

modules

services

genesis init

begin/end block order

🧩 4. Keeper Dependency Graph
4.1 Cosmos Keepers
AccountKeeper → BankKeeper → StakingKeeper → GovKeeper


Bank requires Account.
Staking requires Account + Bank.
Gov requires Staking.

4.2 Ethermint Keepers
EVMKeeper → AccountKeeper, BankKeeper, StakingKeeper
FeeMarketKeeper → EVMKeeper

4.3 PoSS Keeper (Custom)
PoSSKeeper → AccountKeeper
            → BankKeeper
            → StakingKeeper
            → (hooks) BeginBlock, EndBlock


The PoSS module depends on staking and bank to access balances and validator power.

🗃️ 5. Store Layout Map
RootStore (IAVL)
root  
│
├── auth        (KVStore)
├── bank        (KVStore)
├── staking     (KVStore)
├── gov         (KVStore)
├── evm         (KVStore)
├── feemarket   (KVStore)
└── noorsignal  (KVStore)


All stores are mounted into BaseApp at app initialization.

🔄 6. Block Lifecycle Map
6.1 BeginBlock sequence
1. FeeMarket module updates base fee
2. EVM module prepares EVM block context
3. Staking module runs validator updates
4. PoSS module processes signals / rewards (Phase 4B)
5. Governance tallies ongoing proposals

6.2 DeliverTx sequence
1. AnteHandler (signature verification, fees)
2. Msg routing to module
3. State transitions
4. Gas accounting

6.3 EndBlock sequence
1. Staking updates (validator set)
2. Gov updates

6.4 Commit

State root committed via IAVL.

🌍 7. Genesis Lifecycle Map

At genesis:

Accounts created

Balances credited

Staking params set

Gov params set

EVM genesis loaded

FeeMarket genesis loaded

PoSS placeholder initialized

Validator set generated

Chain starts at height 1

🧠 8. Inter-module Interaction Summary
auth ↔ bank

account numbers

balances

signatures

bank ↔ staking

delegation shares

token transfers locked/unlocked

staking ↔ gov

voting power

proposal weights

evm ↔ feemarket

block gas cost calculation

EIP-1559 dynamic base fee

noorsignal ↔ staking

PoSS reward distribution depends on validator power

potential slashing conditions (later)

🎯 9. App Constructor Map (app.go)

The constructor (later implemented in code) must:

Build encoding config

Create BaseApp

Define store keys

Instantiate all keepers

Set keeper relationships

Create ModuleManager

Register services

Register BeginBlock/EndBlock

Set InitGenesis/ExportGenesis functions

Load latest state

Return application instance

This map ensures deterministic construction.

📌 10. Summary Table
Layer	Component	Purpose
Cosmos Base	auth/bank/staking/gov	Core blockchain logic
EVM Layer	evm/feemarket	Ethereum compatibility
Custom Layer	noorsignal	PoSS social consensus
App System	BaseApp, Keepers, Stores	Infrastructure
Runtime	BeginBlock/EndBlock	Chain lifecycle
