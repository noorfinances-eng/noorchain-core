package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd est une commande placeholder pour initialiser un répertoire
// de travail NOORCHAIN (home directory, fichiers de config, etc.).
// Pour l'instant, elle ne fait qu'afficher un message.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise un répertoire NOORCHAIN (placeholder)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("noord init: commande placeholder (configuration locale à implémenter plus tard).")
		return nil
	},
}

func init() {
	// Enregistrer la commande init auprès de la racine.
	rootCmd.AddCommand(initCmd)
}
