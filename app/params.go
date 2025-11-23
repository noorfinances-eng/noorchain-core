package app

// Basic, chain-wide NOORCHAIN parameters.
// These values complement config.go and will be used later when wiring
// Cosmos SDK, Ethermint and when generating genesis files.
//
// IMPORTANT:
// - AppName, ChainID and CoinDenom / CoinDisplayDenom / CoinDecimals
//   are defined in config.go and must NOT be redefined here.

const (
	// Bech32 address prefixes (will be used by Cosmos SDK's bech32 config).
	//
	// Final forms:
	// - noor1...        (accounts)
	// - noorvaloper1... (validators)
	// - noorvalcons1... (validator consensus keys)
	Bech32PrefixAccAddr  = Bech32MainPrefix
	Bech32PrefixAccPub   = Bech32MainPrefix + "pub"
	Bech32PrefixValAddr  = Bech32MainPrefix + "valoper"
	Bech32PrefixValPub   = Bech32MainPrefix + "valoperpub"
	Bech32PrefixConsAddr = Bech32MainPrefix + "valcons"
	Bech32PrefixConsPub  = Bech32MainPrefix + "valconspub"

	// Token parameters
	//
	// NUR is the human-readable token symbol.
	// "unur" is the base denom (like "uatom" for ATOM, "uosmo" for OSMO).
	//
	// 1 NUR = 10^CoinDecimals unur  (micro-nur)
	TokenSymbol  = CoinDisplayDenom
	BaseDenom    = CoinDenom
	CoinDecimals = 6 // must stay in sync with config.go

	// EVM chain ID for Ethermint integration.
	// This is a placeholder and may be adjusted later
	// to avoid collisions with existing EVM chains.
	// (symbolic "speed of light" value)
	EvmChainID uint64 = 299792458
)
