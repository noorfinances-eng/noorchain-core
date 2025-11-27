package app

// Métadonnées globales de NOORCHAIN.
// Ce fichier ne contient AUCUNE dépendance externe.
// Il est donc totalement safe pour la Phase 2.

// Nom de l'application (binaire CLI).
const AppName = "noord"

// Identifiant de chaîne par défaut pour le DEV / Testnet local.
// On pourra en ajouter d'autres plus tard (testnet public, mainnet, etc.).
const DefaultChainID = "noorchain-dev-1"

// Dossier par défaut du node (dans le HOME de l'utilisateur).
// Exemple sur Linux/Mac:
//   ~/.noord
const DefaultNodeHome = ".noord"

// Version simple pour l’instant (placeholder).
// Plus tard, on pourra la brancher sur un tag Git ou un script de build.
const AppVersion = "0.0.1-dev"
