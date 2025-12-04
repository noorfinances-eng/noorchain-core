NOORCHAIN 1.0 — Core Architecture
Phase 3 Documentation
Version 1.1 — English
1. Overview

NOORCHAIN 1.0 is a sovereign Swiss blockchain built with a hybrid Cosmos/EVM foundation and an ethical economic model based on Proof of Signal Social (PoSS).

Core design elements:

Cosmos SDK v0.46.11

Ethermint v0.22.0 (full EVM compatibility)

CometBFT/Tendermint v0.34.27

Fixed token supply: 299,792,458 NUR

Halving period: every 8 years

PoSS reward split: 70% participant / 30% curator

Legal model: Swiss Legal Light CH, non-custodial and non-financial

Status: Phase 4 implementation complete, providing a fully functional chain core

At this stage, the chain runs with all essential modules, stores, keepers, ABCI lifecycle logic and EVM execution.
The PoSS logic exists structurally but is not yet activated or distributing rewards.

2. Core Modules
Module	Status	Description
auth	Complete	Account management, signatures, public key system
bank	Complete	Balances, transfers, supply tracking
staking	Complete	Validators, delegations, baseline staking logic
gov	Store ready	Governance store and params subspace present (keeper wiring planned)
params	Complete	Global parameters, per-module subspaces
evm	Complete	Full EVM execution through Ethermint
feemarket	Complete	EIP-1559 dynamic fee system
noorsignal	Structural skeleton	Base PoSS module scaffolding, no active logic yet

All modules are registered in the application, with their stores mounted and available.

3. Keeper Architecture

The NoorchainApp currently includes the following keepers:

AccountKeeper

BankKeeper

StakingKeeper

ParamsKeeper

FeeMarketKeeper

EvmKeeper

NoorSignalKeeper (structural, no rewards logic yet)

Each module has:

A dedicated key (store key)

Optional transient stores

Correct mounting order in the BaseApp initialization flow

This configuration ensures a fully operational blockchain even before PoSS activation.

4. AnteHandler

NOORCHAIN uses Ethermint’s NewAnteHandler, supporting both Cosmos and Ethereum-style transactions.

Capabilities include:

Cosmos SDK signature checks

EVM signatures (MsgEthereumTx)

Gas consumption and validation rules

Min gas price and dynamic fee enforcement

Ethereum-specific decorators for execution safety

Compatibility with EIP-1559 fee market

This makes NOORCHAIN a dual-transaction blockchain:
Cosmos messages + full Ethereum tx support.

5. ABCI Lifecycle

The following ABCI methods are implemented and functional:

InitChainer

Parses and initializes state from the genesis file

Configures module parameters and initial states

BeginBlocker

Registered for all modules

Currently empty (PoSS logic will later use it)

EndBlocker

Registered for all modules

Currently empty (PoSS reward distribution may connect here in future phases)

Modules are initialized in correct dependency order, ensuring deterministic state execution.

6. NOOR Signal Module (PoSS) — Current State

The NoorSignal module exists structurally but does not yet implement PoSS logic.

Currently implemented (Phase 3–4 structural work):

Module folder structure

Keeper structure

Genesis import and export

Module registration in the ModuleManager

ABCI BeginBlock and EndBlock hooks (empty handlers)

Placeholder event, parameter and state files

Not yet implemented (PoSS Logic Phases 1–5):

Parameters (PoSSEnabled, weights, limits, halving period)

Signal types and validation structure

Reward calculation

Halving mechanism

Curator roles and verification

Daily counters and anti-abuse rules

Storage and indexing of signals

Queries (signal, curator, config, counters)

Msg handlers (SubmitSignal, ValidateSignal, AddCurator, etc.)

These elements are added gradually across PoSS Logic 1 → 10 and will become active only in Phase 6 (activation phase).

7. Status Summary (Phase 3 Completed → Phase 4 Completed)
Fully Completed

BaseApp

Module registration

Store mounting

Keepers initialization

AnteHandler (Cosmos + Ethereum)

EVM execution

Fee Market system

Genesis import/export

Basic ABCI lifecycle

Structured but inactive

PoSS module

Governance keeper wiring

PoSS BeginBlock/EndBlock logic

Reward distribution

Curator validation mechanism

NOORCHAIN is now a fully operational core blockchain ready for:
Testnet, governance integration, and full PoSS activation.

End of Document — Version 1.1
