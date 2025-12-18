> ⚠️ **Status: TECH MIGRATION**  
> This document is being realigned to the **NOORCHAIN 2.0 Technical Baseline**.  
> Reference: `NOORCHAIN_Tech_Baseline_2.0_1.1.md`  
> Branch: `main-3` — Date: 2025-12-18

1. Purpose of This Document

This document defines the ModuleManager architecture of NOORCHAIN:

module registration

service registration

genesis flow

begin/end block execution order

interoperability between modules

requirements for PoSS integration (later: Phase 4B)

This design will be used to implement the real ModuleManager in app/app.go.

🧩 2. ModuleManager Role

The ModuleManager in Cosmos SDK is responsible for:

Initializing all modules

Executing lifecycle callbacks

Registering services (Msg, Query)

Managing the precise block execution order

Handling module interactions in a deterministic way

NOORCHAIN requires strict ordering because of:

EVM execution

Fee market logic

PoSS signal processing

governance & staking dependencies

🏛️ 3. Modules Included in NOORCHAIN

The ModuleManager must manage the following modules:

3.1 Cosmos SDK Modules

auth

bank

staking

gov

3.2 Ethermint Modules

evm

feemarket

3.3 Custom Module

noorsignal (PoSS)

introduced in Phase 4B Blueprint

lifecycle defined here for future integration

🔄 4. Execution Order Requirements

Execution order is critical for chain stability.
The following order has been validated in Phase 3 and in Ethermint reference chains.

4.1 BeginBlock Order

Final BeginBlock sequence:

1. FeeMarket     (dynamic base fee update)
2. EVM           (prepare EVM block context)
3. Staking       (validator updates, power recalculation)
4. NoorSignal    (PoSS reward logic, halving checks)
5. Gov           (tally ongoing proposals)

Reasoning

fee market first → ensures correct gas economics

evm second → EVM block context must be ready

staking third → validator updates before PoSS

PoSS fourth → needs validated validator power

gov last → governance depends on staking results

4.2 EndBlock Order
1. Staking
2. Gov

Reasoning:

validator updates must finalize before governance can finalize voting power.

🌍 5. InitGenesis Order

Genesis initialization must follow:

1. auth
2. bank
3. staking
4. gov
5. evm
6. feemarket
7. noorsignal

Reasoning:

accounts must exist before balances

balances must exist before staking

staking must exist before governance voting power

EVM requires complete state before loading EVM genesis

fee market depends on EVM

PoSS must initialize last (needs full state)

📤 6. ExportGenesis Order
auth → bank → staking → gov → evm → feemarket → noorsignal


Matches InitGenesis order for deterministic state.

🧱 7. Service Registration Order

Services must be registered before the app loads its state:

auth → bank → staking → gov → evm → feemarket → noorsignal


Each module registers:

Msg servers

Query servers

🧠 8. Module Interactions Summary

The ModuleManager ensures correct interaction ordering:

Module	Depends on	Notes
auth	none	base layer
bank	auth	balances depend on accounts
staking	bank	delegation depends on balances
gov	staking	voting depends on validator power
evm	auth, bank, staking	Ethermint core
feemarket	evm	EIP-1559 model
noorsignal	staking, bank	PoSS rewards depend on validator set
🧩 9. ModuleManager Construction Blueprint

The ModuleManager in NOORCHAIN will follow this sequence:

Step 1

Register module basics: auth, bank, staking, gov

Step 2

Register Ethermint modules: evm, feemarket

Step 3

Register NOORCHAIN PoSS module: noorsignal (from Phase 4B)

Step 4

Define service registration (Msg + Query)

Step 5

Define begin-block order (critical)

Step 6

Define end-block order

Step 7

Define InitGenesis order

Step 8

Define ExportGenesis order

Step 9

Register module invariants (if needed)

📌 10. High-Level Diagram
ModuleManager
│
├── RegisterModules()
│     ├── auth
│     ├── bank
│     ├── staking
│     ├── gov
│     ├── evm
│     ├── feemarket
│     └── noorsignal
│
├── SetBeginBlockerOrder()
├── SetEndBlockerOrder()
├── SetInitGenesisOrder()
└── SetExportGenesisOrder()

🎯 11. Summary Table
Stage	Modules
BeginBlock	feemarket → evm → staking → noorsignal → gov
EndBlock	staking → gov
InitGenesis	auth → bank → staking → gov → evm → feemarket → noorsignal
ExportGenesis	same as InitGenesis
