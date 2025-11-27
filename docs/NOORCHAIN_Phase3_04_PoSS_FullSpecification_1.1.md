NOORCHAIN 1.0 — Phase 3.04
Proof of Signal Social (PoSS) — Full System Specification
Version 1.1 — Official Document

Purpose of this document
This document fully specifies the Proof of Signal Social (PoSS) consensus-economic mechanism of NOORCHAIN.
It defines all logic, all parameters, all components, all flows, without code, to serve as:

the reference for implementing the PoSS module (x/noorsignal)

the basis for Testnet genesis parameters

the source of truth for documentation and audits

the validation blueprint for reward calculations.

1. Overview of PoSS

PoSS (Proof of Signal Social) is NOORCHAIN’s native social-consensus mining mechanism, based on:

human-validated signals

curator oversight

transparent on-chain scoring

fixed supply (299,792,458 NUR)

8-year halving

zero inflation (pre-allocated reward pool)

70% rewards to participants, 30% to curators

PoSS replaces energy-based mining and solves the “real-world contribution” problem.

It aligns with NOORCHAIN’s vision:
ethical, human-centric, Swiss-compliant, socially anchored, and economically sustainable.

2. Core Concepts
2.1. Signal

A signal is an on-chain representation of a positive, validated social action.

2.2. Participant

The user who emits the signal.

2.3. Curator

A verified human or organization who validates the signal.
Curators have three tiers:

Bronze

Silver

Gold

Higher tiers have:

higher visibility

higher validation authority

stricter requirements

2.4. PoSS Reward Pool

A fixed pre-allocated pool of NUR used to reward signals.
Defined in Genesis (Phase3_05):

80% of total supply = PoSS Mintable Reserve

2.5. Reward Split

Every validated signal distributes:

70% → participant

30% → curator
(as permanently agreed and stored in memory)

2.6. Epoch

The time window used to:

aggregate stats

check daily limits

apply halving progression

process PoSS module maintenance

Default: 1 day.

3. Types of Signals

NOORCHAIN supports four official signal categories, each with a different weight.

3.1. Micro-donation Signal

A small donation or micro-action validated by a curator.
Examples: 1 CHF donation to an NGO, 1 NUR micro-tip, etc.

3.2. Verified Participation

Confirmed presence or contribution to a real-world activity.
Examples:

volunteering

community event

social participation

3.3. Certified Content

Publication of positive or original content validated by curators.
Examples:

educational posts

constructive online content

CCN (Certified Content Noorchain)

3.4. CCN Signal (Content Collaboration Noorchain)

High-value social or educational content certified by a curator.

4. Signal Weight System

Each signal has a weight determining its reward impact.

Weight Scale (public):
Signal Type	Weight
Micro-donation	x1
Participation	x2
Certified Content	x3
CCN Signal	x5

Higher weight = larger share of the epoch reward bucket.

5. Anti-Abuse & Limits

PoSS includes strong anti-abuse rules:

5.1. Daily Limits

Each participant can emit N signals/day.
Recommended default:

DailyMaxSignals = 10

5.2. Rate Limits

Curators cannot validate more than:

CuratorMaxValidations = 50/day

5.3. Duplicate Prevention

No repeated validation of identical signals.
Hashes stored in module state.

5.4. Tier-based Curator Controls

Higher tiers = higher validation capacity, not higher rewards.

5.5. Economic Protections

The system prevents:

wash signaling

circular curator-participant loops

self-curation

mass spam by NGOs or bots

Sybil identity attacks

6. Reward Mechanism

PoSS uses deterministic, fully public rules.

6.1. Reward Formula

Let:

W = weight of the signal

R_day = daily reward budget for the network

R_signal = (W / totalWeightsOfDay) * R_day

ParticipantReward = R_signal * 0.70

CuratorReward = R_signal * 0.30

Rewards come from the PoSS Reserve.

6.2. Daily Reward Budget

Daily rewards follow halving cycles (Section 8).
Example (illustrative):

R_day = BaseReward / (2 ^ HalvingEpoch)

6.3. Distribution Method

Rewards are distributed instantly:

Participant → direct transfer

Curator → direct transfer

Sender: PoSS Reserve Address
Defined in Phase3_05.

7. PoSS Module Parameters

All parameters are stored in ParamsKeeper.

7.1. Global Parameters
Parameter	Description
DailyMaxSignals	max signals/day per user
CuratorMaxValidations	max validations/day per curator
WeightMicroDonation	1
WeightParticipation	2
WeightCertifiedContent	3
WeightCCN	5
ParticipantRatio	0.70
CuratorRatio	0.30
EpochDuration	24h
HalvingBlocks	derived from 8-year cycle
ReservePoolAddress	core reward wallet
StimulusPoolAddress	incentive pool for early adoption
8. Halving Mechanism

Halving occurs every 8 years.

Formula:

Reward = InitialReward / 2^(Years / 8)


Halving applies to:

daily reward budget

long-term emission curve from the PoSS Reserve

NOORCHAIN uses no inflation: all emission is from pre-minted reserve.

9. Event System

PoSS emits the following events for transparency:

9.1. EventSignalEmitted

sender

signal type

weight

timestamp

9.2. EventSignalValidated

curator

participant

signal hash

timestamp

9.3. EventRewardDistributed

participant reward

curator reward

total weight

reserve address

new reserve balance

9.4. EventAdminParamsUpdated

old parameters

new parameters

admin address

All events are queryable by explorers and dashboards.

10. Queries (gRPC + REST)

PoSS exposes:

10.1. QuerySignals(address)

List all signals for a participant.

10.2. QueryCurator(address)

Get curator tier & stats.

10.3. QueryParams()

Current PoSS parameters.

10.4. QueryStats()

Network-wide PoSS data (daily weights, totals, halving epoch…).

10.5. QueryRewardHistory(address)

Rewards earned by any participant or curator.

11. PoSS Flow (Full Transactional Lifecycle)
Step 1 — Participant emits signal

MsgEmitSignal
Stored with metadata + weight.

Step 2 — Curator validates

MsgValidateSignal
Linked to the participant.

Step 3 — Anti-abuse checks

Daily limits, duplicate detection, hash validation.

Step 4 — Reward computation

Based on weight, ratios, halving epoch.

Step 5 — Reward distribution

70% participant
30% curator

Step 6 — Events emitted

Used for explorers & analytics.

12. Admin Logic

Only designated admin addresses can:

update PoSS parameters

update curator tiers

modify daily limits

manage Stimulus Pool (early adoption incentives)

Admin addresses are defined in Genesis.

13. Module State (Conceptual)

PoSS stores:

signals

signal hashes

curator status

daily counters

reward history

parameters

reserve pool balance (read-only reference to bank module)

14. Security Model

PoSS includes:

strict per-address rate limiting

no self-validation

tier-based gating

hash-based signal uniqueness

reserve-controlled reward issuance

deterministic calculation for auditability

no oracle or external data dependency

15. Summary for header

NOORCHAIN — PoSS Full Specification (Phase3_04, v1.1)
Defines the entire Proof of Signal Social mechanism:

Signal types & weights

Curator system

Reward 70/30 split

Halving every 8 years

Anti-abuse rules

Module parameters

Queries & events

Transaction lifecycle

This document is the canonical reference for implementing NOORCHAIN’s x/noorsignal module and for building Testnet 1.0.