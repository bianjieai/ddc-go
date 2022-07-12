package keeper

import (
	context "context"
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

var _ token.QueryServer = Keeper{}

// ApprovedForAll implements token.QueryServer
func (k Keeper) ApprovedForAll(goctx context.Context, req *token.QueryApprovedForAllRequest) (*token.QueryApprovedForAllResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.Protocol != core.Protocol_NFT && req.Protocol != core.Protocol_MT {
		return nil, sdkerrors.Wrapf(token.ErrInvalidProtocol, "invalid protocol")
	}

	if len(strings.TrimSpace(req.Denom)) == 0 {
		return nil, sdkerrors.Wrapf(token.ErrInvalidDenom, "denom cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %s", err.Error())
	}

	ctx := sdk.UnwrapSDKContext(goctx)

	return &token.QueryApprovedForAllResponse{Operator: k.getAccountsApproval(ctx, req.Protocol, req.Owner, req.Denom)}, nil
}

// TokenApproved implements token.QueryServer
func (k Keeper) TokenApproved(goctx context.Context, req *token.QueryTokenApprovedRequest) (*token.QueryTokenApprovedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if len(strings.TrimSpace(req.Denom)) == 0 {
		return nil, sdkerrors.Wrapf(token.ErrInvalidDenom, "denom cannot be empty")
	}

	if len(strings.TrimSpace(req.TokenId)) == 0 {
		return nil, sdkerrors.Wrapf(token.ErrInvalidDenom, "tokenID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goctx)

	return &token.QueryTokenApprovedResponse{Operator: k.getDDCApproval(ctx, core.Protocol_NFT, req.Denom, req.TokenId)}, nil
}
