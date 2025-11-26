module github.com/noorfinances-eng/noorchain-core

go 1.22

require (
    github.com/cosmos/cosmos-sdk v0.50.14
    github.com/evmos/ethermint v0.22.0
)

replace github.com/evmos/ethermint => github.com/b-harvest/ethermint v0.22.0-sdk50-1
