package main

import (
	"log"

	"github.com/noorfinances-eng/noorchain-core/app"
)

func main() {
	// 1) Configure global Cosmos SDK settings for NOORCHAIN
	//    (Bech32 prefixes: noor1..., noorvaloper1..., etc.)
	app.ConfigureSDK()

	// 2) For now we still create the simple placeholder application.
	//    In the next technical phases, this will be replaced by the full
	//    Cosmos SDK + Ethermint powered NOORCHAIN app.
	noa := app.NewNoorchainApp()

	if err := noa.Start(); err != nil {
		log.Fatalf("failed to start NOORCHAIN node: %v", err)
	}
}
