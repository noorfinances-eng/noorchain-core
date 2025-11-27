**NOORCHAIN — Phase 4A

Store & State Model Blueprint**
Version 1.1 — Architecture Only (No Code)

🔧 1. Purpose of This Document

This document defines the entire state model and KVStore architecture of NOORCHAIN 1.0, including:

store key layout

underlying IAVL structure

module state models

interaction between state trees

PoSS state requirements

indexing requirements

persistence rules

This is the authoritative reference for implementing stores during Phase 4C.

🗂️ 2. Multistore Overview (IAVL)

NOORCHAIN uses the Cosmos SDK Multistore, built on:

IAVL Merkle Trees for persistent KV stores

Memory stores for temporary data

Transient stores for ephemeral block data

The root store contains all module KVStores mounted under unique store keys.

🗄️ 3. Store Key Map
Store Key	Module
auth	AccountKeeper
bank	BankKeeper
staking	StakingKeeper
gov	GovKeeper
evm	EVMKeeper
feemarket	FeeMarketKeeper
noorsignal	PoSSKeeper

Each is associated with an IAVL-backed KV store.

🧱 4. Module State Models

Below is the complete description of every module’s internal state.

4.1 auth Module State

Stores:

account numbers

public keys

account sequences

base account data

Required for:

signature verification

sequence (nonce) handling

EVM compatibility

4.2 bank Module State

Stores:

balances (address → coins)

supply

denomination metadata

Bank state is heavily used by:

staking

evm gas accounting

poss reward distribution

4.3 staking Module State

Stores:

validators

delegations

unbonding delegations

redelegations

staking params

Also provides:

validator power mapping

slashing information (future use)

This state is critical for PoSS.

4.4 gov Module State

Stores:

proposals

deposits

votes

tally results

governance params

PoSS does not directly modify governance, but governance depends on staking.

4.5 evm Module State

Stores:

EVM state DB

contract storage

logs

nonce

code hash

block bloom filters

This is a large subtree inside the multistore.

4.6 feemarket Module State

Stores:

base fee

block gas metrics

EIP-1559 dynamics

Critical for:

EVM execution

block validity

gas pricing

4.7 noorsignal (PoSS) Module State

This is a custom state tree for NOORCHAIN.

Stores:

1. Signal registry

All PoSS signals:

signal_id → {
  sender_address,
  curator_address,
  signal_type,
  timestamp,
  weight,
  block_height
}

2. Reward state

Tracked values:

current reward epoch

reward indexes

last reward block

halving cycle data

3. Anti-abuse state

Stores per-account counters:

daily signal count

last participation timestamp

rate-limit data

4. PoSS parameters

max daily signals

reward weights

halving schedule

70/30 split

5. Statistics (optional)

total signals

total PoSS rewards minted (from reserve)

curator-level stats

🔗 5. Inter-Store Interactions
PoSS ↔ staking

validator power needed for reward weighting

block height used for halving windows

PoSS ↔ bank

distributes NUR from PoSS reserve

checks balances

writes reward transfers

EVM ↔ bank

gas deduction

intrinsic gas cost

tx sender checks

staking ↔ gov

governance voting power derived from validator power

bank ↔ auth

balances linked to account structures

🧠 6. State Consistency Requirements

NOORCHAIN enforces:

deterministic writes in BeginBlock/EndBlock

no circular dependencies

no state modification in EndBlock (except staking/gov)

PoSS NEVER writes inside DeliverTx

only inside BeginBlock

ensures consistency and predictable reward flow

🏗️ 7. State Persistence Guarantees

Each commit:

IAVL stores commit and return hashes

root hash is passed to CometBFT

Ethereum state DB commits separately

PoSS state must commit after EVM state

This ensures:

consistent app hash

correct EVM bloom filters

correct PoSS indexing

🌍 8. Store Architecture Diagram
root (IAVL)
│
├── auth
├── bank
├── staking
├── gov
├── evm
├── feemarket
└── noorsignal


Each module store is a subtree of the IAVL Merkle root.

📌 9. Summary Table
Module	Stores	Critical For
auth	accounts	signatures, sequences
bank	balances, supply	staking, evm, poss
staking	validators, delegations	governance, poss
gov	proposals, votes	governance
evm	stateDB, logs	EVM compatibility
feemarket	base fee	gas pricing
noorsignal	signals, rewards, anti-abuse	PoSS social consensus
🎯 10. Final Notes

This blueprint ensures:

stable data model

deterministic operations

clean integration of EVM + PoSS

predictable reward distribution

safe Testnet 1.0 initialization

No change to this architecture is allowed without updating Phase 3.