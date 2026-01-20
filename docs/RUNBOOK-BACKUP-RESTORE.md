NOORCHAIN 2.1 — Runbook: Backup and Restore

Document ID: RUNBOOK-BACKUP-RESTORE
Version: v1.0
Date: January 2026
Status: Active
Scope: Operational runbook to create, verify, store, and restore backups for NOORCHAIN 2.1 node data directories in controlled deployments (local, pilot, permissioned). This runbook is designed for audit-ready evidence and deterministic recovery.

1. Purpose

This runbook defines a controlled procedure to:

create consistent backups of node state (NOOR DB + EVM state DB)

verify backup integrity (hashes, sizes, inventory)

restore a node from a backup without configuration drift

validate recovery via standard RPC gates and parity checks (if multi-node)

This runbook assumes a conservative “offline backup” approach: nodes are stopped before backup to avoid partial LevelDB snapshots.

2. Definitions

Node data-dir: the filesystem directory passed via -data-dir, containing all node persistent state.

NOOR DB: the node’s internal LevelDB store under the data-dir (implementation-defined paths).

EVM state DB (geth store): geth-compatible LevelDB store under <data-dir>/db/geth (if enabled in the environment).

Backup artifact: a single archive containing a complete data-dir snapshot plus an inventory manifest and hashes.

3. Preconditions and Safety Rules
3.1 Access and Authorization

Operator must have authorized shell access to the host.

For pilot / permissioned environments, backup/restore is a controlled operation and should be pre-approved.

Reference:
docs/governance/GOVERNANCE_MODEL_2.1.md
docs/governance/MULTISIG_OPERATIONS_2.1.md

3.2 Offline Backup Rule (Required)

Do not back up a running node’s LevelDB directories unless a tested snapshot mechanism exists for that deployment.

Default procedure requires a clean stop before creating the archive.

3.3 Data Segregation Rule

Each node has its own distinct -data-dir.

Never merge data from different nodes into the same directory.

Reference:
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/ops/INCIDENTS_2.1.md

3.4 Sensitive Material Rule

Do not include private keys, seed phrases, or secrets in backup artifacts unless the environment explicitly requires it and storage is encrypted and access-controlled.

If a keystore is present under a data-dir, it must be treated as sensitive.

Reference:
docs/SECURITY_TRUST_MODEL_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md

4. Backup Targets (What Must Be Included)

A full node backup MUST include the entire -data-dir, including (if present):

NOOR DB (LevelDB) directories and files

EVM state DB under db/geth

node metadata files required for restart

any block metadata substore (if separate)

configuration overlays stored inside the data-dir (if applicable)

A backup SHOULD also include (as separate artifacts, not necessarily inside the same archive):

node log excerpts covering the last healthy window (bounded time)

the node start command used (as an evidence record)

the environment parameter record (chainId, ports, FollowRPC, boot peers)

Reference:
docs/genesis/PARAMETERS_REFERENCE_2.1.md
docs/AUDIT_READINESS_2.1.md

5. Backup Naming and Inventory Standard
5.1 Artifact Naming

Use a naming format that is deterministic and operator-friendly:

noorchain2.1_<env>_<node>_<yyyy-mm-dd>_<utcHHMM>_<commitOrTag>.tar.gz

Examples:

noorchain2.1_local_node1_2026-01-20_2315_190fbc7.tar.gz

noorchain2.1_pilotA_node2_2026-01-20_2315_M10-MAINNETLIKE-STABLE.tar.gz

5.2 Required Inventory Manifest

Each backup operation MUST produce a manifest (text) stored alongside the archive:

environment name

node identifier (leader/follower)

data-dir absolute path

node start command (as last known good)

git commit hash and tag (if applicable)

archive filename

archive size

archive sha256 hash

timestamp (UTC)

operator identity (internal)

This manifest is part of the evidence pack.

Reference:
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md

6. Backup Procedure (Offline, Deterministic)
6.1 Identify the Node and Its Data Directory

Record:

node role (leader/follower)

-data-dir

bound ports (-p2p-addr, -rpc-addr, -health-addr)

if follower: -follow-rpc

6.2 Stop the Node Cleanly

Stop the node using the standard operator method for the environment.

Confirm the process is gone and ports are closed.

Operational requirement:

Do not proceed until the process is stopped and no LISTEN remains on the node’s ports.

Reference:
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/ops/INCIDENTS_2.1.md

6.3 Create the Archive

Create a compressed archive of the full data-dir.

Ensure permissions are preserved.

Avoid including unrelated directories outside the data-dir.

The archive must represent a single consistent snapshot.

6.4 Generate Integrity Hash

Compute sha256 for the archive.

Record it into the manifest.

6.5 Restart and Validate (Optional but Recommended)

After backup is created, restart the node and run minimal gates:

eth_chainId

eth_blockNumber

eth_getBlockByNumber("latest", false)

If leader/follower setup, re-run parity gates per matrix.

Reference:
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/RPC_SPEC_2.1.md

7. Storage and Retention
7.1 Storage Requirements

Backup artifacts must be stored in:

access-controlled storage

preferably encrypted at rest

with integrity hashes retained separately (inventory manifest)

7.2 Retention Policy (Baseline)

Recommended baseline for controlled deployments:

keep the most recent N backups (policy-defined)

keep at least one “known-good” backup associated with a stable release tag

keep incident-related backups for the incident retention window

Retention policy must be consistent with audit readiness.

Reference:
docs/AUDIT_READINESS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md

8. Restore Procedure (Single Node)
8.1 Restore Preconditions

Before restoring:

identify the target environment and expected chain identity

confirm you have the correct backup artifact and manifest

confirm you have the node start command and intended ports

confirm no other node is using the target data-dir (avoid LevelDB lock conflicts)

Reference:
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/ops/INCIDENTS_2.1.md

8.2 Verify Artifact Integrity

verify sha256 matches the manifest

verify archive size is within expected bounds

if integrity fails: do not restore; treat as incident

8.3 Prepare the Target Data Directory

choose a clean target path (empty directory)

ensure filesystem permissions match operator needs

ensure adequate disk space

8.4 Extract the Archive

extract the archive to the target location

verify the restored directory contains expected subpaths (including db/geth if applicable)

8.5 Start the Node

Start using the recorded “last known good” command, adjusted only for environment-specific constraints (ports, hostnames, bind addresses). Any changes must be recorded.

8.6 Post-Restore Validation Gates

Minimum gates:

node process is running

RPC is reachable

eth_chainId matches expected value

eth_blockNumber returns valid quantity

eth_getBlockByNumber("latest", false) returns coherent structure

If state read gates are part of the environment:

eth_getBalance (known address)

eth_getTransactionCount (known address)

Reference:
docs/RPC_SPEC_2.1.md
docs/STATE_MODEL_2.1.md

9. Restore Procedure (Leader + Follower)
9.1 Restore Order

Recommended order for controlled deployments:

restore leader first

validate leader RPC gates

restore follower second

validate follower routing and parity gates

9.2 Follower Routing Validation

Follower must satisfy:

FollowRPC configured to leader

leader-only reads route correctly (per compat matrix)

parity gates pass for the required set

Reference:
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md

10. Failure Modes and Escalation

If any of the following occur, stop and escalate under incident response:

restore artifact integrity mismatch

LevelDB lock conflicts that persist after confirming single process ownership

chain identity mismatch (eth_chainId unexpected)

repeated internal errors on basic RPC gates

parity gates fail after restore in multi-node setup

Reference:
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/ops/INCIDENTS_2.1.md

11. Evidence Pack (Backup/Restore)

A backup/restore evidence pack MUST include:

backup manifest (inventory + hashes)

git commit hash / release tag reference

node start command(s)

stop/start timestamps (UTC)

validation gate outputs after backup (if performed)

validation gate outputs after restore

parity outputs (if multi-node)

Reference:
docs/AUDIT_READINESS_2.1.md
docs/CHANGELOG_2.1.md
docs/RELEASE_PROCESS_2.1.md

12. References

docs/OPERATIONS_PLAYBOOK_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/ops/INCIDENTS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/rpc/RPC_ERROR_MODEL_2.1.md
docs/RPC_SPEC_2.1.md
docs/STATE_MODEL_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/genesis/PARAMETERS_REFERENCE_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md