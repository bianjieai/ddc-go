package keeper

import (
	"context"

	"github.com/bianjieai/ddc-go/ddc/core/fee"
)

var _ fee.MsgServer = Keeper{}

// DeleteFeeRule implements fee.MsgServer
func (Keeper) DeleteFeeRule(context.Context, *fee.MsgDeleteFeeRule) (*fee.MsgDeleteFeeRuleResponse, error) {
	panic("unimplemented")
}

// Recharge implements fee.MsgServer
func (Keeper) Recharge(context.Context, *fee.MsgRecharge) (*fee.MsgRechargeResponse, error) {
	panic("unimplemented")
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
func (Keeper) SetFeeRule(context.Context, *fee.MsgSetFeeRule) (*fee.MsgSetFeeRuleResponse, error) {
	panic("unimplemented")
}

// Settlement implements fee.MsgServer
func (Keeper) Settlement(context.Context, *fee.MsgSettlement) (*fee.MsgSettlementResponse, error) {
	panic("unimplemented")
}
