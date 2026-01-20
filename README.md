# NOORCHAIN — A Human-Centric Blockchain System (2.1)

NOORCHAIN is a blockchain project built around a simple constraint: **technical systems should be verifiable, governed, and socially legible**, not only fast or profitable.

Version **2.1** is the current engineering track: a controlled, permissioned deployment of a sovereign EVM Layer-1 used to validate operations, governance procedures, and evidence production in real conditions.

---

## Why NOORCHAIN Exists

Blockchain systems solved a fundamental problem—**credible execution without a central operator**—but they often failed at something equally important: **legitimacy**.

In practice, many networks optimize for:
- speculative velocity,
- narrative-driven incentives,
- governance that is informal or captured,
- and metrics that do not map to real-world value.

NOORCHAIN is an attempt to build a system where legitimacy is not a marketing claim but a **product of process**:
- decisions are attributable,
- changes are controlled,
- and outcomes can be evidenced.

This repository exists to implement and document that approach.

---

## What NOORCHAIN Prioritizes

### 1) Evidence Over Narratives
NOORCHAIN treats “proof” as operational: reproducible states, deterministic procedures, and evidence packs that an independent operator can verify.

### 2) Controlled Governance (by Design)
NOORCHAIN 2.1 is operated as a **permissioned system** during its build-and-validate phases. This is not an ideology; it is a risk and quality control measure:
- controlled validator membership,
- explicit upgrade processes,
- incident handling discipline,
- and a clear separation between what is “ready” and what is still experimental.

### 3) Separation of Security and Social Value
NOORCHAIN separates:
- **consensus/security** (what makes the chain reliable),
from
- **social value and legitimacy signals** (how the system measures real contributions).

This separation prevents “social scoring” from becoming consensus power, and keeps security engineering independent from application-layer incentives.

---

## PoSS: Proof of Signal Social (Concept)

PoSS (Proof of Signal Social) is NOORCHAIN’s application-layer mechanism to represent and validate real-world actions through structured evidence and curator governance.

Key principles:
- PoSS is **not consensus**.
- PoSS is designed to be **auditable** and bounded by process.
- Curators validate “signals” and sign snapshots.
- The reward model follows a fixed split: **70% participants / 30% curators** (application-layer rule).

PoSS is intended to support pilots and controlled deployments where legitimacy is measurable, even when the environment is imperfect.

---

## NUR: Native Asset (Context)

NUR is the native asset of NOORCHAIN. Its role is to support network operations (e.g., transaction fees) and internal economic coordination where applicable.

This repository and documentation do not present NUR as an investment product and do not make promises about financial returns.

---

## The 2.1 Track (What This Repo Represents)

This repository is the **engineering and documentation base** for NOORCHAIN 2.1:

- a sovereign EVM L1 implementation operated in controlled environments,
- a set of operational procedures and runbooks to prevent drift,
- and a documentation set that defines governance, security posture, and evidence discipline.

The core objective is not “a demo chain”; the objective is **a mainnet-like operational system** where claims can be tied to verifiable outputs.

---

## Project Ethos (Non-Negotiables)

- **No yield / ROI promises.**
- **No custody representation.**
- **Controlled exposure by default** (security and privacy first).
- **Process discipline**: changes, releases, and incidents must be handled as first-class system events.
- **Legitimacy as an output**: a chain is not legitimate because it says so; it becomes legitimate because it can be independently checked.

---

## Where to Start (Minimal References)

If you need the canonical documentation entry points:

- `docs/NOORCHAIN_Index_2.1.md` (documentation map)
- `docs/ARCHITECTURE_2.1.md` (system overview)
- `docs/legal/LEGAL_LIGHT_POSTURE_2.1.md` (policy constraints)

---

## Disclaimer

NOORCHAIN 2.1 is operated in controlled environments. This repository and its contents are provided for engineering and validation purposes and do not constitute an investment offer, do not promise returns, and do not provide custody services.
