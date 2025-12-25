package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	authmodule "github.com/cosmos/cosmos-sdk/x/auth"
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank"
	genutilmodule "github.com/cosmos/cosmos-sdk/x/genutil"
	paramsmodule "github.com/cosmos/cosmos-sdk/x/params"
	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking"
)

// ModuleBasics fournit les AppModuleBasic pour l’enregistrement (gRPC Gateway, codec, etc.)
var ModuleBasics = module.NewBasicManager(
	authmodule.AppModuleBasic{},
	bankmodule.AppModuleBasic{},
	stakingmodule.AppModuleBasic{},
	genutilmodule.AppModuleBasic{},
	paramsmodule.AppModuleBasic{}, // ajout du module params
)