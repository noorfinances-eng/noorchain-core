package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ------------------------------------------------------------
//  Event Types
// ------------------------------------------------------------

// Event type for submitting a signal.
const EventTypeSignalSubmitted = "poss.signal_submitted"

// Event type for validating a signal.
const EventTypeSignalValidated = "poss.signal_validated"

// ------------------------------------------------------------
//  Attributes Keys
// ------------------------------------------------------------

// Common attributes
const (
	AttrKeySignalID   = "signal_id"
	AttrKeyParticipant = "participant"
	AttrKeyCurator     = "curator"
	AttrKeyWeight      = "weight"
	AttrKeyMetadata    = "metadata"

	AttrKeyRewardTotal       = "reward_total"
	AttrKeyRewardParticipant = "reward_participant"
	AttrKeyRewardCurator     = "reward_curator"
)

// ------------------------------------------------------------
//  Helpers for building events (optional but clean)
// ------------------------------------------------------------

func NewEventSignalSubmitted(
	signalID uint64,
	participant string,
	weight uint32,
	metadata string,
) sdk.Event {
	return sdk.NewEvent(
		EventTypeSignalSubmitted,
		sdk.NewAttribute(AttrKeySignalID, uintToStr(signalID)),
		sdk.NewAttribute(AttrKeyParticipant, participant),
		sdk.NewAttribute(AttrKeyWeight, uintToStr32(weight)),
		sdk.NewAttribute(AttrKeyMetadata, metadata),
	)
}

func NewEventSignalValidated(
	signalID uint64,
	participant string,
	curator string,
	total uint64,
	part uint64,
	cur uint64,
) sdk.Event {
	return sdk.NewEvent(
		EventTypeSignalValidated,
		sdk.NewAttribute(AttrKeySignalID, uintToStr(signalID)),
		sdk.NewAttribute(AttrKeyParticipant, participant),
		sdk.NewAttribute(AttrKeyCurator, curator),
		sdk.NewAttribute(AttrKeyRewardTotal, uintToStr(total)),
		sdk.NewAttribute(AttrKeyRewardParticipant, uintToStr(part)),
		sdk.NewAttribute(AttrKeyRewardCurator, uintToStr(cur)),
	)
}

// ------------------------------------------------------------
//  Internal small helpers
// ------------------------------------------------------------

func uintToStr(v uint64) string {
	return sdk.Uint64ToString(v)
}

func uintToStr32(v uint32) string {
	return sdk.Uint64ToString(uint64(v))
}
