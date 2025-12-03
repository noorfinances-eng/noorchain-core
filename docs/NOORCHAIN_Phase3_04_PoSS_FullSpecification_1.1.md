# NOORCHAIN 1.0  
## Phase 3.04 — PoSS Full Specification  
### Version 1.1  
### Last Updated: 2025-12-03  

---

# 0. Purpose of this Document

This document defines the **complete, authoritative specification** of  
NOORCHAIN’s Proof of Signal Social (PoSS) mechanism.

It serves as the:

- reference for implementing the PoSS module (`x/noorsignal`)
- foundation for Testnet and Mainnet genesis parameters
- normative source for audits and governance reviews
- canonical specification for reward logic, halving, and anti-abuse rules
- public documentation used in Whitepapers and Phase 6 Genesis Pack

No code is included.  
This is a **conceptual and functional specification**.

---

# 1. PoSS Overview

PoSS (Proof of Signal Social) is NOORCHAIN’s **human-centric, non-financial  
reward mechanism** designed to recognize and incentivize verified positive  
social actions.

PoSS is built on the following principles:

- **human-validated signals**
- **curator verification (Bronze / Silver / Gold)**
- **fully transparent on-chain scoring**
- **fixed total supply (299,792,458 NUR)**
- **halving every 8 years**
- **zero inflation (pre-allocated PoSS Reserve)**
- **70% / 30% reward split**
- **Legal Light CH compliance (non-financial, capped, symbolic)**

PoSS is **not** a mining system:  
it is a *distribution model* for a pre-existing pool of tokens.

---

# 2. Core Concepts

## 2.1 Signal
A signal is the on-chain representation of a positive, verifiable social action.  
It is emitted by a participant and later validated by a curator.

## 2.2 Participant
Any user who emits a PoSS signal.

## 2.3 Curator
A verified human or organization responsible for validating signals.  
Curators are assigned a tier:

- **Bronze**
- **Silver**
- **Gold**

Tiers affect **permissions and trust level**, not reward amount.

## 2.4 PoSS Reserve (80% of supply)
A fixed, pre-allocated pool of NUR used for PoSS rewards.  
This reserve is immutable and defined at genesis.

## 2.5 Reward Split (Immutable)
Every validated signal distributes:

- **70%** → Participant  
- **30%** → Curator  

This rule is considered **structural** and cannot be altered by governance.

## 2.6 Epoch (1 day)
The minimal window used for:

- computing daily limits  
- aggregating network weights  
- tracking reward distribution  
- resetting counters  
- applying halving progression  

---

# 3. Signal Types (4 Categories)

PoSS supports four official types of signals, representing verified  
positive contributions.

| Type | Description | Examples |
|------|-------------|----------|
| **1. Micro-donation** | Small financial or symbolic gesture | 1 CHF to NGO, micro-tip |
| **2. Verified Participation** | Presence or contribution to community | volunteering, event participation |
| **3. Certified Content** | Curator-approved positive content | educational posts, social impact content |
| **4. CCN (Content Collaboration Noorchain)** | High-value educational/social content | CCN Studio, long-form content |

---

# 4. Weight System

Each signal type has a predefined weight impacting reward distribution:

| Signal Type | Weight |
|-------------|--------|
| Micro-donation | **x1** |
| Participation | **x2** |
| Certified Content | **x3** |
| CCN Signal | **x5** |

Total weight for the day determines the relative reward share.

---

# 5. Anti-Abuse Measures

PoSS integrates strict safeguards:

## 5.1 Daily Limits
- `DailyMaxSignals` per participant (default: **10**)
- Curators: `CuratorMaxValidations` (default: **50**)

## 5.2 Duplicate Prevention
Signals include a hash.  
Duplicate hashes are rejected.

## 5.3 Tier-based Controls
Higher curator tiers = more validations allowed, not more rewards.

## 5.4 Sybil Resistance
Rate limits + curator gating drastically reduce Sybil risk.

## 5.5 Economic & Behavioural Protections
Prevents:
- circular validation  
- self-curation  
- mass spam  
- NGO over-validation  
- bot activity  

---

# 6. Reward Mechanism

Rewards are deterministic and can be reproduced for audit.

Let:

- `W` = signal weight  
- `W_total_day` = total weights emitted in the epoch  
- `R_day` = epoch reward budget (after halving)

Then:

R_signal = (W / W_total_day) * R_day
ParticipantReward = R_signal * 0.70
CuratorReward = R_signal * 0.30


Rewards are emitted from the **PoSS Reserve Address**.

---

# 7. Module Parameters (ParamsKeeper)

All PoSS parameters are on-chain and governance-adjustable (except immutable rules).

### 7.1 Global Parameters

| Parameter | Description |
|----------|-------------|
| `DailyMaxSignals` | Max signals per participant per day |
| `CuratorMaxValidations` | Max validations per curator per day |
| `WeightMicroDonation` | 1 |
| `WeightParticipation` | 2 |
| `WeightCertifiedContent` | 3 |
| `WeightCCN` | 5 |
| `ParticipantRatio` | 0.70 |
| `CuratorRatio` | 0.30 |
| `EpochDuration` | 24h |
| `HalvingBlocks` | Blocks in 8 years |
| `ReservePoolAddress` | PoSS Reserve |
| `StimulusPoolAddress` | Early adoption pool |

### Immutable parameters:
- 70/30 split  
- fixed supply  
- 80% PoSS reserve  
- 8-year halving

---

# 8. Halving Mechanism

Halving reduces `R_day` every **8 years**.

Formula:
Reward = BaseReward / (2 ^ (Years / 8))


Effects:
- predictable long-term emission  
- sustainable PoSS Reserve usage  
- ~30–40 years of potential PoSS activity  

No inflation is ever introduced.

---

# 9. Event System

PoSS emits transparent events:

## 9.1 EventSignalEmitted
- sender  
- type  
- weight  
- timestamp  

## 9.2 EventSignalValidated
- curator  
- participant  
- hash  
- timestamp  

## 9.3 EventRewardDistributed
- participant reward  
- curator reward  
- total weight  
- reserve remaining  

## 9.4 EventAdminParamsUpdated
- old params  
- new params  
- admin address  

---

# 10. Queries (gRPC + REST)

PoSS exposes:

- `QuerySignals(address)`  
- `QueryCurator(address)`  
- `QueryParams()`  
- `QueryStats()`  
- `QueryRewardHistory(address)`  

Useful for explorers and analytics.

---

# 11. Transaction Lifecycle (Full Flow)

### Step 1 — Participant emits signal  
`MsgEmitSignal`

### Step 2 — Curator validates signal  
`MsgValidateSignal`

### Step 3 — Anti-abuse checks  
Daily limits, duplicate hash, curator tier

### Step 4 — Reward computation  
Based on weight × halving × daily budget

### Step 5 — Reward distribution  
70% to participant  
30% to curator

### Step 6 — Event emission  
Ensures traceability

---

# 12. Administrative Logic

Only authorized admin addresses may:

- update PoSS parameters  
- update curator tiers  
- modify daily limits  
- manage Stimulus Pool (within boundaries)

Admin addresses must be defined in **genesis**.

---

# 13. Module State (Conceptual)

The PoSS module stores:

- signals (with metadata)  
- hashes (for duplicate detection)  
- curator tiers and status  
- daily counters per address  
- reward history  
- PoSS parameters  
- reference to reserve balance (via bank module)  

No external data sources or oracles are required.

---

# 14. Security Model

PoSS ensures:

- deterministic reward calculation  
- no privileged minting  
- no oracle risk  
- Sybil resistance  
- immutable structural rules  
- predictable emission curve  
- strict governance boundaries  
- no custody of user funds  

---

# 15. Summary (Header)

**NOORCHAIN — PoSS Full Specification (Phase3_04, v1.1)**  
Defines:

- all signal types & weights  
- curator structure  
- reward formula  
- anti-abuse mechanisms  
- halving (8 years)  
- PoSS parameters  
- full transactional lifecycle  
- events & queries  
- administrative logic  
- security model  
- compliance alignment  

This is the canonical reference for implementation, testnet, and audits.


