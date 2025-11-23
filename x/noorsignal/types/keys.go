package types

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
var (
	// KeyPrefixSignals indique le préfixe pour les enregistrements de signaux.
	KeyPrefixSignals = []byte{0x01}

	// KeyPrefixCurators indique le préfixe pour les enregistrements de curators.
	KeyPrefixCurators = []byte{0x02}

	// KeyPrefixConfig indique le préfixe pour la configuration globale PoSS.
	KeyPrefixConfig = []byte{0x03}

	// KeyPrefixDailyCounters indique le préfixe pour les compteurs quotidiens
	// "(participant, jour) -> nombre de signaux".
	KeyPrefixDailyCounters = []byte{0x04}
)

// GetSignalStore retourne un store préfixé pour les signaux.
func GetSignalStore(parent sdk.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixSignals)
}

// GetCuratorStore retourne un store préfixé pour les curators.
func GetCuratorStore(parent sdk.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixCurators)
}

// GetConfigStore retourne un store préfixé pour la configuration globale PoSS.
func GetConfigStore(parent sdk.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixConfig)
}

// GetDailyCounterStore retourne un store préfixé pour les compteurs
// quotidiens PoSS.
func GetDailyCounterStore(parent sdk.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixDailyCounters)
}

// SignalKey construit la clé de stockage pour un signal individuel.
func SignalKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// DailyCounterKey construit la clé "(participant, jour)" utilisée pour
// stocker le nombre de signaux émis par un participant donné sur un jour donné.
//
// Paramètres :
// - addr      : adresse du participant
// - dayBucket : indice de jour (ex: timestampUnix / 86400)
//
// Format : key = addr.Bytes() || dayBucket(8 octets big-endian)
func DailyCounterKey(addr sdk.AccAddress, dayBucket uint64) []byte {
	addrBz := addr.Bytes()

	dayBz := make([]byte, 8)
	binary.BigEndian.PutUint64(dayBz, dayBucket)

	return append(addrBz, dayBz...)
}
