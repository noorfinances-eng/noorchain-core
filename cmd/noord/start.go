package main

import (
	"github.com/noorfinances-eng/noorchain-core/app"
)

// Start initializes and runs the NOORCHAIN application server.
// Now uses the Cosmos-based NOORChainApp returned by InitApp().
func Start() error {
	noorApp := app.InitApp()
	return app.BuildServer(noorApp)
}
