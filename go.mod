module github.com/noorfinances-eng/noorchain-core

go 1.22

require (
    github.com/cosmos/cosmos-sdk v0.50.14
    github.com/evmos/ethermint v0.22.0
    cosmossdk.io/core v0.11.0
)

replace github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.50.14

// On utilise le fork Ethermint adapté SDK 0.50.x
replace github.com/evmos/ethermint => github.com/b-harvest/ethermint v0.22.0-sdk50-1

// On force la bonne version de core pour éviter le problème BlockInfo/Info
replace cosmossdk.io/core => cosmossdk.io/core v0.11.0
