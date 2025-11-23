package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// validateGenesisCmd est une commande placeholder pour valider un
// fichier genesis. Plus tard, elle utilisera la logique réelle de
// validation Cosmos SDK.
var validateGenesisCmd = &cobra.Command{
	Use:   "validate-genesis",
	Short: "Valide un fichier genesis NOORCHAIN (placeholder)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("noord validate-genesis: validation réelle à implémenter plus tard.")
		return nil
	},
}

func init() {
	// Enregistrer la commande validate-genesis auprès de la racine.
	rootCmd.AddCommand(validateGenesisCmd)
}
