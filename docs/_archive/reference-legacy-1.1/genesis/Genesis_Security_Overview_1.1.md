NOORCHAIN — Genesis Security Overview

Version 1.1
Last Updated: 2025-XX-XX

1. Purpose of This Document

The Genesis Security Overview defines all security guarantees, restrictions, and invariants enforced at the genesis block of NOORCHAIN.

It ensures:

legal and economic safety

immutability of critical rules

correct initialization of PoSS

protection against malicious governance

long-term resilience for mainnet

This document is part of the Genesis Pack.

2. Security Principles at Genesis

NOORCHAIN follows five foundational security principles:

Immutability of core supply rules

No discretionary minting

Strict separation of governance powers

No custody of user funds by Foundation or Sàrl

Legal Light CH compliance

These principles are non-negotiable and embedded into documentation, governance, and genesis files.

3. Fixed Supply Guarantee

The following invariants are guaranteed at genesis:

Total supply = 299,792,458 NUR

No inflation except PoSS Reserve minting

PoSS Reserve capped at 80% of supply

Genesis allocations (5/5/5/5/80) cannot be altered

No governance proposal can modify:

the global cap

the genesis supply distribution

the existence of the PoSS Reserve

This prevents economic manipulation or dilution.

4. Governance Security

Governance is initialized with strict boundaries:

Foundation Board

Administrative and oversight powers only.

Multi-sig Committee (3/5)

Responsible for executing decisions, not creating supply.

No single actor can:

mint tokens

bypass PoSS

reassign genesis allocations

override legal constraints

Any governance proposal violating Legal Light CH is invalid by design.

5. PoSS Module Security

PoSS is intentionally disabled at genesis (PoSSEnabled = false) for safety.

Security protections:

No automatic minting before activation

Minting only from the PoSS Reserve

70/30 reward split locked forever

Halving schedule immutable

Daily limits configurable but capped

No external oracle dependency

This ensures PoSS cannot be abused or exploited during the early stages of mainnet.

6. EVM Security Measures

Ethermint genesis integrates:

EIP-1559 base fee disabled initially

EVM state empty at genesis

No pre-deployed contracts

No guaranteed gas subsidies

Full compatibility with Ethereum signing rules

Replay protection via ChainID

This eliminates attack vectors related to preloaded contracts or unexpected gas behavior.

7. Staking Security

The staking module security is ensured by:

Minimum self-delegation

Commission limits

Validator set capped

21-day unbonding period

Governance separation

These prevent sudden validator takeovers or rapid centralization.

8. Legal Compliance Enforcement

NOORCHAIN genesis respects Legal Light CH:

No custody

No promised returns

No Fiat gateway

No investment language

PoSS = social reward mechanism, not a yield instrument

Any behavior violating Swiss non-custodial rules is structurally impossible.

9. Attack Surface Reduction at Genesis

The following are intentionally disabled or minimized:

PoSS automatic minting

Preloaded smart contracts

EVM precompiles beyond standards

Treasury-like mechanisms

Unlimited governance changes

Rapid validator set expansion

This ensures a safe, conservative mainnet launch.

10. Summary of Genesis Security Guarantees

The genesis block of NOORCHAIN guarantees:

fixed supply

immutable economic rules

PoSS locked behind governance

strong staking security

safe EVM initialization

strict governance boundaries

Legal Light CH compliance

Version 1.1 is adopted as part of the Genesis Pack.
