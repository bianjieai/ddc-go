package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/core/fee"
)

var _ fee.MsgServer = Keeper{}

// DeleteFeeRule implements fee.MsgServer
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L154
func (k Keeper) DeleteFeeRule(goctx context.Context, msg *fee.MsgDeleteFeeRule) (res *fee.MsgDeleteFeeRuleResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if err = k.requireOperator(ctx, msg.Operator); err != nil {
		return res, err
	}
	k.deleteFee(ctx, msg.Protocol, msg.Denom, msg.Function)
	return
}

// Recharge implements fee.MsgServer
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L72
func (k Keeper) Recharge(goctx context.Context, msg *fee.MsgRecharge) (res *fee.MsgRechargeResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if err = k.checkRechargeAuth(ctx, msg.From, msg.To, msg.Amount); err != nil {
		return
	}
	err = k.recharge(ctx, msg.From, msg.To, msg.Amount)
	return
}

// RechargeBatch implements fee.MsgServer
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L81
func (k Keeper) RechargeBatch(goctx context.Context, msg *fee.MsgRechargeBatch) (res *fee.MsgRechargeBatchResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	for i := range msg.To {
		if err = k.checkRechargeAuth(ctx, msg.From, msg.To[i], msg.Amount[i]); err != nil {
			return res, err
		}
		if err = k.recharge(ctx, msg.From, msg.To[i], msg.Amount[i]); err != nil {
			return res, err
		}
	}
	return
}

// RevokeDDC implements fee.MsgServer
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L164
func (k Keeper) RevokeDDC(goctx context.Context, msg *fee.MsgRevokeDDC) (res *fee.MsgRevokeDDCResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if err = k.requireOperator(ctx, msg.Operator); err != nil {
		return res, err
	}

	store := k.prefixStore(ctx)
	store.Delete(ddcAuthKey(msg.Protocol, msg.Denom))
	return
}

// SetFeeRule implements fee.MsgServer
// implements: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L139
func (k Keeper) SetFeeRule(goctx context.Context, msg *fee.MsgSetFeeRule) (res *fee.MsgSetFeeRuleResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if err = k.requireOperator(ctx, msg.Operator); err != nil {
		return res, err
	}
	k.setFee(ctx, msg.Protocol, msg.Denom, msg.Function, uint64(msg.Fee))
	return
}

// Settlement implements fee.MsgServer
func (k Keeper) Settlement(goctx context.Context, msg *fee.MsgSettlement) (res *fee.MsgSettlementResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if err = k.settlement(ctx, msg.Operator, msg.Protocol, msg.Denom, msg.Amount); err != nil {
		return res, err
	}
	return
}
