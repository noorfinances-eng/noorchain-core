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
// À ce stade, la structure contient déjà :
// - un stockage de configuration PoSS
// - un helper de calcul de récompenses à partir de la config.
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

// ComputeSignalRewardsFromConfig calcule les récompenses PoSS pour un signal
// en utilisant la configuration actuellement stockée dans le module.
//
// Paramètres :
// - ctx    : contexte d'exécution
// - weight : poids du signal (1, 2, 5, etc.)
// - era    : indice d'ère pour le halving (0 = aucune division, 1 = /2, etc.)
//
// Retourne :
// - total       : récompense totale (après halving)
// - participant : part pour le participant
// - curator     : part pour le curator
// - found       : booléen indiquant si une configuration PoSS était présente
//
// Remarque :
// - si aucune configuration n'est trouvée, la fonction retourne
//   (0, 0, 0, false).
func (k Keeper) ComputeSignalRewardsFromConfig(
	ctx sdk.Context,
	weight uint32,
	era uint32,
) (total uint64, participant uint64, curator uint64, found bool) {
	cfg, ok := k.GetConfig(ctx)
	if !ok {
		return 0, 0, 0, false
	}

	total, participant, curator = noorsignaltypes.ComputeSignalRewards(cfg, weight, era)
	return total, participant, curator, true
}

// ComputeSignalRewardsCurrentEra calcule les récompenses PoSS pour un signal
// en utilisant l'ère courante définie dans la configuration PoSS.
//
// Paramètres :
// - ctx    : contexte d'exécution
// - weight : poids du signal (1, 2, 5, etc.)
//
// Retourne :
// - total       : récompense totale (après halving avec cfg.EraIndex)
// - participant : part pour le participant
// - curator     : part pour le curator
// - found       : booléen indiquant si une configuration PoSS était présente
//
// Remarques :
// - cette fonction lit cfg.EraIndex et l'utilise comme paramètre "era"
//   pour le halving.
// - si aucune configuration n'est trouvée, elle retourne (0, 0, 0, false).
func (k Keeper) ComputeSignalRewardsCurrentEra(
	ctx sdk.Context,
	weight uint32,
) (total uint64, participant uint64, curator uint64, found bool) {
	cfg, ok := k.GetConfig(ctx)
	if !ok {
		return 0, 0, 0, false
	}

	era := cfg.EraIndex
	total, participant, curator = noorsignaltypes.ComputeSignalRewards(cfg, weight, era)
	return total, participant, curator, true
}
