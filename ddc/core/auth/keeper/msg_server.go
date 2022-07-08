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
func (k Keeper) AddAccount(goctx context.Context, msg *auth.MsgAddAccount) (res *auth.MsgAddAccountResponse, err error) {
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
		err = k.addAccountByOperator(ctx,
			msg.Address, msg.Name, msg.Did, msg.LeaderDID, account)
	case core.Role_PLATFORM_MANAGER:
		err = k.addAccountByPlatform(ctx,
			msg.Address, msg.Name, msg.Did, account)
	default:
		return &auth.MsgAddAccountResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid operate")
	}
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&auth.EventAddAccount{
		Caller:  msg.Sender,
		Account: msg.Address,
	})
	return
}

// AddBatchAccount implements auth.MsgServer
// implement:
// 	- addBatchAccountByPlatform
// 	- addBatchAccountByOperator
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L103
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L172
func (k Keeper) AddBatchAccount(goctx context.Context, msg *auth.MsgAddBatchAccount) (res *auth.MsgAddBatchAccountResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	account, err := k.GetAccount(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	switch account.Role {
	case core.Role_OPERATOR:
		err = k.addBatchAccountByOperator(ctx,
			msg.Addresses, msg.Names, msg.Dids, msg.LeaderDIDs, account)
	case core.Role_PLATFORM_MANAGER:
		err = k.addBatchAccountByPlatform(ctx,
			msg.Addresses, msg.Names, msg.Dids, account)
	default:
		return &auth.MsgAddBatchAccountResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid operate")
	}
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&auth.EventAddBatchAccount{
		Caller:  msg.Sender,
		Address: msg.Addresses,
	})
	return
}

// AddFunction implements auth.MsgServer
// implement:
// 	- addFunction
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L317
func (k Keeper) AddFunction(goctx context.Context, msg *auth.MsgAddFunction) (res *auth.MsgAddFunctionResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	account, err := k.GetAccount(ctx, msg.Operator)
	if err != nil {
		return nil, err
	}
	if account.Role != core.Role_OPERATOR {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "account: %s no access", msg.Operator)
	}

	if err = k.addFunction(ctx, msg.Role, msg.Protocol, msg.Denom, msg.Function); err != nil {
		return res, err
	}
	ctx.EventManager().EmitTypedEvent(&auth.EventAddFunction{
		Operator: msg.Operator,
		Role:     msg.Role,
		Protocol: msg.Protocol,
		Denom:    msg.Denom,
		Function: msg.Function,
	})
	return
}

// ApproveCrossPlatform implements auth.MsgServer
// implement:
// 	- crossPlatformApproval
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L373
func (k Keeper) ApproveCrossPlatform(goctx context.Context, msg *auth.MsgApproveCrossPlatform) (res *auth.MsgApproveCrossPlatformResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	account, err := k.GetAccount(ctx, msg.Operator)
	if err != nil {
		return nil, err
	}
	if account.Role != core.Role_OPERATOR {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "account: %s no access", msg.Operator)
	}

	if err := k.approveCrossPlatform(ctx, msg.From, msg.To); err != nil {
		return nil, err
	}
	return
}

// DeleteAccount implements auth.MsgServer
func (k Keeper) DeleteAccount(goctx context.Context, msg *auth.MsgDeleteAccount) (*auth.MsgDeleteAccountResponse, error) {
	//TODO
	panic("unimplemented")
}

// DeleteFunction implements auth.MsgServer
// implement:
// 	- delFunction
// reference:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L352
func (k Keeper) DeleteFunction(goctx context.Context, msg *auth.MsgDeleteFunction) (res *auth.MsgDeleteFunctionResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	account, err := k.GetAccount(ctx, msg.Operator)
	if err != nil {
		return nil, err
	}
	if account.Role != core.Role_OPERATOR {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "account: %s no access", msg.Operator)
	}

	if err = k.deleteFunction(ctx, msg.Role, msg.Protocol, msg.Denom, msg.Function); err != nil {
		return res, err
	}
	ctx.EventManager().EmitTypedEvent(&auth.EventDeleteFunction{
		Operator: msg.Operator,
		Role:     msg.Role,
		Protocol: msg.Protocol,
		Denom:    msg.Denom,
		Function: msg.Function,
	})
	return
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
