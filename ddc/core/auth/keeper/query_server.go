package keeper

import (
	context "context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bianjieai/ddc-go/ddc/core"
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
func (k Keeper) Functions(goctx context.Context, req *auth.QueryFunctionsRequest) (*auth.QueryFunctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if _, ok := core.Role_value[req.Role.String()]; !ok {
		return nil, sdkerrors.Wrap(auth.ErrInvalidRole, "role not exist")
	}

	if _, ok := core.Protocol_value[req.Protocol.String()]; !ok {
		return nil, sdkerrors.Wrap(auth.ErrInvalidProtocol, "protocol not exist")
	}

	if len(strings.TrimSpace(req.Denom)) == 0 {
		return nil, sdkerrors.Wrap(auth.ErrInvalidDenom, "denom cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	return &auth.QueryFunctionsResponse{
		Functions: k.getFunction(ctx, req.Role, req.Protocol, req.Denom),
	}, nil
}

// hasFunctionPermission
// switcherStateOfPlatform
// onePlatformCheck
// crossPlatformCheck
