package types

// Le module V1 est minimal :
// - Pas de storage complexe
// - Pas de Params
// - Pas de structures avancées
// - Juste une base propre pour accueillir les Msg + Genesis plus tard

// DefaultIndex est un exemple de valeur par défaut.
// Il servira pour illustrer un premier champ dans le genesis.
const DefaultIndex uint64 = 1

// GenesisState définit l'état du module au lancement (genesis).
// V1 : il est minimal. Ne contient qu'une valeur simple.
type GenesisState struct {
	// Un exemple de champ : un compteur global.
	// (Il peut être utilisé plus tard pour les IDs des signaux PoSS.)
	Counter uint64 `json:"counter" yaml:"counter"`
}

// NewGenesisState crée un nouvel état genesis avec des valeurs données.
func NewGenesisState(counter uint64) *GenesisState {
	return &GenesisState{
		Counter: counter,
	}
}

// DefaultGenesis renvoie un genesis minimal par défaut.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Counter: DefaultIndex,
	}
}
