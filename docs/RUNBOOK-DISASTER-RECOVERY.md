NOORCHAIN 2.1 — Runbook: Disaster Recovery

Document ID: RUNBOOK-DISASTER-RECOVERY
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines disaster recovery (DR) procedures for NOORCHAIN 2.1 controlled and permissioned deployments, including failure scenarios, recovery objectives, decision gates, restoration steps, and evidence requirements. This runbook assumes a permissioned operator context and a Legal-Light posture.

1. Purpose

This runbook defines a controlled approach to recover NOORCHAIN 2.1 operations after severe incidents that materially affect availability, integrity, or operator control.

It is designed to:

provide a deterministic recovery path under operator control

minimize the risk of forks or identity drift

define decision gates for restore vs redeploy

preserve evidence for audit readiness

align recovery actions with governance and Legal-Light posture

This runbook complements the incident response process and the backup/restore runbook.

References:
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/RUNBOOK-BACKUP-RESTORE.md

2. Definitions and Objectives

2.1 Disaster vs Incident

Incident: operational degradation or failure recoverable with standard procedures.

Disaster: loss of node host(s), loss of persistent state, compromised control plane, or systemic failure requiring restoration from backups or redeployment.

Reference:
docs/governance/INCIDENT_RESPONSE_2.1.md

2.2 Recovery Objectives

Objectives must be defined per environment:

RTO (Recovery Time Objective): maximum acceptable downtime.

RPO (Recovery Point Objective): maximum acceptable data loss window.

In controlled deployments, RPO is bounded by the most recent validated backup.

2.3 DR Artifacts

Known-good release: commit/tag used to produce the binary and runbooks.

Genesis anchor: genesis artifact digest for the environment (if applicable).

Backup artifacts: archives of node data-dirs with integrity manifests.

References:
docs/RELEASE_PROCESS_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/AUDIT_READINESS_2.1.md

3. DR Triggers (Non-Exhaustive)

A DR event may be declared when any of the following occur:

total loss of leader host and inability to recover local disk

simultaneous loss of leader and follower hosts

persistent state corruption that cannot be repaired via standard incident procedures

compromise of signing keys or operator access that threatens integrity

confirmed unauthorized modification of binaries or node data

repeated parity gate failures across nodes after attempted recovery

Reference:
docs/ops/INCIDENTS_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md

4. DR Roles and Authority

4.1 Incident Commander (IC)

IC is responsible for:

declaring DR event and severity

coordinating operators

authorizing containment actions

ensuring evidence preservation

4.2 Governance / Multisig Authority (When Required)

For high-impact decisions (e.g., signer rotation, validator set changes, redeployment affecting identity), governance approval may be required.

Reference:
docs/governance/MULTISIG_OPERATIONS_2.1.md
docs/governance/GOVERNANCE_MODEL_2.1.md

4.3 Operators

Operators are responsible for:

executing recovery steps

capturing evidence and validation outputs

preventing configuration drift

5. Decision Gates: Restore vs Redeploy

A DR recovery must choose one of two paths:

Path A — Restore from Backup (Preferred)

Choose Path A when:

genesis identity remains valid and unchanged

backups exist for the environment and pass integrity checks

the known-good release is available and reproducible

control-plane keys are not compromised, or can be rotated safely without identity drift

Path B — Redeploy Environment (Controlled Re-Anchor)

Choose Path B when:

backups are missing or corrupted

integrity cannot be guaranteed

compromise requires re-anchoring identity under governance decision

environment is pilot-only and redeploy is acceptable within scope

Path B is high impact and may require governance Class 3 approval depending on whether genesis changes are involved.

References:
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/governance/UPGRADE_PROCESS_2.1.md

6. DR Preparation Baseline (Must Exist Before Disaster)

This section defines what must already exist to enable DR.

Required:

at least one validated backup per active node (leader and follower if used)

backup integrity manifests with hashes

known-good release tag/commit and reproducible build procedure

recorded node start commands and port bindings

a secure storage location for backups and manifests

a signer registry record (if multi-sig used)

References:
docs/RUNBOOK-BACKUP-RESTORE.md
docs/RELEASE_PROCESS_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/governance/MULTISIG_OPERATIONS_2.1.md

7. DR Procedure — Path A: Restore from Backup
7.1 Containment and Evidence Preservation

Declare DR event (IC).

Isolate affected hosts (if compromise suspected).

Preserve:

logs (bounded time windows)

configuration files used for node start

any forensic artifacts required (if security incident)

Reference:
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md

7.2 Rebuild the Binary from Known-Good Release

check out the approved commit/tag

build the node binary using documented build procedure

record build outputs and binary digest if policy requires

Reference:
docs/RELEASE_PROCESS_2.1.md
docs/AUDIT_READINESS_2.1.md

7.3 Verify Backup Integrity

For each node backup:

verify sha256 matches manifest

verify archive size and inventory

reject any artifact failing integrity

If integrity fails for leader backup, evaluate fallback leader backup (older) vs Path B.

Reference:
docs/RUNBOOK-BACKUP-RESTORE.md

7.4 Restore Leader First

provision a new host or clean host environment

restore leader data-dir from backup to the intended path

ensure only one process uses the restored data-dir

start leader using recorded command (adjusting only host-specific binds; record changes)

7.5 Validate Leader Recovery Gates

Minimum gates:

process running

RPC reachable

eth_chainId matches expected value

eth_blockNumber returns valid quantity

eth_getBlockByNumber("latest", false) returns coherent structure (roots/bloom as applicable)

If world-state gates are used:

eth_getBalance for known funded address

eth_getTransactionCount for known active address

eth_getTransactionReceipt for known recent tx hash (if available)

References:
docs/RPC_SPEC_2.1.md
docs/STATE_MODEL_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md

7.6 Restore Follower (If Applicable)

restore follower data-dir from backup

configure follower FollowRPC to leader

start follower with distinct ports and data-dir

confirm P2P sessions (if used) and routing correctness

7.7 Validate Multi-Node Parity Gates

eth_chainId parity

eth_blockNumber coherence

eth_getBlockByNumber("latest", false) coherent metadata

leader-only methods parity via follower routing:

eth_getBalance

eth_getTransactionCount

other required matrix methods

References:
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md

8. DR Procedure — Path B: Redeploy Environment (Controlled)

Path B is used when restore is not feasible or not safe.

8.1 Governance Decision Gate

Before redeploy:

IC documents why restore is impossible/unsafe

governance approval is obtained if identity changes or validator changes are involved

record the approved plan (commit/tag, genesis policy, parameters)

References:
docs/governance/GOVERNANCE_MODEL_2.1.md
docs/governance/MULTISIG_OPERATIONS_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md

8.2 Generate or Re-Issue Genesis (If Applicable)

If redeploy requires genesis generation:

follow genesis checklist

record genesis digest and intent record

ensure chain identity and EVM chainId consistency

produce a genesis evidence pack

References:
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md

8.3 Deploy Nodes on Clean Data-Dirs

provision clean hosts

use clean data-dirs

start leader then follower per deployment model

keep RPC localhost-only; access via SSH tunnels for pilots

Reference:
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

8.4 Post-Deploy Validation Gates

identity gates (eth_chainId)

liveness gates (blockNumber, block reads)

world-state read gates (if enabled)

tx path gates (if required for environment)

parity gates (leader/follower)

Reference:
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md

9. Key Compromise and Rotation in DR Context

If DR involves suspected key compromise:

treat as SEV-0/SEV-1 under incident response

rotate compromised keys per multi-sig policy

ensure rotation actions are governed and evidenced

do not proceed with upgrades or redeploys until integrity is restored

References:
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/governance/MULTISIG_OPERATIONS_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/THREAT_MODEL_2.1.md

10. Evidence Pack (Disaster Recovery)

A DR evidence pack MUST include:

DR declaration record (incident ID, severity, timestamps UTC)

decision gate record (Path A vs Path B with reasons)

release tag/commit used for recovery

backup manifests and integrity checks (Path A)

restored node start commands and port binds

validation outputs (identity/liveness/state/tx/parity as applicable)

any key rotation records (if applicable)

post-recovery monitoring notes (initial stability window)

Reference:
docs/AUDIT_READINESS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md

11. Post-DR Review and Change Control

After DR stabilization:

run a post-incident review

identify root causes and prevention actions

update runbooks and policies where required

update changelog if operational behavior or procedures changed

References:
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/CHANGELOG_2.1.md
docs/RELEASE_PROCESS_2.1.md

12. References

docs/RUNBOOK-BACKUP-RESTORE.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/ops/INCIDENTS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/governance/GOVERNANCE_MODEL_2.1.md
docs/governance/MULTISIG_OPERATIONS_2.1.md
docs/governance/UPGRADE_PROCESS_2.1.md
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/genesis/PARAMETERS_REFERENCE_2.1.md
docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/STATE_MODEL_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/THREAT_MODEL_2.1.md
docs/legal/LEGAL_LIGHT_POSTURE_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md