package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// AddCurator gère la réception d'un MsgAddCurator.
// Il permet à une "authority" d'ajouter ou de réactiver un Curator
// avec un certain niveau (BRONZE / SILVER / GOLD).
func (s MsgServer) AddCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgAddCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Vérifier que l'authority est présente.
	// TODO (plus tard) : vérifier que cette adresse correspond bien
	// à une fondation / multisig autorisée ou à un module gov.
	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	// 2) Convertir l'adresse du curator.
	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid curator address")
	}

	// 3) Récupérer un éventuel Curator existant.
	curator, found := s.Keeper.GetCurator(ctx, curatorAddr)
	if !found {
		// Nouveau Curator.
		curator = noorsignaltypes.Curator{
			Address:               curatorAddr,
			Level:                 msg.Level,
			TotalSignalsValidated: 0,
			Active:                true,
		}
	} else {
		// Mise à jour d'un Curator existant.
		curator.Level = msg.Level
		curator.Active = true
	}

	// 4) Enregistrer le Curator dans le store.
	s.Keeper.SetCurator(ctx, curator)

	return &sdk.Result{}, nil
}

// RemoveCurator gère la réception d'un MsgRemoveCurator.
// Il permet à une "authority" de désactiver un Curator existant.
func (s MsgServer) RemoveCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgRemoveCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Vérifier que l'authority est présente.
	// TODO (plus tard) : vérifier que cette adresse correspond bien
	// à une fondation / multisig autorisée ou à un module gov.
	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	// 2) Convertir l'adresse du curator.
	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid curator address")
	}

	// 3) Récupérer le Curator.
	curator, found := s.Keeper.GetCurator(ctx, curatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "curator not found")
	}

	// 4) Le désactiver (on garde l'historique).
	curator.Active = false

	// 5) Enregistrer la mise à jour.
	s.Keeper.SetCurator(ctx, curator)

	return &sdk.Result{}, nil
}

// SetConfig gère la réception d'un MsgSetConfig.
// Il permet à une "authority" de mettre à jour la configuration PoSS
// (base_reward, max_signals_per_day, era_index, ratios 70/30, etc.)
// sans toucher au genesis.
func (s MsgServer) SetConfig(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSetConfig,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Vérifier que l'authority est présente.
	// TODO (plus tard) : vérifier que cette adresse correspond bien
	// à une fondation / multisig autorisée ou à un module gov.
	if msg.Authority == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing authority")
	}

	// 2) Parser la base de récompense depuis la string.
	baseReward, err := strconv.ParseUint(msg.BaseReward, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid base_reward")
	}

	// 3) Vérifier les ratios participant / curator.
	totalRatio := msg.ParticipantRatio + msg.CuratorRatio
	if totalRatio == 0 || totalRatio != 100 {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			"participant_ratio + curator_ratio must be exactly 100",
		)
	}

	// 4) Construire une nouvelle configuration PoSS.
	// EraIndex dans PossConfig est un uint32, on caste depuis le uint64 du message.
	if msg.EraIndex > uint64(^uint32(0)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "era_index out of range")
	}

	newCfg := noorsignaltypes.PossConfig{
		BaseReward:       baseReward,
		ParticipantShare: msg.ParticipantRatio,
		CuratorShare:     msg.CuratorRatio,
		MaxSignalsPerDay: msg.MaxSignalsPerDay,
		Enabled:          true,
		EraIndex:         uint32(msg.EraIndex),
	}

	// 5) Enregistrer cette nouvelle config dans le store PoSS.
	s.Keeper.SetConfig(ctx, newCfg)

	return &sdk.Result{}, nil
}
