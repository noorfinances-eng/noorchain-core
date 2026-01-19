NOORCHAIN 1.0 — PoSS Testnet Practical Guide (v1.1)
Practical scenarios for testing PoSS on a development testnet

Last Updated: 2025-12-03

0. Scope of this Document

This guide explains how to practically test the PoSS module (x/noorsignal) on a NOORCHAIN development testnet, based on the current implementation status:

PoSSEnabled = false (default)

PoSS signals can be created

Daily counters are active

Rewards are computed theoretically but always 0/0

No real minting or balance transfers happen yet

PoSS is in counting / simulation mode

This document complements:

NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md

NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md

NOORCHAIN_Phase3_06_MsgCreateSignal_Spec_1.1.md

1. Testnet Setup (Minimal)

This section describes how a developer can instantiate a local testnet.
Important: Phase 3 documentation must not run code in production; the following commands are illustrative only.

1.1. Initialize a local testnet

Example:

noord init noorchain-testnet --chain-id noorchain-1


This creates:

a default home directory

a placeholder genesis file

a basic Tendermint/CometBFT config

No PoSS-specific changes occur yet.

1.2. Create test accounts

Two accounts are required:

participant (sender of PoSS signals)

curator (validator of PoSS signals, receives 30% later)

Example:

noord keys add participant
noord keys add curator


These will produce valid noor1... bech32 addresses compatible with PoSS.

1.3. Add genesis balances

Both accounts need minimal test balances to pay gas.

Example:

noord add-genesis-account participant 1000000unur
noord add-genesis-account curator     1000000unur


This step does not interact with PoSS; it only prepares accounts for transactions.

1.4. Generate and collect gentx

Standard validator creation:

noord gentx participant 1000000unur
noord collect-gentxs


Again, purely Cosmos-level — PoSS is not involved yet.

1.5. Start the testnet
noord start


The chain will start with:

PoSS module loaded

Default GenesisState for PoSS

Daily counters reset

Parameters loaded via DefaultParams()

No minting or reward transfers occur.

2. Creating PoSS Signals on Testnet

Even with PoSSEnabled = false, you can already:

submit PoSS transactions

observe daily counters

inspect theoretical reward computation

check events in logs

This allows complete dry-run testing of the module.

2.1. Build a MsgCreateSignal transaction

A valid transaction includes:

participant address

curator address

signal type (enum)

optional reference

standard fee and gas

Example structure (conceptual):

{
  "type": "noorsignal/MsgCreateSignal",
  "value": {
    "participant": "noor1xxx...",
    "curator":     "noor1yyy...",
    "signal_type": 2,
    "reference":   "event-qr-123"
  }
}


Signal types:

Enum	Meaning
1	MICRO_DONATION
2	PARTICIPATION
3	CONTENT
4	CCN
2.2. Broadcast the transaction

Through CLI, gRPC or any SDK-based tool.

Example (conceptual):

noord tx noorsignal create-signal participant curator 2 "ref" --from participant


The message is accepted if:

addresses are valid

enum is valid

reference ≤ 256 chars

It is NOT rejected for daily limits or PoSS disabled.

2.3. Expected behaviour with PoSSEnabled = false

The following occurs:

participant daily counter increments

curator daily counter increments

TotalSignals increments

theoretical reward is computed internally

final reward = 0/0 (participant/curator)

event is emitted

This allows detailed debugging without any monetary effect.

3. Inspecting PoSS Internal Behaviour
3.1. Checking daily counters

Using queries (once wired):

QueryDailySignals(participant)
QueryDailySignals(curator)


Expected:

increasing counts during the same day

new counters on the next day

No automatic reward change.

3.2. Checking raw theoretical reward

When calling reward helpers internally (unit tests):

ComputeSignalReward(params, signalType, height)


base reward applied

weight applied

halving applied

then split 70/30

but final = 0 because PoSS disabled

This shows that reward math is already fully operational.

3.3. Checking events

Each signal emits an event such as:

noorsignal_signal_created
noorsignal_reward_distributed  (reward=0/0)


These appear in node logs or RPC event streams.

4. Recommended Test Scenarios

The following scenarios validate PoSS stability before any real minting exists.

Scenario A — PoSS disabled routine

Goal:
Verify default behaviour.

Steps:

Send 1–5 signals.

Observe counters increment.

Confirm rewards are always 0.

Confirm events appear normally.

Success criteria:

No minting

No errors

All counters behave as expected

Scenario B — Daily limit threshold

Set MaxSignalsPerDay = 3 (temporary config for testing).

Send 5 signals.

Expected:

first 3 → reward computed but set to 0 (PoSS disabled)

last 2 → also reward = 0

counters = 5

The difference will matter only once PoSS is enabled.

Scenario C — Curator validation limit

Set MaxSignalsPerCuratorPerDay = 2.

Validate 4 signals with same curator.

Expected:

curator counter increments to 4

but last 2 give curator reward = 0

Scenario D — Multi-day behaviour

Create signals on day D.

Advance block time or wait.

Create signals on day D+1.

Expected:

counters(D) unchanged

counters(D+1) start fresh

Scenario E — Halving simulation (unit tests only)

Override:

HalvingPeriodBlocks = very small number
PoSSEnabled = true


Call reward helper at different block heights.

Expected:

block < halving → full base reward

block > halving → reward / 2

never negative

This ensures long-term stability.

5. What This Testnet Guide Does NOT Do

It does not activate PoSS

It does not mint real NUR

It does not configure the real PoSS Reserve

It does not describe Mainnet setup

It does not modify economic or legal parameters

This is strictly:

technical testing

behaviour verification

pre-activation simulation

6. Next Documents Required (Phase 3 Final)

After this Practical Guide, the next Phase 3 documents to write are:

Phase3_08_PoSS_ParamSubspace_Design_1.1.md

Phase3_09_Testnet_CLI_and_Query_Design_1.1.md

Phase3_10_GenesisPack_Overview_1.1.md

These prepare the transition into Phase 4 (code wiring) and Phase 6 (Genesis Pack).

Summary

This guide provides a complete and clean method to:

start a NOORCHAIN testnet

submit PoSS transactions

verify PoSS behaviour

inspect counters

inspect theoretical rewards

ensure complete safety before activation

PoSS enters no-risk test mode here, perfectly aligned with Legal Light CH.
