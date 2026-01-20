NOORCHAIN 2.1 — Documentation

Version: January 2026
Scope: This docs/ directory contains the canonical documentation set for NOORCHAIN 2.1.

1. How to Use This Documentation

This repository contains both implementation code and the documentation required to operate, review, and integrate NOORCHAIN 2.1.

The documentation is structured to support:

Operators running leader/follower deployments

Integrators using JSON-RPC (wallets, tooling, scripts)

Review of architecture, state semantics, and trust assumptions

Governance, compliance, and legal-light posture

Audit readiness and evidence-driven release discipline

If you are unsure where to start, use the index:

docs/NOORCHAIN_Index_2.1.md

2. Audience-Based Start Points
2.1 Operators

Start with:

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

Runbooks under docs/RUNBOOK-*.md

Operators should run validated gate checks after each lifecycle step.

2.2 RPC / Client Integrators

Start with:

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/dev/JSON_RPC_EXAMPLES_2.1.md

docs/dev/TOOLING_HARDHAT_VIEM_2.1.md

2.3 Architecture / State Review

Start with:

docs/ARCHITECTURE_2.1.md

docs/STATE_MODEL_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

2.4 Governance / Compliance

Start with:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/governance/MULTISIG_OPERATIONS_2.1.md

docs/governance/UPGRADE_PROCESS_2.1.md

docs/compliance/COMPLIANCE_FRAMEWORK_2.1.md

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

docs/legal/IP_AND_BRAND_2.1.md

3. Normative vs. Guidance
3.1 Normative Documents

Normative documents define required behavior or process. If a conflict exists, normative documents prevail.

Core normative set includes:

docs/ARCHITECTURE_2.1.md

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

3.2 Guidance Documents

Guidance documents are operational or developer workflows that must remain consistent with the normative layer:

Runbooks under docs/RUNBOOK-*.md

Developer docs under docs/dev/*

Operational docs under docs/ops/*

Audit and release readiness docs

4. Document Set (2.1)

The canonical list of documents for NOORCHAIN 2.1 is maintained in:

docs/NOORCHAIN_Index_2.1.md

This list defines the perimeter for the documentation set and prevents duplication.

5. Change Control

Documentation changes that affect:

state semantics

RPC behavior

stability expectations

operational procedures

must be treated as controlled changes:

update docs/CHANGELOG_2.1.md

follow docs/RELEASE_PROCESS_2.1.md

6. Writing and Maintenance Rules

Prefer explicit, testable statements over prose.

Avoid ambiguous claims (“should usually,” “in most cases”) unless the variance is documented.

Keep “spec” and “runbook” roles separate:

specs define semantics

runbooks define procedures

Keep terminology consistent across files (leader/follower, routing, state roots, receipts, etc.).

7. Repository Pointers

While this README indexes documentation only, the broader repository typically includes:

node implementation under the core directories

scripts for tooling and local workflows

runbooks for validated milestone procedures

Exact paths depend on repository layout and release tags.