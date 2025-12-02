# NOORCHAIN 1.0 — Minimal Local Testnet & PoSS Status (v1.1)

**Scope**  
This document describes the current minimal local testnet setup for NOORCHAIN 1.0, and how it relates to the PoSS module (`x/noorsignal`) in its current “economic OFF” state.

It completes:

- `NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md`
- `NOORCHAIN_Phase3_07_PoSS_Params_and_KeeperTests_1.1.md`

---

## 1. Minimal local testnet structure

A very simple filesystem-based testnet is now available in the repository:

- Directory: `data-testnet/` (generated)
- Genesis template: `testnet/genesis.json`
- Distribution placeholder: `testnet/genesis_distribution.json`
- Helper script: `scripts/testnet.sh`

The goal of this setup is **not** to run a full production-like network, but to:

- provide a deterministic directory layout (`$HOME/config/genesis.json`),
- keep the PoSS genesis state consistent with our specs,
- prepare the ground for future `noord start` / CLI integration.

---

## 2. Genesis template for PoSS (testnet/genesis.json)

The file:

- `testnet/genesis.json`

currently contains a **minimal app_state** with the modules we care about at this stage:

- `auth`  
- `bank`  
- `staking`  
- `evm`  
- `feemarket`  
- `params`  
- `noorsignal` (PoSS)

Example (simplified):

```json
{
  "genesis_time": "2025-01-01T00:00:00Z",
  "chain_id": "noorchain-testnet-1",
  "initial_height": "1",
  "app_hash": "",
  "app_state": {
    "auth": {
      "accounts": []
    },
    "bank": {
      "params": {
        "default_send_enabled": true
      },
      "balances": [],
      "supply": [],
      "denom_metadata": []
    },
    "staking": {
      "params": {
        "bond_denom": "unur"
      }
    },
    "evm": {},
    "feemarket": {},
    "params": {},
    "noorsignal": {
      "total_signals": "0",
      "total_minted": "0",
      "max_signals_per_day": "20",
      "participant_share": "70",
      "curator_share": "30"
    }
  }
}
2.1. PoSS genesis section
The noorsignal section is aligned with GenesisState:

total_signals = "0"

total_minted = "0" (in unur, as a string)

max_signals_per_day = "20"

participant_share = "70"

curator_share = "30"

This is consistent with the structural PoSS rule:

70 % to participant

30 % to curator

and a first anti-abuse limit (max_signals_per_day).

In code, this is handled by:

x/noorsignal/types/genesis.go

x/noorsignal/keeper/keeper.go (InitGenesis / ExportGenesis)

3. Distribution placeholder (genesis_distribution.json)
The file:

testnet/genesis_distribution.json

is a placeholder for the future 5-address economic allocation:

5 % Foundation (NOOR Foundation)

5 % Dev Sàrl (Noor Dev)

5 % PoSS Stimulus

5 % Optional pre-sale

80 % PoSS mintable supply

Structure (simplified):

json
Copier le code
{
  "foundation": {
    "address": "",
    "allocation": "0"
  },
  "dev": {
    "address": "",
    "allocation": "0"
  },
  "stimulus": {
    "address": "",
    "allocation": "0"
  },
  "presale": {
    "address": "",
    "allocation": "0"
  },
  "poss": {
    "address": "",
    "allocation": "0"
  }
}
At this stage:

All addresses are empty strings ("").

All allocations are "0".

Later, for NOORCHAIN Testnet 1.0, the 5 real Bech32 addresses will be injected here and in the final genesis.json / genesis_distribution.json pairing.

4. Testnet script (scripts/testnet.sh)
The helper script:

scripts/testnet.sh

is a very small tool that:

Removes any previous data-testnet/ directory.

Recreates data-testnet/config/.

Copies testnet/genesis.json into data-testnet/config/genesis.json.

Script content (summary):

bash
Copier le code
#!/usr/bin/env bash

set -e

echo "🔧 Initializing NOORCHAIN local testnet (filesystem only)..."

CHAIN_DIR="./data-testnet"

rm -rf "$CHAIN_DIR"
mkdir -p "$CHAIN_DIR/config"

cp testnet/genesis.json "$CHAIN_DIR/config/genesis.json"

echo "✅ Testnet directory initialized in $CHAIN_DIR"
echo "  - genesis.json copied to $CHAIN_DIR/config/genesis.json"
echo
echo "👉 Next step (later):"
echo "   ./noord start --home $CHAIN_DIR"
4.1. How to use
From the repository root:

bash
Copier le code
chmod +x scripts/testnet.sh   # once
scripts/testnet.sh
Result:

data-testnet/config/genesis.json is always in sync with testnet/genesis.json.

No dependency on noord init for now.

Safe to run multiple times (always resets the local testnet directory).

5. PoSS state on this minimal testnet
With the current code and genesis:

PoSS module is initialized via InitGenesis (AppModule).

GenesisState PoSS starts with:

TotalSignals = 0

TotalMinted = "0"

5.1. Params behaviour
On this minimal testnet:

PoSS Params are not explicitly set in genesis.json.

On first use, the keeper’s GetParams(ctx) will:

detect that the ParamSubspace is empty,

write DefaultParams() into the Subspace,

return the defaults.

Therefore:

PoSSEnabled = false by default,

Limits and weights are loaded from DefaultParams,

The system is economically OFF, but structurally ready.

5.2. Signals and rewards
If a PoSS signal is processed internally via:

Keeper.ProcessSignalInternal(...)

then:

TotalSignals will increase by 1,

TotalMinted will increase by participantReward + curatorReward,

daily counters will increment for the participant,

BUT with PoSSEnabled = false, both rewards are 0 unur:

participantReward = 0 unur

curatorReward = 0 unur

So TotalMinted stays "0".

This matches the Legal Light / Safe OFF strategy:

PoSS logic and accounting are live,

real NUR minting will only happen later, once:

PoSSEnabled = true,

a PoSS reserve account exists,

Bank/mint wiring and daily limits are enforced.

6. Next steps (beyond this document)
This minimal testnet is a bridge between:

Phase 3 — Documentation & Specs

Phase 4 — Implementation (coding, PoSS module, tests)

Phase 6 — Genesis Pack & Communication

Next planned steps:

Introduce a proper noord start flow based on data-testnet/ and the Ethermint/Cosmos configuration.

Extend test scenarios:

CLI / gRPC to inspect:

PoSS Params

Global PoSS stats (TotalSignals, TotalMinted)

Simulate PoSS signals in integration tests.

When legal and economic conditions are satisfied:

enable PoSS (PoSSEnabled = true),

wire a real PoSS reserve (module account) with Bank/mint,

enforce PoSS daily limits and halving.

Until then, this minimal testnet remains a safe playground:

no real NUR minting,

no economic impact,

but a fully defined place to observe and validate PoSS behaviour.