package types

import (
	"os"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Variables globales pour les adresses de test (bech32 valides).
var (
	testParticipantAddr string
	testCuratorAddr     string
)

// initBech32Config configure le préfixe bech32 pour les tests.
//
// On utilise le préfixe "cosmos" standard du SDK, ce qui est suffisant
// pour les tests unitaires (en prod, NOORCHAIN utilisera "noor").
func initBech32Config() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("cosmos", "cosmospub")
	cfg.SetBech32PrefixForValidator("cosmosvaloper", "cosmosvaloperpub")
	cfg.SetBech32PrefixForConsensusNode("cosmosvalcons", "cosmosvalconspub")
	cfg.Seal()
}

// newTestAddr construit une adresse bech32 valide à partir d'un byte seed.
//
// On crée simplement un tableau de 20 bytes avec la même valeur
// et on le convertit en AccAddress, puis en string bech32.
func newTestAddr(seed byte) string {
	bz := make([]byte, 20)
	for i := 0; i < 20; i++ {
		bz[i] = seed
	}
	return sdk.AccAddress(bz).String()
}

// TestMain est le point d'entrée global des tests de ce package.
// On y configure le bech32 et on génère les adresses.
func TestMain(m *testing.M) {
	initBech32Config()

	// Deux adresses de test stables mais valides.
	testParticipantAddr = newTestAddr(1)
	testCuratorAddr = newTestAddr(2)

	os.Exit(m.Run())
}

// helper pour construire un MsgCreateSignal de base valide.
func newValidMsg() *MsgCreateSignal {
	return NewMsgCreateSignal(
		testParticipantAddr,
		testCuratorAddr,
		SignalTypeMicroDonation,
		"test metadata",
		time.Now().UTC(),
		"2025-01-01",
	)
}

// Test que ValidateBasic passe bien avec un message complet et valide.
func TestValidateBasic_Valid(t *testing.T) {
	msg := newValidMsg()

	if err := msg.ValidateBasic(); err != nil {
		t.Fatalf("expected ValidateBasic to pass, got error: %v", err)
	}
}

// Participant vide → erreur.
func TestValidateBasic_InvalidParticipant(t *testing.T) {
	msg := newValidMsg()
	msg.Participant = ""

	if err := msg.ValidateBasic(); err == nil {
		t.Fatalf("expected error for empty participant, got nil")
	}
}

// Curator vide → erreur.
func TestValidateBasic_InvalidCurator(t *testing.T) {
	msg := newValidMsg()
	msg.Curator = ""

	if err := msg.ValidateBasic(); err == nil {
		t.Fatalf("expected error for empty curator, got nil")
	}
}

// Type de signal invalide → erreur.
func TestValidateBasic_InvalidSignalType(t *testing.T) {
	msg := newValidMsg()
	msg.SignalType = SignalType("invalid_type")

	if err := msg.ValidateBasic(); err == nil {
		t.Fatalf("expected error for invalid signal type, got nil")
	}
}

// Date vide → erreur.
func TestValidateBasic_EmptyDate(t *testing.T) {
	msg := newValidMsg()
	msg.Date = ""

	if err := msg.ValidateBasic(); err == nil {
		t.Fatalf("expected error for empty date, got nil")
	}
}

// Date mal formatée → erreur.
func TestValidateBasic_BadDateFormat(t *testing.T) {
	msg := newValidMsg()
	msg.Date = "2025/01/01" // mauvais format

	if err := msg.ValidateBasic(); err == nil {
		t.Fatalf("expected error for bad date format, got nil")
	}
}

// GetSigners doit retourner uniquement le participant (et ne pas paniquer).
func TestGetSigners_UsesParticipant(t *testing.T) {
	msg := newValidMsg()

	signers := msg.GetSigners()
	if len(signers) != 1 {
		t.Fatalf("expected 1 signer, got %d", len(signers))
	}
	// Si on arrive ici sans panic, le test est déjà intéressant.
}
