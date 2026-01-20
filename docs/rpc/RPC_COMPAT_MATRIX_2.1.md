NOORCHAIN 2.1 — RPC Compatibility Matrix

Document ID: RPC_COMPAT_MATRIX_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines the JSON-RPC compatibility surface for NOORCHAIN 2.1, including method support level, behavioral constraints, follower routing rules, and validation gates for tooling/wallet compatibility in controlled deployments.

Purpose

This document defines the RPC compatibility matrix for NOORCHAIN 2.1.

It is designed to:

specify which JSON-RPC methods are supported and at what level

define behavioral expectations (inputs, outputs, error behavior)

define leader/follower routing rules for specific methods

provide validation gates for tooling (Hardhat, Viem, wallets)

prevent compatibility drift across releases

This matrix is normative for operational compatibility claims.

Scope and Definitions

2.1 Node Roles

Leader: authoritative state and execution source.

Follower: read-only RPC server that may proxy leader-only methods via FollowRPC.

Reference:

docs/ops/DEPLOYMENT_MODEL_2.1.md

2.2 Support Levels

Each method is assigned a support level:

L0 (Unsupported): method not implemented; must return a stable error per error model.

L1 (Stub/Minimal): method returns structurally valid output but may be incomplete or restricted.

L2 (Operational): method works for controlled deployments and tooling use cases; behavior is stable and documented.

L3 (Parity Target): method is expected to behave Ethereum-compatibly across leader/follower within defined routing semantics.

Support level may differ between leader and follower depending on routing rules.

2.3 Routing Classes

Local: executed on the node receiving the request.

Leader-only: must be served from leader (either called directly or proxied by follower).

Proxy-safe: follower may proxy to leader with identical semantics.

Follower-local: safe to serve locally on follower without proxy.

Compatibility Goals

3.1 Controlled Wallet/Tooling Compatibility

Primary compatibility targets:

Hardhat / ethers / Viem basic flows

transaction submission, receipt retrieval

contract deploy and basic reads

block metadata reads for explorers and evidence packs

3.2 Evidence and Audit Readiness

Compatibility includes:

stable method behavior across restarts

deterministic outputs for identical inputs within a given block context

coherent block header roots exposure (stateRoot, receiptsRoot, logsBloom)

Reference:

docs/AUDIT_READINESS_2.1.md

docs/STATE_MODEL_2.1.md

Method Matrix

The following table defines the current compatibility surface. If a method is not listed, treat it as L0 by default.

For each method:

Leader Support Level

Follower Support Level

Routing Rule

Notes / Constraints

Validation Gate

4.1 Identity and Network

Method: web3_clientVersion
Leader: L1
Follower: L1
Routing: Local
Notes: informational only; must not leak sensitive details.
Gate: returns non-empty string.

Method: net_version (if supported)
Leader: L0/L1 (implementation-defined)
Follower: L0/L1
Routing: Local
Notes: may be omitted; prefer eth_chainId for identity.
Gate: stable error or valid result.

Method: eth_chainId
Leader: L3
Follower: L3
Routing: Follower-local
Notes: must return 0x849 (2121) for NOORCHAIN 2.1 environments.
Gate: leader==follower==0x849.

Method: eth_protocolVersion (if supported)
Leader: L0/L1
Follower: L0/L1
Routing: Local
Notes: optional; stable error acceptable.
Gate: stable behavior.

4.2 Block and Chain Metadata

Method: eth_blockNumber
Leader: L3
Follower: L3
Routing: Follower-local
Notes: must be coherent; follower may track independently but must not regress.
Gate: follower value is coherent with leader (equal or within defined follow semantics).

Method: eth_getBlockByNumber
Leader: L2/L3
Follower: L2/L3
Routing: Proxy-safe
Notes: must return null when requested block > latest; must include stateRoot/receiptsRoot/logsBloom when available.
Gate: leader and follower return coherent block metadata for "latest".

Method: eth_getBlockByHash
Leader: L1/L2
Follower: L1/L2
Routing: Proxy-safe
Notes: optional for some tooling; stable error acceptable if not required by target tools.
Gate: stable behavior.

Method: eth_getBlockTransactionCountByNumber
Leader: L0/L1
Follower: L0/L1
Routing: Proxy-safe
Notes: optional; stable error acceptable.
Gate: stable behavior.

Method: eth_getBlockTransactionCountByHash
Leader: L0/L1
Follower: L0/L1
Routing: Proxy-safe
Notes: optional; stable error acceptable.
Gate: stable behavior.

4.3 Accounts and World-State Reads

Method: eth_getBalance
Leader: L2/L3
Follower: L2/L3
Routing: Leader-only (proxy by FollowRPC)
Notes: follower must route to leader when FollowRPC is configured; do not serve stale local state.
Gate: leader==follower for a known funded address.

Method: eth_getTransactionCount
Leader: L2/L3
Follower: L2/L3
Routing: Leader-only (proxy by FollowRPC)
Notes: nonce must be consistent with mined txs; required for contract deploy flows.
Gate: nonce increments after a mined tx.

Method: eth_getCode
Leader: L0/L2 (depends on implementation)
Follower: L0/L2
Routing: Leader-only (proxy if supported)
Notes: required for tooling that checks deployed bytecode; if unsupported, must return stable error.
Gate: for a deployed contract address, returns non-empty 0x... code.

Method: eth_getStorageAt
Leader: L0/L2
Follower: L0/L2
Routing: Leader-only (proxy if supported)
Notes: required for some contract read flows; if unsupported, stable error.
Gate: for known storage slot, returns expected 32-byte value.

Method: eth_call
Leader: L1/L2 (implementation-defined)
Follower: L1/L2
Routing: Leader-only (proxy recommended)
Notes: if implemented, must execute against specified block tag semantics. If not, stable error or bounded stub must be documented.
Gate: known contract view call returns expected value.

4.4 Transactions and Receipts

Method: eth_sendRawTransaction
Leader: L2/L3
Follower: L0/L1 (recommended: reject or proxy by policy)
Routing: Leader-only
Notes: submission should be directed to leader; follower should not accept tx submission unless explicitly designed to proxy.
Gate: tx hash returned, and tx becomes retrievable.

Method: eth_getTransactionByHash
Leader: L2/L3
Follower: L2/L3
Routing: Proxy-safe
Notes: must return tx fields expected by tooling (from, to, nonce, input, etc.) as available.
Gate: returns non-null for known tx hash.

Method: eth_getTransactionReceipt
Leader: L2/L3
Follower: L2/L3
Routing: Proxy-safe
Notes: must return receipt after mining; contractAddress must be set for CREATE.
Gate: deploy tx receipt includes non-null contractAddress.

Method: eth_getTransactionByBlockNumberAndIndex
Leader: L0/L1
Follower: L0/L1
Routing: Proxy-safe
Notes: optional; stable error acceptable.
Gate: stable behavior.

Method: eth_getTransactionByBlockHashAndIndex
Leader: L0/L1
Follower: L0/L1
Routing: Proxy-safe
Notes: optional; stable error acceptable.
Gate: stable behavior.

4.5 Logs and Filtering

Method: eth_getLogs
Leader: L1/L2
Follower: L1/L2
Routing: Proxy-safe
Notes: if supported, must at minimum allow filtering by address and block range in controlled deployments; otherwise stable error.
Gate: logs for known tx/contract are retrievable.

Method: eth_newFilter / eth_getFilterChanges / eth_uninstallFilter
Leader: L0
Follower: L0
Routing: N/A
Notes: subscriptions/filters often deferred; must return stable error per model.
Gate: stable error.

Method: eth_subscribe (WS)
Leader: L0
Follower: L0
Routing: N/A
Notes: out of scope unless explicitly implemented.
Gate: stable error.

4.6 Gas and Fee Estimation

Method: eth_gasPrice
Leader: L1/L2
Follower: L1/L2
Routing: Follower-local
Notes: may be constant or policy-driven; must be stable and documented.
Gate: returns valid quantity.

Method: eth_estimateGas
Leader: L1/L2
Follower: L1/L2
Routing: Proxy-safe or leader-only depending on implementation
Notes: heuristic estimation may be used; must not return malformed responses.
Gate: returns valid quantity for basic tx.

Method: eth_feeHistory
Leader: L0/L1
Follower: L0/L1
Routing: Proxy-safe
Notes: optional; stable error acceptable.
Gate: stable behavior.

4.7 Client/Chain Helpers

Method: eth_getBlockReceipts (if supported)
Leader: L0/L1
Follower: L0/L1
Routing: Proxy-safe
Notes: optional.
Gate: stable behavior.

Method: eth_syncing
Leader: L1
Follower: L1
Routing: Local
Notes: may return false in controlled environments; must be stable.
Gate: stable result.

Error Model Requirements (Normative)

Unsupported methods (L0) MUST return a stable error structure consistent with the RPC error model document.

Reference:

docs/rpc/RPC_ERROR_MODEL_2.1.md

Validation Gate Set (Compatibility)

For a given environment, the minimum compatibility gate set depends on intended tool usage. The following gates are recommended for “wallet/tooling compatible” controlled deployments:

Gate A — Identity

eth_chainId == 0x849 (leader and follower)

Gate B — Liveness

eth_blockNumber returns valid quantity

eth_getBlockByNumber("latest", false) returns coherent structure

Gate C — Tx Path

eth_sendRawTransaction (leader) returns tx hash

eth_getTransactionByHash returns non-null

eth_getTransactionReceipt returns non-null after mining

Gate D — Deploy

CREATE tx receipt includes contractAddress

eth_getTransactionCount increments for sender

Gate E — World-State Read Parity

eth_getBalance and eth_getTransactionCount match on leader and follower (with follower routing)

Gate F — Contract Read (If eth_call supported)

eth_call against known view returns expected value

Gate G — Persistence

repeat Gate E and Gate C after node restart

Change Control

Any change to support level or behavior of a method is an RPC compatibility change and MUST include:

changelog entry

versioned documentation update

validation gate evidence for the affected methods

References:

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/AUDIT_READINESS_2.1.md

References

docs/RPC_SPEC_2.1.md
docs/rpc/RPC_ERROR_MODEL_2.1.md
docs/STATE_MODEL_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/CHANGELOG_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md