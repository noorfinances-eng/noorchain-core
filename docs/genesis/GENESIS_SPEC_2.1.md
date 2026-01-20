NOORCHAIN 2.1 — Genesis Specification

Document ID: GENESIS_SPEC_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Normative specification defining the required content, invariants, and determinism rules for producing a valid NOORCHAIN 2.1 genesis configuration for controlled deployments (local, pilot, permissioned networks).

Purpose

This specification defines what constitutes a valid NOORCHAIN 2.1 genesis configuration and the mandatory invariants that MUST be preserved across environments.

It exists to:

define the minimum required genesis content and semantics

freeze critical identity and supply invariants

ensure deterministic, reproducible genesis artifacts

provide an auditable baseline that operators can validate independently

prevent configuration drift across nodes and environments

This document is normative. The operational execution sequence is defined in:

docs/genesis/GENESIS_CHECKLIST_2.1.md

System Invariants (Must Hold)

2.1 Chain Identity

The EVM chainId MUST be 2121 (hex: 0x849).

The environment “chain ID” string identifier (operator-facing) MUST be unique per environment, but MUST NOT replace or conflict with the EVM chainId.

2.2 Native Asset

The native asset symbol is NUR.

Transaction fees (gas) are paid in NUR.

Genesis artifacts MUST NOT contain language implying yield, returns, or investment promises.

2.3 Fixed Supply

Total supply is fixed at 299,792,458 NUR.

Genesis + any bootstrapping mechanism used for that environment MUST preserve this invariant.

2.4 Allocation Model (Canonical Proportions)

The canonical allocation proportions are fixed:

80% — PoSS Reserve

5% — Foundation

5% — Dev Sàrl

5% — PoSS Stimulus Pool

5% — Pre-sale Reserve (optional activation depending on environment policy)

Genesis MUST encode allocation addresses and balances consistent with the policy defined in:

docs/genesis/ALLOC_POLICY_2.1.md

2.5 PoSS Separation

Consensus security is not PoSS.

PoSS is an application-layer system for value and governance.

Genesis MUST NOT introduce PoSS logic into consensus.

Genesis Outputs (Artifact Set)

A valid genesis release MUST include, at minimum:

genesis artifact (genesis.json or equivalent canonical configuration artifact)

a cryptographic digest of the canonical genesis artifact (hash)

a genesis intent record (human-readable) describing environment inputs and decisions

3.1 Required Files (Recommended Naming)

genesis.json

genesis.sha256

GENESIS_INTENT.md

If the implementation uses a different artifact naming scheme, the environment MUST still provide an equivalent set with the same semantics.

Inputs (Must Be Defined Before Genesis Build)

The following inputs MUST be defined and recorded in GENESIS_INTENT.md before generation:

Environment name (e.g., local, gdv, controlled VPS, pre-mainnet, mainnet)

Chain ID string identifier (operator-facing)

Network ID / chainId (EVM): 2121 and 0x849

Genesis timestamp policy (fixed or operator-chosen; must be recorded)

Initial validator/operator set (if applicable)

Allocation plan: addresses + balances (including dev funding policy if applicable)

Ports and exposure posture: localhost-only vs controlled private exposure

Any reserved address ranges and any predeploy policy (if applicable)

Genesis Content Requirements (Normative)

This section defines the minimum required content. Field names may vary by implementation, but the semantics MUST be preserved.

5.1 Global Parameters

Genesis MUST define:

EVM chainId = 2121 (0x849)

Genesis timestamp (and timestamp policy)

Initial block height policy (0 or 1; MUST be consistent across nodes)

Protocol configuration versioning (if fork configuration exists, it MUST be deterministic)

5.2 Initial World-State Commitments

Genesis MUST define the initial world-state consistent with the NOORCHAIN 2.1 state model.

At minimum, genesis MUST define:

initial accounts and balances (where funding is required by policy)

nonce initialization policy (typically zero unless explicitly required)

whether code/storage predeploys exist (default: none)

If the environment uses predeploys or reserved ranges, genesis MUST explicitly record:

the reserved address ranges

the list of predeployed addresses (if any)

code hashes (if applicable)

rationale and operational constraints

Reference:

docs/STATE_MODEL_2.1.md

5.3 Initial Validator Set (Permissioned BFT)

For permissioned BFT deployments, genesis MUST define the initial validator set deterministically.

Rules:

Validator identities MUST be public (no private key material in genesis).

The validator set MUST be identical across all nodes.

The validator set representation MUST be stable and unambiguous.

5.4 Allocation Encoding

Genesis MUST include the allocations required by policy and environment.

Rules:

Balances MUST be expressed in wei (or the smallest unit of NUR, per the implementation’s EVM model).

No secrets MUST be present in genesis artifacts.

Dev funding MUST only be included where explicitly permitted by policy.

Reference:

docs/genesis/ALLOC_POLICY_2.1.md

Alloc Overlay (Post-Genesis Bootstrap) — Policy-Bound

Some environments may permit a post-genesis alloc overlay applied at node startup (typically dev/pilot-only). If used:

Its use MUST be explicitly recorded in GENESIS_INTENT.md.

It MUST be deterministic and reproducible.

It MUST be applied consistently across the environment (no hidden overlays).

It MUST be governed by an idempotent rule (apply-once per data-dir) and that rule MUST be documented.

Important: Any environment that relies on post-genesis overlays for permanent distribution MUST be treated as operationally constrained and policy-bound until a formal genesis distribution is produced.

Reference:

docs/genesis/ALLOC_POLICY_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

Determinism and Canonicalization

Genesis determinism is mandatory.

7.1 Canonical Encoding

The canonical bytes of genesis.json used for hashing MUST be produced with:

stable key ordering

stable whitespace rules

no floating point representations

explicit integer encodings

If the implementation cannot guarantee canonical JSON, a canonicalization tool MUST be used and recorded in GENESIS_INTENT.md.

7.2 Digest

genesis.sha256 MUST be computed from the canonical genesis.json bytes.

The digest MUST be distributed out-of-band to operators and recorded in the evidence pack.

Operators MUST reject joining a network if the genesis digest mismatches.

Validation Requirements (Must Pass)

8.1 Identity Gates

eth_chainId returns 0x849.

The operator-facing chain identifier matches the intended environment.

8.2 Allocation Gates

For each required funded address:

eth_getBalance(address, "latest") equals the expected balance.

If nonce initialization is defined:

eth_getTransactionCount(address, "latest") equals the expected nonce.

8.3 Liveness Gates

eth_blockNumber returns a valid quantity.

eth_getBlockByNumber("latest", false) returns a coherent structure per RPC spec.

References:

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

8.4 Multi-Node Coherence Gates (If Applicable)

For leader/follower or multi-node environments:

distinct data directories (no shared LevelDB locks)

no port collisions

P2P sessions established

RPC parity gates pass per matrix definition

Reference:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

Evidence Pack Requirements (Genesis)

A genesis evidence pack MUST include sufficient evidence for an independent operator to reproduce and validate the genesis.

Minimum evidence:

repository commit hash (git rev-parse HEAD)

release tag reference (if applicable)

genesis inputs (parameter values, alloc file versions)

canonical genesis artifact digest (sha256)

node start commands used

RPC identity outputs (eth_chainId, eth_blockNumber)

balance verification outputs for funded accounts

multi-node parity outputs (if applicable)

Reference:

docs/AUDIT_READINESS_2.1.md

Change Control

Genesis changes are high impact.

Rules:

Any change to chainId (EVM), fixed supply, or allocation proportions is breaking and requires a new identity process.

Genesis MUST NOT be edited post-finalization for an environment. Changes MUST go through the release process and be documented.

Any change affecting behavior MUST be reflected in changelog and release records.

References:

docs/CHANGELOG_2.1.md

docs/RELEASE_PROCESS_2.1.md

References

docs/genesis/GENESIS_CHECKLIST_2.1.md
docs/genesis/PARAMETERS_REFERENCE_2.1.md
docs/genesis/ALLOC_POLICY_2.1.md
docs/STATE_MODEL_2.1.md
docs/RPC_SPEC_2.1.md
docs/rpc/RPC_COMPAT_MATRIX_2.1.md
docs/OPERATIONS_PLAYBOOK_2.1.md
docs/AUDIT_READINESS_2.1.md
docs/RELEASE_PROCESS_2.1.md
docs/CHANGELOG_2.1.md
docs/SECURITY_TRUST_MODEL_2.1.md
docs/THREAT_MODEL_2.1.md