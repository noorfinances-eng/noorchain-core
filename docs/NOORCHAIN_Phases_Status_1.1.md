# NOORCHAIN 1.0 — Global Phases Status (v1.1)

**Scope:**  
This document gives a high-level status of all NOORCHAIN 1.0 phases (1 → 9), at the current development stage.  
It is meant as a *single overview file* to quickly understand what is done, what is in progress, and what remains.

It complements the more detailed Phase 3 / PoSS documents in `docs/`:

- `NOORCHAIN_Phase3_01_Architecture_1.1.md`
- `NOORCHAIN_Phase3_02_Genesis_1.1.md`
- `NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md`
- `NOORCHAIN_Phase3_04_Economic_Model_1.1.md`
- `NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md`
- `NOORCHAIN_Phase3_06_PoSS_Testnet_PracticalGuide_1.1.md`

---

## 1. Phase 1 — Cadrage & Décisions

**Status:** ✅ 100% completed

### 1.1. Technical decisions

- Core stack: **Cosmos SDK + Ethermint (EVM) + CometBFT**, monolithic app `noorchain-core`.
- Native token: **NUR**, fixed supply `299 792 458` (speed of light reference).
- Monetary policy:
  - Fixed cap: **299 792 458 NUR**.
  - **Halving every 8 years** (for PoSS issuance tempo), implemented via parameters.

### 1.2. PoSS (Proof of Signal Social)

- 4 social signal families:
  - `micro_donation`
  - `participation`
  - `content`
  - `ccn` (NOOR certified content / CCN Studio)
- Structural reward split (fixed, protocol-level rule):
  - **70%** to the participant.
  - **30%** to the curator.
- Daily limits:
  - per-participant signals per day,
  - per-curator validated signals per day,
  - per-participant max reward per day (cap).

### 1.3. Economic model

- Official allocation model (**Genesis 5 / 5 / 5 / 5 / 80**):
  - **5%** NOOR Foundation (public).
  - **5%** Noor Dev Sàrl (founder / dev structure).
  - **5%** PoSS Stimulus pool.
  - **5%** optional pre-sale (vested, multi-sig, Swiss-only).
  - **80%** PoSS mintables over time.
- Revenue levers:
  - Potential NUR value appreciation.
  - dApps / SaaS revenue (NOOR Pay, Curators Hub, CCN Studio).
  - B2B partnerships and institutional grants.
  - Responsible founder exit (order-book style, capped sales).

### 1.4. Legal framework

- Jurisdiction: **Switzerland**, “FINMA Light / Legal Light CH”.
- No on-chain custody of client fiat.
- No yield promises, no “investment product” marketing.
- Future fiat conversion only via **regulated PSP partners** (e.g. Mt Pelerin, NOWPayments).

> **Phase 1 is fully closed and considered stable.**

---

## 2. Phase 2 — Technical Skeleton (Cosmos/EVM Core)

**Status:** ✅ 100% completed (clean skeleton)

### 2.1. Repository structure

- `cmd/noord` — minimal CLI binary (version / placeholder run).
- `config/` — basic configuration helpers.
- `scripts/` — shell utilities, including `scripts/testnet.sh`.
- `x/auth`, `x/bank`, `x/staking`, `x/gov` — module placeholders.
- `x/evm`, `x/feemarket` — Ethermint modules.
- `x/noorsignal` — PoSS module (see Phase 4).

### 2.2. `app/app.go` — NoorchainApp

Core app struct:

- Embeds `*baseapp.BaseApp`.
- Holds:
  - `appCodec`, `interfaceRegistry`, `txConfig`.
  - `keys` (KV stores) for:
    - `auth`, `bank`, `staking`, `gov`, `params`, `evm`, `feemarket`, `noorsignal`.
  - `tkeys` (transient stores) for:
    - `params`, `evm`, `feemarket`.
- `ParamsKeeper`:
  - Proper KV + transient store.
  - Subspaces for `auth`, `bank`, `staking`, `gov`, `evm`, `feemarket`.
- Keepers:
  - `AccountKeeper`, `BankKeeper`, `StakingKeeper`.
  - `FeeMarketKeeper` (Ethermint) with module address `"gov"`.
  - `EvmKeeper` (Ethermint) with:
    - EVM KV + transient stores.
    - Access to Account / Bank / Staking / FeeMarket.
  - `NoorSignalKeeper` (PoSS) with:
    - codec + store key + ParamSubspace (later).

### 2.3. Module manager

- `module.NewManager(...)` with:
  - `auth`, `bank`, `staking`,
  - `evm`, `feemarket`,
  - `noorsignal` (PoSS).
- Init order:
  - `auth` → `bank` → `staking` → `evm` → `feemarket` → `noorsignal`.
- Lifecycle:
  - `InitChainer`, `BeginBlocker`, `EndBlocker` delegated to `app.mm`.

### 2.4. Encoding & Tx

- `MakeEncodingConfig()`:
  - Legacy Amino (minimal).
  - `InterfaceRegistry`.
  - `ProtoCodec`.
  - `authtx.NewTxConfig`.
- `NoorchainApp` stores `txConfig` for AnteHandler.

> **Phase 2 is considered complete and clean.**  
> All later work builds on this skeleton.

---

## 3. Phase 3 — Documentation & Specifications

**Status:** 🟡 ~80% completed

### 3.1. Documents already in place

- `NOORCHAIN_Phase3_01_Architecture_1.1.md`
  - High-level architecture (Cosmos/EVM modules, stores, keepers).
- `NOORCHAIN_Phase3_02_Genesis_1.1.md`
  - Genesis structure, module roles, allocation principles.
- `NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md`
  - Functional specs for PoSS:
    - signal types,
    - counters,
    - limits,
    - weights,
    - halving,
    - 70/30 rule.
- `NOORCHAIN_Phase3_04_Economic_Model_1.1.md`
  - 5/5/5/5/80 model, fund usage rules, Legal Light alignment.
- `NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md`
  - Technical status of `x/noorsignal` module.
  - Testnet scenario definitions (no mint, counters only).
- `NOORCHAIN_Phase3_06_PoSS_Testnet_PracticalGuide_1.1.md`
  - Practical guide:
    - how to init the testnet,
    - how PoSS behaves with `PoSSEnabled=false`,
    - expected behaviour (0 reward, counters increment).

### 3.2. Still missing to reach 100%

- Cosmos/Ethermint dev setup doc:
  - `go build`, `go test`, `noord` binary, local testnet routines.
- Ante & EVM fees doc:
  - structure of the AnteHandler,
  - gas / fees model,
  - min gas price strategy.
- API / Queries plan:
  - future gRPC/CLI queries for PoSS (stats, daily counters, etc.).
- Phase 3 “Completed pack” checklist:
  - a short file confirming all Phase 3 docs are done.

---

## 4. Phase 4 — Implementation (Code)

**Status:** 🟢 ~95% completed

This is the main coding phase for the core app and PoSS module.

### 4.1. Core app & EVM / FeeMarket

- `app/app.go`:
  - fully wired Cosmos core.
  - Ethermint EVM + FeeMarket integrated with real keepers and stores.
  - `NoorSignalKeeper` registered and module added to `ModuleManager`.
- `app/ante.go`:
  - AnteHandler using Ethermint’s ante stack.
  - Handles signatures, gas, fees, EVM-specific behaviour.
  - Compiles and runs without conflict.

> Core app is stable at this stage.

---

### 4.2. PoSS module — Types (`x/noorsignal/types`)

Implemented:

- `SignalType`:
  - `micro_donation`, `participation`, `content`, `ccn`.
- `Params`:
  - `PoSSEnabled` (master switch, default `false`).
  - `MaxSignalsPerDay`, `MaxSignalsPerCuratorPerDay`.
  - `MaxRewardPerDay` (per-participant daily cap).
  - `BaseReward` (unit reward).
  - Weights:
    - `WeightMicroDonation`, `WeightParticipation`, `WeightContent`, `WeightCCN`.
  - `PoSSReserveDenom` (always `"unur"`).
  - `HalvingPeriodBlocks` (placeholder at `0` for now).
- `DefaultParams()`:
  - “safe off” configuration:
    - PoSS disabled,
    - reasonable default limits,
    - BaseReward/MaxRewardPerDay in `unur`.
- `Params.Validate()`:
  - basic consistency checks on denoms, amounts, weights.

Rewards helpers:

- `WeightForSignalType(...)`:
  - selects the correct weight per signal type.
- `ComputeBaseReward(...)`:
  - `BaseReward * weight(signal_type)`.
- `ApplyHalving(...)`:
  - applies halving based on `HalvingPeriodBlocks` and block height.
- `SplitReward70_30(...)`:
  - structural 70/30 split with exact sum preservation.
- `ComputeSignalReward(...)`:
  - if `PoSSEnabled=false` → returns `0/0` with correct denom.
  - else:
    - base reward → halving → 70/30.

Genesis:

- `GenesisState`:
  - `TotalSignals` (uint64).
  - `TotalMinted` (string in `unur`).
  - `MaxSignalsPerDay` (declarative).
  - `ParticipantShare=70`, `CuratorShare=30`.
- `DefaultGenesis()` and `ValidateGenesis(...)` implemented.

Daily counters / keys:

- Keys for:
  - `KeyGenesisState`.
  - Total counters.
  - Per-address-per-day counters (`DailyCounterKey(addr, date)`).

Unit tests:

- Rewards tests (PoSS enabled/disabled, weights, 70/30) passing.

---

### 4.3. PoSS module — Keeper (`x/noorsignal/keeper`)

Keeper responsibilities:

- Holds:
  - `cdc Codec`,
  - PoSS `storeKey`,
  - PoSS `ParamSubspace`.
- Ensures ParamSubspace has the correct `KeyTable`.

Genesis helpers:

- `InitGenesis(...)` / `ExportGenesis(...)`:
  - store and retrieve genesis-equivalent state via JSON under `KeyGenesisState`.

Global stats:

- `GetGlobalStats(...)`:
  - returns `PoSSStats`:
    - `TotalSignals`, `TotalMinted`,
    - `PoSSEnabled`,
    - daily limits,
    - reserve denom.

Daily counters:

- `GetDailySignalsCount(...)`
- `SetDailySignalsCount(...)`
- `IncrementDailySignalsCount(...)`

Params management:

- `SetParams(ctx, params)`:
  - validates then stores in ParamSubspace.
- `GetParams(ctx)`:
  - if no params stored yet:
    - writes `DefaultParams()` into the ParamSubspace,
    - returns them.
  - if stored params are invalid:
    - returns `DefaultParams()` as a safety fallback.

Rewards wrapper:

- `ComputeSignalRewardForBlock(ctx, signalType)`:
  - fetches PoSS params,
  - uses `ctx.BlockHeight()`,
  - calls `ComputeSignalReward(...)`.

Pending mint queue (planning only):

- `RecordPendingMint(...)`:
  - stores a `PendingMint` JSON entry:
    - `BlockHeight`,
    - `Timestamp`,
    - `Participant`, `Curator`,
    - `SignalType`,
    - `ParticipantReward`, `CuratorReward`.
  - **No real mint or bank transfer yet.**

Signal pipeline:

- `ProcessSignalInternal(ctx, participant, curator, signalType, date)`:
  - computes participant/curator rewards for the signal.
  - enforces:
    - daily participant signals limit,
    - daily curator signals limit,
    - per-day participant max reward cap.
  - increments daily counters.
  - records a `PendingMint` entry.
  - loads GenesisState, increments:
    - `TotalSignals` by 1,
    - `TotalMinted` by `participantReward + curatorReward`.
  - returns the computed rewards to the caller.

> At this stage, **no real minting happens**.  
> PoSS is economically “locked” even when logically enabled, because the Bank / reserve account is not wired yet.

Unit tests (keeper):

- `TestGetParams_DefaultsStored`
- `TestSetParams_RoundTrip`
- `TestProcessSignalInternal_UpdatesCountersAndGenesis`
- Additional tests for daily caps and MaxRewardPerDay.

All tests pass via:

```bash
go test ./...
4.4. PoSS module — AppModule (x/noorsignal/module.go)
AppModuleBasic:

Name()

DefaultGenesis(...)

ValidateGenesis(...)

RegisterGRPCGatewayRoutes(...) (currently empty)

RegisterInterfaces(...) (currently empty)

GetTxCmd(), GetQueryCmd() → nil (no CLI yet).

AppModule:

embeds AppModuleBasic.

holds cdc + keeper.

NewAppModule(cdc, keeper) constructor.

InitGenesis(...) / ExportGenesis(...) implemented via keeper.

BeginBlock(...) / EndBlock(...) currently no PoSS-specific logic.

ConsensusVersion(), RegisterInvariants(...), legacy routing stubs etc.

Module is fully integrated in app/app.go and wired in the ModuleManager.

4.5. Testnet 0.1 — Filesystem Testnet
A minimal local testnet bootstrap is implemented via:

testnet/genesis.json

testnet/genesis_distribution.json

scripts/testnet.sh

Typical routine:

bash
Copier le code
cd /workspaces/noorchain-core
rm -rf data-testnet
go build -o noord ./cmd/noord
go test ./...
scripts/testnet.sh
Result:

data-testnet/config/genesis.json is created from testnet/genesis.json.

git status remains clean:

data-testnet/ and noord are ignored (not committed).

This is a filesystem-only testnet.
noord start is not fully wired as a real node yet (CLI/server work planned later).

4.6. What remains for Phase 4 to reach 100%
Wire a real PoSS reserve module account (with Minter permission).

Implement real minting / bank transfers from the reserve:

credit 70% to participant, 30% to curator.

Implement MsgCreateSignal (proto, types, msg server).

Implement minimal gRPC/CLI queries for PoSS (stats, counters, etc.).

Wire noord CLI for a usable node (start, tx, etc.).

These items mark the transition between Phase 4 final steps and the beginning of Phase 5/6 (governance, communication, mainnet prep).

5. Phase 5 — Legal & Governance
Status: 🟡 Conceptually advanced, not yet fully documented in this repo

Legal model:

Legal Light CH (no yield promises, no fiat custody, transparent rules).

PSP partners for fiat conversion (no internal PSP).

Governance model:

On-chain parameters (PoSS, fees, limits) adjustable by governance.

Public, transparent management of Genesis allocations (Foundation, Dev, Stimulus, Pre-sale).

Investor funds management:

All investor funds handled via Noor Dev Sàrl bank/payment accounts.

Tokens delivered from the 5% pre-sale pool, multi-sig, vested.

To do in this repo:

Dedicated Phase 5 legal/governance doc:

FINMA Light classification.

Governance processes.

Integration of legal rules into Genesis and PoSS.

6. Phase 6 — Genesis Pack & Communication
Status: 🟡 Partially done

Already defined:

Genesis economic model (5/5/5/5/80).

PoSS integration at genesis.

Legal Light constraints reflected in docs.

To be done:

Real files for:

testnet/genesis.json (already minimal, to be extended).

testnet/genesis_distribution.json (fill with real addresses).

Scripts in scripts/ (testnet tooling).

Address alignment:

5 real bech32 addresses (foundation, founder, stimulus, pre-sale, PoSS reserve).

Keep them consistent across:

genesis.json,

genesis_distribution.json,

future types/addresses.go (when created).

Public Genesis Pack:

summarized PDF / doc for external readers.

allocation tables.

7. Phase 7 — Mainnet 1.0
Status: ⚪ 0% (planned)

Next steps (after testnets):

Public PoSS testnet with:

PoSSEnabled = true,

real reserve account,

live caps and halving parameters.

Final parameter calibration (limits, weights, halving).

External audits:

security,

economic,

legal.

Mainnet genesis:

final genesis file,

launch plan,

monitoring / observability setup.

8. Phase 8 — dApps & Ecosystem
Status: ⚪ Vision only (no code yet in this repo)

Target dApps:

NOOR Pay:

simple payment rails on top of NOORCHAIN (PoSS-aware where relevant).

Curators Hub:

tooling for NOOR Curators (Bronze / Silver / Gold),

validation dashboards, stats, PoSS signals.

CCN Studio:

content certification platform integrated with PoSS.

To be defined:

APIs exposed by the core (REST/gRPC/EVM).

UX / front-end projects (separate repos).

Economic integration (fees, PoSS signals, rewards visibility).

9. Phase 9 — Partnerships & Audits
Status: ⚪ 0% (future)

Planned scope:

External audits:

code audits (Cosmos/Ethermint, PoSS module),

economic/game-theory review,

legal opinion on Legal Light CH alignment.

Partnerships:

NGOs, schools, social actors (Curators network).

PSP partners (for fiat conversion).

Ecosystem partners (validators, infra providers).

10. Summary Table
Phase	Name	Status	Notes
1	Cadrage & Décisions	✅ 100%	Vision, PoSS model, economic & legal framework set.
2	Technical Skeleton	✅ 100%	Core app, modules, keepers, Params skeleton done.
3	Documentation & Specs	🟡 ~80%	Main docs done, some dev/fees/API docs pending.
4	Implementation (Code)	🟢 ~95%	Core + PoSS logic coded & tested, no real mint yet.
5	Legal & Governance	🟡 Conceptual	Needs dedicated docs + on-chain governance wiring.
6	Genesis Pack & Communication	🟡 Partial	Genesis files started, needs real addresses & pack.
7	Mainnet 1.0	⚪ 0%	To do after stable testnets & audits.
8	dApps & Ecosystem	⚪ 0%	Vision only (NOOR Pay, Curators Hub, CCN Studio).
9	Partnerships & Audits	⚪ 0%	Future external phase.

11. Current Technical Snapshot (December 2025)
go build ./... → OK

go test ./... → OK, including:

x/noorsignal/types tests.

x/noorsignal/keeper tests (params, counters, genesis, caps).

scripts/testnet.sh:

creates a minimal local testnet filesystem in ./data-testnet.

copies testnet/genesis.json to data-testnet/config/genesis.json.

git status:

clean after each step (binaries & data dirs ignored).

PoSS economic switch:

PoSSEnabled = false by default:

PoSS can count signals and update totals,

but does not mint or move real NUR yet.

This state is intentional and aligned with the Legal Light CH approach.