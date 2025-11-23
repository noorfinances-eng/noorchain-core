package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TypeMsgUpdateCounter est le type de message principal du module V1.
// C'est un exemple simple : un message qui met à jour un compteur global.
const TypeMsgUpdateCounter = "update_counter"

// validateCreatorAddress vérifie qu'une adresse est au bon format Bech32.
func validateCreatorAddress(addr string) error {
	if addr == "" {
		return fmt.Errorf("creator address cannot be empty")
	}
	_, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return fmt.Errorf("invalid bech32 address: %w", err)
	}
	return nil
}

// MsgUpdateCounter représente un message simple qui met à jour le compteur.
// Dans un futur PoSS complet, ceci sera remplacé par des messages
// liés aux signaux, curateurs, etc.
type MsgUpdateCounter struct {
	Creator string `json:"creator" yaml:"creator"`
	Value   uint64 `json:"value" yaml:"value"`
}

// NewMsgUpdateCounter crée une nouvelle instance de MsgUpdateCounter.
func NewMsgUpdateCounter(creator string, value uint64) *MsgUpdateCounter {
	return &MsgUpdateCounter{
		Creator: creator,
		Value:   value,
	}
}

// Route renvoie la route du module (RouterKey défini dans keys.go).
func (m *MsgUpdateCounter) Route() string {
	return RouterKey
}

// Type renvoie le type du message.
func (m *MsgUpdateCounter) Type() string {
	return TypeMsgUpdateCounter
}

// GetSigners renvoie la liste des signataires requis pour ce message.
func (m *MsgUpdateCounter) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		// ValidateBasic est censé avoir déjà validé l'adresse.
		// Si on arrive ici, on panique pour ne pas signer avec une adresse invalide.
		panic(fmt.Sprintf("invalid creator address in GetSigners: %v", err))
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes renvoie les bytes à signer.
// Ici on utilise simplement un JSON canonique sans ModuleCdc.
func (m *MsgUpdateCounter) GetSignBytes() []byte {
	bz, err := json.Marshal(m)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal MsgUpdateCounter: %v", err))
	}
	return sdk.MustSortJSON(bz)
}

// ValidateBasic effectue les vérifications de base sur le message.
func (m *MsgUpdateCounter) ValidateBasic() error {
	if err := validateCreatorAddress(m.Creator); err != nil {
		return fmt.Errorf("invalid creator: %w", err)
	}
	// Exemple de validation simple : on exige une valeur > 0.
	if m.Value == 0 {
		return fmt.Errorf("value must be > 0")
	}
	return nil
}
