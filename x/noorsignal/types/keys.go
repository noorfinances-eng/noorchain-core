package types

// ModuleName est le nom officiel du module PoSS / NOOR Signal.
const ModuleName = "noorsignal"

// StoreKey est la clé du KVStore principal du module.
const StoreKey = ModuleName

// RouterKey est utilisé si on route des messages via le router Cosmos.
const RouterKey = ModuleName

// Préfixes des différents sous-stores.
// On les utilisera avec prefix.NewStore dans le keeper.
var (
    // Configuration globale du PoSS (limits, ratios 70/30, etc.)
    KeyPrefixPoSSConfig = []byte{0x01}

    // Signaux (un enregistrement par signal validé)
    KeyPrefixSignal = []byte{0x02}

    // Curateurs (infos sur les curators enregistrés)
    KeyPrefixCurator = []byte{0x03}
)
