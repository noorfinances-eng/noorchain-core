package app

import (
	"github.com/noorfinances-eng/noorchain-core/x/auth"
	"github.com/noorfinances-eng/noorchain-core/x/bank"
	"github.com/noorfinances-eng/noorchain-core/x/staking"
	"github.com/noorfinances-eng/noorchain-core/x/gov"
	"github.com/noorfinances-eng/noorchain-core/x/evm"
	"github.com/noorfinances-eng/noorchain-core/x/feemarket"
)

// AppModules bundles all NOORCHAIN modules in their placeholder form.
// In Phase 2, they are not wired to Cosmos SDK yet.
type AppModules struct {
	Auth      auth.AppModule
	Bank      bank.AppModule
	Staking   staking.AppModule
	Gov       gov.AppModule
	EVM       evm.AppModule
	FeeMarket feemarket.AppModule
}

// NewAppModules returns the placeholder module collection.
func NewAppModules() AppModules {
	return AppModules{
		Auth:      auth.AppModule{},
		Bank:      bank.AppModule{},
		Staking:   staking.AppModule{},
		Gov:       gov.AppModule{},
		EVM:       evm.AppModule{},
		FeeMarket: feemarket.AppModule{},
	}
}
