NOORCHAIN — Phase 4C

Testnet 1.0 — Node Setup, CLI Commands & Network Specification
Version 1.1 — Architecture Only (No Code)

1. Purpose of This Document

This document defines the complete specification required to:

initialize the NOORCHAIN Testnet 1.0 node

create configuration directories

generate keys for genesis and for the validator

create the genesis file and gentx

configure EVM RPC, WebSocket, gRPC, and REST

configure minimum gas prices

configure peers (local-only for Testnet 1.0)

launch the node

verify correct network behavior

This file completes Phase 4C and will guide the implementation of the real commands later.

2. Testnet 1.0 Node Specification

Chain ID: noorchain-testnet-1
Binary: noord

Directory Structure

~/.noord/config

~/.noord/data

Default Ports (Ethermint-compatible)

JSON-RPC: 8545

WebSocket: 8546

gRPC: 9090

REST (Legacy): 1317

P2P: 26656

RPC (Tendermint): 26657

Genesis File

~/.noord/config/genesis.json

3. Testnet Initialization Flow

The initialization process must follow this strict order:

Initialize node home directory

Create validator key

Create the five genesis allocation wallets

Add genesis accounts

Add bank balances

Declare validator and generate gentx

Collect gentx into final genesis

Inject PoSS genesis state

Validate the genesis

Start the node

This sequence guarantees deterministic state creation for Testnet 1.0.

4. Key Generation Strategy

Keys are created with:

noord keys add <name>

Required keys:

Foundation

Dev (Sàrl)

PoSS Stimulus

Presale (optional)

PoSS Reserve

Validator operator key

Validator consensus key

Each key produces:

bech32 account address (noor1…)

EVM hex address (0x…)

public key

private key (local use only)

These must be stored in the Phase 4C address registry file.

5. Genesis Account & Balance Injection

Each genesis wallet is added using the relevant genesis-account command (implementation phase).

Wallets and balances (integer-rounded during implementation):

Foundation — 5%

Dev Sàrl — 5%

PoSS Stimulus — 5%

Presale Optional — 5%

PoSS Reserve — 80%

Validator account must also be added with sufficient unur for:

self-delegation

gas usage

The total supply must equal exactly 299,792,458 unur.

6. Staking Validator Initialization

A validator declaration is created through a gentx transaction, containing:

validator operator address

self-delegation amount

commission parameters

consensus public key

description metadata

After generating the gentx, the node executes a gentx collection step to insert the validator into staking genesis.

7. EVM & RPC Configuration

JSON-RPC Configuration

JSON-RPC enabled

RPC address: 0.0.0.0:8545

WebSocket address: 0.0.0.0:8546

Enabled APIs: eth, txpool, web3, net

Minimum Gas Price

Testnet default:
0 unur

EVM Denomination

The EVM module must use:

unur

This is the canonical base denomination defined in Phase 4A.

8. gRPC & REST Configuration

gRPC

Enabled by default on port 9090.
Must remain active for CosmJS, dashboards and future NOORCHAIN dApps.

REST (Legacy)

Must remain enabled for compatibility with common Cosmos tooling.

9. Genesis Validation Requirements

Before launching the node, genesis must be validated through:

noord validate-genesis (implementation phase)

Validation must confirm:

correct account structures

correct supply sum

valid staking entries

valid governance params

EVM genesis state consistency

Feemarket parameters valid

PoSS genesis state valid

no duplicate accounts

no malformed entries

If validation fails → the node must not start.

10. Node Start Specification

Once genesis is valid, the node is started using:

noord start (implementation phase)

Expected behavior:

node begins producing blocks

RPC and WebSocket endpoints become active

gRPC endpoint becomes active

validator shows correct voting power

PoSS module initializes and runs BeginBlock logic (empty state at block 1)

Proper startup confirms the integrity of Phase 4C work.

11. Explorer / Tooling Compatibility

Testnet 1.0 must support:

MetaMask (via JSON-RPC at 8545)

CosmJS (via gRPC at 9090)

NOORCHAIN Dashboard (later)

external block explorers (optional in early testnet)

Ethermint compatibility ensures seamless EVM interactions.

12. Checkpoints Before Confirming Testnet Success

A Testnet 1.0 session is considered valid if:

the node starts without crashing

block height increases continuously

EVM transactions execute successfully

Cosmos bank transfers work

PoSS keeper loads without errors

validator stays bonded

JSON-RPC, gRPC, REST endpoints respond correctly

module parameters load without warnings

Meeting all criteria confirms that Testnet 1.0 is operational.

13. Summary

This document defines:

the full Testnet 1.0 node setup

CLI key and account creation workflows

genesis account and balance initialization

staking validator creation and gentx collection

EVM and RPC configuration

gRPC and REST configuration

genesis validation process

node startup and success criteria

It is the official blueprint for implementing the Testnet 1.0 node during Phase 4C.
