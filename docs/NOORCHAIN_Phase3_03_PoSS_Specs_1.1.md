NOORCHAIN 1.0 — Phase 3  
## PoSS Module Specifications (x/noorsignal) — Version 1.1

**Scope of this document**

This document describes the *current* specifications of the PoSS module (`x/noorsignal`) as implemented in the Noorchain Core (Cosmos SDK v0.46.11 + Ethermint v0.22.0).  
It focuses on:

- What is already coded and wired in the app.
- How PoSS state is represented (genesis, params, counters).
- How daily limits and basic safety are enforced (design-level).
- Which parts are still TODO for later phases (rewards, halving, Bank links, etc.).

It is **not** a full economic paper, but a technical spec for Phase 3 + early Phase 4.

---

## 1. Module overview

- Module name: `noorsignal`
- Package path: `x/noorsignal`
- Store key: `noorsignal`
- Router / querier keys: `noorsignal` (reserved for later messages and queries)

The PoSS module is responsible for:

- Tracking social “signals” (Proof of Signal Social / Proof of Light).
- Enforcing **daily limits** per participant and per curator.
- Applying the **70/30 reward split** (later, in Logic phases).
- Staying aligned with the global NOORCHAIN rules:
  - Fixed cap: **299,792,458 NUR**.
  - Halving every **8 years**.
  - Legal Light CH: no hidden yield promises, transparent rules.

At this stage, the module is mostly a **skeleton with real state types**, focused on:

- Genesis state structure.
- Parameter structure.
- Daily counters for addresses.
- Wiring into the app (AppModule, keeper, BeginBlock).

---

## 2. Genesis state

File: `x/noorsignal/types/genesis.go`

```go
type GenesisState struct {
    // Total number of PoSS signals already processed on the chain.
    TotalSignals uint64 `json:"total_signals" yaml:"total_signals"`

    // Total NUR minted via PoSS, in the smallest unit (unur).
    TotalMinted string `json:"total_minted" yaml:"total_minted"`

    // Maximum number of PoSS signals allowed per address per day.
    MaxSignalsPerDay uint32 `json:"max_signals_per_day" yaml:"max_signals_per_day"`

    // Reward split (must always sum to 100).
    // Official NOORCHAIN rule:
    //   - 70% for the participant
    //   - 30% for the curator
    ParticipantShare uint32 `json:"participant_share" yaml:"participant_share"`
    CuratorShare     uint32 `json:"curator_share" yaml:"curator_share"`
}
2.1 DefaultGenesis
go
Copier le code
func DefaultGenesis() *GenesisState {
    return &GenesisState{
        TotalSignals:     0,
        TotalMinted:      "0",
        MaxSignalsPerDay: 50,
        ParticipantShare: 70,
        CuratorShare:     30,
    }
}
This matches the official rule:

70% of each PoSS reward goes to the participant.

30% goes to the curator.

A simple default daily limit of 50 signals per address is used as a starting point.

2.2 ValidateGenesis
The ValidateGenesis function checks:

Genesis is not nil.

participant_share + curator_share == 100.

total_minted is not empty (at least "0").

max_signals_per_day > 0.

If any of these conditions fails, genesis is considered invalid.

3. PoSS parameters (Params)
File: x/noorsignal/types/params.go

The Params struct encapsulates all tunable aspects of PoSS behavior:

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
3.1 Default parameters
Key defaults:

PoSSEnabled = false
→ PoSS is off by default. The chain can run without social mining until governance decides otherwise.

MaxSignalsPerDay = 20

MaxSignalsPerCuratorPerDay = 100

BaseReward = 1 unur

MaxRewardPerDay = 100 unur

Weights (relative importance):

MicroDonation: 5

Participation: 2

Content: 3

CCN: 1

PoSSReserveDenom = "unur"

HalvingPeriodBlocks = 0 (placeholder, to be set later according to real block time).

3.2 Params.Validate
The Validate() method enforces:

Non-empty PoSSReserveDenom.

MaxSignalsPerDay > 0, MaxSignalsPerCuratorPerDay > 0.

BaseReward and MaxRewardPerDay have the correct denom and non-negative amounts.

All weights > 0.

HalvingPeriodBlocks = 0 is accepted as “not yet configured”.

This keeps PoSS configuration coherent, while leaving the exact economic values adjustable by governance.

4. State and keys
Files:

x/noorsignal/types/keys.go

x/noorsignal/types/counter_keys.go

x/noorsignal/types/counters.go

4.1 Module keys
go
Copier le code
const (
    ModuleName   = "noorsignal"
    StoreKey     = ModuleName
    RouterKey    = ModuleName
    QuerierRoute = ModuleName
)

var (
    KeyLastResetDay                = []byte{0x01}
    KeyPrefixParticipantDailyCount = []byte{0x10}
    KeyPrefixCuratorDailyCount     = []byte{0x11}
)
KeyLastResetDay stores the last day (in Unix days) when all daily counters were reset.

KeyPrefixParticipantDailyCount and KeyPrefixCuratorDailyCount are prefixes for the KVStore entries.

4.2 Daily counters
The code defines an explicit model and helper:

go
Copier le code
type DailyCounter struct {
    Address string `json:"address" yaml:"address"`
    Date    string `json:"date" yaml:"date"`   // "YYYY-MM-DD"
    Signals uint32 `json:"signals" yaml:"signals"`
}

const DailyCounterPrefix = "daily_counter:"

func DailyCounterKey(address, date string) []byte {
    return []byte(DailyCounterPrefix + address + ":" + date)
}
This describes the conceptual shape of a per-address-per-day counter, although the actual storage for the first version uses simple prefixes and big-endian integers.

5. Keeper responsibilities
File: x/noorsignal/keeper/keeper.go

The current keeper is minimal but real:

Holds:

cdc codec.Codec

storeKey storetypes.StoreKey

Provides helpers to:

Access the PoSS KVStore.

Manage daily signal counts for a given address & date.

Key methods:

go
Copier le code
func (k Keeper) GetDailySignalsCount(ctx sdk.Context, address, date string) uint32
func (k Keeper) SetDailySignalsCount(ctx sdk.Context, address, date string, count uint32)
func (k Keeper) IncrementDailySignalsCount(ctx sdk.Context, address, date string) uint32
This is the foundation for enforcing MaxSignalsPerDay later in the PoSS logic.

6. AppModule and ABCI hooks
File: x/noorsignal/module.go

The AppModule:

Implements module.AppModule and module.AppModuleBasic.

Exposes a safe skeleton:

Name()

DefaultGenesis() / ValidateGenesis() (basic form for now)

InitGenesis() / ExportGenesis() (simple placeholders, to be upgraded)

BeginBlock() / EndBlock() (no heavy logic yet)

CLI methods (GetTxCmd, GetQueryCmd) returning nil for now.

The module is registered in app/app.go:

go
Copier le code
noorsignalAppModule := noorsignalmodule.NewAppModule(
    app.appCodec,
    app.NoorSignalKeeper,
)
And included in:

module.NewManager(...)

SetOrderInitGenesis(..., noorsignaltypes.ModuleName)

This means PoSS is a first-class module in the app, even if rewards and signals are not yet processed.

7. Current limitations & next steps
At this stage (PoSS Logic 8):

✅ Genesis state structure exists (with 70/30 rule and MaxSignalsPerDay).

✅ Params struct exists with all fields needed for future PoSS logic.

✅ Daily counters and store keys exist.

✅ PoSS module is wired into the app and can evolve without breaking the chain.

Still not implemented yet (reserved for later “PoSS Logic” steps):

Applying Params in real signal processing.

Reward calculation (BaseReward × weight, halving, MaxRewardPerDay).

Actual minting / transfer of unur from a PoSS reserve.

Full BeginBlock logic (daily reset + halving, stats, events).

Messages and queries (Tx/Query gRPC, CLI commands).

This is exactly what the following “PoSS Logic” phases will deliver.