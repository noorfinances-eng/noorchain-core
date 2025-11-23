package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/store/prefix"
)

// Nom officiel du module
const ModuleName = "noorsignal"

// Clé du store principal
const StoreKey = ModuleName

// Router key
const RouterKey = ModuleName

// -----------------------------------------------------------------------------
// Prefixes des sous-stores
// -----------------------------------------------------------------------------

var (
	// Configuration globale PoSS
	KeyPrefixPoSSConfig = []byte{0x01}

	// Signaux
	KeyPrefixSignal = []byte{0x02}

	// Curateurs
	KeyPrefixCurator = []byte{0x03}

	// Compteurs journaliers (anti-abus)
	KeyPrefixDailyCounter = []byte{0x04}

	// Next Signal ID
	KeyNextSignalID = []byte{0x05}
)

// -----------------------------------------------------------------------------
//  Sous-stores (prefix stores)
// -----------------------------------------------------------------------------

func GetConfigStore(parent storetypes.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixPoSSConfig)
}

func GetSignalStore(parent storetypes.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixSignal)
}

func GetCuratorStore(parent storetypes.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixCurator)
}

func GetDailyCounterStore(parent storetypes.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixDailyCounter)
}

// -----------------------------------------------------------------------------
//  Clés utilitaires
// -----------------------------------------------------------------------------

// SignalKey construit la clé d’un signal à partir de son ID.
func SignalKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// DailyCounterKey construit la clé d’un compteur journalier.
func DailyCounterKey(addr sdk.AccAddress, dayBucket uint64) []byte {
	// format : address || dayBucket (8 bytes)
	day := make([]byte, 8)
	binary.BigEndian.PutUint64(day, dayBucket)

	return append(addr.Bytes(), day...)
}
