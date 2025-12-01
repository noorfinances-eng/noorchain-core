package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TypeMsgCreateSignal is the type string used for PoSS "create signal" messages.
const TypeMsgCreateSignal = "create_signal"

// MsgCreateSignal is the core PoSS message used by participants to submit a
// social signal that can later be validated and rewarded.
//
// NOTE: à ce stade, c'est seulement la définition du message :
// - pas encore de handler,
// - pas encore de vérification de limites journalières,
// - pas encore de débit / crédit de NUR.
type MsgCreateSignal struct {
	// Participant is the NOOR bech32 address (noor1...) of the person
	// emitting the signal.
	Participant string `json:"participant" yaml:"participant"`

	// Curator is the NOOR bech32 address (noor1...) of the curator who will
	// validate the signal (Bronze / Silver / Gold curator).
	Curator string `json:"curator" yaml:"curator"`

	// SignalType is the type of PoSS signal (micro-donation, participation,
	// content, CCN).
	//
	// On supposera plus tard qu’il s’agit d’un enum (SignalType) avec
	// des valeurs bien définies, mais on ne force pas encore IsValid()
	// pour ne pas casser la build.
	SignalType SignalType `json:"signal_type" yaml:"signal_type"`

	// Payload is a free-form JSON or string payload describing the signal.
	// Exemple :
	// - un ID de reçu de micro-don,
	// - un ID d’événement,
	// - un hash de contenu certifié,
	// - une référence CCN, etc.
	Payload string `json:"payload" yaml:"payload"`
}

// NewMsgCreateSignal is a small helper to build a MsgCreateSignal.
func NewMsgCreateSignal(
	participant string,
	curator string,
	signalType SignalType,
	payload string,
) *MsgCreateSignal {
	return &MsgCreateSignal{
		Participant: participant,
		Curator:     curator,
		SignalType:  signalType,
		Payload:     payload,
	}
}

// Route implements sdk.Msg and returns the PoSS router key.
func (m *MsgCreateSignal) Route() string { return RouterKey }

// Type implements sdk.Msg and returns the message type.
func (m *MsgCreateSignal) Type() string { return TypeMsgCreateSignal }

// ValidateBasic performs stateless validation on the message fields.
// Ici on fait seulement des checks simples (pas encore les limites PoSS).
func (m *MsgCreateSignal) ValidateBasic() error {
	if m == nil {
		return fmt.Errorf("MsgCreateSignal cannot be nil")
	}

	if m.Participant == "" {
		return fmt.Errorf("participant address cannot be empty")
	}
	if m.Curator == "" {
		return fmt.Errorf("curator address cannot be empty")
	}
	if m.Participant == m.Curator {
		return fmt.Errorf("participant and curator must be different addresses")
	}

	// Validate bech32 addresses (noor1...)
	if _, err := sdk.AccAddressFromBech32(m.Participant); err != nil {
		return fmt.Errorf("invalid participant address: %w", err)
	}
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return fmt.Errorf("invalid curator address: %w", err)
	}

	// Simple sanity check: on s’assure juste qu’un type a été fourni.
	// Le vrai filtrage (enum STRICT) viendra quand on aura la définition
	// complète de SignalType et éventuellement une méthode IsValid().
	if fmt.Sprintf("%v", m.SignalType) == "" {
		return fmt.Errorf("signal_type cannot be empty")
	}

	// Payload can be empty in v1, but on vérifie au moins la longueur max simple
	if len(m.Payload) > 10_000 {
		return fmt.Errorf("payload too long")
	}

	return nil
}

// GetSignBytes returns the bytes to sign for this message.
// On reste volontairement simple : JSON + tri des clés.
func (m *MsgCreateSignal) GetSignBytes() []byte {
	bz, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the list of signers for this message.
// Pour NOORCHAIN : seul le participant signe le message.
func (m *MsgCreateSignal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Participant)
	if err != nil {
		// On panic ici parce que ValidateBasic aurait dû catcher l’erreur avant.
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
