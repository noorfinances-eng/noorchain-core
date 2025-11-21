package main

import (
    "os"

    "github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/cosmos/cosmos-sdk/client/keys"
    "github.com/cosmos/cosmos-sdk/version"
    "github.com/cosmos/cosmos-sdk/x/auth/client/cli"

    // NOORCHAIN modules
    noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

func main() {
    rootCmd := NewRootCmd()

    if err := Execute(rootCmd); err != nil {
        os.Exit(1)
    }
}

// NewRootCmd assemble toutes les commandes CLI de NOORCHAIN.
// - keys
// - tx
// - query
// - PoSS (noorsignal)
// - versions
//
// C’est une version minimaliste mais fonctionnelle.
func NewRootCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "noord",
        Short: "NOORCHAIN Daemon CLI",
    }

    // -----------------------------
    // Initialisation du client SDK
    // -----------------------------
    initClientCtx := client.Context{}.
        WithCodec(noorsignaltypes.ModuleCdc).
        WithInput(os.Stdin).
        WithHomeDir(os.ExpandEnv("$HOME/.noord")).
        WithViper("NOORCHAIN")

    cmd.PersistentFlags().String(flags.FlagHome, initClientCtx.HomeDir, "node home directory")

    // -----------------------------
    // Sous-commandes Cosmos SDK
    // -----------------------------
    cmd.AddCommand(
        keys.Commands(initClientCtx.HomeDir),          // noord keys ...
        cli.NewTxCmd(),                               // noord tx ...
        cli.NewQueryCmd(),                            // noord query ...
        version.NewVersionCommand(),                  // noord version
    )

    // -----------------------------
    // Ajout PoSS (QUERY)
    // -----------------------------
    cmd.AddCommand(
        NewNoorSignalQueryCmd(),                      // noord query noorsignal ...
    )

    // -----------------------------
    // Ajout PoSS (TX)
    // -----------------------------
    cmd.AddCommand(
        NewNoorSignalTxCmd(),                         // noord tx noorsignal ...
    )

    return cmd
}

// Execute lance la commande racine.
func Execute(rootCmd *cobra.Command) error {
    return rootCmd.Execute()
}
