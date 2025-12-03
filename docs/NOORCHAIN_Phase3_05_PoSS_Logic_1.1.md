# NOORCHAIN — PoSS Logic (End-to-End) v1.1  
Phase 3.05 — Target Behaviour Specification  
Last Updated: 2025-12-03

This document describes the **end-to-end behaviour** of the PoSS (Proof of Signal Social) engine in NOORCHAIN 1.0, from the moment a signal is emitted to the final reward distribution.

It is a **Phase 3 (Documentation) specification**:

- No live minting or balance transfers are active yet.
- PoSS is currently **structurally integrated but economically disabled**.
- All behaviours described here are the target model for Mainnet.
- The implementation has been prepared during Phase 4 and will be activated later by governance.

This file complements:

- `NOORCHAIN_Phase3_04_PoSS_FullSpecification_1.1.md`
- all PoSS-related Phase 3 & Phase 4 documents.

---

## 1. High-level lifecycle of a PoSS signal

When PoSS is fully enabled, each signal will follow this logical pipeline:

1. **User emits a signal**  
   - Supported signal types (as per PoSS Full Specification):  
     - `MICRO_DONATION`  
     - `PARTICIPATION` (QR / event / activity)  
     - `CONTENT` (certified content)  
     - `CCN` (Certified Content Noorchain – high-value content)

2. **Curator validates the signal**  
   - A Curator (Bronze / Silver / Gold) receives or reviews the signal.  
   - The Curator confirms that the signal is legitimate (no spam, no abuse, consistent with rules).  

3. **PoSS module receives a PoSS message**  
   - In the live chain, this is a `MsgCreateSignal` (or equivalent) handled by `x/noorsignal`.  
   - The message always includes at least:
     - participant address  
     - curator address  
     - signal type  
     - optional metadata (amount, reference, context, hashes…)

4. **Daily anti-abuse checks**  
   - The PoSS Keeper:
     - reads the **participant daily counter** (signals today),
     - reads the **curator daily counter** (validations today),
     - compares both to:
       - `MaxSignalsPerDay`
       - `MaxSignalsPerCuratorPerDay`
   - If limits are exceeded:
     - the default target policy is:
       - the signal **may be recorded** for history,
       - but **no reward is paid** beyond the daily limits.
     - Alternatively, governance may choose a stricter policy:
       - reject the transaction when limits are exceeded.
   - The infrastructure (counters + params) is already designed to support both policies.

5. **Reward computation**  
   - The Keeper reads PoSS parameters:
     - `BaseReward` (in `unur`)
     - weights:
       - `WeightMicroDonation`
       - `WeightParticipation`
       - `WeightContent`
       - `WeightCCN`
     - halving schedule via `HalvingPeriodBlocks`
     - PoSS reserve denom: `PoSSReserveDenom = "unur"`
   - The Keeper calls a pure helper function, such as:  
     `ComputeSignalReward(params, signalType, blockHeight)`
   - This function:
     - applies the **signal weight**,
     - applies the **halving factor** (every 8 years, approximated in blocks),
     - computes a raw reward,
     - applies the **70% / 30% structural split**:
       - 70% → participant
       - 30% → curator.

6. **Final checks & minting (target behaviour)**  
   When PoSS is effectively enabled:

   - The Keeper:
     - checks that PoSS is enabled: `PoSSEnabled == true`,
     - ensures that global totals (`TotalMinted`) stay **within the supply cap**,
     - optionally checks that the dedicated PoSS Reserve is sufficient (if implemented as a hard pool),
     - updates:
       - `TotalSignals`
       - `TotalMinted`
       - per-address reward history (if configured).
   - Then, the Keeper:
     - mints or transfers the computed NUR:
       - to the participant address,
       - to the curator address.
   - All state changes are **fully auditable** through:
     - the store content,
     - associated events.

7. **Events & transparency**  
   - Each successful PoSS operation emits events, such as:
     - `"noorsignal_signal_created"`
     - `"noorsignal_signal_validated"`
     - `"noorsignal_reward_distributed"`
   - Typical attributes include:
     - participant
     - curator
     - signal type
     - raw reward
     - halving era / epoch info
   - These events are visible to:
     - explorers,
     - indexers,
     - dashboards,
     - governance tools.

---

## 2. Daily limits behaviour

Daily limits are a central component of PoSS Logic and anti-abuse.

### 2.1 Participant limits

- A participant **cannot receive more than `MaxSignalsPerDay` rewarded signals per day**.
- After the limit is reached:
  - target behaviour:
    - signals can still be emitted and recorded for history / statistics,
    - but **they generate zero PoSS rewards** for that day.
  - alternative stricter behaviour (if governance decides):
    - signals beyond the limit are rejected at the message level.

### 2.2 Curator limits

- A curator **cannot validate more than `MaxSignalsPerCuratorPerDay` rewarded signals per day**.
- After the limit is reached:
  - the curator may still validate signals as a social/ethical action,
  - but:
    - **they do not receive additional PoSS rewards** beyond the limit.

### 2.3 Reset and accounting

- Daily counters are keyed by:
  - address, and  
  - day (via block time or a derived epoch index).
- At each new epoch (1 day), counters are reset or replaced by new entries.

The exact reject vs. accept-without-reward policy remains configurable, but **the storage and logic for both options are already anticipated**.

---

## 3. Halving and long-term issuance logic

PoSS Logic is constrained by a **fixed global cap** of `299,792,458 NUR`.

- PoSS does **not** create inflation.  
- All rewards are derived from a pre-allocated **PoSS Reserve** (80% of total supply).

The PoSS engine only controls the **tempo** of distribution via:

- `BaseReward`
- signal weights
- `HalvingPeriodBlocks` (approx. 8 years)
- daily limits:
  - `MaxSignalsPerDay`
  - `MaxSignalsPerCuratorPerDay`
- the global switch: `PoSSEnabled`.

### 3.1 Halving behaviour

Every halving period (~8 years):

- the effective reward per signal decreases,
- the PoSS Reserve emission pace slows down,
- but the **70/30 split remains untouched**.

Pseudo-formula for a given block height:

```text
era              = floor(blockHeight / HalvingPeriodBlocks)
effectiveReward  = BaseReward / (2 ^ era)
This effectiveReward is then modulated by:

the signal type weight,

possible global daily caps.

4. Invariants and safety rules

The following invariants must always hold in the final implementation:

Denomination

PoSSReserveDenom is always "unur" (the base NUR unit).

Supply cap preservation

PoSS never mints above the global cap set at genesis:

TotalSupply <= 299,792,458 NUR at all times.

Daily counters

Participant daily signals and curator daily validations must be:

correctly updated,

correctly reset or rotated at the epoch boundary,

consistently enforced.

Structural reward split

Reward split is hard-anchored:

ParticipantShare = 70

CuratorShare = 30

This split must not be modifiable by governance.

Governance and parameters

BaseReward, daily limits, and halving intervals are:

set at genesis (Phase 3/4),

adjustable via on-chain governance,

fully transparent.

They must never be used to advertise or guarantee:

fixed yield,

APR,

ROI,

or any financial performance.

Legal Light CH compliance

PoSS must remain compatible with the Swiss Legal Light framework:

no promise of financial return,

reward distribution based on social signals, not investment,

transparent and capped supply,

no custody of third-party funds inside PoSS.

5. Current implementation status (end of PoSS Logic 15 / Phase 4)

At the end of Phase 4 (Implementation core completed), the PoSS engine has the following status:

5.1 Types and genesis

The package x/noorsignal/types defines:

GenesisState + DefaultGenesis + ValidateGenesis

Params + DefaultParams + Validate

keys and helpers for:

daily counters,

global totals (TotalSignals, TotalMinted),

pure reward helpers implementing:

base reward input,

signal weight,

halving factor,

70/30 structural split.

5.2 Keeper and core logic

The package x/noorsignal/keeper defines:

a real Keeper with:

access to the KV store,

per-address, per-day counters,

GetParams(ctx) and parameter accessors,

internal helpers for computing block-level rewards
(e.g. ComputeSignalRewardForBlock(...) as a thin wrapper on pure helpers),

the ability to record signals and counters in the store,

a PoSSEnabled parameter (currently false by default).

5.3 Messages and module integration

A MsgCreateSignal (or equivalent) exists with its MsgServer implementation.

Signals can already be submitted and counted in the Testnet-style environment.

The module x/noorsignal is:

wired in app/app.go,

part of ModuleManager,

included in:

InitGenesis,

ExportGenesis,

BeginBlock / EndBlock logic as needed.

5.4 What is not active yet

No real minting or bank transfers are executed from PoSS:

PoSSEnabled = false in default parameters,

reward calculations are theoretical and can be inspected in logs / queries,

no actual NUR distribution occurs from PoSS in the current Testnet state.

The final wiring between:

PoSS reward helpers and

Bank / Mint modules
will be switched on only after governance and legal validation.

6. Target activation path (from Phase 5 onward)

To move from the current “counting only” state to a fully active PoSS, the following steps are expected:

Legal & Governance validation (Phase 5)

Confirm that reward levels remain symbolic and non-financial.

Validate Legal Light CH classification.

Freeze structural invariants (cap, 70/30, 8-year halving).

Genesis parameter finalization (Phase 6 / Genesis Pack)

Set:

initial BaseReward,

MaxSignalsPerDay,

MaxSignalsPerCuratorPerDay,

HalvingPeriodBlocks,

PoSSEnabled = false at genesis.

Document all values in the Genesis Pack.

Testnet 1.0 PoSS dry-run (Phase 7–8)

Enable PoSS counting only (no minting).

Observe volumes and patterns.

Adjust parameters via governance if needed.

Governance activation vote

When the network is ready:

a governance proposal can switch PoSSEnabled from false to true,

and connect PoSS rewards to real NUR movements (within the cap).

7. Summary

This document captures the end-to-end, target PoSS Logic for NOORCHAIN 1.0:

exact signal lifecycle,

daily anti-abuse behaviour,

halving and long-term emission logic,

structural safety invariants,

current implementation status (Phase 4 completed),

and the path toward safe, compliant activation.

It is the bridge document between:

conceptual specification (Phase3_04_PoSS_FullSpecification) and

real implementation (Phase 4, Testnet, Mainnet).
