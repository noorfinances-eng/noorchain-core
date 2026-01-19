# NOORCHAIN 1.0 — PoSS Module Status & Testnet Preparation (v1.1)

**Scope:**  
This document summarizes the current technical state of the `x/noorsignal` (PoSS) module in NOORCHAIN 1.0, and defines a first wave of targeted testnet scenarios.  
It is part of **Phase 3 — Documentation & Specifications** and relies on the previous documents:

- `NOORCHAIN_Phase3_01_Architecture_1.1.md`
- `NOORCHAIN_Phase3_02_Genesis_1.1.md`
- `NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md`
- `NOORCHAIN_Phase3_04_Economic_Model_1.1.md`

---

## 1. Current PoSS Technical State (x/noorsignal)

### 1.1. Module wiring

The PoSS module is implemented as `x/noorsignal` and wired into the core app:

- `app/app.go`
  - KV store key registered: `noorsignaltypes.StoreKey = "noorsignal"`.
  - `NoorSignalKeeper` created with:
    - `cdc codec.Codec`
    - `storeKey storetypes.StoreKey`
  - PoSS AppModule added to the `module.Manager`:
    - `noorsignalmodule.NewAppModule(app.appCodec, app.NoorSignalKeeper)`
  - PoSS included in InitGenesis order:
    - `..., evm, feemarket, noorsignal`

- `x/noorsignal/module.go`
  - `AppModuleBasic` and `AppModule` implemented.
  - Satisfies `module.AppModule` and `module.AppModuleBasic` interfaces (v0.46).
  - `InitGenesis` / `ExportGenesis` implemented using the `GenesisState`.
  - `BeginBlock` calls PoSS logic when needed (for future daily logic / stats).
  - No legacy querier / CLI commands wired yet (kept as `nil`).

**Important:**  
PoSS is **structurally integrated** into the chain, but there is **no real minting** and **no economic effect** yet.

---

### 1.2. PoSS Genesis State (genesis.go)

`x/noorsignal/types/genesis.go` defines a minimal but real `GenesisState`:

- `TotalSignals uint64`
  - Total number of PoSS signals processed by the chain.
  - Starts at `0` for a fresh network.
- `TotalMinted string`
  - Total NUR minted via PoSS in `unur` (as a string).
  - Starts at `"0"`.
- `MaxSignalsPerDay uint32`
  - First anti-abuse guardrail: maximum signals per address per day.
- `ParticipantShare uint32`
- `CuratorShare uint32`
  - MUST always sum to `100`.
  - Official rule: `70 %` participant / `30 %` curator.

`DefaultGenesis()` is aligned with NOORCHAIN official rules:

- `TotalSignals = 0`
- `TotalMinted = "0"`
- `MaxSignalsPerDay = 50`
- `ParticipantShare = 70`
- `CuratorShare = 30`

`ValidateGenesis()` enforces:

- Non-nil state.
- `ParticipantShare + CuratorShare == 100`.
- `TotalMinted` not empty.
- `MaxSignalsPerDay > 0`.

**Economic effect today:** purely declarative.  
PoSS is still “counting only”; it does not mint or move any real NUR.

---

### 1.3. PoSS Params (params.go)

The `Params` struct (PoSS Logic 4+) defines tunable behaviour:

- `PoSSEnabled bool`
  - Master switch for PoSS rewards.
  - **Default:** `false` → PoSS is **economically OFF** by default.

- Limits:
  - `MaxSignalsPerDay uint64`
  - `MaxSignalsPerCuratorPerDay uint64`

- Rewards:
  - `MaxRewardPerDay sdk.Coin`
  - `BaseReward sdk.Coin`
  - Both denominated in `unur`.

- Weights:
  - `WeightMicroDonation uint32`
  - `WeightParticipation uint32`
  - `WeightContent uint32`
  - `WeightCCN uint32`

- Reserve + halving:
  - `PoSSReserveDenom string` (always `"unur"`)
  - `HalvingPeriodBlocks uint64` (placeholder, not configured yet)

`DefaultParams()`:

- `PoSSEnabled = false`
- `MaxSignalsPerDay = 20`
- `MaxSignalsPerCuratorPerDay = 100`
- `BaseReward = 1 unur`
- `MaxRewardPerDay = 100 unur`
- Weights: `5 / 2 / 3 / 1`
- `PoSSReserveDenom = "unur"`
- `HalvingPeriodBlocks = 0` (not active yet)

`Params.Validate()` ensures basic consistency (non-zero limits, correct denoms, non-negative amounts, non-zero weights).

---

### 1.4. Rewards helpers (rewards.go)

PoSS rewards are computed in **three steps**:

1. **Base reward**
   - `ComputeBaseReward(p, signalType)`  
   - Formula: `BaseReward * weight(signalType)`.

2. **Halving**
   - `ApplyHalving(p, height, reward)`  
   - Uses `HalvingPeriodBlocks` and current block height.
   - If `HalvingPeriodBlocks == 0`, reward is unchanged.

3. **70/30 split**
   - `SplitReward70_30(total)`  
   - Structural rule: `70 %` to participant, `30 %` to curator.
   - It ensures `participant + curator == total` (no rounding drift).

High-level helper:

- `ComputeSignalReward(p, signalType, height)`:
  - If `PoSSEnabled == false`:
    - returns `0/0` with correct denom (no economic effect).
  - Else:
    - Base reward → halving → 70/30 split.

---

### 1.5. Keeper overview (keeper/keeper.go)

The PoSS Keeper currently handles:

- Access to the KVStore:
  - `getStore(ctx sdk.Context) sdk.KVStore`.

- **Daily counters** per address and date:
  - `GetDailySignalsCount(ctx, address, date)`
  - `SetDailySignalsCount(ctx, address, date, count)`
  - `IncrementDailySignalsCount(ctx, address, date)`

- **Params & reward helpers:**
  - `GetParams(ctx)`  
    - Returns `DefaultParams()` for now (no ParamSubspace yet).
  - `ComputeSignalRewardForBlock(ctx, signalType)`  
    - Fetches params, gets `ctx.BlockHeight()`, calls `ComputeSignalReward()`.

- **Internal signal pipeline (no minting yet):**
  - `ProcessSignalInternal(ctx, participantAddr, curatorAddr, signalType, date)`:
    - Computes participant/curator rewards (theoretical).
    - Increments participant daily counter.
    - Returns rewards to the caller.
    - **Does NOT:**
      - enforce daily limits yet,
      - update `TotalSignals` / `TotalMinted`,
      - mint or move real coins.

Result:  
Today, PoSS is an **accounting and simulation layer**, ready to be connected to real bank/mint logic later, once the economic and legal conditions are satisfied.

---

## 2. Economic & Safety Guarantees (Current Stage)

At the current code stage:

- `PoSSEnabled = false` by default.
- No real NUR is minted by PoSS.
- No coins are moved from any reserve account.
- PoSS only:
  - counts signals (daily counters),
  - computes **theoretical** rewards,
  - exposes a clean place (`ProcessSignalInternal`) where final logic will live later.

This matches the **Legal Light CH** strategy:

- No promise of yield.
- No automatic on-chain minting from user actions yet.
- All economic switches remain under explicit, visible control (params + governance + reserve wiring, to be added later).

---

## 3. Targeted PoSS Testnet Preparation

The first wave of PoSS-focused tests should **not** try to mint real coins yet.  
Instead, it should validate:

- the **correct behaviour of counters**,  
- the **70/30 math**,  
- the **effect of PoSSEnabled = false**,  
- the **stability of the PoSS module inside the chain**.

### 3.1. Testnet prerequisites

Before running PoSS scenarios, the following must be available:

1. A local or dev testnet using the `noord` binary:
   - `noord init ...`
   - `noord start`
   - Genesis including `noorsignal` module with `DefaultGenesis()`.

2. A way to submit `MsgCreateSignal` transactions:
   - via CLI (once wired),
   - or via tests / manual tx builder in Go,
   - or via gRPC/REST tooling.

3. At least two accounts:
   - `participant` (user doing the PoSS signal),
   - `curator` (curator validating/receiving 30 % later).

---

### 3.2. Scenario A — PoSS disabled (PoSSEnabled = false)

**Goal:**  
Verify that when PoSS is disabled, signals are recorded but rewards are `0/0`.

**Initial conditions:**

- `PoSSEnabled = false` in `Params` (default).
- `BaseReward > 0` and weights > 0 (as in `DefaultParams`).
- `MaxSignalsPerDay` > 0.

**Steps:**

1. Start the testnet.
2. Submit a `MsgCreateSignal` with:
   - `SignalType = micro_donation` (or any valid type),
   - a valid participant and curator.
3. Internally, the keeper calls:
   - `ComputeSignalRewardForBlock` → returns `0/0` because `PoSSEnabled = false`.
   - `IncrementDailySignalsCount` for the participant.

**Expected results:**

- Daily counter for `(participant, today)` is incremented.
- Theoretical reward from `ProcessSignalInternal` is `0 unur / 0 unur`.
- No coins are minted or moved (this is enforced structurally by the current code).

---

### 3.3. Scenario B — PoSS enabled (simulation only)

**Goal:**  
Simulate what would happen **if** `PoSSEnabled` were set to `true`, without actually enabling minting.

**Approach:**

- In a controlled test (unit test / integration test), override `Params` passed to `ComputeSignalReward` with:
  - `PoSSEnabled = true`.
- Call `ComputeSignalReward` directly with different `SignalType` values.

**Checks:**

- For each `SignalType` (`micro_donation`, `participation`, `content`, `ccn`):
  - base reward = `BaseReward * weight`.
  - 70/30 split is correct:
    - participant amount = `total * 70 / 100`
    - curator amount = `total − participant`
  - halving behaviour is correct when `HalvingPeriodBlocks > 0` and height increases.

**Expected results:**

- Reward math is stable and deterministic.
- Small heights (before first halving) = no reduction.
- After N halving periods, reward is divided by `2^N`, but never negative.

---

### 3.4. Scenario C — Daily counters behaviour

**Goal:**  
Validate that daily counters behave correctly across multiple signals and days.

**Steps:**

1. On day D:
   - Send N signals for `participant` (any `SignalType`).
   - Verify counter for `(participant, D)` == N.
2. Advance block time to day D+1 (via test environment).
3. Send M signals on day D+1.
4. Verify:
   - `(participant, D)` stays at N.
   - `(participant, D+1)` == M.

**Note:**  
If a dedicated daily reset mechanism is wired later (with `KeyLastResetDay` and prefixes), we will extend the scenarios to cover automatic resets.

---

### 3.5. Scenario D — Genesis import / export

**Goal:**  
Ensure that `GenesisState` for PoSS is correctly serialized and deserialized.

**Steps:**

1. Create a custom `GenesisState` in tests:
   - `TotalSignals = 123`
   - `TotalMinted = "456000000000000000000"`
   - `MaxSignalsPerDay = 42`
   - `ParticipantShare = 70`
   - `CuratorShare = 30`

2. Call:
   - `InitGenesis` with this state.
   - Then `ExportGenesis`.

3. Compare exported state to the original (struct-level equality).

**Expected result:**

- The exported `GenesisState` is identical to the imported one, proving stable JSON encoding/decoding.

---

## 4. Next Steps (Beyond this Document)

This document freezes the **PoSS Logic 1–26** stage:

- PoSS is wired, parametrized, and safe (no minting).
- Rewards and daily counters are in place, but only used internally.
- Testnet scenarios are defined to validate the behaviour without any economic impact.

**Future PoSS phases will focus on:**

1. Wiring real PoSS parameters into `x/params` with a dedicated ParamSubspace.
2. Adding a real PoSS reserve account and safe minting/transfer logic.
3. Enforcing daily caps and halving in the live pipeline.
4. Exposing CLI / gRPC endpoints for `MsgCreateSignal` and PoSS queries.
5. Activating PoSS (`PoSSEnabled = true`) only when:
   - legal constraints are satisfied,
   - economic simulations are validated,
   - a clear communication/education layer is published.

