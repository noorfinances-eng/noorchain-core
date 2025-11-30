NOORCHAIN 1.0 — Genesis Structure (Phase 3, English)

Version: 1.1

1. Genesis Philosophy

The genesis defines:

Initial supply distribution

Economic model foundation

Network identity

Initial validators

PoSS genesis state

2. Final NOORCHAIN Genesis Allocation
Allocation	% of Supply	Purpose
Foundation	5%	Public good, maintenance, communication
Noor Dev Sàrl	5%	Development, R&D, salary, infrastructure
PoSS Stimulus	5%	Fuel to bootstrap community signals
Pre-Sale (optional)	5%	Private Swiss investors, with vesting
PoSS Mining Reserve	80%	Long-term PoSS emission with halving 8 years

All addresses must be inserted before Genesis Pack in Phase 6.

3. Required Genesis Files in Testnet

genesis.json

genesis_distribution.json

genesis_addresses.json

4. NOOR Signal Module Genesis (Current skeleton)
{
  "noorsignal": {
    "params": {},
    "signals": [],
    "curators": []
  }
}


This will be populated later in PoSS Logic Phase.