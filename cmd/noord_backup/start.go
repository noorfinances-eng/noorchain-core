package main

import (
	"github.com/noorfinances-eng/noorchain-core/app"
)

// Start triggers the Cosmos-based NOORCHAIN application startup.
func Start() error {
	return app.StartNOORChain()
}
