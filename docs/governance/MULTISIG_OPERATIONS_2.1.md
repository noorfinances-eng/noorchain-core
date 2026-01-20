Document ID: MULTISIG_OPERATIONS_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines operational standards for multi-signature custody and approvals used in NOORCHAIN 2.1 controlled deployments. This document addresses authorization controls, signing workflows, key handling requirements, and evidence expectations under a Legal-Light posture.

Purpose

This document defines how multi-signature (multi-sig) control is used to authorize NOORCHAIN 2.1 high-impact actions.

It is designed to:

reduce single-person key risk

formalize approval thresholds and sign-off rules

standardize operational signing workflows

preserve evidence for audit readiness

maintain Legal-Light posture (no custody claims to third-party assets)

This document applies to:

governance approvals (release/upgrade authorizations)

validator membership and key rotation authorizations (where multi-sig is used)

treasury-like administrative addresses used for controlled environments

policy approvals and incident escalations requiring multi-party control

Principles

2.1 Separation of Duties

Multi-sig signers must be distributed across roles to prevent unilateral control. A minimum separation-of-duties posture is:

at least one signer from governance (Foundation or delegate)

at least one signer from implementation (Dev Sàrl or release manager)

at least one signer from operations (operator lead or security reviewer)

2.2 Least Privilege

Only actions requiring multi-party authorization should be guarded by multi-sig. Avoid using multi-sig for routine actions that create operational bottlenecks.

2.3 Evidence and Reproducibility

Every multi-sig-approved action must be:

referenced to a specific change artifact (commit/tag, configuration hash, validator set file digest)

recorded with a decision record and signer confirmations

Reference:

docs/AUDIT_READINESS_2.1.md

docs/RELEASE_PROCESS_2.1.md

2.4 No Custody Claims

Multi-sig is used as an internal authorization control. It does not imply custody services or financial guarantees.

Reference:

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

Multisig Types and Use Cases

3.1 Governance Approval Multisig

Purpose:

approve Class 2 and Class 3 governance actions (releases, upgrades, genesis-critical changes)

Typical guarded actions:

release authorization (tag approval)

upgrade authorization (rollout approval)

validator set change authorization (membership changes)

emergency actions escalation approval (SEV-0 containment measures)

Reference:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/governance/UPGRADE_PROCESS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

3.2 Administrative Contract/Address Multisig (If Applicable)

Purpose:

control administrative addresses used in controlled deployments (e.g., system addresses, curated contract admin roles)

Constraints:

usage must remain compatible with PoSS separation and Legal-Light posture

admin controls should be minimized and documented explicitly

3.3 Validator Coordination Multisig (Optional)

Purpose:

coordinate validator set updates where a multi-party authorization is desired before distributing new validator configuration artifacts

Thresholds and Signer Composition

4.1 Threshold Policy

Threshold is environment-specific but must be recorded and stable.

Recommended baseline (illustrative):

3-of-5 for controlled environments

3-of-7 for permissioned pre-mainnet/mainnet (where more signers are available)

Rules:

thresholds must be recorded in the environment governance record

lowering a threshold is a high-impact action requiring Class 3 approval

replacing signers is a high-impact action requiring governance approval and evidence

4.2 Signer Requirements

Each signer must satisfy:

hardware-backed or otherwise hardened key storage (where feasible)

independent operational control (no shared devices, no shared seed phrases)

ability to respond within defined incident windows (for SEV-0)

Signer identities should be recorded internally in a controlled registry with:

role designation

public signing address

activation date

deactivation date (if removed)

Key Handling and Security Requirements

5.1 Key Generation

Rules:

keys must be generated on a trusted device

seed phrases must never be transmitted digitally in plaintext

backups must be encrypted and stored separately

5.2 Key Storage

Requirements:

prefer hardware wallets for human signers

if software keys are used in controlled CI environments, they must be isolated, encrypted at rest, and access-controlled

avoid storing keys in repositories or shared notebooks

5.3 Key Rotation

Rotation triggers:

suspected compromise

signer change (personnel)

periodic hygiene (policy-driven)

Rotation process must include:

decision record authorizing rotation

new signer onboarding verification

deactivation record for removed signer

Reference:

docs/SECURITY_TRUST_MODEL_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

Signing Workflow (Operational)

This section defines a standardized workflow for multi-sig approvals regardless of tooling.

6.1 Pre-Signing Requirements

Before collecting signatures, create a signing package containing:

action description (what is being approved)

scope and impact classification (Class 0–3)

target artifacts:

git commit hash and/or tag

configuration file digest(s)

genesis digest (if applicable)

validator set digest (if applicable)

rollback/recovery plan (for Class 2/3)

validation gates that must pass after execution

Evidence to capture:

signing package hash (sha256)

distribution record to signers (internal)

References:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

6.2 Signing Session Rules

Rules:

signers must verify the signing package digest matches what they received

signatures must be collected through the standard toolchain used for that environment

signing must not occur under time pressure without an IC decision (SEV-0 exception only)

For SEV-0:

IC may request an expedited signing session

the signing package must still exist, but may be minimal (bounded by containment objective)

6.3 Post-Signing Finalization

After threshold is met:

record the final signed payload (or transaction id)

record the signer set that approved

record execution timestamp (UTC)

execute the approved change per runbook

run validation gates and attach outputs to the evidence pack

Evidence to capture:

final signed payload / tx hash

list of signer addresses

validation outputs after execution

Common Approved Action Types (Mapping)

7.1 Release Authorization (Class 2)

Signing package must include:

git tag reference

changelog entry reference

validation evidence (smoke tests and parity gates)

References:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

7.2 Upgrade Authorization (Class 2)

Signing package must include:

upgrade plan (rolling restart vs coordinated)

validator coordination plan

rollback plan

post-upgrade gates

Reference:

docs/governance/UPGRADE_PROCESS_2.1.md

7.3 Validator Set Change (Class 2/3 depending on scope)

Signing package must include:

reason and risk analysis

updated validator set artifact digest

effective date/time

communication plan to validators

References:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

7.4 Genesis Re-Issuance (Class 3)

Signing package must include:

genesis intent record

genesis artifact digest

full genesis evidence pack checklist mapping

References:

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

Evidence Pack Requirements (Multisig)

A multi-sig evidence pack MUST contain:

signing package (or a reference with hash)

signer list and threshold

signed payload / tx hash

execution record (commands or procedural steps)

validation gate outputs

incident linkage (if related to incident response)

Evidence must be retained per audit readiness policy.

Reference:

docs/AUDIT_READINESS_2.1.md

Failure Modes and Mitigations

9.1 Signer Unavailability

Mitigations:

maintain a documented standby signer policy

define incident-time quorum procedures

rotate signers if repeated unavailability is observed

9.2 Key Compromise Suspected

Actions:

treat as SEV-0 or SEV-1 depending on scope

freeze approvals involving the compromised key

rotate signer set and thresholds as required

record and distribute the new signer registry

Reference:

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

9.3 Tooling Mismatch / Confusion

Mitigations:

enforce one standard signing toolchain per environment

require signers to verify payload digest and artifact references

provide a signer quick verification checklist (internal)

Change Control

Changes to multi-sig thresholds, signer composition policy, or signing workflows are governance-impacting.

Rules:

changes MUST be documented and versioned

changes MUST be referenced in changelog if they affect operational behavior

lowering thresholds requires explicit governance approval

References:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

References

docs/governance/GOVERNANCE_MODEL_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/governance/UPGRADE_PROCESS_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/THREAT_MODEL_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/legal/LEGAL_LIGHT_POSTURE_2.1.md
docs/compliance/COMMUNICATIONS_POLICY_2.1.md