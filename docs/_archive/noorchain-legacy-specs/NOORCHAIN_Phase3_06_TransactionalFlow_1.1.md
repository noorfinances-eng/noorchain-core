NOORCHAIN 1.0 — Phase 3.06
Full Transactional Flow (Cosmos + EVM + PoSS)
Version 1.1 — Official Document
0. Purpose of this Document

This document provides the complete, end-to-end transaction flow for NOORCHAIN 1.0 across:

Cosmos transactions

EVM (Ethermint) transactions

PoSS (x/noorsignal) transactions

and explains how these interact with:

CometBFT consensus

Cosmos SDK BaseApp

Module Manager & Keepers

EVM module

PoSS module

This document represents the high-level runtime architecture of NOORCHAIN and is used for:

implementation validation (Phase 4)

Testnet behavior understanding (Phase 7)

audit preparation

whitepaper & technical documentation

1. Runtime Layer Overview

NOORCHAIN’s runtime is composed of four integrated layers:

1.1 Network & Consensus Layer

P2P networking

CometBFT consensus engine

Block proposal and validation

Finality through BFT voting

CometBFT is responsible for ordering transactions, NOT executing them.

1.2 Application Layer (Cosmos SDK)

BaseApp

Message routing and ante-handling

Module Manager

Keepers and state stores

This is where all state transitions occur.

1.3 Execution Sub-Layers

Cosmos Execution Path (bank, staking, gov…)

EVM Execution Path (Ethermint)

PoSS Execution Path (x/noorsignal)

All execution paths coexist independently inside BaseApp.

1.4 Client API Layer

gRPC

REST (gRPC-gateway)

JSON-RPC (EVM) for MetaMask & tools

2. Transaction Entry Points

Transactions can enter NOORCHAIN via:

2.1 Cosmos SDK Clients

gRPC Msg clients

REST endpoints

CLI (e.g., noord tx …)

These produce Cosmos protobuf transactions.

2.2 EVM / Web3 Clients

MetaMask

Ethers.js / Web3.js

Hardhat, Foundry

Smart-contract UIs

These produce Ethereum RLP-encoded transactions.

2.3 PoSS dApps / frontends

Frontends generate:

MsgEmitSignal

MsgValidateSignal

submitted through gRPC/REST.

3. High-Level Flow: Submission → Finalization
Step-by-Step Lifecycle

Client builds transaction

Cosmos protobuf OR Ethereum RLP

Signed by user

Submit to node

via gRPC / REST / JSON-RPC

Mempool admission (CheckTx)

AnteHandler checks:

account, signature

nonce/sequence

gas & fees

basic validation

Block proposal

proposer selects txs from mempool

Consensus

validators vote

≥2/3 majority → block committed

DeliverTx execution

BaseApp routes tx to correct module

Runs Cosmos, EVM or PoSS logic

EndBlock & Commit

state changes finalized

new root hashed

validators updated

Query availability

gRPC / REST / JSON-RPC return new state

4. Cosmos Transaction Flow
4.1 Example: Bank Transfer (MsgSend)

User signs a Cosmos transaction containing MsgSend.

Broadcast → CheckTx:

protobuf decoding

sequence & signature verified

gas & fee check

DeliverTx:

routed to bank module

BankKeeper updates balances

events are emitted

EndBlock: staking & PoSS hooks run if applicable

4.2 Internals

AnteHandler uses AccountKeeper, BankKeeper

Module Router dispatches message to correct module

Cosmos execution model = strict protobuf msg routing

5. EVM Transaction Flow
5.1 Example: Smart-Contract Call

User sends signed Ethereum transaction via JSON-RPC (eth_sendRawTransaction).

CheckTx (EVM path):

decode RLP

verify ECDSA signature

check gas limit & EIP-1559 fields

check EVM balance (unur)

DeliverTx:

routed to evm module

EVMKeeper runs VM execution

storage writes, logs, bloom filters

gas is charged in unur

Receipts made available through JSON-RPC

5.2 Differences vs Cosmos Tx
Aspect	Cosmos Tx	EVM Tx
Encoding	protobuf	RLP
Signature	Cosmos secp256k1	Ethereum secp256k1
Nonce	Cosmos sequence	EVM nonce
Fees	Cosmos gas	EIP-1559
Execution	module handlers	EVM runtime
6. PoSS Transaction Flow

PoSS is implemented entirely within the Cosmos SDK as module x/noorsignal.

6.1 Main Messages

MsgEmitSignal

MsgValidateSignal

MsgAdminUpdateParams (admin-only)

6.2 Example Flow
Step 1 — Participant emits signal

builds MsgEmitSignal

CheckTx validates sender, gas, msg format

DeliverTx:

stores signal

sets metadata

emits EventSignalEmitted

Step 2 — Curator validates signal

builds MsgValidateSignal

CheckTx validates curator + limits

DeliverTx:

verifies eligibility

runs PoSS anti-abuse logic

computes reward:

weight

70/30 split

halving

daily limits

instructs BankKeeper to transfer rewards (when fully active)

emits:

EventSignalValidated

EventRewardDistributed

7. Block Lifecycle with PoSS
7.1 BeginBlock

PoSS may:

detect epoch boundary

reset or rotate daily counters

run cleanup tasks

7.2 DeliverTx

Cosmos tx → native modules

EVM tx → evm module

PoSS tx → noorsignal module

All paths coexist inside same block.

7.3 EndBlock

PoSS may:

finalize epoch stats

prepare next-day counters

Standard modules execute staking, governance updates, etc.

8. Event & Query Flow
8.1 Event Emission

Modules emit events during DeliverTx.

Sources:

Cosmos modules (bank, staking…)

EVM logs (converted to Ethereum logs + Cosmos events)

PoSS events:

EventSignalEmitted

EventSignalValidated

EventRewardDistributed

Explorers capture these via:

Websocket

gRPC stream

JSON-RPC

8.2 Queries

After commit, clients may query:

Cosmos gRPC / REST

balances

staking

PoSS stats

PoSS parameters

signal history

EVM JSON-RPC

transaction receipts

block logs

smart contract calls

account states

9. End-to-End Example Scenario

Block contains:

1 bank transaction

1 MsgEmitSignal

1 MsgValidateSignal

1 EVM transaction

During DeliverTx:

Bank tx updates balances

EmitSignal stores PoSS signal

ValidateSignal performs PoSS reward logic

anti-abuse

halving

70/30 split

bank transfers (when active)

EVM tx updates contract state

EndBlock finalizes all modules

The entire block produces a deterministic state and event set.

10. Role of CometBFT

CometBFT ensures:

deterministic ordering

voting-based finality

block commitment

CometBFT does not execute application logic.
Everything related to Cosmos, EVM or PoSS is executed inside BaseApp.

11. Summary (Header Block)

NOORCHAIN — Phase3_06 — Full Transactional Flow (Cosmos + EVM + PoSS)

This document defines how all transaction types flow through NOORCHAIN:

Cosmos protobuf transactions

EVM RLP transactions

PoSS module transactions

and explains the interactions between:

CometBFT

Cosmos SDK

Ethermint

PoSS runtime

It is the canonical reference for understanding the integrated execution pipeline of NOORCHAIN before Testnet activation.
