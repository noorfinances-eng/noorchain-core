package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgSubmitSignal représente une transaction par laquelle un participant
// soumet un "signal social" PoSS.
//
// Exemples de signaux :
// - participation vérifiée à un événement
// - micro-don
// - contenu certifié NOORCHAIN
//
// Remarque :
// - La version finale sera générée à partir de fichiers .proto.
// - Cette struct Go sert de point de départ conceptuel.
type MsgSubmitSignal struct {
	// Adresse du participant qui soumet le signal.
	Participant string `json:"participant" yaml:"participant"`

	// Poids du signal (barème PoSS converti en entier : 1, 2, 5, etc.).
	Weight uint32 `json:"weight" yaml:"weight"`

	// Informations supplémentaires (hash de contenu, ID externe, etc.).
	Metadata string `json:"metadata" yaml:"metadata"`
}

// GetParticipantAddress retourne l'adresse du participant sous forme sdk.AccAddress.
// Helper simple pour éviter de répéter la conversion partout.
func (m MsgSubmitSignal) GetParticipantAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(m.Participant)
}

// MsgValidateSignal représente une transaction par laquelle un curator
// valide un signal existant.
//
// Le flux standard pourrait être :
// 1) un participant soumet un signal (MsgSubmitSignal)
// 2) un curator le valide (MsgValidateSignal) → ce qui débloque les récompenses PoSS.
//
// Remarque : comme pour MsgSubmitSignal, la version finale sera définie
// en Protobuf, cette struct est une base conceptuelle.
type MsgValidateSignal struct {
	// Adresse du curator qui valide le signal.
	Curator string `json:"curator" yaml:"curator"`

	// Identifiant du signal à valider.
	SignalId uint64 `json:"signal_id" yaml:"signal_id"`
}

// GetCuratorAddress retourne l'adresse du curator sous forme sdk.AccAddress.
func (m MsgValidateSignal) GetCuratorAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(m.Curator)
}

// MsgUpdatePossConfig représenterait, dans une version future, une
// mise à jour de la configuration PoSS via gouvernance ou un rôle
// administratif strictement contrôlé.
//
// Pour l'instant, cette struct est laissée en commentaire car la
// gouvernance exacte (on-chain / off-chain) n'est pas finalisée.
/*
type MsgUpdatePossConfig struct {
	// Adresse de l'initiateur de la modification (ex: adresse gov module).
	Authority string `json:"authority" yaml:"authority"`

	// Nouvelle configuration PoSS proposée.
	NewConfig PossConfig `json:"new_config" yaml:"new_config"`
}
*/
