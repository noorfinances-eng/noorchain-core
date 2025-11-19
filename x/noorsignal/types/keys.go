package types

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// ModuleName est le nom officiel du module PoSS dans NOORCHAIN.
const ModuleName = "noorsignal"

// StoreKey est la clé principale du KVStore pour le module.
const StoreKey = ModuleName

// RouterKey est utilisé pour router les messages Msg vers ce module.
const RouterKey = ModuleName

// QuerierRoute peut être utilisé pour les requêtes legacy (si besoin).
const QuerierRoute = ModuleName

// Clé spéciale (non préfixée) pour stocker le prochain identifiant
// de signal auto-incrémenté.
var (
	KeyNextSignalID = []byte("next_signal_id")
)

// Prefixes pour différents types de données stockées par le module PoSS.
//
// Remarque :
// - Chaque préfixe est un byte distinct.
// - On pourra les utiliser avec prefix.NewStore(...) dans le keeper.
//
// Exemples futurs :
// - stockage des signaux
// - stockage des curators
// - stockage de la config PoSS (barème, limites...)
var (
	// KeyPrefixSignals indique le préfixe pour les enregistrements de signaux.
	KeyPrefixSignals = []byte{0x01}

	// KeyPrefixCurators indique le préfixe pour les enregistrements de curators.
	KeyPrefixCurators = []byte{0x02}

	// KeyPrefixConfig indique le préfixe pour la configuration globale PoSS.
	KeyPrefixConfig = []byte{0x03}
)

// GetSignalStore retourne un store préfixé pour les signaux.
func GetSignalStore(parent prefix.Store) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixSignals)
}

// GetCuratorStore retourne un store préfixé pour les curators.
func GetCuratorStore(parent prefix.Store) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixCurators)
}

// GetConfigStore retourne un store préfixé pour la configuration globale PoSS.
func GetConfigStore(parent prefix.Store) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixConfig)
}

// SignalKey construit la clé de stockage pour un signal individuel
// à partir de son identifiant numérique.
//
// Format choisi :
// - KeyPrefixSignals (0x01)
// - suivi de l'identifiant encodé sur 8 octets (big-endian).
//
// Cela permet de parcourir les signaux dans l'ordre des IDs.
func SignalKey(id uint64) []byte {
	// 8 octets pour uint64 en big-endian.
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	// Préfixe + ID encodé.
	return append(KeyPrefixSignals, bz...)
}
