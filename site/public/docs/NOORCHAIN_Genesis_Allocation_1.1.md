NOORCHAIN_Genesis_Allocation_1.1.md
Genesis Pack — Phase 6.3
Version 1.1 — Last Updated: 2025-XX-XX
________________________________________
1. Purpose of This Document
This document defines the immutable economic allocation of the NOORCHAIN native token UNUR (denom: unur) at genesis.
It is part of the official Genesis Pack 1.1 and serves as:
•	a regulatory reference (Legal Light CH),
•	an institutional document for partners and curators,
•	a technical source for the genesis.json configuration,
•	a transparency guarantee for the public.
No future governance process may modify these allocations.
________________________________________
2. Total Supply (Immutable)
The total supply of NOORCHAIN at genesis is:
299,792,458 NUR
(symbolic reference to the speed of light: 299,792,458 m/s)
This supply:
•	is fixed forever,
•	can never be increased,
•	cannot be modified through governance,
•	defines the upper limit of token circulation across all future halving cycles.
________________________________________
3. Allocation Model (Immutable)
The genesis distribution follows the permanent 5 / 5 / 5 / 5 / 80 structure:
Allocation Pool	Percentage	Amount (NUR)	Description
Foundation Reserve	5%	14,989,622.9	Supports governance, audits, public infrastructure, transparency
Noor Dev Sàrl (Development)	5%	14,989,622.9	R&D, integrations, infrastructure, tooling
PoSS Stimulus Pool	5%	14,989,622.9	Early incentives for NGOs, schools, curators, pilots
Pre-sale Pool (Optional)	5%	14,989,622.9	Strictly regulated private pre-mainnet funding
PoSS Mintable Reserve	80%	239,833,966.4	Long-term PoSS issuance over 25–40 years
Round-off error is accepted in denominated decimals; exact values are encoded in the genesis file.
________________________________________
4. Allocation Rules & Restrictions
4.1 Immutable Allocation
The 5/5/5/5/80 model is a permanent rule.
It cannot be changed by:
•	governance,
•	Foundation Board,
•	multisig,
•	smart contracts,
•	or any future protocol upgrade.
4.2 No Reallocation Between Pools
Tokens allocated to a specific pool cannot be:
•	transferred to another pool,
•	rebalanced,
•	repurposed,
•	merged or reduced.
Each pool has a defined mission and must remain isolated.
4.3 No Additional Minting
There is zero inflation.
Only the 80% PoSS Reserve may produce tokens, and only through:
•	the halving schedule,
•	the PoSS module,
•	validated actions.
________________________________________
5. Description of Each Allocation Pool
________________________________________
5.1 Foundation Reserve — 5%
Purpose:
•	governance operations
•	legal compliance
•	audits & security
•	documentation & transparency
•	institutional partnerships
•	public-good development
Access requires:
•	multisig 3/5 approval,
•	compliance review,
•	public reporting.
________________________________________
5.2 Noor Dev Sàrl (Development) — 5%
Purpose:
•	protocol development
•	tooling & integrations
•	research & innovation
•	infrastructure & node operations
•	maintenance and long-term engineering
All funds must follow Swiss corporate and accounting rules.
________________________________________
5.3 PoSS Stimulus Pool — 5%
Supports:
•	NGOs
•	educators
•	associations
•	early curators
•	institutional pilots
•	social onboarding programs
These funds bootstrap early adoption without creating financial incentives.
________________________________________
5.4 Pre-sale Pool (Optional) — 5%
This pool is:
•	strictly private
•	only accessible to Swiss institutions, family offices, impact funds
•	subject to vesting (recommended 6 + 18 months)
•	controlled by multisig
•	capped and regulated
This pool does not represent a public offering.
________________________________________
5.5 PoSS Mintable Reserve — 80%
The core reserve powering the NOORCHAIN social economy.
Released into circulation through:
•	PoSS validation
•	weighted signals
•	halving every 8 years
•	daily limit enforcement
Expected duration: 25–40 years depending on adoption.
This reserve may only be used by the PoSS module.
________________________________________
6. Addresses for Genesis (Phase 7)
The actual addresses for:
•	Foundation
•	Dev Sàrl
•	PoSS Stimulus
•	Pre-sale
•	PoSS Reserve
will be added during Phase 7 — Pre-Mainnet, once:
•	the 5 real bech32 addresses are generated,
•	they are cross-synchronized through:
o	testnet/genesis.json
o	genesis_distribution.json
o	x/noorsignal/types/addresses.go
o	governance documents.
Until Phase 7, this section remains a placeholder.
________________________________________
7. Governance Limits
Governance cannot:
•	alter supply
•	modify allocation percentages
•	mint outside PoSS rules
•	extend the PoSS reserve
•	operate liquidity or investment tools
•	reassign pools
Governance can only adjust PoSS parameters, not the foundation of economics.
________________________________________
8. Integration in Genesis Pack
This Allocation Document is used directly to generate:
•	genesis.json
•	genesis_distribution.json
•	PoSS reserve parameters
•	Governance module initialization
It is also integrated into:
•	Whitepaper 1.1
•	Compliance Framework
•	Legal Notices
•	Public website (Phase 6.6)
________________________________________
9. Summary Table
Pool	Percentage	Notes
Foundation	5%	Public-good operations, audits, governance
Dev Sàrl	5%	R&D, infrastructure, tools
PoSS Stimulus	5%	NGOs, schools, associations
Pre-sale (Optional)	5%	Strictly private, Swiss-only, vested
PoSS Reserve	80%	Multi-decade social reward pool
Total = 100% (immutable)
________________________________________
10. Signature
Prepared by:
NOORCHAIN Foundation
Genesis & Governance Division
Version 1.1

