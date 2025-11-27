*NOORCHAIN — Phase 4B

PoSS Block Logic & Reward Mechanism**
Version 1.1 — Architecture Only (No Code)

🔧 1. Purpose of This Document

This document defines the complete functional behavior of the PoSS module during block execution, including:

BeginBlock algorithm

reward engine logic

halving cycle behavior

per-signal processing

daily limit enforcement

curator validation

reward distribution from PoSS Reserve

event emission

This specification is the canonical reference for implementing BeginBlocker and the reward engine in Phase 4C.

🧩 2. PoSS Execution Context

PoSS executes its entire logic in:

BeginBlock


Because:

validator power must be updated first (after staking)

PoSS must run before governance

deterministic PoSS state must be written before DeliverTx of next block

PoSS does not run in DeliverTx or EndBlock.

🚦 3. BeginBlock Algorithm Overview

The PoSS module performs 7 steps, always in this exact order:

1. Load signals collected during previous block
2. Validate signals (curator, anti-abuse, type rules)
3. Compute weighted signal units
4. Compute block reward using halving-adjusted formula
5. Distribute rewards (70/30) from PoSS Reserve
6. Update PoSS state (indexes, epochs, counters)
7. Clear temporary signal buffer


Each step is mandatory and strictly deterministic.

🪪 4. Step 1 — Load Signals from Previous Block

During DeliverTx, signals are registered but not processed.

BeginBlock loads:

signals = keeper.GetPendingSignals()


Each signal record contains:

sender

curator

type

weight

timestamp

block_height

If no signals → block reward = 0.

🛡️ 5. Step 2 — Validate All Signals

Validation must not be skipped.

Validation Rules (Phase 3 Final Specification)
A. Curator Validation

curator must belong to the Curator Registry

curator must have valid level: Bronze, Silver, Gold

curator must not exceed daily validation quota

curator must not validate their own signals

B. Anti-Abuse Validation

sender must not exceed daily signal quota

minimum delay between signals

subject to weight-specific limitations

sender and curator cannot be the same

no duplicate signal in same block

C. Type Validation

Signal types allowed:

Type	Name	Weight
0	micro-donation	0.5×
1	participation QR	1×
2	certified content	3×
3	CCN broadcast	5×
D. Economic Validation

PoSS Reserve must have enough NUR

weight ∈ {0.5, 1, 3, 5}

Invalid signals are discarded silently.

🎚️ 6. Step 3 — Compute Weighted Signal Units

Define:

unit = weight_of_signal × multiplier


Where multiplier = 1 (for now, defined in params)

Total unit count:
total_units = Σ (unit of each valid signal)


If total_units = 0 → reward = 0.

💰 7. Step 4 — Compute Block Reward (Halving-Adjusted)

The base reward per unit is defined in Phase 3.

Define:

base_reward_per_unit


Halving factor:

halving_factor = 1 / (2 ^ current_halving_cycle)


Final reward:

reward_total = total_units × base_reward_per_unit × halving_factor


Guarantees:

deterministic

predictable supply curve

follows 8-year halving cycle

🎁 8. Step 5 — Reward Distribution (70/30)

For each signal:

signal_reward = unit × base_reward_per_unit × halving_factor


Distribution:

participant_reward = signal_reward × 0.70
curator_reward     = signal_reward × 0.30


The keeper performs:

BankKeeper.SendCoinsFromModuleToAccount(PoSSReserve, participant)
BankKeeper.SendCoinsFromModuleToAccount(PoSSReserve, curator)


Rewards must use:

exact integer arithmetic

truncation rules defined in Phase 3

If PoSS Reserve lacks NUR → reward = 0.

🕒 9. Step 6 — Halving Process (Every 8 Years)

Halving cycle defined by:

blocks_per_halving = 8 years × 365 days × seconds per day / block_time


Stored in params.

Condition:

if block_height ≥ next_halving_height:
    halving_cycle += 1
    next_halving_height += blocks_per_halving


Effects:

reward becomes half

event emitted

new reward era begins

Halving must apply before computing rewards of the next block.

🧠 10. Step 7 — PoSS State Updates

After distributing rewards:

Update:

reward index

total signals count

last block height

epoch counters

anti-abuse stats

curator stats

participant stats

global PoSS metrics

These ensure consistency for:

explorers

dApps

governance

future modules

🧹 11. Step 8 — Clear Pending Signals

Once processed:

keeper.ClearPendingSignals()


Important for:

determinism

avoiding double counting

correct indexing

📡 12. Events Emitted

Events emitted for each block:

poss.reward_distributed

poss.signal_accepted

poss.signal_rejected

poss.halving_event

poss.stats_update

Events are essential for:

explorers

indexers

dApps

analytics

🎯 13. Summary of Block Logic
Step	Operation
1	load pending signals
2	validate signals
3	compute weighted units
4	compute halving-adjusted reward
5	send 70/30 rewards
6	update PoSS state
7	clear pending signals

This flow must be strictly deterministic and fully compatible with Cosmos consensus rules.