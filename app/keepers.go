package app

// AppKeepers définit la liste des "keepers" (gestionnaires de modules)
// que NOORCHAIN utilisera.
//
// IMPORTANT :
// - À ce stade, tous les champs sont de type interface{} (placeholders).
// - Dans des phases techniques futures, chaque champ sera remplacé
//   par le type concret du keeper correspondant (ex: authkeeper.AccountKeeper).
//
// L'objectif de ce fichier est de poser une structure claire de
// l'architecture des modules, sans encore entrer dans les détails
// d'implémentation Cosmos SDK / Ethermint.
type AppKeepers struct {
	// --- Modules Cosmos de base ---

	// Gestion des comptes (adresses, comptes de base, etc.)
	AccountKeeper interface{}

	// Gestion des soldes & transferts de tokens.
	BankKeeper interface{}

	// Gestion du staking (validateurs, délégations, etc.).
	StakingKeeper interface{}

	// Module de "mint" standard (NOORCHAIN n'utilisera pas forcément
	// ce module pour la logique PoSS, mais il reste présent par structure).
	MintKeeper interface{}

	// Gestion du slashing des validateurs (sanctions).
	SlashingKeeper interface{}

	// Gouvernance on-chain (votes, propositions, etc.).
	GovKeeper interface{}

	// Gestion centralisée des paramètres de modules.
	ParamsKeeper interface{}

	// Module de gestion des crises (haltes de la chaîne, checks invariants).
	CrisisKeeper interface{}

	// Gestion des mises à jour de la chaîne (upgrades).
	UpgradeKeeper interface{}

	// --- Modules IBC / transfert (future) ---

	// Gestion du protocole IBC.
	IBCKeeper interface{}

	// Module de transfert de tokens via IBC.
	TransferKeeper interface{}

	// --- Modules Ethermint / EVM (future) ---

	// Module EVM principal (execution de smart contracts Solidity).
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
