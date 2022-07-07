package cli

import (
	"github.com/bianjieai/ddc-go/ddc/core/auth"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	FlagsLeaderDID = "leader_did"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "auth",
		Short:                      "ddc transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewAddAccountTxCmd(),
	)
	return cmd
}

// NewAddAccountTxCmd returns a CLI command handler for creating a account transaction.
func NewAddAccountTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [address] [did] [name]",
		Short: "Add a account.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			_, err = sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			leaderDID, err := cmd.Flags().GetString(FlagsLeaderDID)
			if err != nil {
				return err
			}

			msg := &auth.MsgAddAccount{
				Address:   args[0],
				Did:       args[1],
				Name:      args[2],
				LeaderDID: leaderDID,
				Sender:    clientCtx.GetFromAddress().String(),
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagsLeaderDID, "", "the leader did of the account added")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
