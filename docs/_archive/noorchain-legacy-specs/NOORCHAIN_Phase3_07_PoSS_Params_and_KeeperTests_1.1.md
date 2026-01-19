NOORCHAIN_Phase3_07_PoSS_Params_Keeper_Tests_1.1.md
NOORCHAIN 1.0 — PoSS Parameters, Keeper Behaviour & Test Coverage

Version: 1.1
Status: Official Phase 3 Document
Last Updated: 2025-12-03

1. Scope

This document describes the current state of the PoSS Parameters, Keeper logic and unit tests in the x/noorsignal module.
It complements:

NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md

NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md

NOORCHAIN_Phase3_04_PoSS_FullSpecification_1.1.md

This file is part of Phase 3 Documentation and represents the authoritative reference before Phase 4 integration and Testnet activation.

2. PoSS Parameters (Params) — On-Chain Configuration
2.1 Params Structure

Defined in:

x/noorsignal/types/params.go


The PoSS module exposes the following configurable fields:

PoSSEnabled bool
Master economic switch (default: OFF).

Daily limits

MaxSignalsPerDay

MaxSignalsPerCuratorPerDay

Rewards

BaseReward sdk.Coin

MaxRewardPerDay sdk.Coin

Weights

WeightMicroDonation

WeightParticipation

WeightContent

WeightCCN

Reserve & Halving

PoSSReserveDenom (must be "unur")

HalvingPeriodBlocks (placeholder for 8-year halving)

2.2 DefaultParams (Safe Mode — Economic OFF)

DefaultParams() returns:

PoSSEnabled = false

MaxSignalsPerDay = 20

MaxSignalsPerCuratorPerDay = 100

BaseReward = 1 unur

MaxRewardPerDay = 100 unur

Weights:

Micro = 5, Participation = 2, Content = 3, CCN = 1

PoSSReserveDenom = "unur"

HalvingPeriodBlocks = 0

Consequences:
✔ PoSS is fully wired
✔ All reward computations return "0unur"
✔ No economic effect occurs on any signal

2.3 ParamSet & KeyTable

PoSS Params form a real ParamSet, allowing:

persistent storage in the x/params Subspace

future governance updates via x/gov

Generated keys:

KeyPoSSEnabled

KeyMaxSignalsPerDay

KeyMaxSignalsPerCuratorPerDay

KeyMaxRewardPerDay

KeyBaseReward

KeyWeightMicroDonation

KeyWeightParticipation

KeyWeightContent

KeyWeightCCN

KeyPoSSReserveDenom

KeyHalvingPeriodBlocks

2.4 Validation Rules

Params.Validate() enforces:

Non-empty denom

BaseReward & MaxRewardPerDay ≥ 0

All weights > 0

MaxSignalsPerDay > 0

MaxSignalsPerCuratorPerDay > 0

HalvingPeriodBlocks = 0 allowed

Denom must match "unur"

If validation fails → defaults are returned.

3. PoSS Keeper Behaviour

Located in:

x/noorsignal/keeper/keeper.go

3.1 Keeper Structure
type Keeper struct {
    cdc        codec.Codec
    storeKey   storetypes.StoreKey
    paramSpace paramstypes.Subspace
}


The Subspace automatically receives the PoSS KeyTable if missing.

3.2 Genesis Handling

Helper methods:

InitGenesis(ctx, gs)

ExportGenesis(ctx)

Genesis fields managed:

TotalSignals uint64

TotalMinted string (unur)

3.3 Parameter Handling
SetParams(ctx, params)

Validates

Saves into Subspace

GetParams(ctx)

If Subspace empty → loads DefaultParams() and persists them

If corrupted → returns defaults

Otherwise → returns stored values

This makes PoSS safe for Testnet even on empty stores.

3.4 Reward Helper
ComputeSignalRewardForBlock(ctx, signalType)


Steps:

Load Params

Read block height

Compute:

base reward

weight

halving factor

structural 70/30 split

If PoSSEnabled = false → return (0unur, 0unur)

3.5 Daily Counters

Per-address per-day counters:

GetDailySignalsCount

SetDailySignalsCount

IncrementDailySignalsCount

Used for anti-abuse (MaxSignalsPerDay).

Stored using:

DailyCounterKey(address, date)

3.6 Pending Mint Queue (Planning Only)

RecordPendingMint(...) stores a JSON entry describing:

block height

timestamp

participant & curator

signal type

theoretical rewards

No coins are minted.
This is purely diagnostic and for future wiring.

3.7 Internal Pipeline: ProcessSignalInternal

Actions:

Compute reward (theoretical)

Increment daily counter

Record PendingMint entry

Update:

TotalSignals++

TotalMinted += participant + curator

Return theoretical rewards

Not implemented yet:

enforcement of daily limits

enforcement of curator limits

real minting / transfers

4. Unit Testing Coverage
4.1 Reward Tests (x/noorsignal/types/rewards_test.go)

Covers:

PoSSEnabled = false → reward = 0/0

Weight scaling logic

Structural 70/30 split

Example:

BaseReward = 100
Weight = 5 → total = 500
→ participant = 350
→ curator = 150

Ensures no rounding errors.

4.2 Keeper Tests (x/noorsignal/keeper/keeper_test.go)

A full in-memory Cosmos environment is created:

TMDB (NewMemDB)

CommitMultiStore

x/noorsignal and x/params stores

ParamSubspace

Full Keeper

Test 1 — Default Params are persisted

GetParams(ctx) on an empty store:

✔ returns defaults
✔ persists them in Subspace

Test 2 — SetParams / GetParams round-trip

Custom params are stored then retrieved → must match.

Test 3 — ProcessSignalInternal

For PoSSEnabled = true:

participant counter increments

TotalSignals increments

TotalMinted increments exactly by (participant + curator reward)

No panic, no undefined behaviour

5. Impact on Testnet & Phase Progress

With this implementation:

✔ PoSS Params are now fully on-chain
✔ DefaultParams are automatically persisted
✔ Keeper logic is safe, deterministic and test-covered
✔ PoSS remains economically OFF by default
✔ Internal pipeline is validated

Phase 4 status:

4.3 (PoSS Types) → 100%

4.4 (PoSS Keeper) → 80–85%

4.6 (Tests & Build) → Significant progress

6. Next Steps (later phases)

Not executed during Phase 3:

Integrate ParamSubspace into app wiring

Add real PoSS reserve and mint logic

Enforce daily limits in Keeper (reject or accept-without-reward policy)

Expose PoSS gRPC services and CLI commands

Enable PoSS via governance (PoSSEnabled = true)

Testnet simulation of minting and rewards

7. Summary

This document defines the canonical state of:

PoSS Parameters

PoSS Keeper logic

Unit test coverage and behaviour

Test readiness

It serves as the internal reference for Phase 3, Phase 4, and Testnet behaviour.
