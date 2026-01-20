NOORCHAIN 2.1 — Audit Readiness

Document ID: AUDIT_READINESS_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: NOORCHAIN 2.1 node implementation, operational runbooks, and protocol-facing behavior in controlled deployments.

1. Purpose

This document defines the audit-readiness posture of NOORCHAIN 2.1. It specifies:

What “audit readiness” means for this system

What evidence must exist and where it lives

How changes are controlled and traceable

What can be audited today versus what is explicitly out of scope

Minimum gates and reproducibility requirements for an audit engagement

The goal is not to claim “audited” status, but to ensure the project is structured so that an audit can be executed with clarity, controlled scope, and verifiable evidence.

2. Scope and System Boundaries
2.1 In Scope

An audit engagement for NOORCHAIN 2.1 may cover:

Node behavior (consensus, networking, leader/follower mode)

State model and persistence assumptions (world-state, receipts, block metadata)

RPC interface and compatibility promises (JSON-RPC methods, error model)

Operational runbooks (deployment model, monitoring, incident response basics)

Genesis and parameters (alloc policy, genesis checklist/spec, parameter reference)

Governance process for releases and upgrades (permissioned operational control)

2.2 Potentially In Scope (By Explicit Agreement)

PoSS (Proof of Signal Social) application layer logic, registry, and snapshot mechanics

Client-side scripts and tooling (Hardhat/viem workflows)

Internal interfaces between subsystems (execution hooks, database layout)

2.3 Out of Scope by Default

Unless explicitly added to scope, the following are out of scope:

Formal economic claims or market behavior

Custody, exchange integration, or third-party operational systems

Non-core frontends and marketing properties

Any environment outside controlled deployments (e.g., public permissionless networks)

Unspecified methods or behaviors not documented in RPC_SPEC_2.1.md

3. Audit Readiness Principles

NOORCHAIN 2.1 uses the following principles as audit-readiness invariants:

Deterministic traceability: Any shipped behavior can be tied to a commit and a tagged release.

Spec-to-implementation alignment: Documented RPC and state semantics must match implementation.

Reproducibility: An auditor can rebuild, run, and verify the same outputs using runbooks.

Change control: Material changes must go through a disciplined release process and changelog.

Least surprise: Client-facing behavior should be stable, well-scoped, and explicit about limits.

Evidence-first operations: Validation outputs and gates are recorded and reproducible.

4. Evidence Package Requirements

An audit-ready system must provide an evidence package that is sufficient to:

Reconstruct the build and runtime environment

Identify exactly what code was audited

Validate the major functional claims within scope

4.1 Mandatory Artifacts

The following artifacts must exist and be kept current:

docs/ARCHITECTURE_2.1.md (high-level architecture and boundaries)

docs/STATE_MODEL_2.1.md (world-state, receipts, roots, persistence model)

docs/RPC_SPEC_2.1.md (JSON-RPC method semantics)

docs/rpc/RPC_ERROR_MODEL_2.1.md (error taxonomy and semantics)

docs/rpc/RPC_COMPAT_MATRIX_2.1.md (compatibility gates and client coverage)

docs/RELEASE_PROCESS_2.1.md (tagging, releases, version discipline)

docs/CHANGELOG_2.1.md (human-readable change history with breaking change markers)

docs/ops/DEPLOYMENT_MODEL_2.1.md (deployment topologies and assumptions)

Genesis documentation under docs/genesis/* as locked by the docs index

4.2 Mandatory Code Traceability

A release under audit must correspond to:

A Git tag following the project’s release discipline

A commit hash accessible in the canonical repository

The repository must include:

A clean build path (documented command sequence)

Minimal reliance on unstated environment conditions

4.3 Mandatory Operational Evidence

For each release candidate, there must exist:

A reproducible runbook to start a multi-node configuration

Validation gates demonstrating:

Leader/follower connectivity

RPC availability

State persistence across restart

Receipt availability for submitted transactions

Block metadata visibility (roots/bloom) where implemented

Runbook-driven evidence should be captured as command transcripts and stored in a controlled location (internal archive), not as informal screenshots.

5. Audit Targets and Questions

An audit engagement typically decomposes into concrete questions.

5.1 Consensus and Networking

Does the node maintain correct peer connectivity and avoid inconsistent roles?

Are leader/follower semantics explicit and safe?

Are there observable risks from misconfiguration (ports, data-dir reuse, locks)?

5.2 Persistence and Data Integrity

What is persisted, where, and under what keys/layout?

How is the world-state root managed per block?

Are receipts and block metadata consistent across restarts?

Are there integrity checks or invariants that prevent silent corruption?

5.3 Execution Semantics (EVM and Contracts)

What execution environment is actually used in the audited release?

Are nonce/balance/code/storage updates consistent with documented state semantics?

Are receipts/logs generated and persisted in a manner compatible with documented RPC behavior?

5.4 JSON-RPC Interface

Do Stable methods behave exactly as specified?

Are error responses consistent with the error model?

Are ambiguous behaviors explicitly classified (Stable/Beta/Experimental)?

5.5 Genesis / Parameters / Alloc

How is alloc applied and guarded (one-time application markers)?

Can alloc be re-applied only under explicit operator action?

Are chain identifiers and parameter values consistent with documentation?

5.6 Operations and Release Discipline

Is there a clear release process with changelog discipline?

Are upgrades controlled and documented (especially in permissioned mode)?

Is incident response documented and realistic?

6. Audit Modes
6.1 Design Review (Documentation-Led)

A design review audits:

Architecture and threat model assumptions

Documented boundaries and invariants

Spec completeness and internal consistency

Outputs: findings on specification clarity, missing invariants, and risk posture.

6.2 Implementation Review (Code-Led)

A code-led audit focuses on:

Correctness of the implemented semantics

Persistence integrity and state transitions

RPC behavior and error handling

Outputs: issues with severity ranking, recommended fixes, and regression tests.

6.3 Operational Review (Runbook-Led)

An operational audit covers:

Deployment model validity

Monitoring/health endpoints

Failure mode handling and recovery

Configuration hazards and operator safety

Outputs: runbook improvements and operational risk mitigation.

7. Reproducibility Requirements

An auditor must be able to reproduce key behaviors with deterministic steps.

7.1 Build Reproducibility

Minimum requirement:

Build must succeed using documented commands from repository root.

The produced binary must run without requiring hidden local state.

Recommended:

Documented Go version requirement and a verification command.

A stable build output name and location.

7.2 Runtime Reproducibility

Minimum requirement:

Nodes can be started with explicit flags and explicit data directories.

Multi-node connectivity can be validated using documented gates.

RPC responses for core methods can be validated via documented example calls.

7.3 Evidence Reproducibility

Minimum requirement:

Validation transcripts include:

the exact commands executed

the exact output (or hashes where appropriate)

timestamps and node role context (leader/follower)

8. Change Control and Auditability
8.1 Change Logging

All material changes affecting:

state semantics

RPC method behavior

persistence layout

consensus/network behavior

deployment model assumptions

must be recorded in:

docs/CHANGELOG_2.1.md (human-readable)

and tied to a tag under docs/RELEASE_PROCESS_2.1.md.

8.2 Backward Compatibility Discipline

Breaking changes to Stable methods must follow:

docs/API_STABILITY_POLICY_2.1.md

Deprecation windows and migration paths must be explicit.

8.3 “Spec Is Law” for Audited Scope

Within the scope of an audit:

If implementation differs from spec, it is treated as a defect.

If spec is ambiguous, it is treated as a documentation defect that must be resolved before claiming readiness.

9. Risk Register Alignment

Audit readiness requires a clear mapping from risks to controls.

The following documents must exist and be aligned:

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

docs/ops/INCIDENTS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

The audit engagement should validate that:

risks are enumerated

mitigations are concrete

residual risk is acknowledged

10. What “Ready for Audit” Means (Practical Gates)

NOORCHAIN 2.1 is considered “ready for an audit engagement” when:

The audited release is tagged and reproducible.

The RPC spec is complete for the methods claimed as Stable/Beta.

The state model is documented and matches implementation.

Multi-node runbooks exist and pass gates repeatedly.

Change control (changelog + release process) is consistent.

Threat/trust model documents exist and are not contradictory.

This does not imply “no findings.” It implies the engagement can proceed without structural blockers.

11. Engagement Preparation Checklist

Before starting an audit:

Select the exact tag to audit.

Freeze scope: methods, subsystems, deployment mode.

Provide:

build instructions

runbooks for leader/follower

RPC examples for validation

a short description of the expected invariants

Define deliverables:

issue format

severity rubric

remediation expectations

Define re-test window and re-audit criteria.

12. References

docs/ARCHITECTURE_2.1.md

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/ops/INCIDENTS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md