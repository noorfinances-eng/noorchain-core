package app

// Configuration globale de l'application NOORCHAIN.
//
// Ce fichier est la SEULE source de vérité pour :
// - le nom de l'app
// - les chain-id (testnet / mainnet)
// - les préfixes Bech32
// - la monnaie native (NUR / unur)
// - l'ID EVM pour Ethermint.

const (
	// Nom interne de l'application.
	AppName = "NOORCHAIN"

	// ChainID utilisés pour NOORCHAIN.
	// - ChainIDTestnet : premier testnet public
	// - ChainIDMainnet : mainnet cible (plus tard)
	ChainIDTestnet = "noorchain-testnet-1"
	ChainIDMainnet = "noorchain-1"

	// ChainID actuellement utilisé par l'app.
	// Pour Testnet 1.0 on pointe sur ChainIDTestnet.
	ChainID = ChainIDTestnet

	// Préfixe Bech32 de base pour NOORCHAIN.
	// Exemple d'adresses :
	// - noor1...        (comptes)
	// - noorvaloper1... (validateurs)
	// - noorvalcons1... (clés consensus)
	Bech32MainPrefix = "noor"

	Bech32PrefixAccAddr  = Bech32MainPrefix
	Bech32PrefixAccPub   = Bech32MainPrefix + "pub"
	Bech32PrefixValAddr  = Bech32MainPrefix + "valoper"
	Bech32PrefixValPub   = Bech32MainPrefix + "valoperpub"
	Bech32PrefixConsAddr = Bech32MainPrefix + "valcons"
	Bech32PrefixConsPub  = Bech32MainPrefix + "valconspub"

	// -------------------------------------------------------------------------
	//  Configuration de la monnaie native NUR
	// -------------------------------------------------------------------------

	// Base denom interne (comme "uatom" pour ATOM, "uosmo" pour OSMO).
	// Convention : "u" = micro (10^-6).
	// Ici : 1 NUR = 1_000_000 unur
	CoinDenom        = "unur"
	CoinDisplayDenom = "NUR"
	CoinDecimals     = 6

	// BondDenom est utilisé pour le staking, les frais, etc.
	// On utilise la même monnaie partout.
	BondDenom = CoinDenom

	// -------------------------------------------------------------------------
	//  EVM / Ethermint
	// -------------------------------------------------------------------------

	// EvmChainID est l'ID de chaîne EVM pour la partie Ethermint.
	// Valeur symbolique "vitesse de la lumière" 299_792_458.
	// On pourra l'ajuster plus tard si collision avec une autre chaîne.
	EvmChainID = 299792458
)
