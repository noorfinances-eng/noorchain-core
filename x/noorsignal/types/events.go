package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -----------------------------------------------------------------------------
// Noms des événements PoSS
// -----------------------------------------------------------------------------

const (
	// Event émis lorsqu'un participant soumet un nouveau signal PoSS.
	EventTypeSignalSubmitted = "poss.signal_submitted"

	// Event émis lorsqu'un curator valide un signal.
	EventTypeSignalValidated = "poss.signal_validated"

	// Clés d'attributs communes.
	AttributeKeySignalID          = "signal_id"
	AttributeKeyParticipant       = "participant"
	AttributeKeyCurator           = "curator"
	AttributeKeyWeight            = "weight"
	AttributeKeyTotalReward       = "total_reward"
	AttributeKeyRewardParticipant = "reward_participant"
	AttributeKeyRewardCurator     = "reward_curator"
	AttributeKeyBlockHeight       = "block_height"
)

// -----------------------------------------------------------------------------
// Helpers pour construire les événements
// -----------------------------------------------------------------------------

// NewSignalSubmittedEvent construit un événement poss.signal_submitted
// à partir d'un Signal fraîchement créé.
func NewSignalSubmittedEvent(sig Signal, blockHeight int64) sdk.Event {
	return sdk.NewEvent(
		EventTypeSignalSubmitted,
		sdk.NewAttribute(AttributeKeySignalID, fmt.Sprintf("%d", sig.Id)),
		sdk.NewAttribute(AttributeKeyParticipant, sig.Participant.String()),
		sdk.NewAttribute(AttributeKeyWeight, fmt.Sprintf("%d", sig.Weight)),
		sdk.NewAttribute(AttributeKeyBlockHeight, fmt.Sprintf("%d", blockHeight)),
	)
}

// NewSignalValidatedEvent construit un événement poss.signal_validated
// à partir d'un Signal qui vient d'être validé (avec rewards).
func NewSignalValidatedEvent(sig Signal, blockHeight int64) sdk.Event {
	return sdk.NewEvent(
		EventTypeSignalValidated,
		sdk.NewAttribute(AttributeKeySignalID, fmt.Sprintf("%d", sig.Id)),
		sdk.NewAttribute(AttributeKeyParticipant, sig.Participant.String()),
		sdk.NewAttribute(AttributeKeyCurator, sig.Curator.String()),
		sdk.NewAttribute(AttributeKeyWeight, fmt.Sprintf("%d", sig.Weight)),
		sdk.NewAttribute(AttributeKeyTotalReward, fmt.Sprintf("%d", sig.TotalReward)),
		sdk.NewAttribute(AttributeKeyRewardParticipant, fmt.Sprintf("%d", sig.RewardParticipant)),
		sdk.NewAttribute(AttributeKeyRewardCurator, fmt.Sprintf("%d", sig.RewardCurator)),
		sdk.NewAttribute(AttributeKeyBlockHeight, fmt.Sprintf("%d", blockHeight)),
	)
}
