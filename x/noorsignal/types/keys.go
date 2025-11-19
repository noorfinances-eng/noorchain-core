package types

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
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
//
// Remarque :
// - Chaque préfixe est un byte distinct.
// - On les utilise avec prefix.NewStore(...) dans le keeper.
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

// GetDailyCounterStore retourne un store préfixé pour les compteurs
// quotidiens PoSS.
func GetDailyCounterStore(parent prefix.Store) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixDailyCounters)
}

// SignalKey construit la clé de stockage pour un signal individuel
// à partir de son identifiant numérique.
//
// IMPORTANT :
// - le store utilisé dans le keeper est déjà préfixé par KeyPrefixSignals,
//   via GetSignalStore(...) et prefix.NewStore(...).
// - la clé retournée ici est donc UNIQUEMENT l'ID encodé sur 8 octets
//   (sans ré-ajouter le préfixe).
func SignalKey(id uint64) []byte {
	// 8 octets pour uint64 en big-endian.
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
// Format :
//   key = addr.Bytes() || dayBucket(8 octets big-endian)
//
// Le préfixe KeyPrefixDailyCounters est déjà appliqué par
// GetDailyCounterStore, donc on ne le rajoute pas ici.
func DailyCounterKey(addr sdk.AccAddress, dayBucket uint64) []byte {
	addrBz := addr.Bytes()

	dayBz := make([]byte, 8)
	binary.BigEndian.PutUint64(dayBz, dayBucket)

	// concaténation : [addr || day]
	return append(addrBz, dayBz...)
}
