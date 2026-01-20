NOORCHAIN 2.1 — Threat Model

Document ID: THREAT_MODEL_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Threat enumeration and mitigations for NOORCHAIN 2.1 in controlled deployments (permissioned operational control).

1. Purpose

This document enumerates credible threats to NOORCHAIN 2.1 and maps them to mitigations, detection mechanisms, and residual risks.

It is intended to be used with:

docs/SECURITY_TRUST_MODEL_2.1.md (assumptions and boundaries)

docs/OPERATIONS_PLAYBOOK_2.1.md (procedures and gates)

docs/ops/INCIDENTS_2.1.md and docs/governance/INCIDENT_RESPONSE_2.1.md (response playbooks)

release gating and audit readiness materials

This document does not claim to cover all adversaries; it is scoped to the deployment model described.

2. Scope and Assumptions
2.1 Deployment Assumption

NOORCHAIN 2.1 is deployed in a controlled environment with permissioned operational control.

Operators are trusted for:

correct configuration

restricted RPC exposure

disciplined upgrade procedures

2.2 Security Boundary Summary

Threats are analyzed across these boundaries:

RPC interface (JSON-RPC)

networking (P2P)

persistence (databases under -data-dir)

release and supply-chain discipline

application-layer (PoSS) misuse

3. Assets

The model protects the following assets.

3.1 Integrity Assets

Chain history (blocks, headers)

Transaction inclusion semantics (tx → receipt)

World-state correctness and persistence (accounts, code, storage, roots)

Receipt correctness and availability (tx receipts, logs/bloom)

3.2 Availability Assets

Node liveness (leader/follower)

RPC endpoint availability

Health endpoint availability (where enabled)

3.3 Governance and Process Assets

Repository integrity (commits, tags)

Release gating evidence and changelog correctness

Operational runbook correctness

3.4 Confidentiality Assets (Limited)

Private keys and secrets (operator/tooling)

Non-public deployment topology details (optional)

Logs that might expose internal addresses or request payloads

Note: On-chain payloads are public by nature; confidentiality for on-chain content is not assumed.

4. Adversaries
4.1 External Remote Actor

Capabilities:

send malformed or abusive JSON-RPC requests (if RPC exposed)

attempt resource exhaustion

attempt to exploit parsing/validation weaknesses

4.2 Malicious or Misconfigured Peer

Capabilities:

send malformed P2P traffic

attempt connection churn, spam, or inconsistent behavior

4.3 Operator Error (Primary Real-World Risk)

Capabilities:

start two nodes on the same data directory

expose RPC publicly

deploy untagged binaries

break follower routing assumptions

4.4 Repository / Supply Chain Attacker

Capabilities:

compromise maintainer credentials

inject malicious dependencies or code changes

trick release process into tagging unreviewed code

5. Threat Enumeration (STRIDE-Inspired)

Each threat includes:

Threat

Impact

Mitigations

Detection

Residual Risk

Severity is contextual and depends on exposure; in controlled deployments, the highest risks often come from configuration and operational discipline.

6. RPC Threats
T-RPC-01 — RPC exposed publicly without controls

Impact:
Unauthorized access, DoS, data leakage via logs, unexpected load, potential exploitation of parser bugs.

Mitigations:

Bind RPC to localhost by default.

Use SSH tunneling or private networks for remote access.

If public exposure is required: authentication, rate limiting, allowlists.

Detection:

Operator audits of listen addresses (ss -ltnp).

External checks against public IP surfaces.

Monitoring abnormal request volume.

Residual Risk:
If public exposure occurs, the node becomes an internet-facing service and requires stronger hardening.

T-RPC-02 — Resource exhaustion via heavy RPC calls

Impact:
RPC becomes unresponsive, node instability, degraded liveness.

Mitigations:

Strict parameter validation and bounded execution.

Input size limits and request rate limiting (proxy or node-level if implemented).

Prefer leader-only routing for expensive state reads if follower is present.

Detection:

Increased latency and timeout rates.

CPU/memory spikes.

Log patterns indicating abusive calls.

Residual Risk:
Bounded but not eliminated if RPC must serve untrusted clients.

T-RPC-03 — Malformed JSON / parameter confusion

Impact:
Crashes, inconsistent behavior, bypassed validation, incorrect results.

Mitigations:

Strict decoding and explicit param typing.

Reject ambiguous encodings.

Stable error model and negative tests.

Detection:

Panic traces in logs.

Unexpected error patterns.

Regression tests and gate scripts.

Residual Risk:
Reduced by strict parsing; never fully eliminated.

T-RPC-04 — Semantic drift between spec and implementation

Impact:
Clients behave incorrectly, integration breaks, audit failures.

Mitigations:

“Spec is law” discipline.

Release gates requiring spec alignment and compatibility checks.

Changelog entries for semantic changes.

Detection:

Compatibility matrix regressions.

Parity mismatches (leader vs follower).

Auditor review and runbook evidence.

Residual Risk:
Primarily process-driven; requires discipline.

7. Networking (P2P) Threats
T-NET-01 — Peer spoofing in weakly controlled networks

Impact:
Unexpected peers connect, traffic injection, load increase.

Mitigations:

Controlled peer lists (-boot-peers discipline).

Firewall rules restricting allowed peers.

Explicit network segmentation in VPS deployments.

Detection:

Unexpected peer addresses in logs.

Connection churn and unknown peer IDs.

Residual Risk:
Low in properly controlled environments; higher if P2P is exposed broadly.

T-NET-02 — Connection churn / spam

Impact:
Resource exhaustion, instability, reduced availability.

Mitigations:

Persistent session management where applicable.

Rate limiting of inbound connection attempts (OS/firewall level).

Monitoring and auto-block rules in controlled deployments.

Detection:

Frequent connect/disconnect logs.

Socket state anomalies.

Residual Risk:
Depends on P2P exposure.

T-NET-03 — Protocol parsing vulnerabilities

Impact:
Potential crashes or exploitation via malformed messages.

Mitigations:

Defensive decoding.

Fuzzing where feasible (future hardening).

Minimal trusted surface; keep dependencies controlled.

Detection:

Panics in logs.

Unexpected disconnect sequences.

Residual Risk:
Non-zero; addressed via hardening and testing maturity.

8. Persistence and State Threats
T-STATE-01 — Two processes use the same data directory

Impact:
Database corruption, inconsistent state, undefined behavior.

Mitigations:

Hard operational rule: unique -data-dir per node.

LevelDB lock enforcement.

Operator gates before start (process + port + lock awareness).

Detection:

LevelDB lock errors.

Duplicate process args.

Unexplained state corruption after simultaneous runs.

Residual Risk:
Primarily operational; preventable with discipline.

T-STATE-02 — Unsafe backup/restore leading to inconsistent DB snapshots

Impact:
State divergence, corrupted world-state, missing receipts.

Mitigations:

Stop node before backups unless a consistent snapshot procedure is documented.

Validate restore in a non-production environment.

Maintain backup runbooks and test them.

Detection:

Post-restore gate failures (state reads, receipts, roots).

Node errors during startup.

Residual Risk:
Moderate without a mature backup/restore runbook.

T-STATE-03 — State root / receipt root inconsistency

Impact:
Incorrect block metadata, client verification failures, audit issues.

Mitigations:

Persist and validate roots deterministically.

Ensure RPC exposes the same semantics as stored metadata.

Gate checks for roots across restarts and between leader/follower.

Detection:

Inconsistent eth_getBlockByNumber fields.

Regression tests and parity checks.

Residual Risk:
Implementation-dependent; requires continuous gating.

T-STATE-04 — Receipt/log persistence mismatch

Impact:
Tooling breaks (wallets, explorers), inability to prove inclusion, audit findings.

Mitigations:

Stable receipt persistence keys/layout.

RPC must read receipts deterministically and consistently.

Compatibility gates for deploy/tx workflows (Hardhat/viem).

Detection:

Missing receipts for known tx hashes.

Tooling failures during deployments.

Compatibility matrix regressions.

Residual Risk:
Moderate; improves with test coverage.

9. Operational Threats
T-OPS-01 — Port collisions (RPC/Health/P2P)

Impact:
Nodes fail to start, services bind unexpectedly, misrouting.

Mitigations:

Explicit bind flags for each node.

Unique health ports per node (-health-addr).

Gate check for listening ports before and after start.

Detection:

“address already in use” errors.

ss -ltnp shows unexpected owners.

Residual Risk:
Low with discipline.

T-OPS-02 — Deploying untagged or unvalidated code

Impact:
Unknown behavior, inability to audit, regressions.

Mitigations:

Release process discipline: tag only after gates.

Change logging required for behavioral changes.

Operators run only tagged releases in release-like environments.

Detection:

Binary/version mismatch vs expected tag.

Missing changelog entries for observed behavior.

Residual Risk:
Process-driven; requires strong operational governance.

T-OPS-03 — Misconfigured follower routing

Impact:
Follower returns incorrect state reads; clients see inconsistent results.

Mitigations:

Use -follow-rpc as authoritative for routing.

Define parity gates for follower.

Document leader-only reads in RPC spec.

Detection:

Parity mismatch between leader and follower for the same call.

Tooling failures depending on follower reads.

Residual Risk:
Moderate; mitigated by explicit gates.

10. Supply Chain and Governance Threats
T-SC-01 — Repository compromise (maintainer credentials)

Impact:
Malicious code merges, compromised releases/tags.

Mitigations:

Strong access controls on repository.

Mandatory review discipline.

Tag creation restricted to trusted maintainers.

Audit trails for merges and tags.

Detection:

Unexpected commits/tags.

Unreviewed changes in protected branches.

Residual Risk:
Never zero; governed by operational security of maintainers.

T-SC-02 — Dependency compromise

Impact:
Backdoored libraries, build-time or runtime compromise.

Mitigations:

Controlled dependency updates.

Prefer stable, pinned versions.

Review dependency diffs; avoid unbounded upgrades.

Detection:

Unexpected go.mod changes.

Build anomalies.

Residual Risk:
Moderate; depends on dependency hygiene.

11. PoSS Application-Layer Threats
T-POSS-01 — Misinterpretation of PoSS as consensus/security

Impact:
Incorrect trust assumptions, governance confusion, misleading claims.

Mitigations:

Explicit separation in architecture and trust model.

Documentation clarity: PoSS is application-layer and off-consensus.

Communications policy controls external claims.

Detection:

Documentation drift, inconsistent public messaging.

Residual Risk:
Human/process risk; mitigated by strict communications discipline.

T-POSS-02 — Personal data embedded in on-chain snapshot payloads

Impact:
Privacy violations, compliance issues, irreversible publication.

Mitigations:

Privacy policy: avoid PII on-chain.

Store references/hashes; keep sensitive records off-chain.

Pilot-specific data handling rules.

Detection:

Review of snapshot payload formats.

Tooling checks in submission scripts (future).

Residual Risk:
Moderate; requires discipline in application-layer design.

12. Detection and Monitoring Baseline

Operators should maintain baseline detection signals:

process listing and command-line verification

port binding verification

RPC liveness and latency checks

follower parity checks (leader vs follower)

log scanning for panic/error patterns

resource monitoring (CPU/memory/disk growth)

These are operational controls, not protocol-level guarantees.

13. Residual Risk Summary

Under controlled deployments:

The dominant risks are operational misconfiguration, RPC exposure, and spec drift.

Protocol parsing and dependency risks remain non-zero and require incremental hardening.

Privacy risks are primarily about what is placed on-chain and how logs/evidence are handled.

14. Change Control

Any change affecting this threat model requires:

version bump

changelog entry

alignment review against:

security trust model

incident response documents

operations playbook

15. References

docs/SECURITY_TRUST_MODEL_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

docs/ops/INCIDENTS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md