NOORCHAIN_Phase3_08_MinimalTestnet_and_PoSSStatus_1.1.md
NOORCHAIN 1.0 — Minimal Local Testnet & PoSS Status

Version: 1.1
Status: Official Phase 3 Document
Last Updated: 2025-12-03

1. Scope

This document defines the minimal local testnet setup currently available in noorchain-core, and explains how PoSS behaves on this testnet in its economic OFF configuration.

It complements:

NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md

NOORCHAIN_Phase3_07_PoSS_Params_and_KeeperTests_1.1.md

NOORCHAIN_Phase3_04_PoSS_FullSpecification_1.1.md

2. Minimal Local Testnet Structure

A simple filesystem-based testnet is available:

data-testnet/ (generated local chain directory)

testnet/genesis.json (genesis template)

testnet/genesis_distribution.json (economics placeholder)

scripts/testnet.sh (reset & initialize script)

This setup is not intended to run a full node yet.
Its goals are:

provide deterministic directories ($HOME/config/genesis.json),

initialize PoSS genesis state according to specifications,

prepare for future noord start and CLI integration,

serve as a safe Testnet 0 before real Testnet 1.0.

3. Genesis Template (testnet/genesis.json)

The file testnet/genesis.json contains a minimal app_state for:

auth

bank

staking

evm

feemarket

params

noorsignal (PoSS)

3.1 PoSS Genesis Section
"noorsignal": {
  "total_signals": "0",
  "total_minted": "0",
  "max_signals_per_day": "20",
  "participant_share": "70",
  "curator_share": "30"
}


This matches GenesisState from Phase 3:

TotalSignals = 0

TotalMinted = "0" (unur)

MaxSignalsPerDay = 20

Reward split = 70/30 (immutable structural rule)

Handled by:

x/noorsignal/types/genesis.go

x/noorsignal/keeper/keeper.go — Init / Export Genesis

4. Distribution Placeholder (testnet/genesis_distribution.json)

This file prepares the economic distribution for the 5 official NOORCHAIN genesis addresses:

Pool	%	Description
Foundation	5%	NOOR Foundation (non-profit)
Dev Sàrl	5%	Founder / Dev structure
PoSS Stimulus	5%	Curators & early adopters
Pre-Sale	5%	Optional private CH investors
PoSS Reserve	80%	Long-term PoSS emission

Template (simplified):

{
  "foundation": { "address": "", "allocation": "0" },
  "dev":        { "address": "", "allocation": "0" },
  "stimulus":   { "address": "", "allocation": "0" },
  "presale":    { "address": "", "allocation": "0" },
  "poss":       { "address": "", "allocation": "0" }
}


These values will be filled during Phase 6 — Genesis Pack.

5. Testnet Script (scripts/testnet.sh)

Purpose:

delete previous data-testnet/

recreate directory

copy genesis template into data-testnet/config/

Result:

data-testnet/config/genesis.json


Command to run:

chmod +x scripts/testnet.sh
scripts/testnet.sh


This ensures:

consistent genesis state

reproducible directory layout

safe reset for repeated testing

6. PoSS State on the Minimal Testnet
6.1 Genesis Initialization

Upon chain initialization:

InitGenesis loads PoSS section

PoSS state becomes:

TotalSignals = 0
TotalMinted  = "0"

6.2 Params Behaviour

Since PoSS Params are NOT included in genesis.json yet:

the first call to GetParams(ctx) finds the Subspace empty

the keeper writes and returns DefaultParams()

Meaning:

PoSSEnabled = false

PoSS is structurally ON, but economically OFF

Weights, BaseReward, limits all use official defaults

This is exactly the Legal Light safe mode.

7. PoSS Behaviour on Testnet (Current Stage)
7.1 Signals

If PoSS is triggered internally (via tests or direct keeper call):

daily counters increment

TotalSignals increments

reward computation happens theoretically

7.2 Rewards

Because PoSSEnabled = false:

participantReward = 0 unur
curatorReward     = 0 unur


Thus:

TotalMinted remains "0"

No coins are minted

No balances change

This is REQUIRED by Legal Light CH and our technical roadmap.

8. Purpose of This Minimal Testnet

This Testnet 0 exists to:

validate PoSS module behaviour

validate Params persistence

simulate counters, stats, signals

prepare local tooling before real node execution

test JSON serialization, module integration, and state transitions

It is NOT a real blockchain network yet.

Real launch sequence begins in:

Phase 4C → Testnet 1.0 (CLI, app wiring, genesis packing)

Phase 6 → Full Genesis Pack & Documentation

Phase 7 → NOORCHAIN Testnet 1.0

Phase 8 → Testnet expansion

Phase 9 → Public audits / Mainnet readiness

9. Next Steps (After This Document)

Add noord start flow using data-testnet/

Add CLI to inspect PoSS Params (GetParams)

Add CLI to simulate PoSS signals

Implement PoSS Reserve module account

Add Bank transfers when PoSSEnabled = true (later)

Add strict or soft limit enforcement for daily caps

Prepare final Genesis Pack (Phase 6)

10. Summary (Header)

NOORCHAIN 1.0 — Minimal Testnet & PoSS Status (v1.1) defines:

the structure of the minimal local testnet

how genesis.json includes PoSS state

how PoSS Params behave via the Subspace

how the PoSS module behaves in safe OFF mode

the exact relationship with upcoming phases

This is the official reference for internal testing prior to full Testnet 1.0.
