package app

// params.go
//
// Fichier réservé pour de futurs helpers de paramètres globaux NOORCHAIN.
//
// IMPORTANT :
// - Toutes les constantes globales de configuration (AppName, ChainID*,
//   Bech32MainPrefix, Bech32PrefixAccAddr / Val / Cons*, CoinDenom,
//   CoinDisplayDenom, CoinDecimals, EvmChainID, etc.) sont définies
//   une seule fois dans config.go.
// - Ne pas redéfinir ces constantes ici, pour éviter les conflits
//   de compilation ("redeclared in this block").
//
// Si plus tard tu as besoin de fonctions utilitaires liées aux paramètres
// (ex. helpers sur les denoms, conversions d’unités, formatage d’adresses),
// tu pourras les ajouter dans ce fichier, sans y redéclarer de const/var
// déjà présentes dans config.go.
