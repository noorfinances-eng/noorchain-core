package types

// -----------------------------------------------------------------------------
//  Event types for the PoSS (noorsignal) module
// -----------------------------------------------------------------------------
//
// Ces constantes définissent les noms des events et des attributs utilisés
// dans MsgSubmitSignal, MsgValidateSignal, MsgAddCurator, MsgRemoveCurator,
// MsgSetConfig.
// -----------------------------------------------------------------------------

// Event types
const (
	EventTypeSignalSubmitted = "poss.signal_submitted"
	EventTypeSignalValidated = "poss.signal_validated"

	EventTypeCuratorAdded   = "poss.curator_added"
	EventTypeCuratorRemoved = "poss.curator_removed"

	EventTypeConfigUpdated = "poss.config_updated"
)

// Event attribute keys
const (
	AttrKeySignalID          = "signal_id"
	AttrKeyParticipant       = "participant"
	AttrKeyCurator           = "curator"
	AttrKeyWeight            = "weight"
	AttrKeyTimestamp         = "timestamp"
	AttrKeyMetadata          = "metadata"

	AttrKeyTotalReward       = "total_reward"
	AttrKeyRewardParticipant = "reward_participant"
	AttrKeyRewardCurator     = "reward_curator"

	// Curator admin
	AttrKeyLevel    = "level"
	AttrKeyAuthority = "authority"

	// Config admin
	AttrKeyBaseReward       = "base_reward"
	AttrKeyMaxSignalsPerDay = "max_signals_per_day"
	AttrKeyEraIndex         = "era_index"
	AttrKeyParticipantRatio = "participant_ratio"
	AttrKeyCuratorRatio     = "curator_ratio"
)
