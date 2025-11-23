package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd est la commande racine du binaire "noord".
//
// Pour l'instant, il s'agit d'un skeleton simple permettant d'ajouter
// des sous-commandes (init, start, validate-genesis, query, etc.).
var rootCmd = &cobra.Command{
	Use:   "noord",
	Short: "NOORCHAIN node / CLI (prototype)",
	Long: `NOORCHAIN - prototype CLI

Ce binaire sert à :
- démarrer un nœud NOORCHAIN (plus tard),
- valider un fichier genesis,
- exécuter des requêtes (query),
- tester le module PoSS (noorsignal).

Actuellement, seules quelques commandes de base sont implémentées.
`,
}

// Execute lance l'exécution de rootCmd.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return err
	}
	return nil
}
