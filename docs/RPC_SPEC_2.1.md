# NOORCHAIN 2.1 — RPC Specification (JSON-RPC 2.0 / Ethereum Compatibility)

Version: v2.1-draft.0  
Status: Working Specification (ops-first)  
Last updated: 2026-01-19 (Europe/Zurich)

This document specifies the public JSON-RPC surface of NOORCHAIN 2.1 (evm-l1), with explicit compatibility targets, role rules (leader/follower), and known limits.

Normative keywords **MUST**, **SHOULD**, **MAY** are to be interpreted as described in RFC 2119.

---

## 0. Scope

### 0.1 In scope
- Transport, request/response rules, and batch behavior (JSON-RPC 2.0).
- Endpoint topology: **Leader** and **Follower** roles and the **FollowRPC** routing rule.
- Method-level semantics for the supported Ethereum-style RPC subset.
- Operational constraints and explicit limitations (what is intentionally not provided).

### 0.2 Out of scope
- Consensus, validator onboarding, and P2P topology details (covered in runbooks).
- Economic policy / token issuance (covered elsewhere).
- Public internet hardening / authentication layer design (recommended patterns only).

---

## 1. Definitions

- **Leader**: the node producing blocks and maintaining the authoritative head progression.
- **Follower**: a node following the leader’s head and serving read requests; may proxy a defined subset of calls to the leader.
- **FollowRPC**: follower configuration pointing to the leader RPC base URL (e.g., `http://127.0.0.1:8545`).  
  When set, it is the authoritative signal that the follower must proxy leader-only calls.

---

## 2. Transport

### 2.1 Protocol
- The RPC interface **MUST** accept **HTTP POST** requests with JSON bodies.
- The server **MUST** implement JSON-RPC 2.0 envelope fields:
  - `jsonrpc`: `"2.0"`
  - `id`: string|number|null
  - `method`: string
  - `params`: array|object (method-dependent)

### 2.2 Content types
- Clients **SHOULD** send: `Content-Type: application/json`.

### 2.3 Batch requests
- The server **MUST** accept JSON-RPC batch payloads (array of requests).
- Batch responses **MUST** be an array of per-request responses in arbitrary order.

### 2.4 Errors
- Unknown method: `-32601` (Method not found)
- Invalid JSON: `-32700` (Parse error)
- Invalid request: `-32600`
- Invalid params: `-32602`
- Internal error: `-32603`

Implementation-specific error messages **MAY** be returned in `error.message`, but clients must not rely on exact wording.

---

## 3. Endpoint Topology (Leader / Follower)

### 3.1 Addresses
RPC bind addresses are configured by node flags (examples below are conventional defaults used in local ops):

- Leader RPC: `-rpc-addr 127.0.0.1:8545`
- Follower RPC: `-rpc-addr 127.0.0.1:8546`
- Follower leader-link: `-follow-rpc http://127.0.0.1:8545`

The RPC service **SHOULD** bind to loopback by default. Exposing RPC publicly is not a supported default posture.

### 3.2 Leader-only routing rule
When `FollowRPC` is set on a follower:
- The follower **MUST** proxy all **Leader-only** methods to the leader.
- The follower **MUST NOT** attempt to serve leader-only methods from its own local state, even if it *appears* up to date.

Rationale: leader-only methods depend on canonical head world-state and/or canonical block metadata.

---

## 4. Compatibility Targets

The RPC surface is designed to support standard Ethereum tooling in a controlled environment:

- Wallet/tooling read primitives: chain identity, blocks, receipts, basic state reads.
- Transaction path: submit raw tx → obtain tx hash → poll receipt.
- Basic contract interactions via `eth_call` for read-only calls, within supported limits.

Compatibility is **subset-based**: no claim of full Ethereum mainnet parity is made in this document. Deviations are declared explicitly.

---

## 5. Chain Identity

### 5.1 `eth_chainId`
- **Supported**
- Returns the chain ID as hex quantity.

Expected NOORCHAIN 2.1 chain ID:
- Decimal: `2121`
- Hex: `0x849`

### 5.2 `net_version`
- **Supported (compat)**
- Returns the chain ID as a decimal string (e.g., `"2121"`).

### 5.3 `web3_clientVersion`
- **Supported (informational)**
- Returns an implementation string. Format is not stable and must not be parsed.

---

## 6. Block & Header Model

### 6.1 Quantity encoding
All numeric values returned as Ethereum quantities **MUST** be hex, `0x` prefixed, without leading zeros (except `0x0`).

### 6.2 `eth_blockNumber`
- **Supported**
- Returns the latest canonical block height (head).

### 6.3 `eth_getBlockByNumber`
- **Supported**
- Params:
  1) block tag (`"latest"`) or quantity (hex)
  2) boolean `true|false` to include full tx objects

Semantics:
- If requested height is greater than current head, **MUST** return `null`.
- Block object **MUST** include:
  - `number`, `hash`, `parentHash`, `timestamp`
  - `stateRoot`, `receiptsRoot`, `logsBloom` (when available)
- For `true` include tx objects; for `false` include tx hashes only.

Known constraint:
- `"pending"` is treated as `"latest"`.

---

## 7. Transaction Path

### 7.1 `eth_sendRawTransaction`
- **Supported**
- Accepts a raw signed transaction as hex bytes.
- Returns the transaction hash as `0x…`.

Constraints:
- The node does not expose txpool inspection endpoints (`txpool_*` not supported).
- Transaction ordering and inclusion are implementation-defined and tied to the node’s block production loop.

### 7.2 `eth_getTransactionByHash`
- **Supported**
- Returns `null` if unknown.
- When known, returns a tx object sufficient for tooling correlation (hash, from, to, nonce, input, etc.).

### 7.3 `eth_getTransactionReceipt`
- **Supported**
- Returns `null` until mined.
- Receipt object includes:
  - `transactionHash`, `blockHash`, `blockNumber`
  - `from`, `to`
  - `contractAddress` for CREATE transactions
  - `status`
  - `logs` (see §10)
  - `gasUsed` (when available), `effectiveGasPrice` (may be omitted)

Receipt persistence:
- Receipts are persisted and should survive restarts in normal ops mode.

---

## 8. World-State Reads

World-state reads are considered **Leader-only** unless explicitly stated otherwise.

### 8.1 `eth_getBalance`  *(Leader-only)*
- **Supported**
- Params: (address, blockTag)
- Supported tags: `"latest"` (others may be treated as `"latest"`).

### 8.2 `eth_getTransactionCount` *(Leader-only)*
- **Supported**
- Returns account nonce as of head.

### 8.3 `eth_getCode` *(Leader-only)*
- **Status: implementation-dependent**
- If supported, returns deployed code bytes (`0x…`) at `"latest"`.
- If not supported in a given build, method may return `0x` or `-32601`.  
  Ops MUST align this method’s behavior with the currently tagged implementation.

### 8.4 `eth_getStorageAt` *(Leader-only)*
- **Status: implementation-dependent**
- If supported, returns 32-byte slot value at `"latest"`.

---

## 9. Call & Gas

### 9.1 `eth_call` *(Leader-only)*
- **Supported (read-only call path)**
- Executes a call without persisting state changes.
- Intended use: read-only contract methods, tooling reads, and scripted queries.

Constraints:
- `"pending"` treated as `"latest"`.
- Revert data / error surfaces may be partial depending on build tag.

### 9.2 `eth_estimateGas` *(Leader-only)*
- **Supported (heuristic)**
- Gas estimates may be conservative or simplified.
- Clients should tolerate estimates that differ from Ethereum mainnet norms.

---

## 10. Logs

### 10.1 `eth_getLogs`
- **Supported (bounded / ops-first)**
- Intended primarily for:
  - retrieving logs for a bounded block range
  - confirming receipts/log emission for known contracts

Constraints (explicit):
- Large ranges may be rejected or truncated.
- Some advanced filter combinations may be unsupported.

### 10.2 No subscriptions
- `eth_subscribe` / `eth_unsubscribe` are **not supported** (no websocket pubsub in this spec).

---

## 11. Accounts & Signing

The node does not provide personal account management.

- `personal_*` methods: **not supported**
- `eth_sendTransaction`: **not supported** (requires unlocked accounts)
- Clients **MUST** sign locally and submit via `eth_sendRawTransaction`.

---

## 12. Misc Compatibility Methods

These methods exist to satisfy common tooling expectations.

- `web3_sha3`: may be supported (implementation-dependent)
- `eth_getBlockByHash`: may be supported (implementation-dependent)
- `eth_gasPrice`: may be supported with a constant or heuristic value

If a method is not present in the currently tagged build, it **MUST** return `-32601`.

---

## 13. Leader-only Method List (Normative)

A node configured with `FollowRPC` (follower mode) **MUST** proxy the following methods to the leader:

- `eth_getBalance`
- `eth_getTransactionCount`
- `eth_getCode` (if present)
- `eth_getStorageAt` (if present)
- `eth_call`
- `eth_estimateGas`
- `eth_getLogs` (recommended leader-only in ops mode)

This list is conservative by design. It may be reduced only when follower state parity is proven operationally.

---

## 14. Security Posture (Ops Guidance)

- RPC **SHOULD** remain bound to loopback in default runbooks.
- If RPC must be exposed:
  - terminate TLS at a reverse proxy
  - enforce IP allowlists
  - enforce authentication (mTLS or token)
  - rate-limit and request-size limit
- The node does not embed an auth mechanism; exposure without a gateway is not supported.

---

## 15. Examples (curl)

### 15.1 chainId
```bash
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
15.2 blockNumber
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}'

15.3 getBlockByNumber (latest)
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":3,"method":"eth_getBlockByNumber","params":["latest",false]}'

15.4 submit raw tx (signed client-side)
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":4,"method":"eth_sendRawTransaction","params":["0x..."]}'

16. Change Control

This spec is tied to the evm-l1 tagged state. Any behavioral change to an RPC method MUST be reflected here as part of the same change-set (code + docs), and validated in the relevant runbook gates.
