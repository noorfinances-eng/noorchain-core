package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// QueryServer implémente l’interface générée par les fichiers .proto.
// Il expose les endpoints gRPC publics du module PoSS.
type QueryServer struct {
	Keeper
}

// NewQueryServer crée un nouveau QueryServer basé sur un Keeper PoSS.
func NewQueryServer(k Keeper) QueryServer {
	return QueryServer{Keeper: k}
}

// ------------------------------------------------------------
// GET SIGNAL BY ID
// ------------------------------------------------------------
func (qs QueryServer) GetSignal(
	goCtx context.Context,
	req *noorsignaltypes.QueryGetSignalRequest,
) (*noorsignaltypes.QueryGetSignalResponse, error) {
	if req == nil {
		return nil, sdk.ErrInvalidRequest.Wrap("nil request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sig, found := qs.Keeper.GetSignal(ctx, req.SignalId)
	if !found {
		return nil, sdk.ErrNotFound.Wrap("signal not found")
	}

	// Conversion en type proto
	p := noorsignaltypes.SignalToProto(sig)

	return &noorsignaltypes.QueryGetSignalResponse{
		Signal: p,
	}, nil
}

// ------------------------------------------------------------
// LIST SIGNALS (PAGINATED)
// ------------------------------------------------------------
func (qs QueryServer) ListSignals(
	goCtx context.Context,
	req *noorsignaltypes.QueryListSignalsRequest,
) (*noorsignaltypes.QueryListSignalsResponse, error) {
	if req == nil {
		return nil, sdk.ErrInvalidRequest.Wrap("nil request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := qs.signalStore(ctx)

	var results []*noorsignaltypes.Signal

	// Utilisation de la pagination Cosmos standard.
	pageRes, err := query.Paginate(store, req.Pagination, func(key, value []byte) error {
		var sig noorsignaltypes.Signal
		qs.cdc.MustUnmarshal(value, &sig)

		results = append(results, noorsignaltypes.SignalToProto(sig))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &noorsignaltypes.QueryListSignalsResponse{
		Signals:    results,
		Pagination: pageRes,
	}, nil
}

// ------------------------------------------------------------
// GET CURATOR BY ADDRESS
// ------------------------------------------------------------
func (qs QueryServer) GetCurator(
	goCtx context.Context,
	req *noorsignaltypes.QueryGetCuratorRequest,
) (*noorsignaltypes.QueryGetCuratorResponse, error) {
	if req == nil {
		return nil, sdk.ErrInvalidRequest.Wrap("nil request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	cur, found := qs.Keeper.GetCurator(ctx, addr)
	if !found {
		return nil, sdk.ErrNotFound.Wrap("curator not found")
	}

	return &noorsignaltypes.QueryGetCuratorResponse{
		Curator: noorsignaltypes.CuratorToProto(cur),
	}, nil
}

// ------------------------------------------------------------
// LIST CURATORS (PAGINATED)
// ------------------------------------------------------------
func (qs QueryServer) ListCurators(
	goCtx context.Context,
	req *noorsignaltypes.QueryListCuratorsRequest,
) (*noorsignaltypes.QueryListCuratorsResponse, error) {
	if req == nil {
		return nil, sdk.ErrInvalidRequest.Wrap("nil request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := qs.curatorStore(ctx)

	var results []*noorsignaltypes.Curator

	pageRes, err := query.Paginate(store, req.Pagination, func(key, value []byte) error {
		var curator noorsignaltypes.Curator
		qs.cdc.MustUnmarshal(value, &curator)

		results = append(results, noorsignaltypes.CuratorToProto(curator))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &noorsignaltypes.QueryListCuratorsResponse{
		Curators:   results,
		Pagination: pageRes,
	}, nil
}

// ------------------------------------------------------------
// GET CONFIG
// ------------------------------------------------------------
func (qs QueryServer) GetConfig(
	goCtx context.Context,
	req *noorsignaltypes.QueryGetConfigRequest,
) (*noorsignaltypes.QueryGetConfigResponse, error) {
	if req == nil {
		return nil, sdk.ErrInvalidRequest.Wrap("nil request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	cfg, _ := qs.Keeper.GetConfig(ctx) // si absent → DefaultConfig lors du genesis

	return &noorsignaltypes.QueryGetConfigResponse{
		Config: noorsignaltypes.ConfigToProto(cfg),
	}, nil
}
