package app

import (
	"cosmossdk.io/log"
)

// GetLogger retourne le logger principal de l'application.
// V1 : simple wrapper autour du logger de BaseApp.
func (a *App) GetLogger() log.Logger {
	if a == nil || a.BaseApp == nil {
		return log.NewNopLogger()
	}
	return a.BaseApp.Logger()
}
