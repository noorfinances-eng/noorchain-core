**NOORCHAIN — Phase 4B

PoSS Architecture & State Model (x/noorsignal)**
Version 1.1 — Architecture Only, No Code

🔧 1. Purpose of This Document

This document defines the full architecture of the NOORCHAIN Proof of Signal Social (PoSS) module, including:

module purpose and design principles

keeper architecture

store model

reward engine structure

halving model

signal types

validation rules

interaction with Cosmos/EVM layers

This specification is the official reference to implement x/noorsignal during Phase 4C.

🌞 2. Overview of PoSS (Phase 3 Definition)

PoSS is NOORCHAIN’s social consensus mechanism, based on:

social signals emitted by users

curators who validate signals

70% reward to the participant

30% reward to the curator

halving every 8 years

fixed maximum supply: 299,792,458 NUR

no inflation: rewards come from PoSS Reserve (Genesis 80%)

Key principle:

PoSS is processed in BeginBlock, not DeliverTx.

This ensures deterministic rewards and fixed consensus behavior.

🧩 3. Module Goals

The PoSS module provides:

A registry for all signals

A reward engine

A halving schedule

Anti-abuse controls

Curator validation structure

A secure distribution from PoSS Reserve

A query interface for dApps

Pub/sub events for explorers

Deterministic BeginBlock execution

🏛️ 4. PoSS Module Structure
4.1 Files

The module will contain:

x/noorsignal/
│
├── keeper/
│   ├── keeper.go
│   ├── msg_server.go
│   ├── query_server.go
│   ├── grpc_query.go
│   └── reward_engine.go
│
├── types/
│   ├── keys.go
│   ├── types.go
│   ├── signals.go
│   ├── rewards.go
│   ├── halving.go
│   ├── params.go
│   ├── events.go
│   ├── msg.go
│   └── query.go
│
├── module.go
└── genesis.go


(This file lists structure only. Implementation will come in Phase 4C.)

🧱 5. PoSS Store Model

The PoSS module uses a dedicated KVStore:

Store Key: "noorsignal"

5.1 Data Stored in KVStore
A. Signal Registry
signal_id → {
  sender,
  curator,
  signal_type,
  timestamp,
  block_height,
  weight,
  metadata (optional)
}


Signal types (Phase 3 definition):

micro-don

participation QR

certified content

CCN broadcast

Weights: 0.5× → 5× (defined in params)

B. Reward State

Stores global PoSS reward metadata:

last_reward_block

current_epoch

current_halving_cycle

reward_index

accumulated signal units (per block)

latest_state_hash (optional)

C. Anti-Abuse State

Per-address counters:

address → {
  daily_signal_count,
  last_signal_timestamp,
  rate_limit_flags
}


Limits come from Phase 3 rules.

D. PoSS Parameters

Stored once in genesis:

max signals/day

weight table

minimum curator level

halving period (8 years in blocks)

reward scaling factors

E. Statistics (Optional)

Track global stats:

total_signals

total_rewards_distributed

curator-level stats

(Useful for explorers, not consensus-critical.)

🔗 6. Keeper Architecture
PoSSKeeper responsibilities

The keeper:

reads signal registry

validates new signals

checks curator permissions

computes weighted signals

determines per-block PoSS reward

executes reward distribution

updates halving cycle

updates anti-abuse counters

exposes Msg servers

exposes Query servers

emits events

Dependencies
PoSSKeeper → AccountKeeper
           → BankKeeper
           → StakingKeeper


No dependency on Ethermint.

🚦 7. Signal Validation Rules (Phase 3)

Every signal must pass:

A. Structural Validation

valid sender

valid curator

valid signal type

valid signature

B. Curator Validation

curator must be a registered curator

curator must have level: Bronze / Silver / Gold

curator must not exceed validation limits

C. Anti-Abuse

maximum daily signals

minimum block spacing

curator/sender can't be the same

no repeated signals from same sender in same block

D. Economic Validation

PoSS Reserve must have enough NUR

weight must be valid

🔥 8. Reward Engine Architecture

Rewards occur per block inside BeginBlock, not per transaction.

8.1 Reward Inputs

number of valid signals in previous block

weights of signals

validator power (staking)

current halving cycle

8.2 Reward Output

Per block:

reward_total = f(weighted_signal_units, halving_factor)


Distribution:

70% → participant
30% → curator

8.3 Halving Mechanism

Every 8 years (in block height):

reward = reward / 2


PoSSKeeper stores:

last halving height

next halving height

remaining blocks

Halving is applied at block boundary.

🔄 9. PoSS in Block Lifecycle
BeginBlock

PoSS begins its main logic:

load all signals from previous block

compute reward units

compute halving factor

send rewards from PoSS Reserve

update PoSS state

clear last block buffer

emit PoSS events

DeliverTx

Only records signal submission.

No rewards, no state changes, no economics happen here.

EndBlock

PoSS does nothing here → deterministic.

Commit

PoSS commits KVStore updates.

🧠 10. PoSS Queries (for apps & explorers)

Queries needed:

list signals

signal by ID

curator stats

participant stats

reward stats

halving info

module parameters

Necessary for the NOORCHAIN Explorer & NOOR Apps.

🎯 11. Summary

The PoSS module requires:

a KVStore (noorsignal)

a PoSSKeeper with staking, bank, account dependencies

a reward engine running exclusively in BeginBlock

signal registry + anti-abuse system

halving mechanism every 8 years

strict adherence to Phase 3 rules

no circular dependencies

purely deterministic behavior

This document is the canonical specification for Phase 4C module implementation.