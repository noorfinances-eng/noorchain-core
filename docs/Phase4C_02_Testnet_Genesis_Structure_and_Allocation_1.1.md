NOORCHAIN — Phase 4C

Testnet 1.0 Genesis Structure & Allocation
Version 1.1 — Architecture Only (No Code)

1. Purpose of This Document

This document defines the complete genesis structure for NOORCHAIN Testnet 1.0, including:

module initialization order

total supply distribution

genesis accounts

bank balances

staking validator configuration

PoSS reserve setup

Ethermint EVM genesis configuration

Feemarket configuration

mandatory validation checks before chain launch

This is the authoritative reference for creating testnet/genesis.json during Phase 4C.

2. High-Level Genesis Overview

The genesis file must define:

chain metadata (chain-id, genesis time, app version)

all accounts and keys

balances and total supply

validator set

staking module state

governance module parameters

EVM genesis configuration

Feemarket parameters

PoSS module state

parameters for every Cosmos module

Genesis must be executed in the fixed InitGenesis order defined in Phase 4A:

auth → bank → staking → gov → evm → feemarket → noorsignal

This ordering guarantees deterministic state creation.

3. Genesis Total Supply Distribution

Testnet uses the same total supply as mainnet for deterministic PoSS behavior:

299,792,458 unur

Supply must be allocated as follows:

Allocation	Percentage	Amount (unur)
Foundation	5%	14,989,622.9
Dev Sàrl	5%	14,989,622.9
PoSS Stimulus	5%	14,989,622.9
Presale Optional	5%	14,989,622.9
PoSS Reserve	80%	239,833,966.4

Amounts will be integer-rounded using deterministic truncation in Phase 4C_02.

4. Genesis Accounts Structure

Each of the five economic allocation wallets must appear under:

/cosmos.auth.v1beta1.BaseAccount

Fields required:

bech32 address

public key

account_number = 0

sequence = 0

Additionally, the testnet validator account must also appear.

Minimum number of accounts in genesis:

5 genesis supply wallets + 1 validator wallet = 6 accounts

5. Bank Balances Structure

Each genesis wallet must have:

address: bech32

coins: a list containing at least one entry with:

denom: unur

amount: the allocated supply chunk

The validator must be credited with:

minimum self-delegation

additional tokens for gas and test operations

The sum of all balances must equal the total supply.

6. Staking Genesis State

Staking genesis must include:

staking parameters (unbonding, max validators, etc.)

last total power

last validator power

validators array

delegations array

Primary Validator Entry

The validator must be fully defined with:

operator address

consensus public key

bonding status

tokens

delegator shares

commission parameters

minimum self-delegation

A delegation entry must map the validator wallet to the validator operator address.

7. Governance Genesis State

Governance parameters must include:

voting period

tally parameters

deposit parameters

For Testnet:

shorter voting period is recommended

deposit threshold may be minimal

8. EVM Genesis State (Ethermint)

The EVM module requires initialization similar to an Ethereum genesis:

Required:

chain_id = NOORCHAIN Testnet chain-id

denom = unur

gas limit per block

empty initial storage

empty code database

zeroed nonce mapping

Critical flags:

enable_create = true

enable_call = true

This allows smart contract creation and execution.

9. Feemarket Genesis State (EIP-1559)

The feemarket module must define:

base_fee (often zero in Testnet)

minimum gas price

block gas target

elasticity multiplier

Defaults used in Ethermint/Evmos are acceptable.

10. PoSS Module Genesis State

The PoSS module must initialize:

Parameters

PoSS reserve address

base_reward_per_unit

blocks_per_halving (derived from 8-year cycle)

weight table

max_daily_signals

minimum curator level

Reward State

current_epoch = 0

current_halving_cycle = 0

last_reward_block = 0

Curator Set

list of initial curators with level (Bronze/Silver/Gold)

Stats

total_signals = 0

total_rewards_distributed = 0

PoSS must be fully operational from block 1, even if rewards are disabled initially.

11. Genesis Validation Requirements

Before launching the chain, the following conditions must hold:

All module states are valid

Sum of balances equals total supply

PoSS Reserve holds exactly 80%

All five genesis wallets are present

Validator is fully funded and bonded

EVM chain_id is correct

Feemarket has valid parameters

No duplicate accounts

No coins with zero denom

These checks prevent panics during chain startup.

12. Pre-Launch Checkpoints

Before running noord start, you must:

run noord prepare-genesis (Phase 4C tool)

validate each module state (auth, bank, staking, gov, evm, feemarket, noorsignal)

run noord validate-genesis

export genesis and import again for reproducibility tests

This ensures the Testnet launches deterministically.

13. Summary

This document defines:

the full Testnet 1.0 genesis structure

the 5/5/5/5/80 supply allocation

all required accounts

bank and staking configuration

EVM & feemarket initialization

PoSS module genesis state

mandatory validation rules

It is the official blueprint for generating the actual genesis.json in the next Phase 4C document.
