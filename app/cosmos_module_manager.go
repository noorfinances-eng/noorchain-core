package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	authmodule "github.com/cosmos/cosmos-sdk/x/auth/module"
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank/module"
	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking/module"
	govmodule "github.com/cosmos/cosmos-sdk/x/gov/module"

	evmmodule "github.com/evmos/ethermint/x/evm/module"
	feemarketmodule "github.com/evmos/ethermint/x/feemarket/module"
)

// CosmosModuleManager is a wrapper around the Cosmos SDK module manager.
// In Phase 2, it prepares the structure for ordering InitGenesis and block hooks.
type CosmosModuleManager struct {
	Manager *module.Manager
}

// NewCosmosModuleManager returns a CosmosModuleManager with real modules registered (Phase 2 placeholders).
func NewCosmosModuleManager() CosmosModuleManager {

	// Phase 2: no keepers yet → they are all nil.
	mgr := module.NewManager(
		authmodule.NewAppModule(nil, nil),
		bankmodule.NewAppModule(nil, nil, nil),
		stakingmodule.NewAppModule(nil, nil, nil),
		govmodule.NewAppModule(nil, nil, nil),

		evmmodule.NewAppModule(nil, nil),
		feemarketmodule.NewAppModule(nil),
	)

	return CosmosModuleManager{
		Manager: mgr,
	}
}

// SetOrderInitGenesis configures the InitGenesis order.
func (mm *CosmosModuleManager) SetOrderInitGenesis() {
	if mm.Manager == nil {
		return
	}
	mm.Manager.SetOrderInitGenesis(
		authmodule.ModuleName,
		bankmodule.ModuleName,
		stakingmodule.ModuleName,
		govmodule.ModuleName,
		evmmodule.ModuleName,
		feemarketmodule.ModuleName,
	)
}

// SetOrderBeginBlockers configures the BeginBlocker order.
func (mm *CosmosModuleManager) SetOrderBeginBlockers() {
	if mm.Manager == nil {
		return
	}
	mm.Manager.SetOrderBeginBlockers(
		stakingmodule.ModuleName,
		feemarketmodule.ModuleName,
		evmmodule.ModuleName,
	)
}

// SetOrderEndBlockers configures the EndBlocker order.
func (mm *CosmosModuleManager) SetOrderEndBlockers() {
	if mm.Manager == nil {
		return
	}
	mm.Manager.SetOrderEndBlockers(
		stakingmodule.ModuleName,
		govmodule.ModuleName,
	)
}
