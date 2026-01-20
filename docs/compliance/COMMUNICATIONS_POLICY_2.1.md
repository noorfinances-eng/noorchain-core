NOORCHAIN 2.1 — Communications Policy

Document ID: COMMUNICATIONS_POLICY_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Public and external communications about NOORCHAIN 2.1 (technical status, pilots, governance posture, legal-light constraints).

1. Purpose

This policy defines how NOORCHAIN 2.1 is communicated externally in a controlled and compliant manner.

It establishes:

what can be claimed and what must not be claimed

how to describe the system status (development, controlled deployments)

how to communicate about PoSS and token concepts without misleading language

approval and review discipline for public statements

rules for handling sensitive operational details

This policy is designed to support a “legal-light” posture and to reduce reputational, regulatory, and security risk.

2. Scope

This policy applies to:

websites, landing pages, and documentation exposed publicly

social posts (X/Twitter, LinkedIn, etc.)

pitch decks and partnership materials

public talks and interviews

email communications to external parties about NOORCHAIN 2.1

press releases and media interactions

It also applies to internal communications when they are likely to become external (forwarded emails, shared docs).

3. Core Principles

Accuracy over persuasion
Describe what exists and what is validated, not aspirations presented as facts.

No investment framing
Avoid language that constitutes a solicitation or promises financial outcomes.

Controlled disclosure
Do not publish operationally sensitive details that increase attack surface.

Separation of layers
PoSS is application-layer and must not be described as consensus security.

Version and scope discipline
Claims must be tied to specific versions/tags when they describe technical properties.

4. Allowed Claims vs. Prohibited Claims
4.1 Allowed Claims (With Conditions)

The following categories are allowed when true and when phrased conservatively:

Development status: “under active development,” “controlled deployments,” “mainnet-like testing environment.”

Technical scope: “sovereign EVM-compatible L1,” “JSON-RPC interface,” “leader/follower deployment model.”

Validation statements: “validated in controlled environment,” “runbook-driven gates executed,” tied to a tag where relevant.

PoSS description: “application-layer mechanism,” “off-consensus snapshots,” “curator-signed evidence workflow,” without claiming consensus security.

4.2 Prohibited Claims (Always)

The following are prohibited:

Any guarantee of profit, price appreciation, yield, or return.

“Risk-free” or “guaranteed” language of any kind.

Statements implying regulatory approval or endorsement.

Statements that PoSS secures the chain or replaces consensus.

Claims of “audited” unless a published audit exists and is accurately referenced.

Misleading comparisons that imply equivalence to Ethereum mainnet security without qualification.

5. Technical Status Communication Rules
5.1 Use Explicit Status Labels

When describing the system, use explicit labels:

Prototype

Development build

Validated milestone

Controlled deployment

Release tag (if applicable)

Avoid vague terms like “production-ready” unless formally defined and proven.

5.2 Tie Claims to Evidence

For any technical claim that matters (RPC coverage, state persistence, multi-node behavior):

reference the relevant documentation

reference the validated release tag when possible

avoid claiming completeness beyond what has gates

5.3 Do Not Overspecify Unknowns

If a capability is planned but not validated, state:

“planned,” “in progress,” or “not yet implemented”

and do not imply it exists today

6. PoSS Communications Rules
6.1 Required Language (PoSS)

Any external mention of PoSS must include:

PoSS is application-layer, not consensus

PoSS does not provide chain security

PoSS uses curator-signed snapshots as a governance/value mechanism

6.2 Prohibited Language (PoSS)

Do not claim:

PoSS “secures” the chain

PoSS is “proof-of-stake consensus”

PoSS “guarantees rewards” or “pays yield”

If a reward model is discussed at all, it must be described as a design rule subject to governance and must not be framed as a promise.

7. Token / Economic Communications Constraints

NOORCHAIN communications must avoid:

“investment opportunity” framing

implied listing commitments

price predictions

return projections

If token concepts are mentioned:

describe them as protocol design elements

avoid inducement language

include a “not an investment offer” disclaimer where appropriate

8. Disclosure of Operational Details
8.1 Prohibited Operational Disclosures (Public)

Do not publish:

server IP addresses and open port details

exact node topology and peer lists

internal admin endpoints or monitoring URLs

private operational runbook steps that increase exploitability

Public documentation should be ops-safe: it can explain concepts and safe examples but should not provide attack-enabling specifics.

8.2 Allowed Operational Disclosures (Controlled)

In private operator contexts (trusted partners, auditors under NDA):

details may be shared as needed

evidence packs may include port bindings and logs

secrets must still never be shared

9. Approval and Review Process
9.1 Content Classes

Class A (High Risk): token/economics, fundraising, partnerships, claims about audit/security, public roadmap commitments
Requires explicit review and approval.

Class B (Medium Risk): technical updates, milestone announcements, pilot updates, documentation releases
Requires review for accuracy and scope.

Class C (Low Risk): neutral educational content, general philosophy, non-specific design narratives
Requires minimal review; still must follow the prohibited claims list.

9.2 Review Checklist

Before publishing:

Is every claim true today?

Is technical scope tied to a version/tag where relevant?

Does it avoid investment framing?

Does it preserve PoSS separation?

Does it avoid operationally sensitive details?

Does it align with legal-light posture?

10. Disclaimers (Standard Language)

Where appropriate (technical posts, decks, pilots), include a short disclaimer:

NOORCHAIN 2.1 is under active development and operated in controlled environments.

This is not an investment offer.

No guarantees of performance, returns, or future availability are made.

Disclaimers must be short and factual.

11. Incident and Correction Policy

If incorrect information is published:

correct it promptly

issue a factual clarification without defensiveness

avoid expanding the scope of claims during correction

For security-relevant disclosures, coordinate with incident response governance:

docs/governance/INCIDENT_RESPONSE_2.1.md

12. Alignment with Other Policies

This policy must be consistent with:

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

docs/compliance/COMPLIANCE_FRAMEWORK_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

13. Change Control

Any change to communications rules requires:

version bump to this document

changelog entry

review against legal-light posture and compliance framework