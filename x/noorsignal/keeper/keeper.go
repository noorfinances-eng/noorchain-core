package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper est le keeper minimal du module x/noorsignal (PoSS).
// Pour l’instant :
//
// - pas de BankKeeper, pas de StakingKeeper,
// - pas encore de params PoSS,
// - juste un point central propre pour le module et un logger.
//
// On l’étendra plus tard quand on branchera la vraie logique PoSS.
type Keeper struct {
	// Codec binaire (Proto)
	cdc codec.BinaryCodec

	// storeKey permet d’accéder au KVStore du module (x/noorsignal).
	storeKey storetypes.StoreKey
}

// NewKeeper crée un nouveau Keeper PoSS minimal.
// On ajoutera plus tard : params, hooks, liens avec Bank/Staking, etc.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// Logger retourne un logger spécifique au module x/noorsignal.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", noorsignaltypes.ModuleName)
}
