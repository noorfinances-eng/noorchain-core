package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// ApplyEconomicGenesis injecte les comptes économiques NOORCHAIN
// dans le genesis state (section bank).
//
// Les adresses sont centralisées dans x/noorsignal/types/addresses.go
// et peuvent être remplacées plus tard sans toucher au code.
func ApplyEconomicGenesis(
	genesisState map[string]json.RawMessage,
	cdc codec.JSONCodec,
) {

	// 1) Récupérer la partie bank du genesis
	var bankState banktypes.GenesisState
	if bz, ok := genesisState[banktypes.ModuleName]; ok && bz != nil {
		if err := cdc.UnmarshalJSON(bz, &bankState); err != nil {
			panic(err)
		}
	} else {
		bankState = banktypes.GenesisState{}
	}

	// 2) Récupérer les adresses officielles (réelles ou placeholders)
	foundationAddr := noorsignaltypes.GetFoundationAddress()
	devAddr := noorsignaltypes.GetDevWalletAddress()
	stimulusAddr := noorsignaltypes.GetStimulusAddress()
	presaleAddr := noorsignaltypes.GetPresaleAddress()
	possReserveAddr := noorsignaltypes.GetPossReserveAddress()

	// 3) Définir les montants (en unur)
	const (
		totalSupply     uint64 = 299_792_458
		allocation5Pct  uint64 = totalSupply * 5 / 100   // 14 989 622
		allocation80Pct uint64 = totalSupply * 80 / 100  // 239 833 966
		micro           uint64 = 1_000_000               // 1 NUR = 1e6 unur
	)

	amount5 := sdk.NewIntFromUint64(allocation5Pct * micro)
	amount80 := sdk.NewIntFromUint64(allocation80Pct * micro)

	// 4) Helper pour ajouter une balance + mettre à jour la supply
	addBalance := func(addr string, amt sdk.Int) {
		if addr == "" {
			// Si jamais une adresse réelle n'est pas encore renseignée,
			// on ne crée pas d'entrée vide dans le genesis.
			return
		}
		bankState.Balances = append(bankState.Balances, banktypes.Balance{
			Address: addr,
			Coins:   sdk.NewCoins(sdk.NewCoin("unur", amt)),
		})
		bankState.Supply = bankState.Supply.Add(sdk.NewCoin("unur", amt))
	}

	// 5) Appliquer les 5 poches économiques
	addBalance(foundationAddr, amount5)
	addBalance(devAddr, amount5)
	addBalance(stimulusAddr, amount5)
	addBalance(presaleAddr, amount5)
	addBalance(possReserveAddr, amount80)

	// 6) Ré-encoder la section bank du genesis
	bz, err := cdc.MarshalJSON(&bankState)
	if err != nil {
		panic(err)
	}
	genesisState[banktypes.ModuleName] = bz
}
