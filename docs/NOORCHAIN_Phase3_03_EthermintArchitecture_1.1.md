NOORCHAIN 1.0 — Phase 3.03
Ethermint Architecture (EVM Integration) — Version 1.1

Purpose of this document
Provide the full technical specification of the Ethermint (EVM) architecture inside NOORCHAIN.
This document defines what the EVM layer is, how it integrates with Cosmos SDK, how EVM transactions work, how JSON-RPC works, and how gas/fees behave.

NOORCHAIN is an EVM-enabled Cosmos chain; this document formalizes the EVM component before any code is written.

1. Purpose of Ethermint in NOORCHAIN

NOORCHAIN integrates Ethermint to provide:

full EVM compatibility

deployment and execution of smart contracts

compatibility with MetaMask and all EVM wallets

Web3 RPC endpoints

EIP-1559-style fee market

Ethereum-style account & storage model

Ethermint enables NOORCHAIN to function both as a Cosmos chain AND an EVM chain, simultaneously.

2. High-Level Architecture (Cosmos ↔ Ethermint)

Ethermint is implemented as two Cosmos SDK modules:

evm module

handles Ethereum transaction execution

manages EVM accounts, nonces, storage, and bytecode

provides EVM events & logs

integrates with the Cosmos state store via keepers

feemarket module

implements EIP-1559 dynamic gas pricing

manages base fee, tip, priority fees

ensures predictable gas behavior across blocks

These modules live entirely inside the Cosmos SDK application.

High-level flow:

Cosmos Node (NOORCHAIN)
→ CometBFT Consensus
→ BaseApp (Cosmos SDK)
→ evm module
→ EVM execution
→ Cosmos state storage

Ethermint does not replace Cosmos; it extends it.

3. EVM Account Model in NOORCHAIN
3.1. Dual Account System

NOORCHAIN supports:

Cosmos accounts (bech32, e.g., noor1...)

EVM accounts (0x addresses, Ethereum-style)

Both represent the same underlying state, but with different views.

3.2. Mapping Rules

A Cosmos account can have an associated EVM account with:

Ethereum nonce

EVM balance (mirrors BankKeeper)

Contract bytecode

Contract storage

The mapping is deterministic and stored through:

AccountKeeper (Cosmos)

EVMKeeper (Ethermint)

3.3. Native Token

The native token is:

denom: unur
display: NUR
decimals: 18


It is used for:

Cosmos fees

EVM gas

contract execution costs

4. Ethereum Transaction Lifecycle inside NOORCHAIN

EVM transactions follow the Ethereum transaction model, but run inside Cosmos.

4.1. EVM Tx Pipeline (simplified)

The node receives an Ethereum-style signed transaction.

The AnteHandler (EVM path) validates:

signature (ECDSA secp256k1)

nonce

gas limit

gas price or EIP-1559 fields

account balance

Transaction enters mempool if valid.

During block execution (DeliverTx):

VM executes opcode-by-opcode

state changes applied to EVM storage

logs and events generated

Gas fees deducted in unur

Receipt generated & exposed via JSON-RPC

4.2. Supported Ethereum Transaction Types

Ethermint supports:

Legacy gasPrice transactions

EIP-1559 dynamic fee transactions

EIP-2930 access list transactions

4.3. Nonce Management

Each EVM account maintains:

evm_nonce (separate from cosmos sequence)


This ensures strict Ethereum compatibility.

5. EVM Execution Environment
5.1. Components

The Ethermint EVM environment includes:

Opcode interpreter

Precompiled contracts

EVM state database

JIT optimizations (optional)

Log & event emitter

5.2. Smart Contract Support

Fully compatible with:

Solidity

Vyper

EVM bytecode

ABI encoding

Ethereum tooling (Hardhat, Foundry, Remix, etc.)

5.3. Gas Accounting

Gas is charged in unur.

Gas rules follow Ethereum logic:

Opcode cost table

Memory expansion costs

Storage read/write costs

Base fee + priority fee (EIP-1559)

6. feemarket Module (EIP-1559)

The feemarket module implements Ethereum's fee market:

6.1. Key Concepts

Base Fee (dynamic per block)

Priority Fee (tip to validator)

MaxFeePerGas

MaxPriorityFeePerGas

6.2. Base Fee Adjustment

Base fee adjusts depending on gas usage of previous blocks.

6.3. Cosmos Integration

Base fee affects:

antehandler checks

EVMKeeper execution logic

reward distribution (tip → validator)

6.4. Gas Denom

Always paid in:

unur

7. JSON-RPC Endpoints (Web3 Compatibility)

Ethermint exposes standard Ethereum RPC endpoints, enabling:

MetaMask

Ledger

Web3.js / Ethers.js

Foundry

Hardhat

Third-party explorers

7.1. Supported JSON-RPC APIs

The main namespaces include:

eth_* (full Ethereum API)

web3_*

net_*

debug_* (optional)

txpool_*

personal_* (optional)

7.2. Block & Receipt Format

Responses follow Ethereum structures:

block hashes

receipts

logs

bloom filters

gas fields

7.3. EVM Node Identity

NOORCHAIN will identify itself via:

client: "Noorchain-Ethermint"
protocolVersion: evmos-compatible
networkVersion: chainid (set in Phase3_05)

8. Differences Between Cosmos Transactions and EVM Transactions
8.1. Cosmos Transactions

protobuf messages

multiple msgs in a single tx

signed via Cosmos secp256k1

processed through Cosmos antehandler

stored in ABCI format

8.2. EVM Transactions

RLP encoded

single action per tx

Ethereum-style signature

gasPrice or EIP-1559 fields

executed in evm module

8.3. Key Distinctions
Aspect	Cosmos Tx	EVM Tx
Format	Protobuf	RLP
Signature	Cosmos mode	Ethereum mode
Multi-Msgs	Yes	No
Gas Model	Cosmos	Ethereum/EIP-1559
Execution	Module handler	EVM VM
9. State Storage Layout (Cosmos ↔ EVM)

Ethermint stores Ethereum state inside Cosmos SDK stores.

9.1. EVM Storage Components

Account → balance, code hash, storage root

Storage → key-value pairs

Nonce → per EVM account

Bytecode → stored separately

Logs → block-level logs

9.2. Merkle Path

Ethermint maintains Ethereum-compatible Merkle proofs, enabling:

contract verification

RPC proof queries

interoperability with Ethereum tools

10. Consensus Integration (CometBFT)

Ethermint execution occurs during:

DeliverTx

BeginBlock / EndBlock (minor logic)

Consensus does NOT execute EVM; it only:

orders transactions

hashes application state

finalizes blocks

11. Security Model of the EVM Layer
11.1. DoS Protection

gasLimit per block

maxTxGas

mempool checks

11.2. Replay Protection

chainID enforced

strict nonce rules

11.3. Deterministic Execution

All nodes must produce identical:

logs

receipts

storage state

12. Summary (for top-of-file header)

NOORCHAIN 1.0 — Phase3_03 — Ethermint Architecture (EVM)
This document defines the full Ethermint integration:

evm + feemarket modules

EVM account model

Ethereum transaction lifecycle

gas & fees (EIP-1559)

JSON-RPC compatibility

execution environment

interaction with Cosmos stores and keepers

It is the reference document for implementing the EVM layer before Testnet 1.0.