NOORCHAIN — Documentation Index (evm-l1)
Version: 2.1
Last Updated: 2026-01-19

Purpose
This index is the authoritative entry point for the NOORCHAIN 2.1 documentation maintained in /docs.
It focuses on the evm-l1 line (sovereign EVM L1) and the operational runbooks used to run and validate a mainnet-like local stack.

Scope (Active)
The documents listed below are considered active for NOORCHAIN 2.1 operations, validation, and governance/compliance framing:
- Node operations (multi-node pack, RPC, world-state, receipts/logs)
- Minimal dApp/tooling runbooks
- Genesis and governance/legal/compliance reference materials

Legacy documentation
Historical specifications and pilots from earlier iterations (Cosmos/Ethermint-era materials, older phase indices, legacy pilots) have been preserved and moved under:
  docs/_archive/

1) Operational Runbooks (Active)
Use these runbooks for reproducible operations and validation gates.

- RUNBOOK-M5.md
  Minimal multi-node networking and bootstrap discipline.

- RUNBOOK-M9-CONTRACTS-EXECUTION.md
  Minimal contract execution wiring (mining/applyTx hook) and receipt persistence gates.

- RUNBOOK-M10-MAINNETLIKE.md
  Two-node mainnet-like pack: leader/follower behavior, P2P, RPC stability, health endpoints.

- RUNBOOK-M11-DAPPS.md
  Local tooling and dApp scripts used against the running nodes.

- RUNBOOK-M12-WORLDSTATE-RPC.md
  World-state integration (Geth StateDB/triedb), stateRoot exposure, and RPC read parity gates.

- RUNBOOK-M13-VALIDATION.md
  Post-tag validation checklist for the current baseline.

- RUNBOOK-PHASE7-PROOF-OF-LIVENESS.md
  Controlled proof-of-liveness procedure (minimal exposure).

2) Snapshots / Notes (Active)
- M6-A.1-SNAPSHOT.md
  Historical checkpoint notes used during the evm-l1 build-up.

3) Genesis Reference (Active)
These documents capture the intended genesis posture and supporting governance/security framing.

- genesis/Genesis_Overview_1.1_EN.md
- genesis/Genesis_Overview_1.1_FR.md

- genesis/Genesis_Allocation_1.1_EN.md
- genesis/Genesis_Allocation_1.1_FR.md

- genesis/Genesis_Parameters_1.1.md
- genesis/Genesis_Parameters_1.1_FR.md

- genesis/Genesis_Governance_1.1.md
- genesis/Genesis_Governance_1.1_FR.md

- genesis/Genesis_Security_Overview_1.1.md
- genesis/Genesis_Security_Overview_1.1_FR.md

- genesis/Genesis_Migration_Path_1.1.md
- genesis/Genesis_Migration_Path_1.1_FR.md

4) Governance (Active)
- governance/Governance_Charter_1.1.md
- governance/MultiSig_Committee_1.1.md

5) Legal (Active)
- legal/Legal_Architecture_1.1.md
- legal/Legal_Light_2025_1.1.md
- legal/Legal_Notices_1.1.md
- legal/Foundation_Creation_1.1.md
- legal/Foundation_Statutes_1.1.md

6) Compliance (Active)
- compliance/Compliance_Framework_1.1.md

7) French Entry Point
- fr/README_FR.md
  French navigation entry point for selected documents.

