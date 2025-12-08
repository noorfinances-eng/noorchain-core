NOORCHAIN_Genesis_Parameters_1.1.md
NOORCHAIN — Genesis Parameters
Version 1.1
Last Updated: 2025-XX-XX
________________________________________
1. Purpose of This Document
This document defines all technical and economic parameters included at mainnet genesis.
It ensures full consistency between:
•	core protocol configuration,
•	initial PoSS parameters,
•	governance and legal constraints,
•	Swiss Legal Light CH compliance,
•	the 5/5/5/5/80 economic allocation model.
These rules are part of the immutable Genesis Pack.
________________________________________
2. Core Chain Parameters (Cosmos / Ethermint)
2.1 Chain Identity
•	Chain Name: noorchain
•	Chain ID: noorchain-1
2.2 Denomination
•	Base Denom: unur
•	Display Denom: NUR
•	Decimals: 18
2.3 Block Timing
•	Block Time Target: 5s
•	Max Block Gas: 40,000,000
•	EVM Gas Adjustment: enabled
2.4 Governance
•	Deposit Requirement: 10,000 NUR
•	Voting Period: 5 days
•	Quorum: 33%
•	Threshold: 50% + 1
•	Veto Threshold: 33%
________________________________________
3. PoSS Parameters (Genesis)
These parameters initialize the Proof of Signal Social (PoSS) module at genesis.
They may be adjusted later by governance, except for structural invariants.
Parameter	Value
PoSSEnabled	false
BaseReward	1 unur
WeightMicroDonation	2
WeightParticipation	1
WeightContent	3
WeightCCN	4
MaxSignalsPerDay	10
MaxSignalsPerCuratorPerDay	20
MaxRewardPerDay	100 unur
HalvingPeriodYears	8
PoSSReserveDenom	unur
Structural Invariants (non-modifiable)
•	PoSS reward split: 70% participant / 30% curator
•	No inflation beyond the PoSS reserve
•	Total supply is fixed forever
________________________________________
4. Economic Parameters
•	Total Supply: 299,792,458 NUR
•	Inflation: 0% (permanent)
The genesis distribution applies the 5 / 5 / 5 / 5 / 80 model:
Item	Percentage	Notes
Foundation	5%	Multi-sig (3/5)
Dev Sàrl	5%	Functional allocation
PoSS Stimulus	5%	Early ecosystem incentives
Pre-Sale Pool	5%	Optional, vesting required
PoSS Reserve	80%	Mintable over decades
These values are encoded directly in genesis.json.
________________________________________
5. Governance Limits (Hard-Encoded)
The following limits cannot be changed by governance:
•	Total supply
•	PoSS reward split (70/30)
•	Halving period (8 years)
•	Genesis allocation percentages
•	PoSS reserve size
•	Legal Light CH restrictions
•	No discretionary minting
________________________________________
6. EVM Genesis Configuration
•	Base Fee (EIP-1559): 0
•	Min Gas Price: 0
•	ChainConfig: Ethermint default (Shanghai compatible)
•	EVM StateDB: initialized empty
________________________________________
7. Staking Parameters
•	Unbonding Time: 21 days
•	Max Validators: 75
•	Commission Max Rate: 100%
•	Commission Max Change: 1% / day
•	Min Self-Delegation: 1 NUR
________________________________________
8. Summary
These parameters define the initial behavior of the NOORCHAIN network.
They guarantee:
•	deterministic and predictable chain behavior,
•	full compliance with Legal Light CH,
•	PoSS readiness for future activation,
•	clean integration between governance and economics,
•	stable EVM compatibility.
Version 1.1 is valid for mainnet genesis and may only be updated through on-chain governance where allowed.

