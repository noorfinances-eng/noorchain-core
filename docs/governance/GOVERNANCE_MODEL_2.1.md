Document ID: GOVERNANCE_MODEL_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines the governance model for NOORCHAIN 2.1 controlled deployments, including roles, decision rights, change control, and operational safeguards. This document is governance-oriented and does not constitute an investment, custody, or yield promise.

Purpose

This document defines the governance structure used to operate NOORCHAIN 2.1 in controlled environments and through permissioned network phases.

It is designed to:

define roles and responsibilities unambiguously

formalize decision rights for technical and operational changes

enforce separation of concerns (consensus security vs application-layer governance)

support audit readiness through documented controls and evidence

maintain a Legal-Light posture (no custody, no yield promises)

This model applies to:

local controlled deployments

pilot environments

permissioned pre-mainnet and permissioned mainnet phases

Governance Principles

2.1 Security First

All governance actions must preserve:

chain identity invariants

network safety and availability

reproducibility of releases and configurations

controlled exposure posture for infrastructure

2.2 Separation of Concerns

NOORCHAIN 2.1 governance is explicitly layered:

Consensus governance governs the permissioned validator set, network operations, and protocol upgrades.

PoSS governance governs application-layer legitimacy processes (curation, signal aggregation, snapshot policies).

PoSS does not govern consensus. Consensus does not outsource its security decisions to PoSS.

Reference:

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

2.3 Minimalism and Evidence

Every governance action that changes system behavior must leave evidence sufficient for independent review.

Reference:

docs/AUDIT_READINESS_2.1.md

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

2.4 Legal-Light Posture (CH)

Governance language and artifacts must avoid:

yield or return guarantees

custody assumptions

ambiguous promises about token value

Operationally, governance must remain compatible with a controlled, permissioned deployment posture.

Reference:

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

Governance Domains

3.1 Domain A — Protocol and Network Operations

Includes:

node operations and deployment model

validator membership and key rotation

RPC exposure policy

incident response, recovery, and operational constraints

References:

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

3.2 Domain B — Release and Change Control

Includes:

release tagging and artifact integrity

upgrade procedures

change approval thresholds by impact

References:

docs/RELEASE_PROCESS_2.1.md

docs/governance/UPGRADE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

3.3 Domain C — Application-Layer PoSS Governance

Includes:

curator roles and eligibility

snapshot policy and review standards

evidence requirements for submissions

reward split invariants (participants vs curators)

PoSS rules are social/operational and may evolve, but must remain separated from consensus.

References:

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

Roles and Responsibilities

This section defines the canonical roles. A given person or entity may hold multiple roles only if separation-of-duties controls remain satisfied.

4.1 Foundation Governing Body (Strategic Steward)

Responsibilities:

stewards the long-term direction and governance posture

approves high-impact protocol identity changes (normally prohibited within 2.1)

appoints or approves the top-level operational signers (where required)

Decision rights:

final approval for network phase transitions (pre-mainnet → mainnet permissioned)

approval of governance policy documents and changes that affect compliance posture

Evidence to capture:

signed resolution (or board minutes reference) for major approvals

4.2 Dev Sàrl (Implementation Owner)

Responsibilities:

implements protocol changes and releases

maintains the reference repository and build reproducibility

maintains runbooks and operational documentation

Decision rights:

proposes changes, provides technical assessment and risk analysis

executes releases only after approvals per the change-control table

Evidence to capture:

release tag, commit hash, build evidence, changelog entry

4.3 Operators (Infrastructure Maintainers)

Responsibilities:

run nodes in controlled environments

enforce exposure posture (RPC localhost/private)

maintain logs and evidence packs

execute runbooks and respond to incidents

Decision rights:

may perform emergency operational actions within the incident response policy

may not perform protocol upgrades un_toggle without authorized release artifacts

Evidence to capture:

ops logs, runbook outputs, incident tickets/records

4.4 Validators (Permissioned Consensus Participants)

Responsibilities:

provide consensus signatures and availability

follow upgrade schedules and operational constraints

maintain validator keys securely

Decision rights:

may accept or reject upgrades within the permissioned governance process, as defined by threshold rules

may vote to suspend a validator (subject to process)

Evidence to capture:

validator set change records

key rotation records (public info only)

4.5 Curators (PoSS Application-Layer Roles)

Responsibilities:

verify and validate PoSS submissions according to defined policy

sign snapshot summaries (application layer)

maintain evidence packs for validation decisions

Decision rights:

accept/reject items within PoSS workflows

propose policy improvements within PoSS domain

Constraints:

curators do not control consensus decisions

curator signatures do not modify validator membership

Evidence to capture:

signed snapshots

evidence pack export (hashes, timestamps, validated items)

Decision Classes and Thresholds

This section defines governance decision classes with approval thresholds. Thresholds may be implemented via multi-sig or documented approvals.

5.1 Class 0 — Routine Operations (Low Risk)

Examples:

rotating logs

restarting nodes

patching operational scripts that do not change protocol behavior

updating runbooks without changing operational guarantees

Approval:

Operator Lead approval (or documented operator decision)

Evidence:

runbook execution record or ops log snippet

5.2 Class 1 — Operational Policy Changes (Moderate Risk)

Examples:

changing RPC exposure posture (localhost → private network)

changing monitoring/health endpoints policy

changes to backup/restore procedures

Approval:

Operators + Dev Sàrl sign-off

Evidence:

policy diff + reason + validation outputs

References:

docs/PRIVACY_DATA_POLICY_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

5.3 Class 2 — Software Release (High Risk)

Examples:

any code change that affects:

consensus behavior

block structure / roots

RPC semantics

world-state persistence

transaction execution

Approval:

Dev Sàrl release manager + Foundation governance approval (or delegated approval)

Validator coordination required before rollout

Evidence:

release tag

changelog entry

reproduction evidence (build/validation runbook outputs)

References:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

docs/governance/UPGRADE_PROCESS_2.1.md

5.4 Class 3 — Protocol Identity / Genesis (Critical)

Examples:

changes to:

EVM chainId

fixed supply invariants

genesis re-issuance for an environment

validator set base assumptions (permissioned composition)

Approval:

Foundation governing body resolution + Dev Sàrl sign-off + validator set agreement

treated as a controlled re-anchoring event

Evidence:

signed resolution reference

genesis evidence pack per genesis checklist

References:

docs/genesis/GENESIS_CHECKLIST_2.1.md

docs/genesis/GENESIS_SPEC_2.1.md

Validator Set Governance (Permissioned)

6.1 Membership Changes

Validator additions/removals must follow:

documented reason (capacity, security, operational failure)

pre-defined effective time (or block height)

key material verification (public keys only, provenance documented)

Approval threshold:

defined by the permissioned governance charter (multi-sig / policy)

at minimum: Operators + Dev Sàrl + Foundation delegate

Evidence:

updated validator set artifact

rollout evidence (nodes updated, network stable)

6.2 Key Rotation

Rules:

rotation events must be recorded

old keys must be retired safely

new public identifiers must be distributed to operators and validators

Evidence:

rotation record + effective timestamp

confirmation from validators that new keys are active

PoSS Governance (Application Layer)

7.1 Invariants

PoSS must preserve:

strict separation from consensus

reward split model: 70% participants / 30% curators (policy invariant)

signature-based legitimacy (curators sign snapshots)

Changes to PoSS rules are allowed only if they:

do not change consensus security

do not introduce financial promises

are documented and versioned

7.2 Policy Changes

PoSS policy changes must include:

policy version bump

backward compatibility notes (if relevant)

operational rollout plan for curators and tools

Evidence:

updated policy doc

curator acknowledgment record (where applicable)

Documentation Governance

Documentation is a governance control surface.

Rules:

specs and policies must be versioned

changes must be tracked in CHANGELOG

release process must tie documentation versions to code tags

References:

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/NOORCHAIN_Index_2.1.md

Records and Evidence Retention

Minimum record set to retain per environment:

release artifacts (tags, commit hashes)

genesis evidence packs (where applicable)

incident records and postmortems

validator set change logs

PoSS snapshot evidence packs (if applicable)

Retention rules and privacy posture are defined in:

docs/PRIVACY_DATA_POLICY_2.1.md

docs/AUDIT_READINESS_2.1.md

Escalation and Dispute Handling

If conflicts arise between domains:

security and availability take precedence

if a change affects chain identity or invariants, treat it as Class 3

unresolved disputes escalate to Foundation governance resolution

For incidents:

follow incident response policy

prioritize containment, evidence preservation, and service restoration

Reference:

docs/governance/INCIDENT_RESPONSE_2.1.md

References

docs/governance/INCIDENT_RESPONSE_2.1.md
docs/governance/MULTISIG_OPERATIONS_2.1.md
docs/governance/UPGRADE_PROCESS_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/THREAT_MODEL_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md
docs/AUDIT_READINESS_2.1.md