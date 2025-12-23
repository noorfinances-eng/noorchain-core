package main

import (
	"io"
	"os"
	"strings"

	sdklog "cosmossdk.io/log"

	cometcfg "github.com/cometbft/cometbft/config"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	clientcfg "github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/server"
	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	"github.com/noorfinances-eng/noorchain-core/app"
)

// Execute is the main entry point called by main.go.
func Execute() error {
	rootCmd := NewRootCmd()
	return servercmd.Execute(rootCmd, strings.ToUpper(app.AppName), app.DefaultNodeHome)
}

// NewRootCmd builds the root "noorchain" command.
func NewRootCmd() *cobra.Command {
	// Bech32 prefixes
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("noor", "noorpub")
	cfg.SetBech32PrefixForValidator("noorvaloper", "noorvaloperpub")
	cfg.SetBech32PrefixForConsensusNode("noorvalcons", "noorvalconspub")
	cfg.Seal()

	// Encoding config
	enc := app.MakeEncodingConfig()

	accAddrCodec := addresscodec.NewBech32Codec("noor")
	valAddrCodec := addresscodec.NewBech32Codec("noorvaloper")

	// Base client context (Viper non-nil)
	initClientCtx := client.Context{
		Viper:             viper.New(),
		Codec:             enc.Marshaler,
		InterfaceRegistry: enc.InterfaceRegistry,
		TxConfig:          enc.TxConfig,
		LegacyAmino:       enc.Amino,
		Input:             os.Stdin,
		AccountRetriever:  authtypes.AccountRetriever{},
		HomeDir:           app.DefaultNodeHome,
	}

	rootCmd := &cobra.Command{
		Use:   "noorchain",
		Short: "Noorchain core-local node",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// 1) Client context avec flags
			clientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// 2) Client config (client.toml)
			clientCtx, err = clientcfg.ReadFromClientConfig(clientCtx)
			if err != nil {
				return err
			}
			if err := client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}

			// 3) Interception standard des configs (config.toml, app.toml, flags)
			appTemplate := serverconfig.DefaultConfigTemplate
			appCfg := serverconfig.DefaultConfig()
			tmCfg := cometcfg.DefaultConfig()

			return server.InterceptConfigsPreRunHandler(cmd, appTemplate, appCfg, tmCfg)
		},
	}

	// Flag global chain-id (optionnel)
	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	// Ajout des commandes serveur (start, unsafe-reset-all, etc.)
	server.AddCommands(rootCmd, app.DefaultNodeHome, newApp, appExport, addModuleInitFlags)

	// Commandes keys
	rootCmd.AddCommand(keys.Commands())

	// genutil / genesis
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

// addModuleInitFlags is required by server.AddCommands; no custom start flags for core-local.
func addModuleInitFlags(_ *cobra.Command) {}

// newApp is the AppCreator used by the "start" command.
func newApp(
	logger sdklog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	return app.NewNoorchainApp(logger, db, traceStore, appOpts)
}

// appExport is the AppExporter used by "export" (non implémenté pour core-local).
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

	// Export non supporté pour le profil core-local minimal
	return servertypes.ExportedApp{}, io.EOF
}
