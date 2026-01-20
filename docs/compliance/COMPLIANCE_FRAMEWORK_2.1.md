NOORCHAIN 2.1 — Compliance Framework

Document ID: COMPLIANCE_FRAMEWORK_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Baseline compliance posture for NOORCHAIN 2.1 in controlled deployments under a “legal-light” approach (Swiss context), focusing on operational and communications discipline rather than regulated activity claims.

1. Purpose

This document defines a practical compliance framework for NOORCHAIN 2.1 that:

supports controlled deployments and pilots

reduces regulatory and reputational risk

establishes minimum internal discipline for communications, data handling, and governance

clarifies what NOORCHAIN 2.1 is not doing (by default)

This framework is operational. It is not legal advice.

2. Scope and Boundaries
2.1 In Scope

Controlled deployments (local, VPS private, permissioned operator control)

Public communications and partnership materials

Data handling at node/operator level

Governance processes (upgrades, incident response, multisig operations)

Evidence and audit readiness posture

2.2 Out of Scope

Custody services

Exchange operations and broker-dealer activity

KYC/AML programs for third-party services

Promises of returns or investment solicitation

Consumer-facing “financial product” positioning

If NOORCHAIN expands into these areas, separate compliance work is required.

3. Legal-Light Posture (Baseline)

NOORCHAIN 2.1 adopts a legal-light posture intended to keep the system within a conservative boundary:

No custody assumptions by the core protocol or operators

No yield promises or guaranteed reward language

No token sale marketing under this documentation set

Controlled pilots and deployments presented as experimentation and engineering

The canonical statement of this posture is:

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

4. Compliance Principles

Truthfulness and traceability
Claims about the system must be tied to evidence and versions/tags where relevant.

No investment framing
Communications must not resemble solicitation or promises.

Privacy and data minimization
Avoid personal data on-chain; minimize logs and evidence disclosures.

Security-first disclosure
Do not publish operational details that increase attack surface.

Governance discipline
Releases and upgrades must follow process; emergency actions must be recorded.

5. Policy Set (Required Documents)

This framework depends on the following policies and models:

Communications: docs/compliance/COMMUNICATIONS_POLICY_2.1.md

Privacy: docs/PRIVACY_DATA_POLICY_2.1.md

Security trust model: docs/SECURITY_TRUST_MODEL_2.1.md

Threat model: docs/THREAT_MODEL_2.1.md

Incident response: docs/governance/INCIDENT_RESPONSE_2.1.md

Release discipline: docs/RELEASE_PROCESS_2.1.md

Audit posture: docs/AUDIT_READINESS_2.1.md

6. Communications Compliance Controls
6.1 Prohibited Content Categories

External communications must not include:

financial return promises

price predictions

“guaranteed rewards” language

statements implying regulatory approval

statements implying audits that do not exist

6.2 Required Disclaimers (When Applicable)

For technical updates, pilots, or decks:

system is under development / controlled deployments

not an investment offer

no guarantees of returns or future availability

6.3 Approval Discipline

High-risk communications (token/economics, fundraising, audits/security claims) require explicit review under the communications policy.

7. Data and Privacy Compliance Controls
7.1 On-Chain Data Warning

Policy baseline:

on-chain payloads are not private

do not place personal data on-chain

use hashes or references where identity linkage is needed

7.2 Logs and Evidence Handling

Operators must:

avoid logging secrets

treat logs as sensitive infrastructure data

redact internal identifiers when sharing externally

Reference:

docs/PRIVACY_DATA_POLICY_2.1.md

8. Governance and Operational Compliance Controls
8.1 Controlled Deployments

NOORCHAIN 2.1 is operated under permissioned control. This implies:

explicit operator responsibility for exposure and access control

explicit upgrade discipline

clear incident response procedures

8.2 Release Discipline

Compliance requires traceability and evidence:

tags represent release identity

changelog must reflect material behavior changes

spec must match implementation for claimed behaviors

Reference:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

8.3 Multisig and Key Management

Where multisig is used for governance actions:

operations must be documented

keys must be protected and rotated if compromised

no private keys are stored in repositories

Reference:

docs/governance/MULTISIG_OPERATIONS_2.1.md

9. Incident Response and Disclosure
9.1 Incident Handling

Incidents must be handled under documented governance:

classify severity

capture evidence (minimized and redacted)

document root cause and corrective actions

Reference:

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/ops/INCIDENTS_2.1.md

9.2 External Disclosures

External disclosures must:

avoid technical overexposure

avoid attributing blame without evidence

preserve user/partner privacy

Communications must remain factual and aligned with the communications policy.

10. Partner and Pilot Compliance Guidance

When engaging partners (schools, NGOs, institutions):

avoid collecting personal data unless necessary and authorized

prefer off-chain storage for sensitive records

define access control for evidence packs

ensure participants are not misled about financial outcomes

Pilot-specific addenda should be documented separately and should not override the core policies without explicit governance approval.

11. Audit and Assurance Statements
11.1 “Audit” Terminology

Do not claim “audited” unless:

an external audit report exists

the report is accurately described

scope and date are clearly stated

11.2 Audit Readiness

It is acceptable to state “audit-ready” only if:

release tags and evidence exist

documentation is consistent and complete for the claimed scope

runbook-based reproducibility has been demonstrated

Reference:

docs/AUDIT_READINESS_2.1.md

12. Records and Evidence

Minimum records to maintain internally:

release tags and their evidence transcripts

incident logs and post-incident reviews

communications approvals for high-risk content

partner/pilot agreements and data handling notes (if applicable)

This framework prefers minimal record retention, but not at the cost of operational safety.

13. Change Control

Changes to this compliance framework require:

version bump

changelog entry

alignment review against:

communications policy

privacy policy

legal-light posture

14. References

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/governance/MULTISIG_OPERATIONS_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/AUDIT_READINESS_2.1.md

docs/CHANGELOG_2.1.md