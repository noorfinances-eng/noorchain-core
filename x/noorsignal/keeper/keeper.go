package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper est le gestionnaire principal du module PoSS (noorsignal) pour NOORCHAIN.
//
// Rôle général :
// - lire / écrire les signaux dans le store
// - lire / écrire les curators
// - gérer la configuration globale PoSS
//
// À ce stade, la structure est un squelette : aucune logique métier
// (calcul des récompenses, halving, limites) n'est encore implémentée.
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.Codec

	// Plus tard, on pourra ajouter ici des références à d'autres keepers :
	// - BankKeeper (pour envoyer des NUR)
	// - StakingKeeper / GovKeeper (si nécessaire)
	// Pour l'instant, on garde le keeper le plus simple possible.
}

// NewKeeper construit un nouveau Keeper PoSS pour NOORCHAIN.
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// getStore retourne le KVStore brut du module à partir du contexte.
func (k Keeper) getStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

// signalStore retourne un store préfixé pour les signaux PoSS.
func (k Keeper) signalStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetSignalStore(parent)
}

// curatorStore retourne un store préfixé pour les curators PoSS.
func (k Keeper) curatorStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetCuratorStore(parent)
}

// configStore retourne un store préfixé pour la configuration PoSS.
func (k Keeper) configStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetConfigStore(parent)
}

// SetConfig enregistre la configuration globale PoSS dans le store.
//
// Cette méthode écrase simplement la configuration précédente.
func (k Keeper) SetConfig(ctx sdk.Context, cfg noorsignaltypes.PossConfig) {
	store := k.configStore(ctx)

	bz := k.cdc.MustMarshal(&cfg)
	store.Set([]byte("config"), bz)
}

// GetConfig lit la configuration globale PoSS depuis le store.
//
// Retourne :
// - la configuration (PossConfig)
// - un booléen "found" qui indique si une config existe déjà.
func (k Keeper) GetConfig(ctx sdk.Context) (noorsignaltypes.PossConfig, bool) {
	store := k.configStore(ctx)

	bz := store.Get([]byte("config"))
	if bz == nil {
		return noorsignaltypes.PossConfig{}, false
	}

	var cfg noorsignaltypes.PossConfig
	k.cdc.MustUnmarshal(bz, &cfg)

	return cfg, true
}

// InitDefaultConfig initialise la configuration PoSS avec les valeurs
// par défaut si aucune configuration n'est encore présente.
//
// Si une configuration existe déjà, cette fonction ne fait rien.
func (k Keeper) InitDefaultConfig(ctx sdk.Context) {
	_, found := k.GetConfig(ctx)
	if found {
		// Une configuration existe déjà : on ne la remplace pas.
		return
	}

	defaultCfg := noorsignaltypes.DefaultPossConfig()
	k.SetConfig(ctx, defaultCfg)
}
