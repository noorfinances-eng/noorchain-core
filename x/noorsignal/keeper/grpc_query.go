package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
	noorsignalproto "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types" // pour les types proto générés
)

// QueryServer implémente le service gRPC défini dans query.proto.
type QueryServer struct {
	Keeper
}

// NewQueryServer construit une instance de QueryServer.
func NewQueryServer(k Keeper) QueryServer {
	return QueryServer{Keeper: k}
}

// -----------------------------------------------------------------------------
// Query : Signal(id)
// -----------------------------------------------------------------------------

func (qs QueryServer) Signal(
	goCtx context.Context,
	req *noorsignalproto.QuerySignalRequest,
) (*noorsignalproto.QuerySignalResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	sig, found := qs.Keeper.GetSignal(ctx, req.Id)
	if !found {
		return &noorsignalproto.QuerySignalResponse{}, nil
	}

	resp := noorsignalproto.QuerySignalResponse{
		Signal: convertSignalToProto(sig),
	}

	return &resp, nil
}

// -----------------------------------------------------------------------------
// Query : Signals() — V1 simple, tous les signaux
// -----------------------------------------------------------------------------

func (qs QueryServer) Signals(
	goCtx context.Context,
	req *noorsignalproto.QuerySignalsRequest,
) (*noorsignalproto.QuerySignalsResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := qs.Keeper.SignalStore(ctx)

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var list []*noorsignalproto.SignalPoSS

	for ; iterator.Valid(); iterator.Next() {
		var sig noorsignaltypes.Signal
		qs.Keeper.Cdc().MustUnmarshal(iterator.Value(), &sig)
		list = append(list, convertSignalToProto(sig))
	}

	return &noorsignalproto.QuerySignalsResponse{
		Signals: list,
	}, nil
}

// -----------------------------------------------------------------------------
// Query : Curator(address)
// -----------------------------------------------------------------------------

func (qs QueryServer) Curator(
	goCtx context.Context,
	req *noorsignalproto.QueryCuratorRequest,
) (*noorsignalproto.QueryCuratorResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return &noorsignalproto.QueryCuratorResponse{}, err
	}

	cur, found := qs.Keeper.GetCurator(ctx, addr)
	if !found {
		return &noorsignalproto.QueryCuratorResponse{}, nil
	}

	resp := noorsignalproto.QueryCuratorResponse{
		Curator: &noorsignalproto.CuratorPoSS{
			Address:                cur.Address.String(),
			Level:                  cur.Level,
			TotalSignalsValidated:  cur.TotalSignalsValidated,
			Active:                 cur.Active,
		},
	}

	return &resp, nil
}

// -----------------------------------------------------------------------------
// Query : Config()
// -----------------------------------------------------------------------------

func (qs QueryServer) Config(
	goCtx context.Context,
	req *noorsignalproto.QueryConfigRequest,
) (*noorsignalproto.QueryConfigResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	cfg, found := qs.Keeper.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	resp := noorsignalproto.QueryConfigResponse{
		Config: &noorsignalproto.PossConfigPoSS{
			BaseReward:       cfg.BaseReward,
			ParticipantShare: cfg.ParticipantShare,
			CuratorShare:     cfg.CuratorShare,
			MaxSignalsPerDay: cfg.MaxSignalsPerDay,
			Enabled:          cfg.Enabled,
			EraIndex:         cfg.EraIndex,
		},
	}

	return &resp, nil
}

// -----------------------------------------------------------------------------
// Helper : convertir un Signal (Go) vers SignalPoSS (proto)
// -----------------------------------------------------------------------------

func convertSignalToProto(sig noorsignaltypes.Signal) *noorsignalproto.SignalPoSS {
	var curator string
	if sig.Curator != nil {
		curator = sig.Curator.String()
	}

	return &noorsignalproto.SignalPoSS{
		Id:                sig.Id,
		Participant:       sig.Participant.String(),
		Curator:           curator,
		Weight:            sig.Weight,
		Time:              sig.Time.Format(time.RFC3339),
		Metadata:          sig.Metadata,
		TotalReward:       sig.TotalReward,
		RewardParticipant: sig.RewardParticipant,
		RewardCurator:     sig.RewardCurator,
	}
}
