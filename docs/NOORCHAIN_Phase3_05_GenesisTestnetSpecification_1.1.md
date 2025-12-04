NOORCHAIN 1.0 — Phase 3.05
Genesis Testnet 1.0 Specification
Version 1.1 — Official Document
Purpose of this document

This document defines all mandatory genesis parameters required to launch
NOORCHAIN Testnet 1.0, including:

the 5 official genesis addresses

initial token distribution

PoSS reserve configuration

EVM / Ethermint settings

fees, gas, staking

network parameters

final file structure for genesis.json

This is the canonical reference for Phase 4C (Genesis Generation).

1. Genesis Philosophy

NOORCHAIN Testnet 1.0 is designed to:

strictly follow the 5% / 5% / 5% / 5% / 80% economic model

mirror the final mainnet structure

allow testing of Cosmos, EVM, and PoSS components

run on CometBFT

expose full Ethereum JSON-RPC

remain entirely deterministic and auditable

Everything defined here is mandatory for Testnet 1.0.

2. Genesis Addresses (Official 5-Key Structure)

Testnet requires five bech32 addresses representing the immutable economic allocation model:

Allocation	Description	Genesis Role
5% – Foundation Wallet	Public foundation operations	receives 5%
5% – Dev Sàrl Wallet	Founder operational & technical	receives 5%
5% – PoSS Stimulus Pool	Early adoption & NGO incentives	receives 5%
5% – Pre-Sale Pool	Optional Swiss-regulated private investors	receives 5%
80% – PoSS Mintable Reserve	Long-term PoSS emission	receives 80%
2.1 Requirements for these addresses

must use bech32 prefix: noor1

must be freshly generated

must be recorded in:

x/noorsignal/types/addresses.go

testnet/genesis.json

testnet/genesis_distribution.json

must correspond exactly to the 5 pools

will be inserted during Phase 4C

2.2 When they are generated

→ Phase 4C: "Generate 5 real bech32 addresses (noor1…)"

This document defines their roles, not their values.

3. Total Supply & Allocation

NOORCHAIN supply is fixed and immutable:

299,792,458 NUR
3.1 Token denom

base: unur

display: NUR

decimals: 18

3.2 Genesis distribution
Pool	%	Amount NUR	Purpose
Foundation	5%	14,989,622.90	governance + public interest
Dev Sàrl	5%	14,989,622.90	development, infra
PoSS Stimulus	5%	14,989,622.90	early adoption
Pre-Sale	5%	14,989,622.90	regulated investors
PoSS Reserve	80%	239,833,966.40	PoSS long-term rewards

Genesis will use micro-denom (uunur) values.

4. Chain Configuration
4.1 Chain ID

noorchain-testnet-1

4.2 Block Time

Recommended: 5–6 seconds

4.3 Consensus Engine

CometBFT (Tendermint successor)

4.4 Address Prefixes

accounts → noor1…

validators → noorvaloper1…

consensus → noorvalcons1…

pubkeys → noorpub1…

5. Staking & Validators

Testnet 1.0 launches with one validator: the founder.

5.1 Minimum staking denom

1 unur

5.2 Minimum self-delegation

1 unur

5.3 Commission parameters

Rate: 10%

Max Rate: 20%

Max Change Rate: 1%

5.4 Unbonding time

21 days

5.5 Maximum validators

100

6. Fees & Gas
6.1 Minimum gas price

0.0025unur

6.2 EIP-1559 Base Fee (feemarket)

dynamic base fee

priority fee for validators

inherits Ethermint defaults

6.3 Global gas configuration

maxGas = 20,000,000

Node operators may override in config.

7. EVM Configuration (Ethermint)
7.1 JSON-RPC Enabled

Full Ethereum compatibility:

eth_*

web3_*

net_*

txpool_*

debug_* (optional)

7.2 Hard Fork Configuration

Enabled by default:

Berlin

London (EIP-1559)

Future upgrades (Shanghai / Cancun) can be toggled later.

7.3 Pruning Mode

Recommended for Testnet:

pruning = "nothing"

Allows explorers to index all blocks.

8. PoSS Genesis Configuration

PoSS starts with the following base parameters.

8.1 Global PoSS Parameters
Parameter	Value	Notes
EpochDuration	24h	daily reset
DailyMaxSignals	10	per participant
CuratorMaxValidations	50	per curator
WeightMicroDonation	1	
WeightParticipation	2	
WeightCertifiedContent	3	
WeightCCN	5	
ParticipantRatio	0.70	immutable
CuratorRatio	0.30	immutable
HalvingYears	8	every 8 years
HalvingBlocks	computed	based on block time
ReservePoolAddress	PoSS 80% pool	
StimulusPoolAddress	PoSS Stimulus 5%	
8.2 Stimulus Rules

The Stimulus Pool may be used to:

boost first-year adoption

reward NGOs / schools onboarding

activate curator networks

It is admin-triggered, never automatic.

9. Genesis File Structure

Genesis must include:

9.1 app_state

auth

bank

staking

distribution

params

evm

feemarket

noorsignal (PoSS)

9.2 consensus_params

block time

max gas

9.3 chain_id

noorchain-testnet-1

9.4 genesis_time

Valid RFC3339 timestamp.

10. Files Generated During Phase 4C

This specification will be used to generate:

10.1 testnet/genesis.json

Complete application state.

10.2 testnet/genesis_distribution.json

Balances + initial allocations.

10.3 Address Injection

Update in:

x/noorsignal/types/addresses.go

10.4 CLI Commands (executed later)

noord init

noord keys add

noord add-genesis-account

noord gentx

noord collect-gentxs

noord start

🚫 No commands executed now — definition only.

11. Summary (Header Block)

NOORCHAIN — Phase3_05 — Genesis Testnet Specification (1.1)
Defines all components needed to launch Testnet 1.0:

economic distribution (5 / 5 / 5 / 5 / 80)

supply allocation

five official bech32 genesis addresses

PoSS genesis parameters

staking configuration

EVM & EIP-1559 settings

chain configuration

file structure for genesis.json

This is the official foundation for Phase 4C: Testnet Genesis Generation.
