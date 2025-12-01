# NOORCHAIN — PoSS Logic Overview (1–25)

Version: 1.1  
Module: `x/noorsignal` (PoSS — Proof of Signal Social)

This document summarizes the current state of the PoSS logic in NOORCHAIN
after PoSS Logic steps 1 to 25. It is meant for developers, auditors and
product/BD people who need a clear, high-level understanding of what is
already implemented in the core and what is still intentionally missing.

---

## 1. Scope of the PoSS Module (x/noorsignal)

The `x/noorsignal` module is responsible for:

- Tracking social signals (Proof of Light) as **PoSS signals**.
- Enforcing daily social limits (later).
- Computing rewards based on:
  - a **BaseReward** (in `unur`),
  - **per-signal-type weights**,
  - **halving** over time,
  - a **structural 70/30 split** between participant and curator.
- Recording **pending mints** (planned rewards) without touching real balances.
- Maintaining global PoSS counters:
  - `TotalSignals`
  - `TotalMinted` (planned PoSS minting in `unur`).
- Exposing a **read-only global stats view** (`PoSSStats`) for dashboards, CLI
  and monitoring tools.

The module is intentionally **conservative** at this stage:
no real minting, no bank movements, no governance wiring yet.

---

## 2. Core Data Structures

### 2.1 GenesisState

`x/noorsignal/types/genesis.go`

```go
type GenesisState struct {
    TotalSignals     uint64 `json:"total_signals" yaml:"total_signals"`
    TotalMinted      string `json:"total_minted" yaml:"total_minted"`
    MaxSignalsPerDay uint32 `json:"max_signals_per_day" yaml:"max_signals_per_day"`

    ParticipantShare uint32 `json:"participant_share" yaml:"participant_share"`
    CuratorShare     uint32 `json:"curator_share" yaml:"curator_share"`
}
Key decisions:

TotalSignals counts how many PoSS signals have been processed since genesis.

TotalMinted tracks the planned PoSS issuance in unur (as a string).

MaxSignalsPerDay is a first anti-abuse guardrail (per address).

The 70/30 split is structural and must always sum to 100:

ParticipantShare = 70

CuratorShare = 30

The genesis is serialized as plain JSON in the PoSS KVStore
(no protobuf, no gogo/proto dependency).

2.2 Params (PoSS behaviour)
x/noorsignal/types/params.go

go
Copier le code
type Params struct {
    PoSSEnabled bool `json:"poss_enabled" yaml:"poss_enabled"`

    MaxSignalsPerDay           uint64 `json:"max_signals_per_day" yaml:"max_signals_per_day"`
    MaxSignalsPerCuratorPerDay uint64 `json:"max_signals_per_curator_per_day" yaml:"max_signals_per_curator_per_day"`

    MaxRewardPerDay sdk.Coin `json:"max_reward_per_day" yaml:"max_reward_per_day"`
    BaseReward      sdk.Coin `json:"base_reward" yaml:"base_reward"`

    WeightMicroDonation uint32 `json:"weight_micro_donation" yaml:"weight_micro_donation"`
    WeightParticipation uint32 `json:"weight_participation" yaml:"weight_participation"`
    WeightContent       uint32 `json:"weight_content" yaml:"weight_content"`
    WeightCCN           uint32 `json:"weight_ccn" yaml:"weight_ccn"`

    PoSSReserveDenom    string `json:"poss_reserve_denom" yaml:"poss_reserve_denom"`
    HalvingPeriodBlocks uint64 `json:"halving_period_blocks" yaml:"halving_period_blocks"`
}
At this stage:

DefaultParams() returns a “safe off” configuration:

PoSSEnabled = false

conservative daily limits

rewards defined in unur but effectively disabled

The halving is configured structurally (parameter exists) but
the exact HalvingPeriodBlocks is still a placeholder.

Later, these params will be stored in a real ParamSubspace and controlled
by on-chain governance.

2.3 Rewards helpers
x/noorsignal/types/rewards.go

The PoSS reward pipeline (pure functions) is:

Base reward:

go
Copier le code
base := ComputeBaseReward(params, signalType)
= BaseReward * weight(signal_type)

Halving:

go
Copier le code
halved := ApplyHalving(params, height, base)
= divide by 2 every HalvingPeriodBlocks (never below 0).

70/30 split:

go
Copier le code
participant, curator, err := SplitReward70_30(halved)
High-level helper:

go
Copier le code
participant, curator, err := ComputeSignalReward(params, signalType, height)
If PoSSEnabled == false, ComputeSignalReward returns zero rewards
with the correct denom (no surprise minting).

2.4 Daily counters
Daily counters are per address, per date:

Types:

DailyCounter (documented in types/daily_counter.go)

Keys:

DailyCounterKey(address, date) (documented in types/daily_counter_key.go)

Keeper functions:

GetDailySignalsCount(ctx, address, date)

SetDailySignalsCount(ctx, address, date, count)

IncrementDailySignalsCount(ctx, address, date)

They are updated inside the internal signal pipeline but
daily limits are not yet enforced. This will be part of a later step.

3. Pending Mint Queue (planning only)
x/noorsignal/types/pending_mint.go (name may differ depending on repo)

The idea:

Every processed PoSS signal creates a PendingMint entry in the PoSS store.

This entry contains:

participant / curator addresses,

signal type,

planned rewards (participant + curator),

block height & timestamp.

The keeper exposes:

go
Copier le code
func (k Keeper) RecordPendingMint(
    ctx sdk.Context,
    participantAddr string,
    curatorAddr string,
    signalType SignalType,
    participantReward sdk.Coin,
    curatorReward sdk.Coin,
) error
Important: this does not mint coins.
It is a planning mechanism for a later minting phase, which can:

run in a separate handler,

be audited,

be rate-limited,

respect Legal Light constraints.

4. Internal Signal Pipeline (without real minting)
x/noorsignal/keeper/keeper.go

The internal, non-public pipeline is:

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

Compute rewards (base → halving → 70/30).

Increment participant daily counter for the given date.

Record a PendingMint entry.

Update GenesisState:

TotalSignals++

TotalMinted += participantReward + curatorReward

Return the two reward coins to the caller.

It does not:

enforce daily limits yet,

enforce curator limits,

move any Bank balances.

This is the “sandbox pipeline” to keep PoSS deterministic and auditable
before turning real money on.

5. Global PoSS Stats (PoSSStats)
PoSS Logic 24 introduced a simple, consolidated view:

x/noorsignal/types/stats.go

go
Copier le code
type PoSSStats struct {
    TotalSignals               uint64 `json:"total_signals" yaml:"total_signals"`
    TotalMinted                string `json:"total_minted" yaml:"total_minted"`
    PoSSEnabled                bool   `json:"poss_enabled" yaml:"poss_enabled"`
    MaxSignalsPerDay           uint64 `json:"max_signals_per_day" yaml:"max_signals_per_day"`
    MaxSignalsPerCuratorPerDay uint64 `json:"max_signals_per_curator_per_day" yaml:"max_signals_per_curator_per_day"`
    PoSSReserveDenom           string `json:"poss_reserve_denom" yaml:"poss_reserve_denom"`
}
Keeper helper:

go
Copier le code
func (k Keeper) GetGlobalStats(ctx sdk.Context) noorsignaltypes.PoSSStats
This function:

reads GenesisState from the KVStore,

reads PoSS Params (currently via DefaultParams()),

builds a single read-only snapshot.

5.1 Example JSON output (future query)
In a future noord query poss stats, a typical JSON payload could look like:

json
Copier le code
{
  "total_signals": 12345,
  "total_minted": "987654321",
  "poss_enabled": true,
  "max_signals_per_day": 20,
  "max_signals_per_curator_per_day": 100,
  "poss_reserve_denom": "unur"
}
Interpretation:

12345 PoSS signals processed since genesis.

987654321 unur planned to be minted through PoSS (70/30 split included).

PoSS rewards are currently enabled.

Each participant can emit up to 20 signals/day.

Each curator can validate up to 100 signals/day.

Rewards and reserves are denominated in unur.

6. What is Missing on Purpose (Post-Logic-25 Roadmap)
NOT implemented yet (by design):

Real minting of unur or movement of balances in x/bank.

Full enforcement of:

MaxSignalsPerDay

MaxSignalsPerCuratorPerDay

MaxRewardPerDay

Governance wiring for PoSS params (ParamSubspace + x/gov).

Public Msg/Query gRPC endpoints (proto files, Msg server, Query server).

CLI commands:

noord tx poss signal ...

noord query poss stats

Web dashboard integration.

These pieces will be introduced in later PoSS Logic steps and in
Phase 4 (Implementation) / Phase 6 (Genesis Pack & Communication),
with a strong focus on:

Legal Light CH compliance,

full transparency for curators and participants,

deterministic and auditable behaviour.

