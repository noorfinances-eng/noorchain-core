package cmd

import (
	"os"

	"github.com/spf13/cobra"

	tmcfg "github.com/tendermint/tendermint/config"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

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

// NewRootCmd wires a minimal Cosmos SDK CLI for NOORCHAIN (init only).
func NewRootCmd() *cobra.Command {
	cdc, interfaceRegistry, txCfg, amino := MakeEncodingConfig(app.ModuleBasics)

	initClientCtx := client.Context{}.
		WithCodec(cdc).
		WithInterfaceRegistry(interfaceRegistry).
		WithTxConfig(txCfg).
		WithLegacyAmino(amino).
		WithInput(os.Stdin).
		WithHomeDir(app.DefaultNodeHome).
		WithViper(EnvPrefix) // CRITICAL: avoid nil viper in ReadFromClientConfig

	rootCmd := &cobra.Command{
		Use:   "noord",
		Short: "NOORCHAIN node daemon (public testnet) — CLI minimal (init)",
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

			// Cosmos SDK default server config (no start wiring in Phase 8.A)
			return sdkserver.InterceptConfigsPreRunHandler(
				cmd,
				serverconfig.DefaultConfigTemplate,
				serverconfig.DefaultConfig(),
				tmcfg.DefaultConfig(),
			)
		},
	}

	// Minimal command required now: init (writes config/, genesis.json, node keys).
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
	)

	// Basic flags
	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	// Seal SDK config (safe default)
	cfg := sdk.GetConfig()
	cfg.Seal()

	return rootCmd
}
