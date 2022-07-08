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

// CrossPlatformAble implements auth.QueryServer
// implements: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L497
func (k Keeper) CrossPlatformAble(goctx context.Context, req *auth.QueryCrossPlatformAbleRequest) (res *auth.QueryCrossPlatformAbleResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	fromAcc, err := k.GetAccount(ctx, req.From)
	if err != nil {
		return nil, err
	}
	if !k.isActive(fromAcc) {
		return nil, sdkerrors.Wrapf(auth.ErrAccountNotActive, "account: %s is not active", req.From)
	}

	toAcc, err := k.GetAccount(ctx, req.To)
	if err != nil {
		return nil, err
	}
	if !k.isActive(toAcc) {
		return nil, sdkerrors.Wrapf(auth.ErrAccountNotActive, "account: %s is not active", req.To)
	}

	// 2. Check role
	// a. All are platform roles
	if fromAcc.Role == core.Role_PLATFORM_MANAGER && toAcc.Role == core.Role_PLATFORM_MANAGER {
		res.Enabled = (fromAcc.LeaderDID == toAcc.LeaderDID && k.crossPlatformApproval(ctx, fromAcc.DID, toAcc.DID))
		return
	}

	// b. `from` is the platform, `to` is the consumer
	if fromAcc.Role == core.Role_PLATFORM_MANAGER && toAcc.Role == core.Role_CONSUMER {
		res.Enabled = k.crossPlatformApproval(ctx, fromAcc.DID, toAcc.LeaderDID)
		return
	}

	// c. `to` is the platform, `from` is the consumer
	if fromAcc.Role == core.Role_CONSUMER && toAcc.Role == core.Role_PLATFORM_MANAGER {
		res.Enabled = k.crossPlatformApproval(ctx, fromAcc.LeaderDID, toAcc.DID)
		return
	}
	// d. Both are consumers
	if fromAcc.Role == core.Role_CONSUMER && toAcc.Role == core.Role_CONSUMER {
		res.Enabled = k.crossPlatformApproval(ctx, fromAcc.LeaderDID, toAcc.LeaderDID)
		return
	}

	res.Enabled = false
	return
}

// DDCs implements auth.QueryServer
func (k Keeper) DDCs(goctx context.Context, req *auth.QueryDDCsRequest) (*auth.QueryDDCsResponse, error) {
	panic("unimplemented")
}

// SwitcherState implements auth.QueryServer
func (k Keeper) SwitcherState(goctx context.Context, req *auth.QuerySwitcherStateRequest) (*auth.QuerySwitcherStateResponse, error) {
	panic("unimplemented")
}
