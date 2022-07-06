package cli

import (
	"github.com/bianjieai/ddc-go/ddc"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
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
	// TODO
	)

	return cmd
}
