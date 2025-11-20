package types

import (
	"time"

	"github.com/gogo/protobuf/types"
)

// ------------------------------------------------------------
//  Conversion Helper : Signal → Proto (query.proto)
// ------------------------------------------------------------
func SignalToProto(sig Signal) *SignalProto {
	// Convertir time.Time → google.protobuf.Timestamp
	var ts *types.Timestamp
	if !sig.Time.IsZero() {
		ts = &types.Timestamp{
			Seconds: sig.Time.Unix(),
			Nanos:   int32(sig.Time.Nanosecond()),
		}
	}

	return &SignalProto{
		Id:               sig.Id,
		Participant:      sig.Participant.String(),
		Curator:          sig.Curator.String(),
		Weight:           sig.Weight,
		Time:             ts,
		Metadata:         sig.Metadata,
		TotalReward:      sig.TotalReward,
		RewardParticipant: sig.RewardParticipant,
		RewardCurator:     sig.RewardCurator,
	}
}

// ------------------------------------------------------------
//  Conversion Helper : Curator → Proto
// ------------------------------------------------------------
func CuratorToProto(cur Curator) *CuratorProto {
	return &CuratorProto{
		Address:              cur.Address.String(),
		Level:                cur.Level,
		TotalSignalsValidated: cur.TotalSignalsValidated,
		Active:               cur.Active,
	}
}

// ------------------------------------------------------------
//  Conversion Helper : PossConfig → Proto
// ------------------------------------------------------------
func ConfigToProto(cfg PossConfig) *PossConfigProto {
	return &PossConfigProto{
		BaseReward:       cfg.BaseReward,
		ParticipantShare: cfg.ParticipantShare,
		CuratorShare:     cfg.CuratorShare,
		MaxSignalsPerDay: cfg.MaxSignalsPerDay,
		Enabled:          cfg.Enabled,
		EraIndex:         cfg.EraIndex,
	}
}
