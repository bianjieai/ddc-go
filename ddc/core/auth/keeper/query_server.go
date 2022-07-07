package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

var _ auth.QueryServer = Keeper{}

// Account implements auth.QueryServer
func (k Keeper) Account(goctx context.Context, req *auth.QueryAccountRequest) (*auth.QueryAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "address cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %s", err.Error())
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	account, err := k.GetAccount(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	return &auth.QueryAccountResponse{
		AccountInfo: account,
	}, nil
}

// Functions implements auth.QueryServer
func (Keeper) Functions(goctx context.Context, req *auth.QueryFunctionsRequest) (*auth.QueryFunctionsResponse, error) {
	panic("unimplemented")
}
