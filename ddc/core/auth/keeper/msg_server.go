package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/ddc-go/ddc/core"
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
func (k Keeper) AddAccount(goctx context.Context, msg *auth.MsgAddAccount) (*auth.MsgAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if k.isRoot(ctx, msg.Sender) {
		return &auth.MsgAddAccountResponse{}, k.addOperator(ctx, msg.Address, msg.Name, msg.Did)
	}

	account, err := k.GetAccount(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	switch account.Role {
	case core.Role_OPERATOR:
		return &auth.MsgAddAccountResponse{}, k.addAccountByOperator(ctx,
			msg.Address, msg.Name, msg.Did, msg.LeaderDID, account)
	case core.Role_PLATFORM_MANAGER:
		return &auth.MsgAddAccountResponse{}, k.addAccountByPlatform(ctx,
			msg.Address, msg.Name, msg.Did, account)
	default:
		return &auth.MsgAddAccountResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid operate")
	}
}

// AddBatchAccount implements auth.MsgServer
// implement:
// 	- addBatchAccountByPlatform
// 	- addBatchAccountByOperator
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L103
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L172
func (k Keeper) AddBatchAccount(goctx context.Context, msg *auth.MsgAddBatchAccount) (*auth.MsgAddBatchAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	account, err := k.GetAccount(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	switch account.Role {
	case core.Role_OPERATOR:
		return &auth.MsgAddBatchAccountResponse{}, k.addBatchAccountByOperator(ctx,
			msg.Addresses, msg.Names, msg.Dids, msg.LeaderDIDs, account)
	case core.Role_PLATFORM_MANAGER:
		return &auth.MsgAddBatchAccountResponse{}, k.addBatchAccountByPlatform(ctx,
			msg.Addresses, msg.Names, msg.Dids, account)
	default:
		return &auth.MsgAddBatchAccountResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid operate")
	}
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

// SetSwitcherStateOfPlatform implements auth.MsgServer
// implement:
// 	- setSwitcherStateOfPlatform
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L133
func (Keeper) SetSwitcherStateOfPlatform(context.Context, *auth.MsgSetSwitcherStateOfPlatform) (*auth.MsgSetSwitcherStateOfPlatformResponse, error) {
	panic("unimplemented")
}
