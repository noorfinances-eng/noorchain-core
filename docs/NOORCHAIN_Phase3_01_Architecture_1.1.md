NOORCHAIN 1.0 — Core Architecture (Phase 3 Documentation)

Version: 1.1
Language: English

1. Overview

NOORCHAIN 1.0 is a sovereign Swiss blockchain built with:

Cosmos SDK v0.46.11

Ethermint v0.22.0 (EVM compatibility)

CometBFT/Tendermint v0.34.27

PoSS (Proof of Signal Social) economic consensus

Fixed supply: 299,792,458 NUR (speed of light)

Halving: every 8 years

Reward split: 70% participant / 30% curator

The Phase 4 implementation provides a full core chain without business logic yet.

2. Core Modules
Module	Status	Description
auth	✅	Accounts, signatures, public keys
bank	✅	Balances, transfers
staking	✅	Validators, delegations (future PoSS integration)
gov	🟧	Store and param subspace exist, keeper not wired yet
params	✅	Global param keeper and per-module subspaces
evm	✅	Full EVM execution (Ethermint)
feemarket	✅	EIP-1559 dynamic fees
noorsignal	🟧 (skeleton)	PoSS module placeholder
3. Keepers

The NoorchainApp currently includes:

AccountKeeper

BankKeeper

StakingKeeper

ParamsKeeper

FeeMarketKeeper

EvmKeeper

NoorSignalKeeper (empty)

Each store (KV/transient) is mounted and fully functional.

4. AnteHandler

NOORCHAIN uses the Ethermint NewAnteHandler() to support both:

Cosmos SDK transactions

Ethereum-compatible transactions (MsgEthereumTx)

This includes:

Gas checks

Signature verification

EVM-specific decorators

Min gas price enforcement

Dynamic fee rules (FeeMarket)

5. ABCI Lifecycle

The following methods are implemented:

InitChainer — accepts a JSON genesis state

BeginBlocker — empty for now

EndBlocker — empty for now

Each module is initialized in the correct order.

6. NOOR Signal Module (PoSS) [Skeleton]

Currently implemented:

Module structure

Keeper structure

Genesis import/export

ABCI Begin/EndBlock triggers (empty)

Module registration in ModuleManager

Not yet implemented:

Params

Signal types

Reward logic

Halving logic

Curator validation

Storage model

Queries

Msg handlers

These features belong to PoSS Logic 1–5 (next phases).