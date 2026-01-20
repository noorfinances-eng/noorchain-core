NOORCHAIN 2.1 — Genesis Checklist

Document ID: GENESIS_CHECKLIST_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Operational checklist for producing and validating a NOORCHAIN 2.1 genesis configuration for controlled deployments.

1. Purpose

This checklist defines a deterministic procedure to produce and validate a NOORCHAIN 2.1 genesis configuration.

It is designed to:

prevent configuration drift

ensure chain identity is consistent across nodes

ensure allocations and parameters are correct

ensure operator evidence exists for audit readiness

This checklist does not replace the genesis specification. Normative requirements are defined in:

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

2. Inputs (Must Be Defined Before Execution)

Define the following inputs before generating genesis:

Environment name: (e.g., local, controlled VPS, pilot)

Chain ID: (string identifier)

Network ID / chainId (EVM): (numeric) and its hex representation

Genesis timestamp policy: fixed or operator-chosen (must be documented)

Initial validator/operator set (if applicable to the environment)

Allocations: addresses and balances, including dev funding if required

PoSS application-layer configuration (if applicable)

Ports and exposure policy: localhost-only vs controlled private exposure

Record these inputs as the genesis intent record.

3. Pre-Generation Checklist
3.1 Repository and Release Identity

 Work from a clean working tree

 Confirm the intended release tag/commit to be used for this genesis

 Confirm documentation set and spec match the intended release

Evidence to capture:

git rev-parse HEAD

git status

tag reference if applicable

3.2 Parameter Review

 Review all genesis-relevant parameters against PARAMETERS_REFERENCE_2.1.md

 Confirm chain identity and EVM chainId consistency

 Confirm default ports and exposure posture align with deployment model

3.3 Allocation Review

 Confirm all required addresses are present

 Confirm balances are expressed in wei and within bounds

 Confirm no secrets are present in alloc artifacts

 Confirm dev funding is only included where appropriate

Reference:

docs/genesis/ALLOC_POLICY_2.1.md

4. Generation Checklist
4.1 Deterministic Genesis Build Inputs

 Freeze the genesis input files (parameters + alloc)

 Ensure generation tool version is pinned (repository commit/tag)

 Ensure timestamp policy is documented (fixed or recorded value)

4.2 Generate Genesis Artifact

 Generate the genesis file/artifact using the project’s standard procedure

 Record the command used and outputs

 Compute and record a digest of the genesis artifact (hash)

The exact generation mechanism may be environment-specific; the rule is that it must be reproducible and evidenced.

5. Validation Checklist (Single Node)

Run a single node with the genesis artifact and validate:

5.1 Process and Ports

 Node starts without error

 RPC port is listening on expected address

 Health endpoint responds (if enabled)

5.2 RPC Identity

 eth_chainId returns expected chain id value

 eth_blockNumber returns a valid quantity

5.3 Allocation Validation

For each required funded address:

 eth_getBalance(address, "latest") returns the expected balance

If nonce initialization is defined by the genesis model:

 eth_getTransactionCount(address, "latest") returns expected value

5.4 Block Production / Liveness

 Block number progresses (if block production is enabled)

 eth_getBlockByNumber("latest", false) returns a coherent block structure

6. Validation Checklist (Multi-Node) — If Applicable

If deploying leader/follower (or multi-node) environments:

6.1 Leader/Follower Start

 Leader and follower use distinct data directories

 Leader and follower ports do not collide

 P2P sessions are established

6.2 RPC Parity Gates

 eth_chainId matches on leader and follower

 eth_blockNumber is coherent on both

 eth_getBlockByNumber("latest", false) returns consistent metadata

 Parity-gated world-state reads return equivalent results (as specified)

The exact parity scope is defined in:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

7. Evidence Pack (Genesis)

A genesis evidence pack must include:

 repository commit hash and tag reference

 genesis inputs (parameters + alloc file versions)

 genesis artifact digest (hash)

 node start commands used

 RPC identity outputs (eth_chainId, eth_blockNumber)

 balance verification outputs for funded accounts

 multi-node parity outputs (if applicable)

Evidence must be sufficient for an independent operator to verify the same genesis results under the same release tag.

8. Change Control

Genesis changes are high impact.

Any change to genesis inputs must be documented:

record the reason

update changelog if it affects behavior

ensure specifications remain aligned

References:

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

9. References

docs/genesis/GENESIS_SPEC_2.1.md

docs/genesis/PARAMETERS_REFERENCE_2.1.md

docs/genesis/ALLOC_POLICY_2.1.md

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md