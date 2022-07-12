package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L58
func (k Keeper) recharge(ctx sdk.Context, from, to string, amount uint64) error {
	if len(from) != 0 {
		balance := k.balanceOf(ctx, from)
		if balance < amount {
			return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "account's balance %d < %s recharged", balance, amount)
		}
		k.setBalance(ctx, from, balance-amount)

	}
	balance := k.balanceOf(ctx, to)
	k.setBalance(ctx, to, balance+amount)
	return nil
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
