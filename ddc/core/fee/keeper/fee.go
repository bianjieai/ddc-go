package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/fee"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L58
func (k Keeper) recharge(ctx sdk.Context, payer, receiver string, amount uint64) error {
	if len(payer) != 0 {
		balance := k.balanceOf(ctx, payer)
		if balance < amount {
			return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "account's balance %d < %s recharged", balance, amount)
		}
		k.setBalance(ctx, payer, balance-amount)
	}
	balance := k.balanceOf(ctx, receiver)
	k.setBalance(ctx, receiver, balance+amount)
	return nil
}

//implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L106
func (k Keeper) pay(ctx sdk.Context,
	payer string,
	protocol core.Protocol,
	denom string,
	function core.Function,
) error {
	fee, err := k.queryFee(ctx, protocol, denom, function)
	if err != nil {
		return err
	}
	if fee == 0 {
		return nil
	}
	//TODO
	ddcAddress := GetDDCEscrowAddress(protocol, denom).String()
	return k.recharge(ctx, payer, ddcAddress, fee)
}

//implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L125
func (k Keeper) settlement(ctx sdk.Context,
	operator string,
	protocol core.Protocol,
	denom string,
	amount uint64,
) error {
	store := k.prefixStore(ctx)
	if store.Has(ddcAuthKey(protocol, denom)) {
		return sdkerrors.Wrapf(fee.ErrFeeRuleUnavailable, "denom: %s", denom)
	}
	if err := k.requireOperator(ctx, operator); err != nil {
		return err
	}

	ddcAddress := GetDDCEscrowAddress(protocol, denom).String()
	return k.recharge(ctx, ddcAddress, operator, amount)
}

//implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L139
func (k Keeper) setFee(ctx sdk.Context,
	protocol core.Protocol,
	denom string,
	function core.Function,
	fee uint64,
) {
	store := k.prefixStore(ctx)
	store.Set(feeRuleKey(protocol, denom, function), sdk.Uint64ToBigEndian(fee))

	ddcAuthKey := ddcAuthKey(protocol, denom)
	if !store.Has(ddcAuthKey) {
		store.Set(ddcAuthKey, Placeholder)
	}
}

//implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L154
func (k Keeper) deleteFee(ctx sdk.Context,
	protocol core.Protocol,
	denom string,
	function core.Function,
) {
	store := k.prefixStore(ctx)
	store.Delete(feeRuleKey(protocol, denom, function))
}

//implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Charge/Charge.sol#L198
func (k Keeper) queryFee(ctx sdk.Context,
	protocol core.Protocol,
	denom string,
	function core.Function,
) (uint64, error) {
	store := k.prefixStore(ctx)
	if store.Has(ddcAuthKey(protocol, denom)) {
		return 0, sdkerrors.Wrapf(fee.ErrFeeRuleUnavailable, "denom:%s", denom)
	}

	bz := store.Get(feeRuleKey(protocol, denom, function))
	return sdk.BigEndianToUint64(bz), nil
}

func (k Keeper) balanceOf(ctx sdk.Context, address string) uint64 {
	store := k.prefixStore(ctx)
	bz := store.Get(balanceKey(address))
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) setBalance(ctx sdk.Context, address string, amount uint64) {
	store := k.prefixStore(ctx)
	store.Set(balanceKey(address), sdk.Uint64ToBigEndian(amount))
}
