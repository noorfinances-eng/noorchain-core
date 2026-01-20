NOORCHAIN 2.1 — Security and Trust Model

Document ID: SECURITY_TRUST_MODEL_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Security boundaries and trust assumptions for NOORCHAIN 2.1 in controlled deployments (permissioned operational control).

1. Purpose

This document defines the security model and trust assumptions of NOORCHAIN 2.1.

It clarifies:

what the system protects

what it does not protect

which entities are trusted (and for what)

what “secure” means under controlled deployment assumptions

which controls are expected from operators

This is a baseline for:

docs/THREAT_MODEL_2.1.md

incident response procedures

audit readiness and release gating

2. System Context

NOORCHAIN 2.1 is a sovereign EVM-compatible L1 operated in a controlled environment with permissioned operational control. The system is designed to be “mainnet-like” in behavior while being deployed under explicit governance and operator discipline.

2.1 Roles

Operators: run nodes and manage deployments, configuration, and upgrades.

Maintainers: merge changes, create releases/tags, and publish documentation updates.

Clients: wallets, tooling, dApps consuming JSON-RPC.

Curators (PoSS layer): application-layer actors signing and submitting PoSS snapshots (not consensus).

Participants (PoSS layer): application-layer participants referenced by PoSS mechanics.

2.2 Leader/Follower Model

Deployments may use:

a leader node as the authoritative producer

one or more followers that synchronize and may route specific reads to the leader

Routing rules are part of the interface semantics and must be explicitly documented.

3. Security Objectives
3.1 Primary Objectives

Integrity of chain data: blocks, transactions, receipts, and state must be internally consistent and persist correctly.

Correctness of client interface: JSON-RPC results must match documented semantics.

Operational safety: configurations and procedures must prevent avoidable corruption (e.g., data-dir reuse, port collisions).

Controlled upgrade safety: upgrades are executed only via documented, authorized processes.

3.2 Secondary Objectives

Minimize attack surface (localhost-by-default RPC, explicit exposure controls).

Preserve clear separation between consensus/security and application-layer value logic.

4. Trust Assumptions

This system intentionally assumes some trust due to permissioned operational control.

4.1 Trusted Entities
Operators (Trusted for Availability and Configuration Integrity)

Operators are trusted to:

run nodes on controlled machines

keep RPC exposure controlled

manage data directories, backups, and restarts safely

apply upgrades only when authorized and validated

Maintainers (Trusted for Code Integrity)

Maintainers are trusted to:

enforce review discipline

create tags only for validated commits

keep docs/spec aligned with implementation

Deployment Environment (Trusted Baseline)

The environment is assumed to provide:

OS-level process isolation

filesystem integrity sufficient for LevelDB usage

basic network protections (firewall, least exposure)

4.2 Untrusted Inputs

The system must treat as untrusted:

arbitrary JSON-RPC request inputs

transaction payloads (calldata)

peer network traffic (even in controlled networks, inputs can be malformed)

any external client behavior

5. Security Boundaries
5.1 Code Boundary

The node binary and its direct dependencies define the trusted computing base for correctness.

Changes to core behavior require:

release gating

changelog entries

documentation alignment

5.2 RPC Boundary

JSON-RPC is a primary attack surface:

input validation must be strict

errors must not leak secrets

resource usage must be controlled (avoid unbounded work)

RPC is expected to bind to localhost by default.

5.3 Persistence Boundary

On-disk persistence is critical:

data directories are sensitive infrastructure assets

corruption or unauthorized modification can compromise correctness

backups and restore procedures must be controlled

5.4 PoSS Boundary

PoSS is explicitly application-layer and not consensus.

Trust model for PoSS:

PoSS can define value and governance semantics off-consensus

PoSS does not secure the chain

Any PoSS snapshot claims must be evaluated under PoSS rules, not under consensus assumptions

This separation is a security invariant.

6. Threat Classes (High-Level)

This trust model recognizes the following threat classes:

RPC abuse: malformed inputs, heavy queries, resource exhaustion

Persistence faults: data-dir mismanagement, corruption, partial writes, unsafe backups

Operator errors: port collisions, running two nodes on same DB, misconfigured follower routing

Network manipulation: malformed peer messages, connection churn, spoofed peers in poorly controlled networks

Supply chain risks: dependency changes, build environment drift, compromised repository access

Application-layer misuse: embedding personal data in on-chain payloads, PoSS misinterpretation as consensus

Detailed threats and mitigations are enumerated in docs/THREAT_MODEL_2.1.md.

7. Controls and Expectations
7.1 Release and Change Control

Required controls:

tagged releases (tags as identity)

documented release process

changelog discipline

spec-to-implementation alignment

References:

docs/RELEASE_PROCESS_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/CHANGELOG_2.1.md

7.2 Operational Controls

Operators must:

bind RPC locally by default

use distinct data directories for each node

stop nodes cleanly before backup/restore workflows

keep logs and evidence for incident resolution

validate gates after each change

Reference:

docs/OPERATIONS_PLAYBOOK_2.1.md

7.3 Access Controls

At minimum:

restrict RPC exposure

restrict filesystem access to data dirs

restrict who can deploy binaries and change configs

In controlled environments, these are typically handled by OS user separation and firewall rules.

7.4 Key Material Controls

Private keys and secrets must never be committed.

Key exposure is treated as compromise and requires rotation.

Reference:

docs/PRIVACY_DATA_POLICY_2.1.md

8. Security Properties (What Is and Is Not Guaranteed)
8.1 Guaranteed Under This Model (When Controls Are Followed)

Consistent behavior for documented RPC methods at their stability level

Persistence and restart consistency for documented state and receipt models

Follower parity for methods designated as parity-gated (when configured correctly)

Controlled upgrade behavior via release discipline

8.2 Not Guaranteed

Permissionless adversarial security in open networks (not the default model)

Privacy of on-chain payloads

Protection against malicious operators (operators are trusted within this model)

Economic safety (markets, pricing, token liquidity) — out of scope

9. Audit Readiness Alignment

This trust model is intended to be auditable:

assumptions are explicit

controls are documented

evidence expectations exist

Reference:

docs/AUDIT_READINESS_2.1.md

10. Change Management for Security Model

Changes to this security/trust model require:

version bump to this document

CHANGELOG_2.1.md entry

review against the threat model and incident response docs

11. References

docs/ARCHITECTURE_2.1.md

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

docs/THREAT_MODEL_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/ops/INCIDENTS_2.1.md