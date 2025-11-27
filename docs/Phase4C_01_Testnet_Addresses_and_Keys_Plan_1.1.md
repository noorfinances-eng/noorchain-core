**NOORCHAIN — Phase 4C

Testnet 1.0 — Addresses & Keys Plan**
Version 1.1 — Architecture Only (No Code)

🔧 1. Purpose of This Document

This document defines:

the address strategy for Testnet 1.0

the five required bech32 addresses (foundation, dev, stimulus, presale, PoSS Reserve)

the validator keys

the genesis account format

the rules for address generation

the mapping to Cosmos/EVM formats

This is the first step before creating the real genesis in Phase 4C_02.

🧩 2. Address Format Requirements

NOORCHAIN uses:

Bech32 Prefixes

As defined in Phase 2:

Account prefix: noor1...

Validator operator: noorvaloper1...

Consensus key: noorvalcons1...

Public keys: noorpub1...

EVM Address Mapping

Ethermint creates a dual-address model:

Native format (bech32)

EVM-compatible address (0x…)

Every account used in Testnet must support:

Cosmos bech32 address

EVM hex address

Public key

Private key

🏛️ 3. Required Genesis Wallets

Testnet 1.0 requires exactly 5 real wallets, consistent with the official NOORCHAIN Genesis Economics (5/5/5/5/80).

1️⃣ Foundation Address (5%)

Holds funds for chain infrastructure and public transparency.

2️⃣ Dev / Sàrl Address (5%)

Used for internal development operations.

3️⃣ PoSS Stimulus Address (5%)

Funds early PoSS incentives.

4️⃣ Presale Optional Address (5%)

Reserved for future private investors (timelocked in mainnet).

5️⃣ PoSS Reserve (80%)

Main PoSS reward distribution supply.

🔥 VERY IMPORTANT

These 5 addresses MUST be:

valid bech32

valid hex (Ethermint)

indexable by EVM module

stored in genesis as BaseAccounts

funded with full NUR distribution

These 5 addresses will appear in:

genesis.json → auth.accounts

genesis.json → bank.balances

genesis_distribution.json

x/noorsignal params

🔐 4. Validator Keys Requirements

Testnet 1.0 will have:

1 local validator, using:

1 validator operator key → noorvaloper1…

1 consensus key → noorvalcons1…

1 self-delegation account

The validator must be:

included in staking.validators

funded with enough NUR to self-delegate

configured with minimum gas prices

connected to Ethermint RPC

🧱 5. Address Generation Rules

To create the 5 genesis wallets + validator keys:

Option A — Using noord keys add

(Recommended during Phase 4C implementation)

Generates:

bech32 account
hex EVM address
public key
private key

Option B — Offline

Using:

ethermintd keys

gaiad keys

cosmjs script

Option C — Imported Keys

If you want to reuse existing addresses, you can.

🔄 6. Address Mapping Specification

Each address must have:

Field	Description
bech32	noor1…
evm_hex	0x…
public_key	secp256k1
private_key	for local use only
account_number	set by auth module
sequence	starts at 0

These fields are defined in auth.BaseAccount.

🗂️ 7. Genesis Accounts Format

In genesis.json → auth.accounts, we must insert:

{
  "@type": "/cosmos.auth.v1beta1.BaseAccount",
  "address":   "noor1...",
  "pub_key":   { … },
  "account_number": "0",
  "sequence": "0"
}


For all 5 genesis wallets + validator self-delegation key.

💰 8. Genesis Balances Format

In genesis.json → bank.balances:

{
  "address": "noor1…",
  "coins": [
    { "denom": "unur", "amount": "XXXXX" }
  ]
}


The amount is determined by:

Genesis economic split (5/5/5/5/80)

Testnet total supply (same as mainnet supply for deterministic PoSS behavior)

🧠 9. Testnet Address Integrity Rules

All 5 genesis wallets must:

appear in auth.accounts

appear in bank.balances

contain the correct NUR amounts

be referenced by x/noorsignal params

be deterministic between runs

be exportable via noord export

Validator keys must:

match consensus key in validators

match operator key in staking

match account in auth.accounts

📌 10. Summary

This document defines:

✔️ 5 required genesis wallets
✔️ validator keys requirements
✔️ bech32 + EVM dual format
✔️ genesis account structure
✔️ genesis balance structure
✔️ address generation strategy
✔️ integration with PoSS & Testnet

This file is the official prerequisite for creating Testnet 1.0 genesis in the next document