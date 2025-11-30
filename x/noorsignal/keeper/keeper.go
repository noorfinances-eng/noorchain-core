package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Keeper est le keeper minimal du module x/noorsignal (PoSS).
// Pour l’instant, il ne fait rien : on prépare seulement la structure.
type Keeper struct {
	// Codec pour encoder/décoder l’état du module.
	cdc codec.Codec

	// storeKey permet d’accéder au KVStore du module.
	storeKey storetypes.StoreKey
}

// NewKeeper crée un nouveau Keeper PoSS minimal.
// On ajoutera plus tard : params, hooks, liens avec Bank/Staking, etc.
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}
