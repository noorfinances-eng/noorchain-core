package types

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgCreateSignal est le message principal PoSS côté Go.
//
// Pour l’instant, on le définit entièrement en Go (sans s’appuyer sur un
// fichier .pb.go généré). Pour satisfaire sdk.Msg (qui impose proto.Message),
// on ajoute aussi les méthodes Reset / String / ProtoMessage en bas du fichier.
type MsgCreateSignal struct {
	// Participant NOOR (noor1...) — SIGNER principal de la tx.
	Participant string `json:"participant" yaml:"participant"`

	// Curator NOOR (noor1...) — reçoit 30 % du reward PoSS (quand activé).
	Curator string `json:"curator" yaml:"curator"`

	// Type de signal PoSS (micro_donation, participation, content, ccn).
	SignalType SignalType `json:"signal_type" yaml:"signal_type"`

	// Métadonnées libres (URI, hash, description…).
	Metadata string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Timestamp logique fourni par le client (optionnel).
	// La chaîne utilisera quand même son block time pour l’accounting interne.
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`

	// Date logique PoSS "YYYY-MM-DD" utilisée pour les compteurs journaliers.
	Date string `json:"date" yaml:"date"`
}

// NewMsgCreateSignal est un helper pour construire le message côté client.
func NewMsgCreateSignal(
	participant string,
	curator string,
	signalType SignalType,
	metadata string,
	timestamp time.Time,
	date string,
) *MsgCreateSignal {
	return &MsgCreateSignal{
		Participant: participant,
		Curator:     curator,
		SignalType:  signalType,
		Metadata:    metadata,
		Timestamp:   timestamp,
		Date:        date,
	}
}

// MsgCreateSignalResponse : réponse logique renvoyée par le module PoSS.
// À ce stade, il s’agit des REWARDS THÉORIQUES (aucun mint réel encore).
type MsgCreateSignalResponse struct {
	ParticipantReward sdk.Coin `json:"participant_reward" yaml:"participant_reward"`
	CuratorReward     sdk.Coin `json:"curator_reward" yaml:"curator_reward"`
}

// On s’assure que MsgCreateSignal implémente bien sdk.Msg.
var _ sdk.Msg = &MsgCreateSignal{}

// Route retourne le nom du module PoSS.
func (m *MsgCreateSignal) Route() string {
	return ModuleName
}

// Type retourne le type logique du message.
func (m *MsgCreateSignal) Type() string {
	return "CreateSignal"
}

// ValidateBasic fait les checks de base (sans accès au store).
func (m *MsgCreateSignal) ValidateBasic() error {
	if m.Participant == "" {
		return fmt.Errorf("participant address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Participant); err != nil {
		return fmt.Errorf("invalid participant bech32 address: %w", err)
	}

	if m.Curator == "" {
		return fmt.Errorf("curator address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return fmt.Errorf("invalid curator bech32 address: %w", err)
	}

	switch m.SignalType {
	case SignalTypeMicroDonation,
		SignalTypeParticipation,
		SignalTypeContent,
		SignalTypeCCN:
		// ok
	default:
		return fmt.Errorf("invalid PoSS signal type: %s", m.SignalType)
	}

	if m.Date == "" {
		return fmt.Errorf("date cannot be empty (expected YYYY-MM-DD)")
	}
	// Petit check format rapide (on reste léger ici).
	if len(m.Date) != 10 || m.Date[4] != '-' || m.Date[7] != '-' {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD, got: %s", m.Date)
	}

	return nil
}

// GetSignBytes retourne les bytes à signer (JSON trié).
func (m *MsgCreateSignal) GetSignBytes() []byte {
	bz, err := json.Marshal(m)
	if err != nil {
		// En cas de problème, on panique — c’est un bug dev, pas une erreur user.
		panic(fmt.Errorf("cannot marshal MsgCreateSignal to JSON: %w", err))
	}
	return sdk.MustSortJSON(bz)
}

// GetSigners retourne la liste des signers (ici uniquement le participant).
func (m *MsgCreateSignal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Participant)
	if err != nil {
		// Ne devrait pas arriver si ValidateBasic a été appelé avant.
		panic(fmt.Errorf("invalid participant address in GetSigners: %w", err))
	}
	return []sdk.AccAddress{addr}
}

// ------------------------------------------------------------
// Implémentation minimale de proto.Message
// (requis car sdk.Msg embed proto.Message).
// ------------------------------------------------------------

// Reset remet le message à zéro.
func (m *MsgCreateSignal) Reset() { *m = MsgCreateSignal{} }

// String retourne une représentation texte (JSON best-effort).
func (m *MsgCreateSignal) String() string {
	bz, err := json.Marshal(m)
	if err != nil {
		return "MsgCreateSignal{}"
	}
	return string(bz)
}

// ProtoMessage est une méthode vide exigée par l’interface proto.Message.
func (*MsgCreateSignal) ProtoMessage() {}
