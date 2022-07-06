package keeper

import (
	context "context"

	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

var _ auth.MsgServer = Keeper{}

// AddAccount implements auth.MsgServer
func (Keeper) AddAccount(context.Context, *auth.MsgAddAccount) (*auth.MsgAddAccountResponse, error) {
	panic("unimplemented")
}

// AddBatchAccount implements auth.MsgServer
func (Keeper) AddBatchAccount(context.Context, *auth.MsgAddBatchAccount) (*auth.MsgAddBatchAccountResponse, error) {
	panic("unimplemented")
}

// AddFunction implements auth.MsgServer
func (Keeper) AddFunction(context.Context, *auth.MsgAddFunction) (*auth.MsgAddFunctionResponse, error) {
	panic("unimplemented")
}

// ApproveCrossPlatform implements auth.MsgServer
func (Keeper) ApproveCrossPlatform(context.Context, *auth.MsgApproveCrossPlatform) (*auth.MsgApproveCrossPlatformResponse, error) {
	panic("unimplemented")
}

// DeleteAccount implements auth.MsgServer
func (Keeper) DeleteAccount(context.Context, *auth.MsgDeleteAccount) (*auth.MsgDeleteAccountResponse, error) {
	panic("unimplemented")
}

// DeleteFunction implements auth.MsgServer
func (Keeper) DeleteFunction(context.Context, *auth.MsgDeleteFunction) (*auth.MsgDeleteFunctionResponse, error) {
	panic("unimplemented")
}

// SyncPlatformDID implements auth.MsgServer
func (Keeper) SyncPlatformDID(context.Context, *auth.MsgSyncPlatformDID) (*auth.MsgSyncPlatformDIDResponse, error) {
	panic("unimplemented")
}

// UpdateAccountState implements auth.MsgServer
func (Keeper) UpdateAccountState(context.Context, *auth.MsgUpdateAccountState) (*auth.MsgUpdateAccountStateResponse, error) {
	panic("unimplemented")
}

// UpgradeToDDC implements auth.MsgServer
func (Keeper) UpgradeToDDC(context.Context, *auth.MsgUpgradeToDDC) (*auth.MsgUpgradeToDDCResponse, error) {
	panic("unimplemented")
}
