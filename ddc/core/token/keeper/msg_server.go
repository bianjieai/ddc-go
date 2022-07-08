package keeper

import (
	context "context"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ token.MsgServer = Keeper{}

// Approve implements token.MsgServer
// implement:
//  - approve
// reference:
//  - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L224
func (k Keeper) Approve(goctx context.Context, msg *token.MsgApprove) (res *token.MsgApproveResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	err = k.approve(ctx, msg.Denom, msg.TokenID, msg.Operator, msg.To)
	if err != nil {
		return nil, err
	}

	return
}

// ApproveForAll implements token.MsgServer
func (Keeper) ApproveForAll(context.Context, *token.MsgApproveForAll) (*token.MsgApproveForAllResponse, error) {
	panic("unimplemented")
}

// BatchBurn implements token.MsgServer
func (Keeper) BatchBurn(context.Context, *token.MsgBatchBurn) (*token.MsgBatchBurnResponse, error) {
	panic("unimplemented")
}

// BatchTransfer implements token.MsgServer
func (Keeper) BatchTransfer(context.Context, *token.MsgBatchTransfer) (*token.MsgBatchTransferResponse, error) {
	panic("unimplemented")
}

// Freeze implements token.MsgServer
func (Keeper) Freeze(context.Context, *token.MsgFreeze) (*token.MsgFreezeResponse, error) {
	panic("unimplemented")
}

// Unfreeze implements token.MsgServer
func (Keeper) Unfreeze(context.Context, *token.MsgUnfreeze) (*token.MsgUnfreezeResponse, error) {
	panic("unimplemented")
}
