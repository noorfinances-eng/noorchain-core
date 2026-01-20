NOORCHAIN 2.1 — Documentation Index

Document ID: NOORCHAIN_INDEX_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Canonical index for the NOORCHAIN 2.1 documentation set (technical + operations + governance + compliance).

1. Purpose

This file is the single entry point to the NOORCHAIN 2.1 documentation set.

It provides:

The authoritative document map (what exists, what is normative, what is guidance)

A stable reading order for different audiences

A locked list of files that define the 2.1 documentation perimeter

Maintenance rules to keep the set coherent over time

This index is intended to reduce ambiguity, prevent duplicate specs, and support audit readiness.

2. Audience and Reading Paths
2.1 Operators (run and keep it healthy)

Recommended order:

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/ops/INCIDENTS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

Runbooks under docs/RUNBOOK-*

2.2 Integrators (wallets, tooling, clients)

Recommended order:

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/dev/JSON_RPC_EXAMPLES_2.1.md

docs/dev/TOOLING_HARDHAT_VIEM_2.1.md

2.3 Protocol / Core Review (architecture and state)

Recommended order:

docs/ARCHITECTURE_2.1.md

docs/STATE_MODEL_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/AUDIT_READINESS_2.1.md

2.4 Governance / Compliance

Recommended order:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/governance/MULTISIG_OPERATIONS_2.1.md

docs/governance/UPGRADE_PROCESS_2.1.md

docs/compliance/COMPLIANCE_FRAMEWORK_2.1.md

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

docs/legal/IP_AND_BRAND_2.1.md

3. Document Roles (Normative vs. Guidance)
3.1 Normative Documents

Normative documents define required behavior or process. In case of conflict, normative documents prevail.

docs/RPC_SPEC_2.1.md

docs/STATE_MODEL_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

Governance and compliance documents under docs/governance/* and docs/compliance/*

3.2 Guidance Documents

Guidance documents provide operational or developer workflows; they must remain consistent with the normative layer.

Runbooks docs/RUNBOOK-*.md

Developer quickstart and tooling docs under docs/dev/*

Operations playbooks under docs/ops/*

Changelog and index documents

4. Locked File Set (2.1 Perimeter)

The following list is the locked perimeter for the “Docs 2.1 — files-to-fill” workstream.
No additional files are introduced as part of this sequence unless explicitly authorized.

docs/API_STABILITY_POLICY_2.1.md

docs/AUDIT_READINESS_2.1.md

docs/CHANGELOG_2.1.md

docs/NOORCHAIN_Index_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

docs/README.md

docs/RELEASE_PROCESS_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

docs/compliance/COMPLIANCE_FRAMEWORK_2.1.md

docs/dev/DEV_QUICKSTART_2.1.md

docs/dev/JSON_RPC_EXAMPLES_2.1.md

docs/dev/TOOLING_HARDHAT_VIEM_2.1.md

docs/genesis/ALLOC_POLICY_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/governance/MULTISIG_OPERATIONS_2.1.md

docs/governance/UPGRADE_PROCESS_2.1.md

docs/legal/IP_AND_BRAND_2.1.md

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/ops/INCIDENTS_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

5. Current Documentation Map (High-Level)
5.1 Core (Protocol / Node)

docs/ARCHITECTURE_2.1.md — system boundaries, components, roles

docs/STATE_MODEL_2.1.md — world-state, receipts, roots, persistence model

docs/SECURITY_TRUST_MODEL_2.1.md — trust assumptions, security boundaries

docs/THREAT_MODEL_2.1.md — threats, mitigations, residual risks

5.2 RPC (Client Interface)

docs/RPC_SPEC_2.1.md — JSON-RPC behavior and semantics

docs/rpc/RPC_ERROR_MODEL_2.1.md — error taxonomy and handling rules

docs/rpc/RPC_COMPAT_MATRIX_2.1.md — compatibility gates and client coverage

docs/API_STABILITY_POLICY_2.1.md — stability levels and deprecation discipline

5.3 Operations

docs/OPERATIONS_PLAYBOOK_2.1.md — operator baseline, procedures, gates

docs/ops/DEPLOYMENT_MODEL_2.1.md — deployment topologies and assumptions

docs/ops/INCIDENTS_2.1.md — incident catalog and operational responses

Runbooks: docs/RUNBOOK-*.md — step-by-step validated procedures

5.4 Genesis / Parameters

docs/genesis/ALLOC_POLICY_2.1.md — allocation policy and operator rules

docs/genesis/GENESIS_SPEC_2.1.md — genesis format, required fields, invariants

docs/genesis/GENESIS_CHECKLIST_2.1.md — pre-genesis validation checklist

docs/genesis/PARAMETERS_REFERENCE_2.1.md — chain parameters reference

5.5 Governance / Compliance / Legal

Governance: model, multisig operations, upgrades, incident response

Compliance: communications policy and compliance framework

Legal: legal-light posture, IP/brand usage rules

5.6 Developer (Tooling and Examples)

docs/dev/DEV_QUICKSTART_2.1.md — local developer workflows

docs/dev/JSON_RPC_EXAMPLES_2.1.md — canonical request/response examples

docs/dev/TOOLING_HARDHAT_VIEM_2.1.md — supported tooling patterns

6. Release and Change Tracking
6.1 Canonical Change Log

docs/CHANGELOG_2.1.md is the canonical human-readable history for released changes.

All material changes affecting behavior, persistence, or client compatibility must be reflected there.

6.2 Release Discipline

docs/RELEASE_PROCESS_2.1.md defines tagging, release gating, and upgrade discipline.

docs/AUDIT_READINESS_2.1.md defines evidence requirements and reproducibility expectations.

7. Maintenance Rules

To keep the documentation set coherent:

Single source of truth: Do not duplicate normative definitions across files.

No silent drift: If implementation changes, update the spec and changelog in the same release unit.

Stability labeling: RPC methods must be labeled Stable/Beta/Experimental per the stability policy.

Cross-document alignment: Architecture, state model, and RPC spec must not contradict each other.

Controlled scope: Anything out of scope must be explicitly declared (especially for audit readiness).

Operational realism: Runbooks must reflect validated, reproducible steps.

8. Quick “Start Here” Summary

Start here for the system: docs/ARCHITECTURE_2.1.md

Start here for client integration: docs/RPC_SPEC_2.1.md

Start here for state semantics: docs/STATE_MODEL_2.1.md

Start here for stability expectations: docs/API_STABILITY_POLICY_2.1.md

Start here for audit planning: docs/AUDIT_READINESS_2.1.md

Start here for release history: docs/CHANGELOG_2.1.md