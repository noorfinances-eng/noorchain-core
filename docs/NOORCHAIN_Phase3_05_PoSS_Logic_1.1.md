# NOORCHAIN — PoSS Logic (End-to-End) v1.1

This document describes the *end-to-end* behaviour of the PoSS (Proof of Signal Social) engine in NOORCHAIN 1.0, from the moment a signal is emitted to the final reward distribution.

It is a **Phase 3 (Docs)** specification:
- No actual minting logic is live yet.
- All behaviours described here will be implemented in Phase 4 (Implementation) and refined by governance.

---

## 1. High-level flow of a PoSS signal

When PoSS is fully implemented, each signal will follow this logical pipeline:

1. **User emits a signal**  
   - Signal types (at least):  
     - `MICRO_DONATION`  
     - `PARTICIPATION` (QR / event)  
     - `CONTENT` (curated content)  
     - `CCN` (Curated Community Network / social propagation)

2. **Curator validates the signal**  
   - A Curator (Bronze / Silver / Gold) receives or reviews the signal.  
   - The Curator confirms that the signal is legitimate (no spam, no abuse).

3. **PoSS module receives a PoSS message**  
   - In Phase 4, this will be a real `Msg` handled by `x/noorsignal`.  
   - The message will include at least:
     - participant address
     - curator address
     - signal type
     - optional metadata (amount, reference, etc.)

4. **Daily anti-abuse checks**  
   - The Keeper will:
     - read the **participant daily counter** (signals today),
     - read the **curator daily counter** (validations today),
     - compare them to:
       - `MaxSignalsPerDay`
       - `MaxSignalsPerCuratorPerDay`
     - if limits are exceeded:
       - **the signal is recorded but no reward is paid** (or the tx is rejected, depending on the final policy).

5. **Reward computation**  
   - The Keeper reads the PoSS params:
     - `BaseReward` (in `unur`)
     - weights:
       - `WeightMicroDonation`
       - `WeightParticipation`
       - `WeightContent`
       - `WeightCCN`
     - halving schedule via `HalvingPeriodBlocks`
     - PoSS reserve denom: `PoSSReserveDenom = "unur"`
   - The Keeper calls the pure helper:
     - `ComputeSignalReward(params, signalType, blockHeight)`
   - This function:
     - applies the **signal weight**,
     - applies the **halving factor** (every 8 years in block time),
     - applies the **70 % / 30 % split**:
       - 70 % → participant
       - 30 % → curator.

6. **Final checks & minting (future Phase 4)**  
   - The Keeper will:
     - check that the PoSS reserve is sufficient (if we implement a hard reserve),
     - update:
       - `TotalSignals`
       - `TotalMinted`
     - mint or transfer the computed NUR:
       - to the participant address,
       - to the curator address.
   - All state changes will be **fully auditable** in the store and in events.

7. **Events & transparency**  
   - Each successful PoSS operation will emit events:
     - `"noorsignal_signal_created"` with:
       - participant
       - curator
       - signal type
       - raw reward
       - halving era
     - These events will be visible in explorers and indexers.

---

## 2. Daily limits behaviour

Even though PoSS limits are not hard-coded yet, the *intended* behaviour is:

- **Participant:** cannot emit more than `MaxSignalsPerDay` eligible signals per day.
  - After the limit is reached:
    - additional signals are either rejected or accepted with **zero reward**.
- **Curator:** cannot validate more than `MaxSignalsPerCuratorPerDay` rewarded signals per day.
  - Beyond that, the curator can still see signals, but:
    - they will not receive new PoSS rewards.

The exact policy (reject vs. accept-without-reward) will be finalised in a later PoSS Logic step, but the infrastructure (counters, params) is already prepared.

---

## 3. Halving and long-term issuance

PoSS Logic keeps the **fixed cap** of `299,792,458 NUR` intact.

- The PoSS module does **not** change the total cap.
- It only controls the *tempo* of social mining via:
  - `BaseReward`
  - signal weights
  - `HalvingPeriodBlocks`
  - daily limits (`MaxSignalsPerDay`, `MaxSignalsPerCuratorPerDay`)
  - and `PoSSEnabled`.

Every halving period:
- the **effective reward per signal** is reduced,
- the system becomes more conservative,
- but the 70/30 split always remains the same.

---

## 4. Invariants and safety rules

The following rules MUST always hold in implementation:

1. `PoSSReserveDenom` is always `"unur"`.
2. PoSS **never mints above the global cap** defined at genesis.
3. Daily counters are correctly reset and enforced:
   - participant per day,
   - curator per day.
4. The **reward split is structurally fixed**:
   - `ParticipantShare = 70`
   - `CuratorShare = 30`
5. Parameters (BaseReward, limits, halving) are:
   - adjustable by governance,
   - transparent and documented,
   - never used to promise any “yield” or “APR”.

---

## 5. Current implementation status (end of PoSS Logic 15)

As of the end of PoSS Logic 15:

- `x/noorsignal/types` already defines:
  - `GenesisState` + `DefaultGenesis` + `ValidateGenesis`
  - `Params` + `DefaultParams` + `Validate`
  - daily counter helpers and keys
  - pure reward helpers with 70/30 split and halving input
- `x/noorsignal/keeper` already defines:
  - a minimal `Keeper` with:
    - KVStore access,
    - per-address, per-day counters,
    - a `GetParams(ctx)` (currently returns `DefaultParams()`),
    - `ComputeSignalRewardForBlock(...)` as a thin wrapper on helpers.
- The **actual PoSS transaction type** and **state updates** (TotalSignals, TotalMinted, balances) are **not yet implemented**.
- PoSS remains **disabled by default** (`PoSSEnabled = false`).

This document defines the target behaviour so that Phase 4 can implement it step by step without ambiguity.
