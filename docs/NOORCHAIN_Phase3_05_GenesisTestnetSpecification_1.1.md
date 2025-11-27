NOORCHAIN 1.0 — Phase 3.05
Genesis Testnet 1.0 Specification
Version 1.1 — Official Document

Purpose of this document
Define all genesis parameters required to launch NOORCHAIN Testnet 1.0, including:

the 5 official genesis addresses

initial token distribution

PoSS reserve setup

fee, gas & EVM settings

staking & validator parameters

network configuration

This is the canonical reference before generating genesis.json.

1. Genesis Philosophy

NOORCHAIN Testnet 1.0 is designed to:

follow the economic model (5% / 5% / 5% / 5% / 80%)

reflect the final mainnet structure

allow full testing of Cosmos, EVM, and PoSS

run on CometBFT

integrate Ethermint JSON-RPC

be fully deterministic and auditable

Everything defined here is mandatory for Testnet 1.0.

2. Genesis Addresses (Official 5-Key Structure)

Testnet requires 5 bech32 addresses representing the economic allocation model:

Allocation	Description	Genesis Role
5% – Foundation Wallet	Public foundation funding	receives 5% of total supply
5% – Dev Sàrl Wallet	Founder operational wallet	receives 5%
5% – PoSS Stimulus Pool	Early rewards + adoption boosts	receives 5%
5% – Optional Pre-Sale Wallet	Reserved for regulated private investors	receives 5%
80% – PoSS Mintable Reserve	Source of all PoSS emissions	receives 80%
2.1. Requirements for the addresses

must be bech32 with prefix noor1

must be generated fresh and manually recorded

must be stored later inside x/noorsignal/types/addresses.go

must be included in Testnet’s genesis.json and genesis_distribution.json

2.2. When we generate them

During Phase 4C: Testnet 1.0, step:

Generate 5 real bech32 addresses (noor1…)


This document establishes their roles, not their values yet.

3. Total Supply & Allocation

NOORCHAIN’s total supply is fixed:

299,792,458 NUR

3.1. Token denom
base denom: unur
display: NUR
decimals: 18

3.2. Genesis token distribution
Pool	%	Amount (NUR)	Purpose
Foundation	5%	14,989,622.90	public-interest operations
Dev Sàrl	5%	14,989,622.90	development & infra
PoSS Stimulus	5%	14,989,622.90	boost early adoption
Pre-Sale	5%	14,989,622.90	optional investors (vesting)
PoSS Reserve	80%	239,833,966.40	reward engine

Amounts are exact but not stored with decimals in genesis; Genesis will use micro-denom (uunur) equivalents.

4. Chain Configuration
4.1. Chain ID

Testnet Chain ID must follow convention:

noorchain-testnet-1

4.2. Block Time

Recommended:

5–6 seconds

4.3. Consensus Engine
CometBFT (Tendermint successor)

4.4. Address Prefix
Bech32: noor


noor1… → accounts

noorvaloper1… → validators

noorvalcons1… → consensus

noorpub1… → pubkeys

5. Staking & Validators

Testnet 1.0 will start with one validator: the founder.

5.1. Minimum staking denom
1 unur

5.2. Minimum self-delegation
1 unur

5.3. Commission parameters

Rate: 10%

Max Rate: 20%

Max Change Rate: 1%

5.4. Unbonding time
21 days (default Cosmos value)

5.5. Maximum validators
100

6. Fees & Gas
6.1. Minimum gas price

Testnet recommended:

0.0025unur

6.2. EIP-1559 Base Fee (feemarket)

Enabled.

dynamic base fee

priority fee to validator

gas cap inherited from default Ethermint settings

6.3. Global gas configuration

maxGas = 20,000,000

node operators can override via config

7. EVM Configuration (Ethermint)
7.1. JSON-RPC enabled

Full Ethereum RPC:

eth_*
web3_*
net_*
txpool_*
debug_* (optional)

7.2. EVM Chain Configuration

Matches Ethereum defaults:

London hard fork

EIP-1559

Berlin features

Shanghai / Cancun upgrades can be toggled later

7.3. Pruning

Testnet pruning recommended:

pruning = "nothing"


To allow explorers to index everything.

8. PoSS Genesis Configuration

PoSS must start with initial parameters fixed here.

8.1. Global PoSS Parameters
Parameter	Value (default)	Notes
EpochDuration	24h	daily epochs
DailyMaxSignals	10	per participant
CuratorMaxValidations	50	per curator
WeightMicroDonation	1	
WeightParticipation	2	
WeightCertifiedContent	3	
WeightCCN	5	high-value
ParticipantRatio	0.70	fixed
CuratorRatio	0.30	fixed
HalvingYears	8	8-year halving
HalvingBlocks	computed	based on block time
ReservePoolAddress	one of the 5 addresses	80% pool
StimulusPoolAddress	one of the 5 addresses	5% pool
8.2. Stimulus Rules

The PoSS Stimulus pool distributes:

optional boosted rewards

onboarding incentives

curator activation incentives

Stimulus logic is never automatic; it is admin-triggered.

9. Genesis File Structure

The genesis must include:

9.1. app_state

auth accounts

bank balances

staking state

distribution state

params

evm module state

feemarket state

noorsignal (PoSS) module state

9.2. consensus_params

block time

max gas

9.3. chain_id
noorchain-testnet-1

9.4. genesis_time

Any valid RFC3339 timestamp.

10. Files to Generate During Phase 4C

This document defines the content.
Phase 4C will generate:

10.1. testnet/genesis.json

Full state.

10.2. testnet/genesis_distribution.json

Only balances + allocations.

10.3. Update in:
x/noorsignal/types/addresses.go

10.4. CLI Commands

When ready:

noord init
noord keys add
noord add-genesis-account
noord gentx
noord collect-gentxs
noord start


But NO CODE and NO COMMANDS will be run now — only definition.

11. Summary (Header Block)

NOORCHAIN — Phase3_05 — Genesis Testnet Specification (1.1)
Defines all genesis components required for launching Testnet 1.0:

economic allocation (5/5/5/5/80)

supply distribution

5 official bech32 addresses

PoSS parameters

staking configuration

gas & EIP-1559 settings

EVM configuration

chain configuration

Used later in Phase 4C to generate genesis.json and start the first test network.