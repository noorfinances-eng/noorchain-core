NOORCHAIN 2.1 — RPC Error Model

Document ID: RPC_ERROR_MODEL_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines the JSON-RPC error model for NOORCHAIN 2.1, including stable error shapes, codes, and behavioral requirements for unsupported methods, invalid parameters, and node role constraints. This document ensures consistent tooling behavior in controlled deployments.

Purpose

This document defines a stable JSON-RPC error model for NOORCHAIN 2.1.

It is designed to:

ensure unsupported methods fail predictably

provide consistent error shapes and codes for tooling compatibility

define validation rules for parameters and request shapes

prevent “silent success” behavior that breaks clients

define follower behavior when methods must be leader-only

This document is normative for JSON-RPC error behavior.

Scope and Assumptions

Applies to all JSON-RPC endpoints exposed by NOORCHAIN 2.1 nodes.

Applies to leader and follower roles.

Assumes standard JSON-RPC 2.0 envelope.

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/RPC_SPEC_2.1.md

Normative Requirements

3.1 JSON-RPC Envelope

All responses MUST be valid JSON-RPC 2.0:

exactly one of result or error MUST be present

jsonrpc MUST equal "2.0"

id MUST be preserved from request (or null if request id was null)

3.2 Error Object Shape

On error, nodes MUST return:

{
  "jsonrpc": "2.0",
  "id": <same as request>,
  "error": {
    "code": <integer>,
    "message": "<short, stable string>",
    "data": <optional, structured>
  }
}


Requirements:

code MUST be an integer

message MUST be a stable, human-readable category string

data SHOULD be used for structured diagnostics, not free-form logs

3.3 Stability

For a given error class, the node MUST keep:

the error code stable

the error message category stable

the presence/shape of data stable (if used)

Tooling depends on stability. Avoid embedding volatile content (timestamps, stack traces) in message.

Error Code Set

This document defines a canonical code set for NOORCHAIN 2.1. Where possible, it aligns with JSON-RPC and common Ethereum client expectations.

4.1 Standard JSON-RPC Codes

The following codes are reserved and may be used:

-32700 Parse error

-32600 Invalid Request

-32601 Method not found

-32602 Invalid params

-32603 Internal error

4.2 NOORCHAIN-Specific Codes (Recommended)

The following codes are defined for stable policy errors:

-32000 Generic server error (bounded, stable category)

-32001 Not supported (method exists in some clients but is disabled here)

-32002 Leader required (method must be served by leader; follower refusal without proxy)

-32003 Not ready (state unavailable, node warming, or not yet indexed)

-32004 Not found (object absent: block/tx/receipt/log not found)

-32005 Rate limited / temporarily unavailable (if applicable)

If NOORCHAIN-specific codes are used, they MUST remain stable across releases unless explicitly changed via changelog and compatibility gates.

Error Classes and Rules

5.1 Parse Error (-32700)

Condition:

request body is not valid JSON

Message:

"parse error"

Data:

SHOULD be omitted or minimal.

5.2 Invalid Request (-32600)

Condition:

request is not a valid JSON-RPC 2.0 object

missing jsonrpc, missing method, invalid id types, etc.

Message:

"invalid request"

Data:

optional structured explanation (field name, expected type)

5.3 Method Not Found (-32601)

Condition:

method name not implemented

Message:

"method not found"

Data:

SHOULD be omitted.

Normative:

Any method not listed in the compatibility matrix MUST return this error (or -32001 Not supported if method name is reserved but intentionally disabled). The choice must be consistent.

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

5.4 Invalid Params (-32602)

Condition:

parameters present but malformed, wrong type, wrong length, invalid hex quantity, invalid address format, etc.

Message:

"invalid params"

Data:

structured details recommended:

{ "param": "<name or index>", "reason": "<category>" }

Normative:

Do not return success with default values when params are invalid.

5.5 Internal Error (-32603)

Condition:

unhandled exception or unexpected failure

Message:

"internal error"

Data:

may include a bounded error category (not stack traces)

Normative:

internal error should be rare; repeated internal errors are incidents.

Reference:

docs/ops/INCIDENTS_2.1.md

5.6 Not Supported (-32001)

Condition:

method is known but intentionally disabled or deferred (e.g., filters, subscriptions)

Message:

"not supported"

Data:

may include { "method": "eth_newFilter" }

Normative:

Use -32001 when you want stability over -32601 for known-but-disabled methods.

5.7 Leader Required (-32002)

Condition:

follower receives a leader-only method and cannot (or is not configured to) proxy to leader

Message:

"leader required"

Data:

may include { "method": "<name>", "hint": "call leader or configure FollowRPC" }

Normative:

If follower has FollowRPC configured and the method is designated proxy-safe, the follower SHOULD proxy rather than returning -32002.

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/ops/DEPLOYMENT_MODEL_2.1.md

5.8 Not Ready (-32003)

Condition:

method depends on state that is not available yet (e.g., node not fully initialized, no head state, indexing not complete)

Message:

"not ready"

Data:

optional { "reason": "<category>" }

Normative:

Not ready must be temporary; if persistent, treat as incident.

5.9 Not Found (-32004)

Condition:

requested object does not exist:

unknown transaction hash

receipt not yet available

block not found by hash

log query yields no results where query expects existence (optional)

Message:

"not found"

Data:

optional:

{ "type": "transaction|receipt|block|log", "id": "<hash or selector>" }

Normative:

For some Ethereum-like methods, returning null is expected rather than an error (e.g., eth_getTransactionReceipt before mined). If null is used, it must be consistently used and documented in the method-specific behavior. Do not alternate between error and null unpredictably.

Reference:

docs/RPC_SPEC_2.1.md

Method-Specific Error Semantics (Selected)

This section defines specific requirements for commonly used methods.

6.1 eth_getBlockByNumber

If requested height > latest: return result: null (not error).

If params invalid: -32602.

If internal failure: -32603.

6.2 eth_getTransactionReceipt

If tx exists but receipt not yet available: return result: null (not error).

If tx unknown: may return result: null (preferred for compatibility) or -32004; choice must be stable and documented.

If params invalid: -32602.

6.3 eth_sendRawTransaction

If raw tx is malformed: -32602.

If leader required (called on follower): either proxy to leader or return -32002.

If rejected by policy: -32000 with stable message category (e.g., "rejected").

6.4 eth_call (if supported)

If method unsupported: -32001 or -32601 consistently.

If contract execution reverts: return an error with stable code category; do not return success with "0x" unless that is the actual return data.

If invalid params: -32602.

Address and Hex Validation Rules

7.1 Hex Quantity

A hex quantity MUST:

be a hex string with 0x prefix

represent a non-negative integer

have no leading zeros (except "0x0")

Invalid hex quantities MUST trigger -32602.

7.2 Hex Data

Hex data MUST:

start with 0x

have even-length hex after prefix

be lowercase/uppercase tolerant, but output format should be stable per method spec.

7.3 Address Format

Addresses MUST:

be 0x-prefixed 20-byte hex (40 hex chars after 0x)

invalid addresses MUST trigger -32602

output addresses should be consistently checksummed or consistently lowercased (project-defined); do not mix formats unpredictably

Reference:

docs/RPC_SPEC_2.1.md

Logging and Error Privacy

Errors MUST NOT include:

private keys

secrets (env vars)

internal topology details

filesystem paths that reveal sensitive structure (where avoidable)

If such leakage is suspected, treat as an incident.

Reference:

docs/PRIVACY_DATA_POLICY_2.1.md

docs/ops/INCIDENTS_2.1.md

Change Control

Any change to error codes or message categories is a compatibility change and MUST:

be recorded in changelog

be validated against tooling compatibility gates

be referenced in release process evidence

References:

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

References

docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/STATE_MODEL_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md
docs/ops/INCIDENTS_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/CHANGELOG_2.1.md
docs/RELEASE_PROCESS_2.1.md