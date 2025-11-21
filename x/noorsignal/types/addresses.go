package types

// Ce fichier centralise les adresses "officielles" de NOORCHAIN.
// Il permet :
// - d'éviter de dupliquer les adresses dans plusieurs fichiers
// - de pouvoir remplacer facilement les adresses placeholders par les vraies
// - de garder une cohérence entre genesis, PoSS et BankKeeper

// -----------------------------------------------------------------------------
// PLACEHOLDERS (Testnet) — seront remplacées plus tard
// -----------------------------------------------------------------------------

const PlaceholderFoundation   = "noor1foundationxxxxxxxxxxxxxxxxxxxxx"
const PlaceholderDevWallet    = "noor1devxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
const PlaceholderStimulus     = "noor1stimulusxxxxxxxxxxxxxxxxxxxxxxx"
const PlaceholderPresale      = "noor1presalexxxxxxxxxxxxxxxxxxxxxxxx"
const PlaceholderPossReserve  = "noor1possreservexxxxxxxxxxxxxxxxxxxx"

// -----------------------------------------------------------------------------
// FUTURES ADRESSES RÉELLES (vont être renseignées par le fondateur)
// Pour l'instant : elles sont vides, on ne les active pas.
// -----------------------------------------------------------------------------

var AddressFoundation   = ""
var AddressDevWallet    = ""
var AddressStimulus     = ""
var AddressPresale      = ""
var AddressPossReserve  = ""

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

// GetOfficialAddress retourne l'adresse réelle si elle existe,
// sinon retourne le placeholder (testnet).
func GetOfficialAddress(real string, placeholder string) string {
	if real != "" {
		return real
	}
	return placeholder
}

// Accès simplifié pour tous les modules (BankKeeper, PoSS, Genesis, etc.).
func GetFoundationAddress() string  { return GetOfficialAddress(AddressFoundation, PlaceholderFoundation) }
func GetDevWalletAddress() string   { return GetOfficialAddress(AddressDevWallet, PlaceholderDevWallet) }
func GetStimulusAddress() string    { return GetOfficialAddress(AddressStimulus, PlaceholderStimulus) }
func GetPresaleAddress() string     { return GetOfficialAddress(AddressPresale, PlaceholderPresale) }
func GetPossReserveAddress() string { return GetOfficialAddress(AddressPossReserve, PlaceholderPossReserve) }
