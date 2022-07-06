package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/bianjieai/ddc-go/ddc"
	authcli "github.com/bianjieai/ddc-go/ddc/core/auth/cli"
	feecli "github.com/bianjieai/ddc-go/ddc/core/fee/cli"
	tokencli "github.com/bianjieai/ddc-go/ddc/core/token/cli"
)

// GetQueryCmd returns the cli query commands for the module.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        ddc.ModuleName,
		Short:                      "Querying commands for the ddc module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcli.GetQueryCmd(),
		feecli.GetQueryCmd(),
		tokencli.GetQueryCmd(),
	)
	return cmd
}
