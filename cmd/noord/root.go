package main

import (
	"io"
	"os"
	"path/filepath"

	sdklog "cosmossdk.io/log"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
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
	return svrcmd.Execute(rootCmd, app.AppName, app.DefaultNodeHome)
}

func NewRootCmd() *cobra.Command {
	// --- Bech32 prefixes (global config) ---
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("noor", "noorpub")
	cfg.SetBech32PrefixForValidator("noorvaloper", "noorvaloperpub")
	cfg.SetBech32PrefixForConsensusNode("noorvalcons", "noorvalconspub")
	cfg.Seal()

	// NOTE: app.MakeEncodingConfig returns ONE value (no error) on this branch.
	enc := app.MakeEncodingConfig()

	// Address codecs required by v0.50 genutil commands
	accAddrCodec := addresscodec.NewBech32Codec("noor")
	valAddrCodec := addresscodec.NewBech32Codec("noorvaloper")

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
		Short: "NOORCHAIN daemon (Phase 2 skeleton)",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			var err error

			// --- Client context ---
			initClientCtx = initClientCtx.WithCmdContext(cmd.Context())
			initClientCtx, err = client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
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

			// Ensure config/ exists before genutil tries to write node_key.json
			if err := os.MkdirAll(filepath.Join(homeDir, "config"), 0o711); err != nil {
				return err
			}

			// Propagate --chain-id into server Viper (appOpts)
			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			if chainID != "" {
				serverCtx.Viper.Set(flags.FlagChainID, chainID) // "chain-id"
				serverCtx.Viper.Set("chain-id", chainID)        // defensive alias
			}

			return server.SetCmdServerContext(cmd, serverCtx)
		},
	}

	// Global flags
	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	// Standard server commands (start, etc.)
	server.AddCommands(rootCmd, app.DefaultNodeHome, newApp, appExport, addModuleInitFlags)

	// Keys
	rootCmd.AddCommand(keys.Commands())

	// --- GENUTIL commands (explicit codecs + balances iterator) ---
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

func newApp(
	logger sdklog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	_ = appOpts // Phase 2: app.NewNoorchainApp does NOT consume AppOptions on this branch.
	return app.NewNoorchainApp(logger, db, traceStore)
}

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

	// Phase 2: no exporter yet
	return servertypes.ExportedApp{}, io.EOF
}
