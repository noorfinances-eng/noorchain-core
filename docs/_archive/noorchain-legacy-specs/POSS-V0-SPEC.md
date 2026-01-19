# PoSS v0 Specification — NOORCHAIN 2.1

## Overview
PoSS (Proof of Signal Social) v0 is an application-layer mechanism for value signaling and governance.
It does not participate in consensus and has no impact on block production or finality.

## Design Principles
- Strict separation: consensus = security, PoSS = value & governance
- Non-custodial
- Snapshot-based and immutable
- Threshold-signed by authorized curators
- Minimal on-chain footprint
- No yield promise, no financial commitment

## Architecture
PoSS v0 consists of:
- An off-chain engine producing canonical snapshots
- Authorized curators signing snapshot hashes
- On-chain contracts registering snapshots and verifying signatures

## On-chain Contracts
### CuratorSet
- Maintains the list of authorized curators
- Defines the signature threshold
- Admin-controlled (permissioned, governance-ready)

### PoSSRegistry
- Registers immutable snapshots
- Verifies curator threshold signatures
- Stores only snapshot hash and minimal metadata

## Snapshot Model
- Snapshots are generated off-chain
- Each snapshot is identified by a canonical hash
- The hash is signed by curators
- The snapshot is immutable once accepted on-chain

## Manifest v1
- Canonical JSON
- Ordered keys
- UTF-8, minified
- Hash = keccak256(minified JSON bytes)

## Rewards Philosophy
PoSS v0 preserves the NOORCHAIN 1 & 2 reward philosophy:
- 70% participants
- 30% curators

The rewards model is declared, not enforced on-chain.
No distribution logic is part of consensus.

## Security Properties
- Replay protection via chainId
- Threshold signature enforcement
- Duplicate snapshot rejection
- Full auditability from manifest + signatures

## Non-goals
- No validator selection
- No staking or slashing
- No automatic payouts
- No financial guarantees
