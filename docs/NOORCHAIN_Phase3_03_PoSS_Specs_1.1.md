NOORCHAIN 1.0 — Phase 3
PoSS Module Specifications (x/noorsignal) — Version 1.1
0. Scope of this document

This document describes the current specifications of the PoSS module (x/noorsignal) as implemented in the Noorchain Core (Cosmos SDK v0.46.11 + Ethermint v0.22.0).

It focuses on:

what is already coded and wired in the app,

how PoSS state is represented (genesis, params, counters),

how daily limits and basic safety are enforced (design-level),

which parts are still TODO for later phases (rewards, halving, Bank links, etc.).

It is not a full economic paper, but a technical specification for Phase 3 + early Phase 4.

1. Module overview

Module name: noorsignal

Package path: x/noorsignal

Store key: noorsignal

Router / querier keys: noorsignal (reserved for future messages and queries)

The PoSS module is responsible for:

tracking social “signals” (Proof of Signal Social / Proof of Light),

enforcing daily limits per participant and per curator,

applying the 70/30 reward split (later PoSS Logic phases),

staying aligned with the global NOORCHAIN rules:

fixed cap: 299,792,458 NUR,

halving every 8 years,

Legal Light CH: no hidden yield promises, transparent rules.

At this stage, the module is essentially a skeleton with real state types, focused on:

genesis state structure,

parameter structure,

daily counters per address,

wiring into the app (AppModule, keeper, BeginBlock).

2. Genesis state

File: x/noorsignal/types/genesis.go

The GenesisState structure contains:

TotalSignals (uint64)
Total number of PoSS signals already processed on-chain.

TotalMinted (string, amount in unur)
Total NUR minted via PoSS, expressed in the smallest unit (unur).

MaxSignalsPerDay (uint32)
Maximum number of PoSS signals allowed per address per day.

ParticipantShare (uint32)

CuratorShare (uint32)

The two share fields must always sum to 100.
The official NOORCHAIN rule is:

70% of each PoSS reward goes to the participant,

30% goes to the curator.

2.1 DefaultGenesis

The default genesis values are:

TotalSignals = 0

TotalMinted = "0"

MaxSignalsPerDay = 50

ParticipantShare = 70

CuratorShare = 30

This encodes the 70/30 split and a simple initial daily limit of 50 signals per address.

2.2 ValidateGenesis

The ValidateGenesis function enforces:

GenesisState is not nil,

ParticipantShare + CuratorShare == 100,

TotalMinted is non-empty (at least "0"),

MaxSignalsPerDay > 0.

If any of these checks fails, genesis is considered invalid and the chain must not start.

3. PoSS parameters (Params)

File: x/noorsignal/types/params.go

The Params structure encapsulates all tunable aspects of PoSS behaviour:

PoSSEnabled (bool)
Master switch controlling whether PoSS rewards are active.

MaxSignalsPerDay (uint64)
Maximum number of PoSS signals allowed per participant per day.

MaxSignalsPerCuratorPerDay (uint64)
Maximum number of signals a curator can validate per day.

MaxRewardPerDay (sdk.Coin)
Cap on the total PoSS reward per address per day (optional but reserved).

BaseReward (sdk.Coin)
Base unit of reward (in unur), multiplied by weights depending on signal type.

WeightMicroDonation (uint32)

WeightParticipation (uint32)

WeightContent (uint32)

WeightCCN (uint32)

Relative weights for different types of PoSS signals.

PoSSReserveDenom (string)
Denom used for PoSS rewards. Must be unur in NOORCHAIN.

HalvingPeriodBlocks (uint64)
Number of blocks between halving events (derived from “8 years” target).
A value of 0 is accepted as “not configured yet”.

3.1 Default parameters

The default parameter values are:

PoSSEnabled = false
PoSS is disabled by default; the chain can run without social mining until governance decides to enable it.

MaxSignalsPerDay = 20

MaxSignalsPerCuratorPerDay = 100

BaseReward = 1 unur

MaxRewardPerDay = 100 unur

Weights (relative importance):

WeightMicroDonation = 5

WeightParticipation = 2

WeightContent = 3

WeightCCN = 1

PoSSReserveDenom = "unur"

HalvingPeriodBlocks = 0 (placeholder; to be set later according to real block time).

These defaults are not final economics, but a sane technical configuration that can be safely updated by governance.

3.2 Params.Validate

The Validate() method guarantees:

PoSSReserveDenom is non-empty,

MaxSignalsPerDay > 0 and MaxSignalsPerCuratorPerDay > 0,

BaseReward and MaxRewardPerDay:

use the correct denom (unur),

have non-negative amounts,

all weights (WeightMicroDonation, WeightParticipation, WeightContent, WeightCCN) are strictly > 0,

HalvingPeriodBlocks = 0 is allowed (interpreted as “not configured yet”).

This keeps PoSS configuration coherent, while leaving the exact economic values adjustable via governance.

4. State keys and counters

Files:

x/noorsignal/types/keys.go

x/noorsignal/types/counter_keys.go

x/noorsignal/types/counters.go

4.1 Module keys

The module defines:

ModuleName = "noorsignal"

StoreKey = "noorsignal"

RouterKey = "noorsignal"

QuerierRoute = "noorsignal"

And the following store prefixes:

KeyLastResetDay
Stores the last day (in Unix days) when daily counters were reset.

KeyPrefixParticipantDailyCount
Prefix for participant daily counters.

KeyPrefixCuratorDailyCount
Prefix for curator daily counters.

4.2 Daily counters model

Conceptually, the code defines a daily counter structure:

Address (string)

Date (string, formatted "YYYY-MM-DD")

Signals (uint32)

and a textual key prefix such as:

DailyCounterPrefix = "daily_counter:"

The final storage in v1 uses simple prefixes + big-endian integers under the participant / curator prefixes, but the conceptual model remains “per address, per date, number of signals”.

This is the basis for enforcing:

MaxSignalsPerDay (per participant),

MaxSignalsPerCuratorPerDay (per curator).

5. Keeper responsibilities

File: x/noorsignal/keeper/keeper.go

The current Keeper is minimal but real.

It owns:

a codec (cdc),

a storeKey (PoSS KV store).

It provides helper methods to:

access the PoSS store,

read the daily signal count for a given (address, date),

update or increment this count.

Typical responsibilities:

GetDailySignalsCount(ctx, address, date)

SetDailySignalsCount(ctx, address, date, count)

IncrementDailySignalsCount(ctx, address, date)

These helpers are the foundation for enforcing MaxSignalsPerDay and MaxSignalsPerCuratorPerDay later in the PoSS Logic.

6. AppModule and ABCI hooks

File: x/noorsignal/module.go

The PoSS module implements a standard Cosmos SDK AppModule:

Name() returns "noorsignal",

DefaultGenesis() returns a valid GenesisState,

ValidateGenesis() checks the genesis invariants,

InitGenesis() and ExportGenesis() currently use simple placeholder logic,

BeginBlock() and EndBlock() exist but do not yet execute full PoSS logic,

CLI-related methods (GetTxCmd, GetQueryCmd) currently return nil.

The module is registered in app/app.go and included in the main ModuleManager, including:

module registration,

SetOrderInitGenesis(..., noorsignaltypes.ModuleName),

integration in the app’s lifecycle.

This means the PoSS module is a first-class citizen of the application, even though most of the economic logic is still disabled.

7. Current limitations and next steps

At the current stage (PoSS Logic around step 8), the following are already in place:

✅ GenesisState structure exists (with 70/30 rule and MaxSignalsPerDay).

✅ Params structure exists with all required fields for future PoSS logic.

✅ Daily counter keys and helper functions exist.

✅ noorsignal is integrated as a module in noorchain-core (wired in the app).

However, the real PoSS economic behaviour is not implemented yet. The following items are still TODO and are explicitly reserved for later PoSS Logic phases:

Applying Params in real signal processing.

Reward calculation (BaseReward × weight, halving, MaxRewardPerDay).

Actual minting / transfer of unur from a PoSS reserve (BankKeeper integration).

Full BeginBlock / EndBlock logic:

daily reset,

halving counters,

stats and events emission.

Full Msg / Query API:

Tx messages (e.g. MsgCreateSignal),

gRPC queries,

CLI commands.

These parts will be delivered in the upcoming PoSS Logic phases, building on the present skeleton without breaking the chain.

8. Summary (executive)

The PoSS module x/noorsignal is present in the core app, with real genesis, params and counters.

Daily limits, 70/30 split, and Legal Light constraints are already encoded in structures and genesis.

PoSS is disabled by default (PoSSEnabled = false), and no minting occurs yet.

Future phases will add:

reward logic,

PoSS-enabled minting,

messages and queries,

full halving mechanism.

This document is the official technical reference for the PoSS module at Phase 3 + early Phase 4.
