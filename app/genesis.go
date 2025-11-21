package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// ApplyEconomicGenesis injecte les comptes économiques NOORCHAIN
// dans le genesis state (section bank).
//
// Les adresses sont centralisées dans x/noorsignal/types/addresses.go
// et peuvent être remplacées plus tard sans toucher au code.
//
// Cette fonction est appelée dans InitChainer,
// avant le ModuleManager.InitGenesis().
func ApplyEconomicGenesis(genesisState map[string]json.RawMessage, cdc codec.JSONCodec) {

	// 1) Récupérer la partie bank du genesis
	var bankState banktypes.GenesisState
	if genesisState[banktypes.ModuleName] != nil {
		if err := cdc.UnmarshalJSON(genesisState[banktypes.ModuleName], &bankState); err != nil {
			panic(err)
		}
	} else {
		bankState = banktypes.GenesisState{}
	}

	// 2) Récupérer les adresses officielles (réelles ou placeholders)
	foundationAddr := noorsignaltypes.GetFoundationAddress()
	devAddr        := noorsignaltypes.GetDevWalletAddress()
	stimulusAddr   := noorsignaltypes.GetStimulusAddress()
	presaleAddr    := noorsignaltypes.GetPresaleAddress()
	possReserveAddr := noorsignaltypes.GetPossReserveAddress()

	// 3) Définir les montants des comptes (en unur)
	//    Basé sur la répartition 5 / 5 / 5 / 5 / 80 %
	//    1 NUR = 1'000'000 unur

	const (
		totalSupply       uint64 = 299_792_458
		allocation5Pct    uint64 = totalSupply * 5 / 100     // 14 989 622
		allocation80Pct   uint64 = totalSupply * 80 / 100    // 239 833 966
		micro             uint64 = 1_000_000                 // unur per NUR
	)

	amount5pct  := sdk.NewInt(int64(allocation5Pct * micro))
	amount80pct := sdk.NewInt(int64(allocation80Pct * micro))

	// 4) Ajouter les balances au genesis
	addBalance := func(addr string, amount sdk.Int) {
		bankState.Balances = append(bankState.Balances, banktypes.Balance{
			Address: addr,
			Coins:   sdk.NewCoins(sdk.NewCoin("unur", amount)),
		})
	}

	addBalance(foundationAddr, amount5pct)
	addBalance(devAddr, amount5pct)
	addBalance(stimulusAddr, amount5pct)
	addBalance(presaleAddr, amount5pct)
	addBalance(possReserveAddr, amount80pct)

	// 5) Mettre à jour supply totale
	totalCoins := bankState.Supply.Add(
		sdk.NewCoin("unur", amount5pct),
		sdk.NewCoin("unur", amount5pct),
		sdk.NewCoin("unur", amount5pct),
		sdk.NewCoin("unur", amount5pct),
		sdk.NewCoin("unur", amount80pct),
	)

	bankState.Supply = totalCoins

	// 6) Ré-encoder la section bank du genesis
	bz, err := cdc.MarshalJSON(&bankState)
	if err != nil {
		panic(err)
	}
	genesisState[banktypes.ModuleName] = bz
}
