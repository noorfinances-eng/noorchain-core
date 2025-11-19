package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Signal représente un "signal social" PoSS émis par un participant.
//
// Il peut s'agir par exemple :
// - d'une participation vérifiée
// - d'un micro-don
// - d'un contenu certifié NOOR
//
// Remarque :
// - La version finale sera définie en Protobuf (.proto).
// - Cette struct Go sert de point de départ conceptuel.
type Signal struct {
	// Identifiant unique du signal (clé simple auto-incrémentée ou dérivée).
	Id uint64

	// Adresse du participant qui a émis le signal.
	Participant sdk.AccAddress

	// Adresse du curator qui a validé le signal (peut être vide avant validation).
	Curator sdk.AccAddress

	// Poids du signal (barème PoSS : 0.5x, 1x, 2x, etc. → converti ici en entier).
	// Exemple : 1 = 1x, 2 = 2x, 5 = 5x.
	Weight uint32

	// Timestamp approximatif de l'émission / validation du signal.
	Time time.Time

	// Informations additionnelles (hash de contenu, ID externe, etc.).
	Metadata string
}

// Curator représente un "validateur social" dans le système PoSS.
//
// C'est un acteur NOORCHAIN particulier (association, école, ONG, etc.)
// qui valide des signaux et reçoit une part des récompenses (ex: 30%).
type Curator struct {
	// Adresse on-chain du curator.
	Address sdk.AccAddress

	// Niveau / rôle du curator (ex: "bronze", "silver", "gold").
	// La logique précise sera définie plus tard.
	Level string

	// Indicateurs simples pour le suivi (non strictement nécessaires
	// au consensus, mais utiles pour la gouvernance / stats).
	TotalSignalsValidated uint64
	Active                bool
}

// PossConfig représente une configuration globale PoSS.
//
// Elle regroupe les paramètres qui contrôlent :
// - les récompenses de base
// - la répartition 70% / 30%
// - les limites quotidiennes
// - le comportement de halving (géré côté supply globale)
type PossConfig struct {
	// Montant de base en NUR (ou en "unur") attribué à un signal de poids 1x.
	// Pour l'instant, on garde un entier simple; plus tard, ce sera
	// probablement un sdk.Int ou une DecCoin.
	BaseReward uint64

	// Part du participant (en pourcentage entier, ex: 70).
	ParticipantShare uint32

	// Part du curator (en pourcentage entier, ex: 30).
	CuratorShare uint32

	// Limite max de signaux comptabilisés par jour et par participant.
	MaxSignalsPerDay uint32

	// Indicateur si le module PoSS est actif.
	Enabled bool
}
