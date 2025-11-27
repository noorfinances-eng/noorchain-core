package app

// Configuration de base de NOORCHAIN.
// Ici on ne fait que définir des constantes simples, sans aucun import.
// Elles seront utilisées plus tard dans la Phase 2 (initialisation Cosmos).

const (
	// Nom de l'application et ID de chaîne par défaut.
	AppName = "noorchain"
	ChainID = "noorchain-1"

	// Monnaie principale de NOORCHAIN.
	// On garde la logique NOOR : NUR affiché, unur en "micro-unités".
	BondDenom     = "unur" // unité interne (10^6 = 1 NUR)
	DisplayDenom  = "NUR"
	CoinDecimals  = 6      // 1 NUR = 10^6 unur

	// Préfixes Bech32 (adresses NOORCHAIN).
	//
	// Résultat :
	// - noor1...        (comptes)
	// - noorvaloper1... (validateurs)
	// - noorvalcons1... (clés consensus)
	// - noorvalconspub1... (clés consensus publiques)
	Bech32MainPrefix   = "noor"
	Bech32PrefixAccAddr  = Bech32MainPrefix
	Bech32PrefixAccPub   = Bech32MainPrefix + "pub"
	Bech32PrefixValAddr  = Bech32MainPrefix + "valoper"
	Bech32PrefixValPub   = Bech32MainPrefix + "valoperpub"
	Bech32PrefixConsAddr = Bech32MainPrefix + "valcons"
	Bech32PrefixConsPub  = Bech32MainPrefix + "valconspub"
)
