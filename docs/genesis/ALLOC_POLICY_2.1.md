NOORCHAIN 2.1 — Allocation Policy

Document ID: ALLOC_POLICY_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Genesis and bootstrap allocation policy for NOORCHAIN 2.1 controlled deployments, including development funding via alloc files.

1. Purpose

This policy defines how account allocations are handled for NOORCHAIN 2.1 in controlled deployments.

It covers:

allocation intent and constraints

alloc file structure expectations (operator-facing)

application rules (when alloc is applied)

evidence and validation gates

safe re-application procedures (when permitted)

This policy does not define token economics. Economic invariants and governance belong to separate documents.

2. Principles

Determinism
Alloc application must be deterministic and reproducible.

Minimality
Allocations should be the minimum required to operate the environment (e.g., dev funding).

Controlled Re-application
Alloc is not intended to be repeatedly applied without explicit operator intent.

No Secret Material
Alloc files never contain private keys, seed phrases, or secrets.

3. Allocation Categories
3.1 Genesis Allocation (Release-Like)

A genesis allocation is the initial state for a network environment.

Characteristics:

part of the environment definition

versioned and documented

changes require release discipline and evidence

3.2 Development Bootstrap Allocation (Local / Controlled)

A development bootstrap allocation is used to fund one or more dev addresses for:

deployments

transaction flows

tooling and integration tests

This is typically applied at node startup via an operator-provided alloc file.

4. Alloc File Requirements
4.1 Required Fields (Policy Level)

An alloc file must include:

a declared chain identifier scope (environment-specific)

one or more allocation entries:

address (0x-prefixed, 20-byte hex address)

balanceWei (string or numeric as defined by implementation conventions)

4.2 Allowed Values

addresses must be checksummed or normalized consistently; the node must treat addresses case-insensitively.

balanceWei must be non-negative and must fit within expected numeric bounds.

4.3 Prohibited Content

private keys

mnemonics

API tokens or credentials

personal data

5. Application Rules
5.1 When Alloc Is Applied

Alloc is applied only when:

node is started with an explicit alloc file flag (implementation-defined)

the node determines alloc has not been applied to that data-dir before

This prevents accidental repeated injections.

5.2 Idempotence Guard

The system may persist an “alloc applied” marker under the node’s data directory.

Policy expectation:

alloc application is one-time per data-dir by default

re-application requires an explicit operator action (see Section 6)

5.3 Scope

Alloc application is scoped to a specific node data directory:

Applying alloc to leader data-dir affects leader state.

Followers that route state reads to leader should observe leader state via routing, but should not attempt to re-apply alloc independently unless explicitly designed.

6. Re-application Procedure (Controlled)

Re-applying alloc is permitted only in controlled development contexts and must be done deliberately.

Policy requirements:

stop node cleanly

remove the explicit “alloc applied” marker (implementation-defined)

restart node with the alloc file flag

verify results via RPC gates

This procedure must be recorded as evidence when used.

7. Validation Gates

After starting a node with alloc:

7.1 Alloc Application Evidence (Logs)

Operators should confirm via logs that alloc was applied (or skipped due to guard). The exact log strings are implementation-defined, but the evidence must include:

alloc file path used

number of entries applied (or explicit “already applied” statement)

7.2 Balance Gate (RPC)

Verify funded accounts via eth_getBalance.

Example:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_getBalance","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}'
echo


Expected:

a non-zero hex quantity for funded accounts

7.3 Nonce Gate (Optional)

If allocations are used to enable multiple deploys, also confirm nonce reads work:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_getTransactionCount","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}'
echo

8. Operational Considerations
8.1 Separation from Token Economics

Alloc is an operational bootstrap mechanism and does not define the economic model.

Economic invariants must be maintained by governance and genesis specifications.

8.2 Avoiding Drift

Alloc files used for controlled environments should be:

documented

referenced by runbooks

treated as part of the environment configuration

Untracked ad-hoc alloc files increase non-reproducibility.

9. Change Control

Changes to allocation rules or alloc file format require:

changelog entry

spec alignment if RPC-visible behavior changes

release discipline for release-like environments

References:

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md

10. References

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/GENESIS_CHECKLIST_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/dev/TOOLING_HARDHAT_VIEM_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md