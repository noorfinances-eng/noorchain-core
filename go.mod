module github.com/noorfinances-eng/noorchain-core

go 1.21

require (
    github.com/cosmos/cosmos-sdk v0.46.11
    github.com/evmos/ethermint v0.22.0
    github.com/gogo/protobuf v1.3.2
)

replace (
    // Ethermint v0.22 est fait pour Cosmos SDK 0.46.11
    github.com/cosmos/cosmos-sdk/simapp => github.com/cosmos/cosmos-sdk v0.46.11

    // Gogo protobuf : fork regen qui contient le package grpc
    github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

    // Tendermint est remplacé par CometBFT v0.34.27
    github.com/tendermint/tendermint => github.com/cometbft/cometbft v0.34.27

    // IMPORTANT : on force aussi cometbft lui-même à v0.34.27
    github.com/cometbft/cometbft => github.com/cometbft/cometbft v0.34.27

    // Ethermint attend btcd v0.23.4 (chemin btcec ancien)
    github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.23.4

    // On fixe une version unique de la lib Rosetta pour enlever l’ambiguïté
    github.com/coinbase/rosetta-sdk-go => github.com/coinbase/rosetta-sdk-go v0.7.9
)
