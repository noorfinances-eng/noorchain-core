NOORCHAIN 1.0 — Global Phases Status (v1.2, EN)

Scope
This document provides a complete overview of all NOORCHAIN 1.0 phases (1 → 9).
It serves as a single, high-level status reference for internal use.

============================================================

Phase 1 — Framing & Decisions
Status: 100% complete

Technical decisions

Core: Cosmos SDK + Ethermint (EVM) + CometBFT

Native token: NUR

Fixed supply: 299,792,458 (speed of light concept)

PoSS halving every 8 years (param-based)

PoSS model

Four signal families: micro_donation, participation, content, ccn

Structural immutable rule: 70% participant / 30% curator

Daily limits: participant, curator, participant reward cap

Economic model

Official allocation 5 / 5 / 5 / 5 / 80

Investor funds handled via Noor Dev Sàrl only

No internal PSP. Fiat flows exclusively through regulated partners

Fully compliant with Swiss Legal Light CH (no yield promises, no fiat custody)

Phase 1 is locked and will never be modified.

============================================================
2. Phase 2 — Technical Skeleton
Status: 100% complete

Deliverables

Clean repository layout

Complete app/app.go structure

Mounted stores: auth, bank, staking, gov, params, evm, feemarket, noorsignal

Initialized keepers

Minimal but functional ModuleManager

Full encoding configuration

Initial scripts (testnet.sh)

This skeleton is stable and serves as the foundation for all later phases.

============================================================
3. Phase 3 — Documentation & Specifications
Status: ~80% complete

Documents completed

Architecture overview

Genesis structure and allocation principles

Full PoSS functional specification

Economic model (5/5/5/5/80)

PoSS technical status

PoSS testnet behavioural guide

Documents still missing

AnteHandler and EVM gas/fee specification

CLI/API/Query design specification

Final “Phase 3 Completion Pack”

Phase 3 is advanced but not fully closed.

============================================================
4. Phase 4 — Implementation (Code)
Status: ~95% complete

4.1 Core App / Cosmos / EVM

Fully wired Cosmos SDK 0.46.11

Ethermint 0.22.0 EVM integrated

FeeMarket fully operational

AnteHandler functional

Builds and unit tests fully passing

4.2 PoSS module — Types

Complete Params structure

DefaultParams (PoSS OFF safe mode)

Signal weights and reward engine

Halving logic (param-based)

Immutable 70/30 split

GenesisState implemented

Daily counters implemented

Reward cap implemented

All types-level tests passing

4.3 PoSS module — Keeper

Parameter management via real ParamSubspace

Fully functional internal pipeline (without minting)

PendingMint scheduling queue

Daily counters

Global PoSS statistics

Genesis loading/export

All keeper tests passing

4.4 PoSS module — AppModule

AppModuleBasic complete

Module integrated in ModuleManager

BeginBlock / EndBlock placeholders

No MsgServer or QueryServer yet

4.5 Testnet

testnet/genesis.json functional

genesis_distribution.json awaiting real addresses

scripts/testnet.sh working

data-testnet/ created deterministically

4.6 Remaining work to reach 100%

Real PoSS reserve module account

Actual mint/send via BankKeeper

MsgServer (proto)

QueryServer

CLI integration

Activation of PoSS when legally and technically safe

============================================================
5. Phase 5 — Legal & Governance
Status: conceptually advanced (~60%)

Completed conceptually

Full Legal Light CH alignment

Governance parameters defined (voting, tally, deposit)

Strict allocation model (5/5/5/5/80)

Multi-sig rules for sensitive pools

Investor fund-flow rules (via Sàrl only)

Still required in the repo

Dedicated Phase 5 document

FINMA Light classification documentation

Governance processes description

Integration of governance parameters inside genesis.json

============================================================
6. Phase 6 — Genesis Pack & Communication
Status: partly complete (~40%)

Completed

Economic model

Initial genesis templates

Internal documentation

Remaining

Inject 5 real Bech32 addresses (foundation, dev, stimulus, presale, PoSS reserve)

Align all files (genesis.json, genesis_distribution.json)

Produce the public “Genesis Pack”

Public-facing summary documentation

============================================================
7. Phase 7 — Mainnet 1.0
Status: 0%

Will include

Public PoSS testnet (PoSS ON)

Calibration of PoSS parameters

Security and economic audits

Finalized mainnet genesis

Mainnet launch plan

============================================================
8. Phase 8 — dApps & Ecosystem
Status: vision only

Target dApps

NOOR Pay

Curators Hub

CCN Studio

Public APIs

No code expected in the core repo for now.

============================================================
9. Phase 9 — Partnerships & Audits
Status: 0%

Future deliverables

External technical audits

Economic/game-theory review

Legal audit for FINMA Light

Partnerships with NGOs, schools, institutions

PSP integration agreements

============================================================
10. Summary Table

Phase 1 — 100%
Phase 2 — 100%
Phase 3 — 80%
Phase 4 — 95%
Phase 5 — 60%
Phase 6 — 40%
Phase 7 — 0%
Phase 8 — 0%
Phase 9 — 0%

============================================================
11. Current Technical Snapshot (Dec 2025)

go build → OK
go test → OK
PoSS OFF by default
Minimal deterministic testnet working
ModuleManager stable
Legal Light preserved
No conflicting dependencies
No dangling logic

Overall status
NOORCHAIN 1.0 is in a clean, consistent, highly stable state suitable for Swiss-grade blockchain development.
