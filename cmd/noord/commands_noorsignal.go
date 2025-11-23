package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// noorsignalCmd est la commande parente pour toutes les requêtes
// liées au module PoSS (x/noorsignal).
//
// Exemple d'utilisation cible (plus tard, quand le gRPC sera branché):
//   noord query noorsignal config
var noorsignalCmd = &cobra.Command{
	Use:   "noorsignal",
	Short: "Requêtes liées au module PoSS (noorsignal)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Utilise une sous-commande, par ex.:")
		fmt.Println("  noord query noorsignal config")
		return nil
	},
}

// noorsignalConfigCmd sera plus tard branchée sur une requête gRPC
// vers le QueryServer du module PoSS. Pour l'instant, elle affiche
// simplement un message indiquant que la fonctionnalité est à venir.
var noorsignalConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Affiche la configuration PoSS (placeholder)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("noord query noorsignal config: requête non implémentée (placeholder).")
		fmt.Println("Plus tard, cette commande interrogera le QueryServer gRPC du module PoSS.")
		return nil
	},
}

func init() {
	// Rattacher "noorsignal" à "query", puis "config" à "noorsignal".
	queryCmd.AddCommand(noorsignalCmd)
	noorsignalCmd.AddCommand(noorsignalConfigCmd)
}
