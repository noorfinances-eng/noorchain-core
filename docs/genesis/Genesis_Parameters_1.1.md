NOORCHAIN — Genesis Parameters

Version 1.1
Last Updated: 2025-XX-XX

1. Purpose of This Document

This document defines all protocol parameters included at mainnet genesis.
It ensures consistency between:

the core protocol configuration

PoSS initial parameters

governance constraints

Legal Light CH compliance

the 5/5/5/5/80 economic model

These parameters are part of the immutable Genesis Pack.

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

2.4 Governance

Deposit Requirement: 10,000 NUR
Voting Period: 5 days
Quorum: 33%
Threshold: 50% + 1
Veto Threshold: 33%

3. PoSS Parameters (Genesis)

These parameters initialize the PoSS module at genesis.
They are adjustable by governance except structural rules.

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

Structural Invariants (non-modifiable):
– PoSS split = 70% participant / 30% curator
– No inflation beyond PoSS reserve
– Total supply fixed forever

4. Economic Parameters

Total Supply: 299,792,458 NUR

Following the 5/5/5/5/80 distribution model:

Item	Percentage	Notes
Foundation	5%	Multi-sig 3/5
Dev Sàrl	5%	Functional allocation
PoSS Stimulus	5%	Early ecosystem incentives
Pre-Sale Pool	5%	Optional, vesting required
PoSS Reserve	80%	Only mintable source

These values are included directly in genesis.json.

5. Governance Limits (Hard-Encoded)

The following limits cannot be changed:

Total supply

PoSS 70/30 ratio

Halving period (8 years)

Genesis allocation percentages

Reserve size

Legal Light restrictions

No discretionary minting

6. EVM Genesis Configuration

Base Fee (EIP-1559): 0
Min Gas Price: 0
ChainConfig: Ethermint default (compatible with Ethereum Shanghai fork)
EVM StateDB: initialized empty

7. Staking Parameters

Unbonding Time: 21 days
Max Validators: 75
Commission Max Rate: 100%
Commission Max Change: 1% / day
Min Self-Delegation: 1 NUR

8. Summary

These parameters define the initial shape of NOORCHAIN.
They guarantee:

deterministic chain behavior

Legal Light compliance

PoSS readiness

clean integration of governance and economics

stable EVM compatibility

Version 1.1 is valid for mainnet genesis and will only change through on-chain governance where permitted.
