NOORCHAIN — Master Index
Phase 5 — A1: Global Table of Contents
Version 1.1
Last Updated: 2025-12-03
1. Purpose of this Index

This document serves as the master index for all NOORCHAIN documentation stored in the /docs directory. It provides:

a unified reference for navigation,

a mapping between documents and project phases,

a quality control tool for coherence between technical, legal, governance and PoSS materials,

preparation for Phase 6 (Genesis Pack & Website),

preparation for Phase 7 (Mainnet Launch).

Only documents actually present in the repository are listed.

2. Global Project Status

File:
NOORCHAIN_Phases_Status_1.1.md

Includes:

the full list of phases (1 → 9 + Phase 10)

completion percentage per phase

dependencies between phases

3. Phase 1 — Vision, Economic Model & PoSS Foundations

Main documents:

NOORCHAIN_Phase3_01_VersionsBase_1.1.md

NOORCHAIN_Phase3_04_Economic_Model_1.1.md

NOORCHAIN_PoSS_Logic_Overview_1.1.md

Covers:

Vision (ethical Swiss blockchain with PoSS)

Fixed supply: 299,792,458 NUR

Halving every 8 years

Economic allocation model (5 / 5 / 5 / 5 / 80)

Full conceptual basis of PoSS

4. Phase 2 — Technical Skeleton (Base Architecture)

There is no dedicated Phase 2 document.
Its content was fully absorbed into the later technical specifications:

NOORCHAIN_Phase3_01_Architecture_1.1.md

NOORCHAIN_Phase3_03_EthermintArchitecture_1.1.md

Phase4_01_CosmosEVM_Development_Blueprint_1.1.md

Phase4_02_App_Architecture_Map_1.1.md

Covers:

Cosmos SDK + Ethermint structure

High-level system design (modules, keepers, stores)

Application architecture before coding

5. Phase 3 — Specifications & Full PoSS Design
5.1 Core Architecture

NOORCHAIN_Phase3_01_Architecture_1.1.md

NOORCHAIN_Phase3_01_VersionsBase_1.1.md

NOORCHAIN_Phase3_02_ArchitectureProjet_1.1.md

Topics:

Core architecture

Environment, tools, dependencies

Module responsibilities

5.2 Genesis Specifications & Economic Model

NOORCHAIN_Phase3_02_Genesis_1.1.md

NOORCHAIN_Phase3_04_Economic_Model_1.1.md

NOORCHAIN_Phase3_05_GenesisTestnetSpecification_1.1.md

Topics:

Genesis economic rules

Allocation 5/5/5/5/80

Testnet genesis structure

5.3 PoSS Detailed Specifications

NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md

NOORCHAIN_Phase3_04_PoSS_FullSpecification_1.1.md

NOORCHAIN_Phase3_05_PoSS_Logic_1.1.md

NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md

NOORCHAIN_Phase3_07_PoSS_Params_and_KeeperTests_1.1.md

NOORCHAIN_PoSS_Logic_Overview_1.1.md

Topics:

PoSS model

70/30 reward split

Halving (8 years)

Score weights

Daily counters & anti-abuse

Keeper parameter tests

5.4 Testnet Flow & Practical Specifications

NOORCHAIN_Phase3_06_MsgCreateSignal_Spec_1.1.md

NOORCHAIN_Phase3_06_PoSS_Testnet_PracticalGuide_1.1.md

NOORCHAIN_Phase3_06_TransactionalFlow_1.1.md

NOORCHAIN_Phase3_08_Testnet_Minimal_1.1.md

Topics:

Signal transaction lifecycle

Tests & simulation

Minimal testnet configuration

6. Phase 4 — Implementation Architecture & PoSS Module
6.1 Cosmos / EVM App Architecture

NOORCHAIN_Phase4_EVM_Core_1.1.md

Phase4_01_CosmosEVM_Development_Blueprint_1.1.md

Phase4_02_App_Architecture_Map_1.1.md

Phase4_03_ModuleManager_Design_1.1.md

Phase4_04_Keepers_System_Design_1.1.md

Phase4_05_App_Initialization_Flow_1.1.md

Phase4_06_Block_Lifecycle_1.1.md

Phase4_07_Store_and_State_Model_1.1.md

Phase4_08_Final_Requirements_Checklist_1.1.md

Covers:

Ethermint integration

ModuleManager

Stores, state model

BeginBlock / EndBlock logic

6.2 PoSS Module Implementation

NOORCHAIN_Phase4_PoSS_Implementation_Status_1.1.md

NOORCHAIN_Phase4_Gov_Params_1.1.md

NOORCHAIN_Phase4B_PoSS_Msg_Status_1.1.md

NOORCHAIN_Phase4C_PoSS_Msg_Testnet_1.1.md

Phase4B_01_PoSS_Architecture_and_State_Model_1.1.md

Phase4B_02_PoSS_Block_Logic_and_Rewards_1.1.md

Phase4B_03_PoSS_Messages_Queries_and_Genesis_1.1.md

Covers:

PoSS keeper, rewards, daily counters

Messages, queries, events

Admin controls

PoSS genesis parameters

6.3 Testnet Architecture

Phase4C_01_Testnet_Addresses_and_Keys_Plan_1.1.md

Phase4C_02_Testnet_Genesis_Structure_and_Allocation_1.1.md

Phase4C_03_Testnet_Node_Setup_Commands_and_Network_Spec_1.1.md

Covers:

Address planning

Testnet allocation

Node configuration

7. Phase 5 — Governance, Legal & Compliance
7.1 Governance

Governance_Charter_1.1.md

MultiSig_Committee_1.1.md

Genesis_Governance_1.1.md

Foundation_Statutes_1.1.md

Foundation_Creation_1.1.md

Topics:

Foundation bodies

Multi-sig (3/5)

Genesis governance rules

Statutes & legal structure

7.2 Legal & Compliance

Legal_Light_2025_1.1.md

Legal_Notices_1.1.md

Compliance_Framework_1.1.md

Legal_Architecture_1.1.md

Topics:

Legal classification

Transparency rules

Swiss Legal Light model

Mapping between PoSS, governance and law

7.3 Cross-Checks

Phase5_CrossCheck_1.1.md

Ensures consistency across:

PoSS

Legal

Governance

Genesis

Addresses

8. Master PDFs

NOORCHAIN_MASTER_FULL_A.pdf

NOORCHAIN_MASTER_FULL_B.pdf

Used for:

partners

auditors

internal synthesis

9. Phase 6+ (Upcoming)

To be created:

Genesis Pack (public version)

Economic Model (public version)

Legal Light Public PDF

Website content (all sections)

Mainnet Preparation documentation

10. Address Integration Reminder (Phase 7)

When the 5 real bech32 addresses are generated, update:

Genesis_Governance_1.1.md

Legal_Architecture_1.1.md

MultiSig_Committee_1.1.md

Foundation_Statutes_1.1.md

All genesis files

x/noorsignal/types/addresses.go

Whitepaper 1.1

Website pages

11. Signature

Prepared by:
NOORCHAIN Foundation — Governance & Legal Phase 5
Version 1.1
