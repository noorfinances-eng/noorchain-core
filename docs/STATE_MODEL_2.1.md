# NOORCHAIN 2.1 — State Model (World-State, Roots, Persistence)

Version: v2.1-spec.0  
Status: Operator-first specification (normative)  
Last updated: 2026-01-19 (Europe/Zurich)

This document defines the **persistent state model** of NOORCHAIN 2.1 (evm-l1): Ethereum-compatible world-state storage, block roots, auxiliary indices, and the key/value layout required for deterministic operations and RPC correctness.

Normative keywords **MUST**, **SHOULD**, **MAY** are to be interpreted as described in RFC 2119.

---

## 0. Scope

### 0.1 In scope
- Ethereum-compatible **world-state**: accounts, balances, nonces, code, storage.
- State commitment: `stateRoot` per block, head tracking, and restart behavior.
- Auxiliary persistence required by tooling: receipts and block metadata (`receiptsRoot`, `logsBloom`).
- Storage layout: database locations and stable key prefixes.

### 0.2 Out of scope
- Consensus protocol details and validator operations (see runbooks).
- Full historical state access guarantees by arbitrary block number (beyond what is explicitly documented here).
- Public exposure hardening and authentication (recommended patterns are documented in ops docs).

---

## 1. Terminology

- **World-state**: Ethereum account state (nonce, balance, code, storage) persisted via trie-backed storage.
- **State root (`stateRoot`)**: Merkle-Patricia trie root hash of the committed world-state at a specific block.
- **Receipts root (`receiptsRoot`)**: trie root derived from the block’s transaction receipts per Ethereum rules.
- **Logs bloom (`logsBloom`)**: bloom filter derived from all receipt logs in a block.
- **NOOR DB**: node’s primary LevelDB store under the node `-data-dir` (operator-visible indices and metadata).
- **Geth DB**: dedicated LevelDB used for the Ethereum state backend (trie nodes, code, storage).
- **Leader**: canonical block producer for the active head.
- **Follower**: non-producing node configured to follow and (optionally) proxy leader-only reads.

---

## 2. Persistence Overview (Two-Store Model)

NOORCHAIN 2.1 persists data across two distinct stores:

1) **NOOR DB (LevelDB)**  
   Operator-oriented metadata and auxiliary indices required by RPC and runbooks.

2) **Geth DB (LevelDB)**  
   Ethereum-compatible world-state backend (trie nodes, contract code, storage).

This separation is intentional:
- the world-state uses canonical go-ethereum storage semantics,
- operational indices remain explicit, inspectable, and stable in NOOR DB.

---

## 3. Database Locations

Given a node `-data-dir <dataDir>`:

### 3.1 NOOR DB (LevelDB)
- Path: `<dataDir>/db/leveldb`
- Purpose: head pointers, per-height metadata, receipts index, PoSS application objects, bootstrap markers.

### 3.2 Geth DB (LevelDB)
- Path: `<dataDir>/db/geth`
- Purpose: Ethereum world-state backend (StateDB / trie / code / storage).

### 3.3 Locking invariant (normative)
A `<dataDir>` MUST NOT be used by more than one running node process at the same time (LevelDB locking).  
If LevelDB reports “resource temporarily unavailable”, operators MUST stop all processes using the same data-dir.

---

## 4. World-State Model (Ethereum-Compatible)

### 4.1 Accounts
Accounts follow Ethereum semantics:
- `nonce` (transaction count)
- `balance` (Wei)
- `codeHash` (empty if no code)
- `storageRoot` (empty if no storage)

### 4.2 Contract code
Contract bytecode is stored in the world-state backend.  
If `eth_getCode` is supported in the active build, it MUST return the exact deployed bytecode for `"latest"`.

### 4.3 Contract storage
Storage is a mapping of 32-byte slots to 32-byte values, committed via each account’s storage trie.  
If `eth_getStorageAt` is supported in the active build, it MUST return a 32-byte hex value for `"latest"`, defaulting to zero when unset.

### 4.4 No implicit consensus coupling
World-state contents MUST NOT change consensus validity rules beyond what is required to reproduce canonical block results.  
Application-layer objects (e.g., PoSS snapshots) MUST NOT be interpreted as consensus-critical state.

---

## 5. State Commitment and Roots (Normative)

### 5.1 Commit point
For each finalized block height `H`, the node MUST compute and persist:
- `stateRoot(H)` representing the world-state AFTER applying all block transactions,
- `receiptsRoot(H)` derived from receipts of all block transactions,
- `logsBloom(H)` derived from all receipt logs in the block.

The commit MUST be deterministic given the same chain inputs.

### 5.2 `stateRoot` semantics
- `stateRoot(H)` MUST be a 32-byte hash (`0x`-prefixed hex).
- `stateRoot` MUST be persisted as part of the block metadata for height `H`.
- `"latest"` RPC state reads MUST reference the committed head root, not transient in-memory state.

### 5.3 `receiptsRoot` and `logsBloom`
- `receiptsRoot(H)` MUST match the receipts list that the node can serve via `eth_getTransactionReceipt`.
- `logsBloom(H)` MUST reflect the same logs that appear in receipts.

If logs are not supported beyond receipts (e.g., limited `eth_getLogs`), `logsBloom` MUST still be computed from receipt logs to preserve internal consistency.

---

## 6. Canonical Head Root Tracking (NOOR DB)

The canonical head root is tracked explicitly to anchor `"latest"` reads.

### 6.1 Head state root key (required)
- Key: `stateroot/v1/head`
- Value: ASCII hex string, `0x`-prefixed 32-byte hash

### 6.2 Head root invariants
- After at least one committed block, `stateroot/v1/head` MUST exist.
- On each new finalized block, `stateroot/v1/head` MUST be updated to `stateRoot(H)`.
- After restart, the node MUST load `stateroot/v1/head` and use it as the anchor for `"latest"` state reads.

---

## 7. Block Metadata Index (NOOR DB)

Block metadata provides stable RPC mapping for headers/roots.

### 7.1 Key format (required)
- Key: `blkmeta/v1/<heightHexNo0x>`

Where `<heightHexNo0x>` is the block height encoded as lowercase hex WITHOUT the `0x` prefix.

Examples:
- height 1 → `blkmeta/v1/1`
- height 16 → `blkmeta/v1/10`

### 7.2 Value format (required)
The value MUST be UTF-8 JSON, designed for stable RPC mapping:

Minimum required fields:
- `height` (integer)
- `blockHash` (`0x…`)
- `stateRoot` (`0x…`)
- `receiptsRoot` (`0x…`)
- `logsBloomHex` (`0x…`, 256-byte bloom as hex)

### 7.3 RPC mapping invariants
- `eth_getBlockByNumber(<height>, …)` MUST return `null` if `<height>` is greater than the current head.
- For `"latest"` and by-number reads, `stateRoot`, `receiptsRoot`, and `logsBloom` MUST be populated from `blkmeta/v1/...` when available.

---

## 8. Receipts Persistence (NOOR DB)

Receipts MUST be persisted and retrievable by tx hash to support standard tooling.

### 8.1 Key format (required)
- Key: `rcpt/v1/<txHashHex>`

Where `<txHashHex>` is the `0x`-prefixed transaction hash in hex.

### 8.2 Value format (required properties)
Receipt encoding is implementation-defined, but the persisted receipt MUST preserve (at minimum):
- `transactionHash`
- `blockHash`
- `blockNumber`
- `from`
- `to`
- `contractAddress` (non-null for CREATE transactions)
- `status`
- `logs` (topics + data)

### 8.3 Restart invariant
A receipt that exists before shutdown MUST still be retrievable after restart from the same `-data-dir`.

---

## 9. PoSS Persistence (Application-Layer, NOOR DB)

PoSS (Proof of Signal Social) is an **application-layer** mechanism for value and governance. It is **not** part of consensus and MUST NOT affect consensus safety.

### 9.1 Storage separation (required)
PoSS objects MUST be stored in **NOOR DB** (not inside the world-state trie). This ensures:
- explicit auditability,
- independence from EVM execution details,
- operator control (backup/restore/inspection).

### 9.2 Snapshot object (canonical fields)
A PoSS “snapshot” is a persisted object identified by an incrementing `snapshotId`. A canonical snapshot MUST include, at minimum:
- `snapshotId` (uint64)
- `publisher` (address)
- `timestamp` (uint64)
- `payloadHash` (bytes32) — commitment to an off-chain evidence pack or payload
- `curatorSignatures` (list) — signatures over a deterministic digest (see §9.4)

Implementations MAY store additional metadata (campaign ids, pointers, labels), but the canonical fields above MUST remain stable.

### 9.3 Key-space (required)
PoSS persistence MUST use a versioned namespace and MUST NOT reuse prefixes from other subsystems.

Required keys:
- `poss/v1/head` → latest snapshot id (encoding MUST be stable and documented; recommended: ASCII hex quantity)
- `poss/v1/snap/<id>` → snapshot object (encoding MUST be stable)

Optional indices (if enabled):
- `poss/v1/byPublisher/<address>/<id>` → marker/index entry

### 9.4 Signature digest (required)
Curator signatures MUST be computed over a deterministic digest to prevent replay and ambiguity.

Required digest shape:
- `digest = keccak256( domainSeparator || chainId || registryAddress || snapshotId || payloadHash )`

Where:
- `domainSeparator` is a fixed constant for PoSS snapshots (versioned if needed),
- `chainId` prevents cross-chain replay,
- `registryAddress` binds signatures to the active PoSSRegistry instance.

The exact ABI packing MUST be documented alongside the PoSS registry interface and MUST NOT change without a version bump.

### 9.5 RPC exposure invariant
If PoSS RPC methods are present in the active build, their reads MUST be served from persisted PoSS objects and MUST remain stable across restarts.

### 9.6 Non-consensus invariants (required)
- PoSS persistence MUST NOT influence block validity rules.
- PoSS persistence MUST be safe to replay during resync without changing consensus outcomes.
- Economic policy (e.g., reward split 70% participants / 30% curators) is a project invariant but is not consensus logic.

---

## 10. Bootstrap Allocation Marker (NOOR DB)

Some deployments apply a one-time allocation (dev bootstrap) at node start.

### 10.1 Marker key (optional but stable if used)
- Key: `alloc/v1/applied`
- Value: implementation-defined marker (presence indicates allocation already applied for this data-dir)

### 10.2 Operational semantics
- Allocation MUST be applied at most once per `-data-dir` unless the marker is removed deliberately by an operator.
- Allocation application MUST update the committed head root (because balances/nonces are part of world-state).

---

## 11. Leader / Follower Semantics for World-State Reads (Normative)

World-state reads are treated as **Leader-only** unless explicitly proven safe to serve locally.

### 11.1 Authoritative routing rule
If a node is configured with `FollowRPC` (follower mode):
- leader-only state methods MUST be proxied to the leader,
- the follower MUST NOT claim local authority for head world-state reads.

This rule prevents transient or partial parity failures.

### 11.2 Leader-only set (minimum)
The following are leader-only by default:
- `eth_getBalance`
- `eth_getTransactionCount`
- `eth_getCode` (if supported)
- `eth_getStorageAt` (if supported)
- `eth_call`
- `eth_estimateGas`
- `eth_getLogs` (recommended leader-only in ops mode)

The authoritative list is defined in `docs/RPC_SPEC_2.1.md`.

---

## 12. RPC Mapping Summary (State Reads)

When supported in the active build, the following MUST reflect the committed head state rooted at `stateroot/v1/head`:

- `eth_getBalance(address, "latest")` → balance at head
- `eth_getTransactionCount(address, "latest")` → nonce at head
- `eth_getCode(address, "latest")` → code at head
- `eth_getStorageAt(address, slot, "latest")` → storage slot at head
- `eth_call(callObject, "latest")` → execution against head (read-only)

If a method is present but constrained (e.g., `"pending"` treated as `"latest"`), that constraint MUST be documented in `docs/RPC_SPEC_2.1.md`.

---

## 13. Operational Validation Gates (Normative)

A compliant NOORCHAIN 2.1 node MUST satisfy:

1) **Persistence across restarts**  
   After restart with the same `-data-dir`, `"latest"` state reads MUST remain consistent (unless new blocks were produced).

2) **Head root consistency**  
   `stateroot/v1/head` MUST match `eth_getBlockByNumber("latest", false).stateRoot` when both are available.

3) **Receipt durability**  
   A mined tx MUST have a retrievable receipt after restart.

4) **Follower correctness via proxy**  
   With `FollowRPC` configured, follower leader-only reads MUST match leader results.

---

## 14. Change Control

Any change affecting:
- world-state commitment,
- root derivation/exposure,
- key prefixes or encoding,

MUST update this document in the same change-set and be validated using the active runbooks.

