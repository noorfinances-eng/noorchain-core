package app

// ------------------------------------------------------------
//  NOORCHAIN — Configuration centrale de la blockchain
// ------------------------------------------------------------
//
//  Toutes les valeurs officielles sont définies ici.
//  Elles sont utilisées ensuite par :
//   - app.go
//   - encoding.go
//   - module manager
//   - genesis
//   - CLI (plus tard)
// ------------------------------------------------------------

// Nom officiel de l'application (visible dans BaseApp)
const AppName = "noorchain"

// Identifiant de chaîne par défaut (modifiable dans un testnet)
const DefaultChainID = "noorchain-1"

// Bech32 prefix
//
// Résultat :
//   - noor1...        (comptes)
//   - noorvaloper1... (validateurs)
//   - noorvalcons1... (clés consensus)
const (
	Bech32MainPrefix = "noor"
)

// Token et denomination
//
// Décision officielle NOORCHAIN :
// - Token public : NUR
// - Dénomination interne base-unit : unur
// - Décimales : 18 (aligné Ethereum / Cosmos moderne)
const (
	CoinDenom         = "unur" // base denom
	CoinDisplayDenom  = "NUR"  // affichage humain
	CoinDecimals uint8 = 18
)
