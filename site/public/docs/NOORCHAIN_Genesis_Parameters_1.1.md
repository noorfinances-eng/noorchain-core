NOORCHAIN_Genesis_Parameters_1.1.md 
NOORCHAIN — Genesis Parameters
Version 1.1
Last Updated: 10.12.2025

1. Purpose of This Document

This document defines the technical and economic parameters included at mainnet genesis.

It ensures consistency between:

core protocol configuration,

initial PoSS parameters,

governance and legal constraints,

Swiss Legal Light CH compliance,

the 5 / 5 / 5 / 5 / 80 economic allocation model.

These rules form part of the immutable Genesis Pack.

In case of conflict, the Foundation Statutes and Legal Architecture documents prevail.

2. Core Chain Parameters (Cosmos / Ethermint)
2.1 Chain Identity

Chain Name: noorchain

Chain ID: noorchain-1

2.2 Denomination

Base Denom: unur

Display Denom: NUR

Decimals: 18

2.3 Block Timing

Block Time Target: 5s

Max Block Gas: 40,000,000

EVM Gas Adjustment: enabled

2.4 Governance Parameters

Deposit Requirement: 10,000 NUR

Voting Period: 5 days

Quorum: 33%

Threshold: 50% + 1

Veto Threshold: 33%

Governance cannot introduce financial mechanisms or modify Legal Light CH restrictions.

3. PoSS Parameters (Genesis)

These parameters initialize the Proof of Signal Social module.

They may be adjusted later by governance, except for structural invariants.

Parameter	Value
PoSSEnabled	false
BaseReward	1 unur (non-financial social unit)
WeightMicroDonation	2
WeightParticipation	1
WeightContent	3
WeightCCN	4
MaxSignalsPerDay	10
MaxSignalsPerCuratorPerDay	20
MaxRewardPerDay	100 unur (non-financial social cap)
HalvingPeriodYears	8
PoSSReserveDenom	unur
Structural Invariants (Non-Modifiable)

PoSS reward split: 70% participant / 30% curator

Total supply is immutable

No inflation beyond PoSS reserve

HalvingPeriodYears is immutable

Governance cannot mint tokens

PoSS rewards represent validated social signals, not financial compensation or yield.

4. Economic Parameters

Total Supply: 299,792,458 NUR

Inflation: 0% permanent

Genesis distribution applies the 5 / 5 / 5 / 5 / 80 model:

Item	Percentage	Notes
Foundation	5%	Multi-sig (3/5)
Dev Sàrl	5%	Functional allocation
PoSS Stimulus	5%	Non-financial ecosystem incentives
Pre-Sale Pool	5%	Private, regulated, vested
PoSS Reserve	80%	Released only through PoSS module

These values are encoded directly in genesis.json.

Allocation pools cannot be used as collateral or financial guarantees.

5. Governance Limits (Hard-Encoded)

Governance cannot modify:

total supply

PoSS reward split

halving period

genesis allocation percentages

structure or size of the PoSS reserve

Legal Light CH restrictions

minting rules

6. EVM Genesis Configuration

Base Fee (EIP-1559): 0

Min Gas Price: 0

ChainConfig: Ethermint default (Shanghai compatible)

EVM StateDB: initialized empty

7. Staking Parameters

Unbonding Time: 21 days

Max Validators: 75

Commission Max Rate: 100%

Commission Max Change: 1% / day

Min Self-Delegation: 1 NUR

These parameters are technical and may be adjusted via governance, except where restricted by Legal Light CH.

8. Summary

These parameters collectively define NOORCHAIN’s mainnet behavior.

They guarantee:

deterministic and predictable chain operation,

non-custodial, non-financial compliance,

PoSS readiness for future activation,

stable Cosmos + Ethermint integration,

strict alignment with the Genesis Pack.

Version 1.1 is valid for mainnet genesis and may only be updated through governance where permitted
