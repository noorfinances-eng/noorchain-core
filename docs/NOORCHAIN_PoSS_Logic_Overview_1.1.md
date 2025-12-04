NOORCHAIN — PoSS Logic Overview (1–25)

Version: 1.1
Module: x/noorsignal (PoSS — Proof of Signal Social)

This document summarizes the current state of the PoSS logic in NOORCHAIN after PoSS Logic steps 1 to 25.
It is meant for developers, auditors and product/BD people who need a clear, high-level understanding of what is already implemented in the core and what is still intentionally missing.

1. Scope of the PoSS Module (x/noorsignal)

The x/noorsignal module is responsible for:

Tracking social signals (Proof of Light) as PoSS signals.

Enforcing daily social limits (later).

Computing rewards based on:

a BaseReward (in unur),

per-signal-type weights,

halving over time,

a structural 70/30 split between participant and curator.

Recording pending mints (planned rewards) without touching real balances.

Maintaining global PoSS counters:

TotalSignals

TotalMinted (planned PoSS minting in unur).

Exposing a read-only global stats view (PoSSStats) for dashboards, CLI and monitoring tools.

The module is intentionally conservative at this stage:
no real minting, no bank movements, no governance wiring yet.

2. Core Data Structures
2.1. GenesisState

File: x/noorsignal/types/genesis.go

Structure:

TotalSignals (uint64) — counts how many PoSS signals have been processed since genesis.

TotalMinted (string) — tracks the planned PoSS issuance in unur.

MaxSignalsPerDay (uint32) — first anti-abuse guardrail (per address).

ParticipantShare (uint32) — fixed to 70.

CuratorShare (uint32) — fixed to 30.

Key decisions:

The 70/30 split is structural and must always sum to 100:

ParticipantShare = 70

CuratorShare = 30

The genesis is serialized as plain JSON in the PoSS KVStore
(no protobuf, no gogo/proto dependency).

2.2. Params (PoSS behaviour)

File: x/noorsignal/types/params.go

Structure:

PoSSEnabled (bool)

MaxSignalsPerDay (uint64)

MaxSignalsPerCuratorPerDay (uint64)

MaxRewardPerDay (sdk.Coin)

BaseReward (sdk.Coin)

WeightMicroDonation (uint32)

WeightParticipation (uint32)

WeightContent (uint32)

WeightCCN (uint32)

PoSSReserveDenom (string)

HalvingPeriodBlocks (uint64)

Current behaviour:

DefaultParams() returns a “safe off” configuration:

PoSSEnabled = false

conservative daily limits

rewards defined in unur but effectively disabled

The halving is configured structurally (parameter exists) but the exact HalvingPeriodBlocks is still a placeholder.

Later, these params will be stored in a real ParamSubspace and controlled by on-chain governance.

2.3. Rewards helpers

File: x/noorsignal/types/rewards.go

The PoSS reward pipeline (pure functions) is:

Base reward
base = ComputeBaseReward(params, signalType)
→ BaseReward * weight(signal_type)

Halving
halved = ApplyHalving(params, height, base)
→ divide by 2 every HalvingPeriodBlocks (never below 0).

70/30 split
participant, curator, err = SplitReward70_30(halved)

High-level helper
participant, curator, err = ComputeSignalReward(params, signalType, height)

Rules:

If PoSSEnabled == false, ComputeSignalReward returns zero rewards with the correct denom (no surprise minting).

When PoSSEnabled == true, the pipeline is: base reward → halving → 70/30.

2.4. Daily counters

Daily counters are per address, per date.

Types and helpers:

DailyCounter — documented in types/daily_counter.go.

Keys — DailyCounterKey(address, date) in types/daily_counter_key.go.

Keeper functions:

GetDailySignalsCount(ctx, address, date)

SetDailySignalsCount(ctx, address, date, count)

IncrementDailySignalsCount(ctx, address, date)

They are updated inside the internal signal pipeline, but daily limits are not yet enforced. This will be part of a later step.

3. Pending Mint Queue (planning only)

File: x/noorsignal/types/pending_mint.go (exact name may differ depending on repo)

Concept:

Every processed PoSS signal creates a PendingMint entry in the PoSS store.

Each entry contains:

participant and curator addresses,

signal type,

planned rewards (participant + curator),

block height and timestamp.

Keeper helper:

Keeper.RecordPendingMint(ctx, participantAddr, curatorAddr, signalType, participantReward, curatorReward)

Important:

This does not mint coins.

It is a planning mechanism for a later minting phase, which can:

run in a separate handler,

be audited,

be rate-limited,

respect Legal Light constraints.

4. Internal Signal Pipeline (without real minting)

File: x/noorsignal/keeper/keeper.go

Internal, non-public pipeline:

Keeper.ProcessSignalInternal(ctx, participantAddr, curatorAddr, signalType, date) (participantReward, curatorReward, error)

Steps:

Compute rewards (base → halving → 70/30).

Increment participant daily counter for the given date.

Record a PendingMint entry.

Update GenesisState:

TotalSignals++

TotalMinted += participantReward + curatorReward

Return the two reward coins to the caller.

What it does not do yet:

enforce daily limits,

enforce curator limits,

move any Bank balances.

This is the “sandbox pipeline” to keep PoSS deterministic and auditable before turning real money on.

5. Global PoSS Stats (PoSSStats)

PoSS Logic step 24 introduced a simple, consolidated view.

File: x/noorsignal/types/stats.go

Structure PoSSStats:

TotalSignals (uint64)

TotalMinted (string)

PoSSEnabled (bool)

MaxSignalsPerDay (uint64)

MaxSignalsPerCuratorPerDay (uint64)

PoSSReserveDenom (string)

Keeper helper:

Keeper.GetGlobalStats(ctx) PoSSStats

Behaviour:

Reads GenesisState from the KVStore.

Reads PoSS Params (currently via DefaultParams()).

Builds a single read-only snapshot.

5.1 Example JSON output (future query)

In a future noord query poss stats, a typical JSON payload could look like:

total_signals: 12345

total_minted: "987654321"

poss_enabled: true

max_signals_per_day: 20

max_signals_per_curator_per_day: 100

poss_reserve_denom: "unur"

Interpretation:

12 345 PoSS signals processed since genesis.

987 654 321 unur planned to be minted through PoSS (70/30 split included).

PoSS rewards are currently enabled.

Each participant can emit up to 20 signals per day.

Each curator can validate up to 100 signals per day.

Rewards and reserves are denominated in unur.

6. What Is Missing on Purpose (Post-Logic-25 Roadmap)

The following parts are not implemented yet, by design:

Real minting of unur or movement of balances in x/bank.

Full enforcement of:

MaxSignalsPerDay,

MaxSignalsPerCuratorPerDay,

MaxRewardPerDay.

Governance wiring for PoSS params (ParamSubspace + x/gov).

Public Msg/Query gRPC endpoints (proto files, Msg server, Query server).

CLI commands:

noord tx poss signal ...

noord query poss stats

Web dashboard integration.

These pieces will be introduced in later PoSS Logic steps and in:

Phase 4 (Implementation),

Phase 6 (Genesis Pack & Communication),

with a strong focus on:

Legal Light CH compliance,

full transparency for curators and participants,

deterministic and auditable behaviour.
