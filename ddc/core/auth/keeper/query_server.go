package keeper

import (
	context "context"

	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

var _ auth.QueryServer = Keeper{}

// Account implements auth.QueryServer
func (Keeper) Account(context.Context, *auth.QueryAccountRequest) (*auth.QueryAccountResponse, error) {
	panic("unimplemented")
}

// Functions implements auth.QueryServer
func (Keeper) Functions(context.Context, *auth.QueryFunctionsRequest) (*auth.QueryFunctionsResponse, error) {
	panic("unimplemented")
}
