package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	dbm "github.com/tendermint/tm-db"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"github.com/noorfinances-eng/noorchain-core/app"
)

const (
	appName         = "noorchain"
	defaultNodeHome = ".noorchain"
)

func main() {
	rootCmd := NewRootCmd()
	if err := svrcmd.Execute(rootCmd, appName, defaultNodeHome); err != nil {
		os.Exit(1)
	}
}

// NewRootCmd creates the root CLI command for the NOORCHAIN node.
// Example usage after build:
//   ./noord start
//   ./noord unsafe-reset-all
func NewRootCmd() *cobra.Command {
	encodingConfig := app.MakeEncodingConfig()

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithInput(os.Stdin).
		WithHomeDir(defaultNodeHome).
		WithViper(appName)

	rootCmd := &cobra.Command{
		Use:   "noord",
		Short: "NOORCHAIN full node (Cosmos SDK + Ethermint + PoSS)",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Stdout / stderr corrects
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			// Lecture des flags client (--home, --node, etc.)
			clientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// Lecture de la config client (~/.noorchain/config/client.toml si existante)
			clientCtx, err = config.ReadFromClientConfig(clientCtx)
			if err != nil {
				return err
			}

			// Injection dans le contexte Cobra
			if err := client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}

			// Hook standard Cosmos SDK (config, app.toml, etc.)
			return server.InterceptConfigsPreRunHandler(cmd, "", nil)
		},
	}

	// Flag global --home
	rootCmd.PersistentFlags().String(flags.FlagHome, defaultNodeHome, "directory for config and data")

	// Ajout des commandes serveur (start, unsafe-reset-all, tendermint, etc.)
	server.AddCommands(rootCmd, defaultNodeHome, newApp, nil)

	return rootCmd
}

// newApp instancie l'application NOORCHAIN pour le serveur Cosmos.
func newApp(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	// On charge toujours la dernière version du store (loadLatest = true)
	return app.NewNoorchainApp(
		logger,
		db,
		traceStore,
		true,
		appOpts,
	)
}
