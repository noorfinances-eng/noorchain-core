NOORCHAIN_Genesis_Allocation_1.1.md
Genesis Pack — Phase 6.3
Version 1.1 — Last Updated: 10.12.2025

1. Purpose of This Document

This document defines the immutable economic allocation of the NOORCHAIN native token UNUR (denom: unur) at genesis.

It is part of the official Genesis Pack 1.1 and serves as:

a regulatory reference (Swiss Legal Light CH),

an institutional document for partners and curators,

a technical source for the genesis.json configuration,

a transparency guarantee for the public.

No governance mechanism, Foundation decision, upgrade, or protocol change may modify these allocations.

2. Total Supply (Immutable)

The total supply of NOORCHAIN at genesis is:

299,792,458 NUR

(symbolic reference to the speed of light: 299,792,458 m/s)

This supply:

is fixed forever,

cannot be increased,

cannot be modified by governance,

defines the upper limit for PoSS minting across all future halving cycles.

3. Allocation Model (Immutable)

The genesis distribution follows the permanent 5 / 5 / 5 / 5 / 80 structure:

Allocation Pool	Percentage	Amount (NUR)	Description
Foundation Reserve	5%	14,989,622.9	Governance, audits, public infrastructure, transparency
Noor Dev Sàrl	5%	14,989,622.9	R&D, integrations, infrastructure, tooling
PoSS Stimulus	5%	14,989,622.9	Early social pilots for NGOs, schools, curators
Pre-sale Pool (Optional)	5%	14,989,622.9	Strictly private, Swiss-regulated pre-mainnet financing
PoSS Mintable Reserve	80%	239,833,966.4	Long-term PoSS issuance through halving

Rounding differences are accepted. Exact values are encoded in the genesis file.

4. Allocation Rules & Restrictions
4.1 Immutable Allocation

The 5 / 5 / 5 / 5 / 80 model is a permanent, protocol-level rule.

It cannot be changed by:

governance

Foundation Board

multisig

smart contracts

upgrades or forks

4.2 No Reallocation Between Pools

Tokens assigned to a pool must never be:

transferred to another pool

merged or reduced

repurposed

used as financial guarantees

pledged or collateralized

Each pool has a strict mission and must remain isolated.

4.3 No Additional Minting

There is zero inflation.

Only the PoSS Mintable Reserve may introduce tokens into circulation — and only through:

PoSS validation,

halving schedule,

protocol-defined limits.

No module, no smart contract, and no governance vote may mint tokens outside the PoSS module.

5. Description of Each Allocation Pool
5.1 Foundation Reserve — 5%

Supports:

governance operations

legal compliance

audits & security

documentation & transparency

institutional partnerships

public-good development

Access requires:

multisig 3/5 approval

compliance verification

transparent public reporting

5.2 Noor Dev Sàrl (Development) — 5%

Purpose:

protocol development

tooling & integrations

research & innovation

infrastructure, node operations

long-term engineering

All funds must follow Swiss corporate and accounting rules.

5.3 PoSS Stimulus Pool — 5%

Supports:

NGOs

educators

associations

early curators

social pilots

onboarding programs

These funds encourage adoption without creating financial incentives.

5.4 Pre-sale Pool (Optional) — 5%

This pool is:

strictly private

reserved for Swiss institutions, family offices, impact funds

subject to vesting (recommended 6 + 18 months)

controlled by multisig

capped and regulated

Clarification (Legal Light CH):

This pool does not constitute a public sale, ICO, token offering, or investment product.
Any funds raised are received exclusively by Noor Dev Sàrl, not by the Foundation.

5.5 PoSS Mintable Reserve — 80%

The core reserve powering the NOORCHAIN social economy.

Released into circulation through:

PoSS validation

weighted social signals

halving every 8 years

protocol-defined daily limits

Duration:

Expected multi-decade issuance, depending on adoption and use.
This is a technical estimate, not a financial projection or guarantee.

This reserve may only be used by the PoSS module.

6. Addresses for Genesis (Phase 7)

The actual addresses for:

Foundation

Noor Dev Sàrl

PoSS Stimulus

Pre-sale

PoSS Reserve

will be added during Phase 7 — Pre-Mainnet, once:

the 5 real Bech32 addresses are generated,

they are synchronized across:

testnet/genesis.json

genesis_distribution.json

x/noorsignal/types/addresses.go

governance documents

Until Phase 7, this section remains a placeholder.

7. Governance Limits

Governance cannot:

alter supply

change allocation percentages

mint outside PoSS

extend the PoSS reserve

reassign pools

operate liquidity tools

create financial products

Governance can adjust PoSS parameters only.

8. Integration in Genesis Pack

This Allocation Document is used to generate:

genesis.json

genesis_distribution.json

PoSS reserve parameters

governance initialization

It is also integrated into:

Whitepaper 1.1

Compliance Framework

Legal Notices

Public website (Phase 6.6)

9. Summary Table
Pool	Percentage	Notes
Foundation	5%	Governance, audits, transparency
Dev Sàrl	5%	R&D, infrastructure
PoSS Stimulus	5%	NGOs, schools, associations
Pre-sale (Optional)	5%	Private, Swiss-only, vested
PoSS Reserve	80%	Social reward reserve

Total = 100% (immutable)

10. Hierarchy of Legal Documents (NEW)

In case of conflict:

Foundation Statutes

Legal Architecture

Legal Light Framework

Multisig Charter

Compliance Framework

Genesis Allocation (this document)

11. Signature

Prepared by:
NOORCHAIN Foundation — Genesis & Governance Division
Version 1.1
