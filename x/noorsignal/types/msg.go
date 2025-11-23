package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TypeMsgUpdateCounter est le type de message principal du module V1.
// Il sert d'exemple simple : un message qui met à jour un compteur global.
const TypeMsgUpdateCounter = "update_counter"

// Vérifie qu'une adresse est au bon format Bech32.
func validateCreatorAddress(addr string) error {
	_, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return fmt.Errorf("invalid creator address: %w", err)
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
		// En théorie, ValidateBasic doit déjà avoir catch l'erreur,
		// donc on renvoie un slice vide ici en cas de problème.
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes renvoie les bytes à signer (JSON canonique).
func (m *MsgUpdateCounter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic effectue les vérifications de base sur le message.
func (m *MsgUpdateCounter) ValidateBasic() error {
	if err := validateCreatorAddress(m.Creator); err != nil {
		return sdkerrors.Wrap(err, "invalid creator")
	}
	// Exemple de validation simple : on exige une valeur > 0.
	if m.Value == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "value must be > 0")
	}
	return nil
}
