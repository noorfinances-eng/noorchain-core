NOORCHAIN Core

NOORCHAIN Core is the reference implementation of the NOORCHAIN network.

It provides the main components of the blockchain, based on Cosmos SDK and Ethermint.
The system includes account management, balances, staking, governance, EVM execution, fee market management, and the PoSS (Proof of Signal Social) module.

Build

The node is compiled using Go.
The build process targets the entry point available in the cmd/noord directory.
Go version 1.22 or higher is required.

Run a Local Node

After building the binary, a node can be started using the default home directory.
Running the executable is sufficient to launch a local instance.

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
Technical documentation including protocol, genesis, governance and PoSS specifications.

Documentation

All technical documentation is available in the docs directory.
French versions of each document are located in docs/fr.

License

To be defined.
