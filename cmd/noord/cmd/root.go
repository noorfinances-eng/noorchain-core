package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	tmcfg "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	// v0.46: iterator for balances used by gentx/collect-gentxs
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/noorfinances-eng/noorchain-core/app"
)

// EnvPrefix is used by viper/env for server flags.
const EnvPrefix = "NOORCHAIN"

// MakeEncodingConfig builds a minimal encoding config compatible with init/genesis.
func MakeEncodingConfig(bm module.BasicManager) (codec.Codec, codectypes.InterfaceRegistry, client.TxConfig, *codec.LegacyAmino) {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	// Register module interfaces + legacy amino.
	bm.RegisterInterfaces(interfaceRegistry)
	bm.RegisterLegacyAminoCodec(amino)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	txCfg := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	return cdc, interfaceRegistry, txCfg, amino
}

// NewRootCmd wires the Cosmos SDK CLI for NOORCHAIN (init + keys + gentx + collect-gentxs + start).
func NewRootCmd() *cobra.Command {
	cdc, interfaceRegistry, txCfg, amino := MakeEncodingConfig(app.ModuleBasics)

	initClientCtx := client.Context{}.
		WithCodec(cdc).
		WithInterfaceRegistry(interfaceRegistry).
		WithTxConfig(txCfg).
		WithLegacyAmino(amino).
		WithInput(os.Stdin).
		WithHomeDir(app.DefaultNodeHome).
		WithViper(EnvPrefix)

	rootCmd := &cobra.Command{
		Use:   "noord",
		Short: "NOORCHAIN node daemon (public testnet) — minimal CLI",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			var err error
			initClientCtx, err = client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			// Cosmos SDK default server config
			return sdkserver.InterceptConfigsPreRunHandler(
				cmd,
				serverconfig.DefaultConfigTemplate,
				serverconfig.DefaultConfig(),
				tmcfg.DefaultConfig(),
			)
		},
	}

	// init (writes config/, genesis.json, node keys)
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
	)

	// keys (Cosmos SDK v0.46): enables `noord keys add ...`
	rootCmd.AddCommand(
		keys.Commands(app.DefaultNodeHome),
	)

	// gentx + collect-gentxs (Cosmos SDK v0.46)
	// IMPORTANT: pass txCfg directly (v0.46 has no client.TxEncodingConfig type)
	rootCmd.AddCommand(
		genutilcli.GenTxCmd(
			app.ModuleBasics,
			txCfg,
			banktypes.GenesisBalancesIterator{},
			app.DefaultNodeHome,
		),
		genutilcli.CollectGenTxsCmd(
			banktypes.GenesisBalancesIterator{},
			app.DefaultNodeHome,
		),
	)

	// start + server commands (Cosmos SDK v0.46.x signature)
	creator := appCreator{}
	sdkserver.AddCommands(
		rootCmd,
		app.DefaultNodeHome,
		creator.newApp,
		creator.appExport,
		addModuleInitFlags,
	)

	// Basic flags
	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	// Seal SDK config (safe default)
	cfg := sdk.GetConfig()
	cfg.Seal()

	return rootCmd
}

func addModuleInitFlags(startCmd *cobra.Command) {
	// Keep empty for Phase 8.A minimal
}

type appCreator struct{}

func (a appCreator) newApp(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	return app.NewApp(logger, db, traceStore, appOpts)
}

// appExport keeps command compatibility; not used in Phase 8.A.
func (a appCreator) appExport(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	return servertypes.ExportedApp{}, nil
}
