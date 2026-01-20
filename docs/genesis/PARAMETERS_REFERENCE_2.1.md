Document ID: PARAMETERS_REFERENCE_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Canonical reference for genesis- and operations-relevant parameters for NOORCHAIN 2.1 controlled deployments. This document enumerates parameter names, intended values, constraints, and validation methods.

Purpose

This document defines the authoritative parameter set used to configure and validate NOORCHAIN 2.1 environments.

It is designed to:

prevent configuration drift across nodes

make chain identity and EVM identity explicit

provide operator-checkable validation methods

support audit readiness by standardizing parameter recording

This reference is used by:

docs/genesis/GENESIS_CHECKLIST_2.1.md

docs/genesis/GENESIS_SPEC_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

Scope and Conventions

2.1 Scope

This reference covers parameters that affect:

genesis identity and determinism

node runtime behavior (ports, roles, follower routing)

RPC exposure posture

alloc overlay policy inputs (where applicable)

multi-node pack coherence and parity gates

It does not attempt to describe internal consensus implementation details beyond what is operationally verifiable.

2.2 Notation

For each parameter, this document provides:

Name (operator-facing)

Type

Required/Optional

Allowed values / constraints

Default behavior (if applicable)

Validation method (command or RPC)

Evidence to capture (for audit readiness)

Network Identity Parameters

3.1 EVM Chain ID

Name: EVM chainId
Type: uint256 (reported by JSON-RPC)
Required: Yes
Canonical value: 2121
Hex representation: 0x849
Constraints:

MUST be identical on all nodes of the same environment.

MUST remain constant across environments that claim NOORCHAIN 2.1 compatibility.

Validation method:

RPC: eth_chainId

Evidence to capture:

JSON-RPC response body showing 0x849

References:

docs/RPC_SPEC_2.1.md

docs/genesis/GENESIS_SPEC_2.1.md

3.2 Operator Chain Identifier (String)

Name: chain-id (operator string identifier)
Type: string
Required: Yes
Examples:

noorchain-2-1-local

noorchain-2-1-gdv

noorchain-2-1-pre-mainnet

noorchain-2-1-mainnet

Constraints:

MUST be unique per environment to prevent operator confusion.

MUST NOT conflict with EVM chainId (numeric).

MUST be identical across nodes in the same environment.

Validation method:

Node startup evidence (command line)

Optional: node logs where chain-id is printed (if implemented)

Evidence to capture:

exact node start command including -chain-id ...

References:

docs/genesis/GENESIS_CHECKLIST_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

Genesis Determinism Parameters

4.1 Genesis Timestamp Policy

Name: genesis timestamp policy
Type: policy (fixed or operator-chosen)
Required: Yes (policy must be declared)
Allowed values:

fixed (pre-agreed timestamp recorded in intent)

operator-chosen (timestamp recorded at generation time)

Constraints:

Policy MUST be recorded in the genesis intent record.

If operator-chosen, the chosen value MUST be recorded and distributed with the evidence pack.

Validation method:

Compare genesis intent record to the genesis artifact header.

Verify all nodes use the identical genesis artifact hash.

Evidence to capture:

GENESIS_INTENT.md timestamp policy line

genesis artifact digest (sha256)

References:

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

4.2 Genesis Artifact Digest

Name: genesis digest (sha256)
Type: hex string
Required: Yes
Constraints:

MUST be computed from canonical genesis bytes.

MUST match across operators and nodes.

Validation method:

sha256sum of the canonical genesis artifact

Evidence to capture:

sha256 output line

distribution channel record (out-of-band acknowledgement if used)

References:

docs/genesis/GENESIS_SPEC_2.1.md

docs/AUDIT_READINESS_2.1.md

Node Runtime Parameters (Core)

This section lists the operational flags/parameters used when starting nodes in controlled deployments.

5.1 Data Directory

Name: data-dir
Type: filesystem path
Required: Yes
Constraints:

Each node MUST have a unique data directory.

Two processes MUST NOT share the same data directory (LevelDB lock risk).

Validation method:

OS: confirm process args (pgrep / ps)

OS: ensure LOCK files are not contested

Evidence to capture:

pgrep -a noorcore output showing -data-dir ...

References:

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

5.2 P2P Listen Address

Name: p2p-addr
Type: ip:port
Required: Yes (multi-node), Optional (single-node)
Constraints:

MUST be unique per node on the same host.

In controlled VPS setups, SHOULD bind to the intended interface (private IP) and not to public wildcard unless explicitly permitted.

Validation method:

OS: ss -lntp | grep <port>

Node logs: peer sessions established (if multi-node)

Evidence to capture:

ss -lntp line for the p2p port

optional: evidence of ESTAB sessions (ss/conntrack)

References:

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/ops/INCIDENTS_2.1.md

5.3 RPC Listen Address

Name: rpc-addr
Type: ip:port
Required: Yes
Constraints:

In controlled deployments, RPC SHOULD bind to localhost unless a private exposure model is explicitly defined.

If bound beyond localhost, firewall policy MUST be documented.

Validation method:

OS: ss -lntp | grep <rpc-port>

RPC: eth_chainId, eth_blockNumber

Evidence to capture:

ss -lntp line for the rpc port

curl JSON-RPC outputs for eth_chainId/eth_blockNumber

References:

docs/RPC_SPEC_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

5.4 Health Endpoint Address

Name: health-addr
Type: ip:port
Required: Optional (recommended for ops)
Constraints:

MUST be unique per node on the same host.

SHOULD bind to localhost in controlled deployments.

Validation method:

OS: ss -lntp | grep <health-port>

HTTP: GET health endpoint returns success (per implementation)

Evidence to capture:

curl output to health endpoint

References:

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/ops/INCIDENTS_2.1.md

5.5 Boot Peers

Name: boot-peers
Type: list of p2p endpoints
Required: Multi-node bootstrap only
Constraints:

Leader/follower or multi-node packs MUST define bootstraps deterministically.

Boot peers MUST be reachable under the exposure policy.

Validation method:

Node logs or ss showing established sessions

Multi-node parity gates (RPC)

Evidence to capture:

node start command showing -boot-peers ...

ss evidence of established P2P

References:

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

Follower Routing Parameters (RPC Parity)

6.1 Follow-RPC

Name: follow-rpc
Type: URL or host:port (implementation-defined)
Required: Yes for follower mode in multi-node packs
Constraints:

Follower must route leader-only reads to leader when FollowRPC is set.

FollowRPC must be an explicit value (do not infer from “role” alone).

Validation method:

Compare leader and follower outputs for:

eth_chainId

eth_blockNumber

eth_getBlockByNumber("latest", false)

defined world-state reads in the parity matrix

Evidence to capture:

both leader and follower curl outputs side-by-side

References:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/RPC_SPEC_2.1.md

State Model Parameters (World-State)

7.1 World-State Backend

Name: world-state backend
Type: implementation detail with operational implications
Required: Yes (must be stable for the environment)
Canonical posture:

Ethereum-compatible persistent world-state with deterministic stateRoot per block.

Constraints:

stateRoot MUST be non-zero (except genesis where applicable) and stable across restarts.

world-state reads MUST be consistent across leader/follower within parity scope.

Validation method:

RPC: eth_getBlockByNumber("latest", false) includes stateRoot

RPC reads:

eth_getBalance

eth_getTransactionCount

eth_getCode (where supported)

eth_getStorageAt (where supported)

Evidence to capture:

outputs showing roots and reads

restart proof: same read after node restart (within expected chain progression)

References:

docs/STATE_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

Allocation Parameters (Alloc Overlay)

8.1 Alloc File Path

Name: alloc-file
Type: filesystem path
Required: Optional (policy-bound)
Constraints:

Must not contain secrets.

Must match declared chainId and environment policy.

Must be recorded in evidence pack.

Validation method:

Node logs show alloc application

RPC eth_getBalance verifies expected funding

Evidence to capture:

node logs line indicating alloc applied

eth_getBalance result

References:

docs/genesis/ALLOC_POLICY_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

8.2 Alloc Apply-Once Guard

Name: alloc applied marker
Type: internal marker (persistent)
Required: Yes if alloc overlay is used
Constraint:

Alloc overlay MUST be idempotent per data-dir to prevent repeated minting.

Validation method:

restart node; confirm alloc is not applied twice

verify balance remains coherent (no double credit)

Evidence to capture:

first startup log: alloc applied

second startup log: alloc skipped / already applied (if implemented)

References:

docs/genesis/ALLOC_POLICY_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

RPC Exposure and Privacy Parameters

9.1 Exposure Posture

Name: RPC exposure posture
Type: policy
Required: Yes
Allowed values:

localhost-only (default for controlled deployments)

private network only (explicitly documented)

public exposure (discouraged; requires explicit approval and compensating controls)

Constraints:

Public exposure MUST be accompanied by firewall and monitoring controls and documented in deployment model.

Validation method:

OS: verify bind address (127.0.0.1 vs 0.0.0.0)

Network: port scans from non-local context (where applicable)

Evidence to capture:

ss -lntp lines showing bind addresses

firewall rule evidence (if not localhost-only)

References:

docs/PRIVACY_DATA_POLICY_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

docs/ops/INCIDENTS_2.1.md

Parameter Recording Requirements (Audit Readiness)

For each environment, operators MUST maintain a parameter record including:

chain-id string identifier

EVM chainId and hex

ports (p2p, rpc, health) with bind addresses

validator set reference (if applicable)

alloc overlay usage (yes/no) and file digest (if yes)

genesis artifact digest

Evidence to capture:

a complete evidence pack as defined in GENESIS_CHECKLIST_2.1

References:

docs/genesis/GENESIS_CHECKLIST_2.1.md

docs/AUDIT_READINESS_2.1.md

docs/RELEASE_PROCESS_2.1.md

References

docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/genesis/ALLOC_POLICY_2.1.md
docs/STATE_MODEL_2.1.md
docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md