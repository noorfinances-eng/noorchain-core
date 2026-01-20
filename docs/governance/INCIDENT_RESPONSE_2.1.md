Document ID: INCIDENT_RESPONSE_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines the incident response lifecycle, roles, severity classification, evidence handling, communications posture, and recovery rules for NOORCHAIN 2.1 controlled deployments (local, pilot, permissioned networks).

Purpose

This document defines the operational incident response process for NOORCHAIN 2.1.

It is designed to:

reduce time to containment and recovery

preserve evidence for audit readiness

minimize configuration drift during emergencies

formalize communications posture (Legal-Light)

provide clear severity classification and escalation paths

This document applies to:

local controlled deployments

pilot environments

permissioned pre-mainnet and permissioned mainnet phases

Principles

2.1 Safety and Availability First

During an incident, prioritize:

containment and preventing further harm

service restoration for controlled operators and validators

evidence preservation

root cause analysis and corrective actions

2.2 Deterministic Operations

Emergency changes must still be documented and reproducible. Avoid ad-hoc “hot fixes” that cannot be replayed from a tagged release.

Reference:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

2.3 Minimal Exposure

Do not expand RPC or infrastructure exposure during incidents unless strictly necessary, and only with compensating controls.

Reference:

docs/PRIVACY_DATA_POLICY_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

2.4 Legal-Light Communications

External communications must avoid:

promises of returns or yields

claims about token price/value outcomes

disclosure of operationally sensitive details (IPs, key material, internal topology)

Reference:

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

Roles and Responsibilities

3.1 Incident Commander (IC)

Responsibilities:

owns the incident lifecycle and decision-making

assigns roles (Ops Lead, Comms Lead, Scribe)

approves containment actions and recovery milestones

decides when to escalate severity class

3.2 Ops Lead

Responsibilities:

executes technical actions per runbooks

coordinates operator activity across nodes

preserves logs and evidence artifacts

ensures changes are reversible and recorded

3.3 Scribe (Recorder)

Responsibilities:

maintains the incident timeline

records decisions, commands executed, and outputs

collects evidence pack references and hashes

produces the post-incident report draft

3.4 Comms Lead

Responsibilities:

drafts internal and external incident communications

ensures compliance with communications policy

manages stakeholder updates cadence

3.5 Security Reviewer (If Assigned)

Responsibilities:

assesses security impact (keys, integrity, exposure)

recommends containment for suspected compromise

validates evidence preservation scope

Incident Severity Classification

Incidents are classified into four severity levels.

4.1 SEV-0 (Critical)

Definition:

network integrity or chain identity at risk

suspected key compromise affecting validator operations

uncontrolled forks in a permissioned environment

sustained inability to produce blocks in controlled environments (where block production is expected)

confirmed unauthorized access to infrastructure

Immediate actions:

trigger emergency bridge for validators and operators

freeze changes not directly required for containment

rotate credentials if compromise is suspected

preserve full evidence pack

Escalation:

Foundation governance delegate is notified immediately

validators notified immediately

4.2 SEV-1 (High)

Definition:

prolonged outage or major functional impairment

RPC parity gates failing across leader/follower

persistent world-state inconsistency

repeatable transaction execution failure affecting controlled usage

Immediate actions:

contain blast radius (limit endpoints, isolate nodes)

execute known-good rollback if applicable

preserve logs and state evidence

4.3 SEV-2 (Medium)

Definition:

partial degradation, intermittent failures

monitoring/health failures or partial RPC method failures

performance issues that threaten stability but not integrity

Immediate actions:

mitigate, monitor, document

schedule corrective action if not urgent

4.4 SEV-3 (Low)

Definition:

minor issues or documentation/runbook mismatches

non-blocking warnings, cosmetic operational defects

Immediate actions:

document and fix during normal maintenance

Incident Lifecycle

5.1 Detection

Sources:

operator observation

validator alerts

health endpoint alerts

parity gate failures (leader/follower mismatches)

log-based anomaly detection

Required record:

detection timestamp (UTC)

detector identity (role)

initial symptoms and scope (which nodes, which RPCs)

5.2 Triage

Triage goals:

classify severity

define immediate containment objective

identify impacted components:

consensus / block production

RPC service

world-state

networking (P2P)

storage persistence

tooling / deployment scripts

Triage outputs:

severity level

incident owner (IC)

initial hypothesis

initial action plan (containment steps)

5.3 Containment

Containment actions may include:

stopping affected nodes to prevent further divergence

disabling external exposure (RPC bind tightening, firewall rules)

isolating a follower from serving inconsistent reads

reverting to last known-good release tag (if rollback policy permits)

Rules:

containment actions must be recorded with exact commands and outputs

avoid destructive operations unless explicitly authorized (see recovery section)

5.4 Eradication and Recovery

Recovery goals:

restore intended service (as defined by the environment profile)

confirm network coherence (multi-node gates)

confirm identity and invariants (chainId, genesis digest if applicable)

confirm world-state reads and parity (as defined by compat matrix)

Recovery must be validated by:

RPC identity checks

parity gates

liveness checks

References:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/RPC_SPEC_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

5.5 Post-Incident Review

Outputs:

root cause analysis (RCA)

corrective actions (short-term and long-term)

documentation/runbook updates

changelog entry if behavior changed

evidence pack retention confirmation

References:

docs/AUDIT_READINESS_2.1.md

docs/CHANGELOG_2.1.md

Evidence Handling and Preservation

6.1 Evidence Pack Requirements

An incident evidence pack MUST include:

incident timeline (UTC)

impacted nodes list and roles

node start commands (as run)

relevant logs (time bounded)

RPC outputs for identity and failing endpoints

parity gate results (leader vs follower)

version evidence:

git commit hash

release tag reference if applicable

6.2 Evidence Integrity

Rules:

store evidence in append-only storage where possible

compute hashes for log bundles and key artifacts

avoid editing raw logs; if redaction is required, keep originals internally

6.3 Sensitive Data Controls

Never capture or store in incident artifacts:

private keys or seed phrases

secrets in environment variables

full infrastructure topology if it increases attack surface

If secrets might be exposed:

treat the incident as SEV-0 or SEV-1 depending on scope

rotate credentials immediately and record rotation as a security action

Reference:

docs/PRIVACY_DATA_POLICY_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

Communications

7.1 Internal Communications

Internal updates should include:

severity class

current status (containment, recovery, monitoring)

known impacts (what is broken, what is safe)

next update time

7.2 External Communications (If Applicable)

External communication is permitted only under the communications policy and must:

avoid operationally sensitive details

avoid financial promises

remain factual and bounded

Reference:

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

7.3 Update Cadence

Recommended cadence:

SEV-0: frequent updates (e.g., every 30–60 minutes)

SEV-1: hourly updates

SEV-2: daily or as-needed

SEV-3: included in routine maintenance notes

Recovery Policies

8.1 Rollback Policy

Rollback is permitted when:

a known-good release tag exists

rollback does not violate chain identity constraints

operators can reproduce the exact binaries and configs

Rollback must be recorded as:

release tag chosen

reason for rollback

validation gates passed after rollback

References:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

8.2 Destructive Actions Policy (Data Loss)

Actions such as deleting data directories, wiping databases, or reinitializing state are high-risk and require:

IC approval

explicit documentation of consequences

evidence capture prior to the destructive step

If destructive actions affect genesis anchoring or chain identity, treat as SEV-0.

Reference:

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

8.3 Key Compromise Response

If validator keys or operational credentials are suspected compromised:

isolate affected systems

rotate keys/credentials

review logs for unauthorized access

re-issue validator set changes per governance process

References:

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

docs/governance/UPGRADE_PROCESS_2.1.md

Standard Validation Gates During Incidents

During recovery, run the minimum validation set:

eth_chainId

eth_blockNumber

eth_getBlockByNumber("latest", false)

parity gate reads as defined in RPC_COMPAT_MATRIX

service health endpoint (if enabled)

For multi-node:

confirm P2P sessions established

confirm follower routing behavior (if used)

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

Post-Incident Report Template (Minimum)

A post-incident report SHOULD include:

Summary (what happened, impact)

Timeline (UTC, minute resolution)

Root Cause

Detection and Response

Recovery Steps and Gates

Corrective Actions (Owners + deadlines)

Evidence Pack Index (hashes, locations)

Communications Summary (internal/external)

Reference:

docs/AUDIT_READINESS_2.1.md

References

docs/governance/GOVERNANCE_MODEL_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/THREAT_MODEL_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/compliance/COMMUNICATIONS_POLICY_2.1.md
docs/legal/LEGAL_LIGHT_POSTURE_2.1.md