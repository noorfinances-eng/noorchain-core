package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"

	noorsignalmodule "github.com/noorfinances-eng/noorchain-core/x/noorsignal"
	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
)

// AppModules regroupe les éléments liés aux modules Cosmos de NOORCHAIN.
//
// On y conserve un module.Manager (pour Begin/EndBlockers, InitGenesis)
// et on ajoute un Configurator (nouvelle API pour enregistrer MsgServer,
// QueryServer, services gRPC, etc.)
type AppModules struct {
	Manager      *module.Manager
	Configurator module.Configurator
}

// NewAppModuleManager construit le moduleManager + le configurator.
//
func NewAppModuleManager(keepers AppKeepers, encCfg EncodingConfig) AppModules {

	// --- 1) Construire le module PoSS ---
	noorSignalModule := noorsignalmodule.NewAppModule(keepers.NoorSignalKeeper)

	// --- 2) Construire le Manager avec uniquement PoSS pour l'instant ---
	mm := module.NewManager(
		noorSignalModule,
	)

	// --- 3) Construire un Configurator ---
	configurator := module.NewConfigurator(
		encCfg.Marshaler, // codec
		mm,               // manager
	)

	// --- 4) Enregistrer les services du module PoSS ---
	// Le module noorsignal a besoin :
	// - d'un MsgServer (transactions)
	// - d'un QueryServer (queries gRPC)
	//
	// On crée ici le MsgServer avec le BankKeeper pré-câblé.
	noorsignalkeeper.RegisterMsgServer(
		configurator.MsgServer(),                     // service server
		noorsignalkeeper.NewMsgServer(                // instance
			keepers.NoorSignalKeeper,
			keepers.BankKeeper,                      // pré-câblé
		),
	)

	// Query Server (fournit gRPC / CLI "noord query noorsignal ...")
	noorsignalkeeper.RegisterQueryServer(
		configurator.QueryServer(),
		noorsignalkeeper.NewQueryServer(keepers.NoorSignalKeeper),
	)

	return AppModules{
		Manager:      mm,
		Configurator: configurator,
	}
}

// ConfigureModuleManagerOrder applique l'ordre BeginBlock / EndBlock / InitGenesis
// défini dans modules_layout.go
func ConfigureModuleManagerOrder(mm *module.Manager) {
	if mm == nil {
		return
	}

	mm.SetOrderBeginBlockers(BeginBlockerOrder...)
	mm.SetOrderEndBlockers(EndBlockerOrder...)
	mm.SetOrderInitGenesis(InitGenesisOrder...)
}
