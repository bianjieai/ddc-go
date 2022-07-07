package ddc

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState constructs a new GenesisState instance
func NewGenesisState() *core.GenesisState {
	return &core.GenesisState{}
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() *core.GenesisState {
	return &core.GenesisState{}
}

// ValidateGenesis validates the provided identity genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data core.GenesisState) error {
	//TODO
	return nil
}

// InitGenesis - store genesis parameters
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data core.GenesisState) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
	}
	//TODO
}

// ExportGenesis - output genesis parameters
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *core.GenesisState {
	//TODO
	return NewGenesisState()
}
