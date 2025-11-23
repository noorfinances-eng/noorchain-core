module github.com/noorfinances-eng/noorchain-core

go 1.24

require (
    // Noyau Cosmos
    github.com/cosmos/cosmos-sdk v0.50.5
    github.com/cosmos/cosmos-db v1.0.2

    // CometBFT – version forcée compatible avec Cosmos SDK 0.50.x
    github.com/cometbft/cometbft v1.0.0

    // Dépendances utilitaires
    github.com/gogo/protobuf v1.3.2
    github.com/gorilla/mux v1.8.0
    github.com/spf13/cast v1.6.0
    github.com/spf13/cobra v1.8.1
    github.com/stretchr/testify v1.10.0
    google.golang.org/grpc v1.68.1
    google.golang.org/protobuf v1.35.2
)

replace github.com/cometbft/cometbft => github.com/cometbft/cometbft v1.0.0
