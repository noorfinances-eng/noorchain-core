 NOORCHAIN — Phase 5  
## A2 — Official Timestamp Procedure (Version 1.1)
### Category: Legal & Governance  
### Status: Mandatory  
### Last updated: 2025-XX-XX

---

## 1. Purpose of this Document

This document defines the **official timestamping procedure** for NOORCHAIN core materials.  
Timestamping provides **immutable proof of authorship and priority**, and is required before:

- drafting legal documents,  
- publishing governance rules,  
- releasing the Genesis Pack,  
- filing trademarks,  
- launching Phase 6 communication,  
- or raising external funds.

This step ensures **Swiss-grade legal protection** of the intellectual property.

---

## 2. Timestamp Method (Approved Approach)

NOORCHAIN uses a **dual-layer timestamp**:

### **Layer 1 — SHA256 Hash**
A cryptographic fingerprint of each document.

### **Layer 2 — IPFS Immutable Storage**
The hash is uploaded to IPFS, producing:

- Content Identifier (CID)  
- Permanent immutable link  
- Public proof of existence at a given time

This method is:
- Free  
- Decentralized  
- Cryptographically verifiable  
- Compliant with Swiss “proof of anteriority” practice  

---

## 3. Files to Timestamp (Mandatory List)

### **A. Technical Architecture**
- `docs/Phase1_Vision_1.1.md`
- `docs/Phase2_TechSkeleton_1.1.md`
- `docs/Phase3_Documentation_1.1.md`
- `docs/PoSS/PoSS_Params_1.1.md`
- `docs/PoSS/PoSS_Logic_1.1.md`
- `docs/PoSS/PoSS_GenesisRules_1.1.md`

### **B. Governance & Economics**
- `docs/Economics/Genesis_Allocations_1.1.md`
- `docs/Economics/Token_Model_1.1.md`
- `docs/Governance/Governance_Principles_1.1.md`
- `docs/Governance/Curators_Model_1.1.md`

### **C. Legal Structure**
- `docs/Legal/Legal_Light_Framework_1.1.md`
- `docs/Legal/Compliance_Rules_1.1.md`

### **D. Phase 4C Technical Blueprint**
- `docs/Phase4C_CoreArchitecture_1.1.md`
- `docs/Phase4C_AppModule_1.1.md`
- `docs/Phase4C_Noorsignal_Module_1.1.md`

⚠️ **Conversations ChatGPT are excluded** (per founder directive).

---

## 4. Procedure: Step-by-Step

### **Step 1 — Generate SHA256 hashes**
For each file:

```bash
sha256sum filename.ext
Store results in:

docs/Phase5_Timestamp_SHA256_1.1.txt

Format:

php-template
Copier le code
<sha256hash>  <filename>
Step 2 — Upload each file to IPFS
Using a public IPFS gateway (e.g., Web3Storage, Pinata, or local IPFS):

bash
Copier le code
ipfs add filename.ext
Result:

objectivec
Copier le code
added <CID> filename.ext
Store results in:

docs/Phase5_Timestamp_IPFS_1.1.txt

Format:

objectivec
Copier le code
filename.ext
CID: <cid_here>
IPFS URL: https://ipfs.io/ipfs/<cid_here>
Step 3 — Create the official timestamp record
Create:

docs/Phase5_Timestamp_Record_1.1.md

Format:

makefile
Copier le code
Document: <filename>
SHA256: <hash>
CID: <ipfs_hash>
Date: <YYYY-MM-DD>
Authority: NOORCHAIN (Founder)
Step 4 — Commit to GitHub
Commit message:

pgsql
Copier le code
Phase 5 — A2: Timestamp complete (Version 1.1)
Added SHA256, IPFS CIDs, and timestamp record.
Once pushed, GitHub itself provides a secondary timestamp.

5. Verification Procedure
Any third party can verify:

Download the file from GitHub

Compute SHA256

Compare to HASH in Timestamp_SHA256_1.1.txt

Retrieve file from IPFS using CID

Ensure IPFS content matches GitHub version

If both hashes are identical →
✔ Document is authentic
✔ Timestamp is valid
✔ Integrity guaranteed

6. Notes & Compliance
No personal data is stored on IPFS

No private keys are stored

Timestamping is open and transparent

Complies with Swiss “preuve d’antériorité”

This procedure must be completed before Phase 6 begins

7. Signature
Founder (Walid Barhoumi)
NOORCHAIN Project — Legal & Governance Phase
Version 1.1
