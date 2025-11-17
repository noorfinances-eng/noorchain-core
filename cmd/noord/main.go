package main

import (
    "log"

    "github.com/noorfinances-eng/noorchain-core/app"
)

func main() {
    // For now we only create and start a very simple placeholder application.
    // In the next steps, this will be replaced by a full Cosmos SDK + Ethermint node.
    noa := app.NewNoorchainApp()

    if err := noa.Start(); err != nil {
        log.Fatalf("failed to start NOORCHAIN node: %v", err)
    }
}
