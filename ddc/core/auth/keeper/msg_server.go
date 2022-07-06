package keeper

import (
	context "context"

	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

var _ auth.MsgServer = Keeper{}

// AddAccount implements auth.MsgServer
// implement:
// 	- addOperator
// 	- addAccountByPlatform
// 	- addAccountByOperator
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L58
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L81
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L158
func (Keeper) AddAccount(context.Context, *auth.MsgAddAccount) (*auth.MsgAddAccountResponse, error) {
	panic("unimplemented")
}

// AddBatchAccount implements auth.MsgServer
// implement:
// 	- addBatchAccountByPlatform
// 	- addBatchAccountByOperator
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L103
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L172
func (Keeper) AddBatchAccount(context.Context, *auth.MsgAddBatchAccount) (*auth.MsgAddBatchAccountResponse, error) {
	panic("unimplemented")
}

// AddFunction implements auth.MsgServer
// implement:
// 	- addFunction
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L317
func (Keeper) AddFunction(context.Context, *auth.MsgAddFunction) (*auth.MsgAddFunctionResponse, error) {
	panic("unimplemented")
}

// ApproveCrossPlatform implements auth.MsgServer
// implement:
// 	- crossPlatformApproval
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L373
func (Keeper) ApproveCrossPlatform(context.Context, *auth.MsgApproveCrossPlatform) (*auth.MsgApproveCrossPlatformResponse, error) {
	panic("unimplemented")
}

// DeleteAccount implements auth.MsgServer
func (Keeper) DeleteAccount(context.Context, *auth.MsgDeleteAccount) (*auth.MsgDeleteAccountResponse, error) {
	//TODO
	return &auth.MsgDeleteAccountResponse{}, nil
}

// DeleteFunction implements auth.MsgServer
// implement:
// 	- delFunction
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L352
func (Keeper) DeleteFunction(context.Context, *auth.MsgDeleteFunction) (*auth.MsgDeleteFunctionResponse, error) {
	panic("unimplemented")
}

// SyncPlatformDID implements auth.MsgServer
// implement:
// 	- syncPlatformDID
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L146
func (Keeper) SyncPlatformDID(context.Context, *auth.MsgSyncPlatformDID) (*auth.MsgSyncPlatformDIDResponse, error) {
	panic("unimplemented")
}

// UpdateAccountState implements auth.MsgServer
// implement:
// 	- updateAccountState
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L199
func (Keeper) UpdateAccountState(context.Context, *auth.MsgUpdateAccountState) (*auth.MsgUpdateAccountStateResponse, error) {
	panic("unimplemented")
}

// UpgradeToDDC implements auth.MsgServer
func (Keeper) UpgradeToDDC(context.Context, *auth.MsgUpgradeToDDC) (*auth.MsgUpgradeToDDCResponse, error) {
	panic("unimplemented")
}
