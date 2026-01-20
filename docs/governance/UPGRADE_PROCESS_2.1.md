Document ID: UPGRADE_PROCESS_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines the controlled upgrade process for NOORCHAIN 2.1 permissioned deployments, including approval gates, multi-node rollout rules, validation requirements, and rollback constraints. This document is operational and assumes a permissioned BFT environment.

Purpose

This document defines how software upgrades are prepared, approved, deployed, and validated for NOORCHAIN 2.1.

It is designed to:

prevent configuration drift and uncontrolled forks

preserve chain identity and determinism

standardize upgrade approvals and evidence

define safe rollout patterns for permissioned networks

provide rollback rules that do not violate identity constraints

Scope and Assumptions

2.1 Applies To

controlled local environments

pilot environments

permissioned pre-mainnet and permissioned mainnet phases

2.2 Assumptions

operator access is controlled

validator membership is permissioned

releases are tagged and reproducible

upgrade windows can be coordinated

References:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/governance/MULTISIG_OPERATIONS_2.1.md

docs/RELEASE_PROCESS_2.1.md

Upgrade Classes and Impact

Upgrades are classified by impact. The governance decision class must be recorded.

3.1 Class A — Operational Upgrade (No Protocol Behavior Change)

Examples:

runbook improvements

monitoring/health wiring changes that do not affect chain behavior

deployment script improvements (no protocol impact)

Approval:

Ops Lead + Dev Sàrl sign-off

Evidence:

commit hash

validation gates relevant to ops endpoints

3.2 Class B — RPC/Tooling Behavior Change (Protocol-Adjacent)

Examples:

JSON-RPC method behavior changes

follower routing changes

additional RPC methods that affect wallet compatibility expectations

Approval:

Dev Sàrl release manager + Ops Lead

governance approval if impact is broad

Evidence:

release tag

RPC compatibility gates per matrix

changelog entry

References:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/RPC_SPEC_2.1.md

docs/CHANGELOG_2.1.md

3.3 Class C — Protocol Behavior Change (High Risk)

Examples:

transaction execution behavior

world-state model changes

block header roots/bloom behavior changes

consensus-related changes

Approval:

governance Class 2 approval (multi-party)

validator coordination required

Evidence:

release tag

parity gates and state persistence proof

rollback plan

References:

docs/STATE_MODEL_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

3.4 Class D — Identity/Genesis Re-Anchor (Critical)

Examples:

genesis re-issuance for an environment

chain identity change

EVM chainId change (normally prohibited in 2.1)

Approval:

governance Class 3 approval

special process and separate evidence pack

References:

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

Pre-Upgrade Requirements

4.1 Release Preparation

Every upgrade MUST reference a reproducible release artifact:

git commit hash and/or annotated tag

changelog entry describing the change and impact

compatibility notes (if RPC behavior changes)

Reference:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

4.2 Approval Package

Before rollout, an approval package MUST be assembled containing:

upgrade class (A–D)

scope (which environments, which nodes)

risk analysis (security and availability)

validation gates to be executed post-upgrade

rollback plan and constraints

operator window (UTC) and coordination plan

For Class C/D, the approval package MUST be approved under multi-sig policy.

Reference:

docs/governance/MULTISIG_OPERATIONS_2.1.md

4.3 Operator and Validator Coordination

For multi-node permissioned deployments:

define the upgrade window

confirm validator readiness

define the expected network behavior during rollout (degraded vs paused)

Evidence to capture:

coordination record (internal message or signed note)

window start/end timestamps (UTC)

Upgrade Rollout Patterns

5.1 Rolling Upgrade (Preferred When Safe)

Rolling upgrade replaces binaries node-by-node while attempting to maintain service continuity.

Rules:

only allowed if the upgrade is backward compatible at the network protocol level

nodes must not diverge in chain behavior

follower routing must remain coherent

Recommended order:

upgrade follower nodes first (if follower is read-only)

upgrade leader/primary nodes

upgrade remaining validators

Validation after each node upgrade is required (see Section 6).

5.2 Coordinated Restart (Preferred When Not Backward Compatible)

If the upgrade changes protocol behavior (Class C) or risks divergence:

stop nodes in a coordinated window

upgrade binaries and configs

restart in a defined order

re-run full parity and liveness gates

5.3 Emergency Upgrade (Incident-Driven)

Emergency upgrades occur under incident response control.

Rules:

IC must declare the incident severity (SEV-0/SEV-1)

changes must still reference a known-good tag or a minimal emergency patch with recorded commit hash

evidence pack must be preserved and later reviewed

Reference:

docs/governance/INCIDENT_RESPONSE_2.1.md

Post-Upgrade Validation Gates (Must Pass)

Validation gates are mandatory. The exact scope depends on upgrade class.

6.1 Minimum Gates (All Upgrades)

node process starts without error

RPC reachable (if enabled)

eth_chainId matches expected value

eth_blockNumber returns valid quantity

eth_getBlockByNumber("latest", false) returns coherent structure

References:

docs/RPC_SPEC_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

6.2 Multi-Node Gates (If Applicable)

no port collisions

P2P sessions established

leader/follower parity gates pass for:

eth_chainId

eth_blockNumber

eth_getBlockByNumber("latest", false)

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

6.3 World-State and Execution Gates (If Applicable)

If the upgrade affects state or execution:

eth_getBalance parity (leader/follower within defined routing rules)

eth_getTransactionCount parity

state persistence across restart (read-after-restart)

receipts availability via eth_getTransactionReceipt (if tx path relevant)

References:

docs/STATE_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

6.4 RPC Compatibility Gates (If Applicable)

If the upgrade affects RPC behavior:

run the compatibility matrix gates for required wallet/tooling methods

confirm error model expectations for unsupported methods remain stable

References:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

Rollback Policy

Rollback is permitted only if it does not violate chain identity constraints and does not create a fork.

7.1 Conditions for Rollback

Rollback is allowed when:

rollback target is a known-good release tag

rollback does not change genesis anchoring or chainId invariants

operators can reproduce the exact binary

7.2 Rollback Process Requirements

Rollback must include:

IC or governance approval (depending on severity/class)

explicit record of why rollback is safer than continuing

full post-rollback validation gates (Section 6)

References:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

Configuration Management During Upgrades

Rules:

configuration changes must be versioned and recorded

node start commands must be preserved in evidence packs

avoid “silent” changes to ports, exposure posture, or follower routing

References:

docs/genesis/PARAMETERS_REFERENCE_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

Evidence Pack Requirements (Upgrade)

An upgrade evidence pack MUST include:

upgrade approval package reference (hash)

release tag / commit hash

changelog entry reference

per-node:

start/stop timestamps (UTC)

exact command lines used

validation gate outputs

multi-node parity outputs (if applicable)

rollback decision record (if rollback occurred)

Reference:

docs/AUDIT_READINESS_2.1.md

Change Control

This document is part of governance controls.

Changes to upgrade thresholds, validation gates, or rollback policy are governance-impacting and must be reviewed and logged.

References:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

References

docs/governance/GOVERNANCE_MODEL_2.1.md
docs/governance/MULTISIG_OPERATIONS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/STATE_MODEL_2.1.md
docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/rpc/RPC_ERROR_MODEL_2.1.md
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/PARAMETERS_REFERENCE_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md
docs/AUDIT_READINESS_2.1.md