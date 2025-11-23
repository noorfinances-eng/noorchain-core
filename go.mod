module github.com/noorfinances-eng/noorchain-core

go 1.24

require (
    github.com/cometbft/cometbft v1.0.0
    github.com/cosmos/cosmos-sdk v0.50.5
    // ⚠️ Ethermint sera ajouté plus tard quand on choisira une vraie version compatible.
    // github.com/evmos/ethermint v1.0.0
)

require (
    github.com/cosmos/cosmos-db v1.0.0 // indirect
    github.com/gogo/protobuf v1.3.2    // indirect
    github.com/gorilla/mux v1.8.1      // indirect
    github.com/stretchr/testify v1.8.4 // indirect
    github.com/spf13/cobra v1.7.0      // indirect
    github.com/spf13/cast v1.6.0       // indirect
    github.com/tendermint/tendermint v0.34.27         // indirect
    github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
    google.golang.org/grpc v1.64.0                    // indirect
    google.golang.org/protobuf v1.33.0                // indirect
)

replace github.com/gogo/protobuf v1.3.3 => github.com/gogo/protobuf v1.3.2
