NOORCHAIN — Phase 4B

PoSS Messages, Queries & Genesis Specification
Version 1.1 — Architecture Only (No Code)

1. Purpose of This Document

This document defines:

the PoSS message types

the PoSS query API (gRPC)

the genesis structure and requirements

the set of PoSS module parameters

the events emitted by the module

This is the final architectural blueprint before implementing the module during Phase 4C.
It contains no code.

2. Message Types (Msg)

PoSS defines exactly one actionable transaction message.

2.1 MsgSubmitSignal

Purpose: Submit a PoSS signal validated by a curator.

Fields:

sender_address

curator_address

signal_type

metadata (optional)

signature (optional if performed off-chain or externally)

Execution context:

Received in DeliverTx

Stored in a pending buffer

Fully validated and rewarded only in BeginBlock

Reject conditions:

invalid curator

invalid signal type

sender equals curator

anti-abuse limit violations

duplicated signal

undefined or invalid weight

Success effect:

the signal is written to pending storage

event poss.signal_submitted is emitted

3. Query API (gRPC)

The module exposes several read-only endpoints for explorers, dApps, wallets and analytics systems.
All queries must be specified in proto/noorsignal/query.proto.

3.1 QuerySignalByID

Returns:

signal metadata

block height

sender

curator

weight

timestamp

3.2 QuerySignalsByAddress

Returns all signals emitted by a given address.
Used for dashboards and account history.

3.3 QueryCuratorStats

Returns:

total validated signals

curator level (Bronze / Silver / Gold)

daily validation counters

accumulated rewards

3.4 QueryParticipantStats

Returns:

total signals emitted

total rewards received

recent PoSS activity

3.5 QueryRewardState

Returns:

last reward block

current reward epoch

current halving cycle

accumulated reward index

total PoSS rewards distributed since genesis

3.6 QueryParams

Returns all PoSS parameters:

weight table

halving schedule

maximum daily signals

curator-level requirements

PoSS Reserve address

base reward per unit

3.7 QueryModuleState

Returns a complete snapshot of the PoSS module state (for archive snapshots and debugging).

4. Genesis Specification

The genesis file must contain all fields required for PoSS to begin in a deterministic and valid state.

4.1 Required Genesis Fields

poss_reserve_address

initial_halving_cycle

blocks_per_halving

base_reward_per_unit

weight_table

max_daily_signals

curator_set (curators + level)

Optional fields:

initial total signals

initial total rewards

initial curator statistics

4.2 Genesis Initialization Logic (InitGenesis)

During chain boot:

load and validate PoSS params

load curator registry

load PoSS reserve address

initialize halving cycle and next halving height

initialize all PoSS KVStore keys

verify PoSS Reserve balance exists in bank module

emit event poss.genesis_initialized

No rewards are computed during InitGenesis.

4.3 Genesis Export (ExportGenesis)

The module must export:

PoSS params

reward state

curator data

halving data

module statistics

5. Module Parameters (Params)

All parameters are stored in types/params.go and exposed in genesis.

Parameter	Description
max_daily_signals	per-address daily signal limit
weight_table	mapping of signal type → weight
blocks_per_halving	number of blocks between halvings
base_reward_per_unit	base reward per weighted unit
min_curator_level	Bronze / Silver / Gold requirement
poss_reserve_address	module account holding PoSS funds
rate_limit_enabled	toggle for anti-abuse logic

These parameters are adjustable via governance after launch.

6. Events Emitted

The PoSS module emits the following events:

Signal Events

poss.signal_submitted

poss.signal_rejected

poss.signal_validated

Reward Events

poss.reward_distributed

poss.reward_participant

poss.reward_curator

System Events

poss.halving_event

poss.state_update

poss.genesis_initialized

Events are required for:

explorers

indexers

analytics systems

dApps

PoSS dashboards

7. Module Integration Rules
DeliverTx

records incoming signals

performs light validation

does not compute rewards

does not apply halving

economic logic is forbidden here

BeginBlock

executes the full PoSS reward cycle

validates all pending signals

computes weighted units

applies halving

distributes 70/30 rewards

updates PoSS state

clears pending signals

emits events

EndBlock

PoSS performs no operations here

Commit

state writes only

8. Summary

The PoSS module provides:

Messages

one message: MsgSubmitSignal

Queries

by signal ID

by address

curator stats

participant stats

reward state

params

full module snapshot

Genesis

full PoSS parameter set

halving settings

base reward

weight table

curator registry

PoSS Reserve address

Block Logic

deterministic reward engine

BeginBlock execution

70/30 reward split

halving every 8 years

strictly no inflation beyond the PoSS Reserve

This document completes the Phase 4B blueprint and is ready for implementation in Phase 4C.
