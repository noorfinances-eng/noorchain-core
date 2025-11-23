package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// queryCmd est la commande parente pour les requêtes de lecture
// (similaire à "noord query ..." dans les chaînes Cosmos classiques).
//
// Les sous-commandes spécifiques (comme "noord query noorsignal ...")
// seront ajoutées en dessous.
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Commandes de requête (lecture) sur NOORCHAIN",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Utilise une sous-commande, par ex.:")
		fmt.Println("  noord query noorsignal config")
		return nil
	},
}

func init() {
	// Enregistrer la commande query auprès de la racine.
	rootCmd.AddCommand(queryCmd)
}
