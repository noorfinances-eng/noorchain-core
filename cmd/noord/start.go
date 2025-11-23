package main

import (
	"github.com/noorfinances-eng/noorchain-core/app"
)

// Start initializes and runs the NOORCHAIN application server.
// In Phase 2 this is only a thin wrapper around placeholders.
func Start() error {
	a := app.InitApp()
	return app.BuildServer(a)
}
