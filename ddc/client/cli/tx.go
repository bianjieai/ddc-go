package cli

import (
	"github.com/bianjieai/ddc-go/ddc"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        ddc.ModuleName,
		Short:                      "ddc transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
	// TODO
	)
	return cmd
}
