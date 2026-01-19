**NOORCHAIN — Phase 4A

Cosmos/EVM Full Development Blueprint**
Version 1.1 — Architecture & Technical Steps (no code)

🔧 1. Purpose of Phase 4A

This document defines the full development blueprint required to transform the Phase 2 skeleton + Phase 3 specifications into a fully functional NOORCHAIN 1.0 Core, before Testnet 1.0.

No code is included.
Only architecture, design, and sequenced technical steps.

This blueprint governs the implementation of:

Cosmos SDK integration

Ethermint EVM integration

NOORCHAIN parameters

Module wiring

App initialization

Genesis structure

ModuleManager lifecycle

State keepers

PoSS compatibility hooks (but no PoSS code here)

🧩 2. Versions & Compatibility (strict)
Component	Version
Cosmos SDK	0.50.x
CometBFT	0.37.x
Ethermint	0.27.x
IAVL	0.21+
Go	1.22+
CosmJS (later, for dApps)	0.33+

Compatibility was validated in Phase 3:
→ Cosmos SDK 0.50.x requires CometBFT 0.37.x
→ Ethermint 0.27.x requires Cosmos SDK 0.50.x
→ All components build together correctly.

🏗️ 3. High-Level Architecture

NOORCHAIN uses:

3.1 Base Application Layer (Cosmos SDK)

auth

bank

staking

gov

params (implicit through config.go)

3.2 EVM Layer (Ethermint)

evm

feemarket

3.3 Custom Layer

x/noorsignal (PoSS module, defined later in Phase 4B)

3.4 Application Composition

The application is composed under:

/app/app.go
/app/types.go
/config/*.go
/x/* modules
/cmd/noord/main.go


The App struct will include:

BaseApp

Codec (encoding config)

StoreKeys

Keepers

ModuleManager

🏛️ 4. Application Keepers (Full Plan)

In Phase 4, we implement real keeper instantiation.

4.1 Required Keepers

AccountKeeper (auth)

BankKeeper (bank)

StakingKeeper (staking)

GovKeeper (gov)

EVMKeeper (Ethermint)

FeeMarketKeeper (Ethermint)

SlashingKeeper (later optional)

PoSSKeeper (Phase 4B)

4.2 Keeper Wiring Rules

Order is critical:
staking → slashing → gov → bank → evm

EVMKeeper requires:

AccountKeeper

BankKeeper

StakingKeeper

PoSSKeeper requires:

AccountKeeper

BankKeeper

StakingKeeper

(later) Hooks into BeginBlock/EndBlock

📦 5. Store Keys & Subspaces

The following KVStores must be defined:

Module	StoreKey
auth	auth
bank	bank
staking	staking
gov	gov
evm	evm
feemarket	feemarket
noorsignal	noorsignal

Subspaces (x/params) are auto-generated in SDK 0.50.x.

🔄 6. Module Manager Lifecycle

The ModuleManager must define:

6.1 Ordering

BeginBlockers

EndBlockers

InitGenesis

ExportGenesis

6.2 Standard Order (Cosmos + Ethermint)
BeginBlock:   feemarket → evm → staking → noorsignal → gov
EndBlock:     staking → gov
InitGenesis:  auth → bank → staking → gov → evm → feemarket → noorsignal


This order is required for:

fee market updates

EVM state transitions

PoSS reward indexing

staking lifecycle

governance proposals

⚙️ 7. App Initialization Plan

The initialization sequence in app.go must follow this structure:

Create codec & interface registry

Create BaseApp

Define store keys

Create keepers

Set keeper dependencies (SetHooks, SetParams…)

Create ModuleManager

Register services (Msg, Query)

Register BeginBlock, EndBlock, InitGenesis

Load state from disk

Return fully constructed app

This blueprint ensures deterministic, reproducible app lifecycle.

🔐 8. Consensus Parameters (Phase 3 Reference)

Using Phase 3 definitions:

Block time target : ~1.0s

Max gas per block : Cosmos defaults

EVM base fee : dynamic (feemarket)

Halving every 8 years → implemented in PoSS (Phase 4B)

PoSS hooked into BeginBlock

🌍 9. Genesis Structure

The genesis for Testnet 1.0 will require:

9.1 Base Accounts

foundation wallet

dev wallet

stimulus wallet

presale wallet

PoSS reserve wallet

(All 5 real bech32 addresses will be added in Phase 4C)

9.2 Core Genesis Fields

auth.accounts

bank.balances

staking.params

gov.params

evm.params

feemarket.params

noorsignal (placeholder from Phase 4B)

9.3 Initial Supply Distribution

Conforms to official NOORCHAIN Genesis 5 / 5 / 5 / 5 / 80 rule.

🔌 10. Services & API Endpoints

The app must expose:

10.1 REST / gRPC

/cosmos.auth.v1beta1

/cosmos.bank.v1beta1

/cosmos.staking.v1beta1

/cosmos.gov.v1beta1

/ethermint.evm.v1

/ethermint.feemarket.v1

/noorchian.noorsignal.v1 (later in 4B)

10.2 JSON-RPC (EVM)

eth_call

eth_sendRawTransaction

eth_getBalance

eth_estimateGas

web3_clientVersion

All via Ethermint RPC API.

🔨 11. Development Steps (Ordered, No Code)
Step 1 — Finalize store keys

Define all keys in app.go.

Step 2 — Create keepers

Instantiate all Cosmos + Ethermint keepers.

Step 3 — Wire keeper dependencies

staking hooks

evm keeper dependencies

feemarket params

PoSS placeholder keeper

Step 4 — Build ModuleManager

Register all modules with correct ordering.

Step 5 — Register services

Msg servers, query servers.

Step 6 — Register begin/end block hooks
Step 7 — InitGenesis wiring
Step 8 — ExportGenesis wiring
Step 9 — App constructor cleanup

Remove Phase 2 placeholders.

🧪 12. Test Requirements Before Testnet

Before Phase 4C, the app MUST:

compile without warnings

noord start runs (empty chain)

noord genesis export works

EVM RPC starts correctly

gRPC starts correctly

JSON-RPC responds to web3_clientVersion

✅ END OF FILE

Phase4_01_CosmosEVM_Development_Blueprint_1.1.md

🔧 Séquence Git (à exécuter après création du fichier sur GitHub)
git add docs/Phase4_01_CosmosEVM_Development_Blueprint_1.1.md
git commit -m "Phase 4 – File 01: Cosmos/EVM Development Blueprint (1.1)"
git push
