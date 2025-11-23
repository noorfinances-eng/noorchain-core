package app

import (
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
)

// NewBaseNoorchainApp crée une instance minimale de l'application NOORCHAIN.
// Pour l'instant, c'est un constructeur basique qui sera enrichi plus tard
// avec les keepers, modules, etc.
func NewBaseNoorchainApp(
	logger log.Logger,
	base *baseapp.BaseApp,
) *App {
	if base == nil {
		base = baseapp.NewBaseApp("noorchain", logger, nil, nil)
	}

	return &App{
		BaseApp: base,
		Name:    "noorchain",
		Version: "0.1.0",
		// Keepers, Encoding et Modules seront câblés plus tard.
	}
}

// Logger renvoie le logger de l'application en utilisant le logger de BaseApp.
func (a *App) Logger() log.Logger {
	if a == nil {
		return log.NewNopLogger()
	}
	return a.BaseApp.Logger()
}
