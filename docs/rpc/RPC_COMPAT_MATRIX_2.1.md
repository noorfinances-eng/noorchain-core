# NOORCHAIN 2.1 — RPC Compatibility Matrix (Tooling / Wallets)

Version: v2.1-draft.0  
Status: Operator-first reference (non-normative, but precise)  
Last updated: 2026-01-19 (Europe/Zurich)

This document summarizes JSON-RPC method support and behavioral constraints for **NOORCHAIN 2.1 (evm-l1)**, with a focus on real-world tooling and wallet expectations.

Authoritative specification remains:
- `docs/RPC_SPEC_2.1.md`
- `docs/STATE_MODEL_2.1.md`

---

## 0. Compatibility Position

NOORCHAIN 2.1 targets **practical Ethereum tooling compatibility** in controlled environments through a **subset-based JSON-RPC surface**.

Principles:
- Prefer methods required by standard wallets and SDKs (ethers/viem/Hardhat).
- Make leader/follower behavior explicit and operationally safe.
- Unsupported methods fail cleanly (`-32601`) rather than returning misleading placeholders.

Chain identity:
- `chainId` = 2121 (`0x849`)

---

## 1. Role Semantics (Leader / Follower)

Default operational model:
- **Leader** is authoritative for `"latest"` state.
- **Follower** MAY serve safe reads, but MUST proxy **leader-only** methods when configured with `FollowRPC`.

If `FollowRPC` is set:
- Follower MUST proxy all leader-only reads to the leader.
- This avoids “looks-up-to-date” divergence issues and ensures deterministic tooling results.

---

## 2. Method Support Matrix

Legend:
- ✅ Supported
- ⚠️ Supported with constraints
- ❌ Not supported (must return `-32601`)

Role:
- **Any**: can be served locally by leader or follower
- **Leader-only**: must be served by leader; follower proxies if `FollowRPC` is set
- **Follower-proxy**: follower returns leader result (operational rule)

### 2.1 Identity / Network

| Method | Support | Role | Notes |
|---|---:|---|---|
| `eth_chainId` | ✅ | Any | Returns `0x849` |
| `net_version` | ✅ | Any | Decimal string `"2121"` |
| `web3_clientVersion` | ✅ | Any | Informational only |

### 2.2 Blocks / Headers

| Method | Support | Role | Notes |
|---|---:|---|---|
| `eth_blockNumber` | ✅ | Any | Latest head height |
| `eth_getBlockByNumber` | ✅ | Any | If height > head → `null`; includes `stateRoot/receiptsRoot/logsBloom` when available |
| `eth_getBlockByHash` | ⚠️ | Any | Implementation-dependent; document if enabled |

### 2.3 Transaction Path

| Method | Support | Role | Notes |
|---|---:|---|---|
| `eth_sendRawTransaction` | ✅ | Leader-only | Client must sign locally |
| `eth_getTransactionByHash` | ✅ | Any | `null` if unknown |
| `eth_getTransactionReceipt` | ✅ | Any | `null` until mined; receipt persists across restarts |

### 2.4 World-State Reads

| Method | Support | Role | Notes |
|---|---:|---|---|
| `eth_getBalance` | ✅ | Leader-only | `"latest"` anchored to committed head root |
| `eth_getTransactionCount` | ✅ | Leader-only | Nonce at `"latest"` |
| `eth_getCode` | ⚠️ | Leader-only | If enabled: returns bytecode; otherwise `-32601` or `0x` depending on build policy (spec must match reality) |
| `eth_getStorageAt` | ⚠️ | Leader-only | If enabled: returns 32-byte value; otherwise `-32601` |

### 2.5 Call / Gas

| Method | Support | Role | Notes |
|---|---:|---|---|
| `eth_call` | ✅ | Leader-only | Read-only; `"pending"` treated as `"latest"` |
| `eth_estimateGas` | ⚠️ | Leader-only | Heuristic; may be conservative |
| `eth_gasPrice` | ⚠️ | Any | If exposed, may be constant/heuristic |

### 2.6 Logs

| Method | Support | Role | Notes |
|---|---:|---|---|
| `eth_getLogs` | ⚠️ | Leader-only (recommended) | Bounded/ops-first; large ranges may be rejected or truncated |

### 2.7 Not Supported (by design)

| Method family | Support | Rationale |
|---|---:|---|
| `personal_*` | ❌ | No node-managed accounts |
| `eth_sendTransaction` | ❌ | Requires unlocked accounts |
| `eth_subscribe` / `eth_unsubscribe` | ❌ | No websocket pubsub baseline |
| `txpool_*` | ❌ | Not exposed in ops-first baseline |
| `debug_*` / `trace_*` | ❌ | Not part of baseline surface |

If a method is not explicitly supported, the server MUST return `-32601`.

---

## 3. Wallet Compatibility (Practical)

### 3.1 MetaMask / Browser wallets
Expected to work when the node provides:
- `eth_chainId`, `net_version`
- `eth_blockNumber`, `eth_getBlockByNumber`
- `eth_getBalance`, `eth_getTransactionCount`
- `eth_sendRawTransaction`
- `eth_getTransactionReceipt`

Constraints:
- No `personal_*` methods: signing remains inside the wallet.
- If the wallet relies heavily on filters/subscriptions, it may degrade gracefully or require polling.

Operational guidance:
- Point wallet to **Leader RPC**.
- If using follower RPC, ensure follower has `FollowRPC` configured so leader-only reads proxy correctly.

### 3.2 Hardware wallets
Works through the same flow (local signing). No node changes required.

---

## 4. Tooling Compatibility

### 4.1 ethers.js
Works with:
- chain identity calls
- tx submission + receipt polling
- read-only contract calls (`eth_call`)

Sensitive points:
- Gas estimation may be heuristic; scripts should tolerate higher-than-necessary estimates.

### 4.2 viem
Works with:
- explicit chain id
- `eth_getTransactionCount` correctness (nonces must advance)
- `eth_getTransactionReceipt.contractAddress` correctness for CREATE deploy flows

### 4.3 Hardhat
Works with:
- external network config (RPC URL)
- local signing via provided private key
- deployment scripts via `eth_sendRawTransaction`

Non-goals:
- no built-in Hardhat network features (forking, tracing, automining semantics).

### 4.4 Foundry (forge/cast)
Works with:
- `cast send --rpc-url ...` (raw tx submission)
- `cast call --rpc-url ...` (`eth_call`)
- `cast receipt` / polling

If `eth_getCode` and `eth_getStorageAt` are not enabled, some contract inspection workflows will be limited.

---

## 5. Known Failure Modes (Actionable)

### 5.1 “Follower returns different nonce/balance than leader”
Cause:
- follower served `"latest"` locally without proxying.

Fix:
- configure follower with `FollowRPC` and enforce leader-only proxy routing per `RPC_SPEC_2.1`.

### 5.2 “Deploy succeeded but contractAddress is null”
Cause:
- receipt persistence or CREATE address derivation inconsistent.

Fix:
- receipt storage must include `contractAddress` for CREATE receipts and must survive restarts. Validate with a deploy + receipt read.

### 5.3 “eth_call returns empty or inconsistent”
Cause:
- call path not anchored to committed head root or execution is constrained.

Fix:
- ensure `"latest"` anchor is `stateroot/v1/head` and call execution uses committed state. Document any remaining constraints explicitly.

### 5.4 “Tooling expects filters/subscriptions”
Cause:
- `eth_newFilter`, `eth_subscribe` not provided.

Fix:
- use polling patterns (`eth_getLogs` bounded ranges) and receipts. Keep constraints explicit.

---

## 6. Operational Validation Checklist (Minimal)

These checks should pass on the **leader** after a fresh start (and remain consistent after restart):

1) `eth_chainId` returns `0x849`
2) `eth_blockNumber` is non-null and progresses
3) `eth_getBlockByNumber("latest", false)` returns non-null and includes roots/bloom when available
4) `eth_getTransactionCount(addr,"latest")` returns a stable nonce
5) `eth_getBalance(addr,"latest")` returns expected value (including alloc-funded dev accounts if used)
6) Submit a signed tx:
   - `eth_sendRawTransaction` → hash
   - `eth_getTransactionReceipt(hash)` eventually non-null
   - receipt persists across restart
7) If follower is enabled with `FollowRPC`:
   - follower `eth_getBalance` equals leader
   - follower `eth_getTransactionCount` equals leader

---

## 7. Change Control

Any change to RPC behavior (method presence, semantics, constraints) MUST be reflected in:
- `docs/RPC_SPEC_2.1.md` (authoritative)
- and this matrix (practical impact)

This file exists to keep operator expectations aligned with real tooling behavior.
