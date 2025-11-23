package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -----------------------------------------------------------------------------
// Noms des événements PoSS
// -----------------------------------------------------------------------------

const (
	// Signal soumis par un participant
	EventTypeSignalSubmitted = "poss.signal_submitted"

	// Signal validé par un curator
	EventTypeSignalValidated = "poss.signal_validated"

	// Curator ajouté / mis à jour
	EventTypeCuratorAdded = "poss.curator_added"

	// Curator désactivé
	EventTypeCuratorRemoved = "poss.curator_removed"

	// Configuration PoSS mise à jour
	EventTypeConfigUpdated = "poss.config_updated"
)

// -----------------------------------------------------------------------------
// Clés d'attributs pour tous les events PoSS
// -----------------------------------------------------------------------------

const (
	AttrKeySignalID          = "signal_id"
	AttrKeyParticipant       = "participant"
	AttrKeyCurator           = "curator"
	AttrKeyWeight            = "weight"
	AttrKeyTotalReward       = "total_reward"
	AttrKeyRewardParticipant = "reward_participant"
	AttrKeyRewardCurator     = "reward_curator"
	AttrKeyBlockHeight       = "block_height"

	AttrKeyLevel             = "level"
	AttrKeyAuthority         = "authority"
	AttrKeyBaseReward        = "base_reward"
	AttrKeyMaxSignalsPerDay  = "max_signals_per_day"
	AttrKeyEraIndex          = "era_index"
	AttrKeyParticipantRatio  = "participant_ratio"
	AttrKeyCuratorRatio      = "curator_ratio"
)

// -----------------------------------------------------------------------------
// Helpers pour construire les événements de signaux
// -----------------------------------------------------------------------------

// NewSignalSubmittedEvent construit un événement poss.signal_submitted
// à partir d'un Signal fraîchement créé.
func NewSignalSubmittedEvent(sig Signal, blockHeight int64) sdk.Event {
	return sdk.NewEvent(
		EventTypeSignalSubmitted,
		sdk.NewAttribute(AttrKeySignalID, fmt.Sprintf("%d", sig.Id)),
		sdk.NewAttribute(AttrKeyParticipant, sig.Participant.String()),
		sdk.NewAttribute(AttrKeyWeight, fmt.Sprintf("%d", sig.Weight)),
		sdk.NewAttribute(AttrKeyBlockHeight, fmt.Sprintf("%d", blockHeight)),
	)
}

// NewSignalValidatedEvent construit un événement poss.signal_validated
// à partir d'un Signal qui vient d'être validé (avec rewards).
func NewSignalValidatedEvent(sig Signal, blockHeight int64) sdk.Event {
	return sdk.NewEvent(
		EventTypeSignalValidated,
		sdk.NewAttribute(AttrKeySignalID, fmt.Sprintf("%d", sig.Id)),
		sdk.NewAttribute(AttrKeyParticipant, sig.Participant.String()),
		sdk.NewAttribute(AttrKeyCurator, sig.Curator.String()),
		sdk.NewAttribute(AttrKeyWeight, fmt.Sprintf("%d", sig.Weight)),
		sdk.NewAttribute(AttrKeyTotalReward, fmt.Sprintf("%d", sig.TotalReward)),
		sdk.NewAttribute(AttrKeyRewardParticipant, fmt.Sprintf("%d", sig.RewardParticipant)),
		sdk.NewAttribute(AttrKeyRewardCurator, fmt.Sprintf("%d", sig.RewardCurator)),
		sdk.NewAttribute(AttrKeyBlockHeight, fmt.Sprintf("%d", blockHeight)),
	)
}
