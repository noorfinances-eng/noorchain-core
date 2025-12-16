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

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/std"

	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	simdcmd "github.com/cosmos/cosmos-sdk/simapp/simd/cmd"

	"github.com/noorfinances-eng/noorchain-core/app"
)

const EnvPrefix = "NOORCHAIN"

func MakeEncodingConfig(bm module.BasicManager) (codec.Codec, codectypes.InterfaceRegistry, client.TxConfig, *codec.LegacyAmino) {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	std.RegisterLegacyAminoCodec(amino)
	std.RegisterInterfaces(interfaceRegistry)

	cryptocodec.RegisterInterfaces(interfaceRegistry)

	bm.RegisterInterfaces(interfaceRegistry)
	bm.RegisterLegacyAminoCodec(amino)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	txCfg := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	return cdc, interfaceRegistry, txCfg, amino
}

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

			return sdkserver.InterceptConfigsPreRunHandler(
				cmd,
				serverconfig.DefaultConfigTemplate,
				serverconfig.DefaultConfig(),
				tmcfg.DefaultConfig(),
			)
		},
	}

	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
	)

	rootCmd.AddCommand(
		keys.Commands(app.DefaultNodeHome),
	)

	rootCmd.AddCommand(
		simdcmd.AddGenesisAccountCmd(app.DefaultNodeHome),
	)

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

	creator := appCreator{}
	sdkserver.AddCommands(
		rootCmd,
		app.DefaultNodeHome,
		creator.newApp,
		creator.appExport,
		addModuleInitFlags,
	)

	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	cfg := sdk.GetConfig()
	cfg.Seal()

	return rootCmd
}

func addModuleInitFlags(startCmd *cobra.Command) {}

type appCreator struct{}

func (a appCreator) newApp(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	return app.NewApp(logger, db, traceStore, appOpts)
}

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
