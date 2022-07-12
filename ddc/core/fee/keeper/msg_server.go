package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/fee"
)

var _ fee.MsgServer = Keeper{}

// DeleteFeeRule implements fee.MsgServer
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L154
func (k Keeper) DeleteFeeRule(goctx context.Context, msg *fee.MsgDeleteFeeRule) (res *fee.MsgDeleteFeeRuleResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if err = k.requireOperator(ctx, msg.Operator); err != nil {
		return
	}
	store := k.prefixStore(ctx)
	store.Delete(feeRuleKey(msg.Protocol, msg.Denom, msg.Function))
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
func (Keeper) RechargeBatch(context.Context, *fee.MsgRechargeBatch) (*fee.MsgRechargeBatchResponse, error) {
	panic("unimplemented")
}

// RevokeDDC implements fee.MsgServer
func (Keeper) RevokeDDC(context.Context, *fee.MsgRevokeDDC) (*fee.MsgRevokeDDCResponse, error) {
	panic("unimplemented")
}

// SetFeeRule implements fee.MsgServer
// implements: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L139
func (k Keeper) SetFeeRule(goctx context.Context, msg *fee.MsgSetFeeRule) (res *fee.MsgSetFeeRuleResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if err = k.requireOperator(ctx, msg.Operator); err != nil {
		return
	}
	store := k.prefixStore(ctx)
	store.Set(feeRuleKey(msg.Protocol, msg.Denom, msg.Function), sdk.Uint64ToBigEndian(uint64(msg.Fee)))
	return
}

// Settlement implements fee.MsgServer
func (Keeper) Settlement(context.Context, *fee.MsgSettlement) (*fee.MsgSettlementResponse, error) {
	panic("unimplemented")
}

func (k Keeper) requireOperator(ctx sdk.Context, sender string) error {
	return k.authKeeper.CheckAvailableAndRole(ctx, sender, core.Role_OPERATOR)
}
