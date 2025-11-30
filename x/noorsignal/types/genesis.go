package types

import "fmt"

// GenesisState defines the PoSS module genesis state.
// This is the minimal but real structure for NOORCHAIN PoSS.
type GenesisState struct {
	// Total number of PoSS signals already processed on the chain.
	// Starts at 0 on a fresh network and increases over time.
	TotalSignals uint64 `json:"total_signals" yaml:"total_signals"`

	// Total NUR minted via PoSS, in the smallest unit (unur).
	// Stored as a string to avoid precision issues and to stay simple for now.
	TotalMinted string `json:"total_minted" yaml:"total_minted"`

	// Maximum number of PoSS signals allowed per address per day.
	// This is the first anti-abuse guardrail.
	MaxSignalsPerDay uint32 `json:"max_signals_per_day" yaml:"max_signals_per_day"`

	// Reward split (must always sum to 100).
	// Official NOORCHAIN rule:
	//   - 70 % for the participant
	//   - 30 % for the curator
	ParticipantShare uint32 `json:"participant_share" yaml:"participant_share"`
	CuratorShare     uint32 `json:"curator_share" yaml:"curator_share"`
}

// DefaultGenesis returns the default PoSS genesis state.
// These values are aligned with the official NOORCHAIN rules.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TotalSignals:     0,
		TotalMinted:      "0",
		MaxSignalsPerDay: 50, // simple, conservative daily limit per address
		ParticipantShare: 70, // 70 % participant
		CuratorShare:     30, // 30 % curator
	}
}

// ValidateGenesis performs basic validation of the PoSS genesis state.
// It ensures that the reward split is coherent and that mandatory fields are set.
func ValidateGenesis(gs *GenesisState) error {
	if gs == nil {
		return fmt.Errorf("genesis state cannot be nil")
	}

	// Reward split must always be 100 %.
	if gs.ParticipantShare+gs.CuratorShare != 100 {
		return fmt.Errorf("participant_share + curator_share must equal 100 (got %d + %d)",
			gs.ParticipantShare, gs.CuratorShare)
	}

	// TotalMinted must not be empty (at least "0").
	if gs.TotalMinted == "" {
		return fmt.Errorf("total_minted cannot be empty (use \"0\" for a fresh chain)")
	}

	// MaxSignalsPerDay must be > 0 to avoid a "locked" PoSS system.
	if gs.MaxSignalsPerDay == 0 {
		return fmt.Errorf("max_signals_per_day must be greater than 0")
	}

	return nil
}
