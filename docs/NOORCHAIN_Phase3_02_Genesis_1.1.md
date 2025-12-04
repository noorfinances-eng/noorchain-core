NOORCHAIN 1.0 — Genesis Structure
Phase 3 Documentation – Version 1.1
Language: English
1. Genesis Philosophy

The genesis is the foundational blueprint of the NOORCHAIN blockchain.
It defines the initial state of the network and contains:

Initial supply distribution

Economic allocation model (5 / 5 / 5 / 5 / 80)

Chain identity and metadata

Initial validator set (if applicable)

PoSS module genesis state (initial empty state)

The genesis must be immutable, deterministic, and fully reproducible.

2. Final NOORCHAIN Genesis Allocation (Official Model)
Allocation	% of Total Supply	Purpose
Foundation	5%	Public-good operations, transparency, governance, communication
Noor Dev Sàrl	5%	Development, R&D, infrastructure, salaries (vested)
PoSS Stimulus	5%	Bootstrap community actions and early Proof of Signal Social activity
Pre-Sale (optional)	5%	Private Swiss investors, with mandatory vesting and documentation
PoSS Mining Reserve	80%	Long-term PoSS distribution with 8-year halving
Notes

Total supply is 299,792,458 NUR (speed of light).

No minting beyond this amount is allowed.

None of the allocations grant financial rights or yield.

All allocations must align with Swiss Legal Light CH.

Address Injection

The 5 official Bech32 addresses must be added during Phase 6 (Genesis Pack):

Foundation

Noor Dev Sàrl

PoSS Stimulus

Pre-Sale

PoSS Reserve

They must be synchronized across:

genesis.json

genesis_distribution.json

x/noorsignal/types/addresses.go

Phase 6 Genesis Pack documents

Compliance and governance documents

Whitepaper 1.1

Website pages /genesis and /governance

3. Required Genesis Files for Testnet

The Testnet requires the following files:

3.1 genesis.json

The full state of the blockchain at block 0.
Includes all modules: auth, bank, staking, gov, evm, feemarket, noorsignal.

3.2 genesis_distribution.json

Defines the distribution structure:

supply allocation

addresses and balances

vesting schedules (if applied)

3.3 genesis_addresses.json

Stores the 5 official addresses for Testnet,
and later for Mainnet.

Used for automated genesis generation.

4. NOOR Signal (PoSS) — Current Genesis Structure (Phase 3 Skeleton)

The NOORCHAIN PoSS module is defined but empty at Phase 3.

Current expected structure in genesis.json:
{
  "noorsignal": {
    "params": {},
    "signals": [],
    "curators": []
  }
}

Notes

No parameters are active yet.

No signals or curators exist at genesis.

Halving, counters, and PoSS-enabled flag will be added in Phase 4–5.

The structure must remain stable to guarantee deterministic genesis.

5. Evolution in Later Phases
Phase 4

Core PoSS parameters

Default config

Reward logic fields

Halving index

Query/Msg genesis handlers

Phase 5

Governance rules

Legal consistency checks

Validation of PoSS genesis integrity

Phase 6 (Genesis Pack)

Address injection

Supply finalization

Distribution proofs

Public genesis documentation

6. Executive Summary

NOORCHAIN uses a fixed-supply, transparent genesis allocation model.

Five mandatory allocation buckets (5/5/5/5/80).

PoSS module genesis is currently minimal and expands in later phases.

Genesis files must remain synchronized across code, documentation, and governance.

7. Status

Validated:
This document defines the official Genesis Structure for Phase 3 and must not be altered without explicit governance approval.
