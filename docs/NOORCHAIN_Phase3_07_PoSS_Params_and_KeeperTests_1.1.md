# NOORCHAIN 1.0 — PoSS Params, Keeper & Tests (v1.1)

**Scope**  
This document updates Phase 3 with the current state of the PoSS parameters, keeper logic and unit tests in `noorchain-core`.  
It completes:

- `NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md`
- `NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md`

---

## 1. PoSS Params on-chain (x/params Subspace)

### 1.1. Params structure

The PoSS parameters are defined in:

- `x/noorsignal/types/params.go`

The `Params` struct currently includes:

- `PoSSEnabled bool`  
  Master switch for PoSS rewards (economic ON/OFF).

- Limits:
  - `MaxSignalsPerDay uint64`
  - `MaxSignalsPerCuratorPerDay uint64`

- Rewards:
  - `MaxRewardPerDay sdk.Coin`
  - `BaseReward sdk.Coin`

- Weights:
  - `WeightMicroDonation uint32`
  - `WeightParticipation uint32`
  - `WeightContent uint32`
  - `WeightCCN uint32`

- Reserve & halving:
  - `PoSSReserveDenom string` (always `"unur"`)
  - `HalvingPeriodBlocks uint64` (placeholder for an 8-year halving period)

### 1.2. DefaultParams (economic OFF, safe mode)

`DefaultParams()` returns a safe configuration:

- `PoSSEnabled = false`
- `MaxSignalsPerDay = 20`
- `MaxSignalsPerCuratorPerDay = 100`
- `BaseReward = 1 unur`
- `MaxRewardPerDay = 100 unur`
- Weights:
  - MicroDonation = 5
  - Participation = 2
  - Content = 3
  - CCN = 1
- `PoSSReserveDenom = "unur"`
- `HalvingPeriodBlocks = 0` (not configured yet)

**Important**  
With these defaults, PoSS is **fully wired** but **economically OFF**.  
The reward helpers will return `0/0` as long as `PoSSEnabled = false`.

### 1.3. ParamSet & KeyTable

PoSS Params are now a **real ParamSet**:

- `ParamKeyTable()` defines a `KeyTable` with all PoSS fields.
- Each field has an associated parameter key:
  - `KeyPoSSEnabled`
  - `KeyMaxSignalsPerDay`
  - `KeyMaxSignalsPerCuratorPerDay`
  - `KeyMaxRewardPerDay`
  - `KeyBaseReward`
  - `KeyWeightMicroDonation`
  - `KeyWeightParticipation`
  - `KeyWeightContent`
  - `KeyWeightCCN`
  - `KeyPoSSReserveDenom`
  - `KeyHalvingPeriodBlocks`

This allows PoSS parameters to be stored and updated **on-chain** via `x/params` and, later, via governance (`x/gov`).

### 1.4. Validation rules

`Params.Validate()` enforces basic coherence:

- `PoSSReserveDenom` must not be empty.
- `MaxSignalsPerDay > 0`.
- `MaxSignalsPerCuratorPerDay > 0`.
- `BaseReward` and `MaxRewardPerDay`:
  - Same denom as `PoSSReserveDenom`.
  - Non-negative amounts.
- All weights (`Micro`, `Participation`, `Content`, `CCN`) must be `> 0`.
- `HalvingPeriodBlocks = 0` is allowed (means “not configured yet”).

---

## 2. PoSS Keeper behaviour (Params, Genesis, Pipeline)

The PoSS keeper lives in:

- `x/noorsignal/keeper/keeper.go`

### 2.1. Keeper fields

```go
type Keeper struct {
    cdc        codec.Codec
    storeKey   storetypes.StoreKey
    paramSpace paramstypes.Subspace
}
cdc : codec for encoding/decoding state.

storeKey : access to the KVStore of x/noorsignal.

paramSpace : Subspace used to store PoSS Params via x/params.

During keeper creation:

The Subspace is ensured to have the PoSS KeyTable:

If no KeyTable is present, WithKeyTable(ParamKeyTable()) is applied.

2.2. Genesis state helpers
The keeper handles a minimal PoSS GenesisState:

getGenesisState(ctx):

Reads JSON under KeyGenesisState.

Returns DefaultGenesis() if empty.

setGenesisState(ctx, gs):

Validates the state via ValidateGenesis.

Stores JSON under KeyGenesisState.

Public methods:

InitGenesis(ctx, gs):

Validates & persists the initial state.

ExportGenesis(ctx):

Reads the current state and returns it as a GenesisState.

The genesis fields currently used are:

TotalSignals uint64

TotalMinted string (in unur)

2.3. Params management (Subspace)
The keeper exposes two methods for PoSS Params:

SetParams
go
Copier le code
func (k Keeper) SetParams(ctx sdk.Context, params noorsignaltypes.Params)
Validates the params (params.Validate()).

Stores them in the Subspace via SetParamSet.

GetParams
go
Copier le code
func (k Keeper) GetParams(ctx sdk.Context) noorsignaltypes.Params
Behaviour:

First run / empty store

The keeper checks if the Subspace has at least one key:

if !k.paramSpace.Has(ctx, KeyPoSSReserveDenom) { ... }

If there is no stored value:

DefaultParams() is loaded.

SetParams() is called once to persist them.

The defaults are returned.

Subsequent runs

GetParamSet is used to read the PoSS Params from the Subspace.

If the loaded Params fail Validate():

A safe fallback is applied: DefaultParams() is returned.

This pattern ensures:

No panic on an empty Subspace.

A deterministic “safe-off” behaviour, even if Params are corrupted.

2.4. Reward helper (ComputeSignalRewardForBlock)
go
Copier le code
func (k Keeper) ComputeSignalRewardForBlock(
    ctx sdk.Context,
    signalType noorsignaltypes.SignalType,
) (sdk.Coin, sdk.Coin, error)
Steps:

params := k.GetParams(ctx)

height := ctx.BlockHeight()

Calls noorsignaltypes.ComputeSignalReward(params, signalType, height):

BaseReward * weight(signalType)

Halving via HalvingPeriodBlocks

Structural 70/30 split

If PoSSEnabled = false, the helper returns 0/0 (denom unur).

2.5. Daily counters
The keeper manages a per-address, per-day counter of PoSS signals:

GetDailySignalsCount(ctx, address, date string) uint32

SetDailySignalsCount(ctx, address, date string, count uint32)

IncrementDailySignalsCount(ctx, address, date string) uint32

Keys are built via:

noorsignaltypes.DailyCounterKey(address, date)

This is used later to enforce MaxSignalsPerDay and anti-abuse rules.

2.6. Pending mint queue
A planning-only “pending mint” queue is implemented:

go
Copier le code
func (k Keeper) RecordPendingMint(
    ctx sdk.Context,
    participantAddr string,
    curatorAddr string,
    signalType noorsignaltypes.SignalType,
    participantReward sdk.Coin,
    curatorReward sdk.Coin,
) error
Builds a PendingMint struct with:

Block height, timestamp

Participant & curator addresses

SignalType

ParticipantReward / CuratorReward

Serializes to JSON and stores under a simple debug-friendly key:

pending_mint:<height>:<participant>:<timestamp_nano>

Important
This does not mint any coins and does not move balances.
It is only a planning/logging mechanism for future PoSS activation.

2.7. Internal pipeline: ProcessSignalInternal
go
Copier le code
func (k Keeper) ProcessSignalInternal(
    ctx sdk.Context,
    participantAddr string,
    curatorAddr string,
    signalType noorsignaltypes.SignalType,
    date string,
) (sdk.Coin, sdk.Coin, error)
Steps:

Compute rewards at current block:

ComputeSignalRewardForBlock(ctx, signalType)

Increment participant daily counter for date.

Record a PendingMint entry.

Update global PoSS totals in Genesis:

TotalSignals++

TotalMinted += participantReward + curatorReward

Return (participantReward, curatorReward).

Limitations (by design at this stage)

No enforcement of daily limits (MaxSignalsPerDay, MaxSignalsPerCuratorPerDay) yet.

No real minting or Bank transfers.

PoSS remains economically OFF as long as PoSSEnabled = false.

3. Unit tests for PoSS logic (types + keeper)
3.1. Reward logic tests (types)
File:

x/noorsignal/types/rewards_test.go

Covers:

PoSSEnabled = false

ComputeSignalReward returns 0/0 with the correct denom (unur).

PoSSEnabled = true + weights + 70/30

BaseReward = 100 unur

Weight(micro_donation) = 5 → total = 500 unur

Split:

Participant: 350 unur

Curator: 150 unur

Ensures participant + curator == total.

These tests validate:

Correct handling of the master switch PoSSEnabled.

Correct proportional scaling by weights.

Correct structural 70/30 split without rounding drifts.

3.2. Keeper tests
File:

x/noorsignal/keeper/keeper_test.go

Includes an in-memory test setup:

setupKeeperAndContext(t):

In-memory TMDB (NewMemDB).

Cosmos CommitMultiStore.

Store keys for:

x/noorsignal KVStore.

x/params KVStore + transient store.

Minimal codec setup (like MakeEncodingConfig).

ParamsKeeper + PoSS Subspace.

Real Keeper instance.

sdk.Context.

Test 1 — TestGetParams_DefaultsStored
Calls k.GetParams(ctx) on an empty ParamSubspace.

Expects:

Defaults to be returned (PoSSReserveDenom, PoSSEnabled, limits).

Second call returns the same values, proving that defaults are persisted into the Subspace.

Test 2 — TestSetParams_RoundTrip
Builds custom Params from DefaultParams():

PoSSEnabled = true

MaxSignalsPerDay = 42

MaxSignalsPerCuratorPerDay = 200

Calls k.SetParams(ctx, custom).

Then k.GetParams(ctx) and checks that the values match.

Test 3 — TestProcessSignalInternal_UpdatesCountersAndGenesis
Forces a simple PoSS config:

PoSSEnabled = true

BaseReward = 10 unur

All weights = 1

Before calling:

Daily count = 0

TotalSignals = 0

TotalMinted = "0"

Calls ProcessSignalInternal once with:

participant, curator, SignalTypeMicroDonation, date = "2025-01-01"

Checks:

Rewards are positive (PoSS is ON).

Daily counter for participant/date increased by 1.

TotalSignals increased by 1.

TotalMinted increased exactly by participantReward + curatorReward.

4. Impact on Testnet & Phase 4 status
With these changes:

PoSS Params are stored on-chain via x/params.

Defaults are automatically written to the ParamSubspace at first use.

GetParams is safe even on a fresh testnet (no panic when the store is empty).

The keeper pipeline is covered by unit tests (params + internal signal processing).

PoSS remains economically OFF by design (default PoSSEnabled=false).

Phase 4 progress impact:

4.3 (PoSS Types) → ~100 %

4.4 (PoSS Keeper) → ~80–85 % (basic pipeline + Params + tests)

4.6 (Tests & Build) → progress increased (keeper tests added)

Next steps for Phase 4 / Phase 3 / Testnet:

Use this document as the reference for PoSS Params & Keeper behaviour.

Extend testnet documentation (Phase3_05 / Phase3_06) with:

“How to inspect PoSS Params via CLI / queries”

“How to simulate PoSS signals and read stats (TotalSignals / TotalMinted)”

Only later: connect PoSS to Bank (real mint from a PoSS reserve) and enforce limits.