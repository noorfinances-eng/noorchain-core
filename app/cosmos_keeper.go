package app

// CosmosKeepers is a placeholder struct grouping all Cosmos SDK keepers.
// In Phase 2 this remains empty; real keepers will be added in later steps.
type CosmosKeepers struct {
	// AuthKeeper      interface{}
	// BankKeeper      interface{}
	// StakingKeeper   interface{}
	// GovKeeper       interface{}
	// EvmKeeper       interface{}
	// FeeMarketKeeper interface{}
}

// NewCosmosKeepers returns an empty placeholder keeper set.
func NewCosmosKeepers() CosmosKeepers {
	return CosmosKeepers{}
}
