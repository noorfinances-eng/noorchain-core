package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	errorsmod "cosmossdk.io/errors"
)

//
// -----------------------------------------------------------------------------
//  MsgSubmitSignal
// -----------------------------------------------------------------------------

type MsgSubmitSignal struct {
	Participant string `json:"participant" yaml:"participant"`
	Weight      uint32 `json:"weight" yaml:"weight"`
	Metadata    string `json:"metadata" yaml:"metadata"`
}

func (m MsgSubmitSignal) Route() string { return "noorsignal" }
func (m MsgSubmitSignal) Type() string  { return "submit_signal" }

func (m MsgSubmitSignal) ValidateBasic() error {
	if m.Participant == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "participant address cannot be empty")
	}
	if m.Weight == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "weight must be >= 1")
	}
	if m.Weight > 100 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "weight must be <= 100")
	}
	return nil
}

func (m MsgSubmitSignal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Participant)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgSubmitSignal) GetParticipantAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(m.Participant)
}

//
// -----------------------------------------------------------------------------
//  MsgValidateSignal
// -----------------------------------------------------------------------------

type MsgValidateSignal struct {
	Curator  string `json:"curator" yaml:"curator"`
	SignalId uint64 `json:"signal_id" yaml:"signal_id"`
}

func (m MsgValidateSignal) Route() string { return "noorsignal" }
func (m MsgValidateSignal) Type() string  { return "validate_signal" }

func (m MsgValidateSignal) ValidateBasic() error {
	if m.Curator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "curator address cannot be empty")
	}
	if m.SignalId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "signal_id must be >= 1")
	}
	return nil
}

func (m MsgValidateSignal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Curator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgValidateSignal) GetCuratorAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(m.Curator)
}

//
// -----------------------------------------------------------------------------
//  ADMIN MESSAGES
// -----------------------------------------------------------------------------
//  MsgAddCurator
//  MsgRemoveCurator
//  MsgSetConfig
// -----------------------------------------------------------------------------

// -----------------------------
// MsgAddCurator
// -----------------------------
type MsgAddCurator struct {
	Authority string `json:"authority" yaml:"authority"`
	Curator   string `json:"curator" yaml:"curator"`
	Level     string `json:"level" yaml:"level"`
}

func (m MsgAddCurator) Route() string { return "noorsignal" }
func (m MsgAddCurator) Type() string  { return "add_curator" }

func (m MsgAddCurator) ValidateBasic() error {
	if m.Authority == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "authority cannot be empty")
	}
	if m.Curator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "curator cannot be empty")
	}
	if m.Level == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "level cannot be empty")
	}
	return nil
}

func (m MsgAddCurator) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// -----------------------------
// MsgRemoveCurator
// -----------------------------
type MsgRemoveCurator struct {
	Authority string `json:"authority" yaml:"authority"`
	Curator   string `json:"curator" yaml:"curator"`
}

func (m MsgRemoveCurator) Route() string { return "noorsignal" }
func (m MsgRemoveCurator) Type() string  { return "remove_curator" }

func (m MsgRemoveCurator) ValidateBasic() error {
	if m.Authority == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "authority cannot be empty")
	}
	if m.Curator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "curator cannot be empty")
	}
	return nil
}

func (m MsgRemoveCurator) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// -----------------------------
// MsgSetConfig
// -----------------------------
type MsgSetConfig struct {
	Authority        string `json:"authority" yaml:"authority"`
	BaseReward       string `json:"base_reward" yaml:"base_reward"`
	MaxSignalsPerDay uint32 `json:"max_signals_per_day" yaml:"max_signals_per_day"`
	EraIndex         uint64 `json:"era_index" yaml:"era_index"`
	ParticipantRatio uint32 `json:"participant_ratio" yaml:"participant_ratio"`
	CuratorRatio     uint32 `json:"curator_ratio" yaml:"curator_ratio"`
}

func (m MsgSetConfig) Route() string { return "noorsignal" }
func (m MsgSetConfig) Type() string  { return "set_config" }

func (m MsgSetConfig) ValidateBasic() error {
	if m.Authority == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "authority cannot be empty")
	}

	if m.BaseReward == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "base_reward cannot be empty")
	}

	if (m.ParticipantRatio + m.CuratorRatio) != 100 {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("participant_ratio + curator_ratio must equal 100 (got %d)",
				m.ParticipantRatio+m.CuratorRatio),
		)
	}

	return nil
}

func (m MsgSetConfig) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
