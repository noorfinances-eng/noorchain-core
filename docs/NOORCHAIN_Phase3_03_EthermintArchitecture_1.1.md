NOORCHAIN 1.0 — Phase 3.03
Ethermint Architecture (EVM Integration)
Version 1.1 — English
Purpose of this Document

This document provides the complete technical specification of the Ethermint (EVM) architecture inside NOORCHAIN.
It defines:

what the EVM layer is,

how Ethermint integrates with the Cosmos SDK,

how Ethereum transactions behave inside NOORCHAIN,

how gas and fees operate (EIP-1559),

how JSON-RPC exposes NOORCHAIN as an EVM-compatible chain.

This document is the official reference for EVM integration before implementation begins in Phase 4.

1. Role of Ethermint in NOORCHAIN

NOORCHAIN integrates Ethermint to provide:

full EVM compatibility

smart contract deployment & execution

MetaMask + EVM wallet compatibility

Web3 JSON-RPC endpoints

EIP-1559 fee market support

Ethereum account & storage model

Ethermint enables NOORCHAIN to function simultaneously as:

a Cosmos SDK chain

an EVM chain

2. High-Level Architecture (Cosmos ↔ Ethermint)

Ethermint is implemented through two Cosmos SDK modules:

2.1 evm module

Executes Ethereum transactions

Manages EVM accounts, nonces, and balances

Stores smart contract bytecode and storage

Emits EVM events/logs

Integrates directly with Cosmos stores via EVMKeeper

2.2 feemarket module

Implements EIP-1559:

Dynamic base fee per block

Priority fees (tips) to validators

Maximum fee constraints

Predictable fee behavior across blocks

Data Flow
CometBFT Consensus  
→ Cosmos BaseApp  
→ evm module  
→ EVM Execution  
→ Cosmos State Stores


Ethermint extends Cosmos SDK without replacing it.

3. EVM Account Model in NOORCHAIN
3.1 Dual-Account Structure

NOORCHAIN supports two parallel account systems:

Cosmos accounts (bech32: noor1…)

Ethereum accounts (0x-prefixed addresses)

Both are representations of the same underlying actor.

3.2 Mapping Rules

Every Cosmos account may have a corresponding EVM account with:

EVM nonce

Contract code (if contract)

Storage root

EVM balance (mirrors Cosmos bank balance)

Handled via:

AccountKeeper

EVMKeeper

3.3 Native Token

The unified token across Cosmos and EVM is:

denom: unur

display: NUR

decimals: 18

Used for:

EVM gas

Cosmos fees

Smart contract execution

4. Ethereum Transaction Lifecycle in NOORCHAIN
4.1 EVM Tx Pipeline

Node receives a signed Ethereum-style tx

AnteHandler (Ethermint) validates:

signature (ECDSA secp256k1)

nonce

gas limits

fee fields (legacy or EIP-1559)

account balance

If valid → added to mempool

During block execution → EVM executes opcodes

State updated in Cosmos KV stores

Gas fees deducted in unur

Receipt produced (Ethereum format)

Exposed via JSON-RPC

4.2 Supported Ethereum Tx Types

Ethermint supports:

Legacy (gasPrice)

EIP-1559 transactions

EIP-2930 access list transactions

4.3 Nonce Management

Each EVM account keeps its own evm_nonce, separate from Cosmos sequence.

5. EVM Execution Environment
Components

Ethermint provides:

Opcode interpreter

Precompiled contracts

State database

Memory management & expansion costs

Log emitter

(Optional) JIT optimizations

Smart Contract Support

Compatible with:

Solidity

Vyper

Hardhat / Foundry / Remix

EVM bytecode

ABI encoding

Gas Accounting

Gas follows Ethereum rules:

opcode cost table

memory expansion

storage read/write charges

EIP-1559 base + priority fee

Always paid in unur.

6. Fee Market Module (EIP-1559)
6.1 Base Concepts

BaseFee (dynamic)

PriorityFee (tip)

MaxFeePerGas

MaxPriorityFeePerGas

6.2 Base Fee Adjustment

Automatically adjusted per block based on gas usage.

6.3 Cosmos Integration

Affects:

AnteHandler validation

EVM execution costs

Validator rewards (priority fees)

6.4 Gas Denomination

All EVM gas is paid in:

unur

7. JSON-RPC Endpoints (Web3 Compatibility)

Ethermint exposes full Ethereum RPC API:

7.1 Supported Namespaces

eth_*

web3_*

net_*

txpool_*

debug_* (optional)

personal_* (optional)

This enables:

MetaMask

Ledger

WalletConnect

Foundry & Hardhat

Block explorers

7.2 Response Format

Ethereum-compatible:

block hashes

receipts

logs

bloom filters

gas fields

7.3 Node Identity

Example RPC identification:

client: "Noorchain-Ethermint"
protocolVersion: evmos-compatible
networkVersion: <chain_id>

8. Cosmos Tx vs EVM Tx
Aspect	Cosmos Tx	EVM Tx
Format	Protobuf	RLP
Signature	Cosmos secp256k1	Ethereum ECDSA
Multi-msg	Yes	No
Gas model	Cosmos	Ethereum / EIP-1559
Execution	Module handler	EVM bytecode VM
9. State Storage Model

Ethermint stores EVM state inside Cosmos stores.

EVM Storage Includes:

Account balance / nonce

Code hash

Contract bytecode

Storage key-value pairs

Logs

Merkle Proof Support

Ethermint maintains Ethereum-compatible Merkle proof structure for:

RPC proof queries

Light client verification

Explorer compatibility

10. Consensus Integration (CometBFT)

CometBFT handles:

transaction ordering

block finalization

networking

hashing state results

EVM execution happens during:

DeliverTx

Minor handling in BeginBlock/EndBlock

Consensus does not execute the EVM itself.

11. Security Model of the EVM Layer
11.1 DoS Protection

Block gas limit

Per-tx max gas

AnteHandler checks

Mempool validation

11.2 Replay Protection

chainID enforcement

strict nonce rules

11.3 Deterministic Execution

All nodes must return identical:

logs

receipts

storage changes

12. Summary (Header Version)

NOORCHAIN 1.0 — Phase 3.03 — Ethermint Architecture

This document defines:

EVM + FeeMarket module integration

Dual-account model

Ethereum-style transaction lifecycle

JSON-RPC compatibility

Gas/fees (EIP-1559)

Execution environment

Cosmos ↔ EVM store interaction

CometBFT integration

Security model

This is the authoritative reference for implementing EVM support before Testnet 1.0.
