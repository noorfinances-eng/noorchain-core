NOORCHAIN — Phase 4C

Testnet 1.0 — Addresses & Keys Plan
Version 1.1 — Architecture Only (No Code)

1. Purpose of This Document

This document defines the complete address and key strategy for NOORCHAIN Testnet 1.0.
It specifies:

the five mandatory genesis wallets

the bech32 and EVM address formats

the validator keys

the genesis account models

the rules for deterministic key generation

the mapping between Cosmos and Ethermint formats

This specification is required before creating the real genesis during Phase 4C_02.

2. Address Format Requirements

NOORCHAIN uses the following Bech32 prefixes (defined in Phase 2):

Account prefix: noor1…

Validator operator: noorvaloper1…

Validator consensus: noorvalcons1…

Public keys: noorpub1…

Dual Address Model (Ethermint)

Every NOORCHAIN account has:

a native Cosmos Bech32 address

an EVM-compatible hexadecimal address (0x…)

Every account used in Testnet must provide:

bech32 address

EVM hex address

public key

private key

This dual format is required for full RPC/EVM compatibility.

3. Required Genesis Wallets (5-Wallet Model)

Testnet 1.0 requires exactly five real wallets, in accordance with the official NOORCHAIN Genesis Economic Model (5% / 5% / 5% / 5% / 80%).

1. Foundation Address — 5%

Used for public infrastructure, transparency, documentation and operations.

2. Dev / Sàrl Address — 5%

Used for development, tooling, integrations and internal operations.

3. PoSS Stimulus Address — 5%

Used for early PoSS incentives during testing phases.

4. Presale Optional Address — 5%

Reserved for private investors, with vesting in Mainnet 1.0.

5. PoSS Reserve Address — 80%

The primary supply holder for PoSS rewards.
This address must appear in:

x/noorsignal params

genesis.json balances

bank module total supply

Mandatory Requirements

Each of the 5 wallets must:

be a valid bech32 account

have a valid EVM hex address

be stored as a BaseAccount in genesis

be funded with the correct NUR allocation

be referenced by PoSS parameters

be deterministic between exports

These five addresses appear in:

genesis.json → auth.accounts

genesis.json → bank.balances

genesis_distribution.json

x/noorsignal parameters

4. Validator Keys Requirements

Testnet 1.0 requires one validator, consisting of:

a validator operator key (noorvaloper1…)

a consensus key (noorvalcons1…)

a self-delegation account

The validator must be included in:

staking.validators

staking.delegations

auth.accounts

Validator requirements:

must hold sufficient NUR for self-delegation

must define minimum gas prices

must be compatible with Ethermint RPC

5. Address Generation Rules

There are three ways to generate the required keys.

Option A — Using noord keys add (recommended)

Automatically produces:

bech32 address

EVM address

public key

private key

JSON exportable keyfile

This is the recommended deterministic method for Phase 4C.

Option B — Offline Generation

Possible using:

ethermintd keys

gaiad keys

cosmjs or ethermint-compatible scripts

Option C — Importing Existing Keys

Allowed, as long as:

the address is valid

the keys are accessible

the bech32 and EVM formats match

6. Address Mapping Specification

Each address must have:

Field	Description
bech32	Example: noor1…
evm_hex	Example: 0x…
public_key	secp256k1
private_key	kept offline
account_number	set by auth module
sequence	starts at 0

These fields are part of the Cosmos BaseAccount model.

7. Genesis Accounts Format

In genesis.json → auth.accounts, each of the five genesis wallets plus the validator account must appear as BaseAccounts.

Required fields:

bech32 address

public key (optional for some testnet accounts)

account_number = "0"

sequence = "0"

This ensures deterministic behavior during initial blocks.

8. Genesis Balances Format

In genesis.json → bank.balances, each of the five genesis wallets must have:

address: bech32

coins: list of denominations

Example structure:

denom: unur

amount: exact integer allocation

The amounts must follow:

Testnet supply = Mainnet supply

5% Foundation

5% Dev

5% Stimulus

5% Presale

80% PoSS Reserve

This keeps Testnet deterministic for PoSS simulations.

9. Testnet Address Integrity Rules

All five genesis wallets must:

appear in both auth.accounts and bank.balances

contain the correct NUR amounts

appear in PoSS parameters

be exportable using “noord export” (deterministic output)

Validator keys must:

match the staking validator entry

match the consensus public key

appear in auth.accounts

correspond to the operator key

10. Summary

This document defines:

the 5 mandatory genesis wallets

the validator key set

bech32 + EVM dual address format

genesis account structure

genesis balances structure

deterministic key generation rules

integration with PoSS parameters

This file is the official prerequisite for Phase 4C_02 — Testnet Genesis Construction.
