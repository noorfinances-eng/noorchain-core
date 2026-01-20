NOORCHAIN 2.1 — Release Process

Document ID: RELEASE_PROCESS_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: How NOORCHAIN 2.1 changes are packaged, validated, tagged, and promoted into “release-like” use (including controlled deployments).

1. Purpose

This document defines the release process for NOORCHAIN 2.1. It ensures that:

changes are traceable (commit → tag → release)

client-facing behavior (RPC) is stable and gated

operational procedures are reproducible

documentation and implementation do not drift

releases can be audited

This process applies to code and documentation that affect runtime behavior, state semantics, RPC surface, or operational posture.

2. Principles

Tags are the release identity.
A “release” is the Git tag and its referenced commit.

Evidence is required.
A release is not valid without passing defined gates and recording evidence.

Spec alignment is mandatory.
RPC and state semantics must match documentation in the same release unit.

Backwards compatibility is intentional.
Breaking changes are controlled by the API stability policy.

Ops-first.
A release must be runnable via documented procedures.

3. Release Units
3.1 Release Candidate (RC)

A Release Candidate is a commit intended to become a tag once gates pass.

Characteristics:

all changes are committed

documentation is updated

gates are executed and recorded

any known limitations are explicit

3.2 Tagged Release

A Tagged Release is a Git tag referencing an RC commit that passed gates.

Tags are immutable identifiers and must not be re-pointed.

3.3 Milestone Tags

The project may use milestone tags for validated milestones (e.g., mainnet-like pack gates). Milestone tags are treated as release-like for operational use and must follow the same gating discipline.

4. Branch Discipline (Operational)
4.1 Main Branch vs. Work Branches

Work occurs on feature branches or dedicated milestone branches.

Promotion occurs via controlled merges to the target branch (commonly main or a stable line).

4.2 Clean Working Tree Requirement

Before tagging:

working tree must be clean

no untracked artifacts that affect build/runtime behavior

data directories, caches, and logs must remain untracked

5. Required Release Artifacts

For a release to be accepted, the following artifacts must be updated as needed.

5.1 Always Required

docs/CHANGELOG_2.1.md — entry describing what changed

docs/NOORCHAIN_Index_2.1.md — if document set structure changes (rare)

relevant runbooks for any operational procedure changes

5.2 Required When RPC Changes

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md (if error handling changes)

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/API_STABILITY_POLICY_2.1.md (if stability classification or deprecation rules change)

5.3 Required When State/Persistence Changes

docs/STATE_MODEL_2.1.md

relevant runbooks demonstrating persistence and restart behavior

5.4 Required When Trust/Security Model Changes

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

6. Release Gates

A release must pass gates appropriate to its change set. Gates are pass/fail, not advisory.

6.1 Gate A — Build Gate

go build succeeds

binary is produced deterministically

no compilation warnings treated as errors in the environment

Evidence: build command + output.

6.2 Gate B — Single-Node Liveness Gate

node starts cleanly

RPC responds to:

eth_chainId

eth_blockNumber

health endpoint responds (if enabled)

Evidence: process listing, port binds, RPC outputs.

6.3 Gate C — Transaction Path Gate (When Applicable)

If transaction path is within scope:

eth_sendRawTransaction accepted

eth_getTransactionByHash returns the tx

eth_getTransactionReceipt returns a receipt after mining/inclusion

Evidence: tx hash, receipt output, block number change.

6.4 Gate D — Persistence Gate (When Applicable)

If persistence changes are within scope:

state survives restart:

relevant world-state reads

receipt availability

block metadata roots/bloom visibility (where implemented)

Evidence: pre-restart values + post-restart values.

6.5 Gate E — Multi-Node Gate (When Applicable)

If multi-node topology is within scope:

leader and follower run concurrently with distinct data-dirs

P2P connectivity is established and stable

follower parity gates pass for covered methods

Evidence: port binds, P2P established signals, parity call outputs.

6.6 Gate F — RPC Compatibility Gate

required methods behave as specified for their stability level

error model matches documentation

client/tooling gates in compatibility matrix are not regressed

Evidence: curated set of JSON-RPC example calls and outputs.

6.7 Gate G — Documentation Alignment Gate

relevant docs updated in the same release unit

no contradictions introduced across:

architecture

state model

RPC spec

stability policy

Evidence: reference commit includes doc changes; reviewers verify.

7. Evidence Recording Requirements

A release must include an evidence record sufficient to reproduce gates.

Minimum evidence:

exact commands executed

exact outputs captured (or hashed outputs where large)

node role context (leader/follower) and addresses used

tag/commit identifiers

Evidence may be stored internally as transcripts. The changelog must reference the tag and the high-level scope of validation.

8. Breaking Changes and Deprecations
8.1 Breaking Changes

Breaking changes require:

explicit [BREAKING] marker in CHANGELOG_2.1.md

migration notes

stability policy compliance

Stable RPC methods cannot be broken without deprecation discipline.

8.2 Deprecations

Deprecations require:

explicit entry in CHANGELOG_2.1.md

deprecation window as defined by docs/API_STABILITY_POLICY_2.1.md

replacement path and migration guidance

9. Tagging Rules
9.1 Tag Format

Tags should be short, explicit, and stable. The project may use:

milestone tags (e.g., M10-MAINNETLIKE-STABLE)

validated tags (e.g., M13-VALIDATED-FINAL)

release line tags if introduced later

The exact tag naming discipline should remain consistent within the 2.1 line.

9.2 Tag Immutability

Tags must not be moved.

If a tag is created incorrectly, a new tag must be created (and the incorrect one must be documented as invalid).

10. Promotion to Controlled Deployments

Promotion criteria depend on the environment:

Local controlled environment: Gates A–D must pass (and E if multi-node is used).

VPS controlled environment: Gates A–F must pass with operator evidence.

A release that changes RPC behavior or state semantics should not be promoted without updated docs and compatibility gates.

11. Rollback Rules

Rollback is allowed only to a prior tagged release.

Operator rules:

Stop nodes cleanly.

Restore prior binary and configuration.

If persistence formats changed, rollback may require data-dir restoration from backups. Treat persistence changes as high-risk and ensure backup runbooks exist.

Rollback events should be recorded in operational logs and incident documentation if they were triggered by faults.

12. References

docs/CHANGELOG_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/STATE_MODEL_2.1.md

docs/AUDIT_READINESS_2.1.md

Runbooks under docs/RUNBOOK-*.md