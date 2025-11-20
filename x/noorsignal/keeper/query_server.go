package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	query "github.com/cosmos/cosmos-sdk/types/query"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// QueryServer implémente le service gRPC "Query" du module PoSS (noorsignal).
// Il s'appuie sur le Keeper interne pour accéder au store.
type QueryServer struct {
	Keeper
}

// NewQueryServer construit un QueryServer à partir d'un Keeper PoSS.
func NewQueryServer(k Keeper) QueryServer {
	return QueryServer{Keeper: k}
}

// Signal retourne un signal unique via son identifiant.
func (q QueryServer) Signal(
	goCtx context.Context,
	req *noorsignaltypes.QuerySignalRequest,
) (*noorsignaltypes.QuerySignalResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sig, found := q.Keeper.GetSignal(ctx, req.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "signal not found")
	}

	// Ici, nous renvoyons directement la struct Signal Go définie dans types.go.
	// La version proto générée (plus tard) pourra mapper exactement cette structure.
	return &noorsignaltypes.QuerySignalResponse{
		Signal: &sig,
	}, nil
}

// Signals retourne une liste paginée de signaux PoSS.
func (q QueryServer) Signals(
	goCtx context.Context,
	req *noorsignaltypes.QuerySignalsRequest,
) (*noorsignaltypes.QuerySignalsResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := q.signalStore(ctx) // prefix.Store sur les signaux
	var signals []noorsignaltypes.Signal

	pageRes, err := query.Paginate(
		store,
		req.Pagination,
		func(key, value []byte) error {
			var sig noorsignaltypes.Signal
			q.Keeper.cdc.MustUnmarshal(value, &sig)
			signals = append(signals, sig)
			return nil
		},
	)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to paginate signals")
	}

	return &noorsignaltypes.QuerySignalsResponse{
		Signals:    signals,
		Pagination: pageRes,
	}, nil
}

// Curator retourne les informations d'un Curator via son adresse.
func (q QueryServer) Curator(
	goCtx context.Context,
	req *noorsignaltypes.QueryCuratorRequest,
) (*noorsignaltypes.QueryCuratorResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty request")
	}
	if req.Address == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid curator address")
	}

	curator, found := q.Keeper.GetCurator(ctx, addr)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "curator not found")
	}

	return &noorsignaltypes.QueryCuratorResponse{
		Curator: &curator,
	}, nil
}

// Config retourne la configuration PoSS actuelle.
func (q QueryServer) Config(
	goCtx context.Context,
	req *noorsignaltypes.QueryConfigRequest,
) (*noorsignaltypes.QueryConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cfg, found := q.Keeper.GetConfig(ctx)
	if !found {
		// Si aucune config n'est encore enregistrée, on renvoie la config par défaut.
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	return &noorsignaltypes.QueryConfigResponse{
		Config: cfg,
	}, nil
}

// DailyCount retourne le nombre de signaux émis par une adresse pour un
// "day bucket" donné (par ex. block_time / 86400).
func (q QueryServer) DailyCount(
	goCtx context.Context,
	req *noorsignaltypes.QueryDailyCountRequest,
) (*noorsignaltypes.QueryDailyCountResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty request")
	}
	if req.Address == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid address")
	}

	count := q.Keeper.getDailySignalCount(ctx, addr, req.Day)

	return &noorsignaltypes.QueryDailyCountResponse{
		Count: count,
	}, nil
}

// signalStore retourne un prefix.Store explicitement typé ici pour le QueryServer.
// On redélègue au helper existant du Keeper.
func (q QueryServer) signalStore(ctx sdk.Context) prefix.Store {
	return q.Keeper.signalStore(ctx)
}
