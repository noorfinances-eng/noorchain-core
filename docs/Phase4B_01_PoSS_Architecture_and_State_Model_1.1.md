NOORCHAIN — Phase 4B

PoSS Architecture & State Model (x/noorsignal)
Version 1.1 — Architecture Only, No Code

1. Purpose of This Document

This document defines the full architecture of the NOORCHAIN Proof of Signal Social (PoSS) module, including:

module purpose and design principles

keeper architecture

store model

reward engine structure

halving model

signal types

validation rules

interaction with Cosmos/EVM layers

This specification is the official reference for implementing x/noorsignal during Phase 4C.

2. Overview of PoSS (Phase 3 Definition)

PoSS is NOORCHAIN’s social consensus mechanism, based on:

social signals emitted by users

curators who validate signals

70% reward to the participant

30% reward to the curator

halving every 8 years

fixed maximum supply: 299,792,458 NUR

no inflation: rewards come from the PoSS Reserve (Genesis 80%)

Key principle:
PoSS logic executes in BeginBlock, not DeliverTx, ensuring deterministic reward calculation and consensus safety.

3. Module Goals

The PoSS module provides:

a registry for all signals

a reward engine

a halving schedule

anti-abuse controls

curator validation structure

secure distribution from PoSS Reserve

a query interface for dApps

pub/sub events for explorers

deterministic behaviour via BeginBlock execution

4. PoSS Module Structure
4.1. Files

The module includes the following structure:

x/noorsignal
• keeper/
– keeper.go
– msg_server.go
– query_server.go
– grpc_query.go
– reward_engine.go

• types/
– keys.go
– types.go
– signals.go
– rewards.go
– halving.go
– params.go
– events.go
– msg.go
– query.go

• module.go
• genesis.go

This file defines structure only. Implementation will be added in Phase 4C.

5. PoSS Store Model

PoSS uses a dedicated KVStore with the store key "noorsignal".

5.1. Data Stored in the KVStore
A. Signal Registry

Each signal entry contains:

sender

curator

signal type

timestamp

block height

weight

optional metadata

Signal types (from Phase 3):

micro-don

participation QR

certified content

CCN broadcast

Weights range from 0.5× to 5× (defined in params).

B. Reward State

Global reward-related metadata:

last reward block

current epoch

current halving cycle

reward index

accumulated signal units per block

optional state hash

C. Anti-Abuse State

Per-address counters:

daily signal count

last signal timestamp

rate-limit flags

Limits come from Phase 3 rules.

D. PoSS Parameters

Stored once at genesis:

max signals/day

weight table

minimum curator level

halving period (8 years in blocks)

reward scaling factors

E. Statistics (Optional)

Used by explorers and dashboards:

total signals

total rewards distributed

curator-level statistics

Not consensus-critical.

6. Keeper Architecture
Responsibilities

The PoSSKeeper:

reads the signal registry

validates new signals

checks curator permissions

computes weighted signal units

determines per-block PoSS rewards

executes reward distribution

updates halving cycle

updates anti-abuse counters

exposes Msg servers

exposes Query servers

emits module events

Dependencies

PoSSKeeper requires access to:

AccountKeeper

BankKeeper

StakingKeeper

It has no dependency on Ethermint.

7. Signal Validation Rules (Phase 3)

Each signal must satisfy:

A. Structural Validation

valid sender

valid curator

valid signal type

valid signature

B. Curator Validation

curator must be registered

must have a valid level (Bronze / Silver / Gold)

must not exceed validation limits

C. Anti-Abuse Rules

maximum daily signals

minimum spacing between signals

curator and sender cannot be the same

no repeated signal from the same sender within the same block

D. Economic Validation

PoSS Reserve must contain enough NUR

the applied weight must be valid

8. Reward Engine Architecture

Rewards are computed per block inside BeginBlock.

8.1. Reward Inputs

number of valid signals in the previous block

weights of each signal

validator power (staking)

current halving cycle

8.2. Reward Output

Per-block reward formula:

reward_total = f(weighted_signal_units, halving_factor)

Distribution:

70% → participant

30% → curator

8.3. Halving Mechanism

Every 8 years (converted to block height):

PoSS reward is divided by 2

PoSSKeeper maintains:

last halving height

next halving height

blocks remaining

Halving is applied at block boundaries.

9. PoSS in the Block Lifecycle
BeginBlock

PoSS main logic:

load last block's signals

compute reward units

compute halving factor

distribute rewards from PoSS Reserve

update PoSS state

clear previous block buffer

emit PoSS events

DeliverTx

Only records signal submissions.
No economic effects happen here.

EndBlock

PoSS performs no actions.

Commit

KVStore updates are committed.

10. PoSS Queries (for apps & explorers)

Required queries:

list of signals

signal by ID

curator statistics

participant statistics

reward statistics

halving information

module parameters

These are essential for the NOORCHAIN Explorer and NOOR Apps.

11. Summary

The PoSS module requires:

a dedicated KVStore (noorsignal)

a PoSSKeeper with account, bank and staking dependencies

a reward engine executed exclusively in BeginBlock

a full signal registry with anti-abuse controls

a halving mechanism every 8 years

strict compliance with Phase 3 specifications

deterministic, auditable behaviour

zero circular dependencies

This document is the canonical specification to implement the PoSS module in Phase 4C.
