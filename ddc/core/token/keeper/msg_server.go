package keeper

import (
	context "context"

	"github.com/bianjieai/ddc-go/ddc/core/token"
)

var _ token.MsgServer = Keeper{}

// Approve implements token.MsgServer
func (Keeper) Approve(context.Context, *token.MsgApprove) (*token.MsgApproveResponse, error) {
	panic("unimplemented")
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
