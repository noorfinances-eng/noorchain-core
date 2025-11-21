package main

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// NewNoorSignalQueryCmd assemble les commandes "noord query noorsignal ..."
func NewNoorSignalQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "noorsignal",
		Short: "Querying commands for the PoSS module",
	}

	// query noorsignal config
	cmd.AddCommand(
		&cobra.Command{
			Use:   "config",
			Short: "Query the PoSS configuration",
			RunE: func(cmd *cobra.Command, args []string) error {
				clientCtx, err := client.GetClientQueryContext(cmd)
				if err != nil {
					return err
				}

				queryClient := noorsignaltypes.NewQueryClient(clientCtx)
				res, err := queryClient.Config(cmd.Context(), &noorsignaltypes.QueryConfigRequest{})
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			},
		},
	)

	return cmd
}

// NewNoorSignalTxCmd assemble les commandes "noord tx noorsignal ..."
func NewNoorSignalTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "noorsignal",
		Short: "Transaction commands for the PoSS module",
	}

	// tx noorsignal submit
	cmd.AddCommand(
		&cobra.Command{
			Use:   "submit",
			Short: "Submit a new PoSS signal",
			RunE: func(cmd *cobra.Command, args []string) error {
				// Ce sera rempli plus tard (testnet)
				return nil
			},
		},
	)

	return cmd
}
