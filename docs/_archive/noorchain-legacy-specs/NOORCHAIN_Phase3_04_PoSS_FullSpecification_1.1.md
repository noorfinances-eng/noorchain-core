NOORCHAIN 1.0
Phase 3.04 — PoSS Full Specification
Version 1.1
Last Updated: 2025-12-03
0. Purpose of this Document

This document defines the complete, authoritative specification of
NOORCHAIN’s Proof of Signal Social (PoSS) mechanism.

It is the official reference for:

implementing the PoSS module (x/noorsignal)

Testnet and Mainnet genesis parameter definition

internal and external audits

governance and compliance reviews

whitepapers and Phase 6 Genesis Pack documentation

This is a conceptual + functional specification, not code.

1. PoSS Overview

PoSS (Proof of Signal Social) is NOORCHAIN’s human-centered, non-financial reward mechanism
that recognises verified positive social actions.

PoSS is based on:

human-validated signals

curator verification (Bronze / Silver / Gold)

transparent on-chain scoring

fixed supply: 299,792,458 NUR

halving every 8 years

zero inflation (pre-allocated PoSS Reserve)

immutable 70% / 30% reward split

Swiss Legal Light CH compliance

PoSS is not mining.
It is a distribution model for tokens already allocated at genesis.

2. Core Concepts
2.1 Signal

The on-chain representation of a positive, verifiable social action.
It is emitted by a participant and later validated by a curator.

2.2 Participant

Any user who emits a PoSS signal.

2.3 Curator

A verified human or organization responsible for signal validation.
Curators have three tiers:

Bronze

Silver

Gold

Tiers influence trust and permissions, not reward amount.

2.4 PoSS Reserve (80% of total supply)

A fixed, immutable reserve dedicated exclusively to PoSS rewards.

2.5 Reward Split (Immutable)

70% → participant

30% → curator

This rule cannot be changed by governance.

2.6 Epoch (24h)

The daily window used for counters, limits, halving progression and budgets.

3. Signal Types (4 Categories)
Type	Description	Examples
Micro-donation	Small symbolic or financial act	micro-tip, NGO support
Verified Participation	Presence or contribution	volunteering, community events
Certified Content	Curator-approved content	educational or social impact
CCN (Content Collaboration Noorchain)	High-value long-form content	CCN Studio
4. Weight System
Signal Type	Weight
Micro-donation	1
Participation	2
Certified Content	3
CCN	5

Weights determine the relative share of the daily reward budget.

5. Anti-Abuse Measures

PoSS incorporates strict safeguards to ensure fairness and prevent manipulation.

5.1 Daily Limits

DailyMaxSignals per participant (default: 10)

CuratorMaxValidations per curator (default: 50)

5.2 Duplicate Prevention

Every signal contains a unique hash → duplicates are rejected.

5.3 Tier Controls

Higher-tier curators can validate more signals but gain no extra rewards.

5.4 Sybil Resistance

rate limits

curator verification

human validation requirement

5.5 Behavioural Protections

Prevents:

self-curation

circular validation

bot activity

mass spam

NGO over-validation abuses

6. Reward Mechanism

Rewards are deterministic and auditable.

Let:

W = weight of a signal

W_total_day = total weights for the epoch

R_day = epoch reward budget (after halving)

Then:

R_signal = (W / W_total_day) × R_day
ParticipantReward = R_signal × 0.70
CuratorReward = R_signal × 0.30

Rewards are paid from the PoSS Reserve Address.

7. Module Parameters

All PoSS parameters (except immutable rules) are governed through on-chain governance.

7.1 Global Parameters
Parameter	Description
DailyMaxSignals	Participant max signals/day
CuratorMaxValidations	Curator max validations/day
WeightMicroDonation	1
WeightParticipation	2
WeightCertifiedContent	3
WeightCCN	5
ParticipantRatio	0.70
CuratorRatio	0.30
EpochDuration	24h
HalvingBlocks	Blocks in 8 years
ReservePoolAddress	PoSS reserve
StimulusPoolAddress	Early-adoption pool
Immutable

70/30 split

fixed supply

80% PoSS reserve

8-year halving

8. Halving Mechanism

Every 8 years, PoSS rewards are divided by two.

Formula:
Reward = BaseReward / (2^(Years / 8))

Results:

predictable reward curve

long-term sustainability

emission potential over 30–40 years

PoSS never introduces inflation.

9. Event System

PoSS emits transparent events for explorers, dApps and audits.

9.1 EventSignalEmitted

sender

type

weight

timestamp

9.2 EventSignalValidated

curator

participant

hash

timestamp

9.3 EventRewardDistributed

participant reward

curator reward

total weight

remaining reserve

9.4 EventAdminParamsUpdated

old params

new params

admin address

10. Queries (gRPC / REST)

QuerySignals(address)

QueryCurator(address)

QueryParams()

QueryStats()

QueryRewardHistory(address)

Compatible with explorers, dashboards and analytics tools.

11. Transaction Lifecycle
Step 1 — Emit Signal

Participant submits MsgEmitSignal.

Step 2 — Curate Signal

Curator validates through MsgValidateSignal.

Step 3 — Anti-Abuse Checks

Daily limits, hash uniqueness, curator tier.

Step 4 — Compute Reward

Weight × halving × daily budget.

Step 5 — Distribute Reward

70% to participant, 30% to curator.

Step 6 — Emit Events

Traceability guaranteed.

12. Administrative Logic

Admin addresses (defined in genesis) may:

update PoSS parameters

adjust curator tiers

modify daily limits

operate Stimulus Pool (within strict boundaries)

They cannot modify immutable rules.

13. Module State (Conceptual)

PoSS stores:

signal records

signal hashes

curator tiers

daily counters

reward history

PoSS parameters

pointer to reserve balance (Bank module)

No oracles or external systems required.

14. Security Model

PoSS is designed to ensure:

deterministic reward calculation

no privileged minting

halving strictly enforced

Sybil resistance

no oracle dependence

transparent state

no custody of user funds

15. Summary

NOORCHAIN — PoSS Full Specification (Phase 3.04, v1.1) defines:

all signal types & weights

complete reward model

reward split (immutable 70/30)

halving every 8 years

daily limits

anti-abuse systems

events and queries

transaction lifecycle

governance parameters

compliance alignment

the complete conceptual model for implementation

This is the canonical reference for Testnet, Mainnet, governance and auditors.
