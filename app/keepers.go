package app

import (
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
)

// AppKeepers définit la liste des "keepers" (gestionnaires de modules)
// que NOORCHAIN utilisera.
//
// PROVISOIREMENT :
// - Certains champs ont déjà leur type concret (modules de base).
// - Les autres restent en interface{} (placeholders) pour simplifier.
// - Plus tard, tous les champs seront remplacés par leurs types réels.
//
// L'objectif de ce fichier est de poser une structure claire de
// l'architecture des modules, en avançant étape par étape.
type AppKeepers struct {
	// --- Modules Cosmos de base ---

	// Gestion des comptes (adresses, comptes de base, etc.).
	AccountKeeper authkeeper.AccountKeeper

	// Gestion des soldes & transferts de tokens.
	BankKeeper bankkeeper.Keeper

	// Gestion du staking (validateurs, délégations, etc.).
	StakingKeeper stakingkeeper.Keeper

	// Module de "mint" standard (NOORCHAIN n'utilisera pas forcément
	// ce module pour la logique PoSS, mais il reste présent par structure).
	// TODO: remplacer interface{} par le type concret du MintKeeper.
	MintKeeper interface{}

	// Gestion du slashing des validateurs (sanctions).
	// TODO: remplacer interface{} par le type concret du SlashingKeeper.
	SlashingKeeper interface{}

	// Gouvernance on-chain (votes, propositions, etc.).
	GovKeeper govkeeper.Keeper

	// Gestion centralisée des paramètres de modules.
	ParamsKeeper paramskeeper.Keeper

	// Module de gestion des crises (haltes de la chaîne, checks invariants).
	// TODO: remplacer interface{} par le type concret du CrisisKeeper.
	CrisisKeeper interface{}

	// Gestion des mises à jour de la chaîne (upgrades).
	// TODO: remplacer interface{} par le type concret de l'UpgradeKeeper.
	UpgradeKeeper interface{}

	// --- Modules IBC / transfert (future) ---

	// Gestion du protocole IBC.
	IBCKeeper interface{}

	// Module de transfert de tokens via IBC.
	TransferKeeper interface{}

	// --- Modules Ethermint / EVM (future) ---

	// Module EVM principal (exécution de smart contracts Solidity).
	EvmKeeper interface{}

	// Module de gestion du "fee market" EVM (EIP-1559 style).
	FeeMarketKeeper interface{}

	// --- Module NOORCHAIN spécifique (PoSS) ---

	// Module "noorsignal" pour la logique Proof of Signal Social :
	// - réception & validation des signaux sociaux
	// - application des règles anti-abus
	// - mint NUR sous cap fixe + halving
	// - distribution 70% / 30%
	NoorSignalKeeper interface{}
}
