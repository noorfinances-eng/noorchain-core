NOORCHAIN 2.1 — Deployment Model

Document ID: DEPLOYMENT_MODEL_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines the deployment model for NOORCHAIN 2.1 controlled and permissioned environments, including node roles, network topology, exposure posture, ports, operational boundaries, and validation gates.

Purpose

This document defines the canonical deployment model for NOORCHAIN 2.1.

It is designed to:

standardize controlled deployment topologies

prevent exposure drift (especially RPC exposure)

define node roles and expected behavior (leader/follower)

specify port and binding posture

define minimum validation gates for each environment

support audit readiness by making deployment assumptions explicit

This document is operational and is read together with the operations playbook and RPC compatibility matrix.

References:

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

Deployment Principles

2.1 Controlled Exposure by Default

Default posture:

RPC binds to localhost unless explicitly approved otherwise

health endpoints bind to localhost

P2P exposure is minimized and restricted (private networking preferred)

2.2 Determinism and Reproducibility

deployment configurations must be versioned

node start commands must be captured

environment parameters must be recorded (chain identity, ports, routing)

Reference:

docs/AUDIT_READINESS_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

2.3 Separation of Concerns (Operational)

leader node: authoritative execution and state writes

follower node(s): read-only RPC service and redundancy (where configured)

validator set governance remains permissioned and controlled

Reference:

docs/governance/GOVERNANCE_MODEL_2.1.md

Environment Profiles

This document defines three canonical profiles. Specific deployments must state which profile they follow.

3.1 Local Profile (Developer Controlled)

Characteristics:

single host (local workstation or codespace)

leader node only, or leader + follower for parity testing

RPC bound to 127.0.0.1

P2P bound to 127.0.0.1 for multi-node local testing

Primary goals:

functional validation

tooling compatibility

reproducible runbook execution

3.2 Pilot Profile (Controlled VPS / Partner Pilot)

Characteristics:

private VPS (or controlled cloud instances)

leader + follower recommended

RPC bound to localhost, accessed via SSH tunnels

P2P restricted to private networking or explicit allowlist

strict logging and evidence pack requirements

Primary goals:

controlled proof-of-liveness

audit-ready evidence generation

minimal exposure posture compatible with Legal-Light

References:

docs/legal/LEGAL_LIGHT_POSTURE_2.1.md

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

3.3 Permissioned Network Profile (Pre-Mainnet / Mainnet Permissioned)

Characteristics:

multi-validator, permissioned BFT

coordinated upgrade windows

explicit validator identity and membership governance

RPC remains restricted (localhost/private) unless explicitly authorized

stronger monitoring and incident response controls

Primary goals:

stable permissioned network operations

deterministic releases and upgrades

controlled governance and change management

References:

docs/governance/UPGRADE_PROCESS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

Node Roles

4.1 Leader Node

Definition:

the leader is the authoritative node for execution and world-state persistence in a leader/follower setup.

Responsibilities:

accepts transactions (where enabled)

produces blocks (where enabled)

persists receipts and world-state

exposes JSON-RPC (preferably localhost only)

Operational constraints:

leader data-dir must be unique

leader RPC is the authoritative source for routed methods

4.2 Follower Node (Read-Only / Routed Reads)

Definition:

a follower is a node configured to serve read-only RPC while routing leader-only methods to the leader via FollowRPC.

Responsibilities:

serve read-only JSON-RPC

proxy leader-only world-state reads and other routed methods

provide redundancy and parity validation surface

Operational constraints:

follower must have a distinct data-dir

follower must be started with a FollowRPC reference

follower must not silently diverge in world-state reads

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

Topologies

5.1 Single-Node Topology (Local)

1 leader node

RPC: localhost

P2P: optional, usually not exposed

When used:

early functional tests

single-node smoke validation of genesis and RPC

Minimum gates:

eth_chainId

eth_blockNumber

eth_getBlockByNumber("latest", false)

5.2 Two-Node Topology (Leader + Follower)

leader node: authoritative

follower node: read-only + routed reads

P2P session established between nodes

When used:

parity testing

controlled pilot deployments

operational redundancy (read-side)

Minimum gates:

leader and follower eth_chainId match

leader and follower eth_blockNumber coherent

parity for required methods per compat matrix

P2P sessions established

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

5.3 Multi-Node Permissioned Validators

multiple validator nodes

explicit validator set (permissioned)

coordinated changes and upgrades

When used:

pre-mainnet and mainnet permissioned phases

Minimum gates:

validator membership is consistent

network produces blocks deterministically under operational constraints

upgrade process gates are executed under governance approval

References:

docs/governance/UPGRADE_PROCESS_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

Ports and Binding Posture

This section defines canonical port categories. Exact port numbers may vary by environment, but binding rules are strict.

6.1 Port Categories

P2P: peer networking sessions

RPC: JSON-RPC interface

Health: lightweight HTTP endpoint for liveness/health (if enabled)

6.2 Binding Rules (Default)

Default posture for controlled deployments:

RPC binds to 127.0.0.1

Health binds to 127.0.0.1

P2P binds to:

127.0.0.1 for local multi-node testing, or

private interface for VPS environments (not public wildcard)

Public wildcard binds (0.0.0.0) are discouraged and require explicit approval and compensating controls.

Reference:

docs/PRIVACY_DATA_POLICY_2.1.md

docs/governance/GOVERNANCE_MODEL_2.1.md

6.3 Exposure Policy and Access Patterns

Preferred access patterns:

local: direct access to localhost ports

VPS pilot: RPC accessed via SSH tunnel; no public RPC

permissioned network: private networking for validator communications; RPC restricted

Configuration and Parameter Recording

For each deployed environment, operators MUST record:

chain-id string identifier

EVM chainId (2121 / 0x849)

per-node:

data-dir

p2p-addr

rpc-addr

health-addr (if enabled)

follow-rpc (for follower nodes)

boot peers list (if used)

alloc overlay usage (yes/no, file digest)

genesis artifact digest (sha256) where applicable

Reference:

docs/genesis/PARAMETERS_REFERENCE_2.1.md

docs/AUDIT_READINESS_2.1.md

Minimum Operational Gates (Deployment Validation)

8.1 Single Node Gates

process starts without error

RPC reachable

eth_chainId == 0x849

eth_blockNumber returns a valid quantity

eth_getBlockByNumber("latest", false) returns coherent structure

Reference:

docs/RPC_SPEC_2.1.md

8.2 Two Node Gates (Leader/Follower)

All single-node gates plus:

no port collisions (distinct binds)

P2P sessions established

follower routing correct:

methods designated leader-only must route to leader

parity gates per matrix pass

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

8.3 Permissioned Network Gates

In addition to the above:

validator membership consistent with permissioned governance artifacts

upgrade windows and release process enforced

incident response readiness confirmed

References:

docs/governance/UPGRADE_PROCESS_2.1.md

docs/governance/INCIDENT_RESPONSE_2.1.md

Operational Boundaries and Forbidden Patterns

The following patterns are forbidden in controlled deployments:

two nodes sharing the same data-dir

public RPC exposure without explicit approval and compensating controls

follower nodes serving world-state reads without routing when required

ad-hoc configuration changes without evidence and versioning

silent genesis mismatch across nodes

References:

docs/genesis/GENESIS_SPEC_2.1.md

docs/PRIVACY_DATA_POLICY_2.1.md

docs/AUDIT_READINESS_2.1.md

Incident Linkage

Deployment model violations are often incidents:

port collisions

data-dir lock conflicts

parity gate failures

unexpected public exposure

Response should follow incident response policy.

Reference:

docs/governance/INCIDENT_RESPONSE_2.1.md

References

docs/OPERATIONS_PLAYBOOK_2.1.md
docs/genesis/PARAMETERS_REFERENCE_2.1.md
docs/genesis/GENESIS_SPEC_2.1.md
docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/governance/GOVERNANCE_MODEL_2.1.md
docs/governance/UPGRADE_PROCESS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md
docs/legal/LEGAL_LIGHT_POSTURE_2.1.md
docs/compliance/COMMUNICATIONS_POLICY_2.1.md