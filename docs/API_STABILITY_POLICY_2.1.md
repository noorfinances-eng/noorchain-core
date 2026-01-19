NOORCHAIN 2.1 — API Stability Policy

Document ID: API_STABILITY_POLICY_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: NOORCHAIN 2.1 JSON-RPC API surface (node-facing), including leader/follower routing behavior where applicable.

1. Purpose

This document defines how NOORCHAIN 2.1 manages API stability over time for its JSON-RPC interface. It establishes:

Stability levels (Stable / Beta / Experimental / Deprecated)

Compatibility guarantees

Rules for changes and deprecations

Expectations for clients (wallets, tooling, dApps, operators)

Release-time gating requirements for API changes

The objective is to enable integration without ambiguity, while allowing controlled evolution of the node and its RPC surface.

2. Audience

This policy is written for:

Client developers (wallets, dApps, SDKs, scripts)

Tooling maintainers (Hardhat, viem-based apps, custom relayers)

Operators (leader/follower setups, monitoring, incident response)

Auditors evaluating interface stability and change control

3. Definitions
3.1 API Surface

For NOORCHAIN 2.1, “API” refers to the JSON-RPC methods and their observable behavior, including:

Request parameters and accepted encodings

Result payload shapes

Error model (codes/messages/data)

State semantics (e.g., block tags, persistence expectations)

Routing semantics in leader/follower mode (where relevant)

3.2 Normative References

The canonical method definitions are maintained in:

docs/RPC_SPEC_2.1.md (method behavior and semantics)

docs/rpc/RPC_ERROR_MODEL_2.1.md (errors and categorization)

docs/rpc/RPC_COMPAT_MATRIX_2.1.md (compatibility gates per client/tooling)

docs/RELEASE_PROCESS_2.1.md (tagging, releases, upgrade discipline)

docs/CHANGELOG_2.1.md (human-readable change history)

This policy defines the rules; the spec files define the exact shapes.

4. Stability Levels

Every JSON-RPC method in NOORCHAIN 2.1 is classified as one of the following.

4.1 Stable

A Stable method is safe for production integration.

Guarantees:

The method name remains unchanged.

The parameter schema remains backward compatible.

Result fields are not removed or retyped.

Semantics remain consistent across releases (see Section 6).

Any deprecation follows the rules in Section 7.

Stable methods may still improve performance, tighten validation, and add optional fields, provided compatibility is preserved.

4.2 Beta

A Beta method is intended for real use but may evolve.

Guarantees:

Best-effort compatibility within the same major line.

Breaking changes are allowed but must be documented clearly in the changelog and release notes.

A reasonable migration path is provided when feasible.

Beta methods are permitted in pilots, internal deployments, and controlled environments.

4.3 Experimental

An Experimental method is for development and diagnostics.

Characteristics:

No compatibility guarantees.

May change or be removed at any time.

Not recommended for external client dependencies.

Experimental methods should be treated as internal implementation details.

4.4 Deprecated

A Deprecated method is scheduled for removal.

Rules:

It remains available for a defined deprecation window.

It emits stable behavior during that window (no semantic drift).

It is removed only after the deprecation conditions are met.

5. Classification Rules
5.1 Default Classification

Methods explicitly documented in RPC_SPEC_2.1.md as Stable are Stable.

Methods documented as Beta are Beta.

Any RPC method not documented in RPC_SPEC_2.1.md is Experimental by default.

5.2 Promotion Criteria (Experimental → Beta → Stable)

A method may be promoted only if:

Its request/response schema is fully documented.

The error model is defined (including validation errors).

Compatibility gates are specified in RPC_COMPAT_MATRIX_2.1.md.

It passes the release gating process (Section 8).

6. Compatibility Guarantees

This section defines what “backward compatible” means in practice.

6.1 Backward-Compatible Changes (Allowed for Stable)

The following are compatible:

Adding new optional fields to result objects.

Adding new optional parameters (defaulting behavior must preserve old semantics).

Expanding accepted input formats when unambiguous (e.g., allowing additional encodings).

Fixing incorrect behavior when it was clearly a bug and the corrected behavior matches the written spec.

6.2 Breaking Changes (Not Allowed for Stable Without Deprecation)

The following are breaking:

Removing a method or renaming it.

Removing, retyping, or repurposing a parameter.

Changing a result field type or meaning.

Changing default behavior for omitted parameters.

Altering block-tag interpretation (latest, pending, explicit heights) in a way that changes results for existing calls.

Changing error codes/categories in a way that breaks client handling.

If a breaking change is required, the method must be deprecated first (Section 7) or a new method must be introduced.

6.3 Semantic Stability (Stable Methods)

For Stable methods, semantics must remain consistent across releases for:

Account reads (balance/nonce/code/storage)

Block/transaction/receipt retrieval semantics

“Latest” behavior relative to finalized/mined blocks in the node’s model

Leader/follower routing semantics when FollowRPC is configured (if applicable)

If a semantic gap exists between spec and implementation, the spec must be updated first and the change must be treated as a controlled evolution (never as an unannounced behavior shift).

7. Deprecation and Removal Policy
7.1 Deprecation Requirements

A method can be marked Deprecated only if:

A replacement exists (new method, new parameter, or clarified behavior).

The migration path is documented.

The deprecation is recorded in CHANGELOG_2.1.md.

7.2 Deprecation Window

For Stable methods, the minimum window is:

At least two tagged releases, and

At least 90 calendar days from the first release that marks it Deprecated,

whichever is longer.

If the network is in a controlled environment and operators explicitly accept a shorter window, this must be documented as an exception in the release notes (still discouraged).

7.3 Removal Conditions

A Deprecated method may be removed only after:

The deprecation window has elapsed, and

Compatibility gates confirm major client/tooling migration (as defined in RPC_COMPAT_MATRIX_2.1.md), and

The removal is recorded prominently in the changelog.

8. Release Gating for API Changes

No release that modifies Stable or Beta behavior is considered acceptable unless it passes the following gates.

8.1 Required Artifacts

Update RPC_SPEC_2.1.md (method semantics) when behavior changes.

Update RPC_ERROR_MODEL_2.1.md if error handling changes.

Update RPC_COMPAT_MATRIX_2.1.md if client coverage changes.

Update CHANGELOG_2.1.md with an explicit entry.

8.2 Required Validation Gates

At minimum:

Schema validation gate: responses match documented shapes for the covered methods.

Behavior gate: key semantics validated (e.g., state reads match persisted world-state; receipts are stable).

Leader/follower gate: when follower mode is used, routing behavior matches the spec for the covered methods.

Regression gate: previously stable methods remain stable under the same inputs.

Where possible, gates should be automated. If automation is not available, a runbook-driven manual gate must be recorded with evidence.

9. Client Expectations
9.1 Wallet and Tooling Compatibility

NOORCHAIN 2.1 aims for operational compatibility with standard Ethereum tooling where explicitly stated in the RPC spec. Clients must assume:

Stable methods adhere to the documented behavior.

Beta methods may change; clients should pin versions or tolerate changes.

Experimental methods are not safe dependencies.

9.2 Error Handling Expectations

Clients should implement robust error handling:

Validate JSON-RPC envelope fields (jsonrpc, id, result/error).

Handle errors according to RPC_ERROR_MODEL_2.1.md.

Never rely on free-form error strings as stable identifiers.

9.3 “Unknown Field” Handling

Clients must tolerate additional fields in result objects for Stable methods. New fields may appear over time.

10. Leader/Follower Semantics

When NOORCHAIN 2.1 is deployed in a leader/follower configuration, the RPC spec may define routing behavior for certain methods (especially world-state reads).

Policy:

Routing behavior is part of API semantics when documented.

If a follower proxies a method to the leader, this must be transparent to the client (result equivalence is the expected property).

Any method whose correctness depends on authoritative state must either:

be routed, or

be explicitly documented as “leader-only” or “not supported on follower.”

The exact routing rules are normative in RPC_SPEC_2.1.md.

11. Security and Abuse Considerations

Stability does not override safety:

Methods may enforce stricter limits over time (rate limits, payload size caps, parameter validation) if required to prevent abuse or protect node stability.

Such changes must be documented and treated as API behavior changes (gated via Section 8), especially if they affect legitimate clients.

12. Exception Handling

Exceptional breaking changes to Stable methods are strongly discouraged. If an exception is required (security incident, critical correctness flaw), the release must include:

A written justification

A migration guide

Explicit version pinning recommendations for clients

A post-incident review entry (see incident governance documents)

13. Governance of This Policy

Changes to this policy are treated as governance-level changes:

Any modification must be recorded in CHANGELOG_2.1.md.

The document version must be incremented.

Related documents (release process, threat model, trust model) should be checked for alignment.

14. Quick Summary (Non-Normative)

Stable methods: no breaking changes without deprecation.

Beta methods: can change, but must be documented clearly.

Experimental: no guarantees.

Deprecation: minimum two releases and 90 days (whichever longer).

Every meaningful API change must update spec + changelog and pass compatibility gates.

15. References

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/RELEASE_PROCESS_2.1.md

docs/CHANGELOG_2.1.md