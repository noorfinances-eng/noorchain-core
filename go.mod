module github.com/noorfinances-eng/noorchain-core

go 1.21

require (
    github.com/cosmos/cosmos-sdk v0.47.11
    github.com/evmos/ethermint v0.22.0
    github.com/gogo/protobuf v1.3.2
)

replace (
    github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
    github.com/cosmos/cosmos-sdk/simapp => github.com/cosmos/cosmos-sdk v0.47.11
    github.com/tendermint/tendermint => github.com/tendermint/tendermint v0.34.21

)
