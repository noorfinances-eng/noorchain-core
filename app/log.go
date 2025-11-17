package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/log"
)

// NewNoorchainLogger returns a basic NOORCHAIN logger.
//
// For now, it returns a no-op logger (it does not print anything).
// Later, this can be changed to log to stdout, files, or structured logs
// depending on how NOORCHAIN nodes are deployed.
func NewNoorchainLogger() sdk.Logger {
	return log.NewNopLogger()
}
