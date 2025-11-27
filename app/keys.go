package app

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// Bech32 prefix root
const (
    Bech32Prefix = "noor"
)

// ConfigureBech32Prefixes sets the Bech32 prefixes for NOORCHAIN.
func ConfigureBech32Prefixes() {
    config := sdk.GetConfig()

    config.SetBech32PrefixForAccount(Bech32Prefix, Bech32Prefix+"pub")
    config.SetBech32PrefixForValidator(Bech32Prefix+"valoper", Bech32Prefix+"valoperpub")
    config.SetBech32PrefixForConsensusNode(Bech32Prefix+"valcons", Bech32Prefix+"valconspub")

    config.Seal()
}
