package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -----------------------------------------------------------------------------
// Adresses économiques PoSS (version TEST)
// -----------------------------------------------------------------------------
//
// Ces adresses sont des placeholders pour le développement.
// Avant mainnet, elles seront remplacées par :
// - Fondation NOOR (association)
// - Wallet fondateur (5 %)
// - Wallet Stimulus
// - Wallet Pre-sale
// - Réserve PoSS (multi-sig)
// -----------------------------------------------------------------------------

var (
	// 80 % — Réserve PoSS
	TestPoSSReserveAddr      = mustAcc("noor1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqjx0pcf")

	// 5 % — Fondation NOOR (association light)
	TestFoundationAddr       = mustAcc("noor1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqf5x4s4")

	// 5 % — Dev Wallet (ton wallet perso en mainnet)
	TestDevWalletAddr        = mustAcc("noor1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7hj7q")

	// 5 % — PoSS Stimulus
	TestPoSSStimulusAddr     = mustAcc("noor1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq8lf9za")

	// 5 % — Optional Pre-sale
	TestPreSaleAddr          = mustAcc("noor1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2vmvke")
)

// Helper simple pour convertir une adresse Bech32 en sdk.AccAddress.
// Si l’adresse est invalide → panic (genesis doit être propre).
func mustAcc(bech string) sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(bech)
	if err != nil {
		panic(err)
	}
	return addr
}
