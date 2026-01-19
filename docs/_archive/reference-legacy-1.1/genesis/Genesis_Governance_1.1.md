# NOORCHAIN FOUNDATION  
## Genesis Governance  
### Phase 5 — C3  
### Version 1.1  
### Last Updated: 2025-XX-XX  

---

# 1. Purpose of Genesis Governance

This document defines the **initial governance configuration** included in the  
**Mainnet Genesis Block** (“genesis.json”).  
It ensures:

- Transparent initial allocations  
- Clear governance responsibilities  
- Legal Light CH compliance  
- Mission-oriented distribution  
- Proper initialization of PoSS parameters  
- Correct mapping of institutional addresses  
- Irreversibility of core governance principles  

Genesis Governance is the backbone of NOORCHAIN’s credibility and legal safety.

---

# 2. Immutable Genesis Principles

The following principles are embedded and cannot be changed post-launch:

1. **Fixed Supply:**  
   Total cap = **299,792,458 NUR** (“speed of light supply”).

2. **Halving Every 8 Years:**  
   PoSS issuance decreases on a fixed cycle.

3. **No Inflation Beyond PoSS:**  
   No discretionary minting is allowed.

4. **5/5/5/5/80 Allocation:**  
   - 5% Foundation  
   - 5% Noor Dev Sàrl (on-chain functional allocation)  
   - 5% PoSS Stimulus  
   - 5% Pre-sale (optional, vesting required)  
   - 80% PoSS Mintable Reserve  

5. **PoSS Reward Split:**  
   - **70% Participant**  
   - **30% Curator**  
   (Per-signal split is structural and immutable)

6. **No Custody:**  
   The Foundation may not hold user assets.

7. **No Investment Product:**  
   The supply distribution cannot be used for any yield-guaranteeing scheme.

8. **Legal Light CH Compatibility:**  
   Everything in the genesis follows Swiss non-profit and non-custodial rules.

---

# 3. Genesis Addresses (to be updated in Phase 7)

The following addresses are placeholders and will be replaced with  
the real bech32 addresses in Phase 7.

### 3.1 Foundation Address  
- `noor1foundationxxxxxxxxxxxxxxxxxxxxx`

Purpose: Holds 5% supply, managed only through Multi-sig (3/5).

### 3.2 Noor Dev Sàrl Address (on-chain functional allocation)  
- `noor1devsarlxxxxxxxxxxxxxxxxxxxxxxxxxxx`

Purpose: Receives 5% supply.  
Represents the on-chain economic allocation of the development entity.

### 3.3 PoSS Stimulus Reserve  
- `noor1stimulusxxxxxxxxxxxxxxxxxxxxxxxx`

Purpose: 5% supply locked for the first ecosystem partners.

### 3.4 Pre-sale Optional Pool  
- `noor1presalexxxxxxxxxxxxxxxxxxxxxxxxxx`

Purpose: 5% supply with mandatory vesting rule.  
Locked until explicit governance approval.

### 3.5 PoSS Mintable Reserve (80%)  
- `noor1possreservexxxxxxxxxxxxxxxxxxxxxx`

Purpose: Main PoSS issuance pool.  
Minted gradually under PoSS module rules.

---

# 4. Genesis Allocations (Hard-coded)

### 4.1 Summary Table

| Category                  | Percentage | Supply (NUR)               |
|--------------------------|------------|-----------------------------|
| Foundation               | 5%         | 14,989,623                 |
| Noor Dev Sàrl (on-chain) | 5%         | 14,989,623                 |
| PoSS Stimulus            | 5%         | 14,989,623                 |
| Pre-sale Optional        | 5%         | 14,989,623                 |
| PoSS Mintable Reserve    | 80%        | 239,833,984                |
| **Total**                | **100%**   | **299,792,458**            |

These allocations are immutable and included directly in the genesis.json.

---

# 5. Genesis Governance Powers

At the genesis block, governance powers are split as follows:

### 5.1 Foundation (Board)
- Approves administrative actions  
- Oversees documentation and mission  
- Supervises Curators  
- Publishes transparency reports  

### 5.2 Multi-sig Committee (3/5)
- Executes all on-chain Foundation decisions  
- Controls the Foundation Address  
- Can modify PoSS parameters (within governance limits)  
- Can pause PoSS in emergencies  
- Approves deployment of new modules  

### 5.3 Noor Dev Sàrl
- No governance power at genesis  
- Holds on-chain functional allocation (5%)  
- Contributes technical improvements  
- May submit proposals  

### 5.4 Curators
- Zero governance power at genesis  
- Advisory voting only  
- Influence future governance decentralization  

---

# 6. Genesis Parameters (PoSS & Protocol)

### 6.1 PoSS Parameters at Genesis
(Values adjustable by governance post-launch)

- `PoSSEnabled`: **false** (PoSS activates after mainnet stability)  
- `BaseReward`: `1 unur`  
- `WeightMicroDonation`: `2`  
- `WeightParticipation`: `1`  
- `WeightContent`: `3`  
- `WeightCCN`: `4`  
- `MaxSignalsPerDay`: `10`  
- `MaxSignalsPerCuratorPerDay`: `20`  
- `MaxRewardPerDay`: `100 unur`  
- `HalvingPeriodYears`: `8`  

### 6.2 Economic Constraints
- PoSS cannot exceed the 80% reserve  
- No external minting  
- No discretionary pool extension  

---

# 7. Governance Limits (Hard-coded)

The following restrictions are encoded in governance logic:

1. No change to fixed supply  
2. No change to PoSS structural split (70/30)  
3. No allocation minting beyond genesis  
4. No deletion or reduction of Foundation 5% or Sàrl 5%  
5. No override of Legal Light restrictions  
6. No reassignment of genesis addresses without full governance quorum  

---

# 8. Upgrade Path (Post-Genesis)

Genesis Governance enables the following future upgrades:

1. Enable PoSS after stability phase  
2. Add new Curators  
3. Update PoSS limits and base reward  
4. Deploy additional modules (Hub, Studio, Pay)  
5. Add secondary governance mechanisms  
6. Extend multi-sig signers  
7. Introduce advisory voting by Curators  

---

# 9. Documentation Requirements

The following documents must be published at or before mainnet:
- Governance Charter (C2)  
- Multi-sig Committee Charter (C1)  
- Legal Light PDF  
- Genesis Allocation PDF  
- Foundation Statutes  

These form the complete **Genesis Governance Pack**.

---

# 10. Adoption

This document is adopted by the  
**NOORCHAIN Foundation Board**  
and becomes part of the **Genesis Pack** for mainnet launch.

Signatures:

_____________________________  
_____________________________  
_____________________________

Version 1.1  
Prepared for NOORCHAIN Foundation — Governance Phase
