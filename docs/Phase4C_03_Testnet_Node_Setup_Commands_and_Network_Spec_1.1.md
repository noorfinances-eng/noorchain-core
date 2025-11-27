Phase4C_03_Testnet_Node_Setup_Commands_and_Network_Spec_1.1.md
**NOORCHAIN — Phase 4C

Testnet 1.0 — Node Setup, CLI Commands & Network Specification**
Version 1.1 — Architecture Only (No Code)

🔧 1. Purpose of This Document

This document provides the full plan and specification required to:

initialize the NOORCHAIN Testnet 1.0 node

create directories

generate validator keys

generate genesis & gentx

configure EVM RPC / gRPC / REST

configure minimum gas prices

configure peers (local-only for now)

start the Testnet node

verify correct boot

This is the final file of Phase 4C and will guide the coding + real commands execution later.

🧩 2. Testnet 1.0 Node Specification
Chain ID
noorchain-testnet-1

Binary
noord

Directories
~/.noord/config
~/.noord/data

Ports (default Ethermint-compatible)

JSON-RPC: 8545

WebSocket: 8546

gRPC: 9090

REST (Legacy): 1317

P2P: 26656

RPC (Tendermint): 26657

Genesis File
~/.noord/config/genesis.json

🏗️ 3. Testnet Initialization Flow

The node setup for Testnet 1.0 follows these steps:

1. Initialize node (home folder)
2. Create validator key
3. Add genesis accounts (5+1)
4. Add balances
5. Add staking validator declaration
6. Generate and collect gentx
7. Inject PoSS genesis state
8. Validate genesis
9. Start node


These steps will be executed in the coding phase (4C Implementation).

🔐 4. Key Generation Strategy

Keys created via:

noord keys add <name>


For each of the 5 genesis wallets:

Foundation

Dev

Stimulus

Presale

PoSS Reserve

Plus:

Validator operator key

Validator consensus key

The command outputs:

bech32 address (noor1…)

hex address (0x…)

public key

private key (for local use only)

Keys must be copied into Phase4C_01 addresses file.

🗂️ 5. Genesis Account & Balance Injection
Genesis Account

Using command later:

noord add-genesis-account <address> <amount>unur


Repeat for:

5 genesis wallets

validator account

Genesis Balances

Amounts must match the exact distribution:

Wallet	%	Amount
Foundation	5	14,989,622.9
Dev Sàrl	5	14,989,622.9
Stimulus	5	14,989,622.9
Presale	5	14,989,622.9
PoSS Reserve	80	239,833,966.4

Amounts truncated to integer during implementation.

🏦 6. Staking Validator Initialization

Using:

noord gentx <keyname> 1000000000unur --chain-id noorchain-testnet-1


Then collect:

noord collect-gentxs


This will create:

validator operator entry in staking genesis

self-delegation entry

consensus pubkey entry

🚀 7. EVM & RPC Configuration
Enable JSON-RPC

In config file:

json-rpc:
  enable: true
  address: "0.0.0.0:8545"
  ws-address: "0.0.0.0:8546"
  api: "eth,txpool,web3,net"

Minimum Gas Price

Set minimum gas price in app config:

minimum-gas-prices = "0unur"

EVM Denom

Must be "unur" as defined in Phase 4A.

🌐 8. gRPC & REST Configuration
gRPC

Enabled by default on 9090.

REST

Legacy REST should stay enabled for completeness.

🧪 9. Genesis Validation (Before Node Start)

Run the following (in implementation phase):

noord validate-genesis


This must validate:

auth

bank

staking

gov

evm

feemarket

noorsignal

If invalid → Testnet cannot start.

🔥 10. Node Start Specification

Once genesis is valid:

noord start


Expected output:

block height increasing

EVM JSON-RPC ready

gRPC ready

PoSS BeginBlock logic runs (empty at block 1)

validator voting power appears in Tendermint output

📡 11. Explorer / Tooling Compatibility

Testnet 1.0 must support:

MetaMask (via JSON-RPC)

CosmJS (via gRPC)

NOORCHAIN Dashboard (later)

Block explorers (optional at first)

Nothing special required; Ethermint provides auto-compatibility.

📌 12. Checkpoints Before Confirming Testnet Success

A Testnet 1.0 is considered valid if:

✔️ Node starts without crash
✔️ Chain produces blocks
✔️ EVM tx works (eth_call test)
✔️ bank tx works (cosmos send test)
✔️ PoSS store loads
✔️ PoSS keeper initialized
✔️ validator is bonded
✔️ noord status returns correct info
✔️ RPC endpoints live
🎯 13. Summary

This document defines:

✔️ full Testnet node setup
✔️ full command flow
✔️ full RPC/gRPC/EVM configuration
✔️ validator init
✔️ genesis account injection
✔️ genesis validation
✔️ node start sequence
✔️ success criteria

It is the official guide for implementing Testnet NOORCHAIN 1.0 during coding.