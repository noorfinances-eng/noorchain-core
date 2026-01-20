NOORCHAIN 2.1 — Privacy and Data Policy

Document ID: PRIVACY_DATA_POLICY_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Data handling and privacy posture for NOORCHAIN 2.1 nodes, operators, and associated developer tooling in controlled environments.

1. Purpose

This policy defines how NOORCHAIN 2.1 treats privacy and data across:

node operation and telemetry

JSON-RPC access and logs

persistence layers (databases and on-disk state)

incident response and evidence collection

development and pilot deployments (controlled environments)

It focuses on practical operator rules and system properties, not marketing statements.

2. Scope and Non-Goals
2.1 In Scope

Node logs and operational artifacts (process args, ports, errors)

On-disk persistence (state DB, receipts, metadata, PoSS snapshot storage where applicable)

JSON-RPC request/response handling and logging posture

Operator evidence collection procedures

Data minimization in pilots (where operationally relevant)

2.2 Out of Scope

Third-party services (cloud providers, hosting, CI, analytics) unless explicitly integrated

End-user application privacy policies for external frontends (separate documents)

Legal advice and jurisdiction-specific compliance determinations

3. Principles

NOORCHAIN 2.1 follows the principles below.

3.1 Data Minimization

Collect and retain only what is required for:

correct system operation

security and incident investigation

reproducibility and audit readiness

3.2 Local-First and Operator-Controlled

Default deployments are intended to be controlled:

RPC binds to localhost by default

node logs are local files under operator control

no mandatory external telemetry

3.3 Transparency of What Is Stored

Persistence layouts and semantics must be documented:

state model and persistence: docs/STATE_MODEL_2.1.md

any PoSS persistence semantics (if in scope): relevant spec/runbooks

3.4 No Hidden Data Exfiltration

NOORCHAIN 2.1 does not require outbound telemetry for correctness.

If optional outbound reporting is introduced in the future, it must be:

explicitly opt-in

documented as such

covered by a change log entry and release gating

4. Data Categories
4.1 On-Chain Data (Public by Nature)

By design, blockchains produce data that can be observed by network participants. On-chain data may include:

transaction hashes, sender/recipient addresses

calldata (which may embed arbitrary user-provided content)

receipts/logs

block headers and metadata (roots, bloom, etc.)

Policy: Operators and integrators must assume that on-chain data is not private, even in controlled deployments.

4.2 Node Operational Data (Local)

Node operational data includes:

process arguments (flags)

P2P peer addresses

port bindings and PIDs

error logs, warnings, stack traces (if enabled)

This data is typically stored in local log files and is under operator control.

4.3 RPC Access Data

JSON-RPC access involves:

method names

request parameters (may include addresses, tx payloads)

response payloads (may include sensitive operational context)

By default, NOORCHAIN 2.1 does not require logging full RPC request bodies. If an operator enables verbose request logging, it becomes the operator’s responsibility to treat those logs as potentially sensitive.

4.4 PoSS Application-Layer Data (If Enabled)

PoSS may store snapshot-related metadata. Policy implications:

snapshot contents and metadata must be treated as potentially sensitive if they encode real-world identities or events

avoid embedding personal data in on-chain payloads

prefer references/hashes over raw personal information

5. Data Storage and Retention
5.1 Persistence on Disk

Nodes may persist:

world-state database content

receipts and transaction artifacts

block metadata (roots, bloom, etc.)

PoSS snapshots (where applicable)

These are stored under the node’s -data-dir.

Policy:

Operators must protect -data-dir as sensitive infrastructure data.

Backups must be controlled and access-limited.

Never share raw databases publicly.

5.2 Log Retention

Default guidance:

Keep logs long enough to diagnose incidents and validate gates.

Rotate logs to avoid uncontrolled growth.

Retain incident-relevant excerpts separately (with access control).

Retention windows depend on deployment posture. In pilot environments, shorter retention may be appropriate if it does not compromise incident response.

5.3 Development Artifacts

Development environments may produce:

build artifacts

cache directories

local script outputs

temporary logs

Policy:

Do not commit sensitive artifacts.

Keep .gitignore updated for local caches and data directories.

Never commit private keys or secrets.

6. Secrets and Key Material
6.1 Private Keys

Policy:

Private keys must never be committed to the repository.

Keys must not be logged.

Keys should be provided at runtime via:

environment variables, or

interactive shell input, or

an operator-controlled secret manager

If a key is suspected to be exposed, treat it as compromised and rotate immediately.

6.2 RPC Credentials (If Used)

If RPC endpoints are protected by credentials or reverse proxies:

credentials must not be stored in plaintext in the repository

credentials must not appear in command history when possible

access logs must be treated as sensitive

7. RPC Exposure and Access Control
7.1 Default Posture

Bind JSON-RPC to 127.0.0.1 by default.

Use SSH tunneling or private networks for remote access.

7.2 Public Exposure

Public exposure of RPC is discouraged unless the following are in place:

authentication and rate limiting (reverse proxy or gateway)

strict method allowlists where feasible

monitoring and alerting

documented incident response procedures

If a public endpoint is required, treat RPC usage data and access logs as sensitive.

8. Incident Handling and Evidence Collection

During incidents, operators may need to collect and share evidence.

8.1 Evidence Minimization

Share only what is necessary:

RPC outputs for a specific method call

log excerpts relevant to the fault window

process and port listings

Avoid sharing:

full databases

full verbose RPC request logs (unless essential)

any secrets or tokens

8.2 Redaction Rules

Before sharing evidence externally:

redact private keys and tokens

redact internal IPs if they are not necessary to diagnose

avoid disclosing user-identifying payloads if present

8.3 Chain Data vs. Human Data

Even if the chain is in a controlled environment, do not treat it as a safe place for personal data.

If the application layer needs to reference real identities, use:

off-chain records with access control

on-chain hashes/commitments

minimal identifiers that do not encode personal information

9. Privacy Considerations for Pilots

In pilot deployments involving institutions or communities:

avoid collecting unnecessary personal data

store sensitive operational or participant data off-chain

ensure evidence packs do not contain personal data unless explicitly authorized and protected

document the data flow and retention rules for the pilot

This policy provides the baseline; pilot-specific privacy constraints should be documented separately as operational addenda.

10. Compliance and Legal-Light Posture

NOORCHAIN 2.1 operates under a “legal-light” posture for controlled deployments:

no custody assumptions

no yield promises

no mandatory KYC/PII collection at the node layer by default

The legal framing is documented in:

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

This privacy policy is technical and operational; it is not a substitute for legal review.

11. Change Control

Any change that affects data handling must be:

documented in docs/CHANGELOG_2.1.md

reviewed as part of release gating under docs/RELEASE_PROCESS_2.1.md

If a change increases data collection or introduces new telemetry, it must be treated as a high-scrutiny change and explicitly highlighted.

12. References

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/AUDIT_READINESS_2.1.md

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/ops/INCIDENTS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md