# NOORCHAIN — Phase 4B  
PoSS Msg Status (MsgCreateSignal) — v1.1

> This document summarizes the current status of the PoSS message layer
> in the `x/noorsignal` module (December 2025).  
> Scope: **Go-level Msg**, **legacy handler**, **no protobuf MsgServer yet**.

---

## 1. Scope of Phase 4B

The goal of Phase 4B is **not** to expose a full public tx API yet.

Instead, Phase 4B delivers:

- A **Go-native PoSS message type**: `MsgCreateSignal`.
- A **legacy handler** `NewHandler` in `x/noorsignal/handler.go` able to process it.
- A clear definition of what PoSS signals look like at the tx level.
- Internal consistency with the existing PoSS pipeline:
  - Params (with daily limits and MaxRewardPerDay),
  - daily counters,
  - pending mint queue,
  - genesis totals (TotalSignals / TotalMinted).

Public exposure (CLI, gRPC, REST, MsgServer protobuf) is explicitly postponed
to a later phase.

---

## 2. What exists today

### 2.1 MsgCreateSignal (Go struct)

File: `x/noorsignal/types/msg.go`

The core PoSS tx is defined as a **plain Go struct**:

- `Participant` (string, bech32 `noor1...`)
- `Curator` (string, bech32 `noor1...`)
- `SignalType` (`SignalTypeMicroDonation`, `SignalTypeParticipation`, `SignalTypeContent`, `SignalTypeCCN`)
- `Metadata` (string, optional)
- `Timestamp` (`time.Time`, logical client timestamp)
- `Date` (string `"YYYY-MM-DD"`, used for daily counters and limits)

Provided helpers and methods:

- `NewMsgCreateSignal(...)` helper constructor.
- `ValidateBasic()`:
  - checks non-empty participant / curator,
  - validates both as valid bech32 addresses,
  - validates `SignalType` against the 4 allowed values,
  - validates `Date` format (`YYYY-MM-DD`).
- `GetSignBytes()`:
  - JSON marshaling + `sdk.MustSortJSON`.
- `GetSigners()`:
  - returns a single signer: the **participant**.
- Minimal `proto.Message` compliance (`Reset`, `String`, `ProtoMessage`),
  required because `sdk.Msg` embeds `proto.Message`.

At this stage, **there is no `.proto` definition for MsgCreateSignal**.  
The type lives purely in Go and is mainly useful for:

- internal tests,
- future CLI / MsgServer wiring,
- documentation of the PoSS tx model.

---

### 2.2 Legacy handler for MsgCreateSignal

File: `x/noorsignal/handler.go`

A **legacy handler** exists to process `MsgCreateSignal`:

- It pattern-matches on `*noorsignaltypes.MsgCreateSignal`.
- Steps:
  1. Calls `ValidateBasic()`.
  2. Calls `Keeper.ProcessSignalInternal(...)` with:
     - participant address,
     - curator address,
     - signal type,
     - date (`YYYY-MM-DD`).
  3. Receives `(participantReward, curatorReward, error)`.
  4. Emits a `noorsignal_create_signal` event with:
     - `participant`,
     - `curator`,
     - `signal_type`,
     - `participant_reward`,
     - `curator_reward`.
  5. Returns a minimal `sdk.Result`.

Important points:

- The handler uses the **internal PoSS pipeline**:
  - PoSS Params via `GetParams()` (with `DefaultParams()` fallback),
  - daily signal counters,
  - `MaxRewardPerDay` cap for participant rewards,
  - pending mint queue,
  - genesis totals (TotalSignals / TotalMinted).
- It **does not mint any real coins** and does not touch the Bank module.
- It is designed to be attached to the legacy router later, when needed.

---

### 2.3 Keeper-side PoSS logic used by the handler

File: `x/noorsignal/keeper/keeper.go`

Relevant parts:

- `GetParams(ctx)`:
  - uses a real `ParamSubspace`,
  - if no params are stored, writes `DefaultParams()` and returns them,
  - on validation error, returns `DefaultParams()` (fails safe).
- `ComputeSignalRewardForBlock(ctx, signalType)`:
  - wraps `ComputeSignalReward` from `types/rewards.go`,
  - uses current block height,
  - applies PoSS halving and the 70/30 split.
- `ProcessSignalInternal(ctx, participant, curator, signalType, date)`:
  - computes raw rewards,
  - enforces `MaxRewardPerDay` **for the participant**:
    - if the cap is zero → cap disabled,
    - if the participant is already at or above the cap → this signal gives `0/0`,
    - if this signal would exceed the cap → this signal gives `0/0`,
    - otherwise → reward is accounted and stored in a daily reward counter,
  - increments the participant’s daily signal counter,
  - records a `PendingMint` entry (for future real minting),
  - updates `TotalSignals` and `TotalMinted` in `GenesisState`,
  - returns `(participantReward, curatorReward)`.

This makes `MsgCreateSignal` **economically meaningful** at the logic level, but still
safe (no real mint, PoSS disabled by default, etc.).

---

## 3. What does NOT exist yet

The following pieces are **NOT implemented yet**, by design:

1. **Protobuf definitions**:
   - No `.proto` file for `MsgCreateSignal` or `MsgCreateSignalResponse`.
   - No generated `*.pb.go` for PoSS messages.

2. **MsgServer (gRPC)**:
   - No `MsgServer` implementation.
   - No `RegisterMsgServer` call in `RegisterServices`.

3. **Public tx API**:
   - No CLI command (no `Tx` subcommands for PoSS).
   - No REST or gRPC-JSON gateway routes.
   - No public client library / SDK.

4. **Router wiring**:
   - The legacy `NewHandler` is **not yet attached** to the app’s `Router`.
   - This is intentional: Msg routing will be done during a later sub-phase,
     once the protobuf / gRPC story is clearer.

As a result, **no external client can currently broadcast a PoSS tx** to a node
and have it processed as a normal transaction. The message type and handler
are ready, but they are **internal only** for now.

---

## 4. Safety status

- `PoSSEnabled` is `false` in `DefaultParams()`.  
  → PoSS rewards are effectively “OFF” by default.
- The PoSS module:
  - never mints real `unur`,
  - never moves balances in the Bank module,
  - only:
    - counts signals,
    - enforces daily caps,
    - maintains genesis-level counters,
    - records `PendingMint` entries for future use.
- All PoSS parameters are stored in a standard `ParamSubspace`
  and are governed via the PoSS `Params` struct.  
  This design stays compatible with future on-chain governance.

---

## 5. Next steps (beyond Phase 4B)

The next logical steps after Phase 4B are:

1. **Phase 4C — PoSS MsgTestnet / Simulation doc**:
   - describe how a PoSS tx *would* be built and processed on a local testnet,
   - define the expected behavior:
     - reward calculation,
     - daily counters,
     - MaxRewardPerDay,
     - pending mint queue,
     - genesis totals.

2. **Later PoSS Messaging Phase** (not part of Phase 4B):
   - introduce `.proto` definitions for MsgCreateSignal,
   - generate `*.pb.go`,
   - implement `MsgServer`,
   - wire `RegisterServices`,
   - add CLI / gRPC / REST access,
   - update testnet scripts and docs for real PoSS tx testing.

For now, **Phase 4B is considered complete** when:

- the Go-level `MsgCreateSignal` and its tests are stable,
- the legacy handler exists and compiles,
- the PoSS keeper logic (Params, counters, caps, pending mint, totals)
  is integrated and tested,
- this document (`NOORCHAIN_Phase4B_PoSS_Msg_Status_1.1.md`) is in the repo.
