NOORCHAIN 2.1 — PHASE 7 (Mini-Explorer / Proof-of-Liveness)
Operational Reference Runbook — Frozen Baseline

Document Class: Operations (internal)
Scope: Production-controlled proof-of-liveness only
Public Surface: Single endpoint GET /liveness.json
Compliance Posture: Swiss Legal-Light CH (no custody, no yield, no financial services, no investment language)

0. Purpose and Non-Goals
Purpose

Maintain a minimal, controlled, and verifiable proof-of-liveness for the NOORCHAIN 2.1 network, using:

two VPS nodes (Leader / Follower),

private RPC access,

a single public status endpoint (/liveness.json) updated every 30 seconds.

Non-Goals

This baseline does not provide:

a public explorer,

public RPC,

transaction or address browsing,

performance claims,

availability guarantees beyond best-effort operations.

1. Baseline Invariants (Frozen)

These invariants define the frozen baseline. Any change to them is a scope change and requires an explicit decision.

1.1 Infrastructure

Two VPS instances:

Leader: block production and canonical reference

Follower: replication / redundancy / observation

P2P: enabled and stable between Leader and Follower.

RPC: private only (restricted access).

1.2 Public Exposure (Hard Limit)

Single public endpoint: GET /liveness.json

Update cadence: 30 seconds

No other public endpoints (including indirect exposure via site assets or APIs).

1.3 Content Policy (Public)

Public payload must remain minimal and non-sensitive:

chain identifier,

leader height,

observed timestamp,

uptime seconds (as computed by the reporter).

No:

peer lists,

IPs,

ports,

node IDs,

RPC URLs,

internal metrics,

transaction or address data.

2. Roles and Responsibilities
2.1 Leader Node

Produces blocks.

Serves as the canonical reference for the liveness reporter.

RPC accessible only through private controls.

2.2 Follower Node

Maintains P2P connection to Leader.

Maintains RPC privately.

Provides redundancy and operational confidence, without public exposure.

2.3 Liveness Reporter

Queries the Leader privately.

Writes /liveness.json to the public web surface.

Runs on a fixed schedule (30 seconds).

3. Operational Boundaries
3.1 Allowed Changes (During Freeze)

Only the following actions are permitted without a scope change:

restart services to recover from failure,

redeploy the same build to restore baseline,

apply emergency security fixes (limited to preventing compromise or outage),

rotate credentials/keys strictly for security.

3.2 Prohibited Changes (During Freeze)

Not permitted under the frozen baseline:

adding any new public endpoints,

exposing RPC publicly (directly or via proxy),

expanding liveness.json content beyond the minimal payload,

publishing operational identifiers (node IDs, peers, ports, IPs),

enabling “explorer-like” UI features on the public site,

enabling public logs, traces, or metrics.

4. Evidence and Verification (What “OK” Means)

The baseline is considered healthy when all conditions below hold:

4.1 Network Health

Leader and Follower are running.

P2P is established between them.

Leader height increases over time.

4.2 Public Liveness

GET /liveness.json is reachable publicly.

JSON is valid.

observed_at updates every ~30 seconds.

leader_height is non-decreasing across updates.

5. Incident Handling (Freeze-Compatible)
5.1 If Public Endpoint Stops Updating

Confirm the reporter process is running.

Confirm it can still query the Leader privately.

Restart reporter only (preferred) before any node restarts.

5.2 If Leader Stops Advancing

Confirm Leader process health.

Confirm disk and resource headroom.

Restart Leader if required.

Confirm Follower reconnects and the public liveness resumes normal updates.

5.3 If Follower Disconnects

Treat as redundancy degradation, not a public incident (unless it affects liveness).

Restore P2P connectivity and confirm steady state.

6. Change Control (Freeze Governance)

A change is considered out of baseline if it modifies any of:

number of VPS nodes,

public surface area,

cadence or content of liveness.json,

public hostname / routing used for liveness delivery,

RPC exposure rules.

Out-of-baseline changes require an explicit decision and a documented update to this runbook.

7. Required Recorded References (To Complete)

The frozen baseline is not formally complete until the following references are recorded here:

Chain ID: ______________

Leader observed RPC (private): ______________

Follower observed RPC (private): ______________

Public liveness URL: ______________

Code reference (branch + commit hash + tag): ______________

Reporter schedule definition (cron/systemd interval): ______________

Until these are filled, the runbook remains operationally incomplete.

8. Public Language (Site Integration Constraint)

Any site-visible mention of this proof-of-liveness must:

describe it as a technical status signal only,

avoid financial framing,

avoid claims of permanence, performance, or guarantees,

avoid implying public participation or public RPC access.

9. Freeze Statement

As of the acceptance of this runbook, the Phase 7 proof-of-liveness baseline is considered frozen:

two-node VPS topology (Leader/Follower),

private RPC,

single public GET /liveness.json updating every 30 seconds,

no additional public surfaces.

Any deviation is a scope change and must be explicitly authorized.
