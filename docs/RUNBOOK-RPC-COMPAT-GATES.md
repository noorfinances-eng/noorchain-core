NOORCHAIN 2.1 — Runbook: RPC Compatibility Gates

Document ID: RUNBOOK-RPC-COMPAT-GATES
Version: v1.0
Date: January 2026
Status: Active
Scope: Defines a deterministic RPC compatibility gate suite for NOORCHAIN 2.1 controlled deployments. The gates verify identity, liveness, block metadata coherence, transaction path, receipt behavior, and leader/follower routing parity. This runbook is designed to be executed as an evidence-producing checklist.

1. Purpose

This runbook defines the canonical RPC compatibility gate suite for NOORCHAIN 2.1.

It is designed to:

prove RPC behavior is stable for tooling/wallet usage

verify leader/follower routing semantics are correct

prevent regressions across releases and upgrades

produce an evidence pack suitable for audit readiness

This runbook is normative for the operational claim: “RPC surface is compatible at the declared level”.

References:
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/rpc/RPC_ERROR_MODEL_2.1.md
docs/AUDIT_READINESS_2.1.md

2. Preconditions

2.1 Environment Inputs (Must Be Recorded)

Record the following before running gates:

environment name (local / pilot / permissioned)

leader RPC URL (host:port)

follower RPC URL (host:port) if applicable

expected eth_chainId (hex)

known funded test address (optional but recommended)

a known tx hash (optional) or a plan to submit a test tx

Store these inputs in the evidence pack header.

2.2 Exposure Posture

RPC must be bound to localhost or controlled private access per deployment model. Do not widen exposure to run gates.

Reference:
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/PRIVACY_DATA_POLICY_2.1.md

2.3 Tooling Requirements

Gates may be executed with:

curl (preferred for deterministic evidence)

optional project tooling (Hardhat/Viem) to generate a tx and receipt

All outputs included in evidence should be captured verbatim.

3. Gate Suite Overview

Gates are grouped by category:

Gate A: Identity

Gate B: Liveness and Block Reads

Gate C: Block Metadata Coherence (roots/bloom)

Gate D: Transaction Path

Gate E: Receipt Behavior

Gate F: World-State Read Parity (leader/follower routing)

Gate G: Error Model and Unsupported Methods

Each gate has:

objective

request(s)

pass/fail criteria

evidence to capture

4. Gate A — Identity

Objective: prove chain identity is correct and stable.

Requests:

eth_chainId on leader

eth_chainId on follower (if applicable)

Pass criteria:

leader returns expected chainId (e.g., 0x849)

follower returns the same value

responses are valid JSON-RPC

Evidence:

full request/response JSON for each endpoint

Reference:
docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md

5. Gate B — Liveness and Block Reads

Objective: prove node is responsive and block read works.

Requests:

eth_blockNumber on leader

eth_getBlockByNumber("latest", false) on leader

If follower exists:

repeat on follower

Pass criteria:

eth_blockNumber returns a valid hex quantity

eth_getBlockByNumber returns:

non-null result

coherent fields (at minimum: number, hash, parentHash)

follower results are coherent with leader (same or explainable under follow semantics)

Evidence:

full request/response JSON

6. Gate C — Block Metadata Coherence (Roots/Bloom)

Objective: ensure block metadata fields required by evidence packs are present and stable (where implemented).

Requests:

eth_getBlockByNumber("latest", false) on leader

optionally also for a fixed height (if you want stable comparison)

If follower exists:

repeat on follower

Pass criteria:

response includes:

stateRoot (non-null, 0x-prefixed 32 bytes)

receiptsRoot (non-null, 0x-prefixed 32 bytes)

logsBloom (non-null, 0x-prefixed 256 bytes)

follower values are coherent with leader for the same block number/hash

If a field is not supported in the declared compatibility level, that must be documented in the matrix and the gate must be marked “N/A” for that environment.

References:
docs/STATE_MODEL_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md

7. Gate D — Transaction Path (Leader)

Objective: prove the basic tx path works end-to-end on the leader.

This gate requires a test transaction.

Two execution options:

Option 1 (Recommended): Submit a Known Test Tx

submit a raw tx to leader via eth_sendRawTransaction

capture returned tx hash

Option 2: Use an Existing Known Tx Hash

use a previously mined tx hash (must be recorded in evidence header)

Requests:

eth_sendRawTransaction (if generating a new tx)

eth_getTransactionByHash(txHash)

eth_getTransactionReceipt(txHash)

Pass criteria:

sendRawTransaction returns a 0x-prefixed tx hash

getTransactionByHash returns non-null result with coherent fields

getTransactionReceipt returns non-null after mining (may require polling with bounded retries documented in evidence)

Evidence:

raw request/response JSON for each call

if polling, record each attempt with timestamp (UTC) and stop after policy limit

Reference:
docs/RPC_SPEC_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md

8. Gate E — Receipt Behavior (Contract Deploy Compatibility)

Objective: ensure receipts are usable by tooling (deploy flows).

This gate requires a CREATE transaction or a known deploy tx hash.

Requests:

eth_getTransactionReceipt(deployTxHash)

Pass criteria:

receipt is non-null after mining

receipt includes:

status (0x1 success for successful deploy)

transactionHash matches input

blockNumber non-null

contractAddress non-null and 20-byte hex (for CREATE)

logs may be empty depending on tx, but structure must be valid

Evidence:

receipt JSON

Reference:
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/rpc/RPC_ERROR_MODEL_2.1.md

9. Gate F — World-State Read Parity (Leader/Follower Routing)

Objective: prove follower routing semantics are correct for leader-only reads.

Precondition:

follower is configured with FollowRPC pointing to leader

follower is intended to serve clients

Requests (leader then follower):

eth_getBalance(testAddress, "latest")

eth_getTransactionCount(testAddress, "latest")

Optional (if implemented and required):

eth_getCode(contractAddress, "latest")

eth_getStorageAt(contractAddress, slot, "latest")

Pass criteria:

follower values match leader values for each method

no evidence of follower serving stale local state for leader-only reads

failures are treated as deployment model or routing incidents

Evidence:

side-by-side request/response JSON for leader and follower

record the chosen testAddress and why it is relevant (funded, active)

References:
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/STATE_MODEL_2.1.md

10. Gate G — Error Model and Unsupported Methods

Objective: verify unsupported or deferred methods fail predictably.

Select a small set of unsupported methods (examples):

eth_newFilter

eth_subscribe

eth_getFilterChanges

Requests:

call each method with minimal valid params shape

Pass criteria:

node returns a JSON-RPC error object with:

stable code

stable category message

optional bounded data

no malformed JSON-RPC envelopes

Evidence:

request/response JSON

record which error codes/messages were returned

Reference:
docs/rpc/RPC_ERROR_MODEL_2.1.md

11. Evidence Pack Format (RPC Gates)

The RPC gates evidence pack MUST include:

header inputs (env, endpoints, expected chainId, addresses used, tx hash used)

a timestamped log of gate execution (UTC)

verbatim JSON request/response pairs for each gate

pass/fail conclusion per gate

notes for any “N/A” gates with matrix justification

operator identity (internal)

release tag/commit reference for the node binary

References:
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md

12. Failure Handling and Escalation

If any gate fails:

stop further gates if failure indicates identity drift or severe inconsistency

preserve evidence of failure before restarting nodes

classify incident and follow incident response procedure

References:
docs/ops/INCIDENTS_2.1.md
docs/governance/INCIDENT_RESPONSE_2.1.md

13. Change Control

Any change to this gate suite is a compatibility control change and must:

be versioned

be aligned with the RPC compatibility matrix

be recorded in changelog if it changes operational expectations

References:
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/CHANGELOG_2.1.md
docs/RELEASE_PROCESS_2.1.md

14. References

docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/rpc/RPC_ERROR_MODEL_2.1.md
docs/RPC_SPEC_2.1.md
docs/STATE_MODEL_2.1.md
docs/ops/DEPLOYMENT_MODEL_2.1.md
docs/ops/INCIDENTS_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md