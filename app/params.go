package app

// Basic, chain-wide NOORCHAIN parameters.
// These values will be used later when wiring Cosmos SDK, Ethermint
// and when generating genesis files.

const (
	// App / chain identity
	AppName  = "noorchain"
	ChainID  = "noorchain-1" // will be used for mainnet; testnets will use suffixes

	// Bech32 address prefixes (will be used by Cosmos SDK's bech32 config).
	//
	// Example final forms (later):
	// - noor1...        (accounts)
	// - noorvaloper1... (validators)
	// - noorvalcons1... (validator consensus keys)
	Bech32PrefixAccAddr = "noor"
	Bech32PrefixAccPub  = "noorpub"
	Bech32PrefixValAddr = "noorvaloper"
	Bech32PrefixValPub  = "noorvaloperpub"
	Bech32PrefixConsAddr = "noorvalcons"
	Bech32PrefixConsPub  = "noorvalconspub"

	// Token parameters
	//
	// NUR is the human-readable token symbol.
	// "unur" is the base denom (like "uatom" for ATOM, "uosmo" for OSMO).
	//
	// 1 NUR  = 1_000_000 unur  (micro-nur)
	TokenSymbol   = "NUR"
	BaseDenom     = "unur"
	CoinDecimals  = 6 // 1 NUR = 10^6 unur

	// EVM chain ID for Ethermint integration.
	// This is a placeholder and may be adjusted later
	// to avoid collisions with existing EVM chains.
	EvmChainID = 299792458 // "speed of light" symbolic value (can be changed later if needed)
)
