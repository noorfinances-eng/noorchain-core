# NOORCHAIN — Master Index  
## Phase 5 — A1: Global Table of Contents  
### Version 1.1  
### Last Updated: 2025-12-03  

---

# 1. Purpose of this Index

This document is the **master index (v1.1)** for all official NOORCHAIN materials currently stored in the `/docs` folder.

It is used to:

- navigate the documentation efficiently  
- link each file to the corresponding project phase  
- check consistency between technical, legal, governance and PoSS documents  
- prepare Phase 6 (Genesis Pack + Website) and Phase 7 (Mainnet)  
- present a clear structure to external reviewers, partners and auditors  

Only files that **actually exist** in the repository are listed here.

---

# 2. Global Project Status

**File:**

- `docs/NOORCHAIN_Phases_Status_1.1.md`

Contents:

- Official list of phases (1 → 9 + Phase 10 Interop)  
- Current status (Done / In progress / Not started)  
- High-level description of each phase  
- Dependencies between phases  

---

# 3. Phase 1 — Vision, Model & PoSS Concept

Phase 1 (Cadrage & Decisions) is documented mainly inside the Phase 3 conceptual files.

**Key references:**

- `docs/NOORCHAIN_Phase3_01_VersionsBase_1.1.md`  
- `docs/NOORCHAIN_Phase3_04_Economic_Model_1.1.md`  
- `docs/NOORCHAIN_PoSS_Logic_Overview_1.1.md`  

Contents:

- Vision of NOORCHAIN (Swiss ethical blockchain with PoSS)  
- Fixed supply (299,792,458 NUR)  
- Halving every 8 years  
- Economic allocation model (5 / 5 / 5 / 5 / 80)  
- Social impact model (PoSS, curators, participants)  
- Non-speculative, Legal Light CH approach  

---

# 4. Phase 2 — Technical Skeleton (Base Architecture)

There is no separate “Phase2_*.md” file because the skeleton has been captured later inside Phase 3 and Phase 4 documents.

**Main references for the skeleton:**

- `docs/NOORCHAIN_Phase3_01_Architecture_1.1.md`  
- `docs/NOORCHAIN_Phase3_03_EthermintArchitecture_1.1.md`  
- `docs/Phase4_01_CosmosEVM_Development_Blueprint_1.1.md`  
- `docs/Phase4_02_App_Architecture_Map_1.1.md`  

Contents:

- Cosmos SDK + Ethermint architecture  
- App layout (modules, keepers, stores)  
- General EVM integration plan  
- High-level technical blueprint before coding  

---

# 5. Phase 3 — Full Specifications & PoSS Design

## 5.1 Core Architecture & Versions

- `docs/NOORCHAIN_Phase3_01_Architecture_1.1.md`  
- `docs/NOORCHAIN_Phase3_01_VersionsBase_1.1.md`  
- `docs/NOORCHAIN_Phase3_02_ArchitectureProjet_1.1.md`  

Contents:

- Core NOORCHAIN architecture  
- Environment, versions and technical constraints  
- Module layout and responsibilities  

## 5.2 Genesis & Economic Model

- `docs/NOORCHAIN_Phase3_02_Genesis_1.1.md`  
- `docs/NOORCHAIN_Phase3_04_Economic_Model_1.1.md`  
- `docs/NOORCHAIN_Phase3_05_GenesisTestnetSpecification_1.1.md`  

Contents:

- Genesis economic rules  
- Allocation model 5 / 5 / 5 / 5 / 80  
- Testnet genesis structure and constraints  

## 5.3 PoSS Specifications

- `docs/NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md`  
- `docs/NOORCHAIN_Phase3_04_PoSS_FullSpecification_1.1.md`  
- `docs/NOORCHAIN_Phase3_05_PoSS_Logic_1.1.md`  
- `docs/NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md`  
- `docs/NOORCHAIN_Phase3_07_PoSS_Params_and_KeeperTests_1.1.md`  
- `docs/NOORCHAIN_PoSS_Logic_Overview_1.1.md`  

Contents:

- Full PoSS model  
- Rewards logic (including 70% / 30%)  
- Halving rule (8 years)  
- Anti-abuse mechanisms  
- Keeper tests and parameterisation  
- PoSS testnet status  

## 5.4 Transactional & Testnet Flow

- `docs/NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md`  
- `docs/NOORCHAIN_Phase3_06_MsgCreateSignal_Spec_1.1.md`  
- `docs/NOORCHAIN_Phase3_06_PoSS_Testnet_PracticalGuide_1.1.md`  
- `docs/NOORCHAIN_Phase3_06_TransactionalFlow_1.1.md`  
- `docs/NOORCHAIN_Phase3_08_Testnet_Minimal_1.1.md`  

Contents:

- MsgCreateSignal specification  
- Transactional life-cycle of signals  
- Minimal testnet requirements  
- Testnet practical guides (how to send signals, etc.)  

---

# 6. Phase 4 — Implementation & Core App Architecture

## 6.1 Cosmos/EVM Core & App Lifecycle

- `docs/NOORCHAIN_Phase4_EVM_Core_1.1.md`  
- `docs/Phase4_01_CosmosEVM_Development_Blueprint_1.1.md`  
- `docs/Phase4_02_App_Architecture_Map_1.1.md`  
- `docs/Phase4_03_ModuleManager_Design_1.1.md`  
- `docs/Phase4_04_Keepers_System_Design_1.1.md`  
- `docs/Phase4_05_App_Initialization_Flow_1.1.md`  
- `docs/Phase4_06_Block_Lifecycle_1.1.md`  
- `docs/Phase4_07_Store_and_State_Model_1.1.md`  
- `docs/Phase4_08_Final_Requirements_Checklist_1.1.md`  

Contents:

- Ethermint integration  
- App structure (BaseApp, keepers, modules)  
- Module manager design  
- Store and state model  
- Block lifecycle and handlers  
- Final checklist for core implementation  

## 6.2 PoSS Implementation Status

- `docs/NOORCHAIN_Phase4_PoSS_Implementation_Status_1.1.md`  
- `docs/NOORCHAIN_Phase4_Gov_Params_1.1.md`  
- `docs/NOORCHAIN_Phase4B_PoSS_Msg_Status_1.1.md`  
- `docs/NOORCHAIN_Phase4C_PoSS_Msg_Testnet_1.1.md`  
- `docs/Phase4B_01_PoSS_Architecture_and_State_Model_1.1.md`  
- `docs/Phase4B_02_PoSS_Block_Logic_and_Rewards_1.1.md`  
- `docs/Phase4B_03_PoSS_Messages_Queries_and_Genesis_1.1.md`  

Contents:

- PoSS messages, queries and genesis wiring  
- PoSS rewards logic implementation  
- Current implementation status of PoSS module  
- Governance parameters related to PoSS  

## 6.3 Testnet Architecture & Node Setup

- `docs/Phase4C_01_Testnet_Addresses_and_Keys_Plan_1.1.md`  
- `docs/Phase4C_02_Testnet_Genesis_Structure_and_Allocation_1.1.md`  
- `docs/Phase4C_03_Testnet_Node_Setup_Commands_and_Network_Spec_1.1.md`  

Contents:

- Testnet address & keys plan  
- Testnet genesis structure & allocation details  
- Node setup commands and minimal network specification  

---

# 7. Phase 5 — Governance, Legal & Compliance

Phase 5 is fully documented in the following files:

## 7.1 Index & Timestamp

- `docs/NOORCHAIN_Index_1.1.md`  *(this file)*  
- `docs/Phase5_Timestamp_1.1.md`  

## 7.2 Foundation & Governance

- `docs/Foundation_Statutes_1.1.md`  
- `docs/Foundation_Creation_1.1.md`  
- `docs/Governance_Charter_1.1.md`  
- `docs/MultiSig_Committee_1.1.md`  
- `docs/Genesis_Governance_1.1.md`  

Contents:

- Foundation structure and statutes  
- Foundation creation process  
- Governance charter (bodies, roles, quorum)  
- Multi-sig committee definition (3/5)  
- Genesis governance rules  

## 7.3 Legal & Compliance

- `docs/Legal_Light_2025_1.1.md`  
- `docs/Legal_Notices_1.1.md`  
- `docs/Compliance_Framework_1.1.md`  
- `docs/Legal_Architecture_1.1.md`  

Contents:

- Legal Light CH classification  
- Legal notices for the website and public docs  
- Compliance framework (no custody, no yield, no PSP internal)  
- Mapping between legal framework, PoSS and governance  

## 7.4 Final Cross-Check

- `docs/Phase5_CrossCheck_1.1.md`  

Contents:

- Final validation of legal, governance, PoSS and genesis consistency  
- List of items to update when real addresses are provided  

---

# 8. Master PDFs (Global View)

These PDFs compile multiple text documents into complete reference packs:

- `docs/NOORCHAIN_MASTER_FULL_A.pdf`  
- `docs/NOORCHAIN_MASTER_FULL_B.pdf`  

Usage:

- Internal review  
- External sharing with advisors / partners  
- High-level view of the technical and legal design  

---

# 9. Future Phase 6+ Documents (To Be Created)

The following documents are **not yet present** in `/docs` but are planned:

- Phase 6 — Genesis Pack (public-oriented):  
  - Genesis Allocation Pack 1.1  
  - Public Economic Model 1.1  
  - Governance & Curators Overview  
  - Legal Light Public PDF  

- Phase 6 — Website content:  
  - Site architecture  
  - Site content draft (Home, /tech, /poss, /curators, /genesis, /roadmap, /docs, /legal, /foundation, /governance, /compliance)  

- Phase 7 — Mainnet Preparation:  
  - `Docs/Mainnet_Preparation_1.1.md` (to be created)  

---

# 10. Address Update Reminder (Phase 7)

When the 5 real bech32 addresses are generated (Foundation, Dev Sàrl, PoSS Stimulus, Pre-sale, PoSS Reserve), they must be injected into:

- `docs/Genesis_Governance_1.1.md`  
- `docs/Legal_Architecture_1.1.md`  
- `docs/MultiSig_Committee_1.1.md`  
- `docs/Foundation_Statutes_1.1.md`  
- `testnet/genesis.json`  
- `mainnet/genesis.json`  
- `x/noorsignal/types/addresses.go`  
- Phase 6 Genesis Pack documents  
- Whitepaper 1.1  
- Website pages `/genesis`, `/governance`, `/legal`, `/compliance`  

---

# 11. Signature

Prepared by:  
**NOORCHAIN Foundation — Governance & Legal Phase 5**  
Version 1.1  

