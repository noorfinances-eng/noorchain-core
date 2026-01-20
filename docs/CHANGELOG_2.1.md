NOORCHAIN 2.1 — Changelog

Document ID: CHANGELOG_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: NOORCHAIN 2.1 core node, RPC surface, state model, and operational runbooks.

0. Conventions
0.1 Versioning

NOORCHAIN 2.1 uses tagged releases as the canonical source of version identity.

Each release is represented by a Git tag.

This changelog records auditable summaries of changes per tag or milestone tag.

If a change is not in a tag, it is not considered a released change.

0.2 Change Categories

Entries use the following categories:

Added — new capability

Changed — modification to existing behavior

Fixed — bug fix with compatibility intent

Deprecated — scheduled for removal

Removed — removed capability

Security — security-relevant changes

Ops — operations/runbook/deployment changes

Docs — documentation changes

0.3 Compatibility Markers

Each entry may include:

[BC] Backward compatible

[BREAKING] Breaking change (must include migration note and reference stability policy)

[BEHAVIOR] Semantic behavior change (even if schema unchanged)

[RPC] JSON-RPC surface impact

[STATE] State/persistence impact

[NET] Networking/consensus impact

1. Released Tags and Milestones

Note: This changelog captures major NOORCHAIN 2.1 milestones and documentation state as of January 2026. For precise code diffs, use the referenced tags/commits in the repository history.

2026-01 — docs(2.1): API stability policy + core specs alignment

Commit: 9067200
Categories: Docs
Notes: Documentation alignment update covering stability policy and core specs.

Added: docs/API_STABILITY_POLICY_2.1.md defining method stability levels, deprecation rules, and release gating.

Changed: Updates/alignments to core documentation files:

docs/ARCHITECTURE_2.1.md

docs/RPC_SPEC_2.1.md

docs/STATE_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md
Compatibility: [BC] (documentation-only).
Impact: None on runtime behavior.

2026-01 — M12.3 VALIDATED — follower RPC world-state reads routed via leader

Tag: (milestone tag per repository)
Categories: Changed, Fixed, Ops
Markers: [RPC] [BEHAVIOR]

Changed: In follower mode (FollowRPC configured), world-state reads are routed to the leader for correctness (as specified).

Validated: eth_getTransactionCount and eth_getBalance parity between leader and follower.
Compatibility: [BC] for clients; behavior becomes more consistent in follower deployments.

2026-01 — M12.2 WORLD STATE GETH integration (StateDB + triedb) — head root persistence

Tag: (milestone tag per repository)
Categories: Added, Changed
Markers: [STATE] [RPC]

Added: Integration of geth-style world-state database store (isolated under node data-dir).

Added: State root tracking and persistence for head state.

Changed: eth_getTransactionCount and eth_getBalance backed by StateDB where implemented.
Compatibility: [BC] for callers; semantics become real-state-derived rather than stubbed.

2025-12 — M10-MAINNETLIKE-STABLE — multi-node mainnet-like pack validated

Tag: M10-MAINNETLIKE-STABLE
Categories: Added, Changed, Ops
Markers: [NET] [RPC]

Added: Two-node leader/follower runbook and validation discipline.

Changed: Health endpoint configuration made explicit and safe via -health-addr.

Ops: Introduced stable operations runbook for multi-node local deployments.
Compatibility: [BC] for clients; improved operator safety.

2025-12 — M9-CONTRACTS-EXECUTION-STABLE — minimal contract execution at mining + receipts persistence

Tag: M9-CONTRACTS-EXECUTION-STABLE
Categories: Added, Changed
Markers: [STATE] [RPC]

Added: Minimal execution hook enabling contract execution at mining time (scope-limited).

Added: Receipts persisted under deterministic keys and served by eth_getTransactionReceipt.

Validated: submitSnapshot workflow produces receipts and persisted PoSS state.
Compatibility: [BC] for existing RPC usage; adds expected receipt availability.

2025-12 — M8-RPC-TXPATH-STABLE — tx path completion

Tag: M8-RPC-TXPATH-STABLE
Categories: Added, Changed
Markers: [RPC]

Added: Full transaction path support:

eth_sendRawTransaction

eth_getTransactionByHash

eth_getTransactionReceipt

Validated: Inclusion/mining minimal active; block numbers non-null; PoSSRegistry read path consistent.
Compatibility: [BC] for clients; enables standard tooling workflows.

2025-12 — M7-PERSIST-STABLE — PoSS snapshot persistence + RPC wiring

Tag: M7-PERSIST-STABLE
Categories: Added, Changed
Markers: [STATE] [RPC]

Added: Persistent PoSS snapshots in LevelDB.

Changed: eth_call read methods return stable values across restarts for PoSS registry reads.
Compatibility: [BC]; persistence correctness improved.

2025-12 — M6A3-STABLE — PoSS v0 contracts + dev JSON-RPC shim (batch-ready)

Tag: M6A3-STABLE
Categories: Added
Markers: [RPC]

Added: PoSS v0 Solidity contracts (CuratorSet, PoSSRegistry).

Added: Deployment and snapshot submit scripts (Hardhat/ethers).

Added: Dev-only JSON-RPC shim enabling tooling compatibility in early phases.
Compatibility: Experimental-to-Beta evolution groundwork; later superseded by real RPC paths.

2025-12 — M5-STABLE — persistent P2P sessions + boot peers

Tag: (milestone tag per repository)
Categories: Added, Changed, Ops
Markers: [NET]

Added: -boot-peers flag and persistent peer sessions.

Changed: Node startup can bootstrap peers directly; stability gates documented.
Compatibility: [BC]; operational behavior improved.

2. Deprecations

No formal deprecations recorded in this changelog yet.

Deprecations must follow:

docs/API_STABILITY_POLICY_2.1.md

3. Breaking Changes

No breaking changes recorded in this changelog yet.

Any breaking change must include:

migration guidance

explicit [BREAKING] marker

stability policy reference

4. Security Notes

Security-relevant changes must be recorded here with sufficient operator impact context and references to:

docs/SECURITY_TRUST_MODEL_2.1.md

docs/THREAT_MODEL_2.1.md

No explicit security advisories recorded in this changelog yet.

5. Operational Notes

Operators should rely on runbooks for canonical procedures:

docs/OPERATIONS_PLAYBOOK_2.1.md (when filled)

Existing runbooks under docs/RUNBOOK-* and docs/ops/*