package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethermint "github.com/evmos/ethermint/types"

	noorcmd "github.com/noorfinances-eng/noorchain-core/cmd/noord/cmd"
	"github.com/noorfinances-eng/noorchain-core/app"
)

func main() {
	setupConfig()
	registerDenoms()

	rootCmd := noorcmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, noorcmd.EnvPrefix, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}

// setupConfig sets address prefixes + BIP44 coin type (Ethermint-compatible).
func setupConfig() {
	cfg := sdk.GetConfig()

	// Bech32 prefixes (NOORCHAIN)
	cfg.SetBech32PrefixForAccount("noor", "noorpub")
	cfg.SetBech32PrefixForValidator("noorvaloper", "noorvaloperpub")
	cfg.SetBech32PrefixForConsensusNode("noorvalcons", "noorvalconspub")

	// Ethermint coin type (60)
	cfg.SetCoinType(ethermint.Bip44CoinType)

	cfg.Seal()
}

// registerDenoms registers display/base denoms for the CLI.
func registerDenoms() {
	// Display denom: nur
	if err := sdk.RegisterDenom("nur", sdk.OneDec()); err != nil {
		panic(err)
	}
	// Base denom: unur (6 decimals)
	if err := sdk.RegisterDenom(app.BondDenom, sdk.NewDecWithPrec(1, 6)); err != nil {
		panic(err)
	}
}
