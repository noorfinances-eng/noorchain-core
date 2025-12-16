package cmd

import (
	"encoding/json"
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

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

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

// -----------------------------------------------------------------------------
// GenesisBalancesIterator (Cosmos SDK v0.46)
// Required by: gentx / collect-gentxs commands.
// -----------------------------------------------------------------------------

type noorGenesisBalancesIterator struct{}

func (noorGenesisBalancesIterator) IterateGenesisBalances(
	cdc codec.JSONCodec,
	appState map[string]json.RawMessage,
	cb func(addr sdk.AccAddress, coins sdk.Coins) (stop bool),
) error {
	raw, ok := appState[banktypes.ModuleName]
	if !ok || len(raw) == 0 {
		return nil
	}

	var bankGen banktypes.GenesisState
	if err := cdc.UnmarshalJSON(raw, &bankGen); err != nil {
		return err
	}

	for _, b := range bankGen.Balances {
		addr, err := sdk.AccAddressFromBech32(b.Address)
		if err != nil {
			// ignore invalid bech32 in genesis to stay robust
			continue
		}
		if cb(addr, b.Coins) {
			break
		}
	}
	return nil
}

// NewRootCmd wires the Cosmos SDK CLI for NOORCHAIN (init + start + genesis tooling).
func NewRootCmd() *cobra.Command {
	cdc, interfaceRegistry, txCfg, amino := MakeEncodingConfig(app.ModuleBasics)

	// v0.46 expects this struct for genutil commands
	txEncCfg := client.TxEncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         cdc,
		TxConfig:          txCfg,
		Amino:             amino,
	}

	genBalIterator := noorGenesisBalancesIterator{}

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
		Short: "NOORCHAIN node daemon (public testnet) — init/start + genesis tools (SDK v0.46)",
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

	// -------------------
	// Core CLI commands
	// -------------------

	// init (writes config/, genesis.json, node keys)
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
	)

	// keys (enables: `noord keys add ...`)
	rootCmd.AddCommand(
		keys.Commands(app.DefaultNodeHome),
	)

	// add-genesis-account / gentx / collect-gentxs (SDK v0.46 needs a balances iterator)
	rootCmd.AddCommand(
		genutilcli.AddGenesisAccountCmd(app.DefaultNodeHome),
		genutilcli.GenTxCmd(app.ModuleBasics, txEncCfg, genBalIterator, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(app.ModuleBasics, genBalIterator, app.DefaultNodeHome),
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
	// Minimal safe export (unused here)
	return servertypes.ExportedApp{
		AppState: nil,
		Height:  height,
	}, nil
}

// Ensure we satisfy genutiltypes.GenesisBalancesIterator at compile-time.
var _ genutiltypes.GenesisBalancesIterator = noorGenesisBalancesIterator{}
