**NOORCHAIN — Phase 4C

Testnet 1.0 Genesis Structure & Allocation**
Version 1.1 — Architecture Only (No Code)

🔧 1. Purpose of This Document

This document defines the complete genesis layout for NOORCHAIN Testnet 1.0, including:

module ordering

initial supply distribution

genesis accounts

genesis balances

staking validator setup

PoSS reserve initialization

Ethermint EVM genesis config

feemarket genesis config

required checks before starting the chain

This specification is the canonical reference for creating testnet/genesis.json during Phase 4C implementation.

🧩 2. High-Level Genesis Overview

The genesis file must define:

Chain info (chain-id, time, app version)

Accounts & keys

Bank balances

Staking validator

Governance parameters

EVM genesis state

Feemarket parameters

PoSS module state

Module parameters for all Cosmos modules

Supply distribution according to 5/5/5/5/80

Genesis must follow the strict InitGenesis order defined in Phase 4A:

auth → bank → staking → gov → evm → feemarket → noorsignal

🏛️ 3. Genesis Total Supply Distribution

The Testnet uses the same supply as mainnet for deterministic PoSS behavior:

299,792,458 unur

Distribution:
Allocation	Percentage	Amount (unur)
Foundation	5%	14,989,622.9
Dev Sàrl	5%	14,989,622.9
PoSS Stimulus	5%	14,989,622.9
Presale Optional	5%	14,989,622.9
PoSS Reserve	80%	239,833,966.4

Amounts will be integer-rounded in Phase 4C code using truncation.

👛 4. Genesis Accounts Structure

Each of the 5 genesis wallets (plus validator wallet) must appear as:

{
  "@type": "/cosmos.auth.v1beta1.BaseAccount",
  "address": "noor1...",
  "pub_key": { ... },
  "account_number": "0",
  "sequence": "0"
}


Additionally:

The validator operator account must also be present.

Total accounts in Testnet genesis:

5 genesis wallets
+ 1 validator account
= 6 accounts minimum

💰 5. Bank Balances Structure

Each genesis wallet must have:

{
  "address": "noor1…",
  "coins": [
    { "denom": "unur", "amount": "<allocated_amount>" }
  ]
}


The validator must be funded with:

minimum self-delegation

extra tokens for gas fees

🏦 6. Staking Genesis State
Required fields:

params (unbonding, min delegation, etc.)

last total power

last validator power

validators list

delegations

The primary validator must be included:

{
  "operator_address": "noorvaloper1...",
  "consensus_pubkey": { ... },
  "jailed": false,
  "status": "BOND_STATUS_BONDED",
  "tokens": "1000000000",
  "delegator_shares": "1000000000.000000000000000000",
  "description": { ... },
  "unbonding_height": "0",
  "unbonding_time": "0001-01-01T00:00:00Z",
  "commission": { ... },
  "min_self_delegation": "1"
}


Then add delegations:

delegator_address (validator wallet) → validator operator

🗳️ 7. Governance Genesis State

Include:

voting params

tally params

deposit params

Defaults recommended for Testnet:

shorter voting period

minimal deposit requirement

⛽ 8. EVM Genesis State (Ethermint)

Ethermint requires EVM genesis data similar to Ethereum:

Must include:

chain_id = NOORCHAIN Testnet ID

evm.Denom: "unur"

block gas limit

empty or minimal contract storage

empty code database

zeroed nonce mapping

Essential:

"enable_create": true
"enable_call": true


This allows EVM smart contracts to run.

🔥 9. Feemarket Genesis State

Feemarket (EIP-1559) requires:

base_fee (default: "0")

min_gas_price (default: "0")

block gas target

elasticity multiplier

Defaults used by Evmos/Ethermint are compatible.

🌞 10. PoSS Module Genesis State

The PoSS module (x/noorsignal) must initialize:

{
  "params": {
    "poss_reserve_address": "noor1...",
    "base_reward_per_unit": "X",
    "blocks_per_halving": "<calculated_value>",
    "weight_table": [...],
    "max_daily_signals": "X",
    "min_curator_level": "bronze"
  },
  "reward_state": {
    "current_epoch": "0",
    "current_halving_cycle": "0",
    "last_reward_block": "0"
  },
  "curators": [
    { "address": "...", "level": "gold" },
    ...
  ],
  "stats": {
    "total_signals": "0",
    "total_rewards_distributed": "0"
  }
}


PoSS must be functional from block 1.

🧪 11. Genesis Validation Requirements

Before starting the chain:

✔️ All module states must match expected schemas
✔️ Total supply must equal sum of all balances
✔️ PoSS Reserve must have 80% exactly
✔️ 5 genesis wallets must be present
✔️ validator must be funded and bonded
✔️ EVM genesis must have valid chain_id
✔️ Feemarket must have a valid base fee
✔️ no duplicate accounts
✔️ no zero-denom coins

These rules prevent chain panic.

📌 12. Checkpoints Before Launch

Before executing noord start, you MUST:

run noord prepare-genesis (Phase 4C tool)

validate structure:

auth OK

bank OK

staking OK

evm OK

feemarket OK

noorsignal OK

run noord validate-genesis

export and re-import genesis for integrity test

🎯 13. Summary

This document defines:

✔️ full genesis layout
✔️ economic distribution 5/5/5/5/80
✔️ necessary accounts
✔️ staking validator config
✔️ EVM + feemarket config
✔️ PoSS genesis state
✔️ final validation requirements

It is the complete reference for generating the real genesis.json in the next file.