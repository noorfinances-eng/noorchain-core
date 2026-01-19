# NOORCHAIN 2.1 — Architecture (System View and Core Flows)

Version: v2.1-spec.0  
Status: Operator-first specification (normative where stated)  
Last updated: 2026-01-19 (Europe/Zurich)

This document provides the **system architecture** of NOORCHAIN 2.1 (evm-l1): components, responsibilities, runtime roles (leader/follower), and the end-to-end flows that define “mainnet-like” behavior in controlled environments.

Normative keywords **MUST**, **SHOULD**, **MAY** are to be interpreted as described in RFC 2119.

---

## 0. Scope

### 0.1 In scope
- High-level system composition (node, RPC, persistence, execution).
- Role model (Leader / Follower) and routing expectations.
- Primary flows: node boot, block production, tx path, state commitment, receipts/logs, PoSS persistence.
- Deployment topology patterns used in operator runbooks.

### 0.2 Out of scope
- Permissioning policy, validator lifecycle, and governance operations.
- Public network hardening and exposure strategy (covered by ops/security docs).
- Full Ethereum parity claim (subset-based compatibility; see RPC spec).

---

## 1. System Overview

NOORCHAIN 2.1 is a **sovereign EVM Layer-1** with:
- a **permissioned BFT** consensus model (details in runbooks / internal design notes),
- an Ethereum-style **JSON-RPC** interface for tooling compatibility,
- an Ethereum-compatible **world-state** backend (trie-based) persisted via a dedicated storage layer,
- operator-oriented persistence indices and metadata stored in a separate chain-level database.

PoSS (Proof of Signal Social) is an **application-layer** mechanism. It is explicitly **not** part of consensus and MUST NOT influence block validity.

---

## 2. Component Model

A single node process (“noorcore”) is the primary unit.

### 2.1 Node Runtime
Responsibilities:
- P2P networking (peer discovery / sessions) for multi-node operation.
- Block production loop (leader mode) or follow mode.
- Transaction ingestion and inclusion (ops-first behavior).
- Execution hooks for transaction processing (EVM path as implemented by active tags).
- Persistence and recovery on restart.

### 2.2 JSON-RPC Server
Responsibilities:
- Provide Ethereum-style JSON-RPC subset for tooling and wallet compatibility.
- Enforce leader/follower routing rules (proxy leader-only reads when configured).
- Serve receipts/logs and block/header information.

Source of truth:
- RPC behavior is specified in `docs/RPC_SPEC_2.1.md`.

### 2.3 Persistence Layer (Two-store design)
- **NOOR DB**: operator-friendly indices and metadata.
- **Geth DB**: Ethereum-compatible world-state backend.

Source of truth:
- Persistence schema is specified in `docs/STATE_MODEL_2.1.md`.

### 2.4 Application-layer PoSS
Responsibilities:
- Persist PoSS objects (snapshots) as audit-grade records.
- Provide read paths that survive restarts and remain stable.

Constraints:
- PoSS MUST NOT impact consensus safety.
- PoSS persistence MUST be separable from world-state.

---

## 3. Runtime Roles (Leader / Follower)

NOORCHAIN 2.1 defines two operational roles for multi-node packs.

### 3.1 Leader
The Leader is the canonical producer for the active head. It MUST:
- advance the chain head by producing blocks,
- commit world-state roots per block,
- persist receipts and block metadata,
- serve authoritative `"latest"` reads.

### 3.2 Follower
The Follower tracks the leader. It MUST:
- maintain network connectivity and follow head progression,
- serve RPC where safe,
- proxy leader-only reads when configured with `FollowRPC`.

### 3.3 FollowRPC routing rule (normative)
When the follower is configured with `FollowRPC`:
- the follower MUST proxy all **leader-only** methods to the leader,
- the follower MUST NOT claim local authority for `"latest"` world-state reads.

This is an ops-first safety rule to avoid partial parity issues.

---

## 4. Data Model (Architectural View)

This section summarizes how persisted state and indices relate.

### 4.1 World-state (Geth DB)
The Ethereum-compatible world-state includes:
- account nonce
- account balance
- contract code
- contract storage

It is committed per block, producing `stateRoot(H)`.

### 4.2 Chain indices (NOOR DB)
NOOR DB maintains:
- `stateroot/v1/head` (canonical head root pointer)
- `blkmeta/v1/<height>` (per-height metadata: blockHash, roots, bloom)
- `rcpt/v1/<txHash>` (receipt lookup index)
- optional PoSS namespaces (application objects)

The exact prefixes and constraints are defined in `docs/STATE_MODEL_2.1.md`.

---

## 5. Core Flows

### 5.1 Node Boot & Recovery

**Goal:** start a node deterministically using a given `-data-dir`.

Flow:
1) Node opens NOOR DB under `<data-dir>/db/leveldb`.
2) Node opens Geth DB under `<data-dir>/db/geth`.
3) Node loads head pointers (e.g., `stateroot/v1/head`) if present.
4) Node initializes runtime services:
   - P2P service
   - JSON-RPC service
   - health endpoint (if enabled by flags)
5) Node enters main loop:
   - leader mode: produce blocks
   - follower mode: follow leader + proxy rules

**Restart invariant (normative):**  
If the node is restarted with the same `-data-dir`, `"latest"` reads MUST match pre-restart results unless new blocks were produced.

---

### 5.2 Transaction Submission Path (Tooling-Compatible)

**Goal:** satisfy the standard Ethereum tooling loop:
`eth_sendRawTransaction → eth_getTransactionByHash → eth_getTransactionReceipt`

Flow:
1) Client signs a transaction locally.
2) Client submits raw tx via `eth_sendRawTransaction`.
3) Node accepts the raw tx and returns its hash.
4) Node includes the tx in a produced block (leader).
5) Node persists:
   - receipt (`rcpt/v1/<txHash>`)
   - block metadata (`blkmeta/v1/<height>`)
   - updated head root pointer (`stateroot/v1/head`)
6) Client polls:
   - `eth_getTransactionByHash` (optional)
   - `eth_getTransactionReceipt` until non-null

**Receipt durability (normative):**  
Once a tx is mined, the receipt MUST be retrievable after restart.

---

### 5.3 Block Production (Leader)

**Goal:** produce a new canonical block and persist all required artifacts.

High-level sequence:
1) Leader selects transactions for inclusion.
2) For each transaction:
   - decode
   - determine sender
   - apply execution hook (EVM execution as implemented by active tags)
   - produce receipt and logs
3) Finalize block:
   - commit world-state → `stateRoot(H)`
   - derive `receiptsRoot(H)`
   - derive `logsBloom(H)`
4) Persist:
   - NOOR DB: `blkmeta/v1/<H>` with roots/bloom
   - NOOR DB: update `stateroot/v1/head = stateRoot(H)`
   - NOOR DB: `rcpt/v1/<txHash>` entries
   - Geth DB: trie/state updates

**Consistency invariant (normative):**  
`eth_getBlockByNumber("latest").stateRoot` MUST match `stateroot/v1/head`.

---

### 5.4 World-State Reads (RPC “latest”)

**Goal:** ensure `"latest"` reads reflect the committed head state.

Authoritative anchor:
- `"latest"` MUST reference the head root stored in `stateroot/v1/head`.

Reads:
- `eth_getBalance(address,"latest")` → world-state at head root
- `eth_getTransactionCount(address,"latest")` → nonce at head root
- `eth_getCode(address,"latest")` → code at head root (if supported)
- `eth_getStorageAt(address,slot,"latest")` → storage at head root (if supported)
- `eth_call(…,"latest")` → execution against head state (read-only semantics)

**Leader-only rule:**  
These methods are leader-only unless explicitly marked otherwise in `docs/RPC_SPEC_2.1.md`.

---

### 5.5 Follower Proxy Flow (Leader-only reads)

**Goal:** provide consistent results across leader and follower in multi-node packs.

Flow:
1) Client calls follower RPC.
2) Follower checks if `FollowRPC` is configured.
3) If method is leader-only:
   - follower proxies the request to leader RPC
   - follower returns the leader response verbatim (JSON-RPC envelope preserved)

**Normative constraint:**  
Follower MUST NOT “best-effort” answer leader-only reads locally when `FollowRPC` is set.

---

### 5.6 Block Metadata Reads (Headers and Roots)

**Goal:** serve block header/root fields deterministically.

Flow:
1) RPC receives `eth_getBlockByNumber(heightOrTag, includeTx)`.
2) RPC resolves tag/height:
   - `"latest"` → current head height
   - if requested height > head → return `null`
3) RPC loads `blkmeta/v1/<height>` (NOOR DB).
4) RPC returns block object with:
   - `stateRoot`, `receiptsRoot`, `logsBloom` populated from blkmeta
   - other fields as supported by active build

---

### 5.7 Logs & Receipts Consistency

**Goal:** ensure `eth_getTransactionReceipt` and log retrieval represent the same ground truth.

Rules:
- Receipt logs returned by `eth_getTransactionReceipt` are canonical.
- `logsBloom` MUST reflect the union of receipt logs for the block.
- `eth_getLogs` (if supported) MUST not contradict receipt logs.

Where the implementation constrains `eth_getLogs` (bounded queries, limited filters), those constraints MUST be declared in `docs/RPC_SPEC_2.1.md`.

---

### 5.8 PoSS Snapshot Flow (Application-layer persistence)

**Goal:** persist PoSS snapshots as audit-grade objects without affecting consensus.

Flow (conceptual):
1) A client prepares an evidence pack off-chain (files, attestations, references).
2) Client computes `payloadHash` (canonical commitment).
3) Curators produce signatures over a deterministic digest (domain-separated).
4) Snapshot object is submitted through the supported path (contract call or dedicated handler as implemented).
5) Node persists snapshot object under the PoSS namespace in NOOR DB.
6) RPC methods (if enabled) read snapshots from persisted objects.

**Non-consensus invariant (normative):**
- PoSS objects MUST NOT affect block validity rules.
- PoSS persistence MUST be replay-safe and restart-stable.

The canonical fields, key-space, and digest constraints are defined in `docs/STATE_MODEL_2.1.md`.

---

## 6. Deployment Topologies (Ops-first)

### 6.1 Single-node local
Used for development and minimal validation.
- One node binds P2P/RPC/health to loopback.
- Tooling connects directly to the node RPC.

### 6.2 Two-node “mainnet-like” pack (Leader/Follower)
Used to validate:
- P2P connectivity and follower tracking,
- proxy correctness for leader-only reads,
- operational stability across restarts.

Normative behavior:
- follower is configured with `FollowRPC` pointing to leader RPC,
- leader-only reads on the follower are proxied to the leader.

The precise commands, port plans, and validation gates are runbook-defined.

---

## 7. Constraints and Non-Goals (Explicit)

- This architecture is **subset-based** for Ethereum RPC compatibility. It does not claim full parity with Ethereum mainnet.
- Account management via `personal_*` and unlocked node keys is not part of the node model; clients MUST sign locally.
- Websocket subscriptions (`eth_subscribe`) are not part of the baseline surface unless explicitly added and documented.
- Public RPC exposure is not supported without an external security gateway (TLS/auth/rate-limits).

---

## 8. Document Links and Source of Truth

- RPC method surface, routing rules, and constraints: `docs/RPC_SPEC_2.1.md`
- Persistence schema, roots, and key prefixes: `docs/STATE_MODEL_2.1.md`
- Operational start/stop and validation gates: runbooks under `docs/`

This architecture document is a system map. When a discrepancy exists, the method-level spec and runbooks are authoritative for operations.

---

## 9. Change Control

Any change affecting:
- leader/follower semantics,
- persistence structure (DB paths, key prefixes, encodings),
- method routing rules or `"latest"` anchoring,

MUST update the relevant core documents (`RPC_SPEC_2.1`, `STATE_MODEL_2.1`, this file) in the same change-set and be validated via the active runbooks.

