package cli

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// GetQueryCmd returns the cli query commands for the module.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "token",
		Short:                      "Querying commands for the ddc module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetApprovedForAll(),
		GetTokenApproved(),
	)
	return cmd
}

func GetApprovedForAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-account [protocol] [denomID] [owner]",
		Short: "query approved accounts",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := token.NewQueryClient(clientCtx)
			res, err := queryClient.ApprovedForAll(cmd.Context(), &token.QueryApprovedForAllRequest{
				Protocol: core.Protocol(core.Protocol_value[args[0]]),
				Denom:    args[1],
				Owner:    args[2],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetTokenApproved() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-account [denomID] [tokenID]",
		Short: "query approved account of a ddc721",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := token.NewQueryClient(clientCtx)
			res, err := queryClient.TokenApproved(cmd.Context(), &token.QueryTokenApprovedRequest{
				Denom:   args[0],
				TokenId: args[1],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
