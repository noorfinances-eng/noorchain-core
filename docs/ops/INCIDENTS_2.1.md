NOORCHAIN 2.1 — Incidents (Operational Catalog)

Document ID: INCIDENTS_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Provides an operational incident catalog for NOORCHAIN 2.1, including common incident types, detection signals, immediate containment actions, required evidence, and recovery validation gates. This document complements the incident response process and is intended for operators.

Purpose

This document defines a practical incident catalog for NOORCHAIN 2.1 operations.

It is designed to:

standardize diagnosis and containment for common failure modes

reduce recovery time by mapping symptoms to actions

enforce evidence preservation and audit readiness

define minimum validation gates after recovery

This document does not replace the incident response lifecycle; it complements it with operator-facing patterns.

References:

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

Incident Handling Conventions

2.1 Severity Mapping

Use severity classification from:

docs/governance/INCIDENT_RESPONSE_2.1.md

This catalog focuses on likely SEV-1/SEV-2/SEV-3 cases, but includes SEV-0 triggers.

2.2 Operator Discipline

During incidents:

execute one action at a time

record exact commands and outputs

avoid destructive actions unless explicitly approved

preserve logs before restarting or reconfiguring nodes

Reference:

docs/AUDIT_READINESS_2.1.md

Catalog of Common Incidents

3.1 Node Fails to Start

Symptoms:

process exits immediately

log shows panic / fatal error

ports are not listening

Detection:

pgrep -a noorcore does not show expected process

ss -lntp shows no LISTEN on expected ports

logs contain startup failure lines

Immediate containment:

stop any partial processes

prevent repeated restart loops until cause is identified

Diagnosis checklist:

verify binary exists and is correct build

verify flags: chain-id, data-dir, p2p-addr, rpc-addr, health-addr

verify data-dir permissions

Evidence to capture:

node log excerpt around failure

pgrep -a noorcore and ss -lntp outputs

git rev-parse HEAD + tag reference (if applicable)

Recovery gates:

node starts

RPC reachable

eth_chainId correct

health endpoint responds (if enabled)

References:

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

3.2 LevelDB Lock Conflict (Data Directory Shared)

Symptoms:

error indicating LevelDB “LOCK” or “resource temporarily unavailable”

node refuses to start

Detection:

logs show lock failure

multiple processes share the same -data-dir

Immediate containment:

stop the duplicate process

confirm only one process uses each data-dir

Diagnosis checklist:

run process listing and inspect args

inspect data-dir for LOCK files (do not delete unless approved)

Evidence to capture:

pgrep -a noorcore

error log line showing lock failure

mapping of node → data-dir

Recovery gates:

node starts successfully with unique data-dir

ports listening correctly

References:

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

3.3 Port Collision (RPC / P2P / Health)

Symptoms:

bind errors on startup

one node starts, other fails

unexpected service behavior due to wrong process holding port

Detection:

logs show “bind: address already in use”

ss -lntp shows port already LISTEN by a different PID

Immediate containment:

identify port owner PID

stop conflicting process (if authorized)

reassign ports only if necessary and recorded

Evidence to capture:

ss -lntp lines for collided port

process command lines for involved PIDs

Recovery gates:

each node listens on intended unique ports

RPC identity gates pass

References:

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

3.4 RPC Down / Connection Refused

Symptoms:

curl fails to connect

tooling errors (Hardhat/Viem) reporting ECONNREFUSED

health endpoint may still be up (depending on wiring)

Detection:

ss -lntp shows no LISTEN on rpc port

curl to JSON-RPC fails

Immediate containment:

confirm node is running and arguments are correct

if node is running but RPC not bound, stop and restart with correct rpc-addr

avoid exposing RPC publicly as a workaround

Evidence to capture:

ss -lntp output

curl error output

node logs around RPC bind

Recovery gates:

eth_chainId returns expected value

eth_blockNumber returns valid quantity

References:

docs/RPC_SPEC_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

3.5 Leader/Follower Parity Failure

Symptoms:

follower returns different values for identity or read methods

parity gates fail (stateRoot mismatch, block metadata mismatch, balance/nonce mismatch)

Detection:

compare leader vs follower outputs per parity matrix

follower may be missing follow-rpc configuration or routing is broken

Immediate containment:

if follower serves clients, restrict access temporarily

validate FollowRPC configuration

route leader-only methods to leader and retest

Evidence to capture:

side-by-side curl outputs showing mismatch

follower start command including follow-rpc

logs indicating routing behavior (if available)

Recovery gates:

parity gates defined by RPC_COMPAT_MATRIX pass

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

3.6 P2P Sessions Not Established

Symptoms:

nodes run but no peer connections

follower cannot follow or parity degrades due to isolation

“boot peers unreachable” behavior

Detection:

ss shows no ESTAB sessions on p2p ports

logs show repeated dial attempts or timeouts

Immediate containment:

confirm p2p-addr bindings and boot peers list

confirm firewall / security groups allow P2P (private network)

restart nodes only after configuration is verified

Evidence to capture:

ss evidence (LISTEN + ESTAB state)

start commands (boot-peers)

any firewall evidence (if applicable)

Recovery gates:

at least one stable P2P ESTAB session for required topology

parity gates pass

References:

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

3.7 World-State Reads Return Zero/Empty Unexpectedly

Symptoms:

eth_getBalance returns 0 unexpectedly

eth_getTransactionCount returns 0 unexpectedly

eth_getCode returns 0x unexpectedly for deployed contract

eth_getStorageAt returns 0x unexpectedly for known storage

Detection:

compare against expected evidence pack outputs

check if reads are being served from follower without routing

check if wrong block tag or wrong address/contract used

Immediate containment:

confirm chainId and correct RPC endpoint

verify address and contract addresses match the environment

run state read on leader directly

verify follower routing (FollowRPC) for leader-only reads

Evidence to capture:

exact JSON-RPC request/response pairs

block tag used ("latest" vs specific)

environment contract address registry (if maintained)

Recovery gates:

reads match expected values on leader

follower parity matches per matrix

References:

docs/STATE_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

3.8 Receipt Missing or Incomplete

Symptoms:

eth_getTransactionReceipt returns null for a known submitted tx

receipt missing contractAddress/logs unexpectedly

tooling fails on deploy due to missing receipt fields

Detection:

tx hash exists but receipt retrieval fails

compare leader vs follower receipt results

Immediate containment:

ensure tx was accepted and mined (check blockNumber/tx index if available)

query on leader directly

verify persistence storage and receipt path compatibility

Evidence to capture:

tx hash and raw request

receipt responses from leader and follower

blockNumber around submission time

Recovery gates:

receipt is available and stable across restart

required fields are present per compatibility expectations

References:

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/STATE_MODEL_2.1.md

3.9 Suspected Key Compromise (SEV-0 Trigger)

Symptoms:

unauthorized signer activity

unexpected validator signatures or governance actions

unauthorized access to node hosts

Immediate containment:

isolate affected systems

rotate credentials and keys per policy

stop network if integrity is at risk (permissioned context)

preserve full evidence pack

Evidence to capture:

logs and access records

timeline and impacted assets

signer registry and recent actions

References:

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

Standard Evidence Pack Checklist (Incidents)

Every incident evidence pack should include:

severity class and incident ID

timeline (UTC)

impacted nodes and roles

node start commands

logs (bounded time)

ss -lntp output

JSON-RPC request/response pairs for failing checks

git commit hash / tag reference

recovery actions executed and validation outputs

Reference:

docs/AUDIT_READINESS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

Recovery Validation Gates (Minimum)

After recovery, always execute:

eth_chainId

eth_blockNumber

eth_getBlockByNumber("latest", false)

If multi-node:

leader/follower parity gates per matrix

P2P session established

If state-related:

eth_getBalance and eth_getTransactionCount for known addresses

eth_getCode for known contract addresses (if applicable)

References:

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/STATE_MODEL_2.1.md

Change Control

Updates to this incident catalog should be:

versioned

aligned with incident response policy

reflected in changelog if they alter operational expectations

References:

docs/governance/INCIDENT_RESPONSE_2.1.md

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

References

docs/governance/INCIDENT_RESPONSE_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/rpc/RPC_ERROR_MODEL_2.1.md
docs/STATE_MODEL_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/THREAT_MODEL_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/genesis/PARAMETERS_REFERENCE_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md
docs/CHANGELOG_2.1.md
docs/RELEASE_PROCESS_2.1.md