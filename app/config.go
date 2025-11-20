package app

// Configuration globale de l'application NOORCHAIN.
//
// Ce fichier définit :
// - le nom de la chaîne
// - le chain-id
// - le prefix Bech32
// - la monnaie native (denom) utilisée pour les frais, le staking et PoSS.

const (
	// Nom interne de l'application.
	AppName = "NOORCHAIN"

	// ChainID utilisé pour le testnet actuel.
	// Il sera ajusté plus tard pour mainnet (ex: "noorchain-1").
	ChainID = "noorchain-testnet-1"

	// Préfixe Bech32 pour les adresses NOORCHAIN.
	// Exemple d'adresse : noor1xxxx...
	Bech32MainPrefix = "noor"
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
)
