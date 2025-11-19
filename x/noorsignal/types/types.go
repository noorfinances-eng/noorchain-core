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
// - l'activation / désactivation du module
// - l'indice d'ère pour le halving (tous les 8 ans).
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

	// EraIndex représente l'"ère" PoSS actuelle pour le halving :
	// - 0 : première période de 8 ans (aucun halving, facteur 1)
	// - 1 : deuxième période (premier halving, facteur 2)
	// - 2 : troisième période (deuxième halving, facteur 4)
	// etc.
	//
	// La gestion de l'évolution de ce champ (tous les 8 ans) pourra être
	// faite via gouvernance, paramètres on-chain ou logique externe.
	EraIndex uint32
}

// DefaultPossConfig retourne une configuration PoSS par défaut
// cohérente avec le modèle NOORCHAIN :
// - 70% pour le participant
// - 30% pour le curator
// - module activé
// - baseReward et limite journalière placés à des valeurs symboliques
//   (ajustables plus tard dans le genesis ou par gouvernance).
func DefaultPossConfig() PossConfig {
	return PossConfig{
		// Exemple symbolique : 100 unités de base par signal 1x.
		// L'unité réelle (NUR vs unur) sera clarifiée plus tard.
		BaseReward: 100,

		ParticipantShare: 70,
		CuratorShare:     30,

		// Exemple : 50 signaux max / jour / participant.
		MaxSignalsPerDay: 50,

		Enabled: true,

		// Au lancement de NOORCHAIN (année 0 à 8),
		// aucun halving n'a encore eu lieu.
		EraIndex: 0,
	}
}
