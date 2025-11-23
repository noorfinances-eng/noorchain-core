package app

// Configuration globale de l'application NOORCHAIN.
//
// Ce fichier définit :
// - le nom de la chaîne
// - le chain-id par défaut (testnet)
// - les préfixes Bech32
// - la monnaie native (denom) utilisée pour les frais, le staking et PoSS
// - quelques paramètres EVM de base (chain ID).

// -----------------------------------------------------------------------------
//  Identité de la chaîne
// -----------------------------------------------------------------------------

const (
	// Nom interne de l'application.
	AppName = "NOORCHAIN"

	// ChainID utilisé pour le testnet actuel.
	// Il sera ajusté plus tard pour le mainnet (ex: "noorchain-mainnet-1").
	ChainID = "noorchain-testnet-1"
)

// -----------------------------------------------------------------------------
//  Préfixes Bech32
// -----------------------------------------------------------------------------

// Préfixe principal pour les adresses NOORCHAIN.
// Exemple d'adresse compte : noor1xxxx...
const Bech32MainPrefix = "noor"

const (
	// Comptes utilisateurs (acc)
	Bech32PrefixAccAddr = Bech32MainPrefix
	Bech32PrefixAccPub  = Bech32MainPrefix + "pub"

	// Adresses de validateurs (valoper)
	Bech32PrefixValAddr = Bech32MainPrefix + "valoper"
	Bech32PrefixValPub  = Bech32MainPrefix + "valoperpub"

	// Adresses de consensus (valcons)
	Bech32PrefixConsAddr = Bech32MainPrefix + "valcons"
	Bech32PrefixConsPub  = Bech32MainPrefix + "valconspub"
)

// -----------------------------------------------------------------------------
//  Configuration de la monnaie native NUR
// -----------------------------------------------------------------------------

const (
	// CoinDenom est le denom interne de la monnaie native.
	// Convention Cosmos : utiliser un "u" pour 10^-6 (micro).
	// Ici : 1 NUR = 1_000_000 unur.
	CoinDenom = "unur"

	// CoinDisplayDenom est la façon "humaine" d'afficher la monnaie.
	// Exemple : 12.345678 NUR
	CoinDisplayDenom = "NUR"

	// CoinDecimals définit le nombre de décimales pour la représentation
	// humaine du token. Avec 6 décimales :
	//   1 NUR = 1_000_000 unur
	CoinDecimals = 6

	// BondDenom est le denom utilisé pour le staking, les frais, etc.
	// On le met égal à CoinDenom pour garder une seule monnaie native.
	BondDenom = CoinDenom

	// TokenSymbol est le symbole lisible (alias de CoinDisplayDenom).
	TokenSymbol = CoinDisplayDenom
)

// -----------------------------------------------------------------------------
//  Paramètres EVM (Ethermint)
// -----------------------------------------------------------------------------

// EvmChainID est l'ID de chaîne EVM utilisé par Ethermint.
// Valeur symbolique "vitesse de la lumière" pour NOORCHAIN.
// (On pourra l'ajuster plus tard pour éviter tout conflit si nécessaire.)
const EvmChainID = 299792458
