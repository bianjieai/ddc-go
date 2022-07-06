package keeper

import (
	context "context"

	"github.com/bianjieai/ddc-go/ddc/core/fee"
)

var _ fee.QueryServer = Keeper{}

// BalanceOf implements fee.QueryServer
func (Keeper) BalanceOf(context.Context, *fee.QueryBalanceOfRequest) (*fee.QueryBalanceOfResponse, error) {
	panic("unimplemented")
}

// BalanceOfBatch implements fee.QueryServer
func (Keeper) BalanceOfBatch(context.Context, *fee.QueryBalanceOfBatchRequest) (*fee.QueryBalanceOfBatchResponse, error) {
	panic("unimplemented")
}

// FeeRule implements fee.QueryServer
func (Keeper) FeeRule(context.Context, *fee.QueryFeeRuleRequest) (*fee.QueryFeeRuleResponse, error) {
	panic("unimplemented")
}

// TotalSupply implements fee.QueryServer
func (Keeper) TotalSupply(context.Context, *fee.QueryTotalSupplyRequest) (*fee.QueryTotalSupplyResponse, error) {
	panic("unimplemented")
}
