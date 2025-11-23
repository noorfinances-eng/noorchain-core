package types

// ModuleName définit le nom du module NOORSIGNAL.
// Il doit être utilisé partout où le module est référencé
// (AppModule, genesis, routing, etc.).
const (
	ModuleName = "noorsignal"

	// StoreKey est la clé principale du KVStore pour ce module.
	StoreKey = ModuleName

	// RouterKey est utilisé pour router les messages vers ce module.
	RouterKey = ModuleName

	// QuerierRoute est le nom utilisé pour les requêtes (legacy).
	// Même si on n'utilise pas gRPC ici, on garde ce champ pour compatibilité.
	QuerierRoute = ModuleName
)

// On peut définir ici des préfixes de clés pour le KVStore.
// Pour l'instant, le module V1 est minimal, mais ces préfixes
// serviront dès qu'on ajoutera un stockage (signals, curators, etc.).
var (
	// KeyPrefixParams est un préfixe générique pour les paramètres du module.
	KeyPrefixParams = []byte{0x01}

	// KeyPrefixExample est un préfixe d'exemple pour illustrer une future
	// clé de stockage (par ex. un compteur global).
	KeyPrefixExample = []byte{0x02}
)

// KeyPrefix transforme une string en préfixe binaire.
// Utile pour définir des clés dérivées dans le KVStore.
func KeyPrefix(p string) []byte {
	return []byte(p)
}
