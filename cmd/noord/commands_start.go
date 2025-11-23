package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd est une commande placeholder pour démarrer un nœud NOORCHAIN.
// Elle sera remplacée plus tard par une intégration réelle avec l'App.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Démarre un nœud NOORCHAIN (placeholder)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("noord start: démarrage du nœud non implémenté dans cette version (prototype CLI).")
		return nil
	},
}

func init() {
	// Enregistrer la commande start auprès de la racine.
	rootCmd.AddCommand(startCmd)
}
