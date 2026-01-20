NOORCHAIN 2.1 — Legal-Light Posture (CH)

Document ID: LEGAL_LIGHT_POSTURE_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines the Legal-Light posture for NOORCHAIN 2.1 under a Swiss-aligned compliance approach for controlled and permissioned deployments. This document sets communication and operational constraints intended to reduce regulatory risk exposure. It is not legal advice.

Purpose

This document defines the Legal-Light posture for NOORCHAIN 2.1.

It is designed to:

provide a consistent compliance-aligned operational posture

constrain communications to avoid regulated financial promises

define guardrails for pilots and permissioned phases

clarify non-custodial assumptions and responsibilities

support audit readiness by making posture explicit and enforceable

This document is not legal advice. It defines internal policy and operational constraints.

Core Posture (Legal-Light)

NOORCHAIN 2.1 operates under a controlled and permissioned posture with the following baseline assumptions:

the network is operated in controlled environments (local, pilot, permissioned networks)

there is no public solicitation of investment returns

there is no custody service provided to the public

the native asset (NUR) is treated as a network utility/fee asset within the system context

communications remain factual, bounded, and evidence-based

Non-Negotiable Constraints

3.1 No Yield / No Return Promises

All communications and documents must avoid:

explicit yield, APY, ROI, “guaranteed returns”, “profit” framing

implicit financial promises through language suggesting inevitable appreciation

marketing statements that convert technical progress into investment claims

Allowed:

describing technical function (fees are paid in NUR)

describing allocation and supply in factual terms

describing controlled pilots and evidence-based outcomes (without valuation claims)

Reference:

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

3.2 No Custody Representation

NOORCHAIN 2.1 operations must avoid representing that:

the project holds assets on behalf of users

the project operates as a custodian, broker, or managed wallet provider

the project guarantees recovery of third-party keys or funds

If controlled accounts exist (e.g., admin multi-sig for operational governance), they are internal authorization controls, not public custody services.

Reference:

docs/governance/MULTISIG_OPERATIONS_2.1.md

3.3 Controlled Exposure

Infrastructure must be operated with conservative exposure posture:

JSON-RPC should be bound to localhost or a controlled private network

public exposure requires explicit governance approval and compensating controls

Reference:

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

3.4 Permissioned Network Phases

Until explicitly declared otherwise through governance and release processes:

validator membership is permissioned

network access and operational topology are controlled

public language must reflect “controlled / permissioned” posture

Reference:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/RELEASE_PROCESS_2.1.md

Communications Guardrails

4.1 Allowed Communication Categories

Communications may cover:

technical specifications and runbooks

controlled deployment evidence (liveness, parity gates, audit readiness)

governance processes and policy documents

pilot descriptions focused on process and legitimacy (not financial outcomes)

4.2 Prohibited Communication Categories

Communications must not include:

token price forecasts, “target prices”, “guaranteed liquidity”

claims of listing certainty or market performance

marketing narratives framed as investment opportunities

statements implying regulatory approval or registration without formal evidence

4.3 Required Disclaimers (When Applicable)

For documents or materials visible beyond internal operators:

include a bounded disclaimer that the system is under controlled deployment and is not an investment offer

reference the version and status of the document

Reference:

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

docs/README.md

Token Posture (NUR)

5.1 Functional Description

NUR is the native asset used for:

transaction fees (gas)

internal network economic operations (where applicable)

5.2 Supply and Allocation

Supply and allocation may be stated factually:

fixed supply: 299,792,458 NUR

allocation proportions are defined in governance documentation and genesis policy

Statements about supply and allocation must not be framed as investment value guarantees.

References:

docs/genesis/ALLOC_POLICY_2.1.md

docs/genesis/GENESIS_SPEC_2.1.md

5.3 No Public Investment Language

Avoid describing NUR as:

a “security”, “share”, “bond”, “guaranteed appreciating asset”

a product with promised returns or dividends

Pilot and Partner Engagement Posture

6.1 Controlled Pilot Framing

Pilots must be framed as:

controlled experiments validating process and evidence flows

legitimacy and audit-readiness demonstrations

limited scope operational deployments

Avoid:

presenting pilots as “public adoption” proof tied to token appreciation

implying partners endorse NUR as a financial instrument

Reference:

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

6.2 Data and Privacy

Pilot operations must comply with privacy and data policy:

minimize personal data collection

store only what is necessary for evidence

control access and retention

Reference:

docs/PRIVACY_DATA_POLICY_2.1.md

Operational Controls that Support Legal-Light

7.1 Evidence Packs and Audit Readiness

Maintaining audit-ready evidence reduces risk of ambiguous claims:

evidence packs for genesis, releases, upgrades, pilots

reproducibility of results under tagged releases

bounded claims tied to verifiable outputs (hashes, RPC outputs)

Reference:

docs/AUDIT_READINESS_2.1.md

docs/RELEASE_PROCESS_2.1.md

7.2 Change Control Discipline

Avoid untracked changes that could create uncontrolled behavior:

all changes go through release process

changelog updated for behavior changes

operational deviations recorded and reviewed

Reference:

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

7.3 Incident Response

If incidents occur that could impact trust posture:

treat communications carefully

preserve evidence

avoid speculation

Reference:

docs/governance/INCIDENT_RESPONSE_2.1.md

Risk Areas and Mitigations (Non-Exhaustive)

This section identifies typical risk areas and the internal posture to mitigate them.

8.1 Misleading Public Claims

Mitigation:

communications policy enforcement

approved channel control

versioned documents with disclaimers

8.2 Uncontrolled Public Exposure of RPC

Mitigation:

bind RPC to localhost/private only

firewall controls

monitoring and incident response

8.3 Key Handling and Custody Confusion

Mitigation:

multi-sig ops policy

no storage of third-party keys

clear internal role boundaries

References:

docs/governance/MULTISIG_OPERATIONS_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

Change Control

This posture is a governance control document.

Changes require:

version bump

governance review and approval if it affects external-facing posture

changelog entry if it materially affects policy constraints

References:

docs/governance/GOVERNANCE_MODEL_2.1.md

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

References

docs/compliance/COMMUNICATIONS_POLICY_2.1.md
docs/compliance/COMPLIANCE_FRAMEWORK_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md
docs/governance/GOVERNANCE_MODEL_2.1.md
docs/governance/MULTISIG_OPERATIONS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/THREAT_MODEL_2.1.md
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/ALLOC_POLICY_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md