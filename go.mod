module github.com/noorfinances-eng/noorchain-core

go 1.22

require (
    github.com/cosmos/cosmos-sdk v0.47.11
    github.com/evmos/ethermint v0.22.0
)

replace github.com/tendermint/tendermint => github.com/cometbft/cometbft v0.37.2
replace github.com/cosmos/cosmos-sdk/simapp => cosmossdk.io/simapp v0.0.0-20231103111158-e83a20081ced
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2

