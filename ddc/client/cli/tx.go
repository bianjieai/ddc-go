package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/bianjieai/ddc-go/ddc"
	authcli "github.com/bianjieai/ddc-go/ddc/core/auth/cli"
	feecli "github.com/bianjieai/ddc-go/ddc/core/fee/cli"
	tokencli "github.com/bianjieai/ddc-go/ddc/core/token/cli"
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
		authcli.NewTxCmd(),
		feecli.NewTxCmd(),
		tokencli.NewTxCmd(),
	)
	return cmd
}
