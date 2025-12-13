NOORCHAIN Core

NOORCHAIN Core is the reference implementation of the NOORCHAIN blockchain network.

It provides the full core protocol built on Cosmos SDK with Ethermint (EVM compatibility), including governance, staking, fee market logic, and the native PoSS (Proof of Signal Social) module.

NOORCHAIN is designed as a non-custodial, non-financial, utility-focused blockchain, aligned with the Swiss Legal Light CH framework.

Overview

The NOORCHAIN Core includes:

• account and balance management
• staking and validator infrastructure
• on-chain governance
• EVM execution layer (Ethermint)
• EIP-1559 compatible fee market
• PoSS (Proof of Signal Social) module for social signal validation

The network uses a fixed supply model and does not rely on inflationary mining or yield-based incentives.

Testnet Status

NOORCHAIN Core is currently running in a controlled testnet and local development environment.

The testnet is intended for:

• protocol inspection
• module-level testing
• governance and PoSS validation
• institutional and pilot experimentation

At this stage:

• the testnet is not publicly advertised
• no public RPC endpoints are exposed
• no public block explorer is provided
• access is limited to internal testing and selected pilot programs

A public testnet and related endpoints may be exposed at a later stage, once protocol stability and governance parameters are fully validated.

This controlled approach is intentional and aligned with the non-financial, compliance-first positioning of the project.

Build

The node is compiled using Go.

Required Go version: 1.22 or higher
Entry point: cmd/noord

The project builds using standard Go tooling.

Run a Local Node

After building the binary, a local node can be started using the default home directory.

Running the executable launches a local development instance suitable for inspection, testing, and protocol review.

This local node is intended for development and analysis purposes only.

Repository Structure

app
Contains the application wiring, keepers, and module initialization.

cmd/noord
Entry point for the NOORCHAIN node.

config
Configuration helpers and utilities.

scripts
Operational scripts used for testnet support and maintenance tasks.

x/auth
Module responsible for authentication and account identity.

x/bank
Module handling balances and token transfers.

x/staking
Module governing staking operations and validator functionality.

x/gov
Module managing governance and chain parameters.

x/evm
Ethermint EVM compatibility layer.

x/feemarket
Fee market logic based on EIP-1559.

x/noorsignal
PoSS (Proof of Signal Social) module.

docs
Technical documentation including protocol architecture, genesis parameters, governance rules, economic model, and PoSS specifications.

docs/fr
French versions of the documentation.

Proof of Signal Social (PoSS)

PoSS is a social validation mechanism, not a financial consensus model.

Its key characteristics are:

• no staking
• no APR, APY, or yield
• no promised returns
• capped issuance via a dedicated reserve
• structural reward split: 70% participant / 30% curator
• halving every 8 years

PoSS is implemented as a native Cosmos SDK module and can be enabled or disabled through on-chain governance.

Documentation

All technical documentation is available in the docs directory.

This includes protocol architecture, PoSS logic and parameters, genesis governance, economic model, and compliance framework.

Public-facing PDFs and summaries are published on the official NOORCHAIN website.

Legal & Compliance Notice

NOORCHAIN does not custody user assets, does not offer financial products, does not operate fiat services, and does not guarantee returns.

This repository is provided for transparency, technical review, and educational purposes.

License

License information will be defined prior to mainnet launch.
