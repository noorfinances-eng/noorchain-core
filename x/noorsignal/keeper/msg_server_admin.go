package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// ---------------------------------------------------------
// Messages admin : AddCurator / RemoveCurator / SetConfig
// ---------------------------------------------------------

// AddCurator : ajoute (ou réactive) un curator.
func (s MsgServer) AddCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgAddCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validation basique du message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// ⚠️ V1 : on ne fait pas encore de vérification forte de "Authority"
	// (gov, adresse fondation, etc.). On vérifiera ça plus tard.
	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	// Créer ou réactiver le curator
	curator := noorsignaltypes.Curator{
		Address: curatorAddr,
		Active:  true,
		// TotalSignalsValidated reste à 0 par défaut
	}

	s.Keeper.SetCurator(ctx, curator)

	// V1 : pas encore d'events spécifiques, on se contente du résultat OK.
	return &sdk.Result{}, nil
}

// RemoveCurator : désactive un curator existant.
func (s MsgServer) RemoveCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgRemoveCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// V1 : pas de check avancé sur Authority, on gère plus tard.
	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	curator, found := s.Keeper.GetCurator(ctx, curatorAddr)
	if !found {
		return nil, errors.New("curator not found")
	}

	curator.Active = false
	s.Keeper.SetCurator(ctx, curator)

	return &sdk.Result{}, nil
}

// SetConfig : met à jour la configuration globale PoSS.
// V1 : on se contente de s'assurer qu'une config existe, sans encore
// appliquer tous les champs du message (BaseReward, ratios, etc.).
func (s MsgServer) SetConfig(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSetConfig,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// V1 : on ne fait pas encore une logique fine d'authority.
	// On se contente de garantir qu'une config existe.
	cfg, found := s.Keeper.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	// TODO (plus tard) : appliquer réellement les champs de msg à cfg.

	s.Keeper.SetConfig(ctx, cfg)
	return &sdk.Result{}, nil
}
