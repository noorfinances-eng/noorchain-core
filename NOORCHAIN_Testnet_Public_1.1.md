NOORCHAIN — Testnet (Public)

NOORCHAIN is a Swiss social blockchain built on the Proof of Signal Social (PoSS) model.

This repository contains the core implementation of the NOORCHAIN blockchain, developed using Cosmos SDK with Ethermint (EVM compatibility).

The NOORCHAIN testnet exists to support:

protocol inspection and review,

controlled technical validation,

educational and institutional pilots,

early ecosystem experimentation.

NOORCHAIN operates under a Swiss Legal Light CH framework:

non-custodial architecture,

non-financial positioning,

no promised returns,

utility-only token model.

Status

Core blockchain implementation: ✅ completed

PoSS module (x/noorsignal): ✅ implemented (disabled by default)

Testnet: ✅ operational (controlled access)

Mainnet: ⏳ not launched

This testnet is not intended for production use.

Architecture Overview

Consensus base: Cosmos SDK (Tendermint / CometBFT)

EVM compatibility: Ethermint

Social layer: Proof of Signal Social (PoSS)

Token supply: fixed at 299,792,458 NUR

Inflation: none (PoSS uses a capped reserve)

Governance: on-chain governance with multi-sig execution (3/5)

Proof of Signal Social (PoSS)

PoSS is a social validation mechanism, not a financial consensus model.

Its purpose is to record and reward verified social signals within the NOORCHAIN ecosystem.

Key principles:

no staking,

no APR, APY, or yield,

no financial promises,

rewards are symbolic and capped,

structural reward split: 70% participant / 30% curator,

halving cycle every 8 years,

PoSS activation is controlled through governance.

PoSS is implemented as a native Cosmos SDK module (x/noorsignal).

Testnet Information

Network parameters:

Chain ID: noorchain-testnet

Base token denom: unur (display: NUR)

EVM enabled: yes

PoSS enabled by default: no

Access policy:

At this stage, the testnet is:

deployed internally,

accessible for controlled pilots and reviews,

documented publicly for transparency.

Public RPC endpoints and block explorers may be exposed at a later stage, once protocol stability and governance parameters are fully validated.

This gradual exposure is intentional and aligned with the compliance-first approach of the project.

Local Node (Developer & Auditor Preview)

A local NOORCHAIN node can be built and executed for inspection and review purposes.

This local environment is intended for:

developers,

auditors,

protocol reviewers.

It is not designed for production or public usage.

Detailed build and execution instructions are available in the main README of the repository.

Legal & Compliance Notice

NOORCHAIN does not custody user assets, does not offer financial products, does not operate fiat services, and does not guarantee returns.

The testnet and this repository are provided strictly for transparency, technical review, research, and educational purposes.
