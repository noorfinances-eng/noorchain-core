package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PendingMint represents an internal record of a PoSS reward that
// SHOULD be minted and distributed later.
//
// PoSS Logic 21: this is a pure "planning" structure. There is still
// NO real minting, NO bank transfers. It is only stored in the KVStore
// for future processing (when NOORCHAIN is live and PoSS is enabled).
type PendingMint struct {
	// BlockHeight is the block at which the signal was processed.
	BlockHeight int64 `json:"block_height" yaml:"block_height"`

	// Timestamp is the block time when the signal was processed.
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`

	// Participant is the NOOR account (noor1...) that should receive 70 %.
	Participant string `json:"participant" yaml:"participant"`

	// Curator is the NOOR curator account (noor1...) that should receive 30 %.
	Curator string `json:"curator" yaml:"curator"`

	// SignalType is the PoSS signal family (micro_donation, participation, ...).
	SignalType SignalType `json:"signal_type" yaml:"signal_type"`

	// ParticipantReward is the 70 % coin that should be distributed later.
	ParticipantReward sdk.Coin `json:"participant_reward" yaml:"participant_reward"`

	// CuratorReward is the 30 % coin that should be distributed later.
	CuratorReward sdk.Coin `json:"curator_reward" yaml:"curator_reward"`
}
