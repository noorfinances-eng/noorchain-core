NOORCHAIN 2.1 — Operations Playbook

Document ID: OPERATIONS_PLAYBOOK_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Day-2 operations for NOORCHAIN 2.1 nodes (leader/follower), JSON-RPC availability, health endpoints, logs, and operator validation gates.

1. Purpose

This playbook defines operator-standard procedures for running NOORCHAIN 2.1 in controlled environments.

It focuses on:

Safe start/stop discipline

Leader/follower topology and routing expectations

Validation gates (what must be true before proceeding)

Log handling and evidence capture

Common failure modes and operator responses

This playbook is intended to be used alongside runbooks. When a runbook exists for a workflow, it is the preferred step-by-step reference.

2. Operating Principles
2.1 One Step at a Time (Command Discipline)

Operational work must follow a strict discipline:

Verify with one command

Change one thing (edit or start/stop one process)

Re-verify with one command

Only then proceed

This reduces ambiguity, prevents cascading errors, and produces auditable evidence.

2.2 Terminal Roles

Use explicit terminal roles:

T1 (Nodes): start/stop node processes and read node logs only

T2 (Tooling): curl, deploy scripts, Hardhat/viem, inspections, git, docs

T3 (Optional): monitoring dashboards, secondary tail/grep, long-running observers

Avoid mixing node lifecycle commands with tooling calls in the same terminal session.

2.3 Evidence-First Operations

When validating behavior:

Capture command + output (copy/paste logs or transcript)

Prefer deterministic signals (RPC responses, port bindings, receipts, roots)

Avoid “it seems fine” reasoning; require explicit gates

3. System Overview (Operational)
3.1 Node Roles

NOORCHAIN 2.1 supports a leader/follower topology:

Leader node: authoritative for mining/production and authoritative state

Follower node: follows the leader for state correctness where routing is required

Follower behavior may proxy specific RPC calls to the leader when configured. Exact routing rules are normative in docs/RPC_SPEC_2.1.md.

3.2 Data Directories and Locks

Each node uses a distinct -data-dir.

Hard rule: never run two node processes against the same -data-dir.
LevelDB locks are expected and are a safety mechanism.

3.3 RPC and Exposure

Default operator posture:

Bind JSON-RPC to localhost only.

Expose RPC externally only via explicit tunneling (e.g., SSH tunnel) or controlled private networking.

This reduces attack surface and limits accidental public exposure.

4. Configuration Baseline
4.1 Canonical Flags (Operator-Facing)

The following flags are operationally significant:

-chain-id <string> — chain identity (human-readable)

-data-dir <path> — node data directory (must be unique per node)

-p2p-addr <ip:port> — P2P bind address

-rpc-addr <ip:port> — JSON-RPC bind address

-health-addr <ip:port> — health server address (must not collide)

-boot-peers <comma-separated peers> — peer bootstrap list (leader often points to follower and/or vice versa)

-follow-rpc <url> — follower-only; when set, follower routes leader-only reads via this URL

The authoritative list and semantics for all supported flags are maintained by the node binary; runbooks document validated combinations.

4.2 Recommended Defaults

RPC: 127.0.0.1:<port>

Health: 127.0.0.1:<port>

P2P: 127.0.0.1:<port> (local) or private interface (remote)

4.3 Log Placement

For controlled operation, route logs explicitly:

Local dev: ./logs/node{1,2}.log

Temporary: /tmp/noorcore_node{1,2}.log

Avoid relying on inherited stdout/stderr without knowing where it points.

5. Standard Gates (Must-Pass Checks)

These gates define a minimal “node is operational” state.

G0 — Build Gate

Build succeeds

The produced binary is the one being run

Recommended command (T2):

go build -o noorcore ./core

G1 — Process Gate

Node process exists and command line includes expected flags

Recommended command (T2):

pgrep -a noorcore

G2 — Port Bind Gate

P2P/RPC/Health ports are listening on expected addresses

Recommended command (T2):

ss -ltnp | egrep 'noorcore|:3030|:854|:808'

(Adjust ports to your environment.)

G3 — RPC Liveness Gate

eth_chainId returns expected value

eth_blockNumber returns a valid hex quantity

Recommended commands (T2):

curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'

curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}'

G4 — Health Gate

Health endpoint responds successfully on configured address

Recommended command (T2):

curl -s http://127.0.0.1:8081/health || true

(Exact health paths and responses may be defined by the node; treat this as a liveness signal.)

G5 — Multi-Node Connectivity Gate (When Running 2 Nodes)

P2P sessions are established between leader and follower

No repeated connect/disconnect loops

No accumulation of pathological socket states over time

Use runbooks for validated expectations.

G6 — Follower Parity Gate (When Running Follower)

The follower returns the same results as the leader for covered read methods

For leader-routed reads, follower results must be result-equivalent to leader

Normative parity expectations are defined in docs/RPC_SPEC_2.1.md and docs/rpc/RPC_COMPAT_MATRIX_2.1.md.

6. Canonical Procedures
6.1 Stop Nodes (Clean Shutdown)

Goal: stop nodes without leaving ports bound.

Procedure:

Signal termination

Verify process exit

Verify ports closed

Recommended commands (T2):

pkill -INT noorcore || true

sleep 1; pgrep -a noorcore || true

ss -ltnp | egrep 'noorcore|:3030|:854|:808' || true

If a process remains, escalate carefully (document the reason). Prefer graceful handling first.

6.2 Start Leader (Single Node)

Start leader with explicit bind addresses and explicit logs.

Example (T1):

Run from repository root

Provide explicit -data-dir

Redirect logs to a known file

(Use the exact command from the relevant runbook for your milestone/tag.)

Minimum gates after start: G1, G2, G3, G4.

6.3 Start 2-Node Pack (Leader + Follower)

Use the validated runbook for your release tag. As a general model:

Leader runs without -follow-rpc

Follower runs with -follow-rpc <leaderURL> (full URL, including scheme and port)

Each node uses distinct -data-dir and non-colliding P2P/RPC/health ports

Minimum gates after start: G1–G6.

Primary reference:

Multi-node validated runbook(s) under docs/RUNBOOK-* (notably the mainnet-like pack runbook).

7. Alloc and Dev Funding (Operator Notes)

When using an alloc file (development bootstrap):

Alloc is applied at node startup when configured

Alloc application is typically guarded to prevent accidental re-application

To re-apply alloc, an explicit operator action is required (runbook-controlled)

Operational rule: never commit private keys.
Use environment variables or interactive shell input for secrets in tooling sessions.

8. Logs and Diagnostics
8.1 What to Capture in Incidents

Always capture:

pgrep -a noorcore output

ss -ltnp filtered output showing listening ports and PIDs

RPC liveness calls (eth_chainId, eth_blockNumber)

Tail of node logs (last 100–200 lines) from the log files you configured

If relevant, follower and leader outputs for the same RPC request

8.2 Log Tail Commands

Recommended (T2):

tail -n 120 ./logs/node1.log

tail -n 120 ./logs/node2.log

If logs are in /tmp, adjust paths accordingly.

8.3 Avoiding Ambiguous Diagnoses

Do not rely on:

assumptions about which binary is running

assumptions about port ownership

assumptions that “a restart fixed it”

Always show gates.

9. Common Failure Modes and Responses
9.1 “Address already in use” (Port Collision)

Symptoms:

Node fails to start

Health endpoint conflicts (frequent if hardcoded)

Response:

Run G2 gate: identify which PID owns the port

Stop the conflicting process

Ensure each node has a unique -health-addr, -rpc-addr, -p2p-addr

9.2 LevelDB Lock Error

Symptoms:

Node fails to open DB

Error mentions “LOCK” or “resource temporarily unavailable”

Response:

Verify no other node is using the same -data-dir (G1)

Stop the other process or correct the -data-dir

Do not delete locks unless you have proven the process is not running

9.3 RPC “Connection refused”

Symptoms:

curl fails to connect

tooling fails (Hardhat/viem)

Response:

G1: confirm node process exists

G2: confirm RPC port is listening

If not listening, inspect logs for startup failure

Restart only after identifying the cause

9.4 Follower Parity Mismatch

Symptoms:

leader and follower return different values for parity-gated reads

Response:

Verify follower was started with -follow-rpc correctly configured

Compare the same call on leader and follower (record outputs)

Confirm which methods are leader-routed per RPC spec

10. Change Control (Operational)

Operators must not deploy untagged changes into an environment that is treated as “release-like”.

Use the release process in docs/RELEASE_PROCESS_2.1.md

Record changes in docs/CHANGELOG_2.1.md

For RPC behavior changes, comply with docs/API_STABILITY_POLICY_2.1.md

11. Backups and Recovery (Baseline Posture)

NOORCHAIN 2.1 expects operators to treat node data as critical.

Baseline expectations:

Backups must be performed with the node stopped or via a documented consistent snapshot procedure.

Backup/restore workflows must be validated on non-production environments before relying on them.

Detailed backup and disaster recovery runbooks may exist separately; if they do, this playbook defers to them as normative for procedures.

12. References

Normative and operational references:

docs/ARCHITECTURE_2.1.md

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/RELEASE_PROCESS_2.1.md

Runbooks under docs/RUNBOOK-*.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/ops/INCIDENTS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md