module github.com/noorfinances-eng/noorchain-core

go 1.22

require (
    // Cosmos SDK core (v0.50.x)
    github.com/cosmos/cosmos-sdk v0.50.5

    // DB Cosmos (utilisé dans app_builder.go : dbm "github.com/cosmos/cosmos-db")
    github.com/cosmos/cosmos-db v1.0.2

    // CLI
    github.com/spf13/cobra v1.8.1

    // Protobuf gogo (utilisé par Cosmos)
    github.com/gogo/protobuf v1.3.2

    // Router HTTP simple (si tu l’utilises plus tard)
    github.com/gorilla/mux v1.8.0

    // Tests
    github.com/stretchr/testify v1.10.0
)

// ⚠️ IMPORTANT : PAS DE "replace" sur github.com/cometbft/cometbft ici.
// go mod tidy téléchargera la version attendue par le Cosmos SDK.
// Si tu avais des "replace" avant pour CometBFT ou Ethermint, on les enlève.
