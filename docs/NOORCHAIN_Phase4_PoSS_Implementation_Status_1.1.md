NOORCHAIN — Phase 4 PoSS Implementation Status (v1.1)

Full technical status of the PoSS module (x/noorsignal) — December 2025

✅ 1. Purpose of Phase 4

Phase 4 is the actual implementation phase of the NOORCHAIN PoSS module, without enabling minting or governance yet.
The objective is to build a fully functional internal PoSS logic, deterministic and validated, ready for future wiring with BankKeeper, EVM, and governance.

This document captures the exact state of the PoSS module in the repository at this moment.

✅ 2. PoSS Architecture — Overview

The x/noorsignal module contains:

types/

Params

GenesisState

SignalType enum

Reward helpers (ComputeSignalReward)

MsgCreateSignal

MsgCreateSignalResponse

Store keys, daily counters

Pure Go validation

keeper/

Core Keeper logic

Daily per-user counters

Daily reward caps

PendingMint planning queue

Genesis load/store

Global PoSS statistics view

module.go

AppModuleBasic / AppModule

InitGenesis / ExportGenesis

BeginBlock (currently empty)

Fully integrated in app/app.go

handler.go

Minimal stub (required for SDK 0.46)

Does not process any tx yet

✅ 3. PoSS Parameters (fully implemented)
Available parameters:

PoSSEnabled

MaxSignalsPerDay

MaxSignalsPerCuratorPerDay

MaxRewardPerDay (sdk.Coin)

BaseReward (sdk.Coin)

WeightMicroDonation

WeightParticipation

WeightContent

WeightCCN

HalvingPeriodBlocks

PoSSReserveDenom = "unur"

Behavior:

If PoSSEnabled = false, PoSS counts signals but produces 0/0 rewards.

All parameters validated before storage.

GetParams() auto-initializes defaults if no params exist.

Governance-ready.

✅ 4. Reward Engine (70/30 split)

Function:
ComputeSignalReward(params, signalType, height)

Steps:

BaseReward × weight

Halving every 8 years (via block height)

Fixed split:

70% to participant

30% to curator

Pure, deterministic, no store access.

✅ 5. GenesisState (PoSS)
Stored fields:

TotalSignals

TotalMinted

Stable 70/30 rules

Reserved space for PendingMint

Implemented:

DefaultGenesis()

ValidateGenesis()

InitGenesis()

ExportGenesis()

Stored as pure JSON, no proto dependency.

✅ 6. Daily Counters
Storage:

Participant daily signal counter
Key → DailyCounterKey(address,date)
Value → uint64

Participant daily reward tracker
Key → daily_reward:<addr>:<date>
Value → uint64 (big-endian)

Functions:

GetDailySignalsCount

IncrementDailySignalsCount

getDailyRewardAmount

setDailyRewardAmount

✅ 7. MaxRewardPerDay Cap

Behavior:

If MaxRewardPerDay = 0, cap disabled.

Otherwise:

If participant already reached the cap → reward = 0/0

The signal still increases the daily counter

If adding new reward would exceed the cap → reward = 0/0

Otherwise → store updated reward sum

Handled entirely inside ProcessSignalInternal().

✅ 8. Internal PoSS Pipeline (ProcessSignalInternal)

This is the core functional engine of PoSS v1.

Steps performed:

Compute raw reward (with halving + weights)

Apply MaxRewardPerDay

Increment participant daily counter

Create PendingMint entry

Update global:

TotalSignals

TotalMinted

Return theoretical rewards to caller

What it does not do yet:

enforce MaxSignalsPerDay

enforce MaxSignalsPerCuratorPerDay

real minting (BankKeeper)

sending coins

curator counters

event emission (ABCI)

MsgServer proto implementation

QueryServer implementation

These will be implemented in Phase 4D → 4F → 6.

✅ 9. PendingMint Queue

PoSS supports scheduling future mint operations (not active yet).

Key format:

pending_mint:<height>:<participant>:<timestamp_nano>


Stored JSON object:

block height

timestamp

participant

curator

signal type

participantReward

curatorReward

Mint is not executed yet (Legal Light).

✅ 10. MsgCreateSignal (pure Go)

Message fields:

Participant (bech32 noor1…)

Curator (bech32)

SignalType

Metadata

Timestamp

Date (YYYY-MM-DD)

Implements:

sdk.Msg

proto.Message (minimal stub)

Validation:

full bech32 check

signal type check

date format check

Unit tests → PASS.

✅ 11. Testing Status

Keeper tests → PASS

Msg tests → PASS

Module compiles cleanly

go test ./... → ALL GREEN

Handler stub does not break build

✅ 12. Testnet infrastructure

Script:
scripts/testnet.sh

Creates:

data-testnet/

data-testnet/config/genesis.json
(copied from testnet/genesis.json)

Result:

A clean home directory ready for ./noord start

Note:
noord start isn't wired yet → expected at this stage.

🎯 13. Phase 4 Completion Percentage

Current PoSS internal logic completion:

≈ 85–90% DONE

Remaining:

MsgServer proto + routing

QueryServer (gRPC + REST)

Daily curator counters

Bank wiring (mint/send)

Events

Integration with EVM hooks

Governance activation of parameters

Testnet activation (PoSS OFF)

🔜 14. Next Recommended Steps (Phase 4D–4F)

PoSS MsgServer (full proto)

PoSS QueryServer

BankKeeper wiring (mint + transfer)

Events & telemetry

Daily curator limits

Full PoSS flowchart / diagrams

Testnet 1.0 with PoSS OFF

🤝 15. Summary

The PoSS module is now:

structurally complete

deterministic

compliant with Cosmos SDK 0.46

Legal Light compliant (no mint yet)

fully testable

ready for progressive activation in later phases

NOORCHAIN now has a solid, professional, Swiss-grade foundation for PoSS.