NOORCHAIN 2.1 — Tooling (Hardhat / viem)

Document ID: TOOLING_HARDHAT_VIEM_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Supported patterns and known pitfalls when using Hardhat and viem against NOORCHAIN 2.1 JSON-RPC in controlled environments.

1. Purpose

This document provides operationally useful guidance for integrating standard Ethereum tooling with NOORCHAIN 2.1:

Hardhat (v3+ patterns)

viem-based scripts and wallet-style workflows

It focuses on:

stable configuration patterns

known failure modes and their causes

evidence-first debugging gates

Normative RPC semantics are defined in:

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

2. General Assumptions

JSON-RPC endpoint is reachable (localhost or via tunnel).

Chain ID is known (e.g., 0x849 for a given environment) and must match the tooling configuration.

Private keys are managed out of band and are never committed.

3. Hardhat (v3+) — Supported Configuration Pattern
3.1 Config File Format

Hardhat v3 expects modern configuration patterns. If tasks/plugins are not loading:

prefer ESM configuration (hardhat.config.mjs)

use explicit plugin registration as required

3.2 Network Naming Discipline

Use the repo-defined network name. If a network is not defined, Hardhat will error.

Common errors:

HHE705: network <name> not defined

HHE3: No Hardhat config file found

Operator rule:

run from repository root

pass --config hardhat.config.mjs when required

use the correct network name (example: noorcore)

3.3 Private Key Handling

Hardhat deployments typically require a funded deployer key.

If NOOR_PRIVATE_KEY (or the configured key env var) is unset or invalid:

ethers.getSigners() may return an empty list

deploy scripts may fail with:

TypeError: Cannot read properties of undefined (reading 'address')

Policy:

do not hardcode keys in code

export the private key via environment variable at runtime

do not log the private key

4. viem — Supported Pattern

viem can be used as a client when NOORCHAIN 2.1 provides the required RPC methods for:

chain identity

account nonce reads

gas estimation (if required by the flow)

sendRawTransaction and receipt polling (for tx flows)

Compatibility coverage is tracked in:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

5. Operational Gate Checks Before Tooling

Before running Hardhat/viem scripts, verify the RPC endpoint.

5.1 RPC Liveness Gate
curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
echo

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}'
echo


If either fails, fix node liveness before debugging tooling.

5.2 Port Binding Gate
ss -ltnp | egrep 'noorcore|:8545'

5.3 Process Gate
pgrep -a noorcore

6. Common Failure Modes (Hardhat)
6.1 RPC Down → ECONNREFUSED

Symptom:

Hardhat throws a connection refusal error

viem cannot reach RPC

Cause:

node not running or RPC not bound on expected address

Fix:

run liveness and port gates

inspect node logs

restart only after confirming cause

6.2 “No Hardhat config file found” (HHE3)

Symptom:

Hardhat cannot locate config

Cause:

running from the wrong working directory

config file name mismatch

missing --config flag when required

Fix:

run from repo root

pass --config hardhat.config.mjs

6.3 “network not defined” (HHE705)

Symptom:

Hardhat does not know the network name

Cause:

using an incorrect network identifier

Fix:

use the configured network name (example: --network noorcore)

6.4 Signer Undefined (Private Key Missing)

Symptom:

deploy script fails because deployer is undefined

Cause:

private key env var unset or invalid

account not configured in Hardhat network settings

Fix:

set the private key via env var at runtime

verify balance and nonce via RPC before deploy

7. Common Failure Modes (viem)
7.1 Deploy Fails Because Receipt contractAddress Is Null

Symptom:

viem deploy returns tx hash but cannot derive contract address

Cause:

node returns incomplete receipts or does not compute contractAddress for CREATE

nonce handling mismatch

Fix:

ensure eth_getTransactionReceipt includes contractAddress for CREATE txs

ensure nonce increments are consistent with eth_getTransactionCount

This is a compatibility gate item tracked in the compatibility matrix.

7.2 eth_estimateGas Returns Error or Unusable Output

Symptom:

viem refuses to send tx because gas estimate fails

Cause:

eth_estimateGas not implemented or returns non-standard behavior

Fix:

confirm whether eth_estimateGas is supported per compatibility matrix

use explicit gas fields in tooling where allowed (controlled testing)

align node behavior with spec if method is claimed supported

8. Funding and Alloc (Developer Baseline)

Tooling requires funded accounts.

In controlled environments, funding is typically achieved via an alloc file applied at node startup.

Operator guidance:

ensure alloc file includes deployer address and balance

restart node as required for alloc application

confirm funding via eth_getBalance

Balance check example:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":10,"method":"eth_getBalance","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}'
echo


Alloc policy is documented in:

docs/genesis/ALLOC_POLICY_2.1.md

9. Recommended Debug Workflow (Evidence-First)

When a deploy or tx flow fails:

Verify node liveness and ports (RPC gates)

Verify chain id matches tooling config

Verify deployer balance and nonce (RPC reads)

Send tx and record tx hash

Poll receipt (record receipt output)

If mismatch persists, compare leader vs follower behavior if follower is used

Record:

commands

outputs

node logs (relevant lines)

10. Security Notes for Tooling

Never commit private keys.

Avoid printing secrets to console.

Treat RPC endpoints as sensitive infrastructure.

Prefer localhost RPC + SSH tunnel for remote development.

These are policy-level constraints:

docs/PRIVACY_DATA_POLICY_2.1.md

docs/compliance/COMMUNICATIONS_POLICY_2.1.md

11. References

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/dev/JSON_RPC_EXAMPLES_2.1.md

docs/genesis/ALLOC_POLICY_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md