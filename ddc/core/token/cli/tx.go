package cli

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "token",
		Short:                      "ddc transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewApproveCmd(),
		NewApproveForAllCmd(),
		NewBatchBurnCmd(),
		NewBatchTransferCmd(),
		NewFreezeCmd(),
		NewUnFreezeCmd(),
	)
	return cmd
}

// NewApproveCmd returns a CLI command handler for approving a ddc721
func NewApproveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-ddc [denomID] [tokenID] [to]",
		Short: "approve a ddc721",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &token.MsgApprove{
				Denom:    args[0],
				TokenID:  args[1],
				Operator: clientCtx.GetFromAddress().String(),
				To:       args[2],
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewApproveForAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-account [protocol] [denomID] [operator]",
		Short: "approve an account",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &token.MsgApproveForAll{
				Protocol: core.Protocol(core.Protocol_value[args[0]]),
				Denom:    args[1],
				Operator: args[1],
				Sender:   clientCtx.GetFromAddress().String(),
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewBatchBurnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch-burn [protocol] [denomID] [tokenIDs...]",
		Short: "batch burn ddc",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenIDs := parseTokenIDs(args[2])
			if tokenIDs == nil {
				return sdkerrors.Wrapf(token.ErrInvalidTokenID, "invalid tokenIDs syntax")
			}

			msg := &token.MsgBatchBurn{
				Protocol: core.Protocol(core.Protocol_value[args[0]]),
				Denom:    args[1],
				TokenIDs: nil,
				Operator: clientCtx.GetFromAddress().String(),
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewBatchTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch-burn [protocol] [denomID] [tokenID1,tokenID2,tokenID3] [amount1,amount2,amount3] [from] [to]",
		Short: "batch burn ddc",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenIDs := parseTokenIDs(args[2])
			if tokenIDs == nil {
				return sdkerrors.Wrapf(token.ErrInvalidTokenID, "invalid tokenIDs syntax")
			}
			amounts := parseAmounts(args[3])
			if amounts == nil {
				return sdkerrors.Wrapf(token.ErrInvalidAmount, "invalid amounts syntax")
			}

			msg := &token.MsgBatchTransfer{
				Protocol: core.Protocol(core.Protocol_value[args[0]]),
				Denom:    args[1],
				TokenIDs: tokenIDs,
				Amount:   amounts,
				From:     args[4],
				To:       args[5],
				Sender:   clientCtx.GetFromAddress().String(),
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewFreezeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freeze [protocol] [denomID] [tokenID]",
		Short: "freeze a ddc",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &token.MsgFreeze{
				Protocol: core.Protocol(core.Protocol_value[args[0]]),
				Denom:    args[1],
				TokenID:  args[2],
				Operator: clientCtx.GetFromAddress().String(),
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewUnFreezeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unfreeze [protocol] [denomID] [tokenID]",
		Short: "unfreeze a ddc",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &token.MsgUnfreeze{
				Protocol: core.Protocol(core.Protocol_value[args[0]]),
				Denom:    args[1],
				TokenID:  args[2],
				Operator: clientCtx.GetFromAddress().String(),
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func parseTokenIDs(arg string) []string {
	s := strings.Split(arg, ",")
	res := make([]string, len(s), len(s))
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
		if s[i] != "" {
			res = append(res, s[i])
		}
	}
	return res
}

func parseAmounts(arg string) []uint64 {
	s := strings.Split(arg, ",")
	res := make([]uint64, len(s), len(s))
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
		val, err := strconv.ParseUint(s[i], 10, 64)
		if err != nil {
			return nil
		}
		res = append(res, val)
	}
	return res
}
