NOORCHAIN — Phase 4C
PoSS Msg — Testnet Simulation Guide (v1.1)

(English — official documentation)

1. Overview

This document explains how to simulate PoSS messages (MsgCreateSignal) on a testnet chain running the current Phase 4 Noorchain Core.

It reflects the real state of the code:

MsgCreateSignal exists (pure Go, no protobuf)

Legacy handler is active (handler.go)

No MsgServer yet

No real minting (PoSS economics OFF)

Internal logic computes rewards, daily counters, pending-mint, and updates genesis totals

This is the first fully functional PoSS message pipeline on Noorchain.

2. Current Capabilities (Phase 4C)
✔ Supported

Creating a PoSS signal via MsgCreateSignal

Per-signal reward computation

Halving-ready logic

70/30 reward splitting (participant/curator)

Daily counters per participant

Daily reward caps (MaxRewardPerDay)

PendingMint queue (internal planning)

Global counters (TotalSignals, TotalMinted)

ABCI events emitted for indexers

Fully deterministic behaviour

✘ Not yet supported

No real minting on BankKeeper

No MsgServer (protobuf)

No CLI tx command

No verification of curator daily limits

No multi-signal batch

No hooks

No on-chain governance for params (only subspace defaults)

These will come in future Phases (Phase 5–6).

3. MsgCreateSignal — Structure Recap
type MsgCreateSignal struct {
    Participant string
    Curator     string
    SignalType  SignalType
    Metadata    string
    Timestamp   time.Time
    Date        string // "YYYY-MM-DD"
}


Used through the helper:

NewMsgCreateSignal(...)

4. Testnet Usage — How to Build the Tx

Because MsgServer does not exist yet, you create a tx using the legacy router path.

Message route
route = "noorsignal"
type  = "CreateSignal"

Example JSON message (broadcast in testnet shell)
{
  "@type": "/noorsignal.CreateSignal",
  "participant": "noor1abcde...",
  "curator": "noor1vwxyz...",
  "signal_type": "micro_donation",
  "metadata": "ipfs://example",
  "timestamp": "2025-01-01T12:00:00Z",
  "date": "2025-01-01"
}

5. Simulation Scenarios

Below are the expected behaviours when broadcasting MsgCreateSignal to a local or remote testnet.

Scenario 1 — Basic Signal (no daily limits reached)
Input

participant = A

curator = C

type = micro_donation

date = 2025-01-01

BaseReward = 1 unur (example)

Weight(micro_donation) = 2

Expected pipeline

Compute reward

raw = base × weight

1 × 2 = 2 unur

70% = 1.4 → floor = 1 unur

30% = 0.6 → floor = 0 unur (current decimals = 0)

DailyReward(A, 2025-01-01) = increment by 1

DailySignals(A, 2025-01-01) = +1

Record PendingMint entry

Genesis totals:

TotalSignals +1

TotalMinted +1

Expected Events (ABCI)
type: noorsignal_create_signal
  participant: noor1abc...
  curator: noor1xyz...
  signal_type: micro_donation
  participant_reward: 1unur
  curator_reward: 0unur

Scenario 2 — MaxRewardPerDay cap reached
Params

MaxRewardPerDay = 3 unur

Steps

Participant sends 4 valid signals same day.

Expected

Signals 1,2,3 → rewarded

Signal 4 → 0 / 0 reward

DailySignals increments to 4

PendingMint still records entries

TotalSignals increments by 4

TotalMinted increments by the reward of the first three only

Scenario 3 — PoSS disabled (PoSSEnabled = false)
Expected

Rewards always zero

Daily limits still enforced (but irrelevant since reward = 0)

PendingMint entries created

TotalSignals increments

TotalMinted does NOT increase

The chain behaves as a pure accounting system.

6. Testnet Validation Checklist

Before approving Phase 4C, the following must work:

✔ Tx accepted by CheckTx
✔ handler.go triggered
✔ ValidateBasic() runs correctly
✔ Internal pipeline executed
✔ Daily counters updated
✔ TotalSignals increases
✔ TotalMinted increases ONLY if reward > 0
✔ Events emitted
✔ No panic
✔ No real mint on BankKeeper
✔ PendingMint entries stored (readable via debug)

If all items are green → Phase 4C is validated.

7. How to Verify State (Debug Instructions)

Inside a testnet environment, attach a debugger or print state using custom queries:

Show global PoSS stats:
GetGlobalStats()


Expected fields:

{
  "total_signals": "...",
  "total_minted": "...",
  "possenabled": false,
  "max_signals_per_day": ...,
  "max_signals_per_curator_per_day": ...,
  "posss_reserve_denom": "unur"
}

Check daily counters:
GetDailySignalsCount(address, date)

Check pending mint:

Scan KV store keys:

pending_mint:<height>:<address>:<timestamp>

8. Final Notes

This version is perfectly aligned with Legal Light CH: no real minting.

The pipeline is ready for Phase 5 — enabling PoSS economics.

No module or file conflicts remain.

Code compiles: go test ./... passes.

Noorchain Core is stable at Phase 4C level.

This document is the official Testnet Simulation Guide for MsgCreateSignal.