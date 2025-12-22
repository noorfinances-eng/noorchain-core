package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	sdklog "cosmossdk.io/log"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/server"
	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	"github.com/noorfinances-eng/noorchain-core/app"
)

func Execute() error {
	rootCmd := NewRootCmd()
	// Cosmos SDK v0.53 standard executor (ajoute les flags log + contextes client/server)
	return servercmd.Execute(rootCmd, strings.ToUpper(app.AppName), app.DefaultNodeHome)
}

func NewRootCmd() *cobra.Command {
	// --- Prefixes Bech32 (legacy global config) ---
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("noor", "noorpub")
	cfg.SetBech32PrefixForValidator("noorvaloper", "noorvaloperpub")
	cfg.SetBech32PrefixForConsensusNode("noorvalcons", "noorvalconspub")
	cfg.Seal()

	// Encodage (Proto + Amino) de l'app
	enc := app.MakeEncodingConfig()

	// Address codecs (v0.53) pour les commandes genutil
	accAddrCodec := address.NewBech32Codec("noor")
	valAddrCodec := address.NewBech32Codec("noorvaloper")

	// Contexte client de base
	initClientCtx := client.Context{}.
		WithCodec(enc.Marshaler).
		WithInterfaceRegistry(enc.InterfaceRegistry).
		WithTxConfig(enc.TxConfig).
		WithLegacyAmino(enc.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome)

	rootCmd := &cobra.Command{
		Use:   app.AppName,
		Short: app.AppName,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			var err error

			// --- Client context ---
			clientCtx := initClientCtx.WithCmdContext(cmd.Context())
			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			if err := client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}

			// --- Server context ---
			serverCtx := server.GetServerContextFromCmd(cmd)

			homeDir, _ := cmd.Flags().GetString(flags.FlagHome)
			if homeDir == "" {
				homeDir = app.DefaultNodeHome
			}

			// Sync RootDir with --home
			serverCtx.Config.SetRoot(homeDir)

			// Ensure config/ exists avant que genutil essaie d'écrire node_key.json
			if err := os.MkdirAll(filepath.Join(homeDir, "config"), 0o711); err != nil {
				return err
			}

			// Propagate --chain-id into server Viper (appOpts)
			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			if chainID != "" {
				serverCtx.Viper.Set(flags.FlagChainID, chainID) // "chain-id"
				serverCtx.Viper.Set("chain-id", chainID)        // alias défensif
			}

			return server.SetCmdServerContext(cmd, serverCtx)
		},
	}

	// Flag global
	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	// Commandes serveur standard (start, etc.)
	server.AddCommands(rootCmd, app.DefaultNodeHome, newApp, appExport, addModuleInitFlags)

	// Commandes keys
	rootCmd.AddCommand(keys.Commands())

	// --- GENUTIL commands (init, add-genesis-account, gentx, collect, validate) ---
	var mbm module.BasicManager = app.ModuleBasics
	var genBalIterator genutiltypes.GenesisBalancesIterator = banktypes.GenesisBalancesIterator{}
	msgValidator := genutiltypes.DefaultMessageValidator

	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.AddGenesisAccountCmd(app.DefaultNodeHome, accAddrCodec),
		genutilcli.GenTxCmd(mbm, enc.TxConfig, genBalIterator, app.DefaultNodeHome, valAddrCodec),
		genutilcli.CollectGenTxsCmd(genBalIterator, app.DefaultNodeHome, msgValidator, valAddrCodec),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
	)

	return rootCmd
}

func addModuleInitFlags(startCmd *cobra.Command) {}

// newApp construit NoorchainApp avec le socle minimal pour Phase 2
func newApp(
	logger sdklog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	_ = appOpts
	return app.NewNoorchainApp(logger, db, traceStore)
}

// appExport est un stub minimal pour satisfaire l'interface Application
func appExport(
	logger sdklog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	_ = logger
	_ = db
	_ = traceStore
	_ = height
	_ = forZeroHeight
	_ = jailAllowedAddrs
	_ = appOpts
	_ = modulesToExport

	return servertypes.ExportedApp{}, io.EOF
}
