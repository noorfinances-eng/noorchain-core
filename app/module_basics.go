package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	// Modules Cosmos SDK
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	bank "github.com/cosmos/cosmos-sdk/x/bank"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	gov "github.com/cosmos/cosmos-sdk/x/gov"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	genutil "github.com/cosmos/cosmos-sdk/x/genutil"
	genutilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

// ModuleBasics définit les modules de base utilisés par NOORCHAIN.
// À ce stade (Phase 2), nous n'activons que les modules standard,
// sans keeper et sans logique spécifique.
var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	bank.AppModuleBasic{},
	staking.AppModuleBasic{},
	gov.AppModuleBasic{},
	genutil.AppModuleBasic{},
)

// RegisterInterfaces permet aux modules d’enregistrer leurs types
// dans le registre global d’interfaces.
func RegisterInterfaces(cfg EncodingConfig) {
	ModuleBasics.RegisterInterfaces(cfg.InterfaceRegistry)

	// Enregistrement des interfaces standards
	authTypes.RegisterInterfaces(cfg.InterfaceRegistry)
	bankTypes.RegisterInterfaces(cfg.InterfaceRegistry)
	stakingTypes.RegisterInterfaces(cfg.InterfaceRegistry)
	govTypes.RegisterInterfaces(cfg.InterfaceRegistry)
	genutilTypes.RegisterInterfaces(cfg.InterfaceRegistry)
}
