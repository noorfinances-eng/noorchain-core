package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ConfigureSDK applies the global Cosmos SDK configuration for NOORCHAIN.
//
// This mainly sets the Bech32 prefixes for:
// - accounts (noor1...)
// - validators (noorvaloper1...)
// - consensus nodes (noorvalcons1...)
//
// IMPORTANT:
// - This function must be called once at application startup
//   (for example in cmd/noord/main.go) *before* using Bech32 addresses.
// - For now, we only define it and we will plug it into the startup flow
//   in a later step.
func ConfigureSDK() {
	cfg := sdk.GetConfig()

	cfg.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)

	// Optionally, we could also configure coin type and full HD path here.
	// For now we keep it simple.
	//
	// Example (if needed later):
	// cfg.SetCoinType(118) // Cosmos default
	// cfg.SetFullFundraiserPath("44'/118'/0'/0/0")

	cfg.Seal()
}
