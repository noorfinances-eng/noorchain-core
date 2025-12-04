NOORCHAIN — Phase 4B

PoSS Block Logic & Reward Mechanism
Version 1.1 — Architecture Only (No Code)

1. Purpose of This Document

This document defines the complete functional behavior of the PoSS module during block execution. It specifies:

the BeginBlock algorithm

reward engine logic

halving cycle behavior

per-signal processing

daily limit enforcement

curator validation

reward distribution from the PoSS Reserve

event emission

This file is the canonical specification to implement the BeginBlocker and the PoSS reward engine in Phase 4C.

2. PoSS Execution Context

PoSS logic runs entirely in BeginBlock, because:

validator power must already be updated (after staking)

PoSS must run before governance logic

deterministic PoSS state must be written before the next block processes DeliverTx

PoSS does not run in DeliverTx or EndBlock.

3. BeginBlock Algorithm Overview

PoSS always performs the following steps in this exact order:

Load signals collected during the previous block

Validate signals (curator, anti-abuse, type validation)

Compute weighted signal units

Compute block reward using the halving-adjusted formula

Distribute rewards (70/30) from the PoSS Reserve

Update PoSS state (indexes, epochs, counters)

Clear the temporary signal buffer

All steps are mandatory and strictly deterministic.

4. Step 1 — Load Signals from Previous Block

During DeliverTx, signals are registered but not processed.

BeginBlock retrieves them:

sender

curator

signal type

weight

timestamp

block height

If no signals exist, the block reward is zero.

5. Step 2 — Validate All Signals

Validation cannot be skipped.

A. Curator Validation

curator must be registered

curator must have level Bronze, Silver, or Gold

curator must not exceed their daily validation quota

curator must not validate their own signals

B. Anti-Abuse Validation

sender must not exceed their daily signal quota

minimum delay between signals

weight-specific limitations

sender and curator must be different

no duplicate signal from the same sender within the same block

C. Type Validation

Allowed signal types and weights:

micro-donation → weight 0.5

participation QR → weight 1

certified content → weight 3

CCN broadcast → weight 5

D. Economic Validation

the PoSS Reserve must have enough NUR

weight must belong to the allowed set {0.5, 1, 3, 5}

Invalid signals are silently discarded.

6. Step 3 — Compute Weighted Signal Units

Each valid signal contributes:

unit = weight × multiplier

The multiplier is currently 1 (defined in parameters).

The total unit count for the block is:

total_units = sum of all units from valid signals

If total_units is zero, the block reward is zero.

7. Step 4 — Compute Block Reward (Halving-Adjusted)

The block reward is calculated using:

the base reward per unit (Phase 3 definition)

the current halving factor

Halving factor:

halving_factor = 1 / (2 ^ current_halving_cycle)

Final block reward:

reward_total = total_units × base_reward_per_unit × halving_factor

This ensures:

deterministic rewards

a predictable supply curve

a fixed 8-year halving cycle

8. Step 5 — Reward Distribution (70/30)

For each signal:

signal_reward = unit × base_reward_per_unit × halving_factor

Distribution:

70% to the participant

30% to the curator

The keeper executes transfers from the PoSS Reserve:

send reward to the participant

send reward to the curator

Rules:

integer arithmetic only

truncation rules defined in Phase 3

if the PoSS Reserve lacks funds → reward becomes zero

9. Step 6 — Halving Process (Every 8 Years)

Halving uses:

blocks_per_halving = 8 years (converted into block count)

If the current block height reaches the next halving point:

halving_cycle increases by 1

next_halving_height is shifted by blocks_per_halving

Effects:

reward is reduced by half

a halving event is emitted

a new reward era begins

Halving always applies before rewards of the next block are computed.

10. Step 7 — PoSS State Updates

After reward distribution, the module updates:

reward index

total signal count

last processed block height

epoch counters

anti-abuse statistics

curator statistics

participant statistics

global PoSS metrics

This ensures consistency for explorers, dApps, and governance.

11. Step 8 — Clear Pending Signals

After successful processing, the pending-signal buffer is cleared.

This prevents:

double counting

inconsistent indexing

non-deterministic behavior

12. Events Emitted

PoSS emits events during BeginBlock processing:

poss.reward_distributed

poss.signal_accepted

poss.signal_rejected

poss.halving_event

poss.stats_update

These events are required for indexers, explorers, analytics, and external systems.

13. Summary of Block Logic

Step | Operation
1 | load pending signals
2 | validate signals
3 | compute weighted units
4 | compute halving-adjusted reward
5 | distribute 70/30 rewards
6 | update PoSS state
7 | clear pending signals

This flow must be strictly deterministic and fully aligned with Cosmos consensus requirements.
