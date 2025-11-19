package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
)

// App est la structure principale de l'application NOORCHAIN.
//
// Elle encapsule :
// - BaseApp : le cœur Cosmos SDK (ABCI, stores, etc.)
// - Name / Version : méta-informations simples
// - Keepers : les gestionnaires de modules (auth, bank, PoSS, etc.)
// - Encoding : la configuration d'encodage (codec, TxConfig, etc.)
// - Modules : le module.Manager et structures associées.
type App struct {
	*baseapp.BaseApp

	Name    string
	Version string

	Keepers  AppKeepers
	Encoding EncodingConfig
	Modules  AppModules
}
