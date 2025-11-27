NOORCHAIN 1.0 — Phase 3.06
Full Transactional Flow (Cosmos + EVM + PoSS)
Version 1.1 — Official Document

Purpose of this document
Describe, end-to-end, how a transaction flows through NOORCHAIN:

standard Cosmos transactions

EVM (Ethermint) transactions

PoSS transactions
and how they interact with:

CometBFT

Cosmos SDK (BaseApp, modules, keepers)

EVM module

PoSS module (x/noorsignal)

This document is the “big picture” view of NOORCHAIN’s runtime behavior.

1. Layers Overview

NOORCHAIN is structured in four main runtime layers:

Network & Consensus Layer

P2P networking

CometBFT consensus

block proposal, voting, finalization

Application Layer (Cosmos SDK)

BaseApp

Module Manager

Keepers & Stores

Execution Sub-Layers

Cosmos Msg pipeline (native modules)

EVM pipeline (Ethermint)

PoSS pipeline (noorsignal module)

API / Client Layer

gRPC

REST (gRPC-Gateway)

JSON-RPC (EVM)

This document focuses on what happens from the moment a transaction is created, until it is finalized and queryable.

2. Transaction Entry Points

Transactions can enter NOORCHAIN through different gateways:

2.1. Cosmos SDK Clients

cosmos-sdk gRPC Msg clients

REST JSON endpoints

CLI (later: noord tx ...)

These produce Cosmos transactions (protobuf-encoded).

2.2. EVM / Web3 Clients

MetaMask

Ethers.js / Web3.js

Hardhat / Foundry

Smart contract UIs

These produce Ethereum-style RLP-encoded transactions.

2.3. PoSS UI / dApps

Custom frontends or apps that:

build Cosmos messages (MsgEmitSignal, MsgValidateSignal)

send them via gRPC / REST / CLI.

In all cases, the transaction ends up submitted to the node, which then passes it into CometBFT’s mempool.

3. High-Level Flow: From Submission to Finalization
3.1. Step-by-Step Overview

Client builds transaction

Cosmos or EVM format

signed by the user

Transaction submission

via gRPC, REST or JSON-RPC

node receives the raw transaction

Mempool admission (CheckTx)

AnteHandler runs basic checks

invalid txs are rejected

valid txs enter mempool

Block proposal

proposer node selects txs from mempool

creates a block proposal

Consensus

validators vote on the proposed block

once ≥2/3 voting power agrees → block is committed

Block execution (DeliverTx)

BaseApp processes each tx

state transitions apply via modules (Cosmos / EVM / PoSS)

EndBlock & Commit

validators updated

state root computed

committed to CometBFT

Query availability

new state is queryable via gRPC / REST / JSON-RPC

4. Cosmos Transaction Flow
4.1. Example: Simple Bank Transfer (Cosmos Tx)

User signs a MsgSend (Cosmos bank module).

Tx is broadcast to a node.

CheckTx (Cosmos path):

decode tx (protobuf)

account sequence & signature checked

gas & fees checked (in unur)

If valid → stored in mempool.

On block inclusion:

DeliverTx routes to the bank module

BankKeeper updates balances

events are emitted (transfer event)

BeginBlock / EndBlock logic executes (staking, PoSS daily hooks if any).

New balances are committed and visible in queries.

4.2. Cosmos Tx Internals

AnteHandler uses:

AccountKeeper

BankKeeper

ParamsKeeper

Router directs to:

bank, staking, gov, noorsignal, etc.

5. EVM Transaction Flow
5.1. Example: Smart Contract Call

User sends a signed Ethereum transaction (e.g., eth_sendRawTransaction).

Node receives RLP transaction.

CheckTx (EVM path):

EVM AnteHandler decodes tx

verifies ECDSA signature

checks nonce

checks gas limit and max fees

checks EVM account balance in unur

If valid → tx is added to the EVM-aware mempool.

Block proposer picks txs.

During DeliverTx:

BaseApp routes tx to evm module

EVMKeeper executes contract call

EVM state is updated (storage, balances, nonces)

logs and bloom filters are generated

Gas consumed is charged in unur.

Receipts become accessible via JSON-RPC (eth_getTransactionReceipt).

5.2. Differences vs Cosmos Tx

EVM txs are opaque to Cosmos modules except evm & feemarket.

They use EVM account nonce, not Cosmos sequence.

They pay gas using EIP-1559 rules via feemarket.

6. PoSS Transaction Flow

PoSS is implemented as a Cosmos SDK module (noorsignal).

6.1. Main PoSS Messages

MsgEmitSignal

participant emits a signal

includes type, metadata, optional link

MsgValidateSignal

curator validates a specific signal

triggers reward logic

MsgAdminUpdateParams

admin-only

updates parameters (limits, weights, etc.)

6.2. Example: Emitting and Validating a PoSS Signal

Step 1 — Emit Signal

Participant builds a Cosmos tx containing MsgEmitSignal.

Tx is broadcast via gRPC/REST.

CheckTx (Cosmos pipeline) checks:

sender account

fees & gas

format of MsgEmitSignal

In DeliverTx:

noorsignal module stores the signal

calculates its weight

marks it as “awaiting validation”

emits EventSignalEmitted.

Step 2 — Validate Signal

Curator builds a tx with MsgValidateSignal.

CheckTx ensures:

curator is a registered curator

curator hasn’t exceeded daily limits

signal exists and is not already validated

In DeliverTx:

PoSS executes full anti-abuse logic

computes rewards based on:

weight

daily epoch totals

70/30 split

instructs BankKeeper to send rewards:

from ReservePoolAddress

to participant & curator

Emits:

EventSignalValidated

EventRewardDistributed.

7. Block Lifecycle with PoSS
7.1. BeginBlock

PoSS may perform:

epoch boundary checks (new day?)

reset daily counters when epoch changes

light housekeeping (e.g., prune old pending signals)

7.2. DeliverTx

Cosmos txs → run in their module

EVM txs → run in evm module

PoSS txs → run in noorsignal

Multiple tx types can coexist in the same block.

7.3. EndBlock

PoSS may:

finalize stats for the epoch

emit aggregated events (optional)

prepare next epoch metadata

Staking and validator updates also occur in EndBlock through the staking module.

8. Event & Query Flow
8.1. Events

Every transaction can emit multiple events:

Cosmos modules → bank, staking, gov events

EVM module → logs (converted to Ethereum logs, plus Cosmos events)

PoSS → signals, validation, rewards

Explorers and off-chain tools subscribe to:

WebSocket RPC

gRPC streams

JSON-RPC logs

8.2. Queries

After commit, clients can query:

Cosmos gRPC / REST

balances

stakes

PoSS stats

JSON-RPC

EVM receipts

EVM logs

block and tx data

PoSS-specific queries include:

QuerySignals(address)

QueryCurator(address)

QueryParams()

QueryRewardHistory(address)

QueryStats()

9. Example End-to-End Scenario

Scenario:

A user sends NUR to another user (Cosmos bank tx).

The same user emits a PoSS signal for a good action.

A curator validates that signal.

A dApp updates a smart contract via EVM.

Block behavior:

Block contains:

1 bank tx

1 MsgEmitSignal

1 MsgValidateSignal

1 EVM tx

During DeliverTx:

Bank tx updates balances.

EmitSignal stores PoSS signal.

ValidateSignal:

checks anti-abuse

sends rewards from PoSS Reserve to both parties.

EVM tx updates contract storage.

At the end of the block:

staking & other hooks run

PoSS updates daily stats if needed

state is committed and exposed for queries.

10. Role of CometBFT in the Flow

CometBFT:

orders all transactions

ensures Byzantine-fault-tolerant consensus

provides finality for each block

It does not know about:

PoSS logic

EVM semantics

bank, staking, etc.

All of that is handled by the Cosmos application (NOORCHAIN).

11. Summary (Header-Style)

NOORCHAIN — Phase3_06 — Full Transactional Flow (Cosmos + EVM + PoSS)
This document describes how:

Cosmos transactions

EVM transactions

PoSS transactions
move through NOORCHAIN from client to final state, including:

mempool & CheckTx

block proposal & consensus

DeliverTx routing to modules

PoSS reward distribution

event emission & queries.

It is the reference for understanding NOORCHAIN’s runtime behavior and for validating that all layers (CometBFT, Cosmos SDK, Ethermint, PoSS) integrate correctly.