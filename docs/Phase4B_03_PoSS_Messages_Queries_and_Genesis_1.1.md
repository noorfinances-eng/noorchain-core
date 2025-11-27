**NOORCHAIN — Phase 4B

PoSS Messages, Queries & Genesis Specification**
Version 1.1 — Architecture Only (No Code)

🔧 1. Purpose of This Document

This document defines:

the messages (Msg) accepted by PoSS

the query API exposed by PoSS

the genesis requirements

the minimal set of parameters the module must support

the events emitted by the module

This is the final blueprint required before starting to code the module in Phase 4C.

No code is included.

🧩 2. Message Types (Msg)

PoSS has only one actionable message, as defined in Phase 3:

2.1 MsgSubmitSignal
Purpose

Submit a PoSS signal validated by a curator.

Fields
sender_address
curator_address
signal_type
metadata (optional)
signature (optional if handled outside)

Requirements

binding to DeliverTx (not BeginBlock)

stored in pending buffer

validated later in BeginBlock

no reward distributed here

Reject Conditions

invalid curator

invalid signal type

anti-abuse violations

sender == curator

duplicated signal

weight not defined

Success Effect

signal stored in pending state

event emitted

🔍 3. Query API (gRPC)

The module exposes several query endpoints for dApps, explorers, and wallets.

All queries MUST be defined in proto/noorsignal/query.proto.

3.1 QuerySignalByID

Fetch a signal by its unique ID.

Returns:

signal details

block height

curator

sender

weight

timestamp

3.2 QuerySignalsByAddress

Return all signals from a specific address.

Useful for:

user dashboards

explorers

stats pages

3.3 QueryCuratorStats

Return:

total validated signals

curator level (Bronze/Silver/Gold)

daily validation counters

rewards received

3.4 QueryParticipantStats

Return:

total emitted signals

rewards received

recent activity

3.5 QueryRewardState

Return:

last_reward_block

current reward epoch

current halving cycle

accumulated reward index

total PoSS rewards given

3.6 QueryParams

Return PoSS module parameters, including:

weight table

halving schedule

max daily signals

curator rules

PoSS Reserve address

base_reward_per_unit

3.7 QueryModuleState

Return the entire PoSS module state for snapshots.

🧱 4. Genesis Specification

The genesis file must include all fields necessary for PoSS to begin in a valid state.

4.1 PoSS Genesis Fields
Required fields:
poss_reserve_address
initial_halving_cycle
blocks_per_halving
base_reward_per_unit
weight_table
max_daily_signals
curator_set (list of curators + level)
initial_stats (optional)

Optional fields:

initial total signals

initial accumulated rewards

initial curator stats

4.2 Genesis Initialization Logic (InitGenesis)

During chain initialization:

Load params

Load curator registry

Load PoSS Reserve address

Set halving cycle + next halving height

Set base reward per unit

Initialize KVStore keys

Validate reserve balance in bank module

Emit event poss.genesis_initialized

No rewards are processed during InitGenesis.

4.3 Genesis Export (ExportGenesis)

Module should export:

params

reward state

curator data

stats

halving data

📦 5. Module Parameters (Params)

Stored in types/params.go.

The following parameters must be defined:

Parameter	Description
max_daily_signals	per-address limit
weight_table	mapping of signal type → weight
blocks_per_halving	number of blocks per halving
base_reward_per_unit	reward unit multiplier
min_curator_level	Bronze/Silver/Gold
poss_reserve_address	funding module account
rate_limit_enabled	anti-abuse toggle

All parameters are part of genesis.

📡 6. Events Emitted

The PoSS module must emit:

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

Events enable:

explorers

dApps

indexers

analytics

🧠 7. Module Integration Rules
DeliverTx

only records signals

NO rewards

NO halving

NO economic logic

BeginBlock

runs full PoSS reward cycle

deterministic

reward distribution

halving

anti-abuse updates

EndBlock

PoSS does nothing here

Commit

state commit only

🎯 8. Summary

The PoSS module provides:

Msg

1 Msg: MsgSubmitSignal

Queries

by signal

by address

curator stats

participant stats

reward state

params

full module snapshot

Genesis

includes all parameters

halving schedule

reward base

curator registry

PoSS reserve settings

Block Logic

BeginBlock-only reward engine

deterministic operations

70/30 reward split

halving every 8 years

no inflation

This file completes the PoSS Blueprint and prepares for coding in Phase 4C.