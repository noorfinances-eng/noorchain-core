NOORCHAIN — Governance Parameters (Phase 4)

Document version: 1.1
Module: x/gov (Cosmos SDK v0.46.11)
Scope: Governance configuration for NOORCHAIN Testnet 1.0 and Mainnet 1.0
Phase: Phase 4 — Implementation (GOV6)

1. Introduction

This document defines the governance parameters used by the NOORCHAIN blockchain.

Governance in NOORCHAIN is designed to be:

Transparent

Community-driven (Curators + Validators)

Legally compliant (Swiss Legal Light CH)

Efficient and predictable

Compatible with PoSS and the long-term economic model

The parameters below apply first to Testnet 1.0, and will be refined for Mainnet 1.0 during the Genesis Pack phase.

2. Governance Parameters Overview

Governance in the Cosmos SDK uses three sets of parameters:

VotingParams

DepositParams

TallyParams

These parameters define:

How long votes last

How much deposit is required to submit a proposal

How proposals are accepted or rejected

All parameters below are 100% compatible with Cosmos SDK v0.46.11.

3. Voting Parameters

Governance proposals must remain open for a fixed period.
NOORCHAIN uses a shorter period on Testnet to accelerate development, and a longer period on Mainnet for stability.

Testnet 1.0

Voting period: 48 hours

Value (Cosmos format): "172800s"

Mainnet 1.0 (tentative)

Voting period: 7 days

Value: "604800s"

Genesis JSON example
"voting_params": {
  "voting_period": "172800s"
}

4. Deposit Parameters

A minimum deposit prevents spam proposals and ensures that proposers have minimal commitment.

Design principles

Testnet deposits should be symbolic

Mainnet deposits should be realistic and scalable

NOORCHAIN uses unur (micro-denom of NUR)

Testnet 1.0

Minimum deposit: 1000 unur

Max deposit period: 48 hours

Mainnet 1.0 (tentative)

Minimum deposit: to be set depending on early market (likely between 1,000 to 10,000 NUR)

Max deposit period: 7 days

Genesis JSON example
"deposit_params": {
  "min_deposit": [
    { "denom": "unur", "amount": "1000" }
  ],
  "max_deposit_period": "172800s"
}

5. Tally Parameters

These parameters determine how proposals pass or fail.

NOORCHAIN adopts a model based on Cosmos Hub defaults, adjusted for a balanced and safe decision-making process.

Testnet 1.0 and Mainnet 1.0 (tentative)
Quorum

Percentage of staked tokens that must vote for the proposal to be valid.

20%

Prevents decision-making by very small active groups.

Threshold

Percentage of “Yes” votes needed for a proposal to pass.

50% + 1 vote

Simple, democratic majority.

Veto Threshold

Percentage of “NoWithVeto” votes needed to reject the proposal regardless of Yes votes.

33.4%

Protects the chain against dangerous changes.

Genesis JSON example
"tally_params": {
  "quorum": "0.200000000000000000",
  "threshold": "0.500000000000000000",
  "veto_threshold": "0.334000000000000000"
}

6. Notes for Implementation

These parameters will be:

Documented here (Phase 4)

Injected into:

testnet/genesis.json (Phase 4C — Testnet 1.0)

mainnet/genesis.json (Phase 7 — Mainnet 1.0)

Adjustable later through governance proposals once the chain is live

No changes are required in app.go for GOV6.
All modifications occur in the genesis file.

7. Compatibility

These parameters are fully compatible with:

Cosmos SDK v0.46.11

Ethermint v0.22.0

Tendermint/CometBFT v0.34.27

NOORCHAIN PoSS and economic architecture

Swiss Legal Light CH (no custodial governance, fully on-chain)

8. Conclusion

GOV6 finalizes the governance layer design for NOORCHAIN’s Phase 4.
You now have:

Governance stores

GovKeeper

Gov module included in ModuleManager

Genesis order configured

Full governance parameter specification

Future-proof structure for Testnet and Mainnet