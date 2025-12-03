# NOORCHAIN FOUNDATION  
## Phase 5 — C1: Multi-sig Committee (3/5) Structure  
### Version 1.1  
### Last Updated: 2025-XX-XX  

---

## 1. Purpose of the Multi-sig Committee

The **Multi-sig Committee** is the on-chain executive organ of the  
**NOORCHAIN Foundation**, responsible for managing the Foundation Allocation,  
executing approved governance actions, and ensuring transparent stewardship  
of on-chain resources.

It operates with **3-out-of-5** signatures (3/5), following industry best practices for  
Swiss blockchain foundations (e.g., Tezos, Nym, NEAR).

The Committee serves as:
- The **operational arm** of the Foundation on-chain,  
- A **safety layer** for protocol actions,  
- A **mechanism of oversight** preventing unilateral decisions.

---

## 2. Multi-sig Configuration

### **2.1 Signers (Five Members)**  
Recommended composition:

1. **Founder / Project Lead**  
   - Technical and strategic oversight  
   - Ensures vision consistency  

2. **NGO / Social Impact Representative**  
   - Represents the Curators and PoSS ecosystem  
   - Ensures social alignment  

3. **Technical Expert / Developer**  
   - Reviews technical proposals  
   - Validates PoSS and governance updates  

4. **Legal / Compliance Observer**  
   - Ensures strict adherence to Swiss Legal Light CH  
   - Verifies no-custody, no-yield, no-investment messaging  

5. **Educational / Community Representative**  
   - Represents schools, educational partners, and community interest  

**Rotation:**  
Members 2–5 may rotate every 12–24 months.

### **2.2 Threshold Rule**
- **3-of-5** signatures required for any on-chain action.  
- Guarantees decentralization and accountability.  

---

## 3. On-chain Responsibilities

The Multi-sig Committee is authorised to:

### **3.1 Manage the Foundation Allocation**
5% of the Genesis supply, designated for:
- audits,  
- research,  
- documentation,  
- NGO partnerships,  
- outreach,  
- ecosystem grants.

### **3.2 Execute Approved Governance Actions**
Examples:
- Updating PoSS parameters (within allowed governance rules),  
- Enabling or disabling PoSS at genesis,  
- Initiating network upgrades,  
- Deploying approved contracts.

### **3.3 Publish Transparency Reports**
Frequency: every **6 or 12 months**.

Reports include:
- Multi-sig operations summary  
- Funds allocated  
- PoSS parameter changes  
- Governance actions executed  

---

## 4. Limitations & Legal Constraints

The Multi-sig MUST adhere to the following restrictions:

### **4.1 No Custody of User Funds**
The Multi-sig cannot:
- hold tokens belonging to users,  
- manage third-party assets,  
- operate as a custodial service.

### **4.2 No Promised Returns**
The Multi-sig cannot:
- distribute rewards beyond PoSS rules,  
- guarantee financial returns,  
- operate staking-as-a-service.

### **4.3 No Market Manipulation**
Prohibitions include:
- price control,  
- artificial liquidity operations,  
- speculative activities.

### **4.4 Compliance With Swiss Legal Light CH**
All actions must align with:
- non-profit purpose,  
- transparency requirements,  
- no investment offering,  
- no financial intermediation.

---

## 5. Allowed Operations (Explicit)

The Committee **may** execute:
1. Funding for audits  
2. Payment for documentation  
3. NGO / school partnership grants  
4. Technical bounties  
5. Infrastructure expenses (RPC, explorer, hosting)  
6. On-chain governance parameters update  
7. Operational support costs (within Foundation scope)

---

## 6. Forbidden Operations (Explicit)

The Committee **may not**:
1. Hold user tokens  
2. Provide loans or credit  
3. Issue promised or fixed-yield instruments  
4. Sell tokens as investment products  
5. Use Foundation funds for speculation  
6. Lock liquidity without governance approval  
7. Modify Genesis allocations post-mainnet  

---

## 7. Transparency Rules

The Multi-sig must:
- publish all transactions publicly,  
- maintain a clear operational log,  
- release periodic reports,  
- disclose any conflicts of interest,  
- document each parameter update,  
- maintain public addresses of signers.

---

## 8. Replacement & Rotation Rules

### **8.1 Replacement**
A signer may be replaced if:
- they resign,  
- they violate ethical principles,  
- they become inactive,  
- a conflict of interest arises.

### **8.2 Rotation**
Representatives ideally rotate every **12–24 months**, except the Founder.

Rotation ensures:
- decentralization,  
- independence,  
- continuity of governance.

---

## 9. Technical Implementation Notes

- The multi-sig will be implemented as a **Cosmos SDK multi-sig account**  
  (or equivalent Ethermint-compatible multi-sig mechanism).

- Signing keys are generated in Phase 7 (Mainnet Setup).

- The Foundation must maintain:
  - `multisig.json` (keyfile)  
  - Backup of each signer’s public key  
  - A registry of active and past members  

---

## 10. Adoption

This Multi-sig Committee structure is adopted by the  
**NOORCHAIN Foundation Board** as part of Phase 5.

### Signatures (Founding)

____________________________  
____________________________  
____________________________  

Version 1.1  
Prepared for the NOORCHAIN Foundation  
Governance & Legal Phase
