package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Adresses Bech32 officielles NOORCHAIN pour le genesis 5/5/5/5/80.
//
// Chaque adresse correspond à un wallet distinct, avec une seed
// différente, comme défini dans le plan NOORCHAIN 1.0.
//
// IMPORTANT :
// - Ne pas modifier ces valeurs après le lancement du mainnet.
// - Ces adresses seront utilisées pour :
//   * le genesis mainnet
//   * la documentation publique
//   * la configuration PoSS / BankKeeper.
const (
	FoundationAddressBech32  = "noor1dwzpnw9g9p2cucj2w2dnxk2at4amaqsrvmyka0"
	DevAddressBech32         = "noor1s2gzjrec9elucycpj66d2eyyw09tjuqfcasy7k"
	StimulusAddressBech32    = "noor1qt2h4crdtngyw4yn3yy8sqraqpzh4ghx0c4l46"
	PresaleAddressBech32     = "noor1pxt0wlq8xswj0l6jm2spzcspjwjnnvaggatalz"
	PossReserveAddressBech32 = "noor1c4gg5n37ycnfvaalsaa4nl22pfzp9ujqwr0mne"
)

// mustAccAddressFromBech32 convertit une string Bech32 en sdk.AccAddress
// et panic si l'adresse est invalide. Cela doit uniquement échouer si
// le code contient une erreur, pas en production.
func mustAccAddressFromBech32(addr string) sdk.AccAddress {
	acc, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		panic(err)
	}
	return acc
}

// Adresses AccAddress prêtes à l'emploi pour les modules qui en ont besoin
// (par exemple pour la réserve PoSS dans BankKeeper).
var (
	FoundationAccAddress  = mustAccAddressFromBech32(FoundationAddressBech32)
	DevAccAddress         = mustAccAddressFromBech32(DevAddressBech32)
	StimulusAccAddress    = mustAccAddressFromBech32(StimulusAddressBech32)
	PresaleAccAddress     = mustAccAddressFromBech32(PresaleAddressBech32)
	PossReserveAccAddress = mustAccAddressFromBech32(PossReserveAddressBech32)
)
