package main

import (
	"os"
)

// main est le point d'entrée de l'exécutable "noord".
// Il délègue simplement l'exécution au rootCmd défini dans root.go.
func main() {
	if err := Execute(); err != nil {
		// Pour l'instant on se contente d'un code de retour non nul.
		// Plus tard on pourra logger l'erreur.
		os.Exit(1)
	}
}
