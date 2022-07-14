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
func (k Keeper) Approve(goctx context.Context,
	msg *token.MsgApprove,
) (res *token.MsgApproveResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	err = k.approve(ctx, msg.Denom, msg.TokenID, msg.Operator, msg.To)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvents(&token.EventApprove{
		Denom:    msg.Denom,
		TokenID:  msg.TokenID,
		Operator: msg.Operator,
		To:       msg.To,
	})

	return
}

// ApproveForAll implements token.MsgServer
// implement:
// - setApprovalForAll
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L274
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L191
func (k Keeper) ApproveForAll(goctx context.Context,
	msg *token.MsgApproveForAll,
) (res *token.MsgApproveForAllResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	switch msg.Protocol {
	case core.Protocol_NFT:
		err = k.setApproveForAllDDC721(ctx, msg.Denom, msg.Sender, msg.Operator, msg.Protocol)
	case core.Protocol_MT:
		err = k.setApproveForAllDDC1155(ctx, msg.Denom, msg.Sender, msg.Operator, msg.Protocol)
	}

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&token.MsgApproveForAll{
		Protocol: msg.Protocol,
		Denom:    msg.Denom,
		Operator: msg.Operator,
		Sender:   msg.Sender,
	})

	return
}

// BatchBurn implements token.MsgServer
// implement:
// - burnBatch
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L444
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L315
func (k Keeper) BatchBurn(goctx context.Context,
	msg *token.MsgBatchBurn,
) (res *token.MsgBatchBurnResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	switch msg.Protocol {
	case core.Protocol_NFT:
		err = k.batchBurnDDC721(ctx, msg.Denom, msg.TokenIDs, msg.Operator, msg.Protocol)
	case core.Protocol_MT:
		err = k.batchBurnDDC1155(ctx, msg.Denom, msg.TokenIDs, msg.Operator, msg.Protocol)
	}

	if err != nil {
		return nil, err
	}

	// TODO: redefine proto
	ctx.EventManager().EmitTypedEvent(&token.MsgBatchBurn{
		Protocol: msg.Protocol,
		Denom:    msg.Denom,
		TokenIDs: nil,
		Operator: msg.Operator,
	})

	return
}

// BatchTransfer implements token.MsgServer
// implement:
// - safeBatchTransferFrom
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L250
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L377
func (k Keeper) BatchTransfer(goctx context.Context,
	msg *token.MsgBatchTransfer,
) (res *token.MsgBatchTransferResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	switch msg.Protocol {
	case core.Protocol_NFT:
		err = k.batchTransferDDC721(ctx, msg.Denom, msg.TokenIDs, msg.From, msg.To, msg.Sender, msg.Protocol)
	case core.Protocol_MT:
		err = k.batchTransferDDC1155(ctx, msg.Denom, msg.TokenIDs, msg.Amount, msg.From, msg.To, msg.Sender, msg.Protocol)
	}

	// TODO: redefine protob
	ctx.EventManager().EmitTypedEvent(&token.MsgBatchTransfer{
		Protocol: msg.Protocol,
		Denom:    msg.Denom,
		TokenIDs: nil,
		Amount:   nil,
		From:     msg.From,
		To:       msg.To,
		Sender:   msg.Sender,
	})

	return
}

// Freeze implements token.MsgServer
// implement:
// - freeze
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L405
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L281
func (k Keeper) Freeze(goctx context.Context,
	msg *token.MsgFreeze,
) (res *token.MsgFreezeResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	switch msg.Protocol {
	case core.Protocol_NFT:
		err = k.freezeDDC721(ctx, msg.Denom, msg.TokenID, msg.Operator, msg.Protocol)
	case core.Protocol_MT:
		err = k.freezeDDC1155(ctx, msg.Denom, msg.TokenID, msg.Operator, msg.Protocol)
	}

	ctx.EventManager().EmitTypedEvent(&token.EventFreeze{
		Protocol: msg.Protocol,
		Denom:    msg.Denom,
		TokenID:  msg.TokenID,
		Operator: msg.Operator,
	})
	return
}

// Unfreeze implements token.MsgServer
// implement:
// - freeze
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L405
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L281
func (k Keeper) Unfreeze(goctx context.Context,
	msg *token.MsgUnfreeze,
) (res *token.MsgUnfreezeResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	switch msg.Protocol {
	case core.Protocol_NFT:
		err = k.unfreezeDDC721(ctx, msg.Denom, msg.TokenID, msg.Operator, msg.Protocol)
	case core.Protocol_MT:
		err = k.unfreezeDDC1155(ctx, msg.Denom, msg.TokenID, msg.Operator, msg.Protocol)
	}

	ctx.EventManager().EmitTypedEvent(&token.EventUnfreeze{
		Protocol: msg.Protocol,
		Denom:    msg.Denom,
		TokenID:  msg.TokenID,
		Operator: msg.Operator,
	})
	return
}
