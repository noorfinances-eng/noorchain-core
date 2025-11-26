package app

// CosmosKeepers groups all Cosmos SDK and Ethermint keepers.
// Phase 2: all fields are placeholders (nil), real keepers come in Phase 3.
type CosmosKeepers struct {
	AuthKeeper        interface{}
	BankKeeper        interface{}
	StakingKeeper     interface{}
	GovKeeper         interface{}
	EvmKeeper         interface{}
	FeeMarketKeeper   interface{}
}

// NewCosmosKeepers returns a keeper set with placeholder (nil) keepers.
// This allows the ModuleManager to be instantiated in Phase 2.
func NewCosmosKeepers() CosmosKeepers {
	return CosmosKeepers{
		AuthKeeper:        nil,
		BankKeeper:        nil,
		StakingKeeper:     nil,
		GovKeeper:         nil,
		EvmKeeper:         nil,
		FeeMarketKeeper:   nil,
	}
}
