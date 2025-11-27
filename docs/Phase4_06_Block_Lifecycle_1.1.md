**NOORCHAIN — Phase 4A

Block Lifecycle Specification (Cosmos SDK + Ethermint + PoSS)**
Version 1.1 — Architecture Only, No Code

🔧 1. Purpose of This Document

This document defines the full, detailed block lifecycle of the NOORCHAIN application, including:

BeginBlock logic

DeliverTx logic

EndBlock logic

Commit cycle

EVM integration points

PoSS integration points

Ordering constraints

This file ensures deterministic block execution during Phase 4C coding.

🧩 2. Block Lifecycle Overview

Each block follows this order:

BeginBlock → DeliverTx → EndBlock → Commit


Inside this flow, Cosmos modules, Ethermint modules and PoSS must run at very precise points.

🏁 3. BeginBlock Lifecycle (Critical Order)
BeginBlock Execution Order
1. FeeMarket
2. EVM
3. Staking
4. PoSS (noorsignal)
5. Governance

3.1 FeeMarket BeginBlock

updates the dynamic base fee (EIP-1559 style)

adjusts global gas prices

3.2 EVM BeginBlock

prepares EVM block context

resets logs

sets block time, base fee, coinbase

3.3 Staking BeginBlock

processes validator updates

recomputes voting power

handles slashing/unjailing events

3.4 PoSS BeginBlock

Uses up-to-date validator power to:

read aggregated signals (later Phase 4B)

compute PoSS rewards

apply 70/30 distribution model

check for halving epochs

record PoSS block index

3.5 Governance BeginBlock

checks proposal status

handles voting periods

🚀 4. DeliverTx Lifecycle

For each transaction:

4.1 AnteHandler

signature checks

account sequence and number

account balance for gas

EVM-specific validation (if eth tx)

4.2 Message Routing

Routes msg types:

cosmos SDK messages

MsgEthereumTx (EVM)

PoSS custom messages

staking & governance msgs

4.3 State Transitions

KVStore writes

event emission

gas accounting

4.4 Special case: Ethereum transactions

executed in the EVM

logs stored

receipts generated

failures revert state via EVM state DB

🧱 5. EndBlock Lifecycle
5.1 Staking EndBlock

finalize validator updates

produce new validator set

prepare diffs for CometBFT consensus

5.2 Governance EndBlock

finalize proposals whose voting ended

apply proposal results

Ordering:

staking → gov

🔄 6. Commit Phase
6.1 State Commitment

Multistore commits via IAVL

produces new app hash

stored for next block

6.2 EVM State Commit

flush EVM state DB

write bloom filter

index logs

6.3 PoSS Commit

(later Phase 4B)

commit epoch counters

commit reward update indexes

🌐 7. Lifecycle Diagram
                ▼
        ┌──────────────────┐
        │   BeginBlock     │
        └──────────────────┘
   feemarket → evm → staking → poss → gov
                ▼
        ┌──────────────────┐
        │    DeliverTx     │
        └──────────────────┘
        ante → msg → state writes
                ▼
        ┌──────────────────┐
        │    EndBlock      │
        └──────────────────┘
         staking → gov
                ▼
        ┌──────────────────┐
        │     Commit       │
        └──────────────────┘
 state commit → evm commit → poss commit

🧠 8. Determinism Constraints

The following rules are mandatory:

fee market must run before EVM

staking must run before PoSS

PoSS must run before governance

EVM must commit before PoSS can finalize

EndBlock must only finalize staking & governance

PoSS must NEVER run during EndBlock (only BeginBlock)

These rules prevent:

reward inconsistencies

EVM mismatch

invalid validator power

governance misalignment

🎯 9. Summary

NOORCHAIN’s block lifecycle is:

Phase	Order
BeginBlock	feemarket → evm → staking → poss → gov
DeliverTx	ante → msg → state
EndBlock	staking → gov
Commit	state → evm → poss

This document is the authoritative reference for block execution during coding.