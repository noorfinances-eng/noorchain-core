package types

// Basic identifiers for the PoSS module (x/noorsignal).
// For l'instant, on ne fait que définir les constantes de base.
// La logique réelle viendra dans les blocs suivants.

const (
	// ModuleName is the name of the PoSS module.
	ModuleName = "noorsignal"

	// StoreKey is the primary KVStore key for x/noorsignal.
	StoreKey = ModuleName

	// RouterKey is used for routing messages (Msg) to this module.
	RouterKey = ModuleName

	// QuerierRoute is used for legacy querier routing.
	QuerierRoute = ModuleName
)
