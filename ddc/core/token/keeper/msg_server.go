package keeper

import (
	"context"
	"github.com/bianjieai/ddc-go/ddc/core"
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

	// TODO: event

	return
}

// ApproveForAll implements token.MsgServer
// implement:
// - setApprovalForAll
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L274
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L191
func (k Keeper) ApproveForAll(goctx context.Context, msg *token.MsgApproveForAll) (res *token.MsgApproveForAllResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	switch msg.Protocol {
	case core.Protocol_NFT:
		err = k.setApproveForAllDDC721(ctx, msg.Denom, msg.Sender, msg.Operator)
	case core.Protocol_MT:
		err = k.setApproveForAllDDC1155(ctx, msg.Denom, msg.Sender, msg.Operator)
	}

	if err != nil {
		return nil, err
	}

	// TODO: event

	return
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
