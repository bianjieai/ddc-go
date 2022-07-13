package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

func (k Keeper) funcStore(ctx sdk.Context, role core.Role, protocol core.Protocol, denom string) prefix.Store {
	store := k.prefixStore(ctx)
	return prefix.NewStore(store, prefixRoleAndFunBindKey(role, protocol, denom))
}

// addFunction save function by role
func (k Keeper) addFunction(ctx sdk.Context,
	role core.Role,
	protocol core.Protocol,
	denom string,
	function core.Function,
) error {
	store := k.funcStore(ctx, role, protocol, denom)
	key := funKey(function)
	if store.Has(key) {
		return sdkerrors.Wrapf(auth.ErrFunctionHasExist, "function: %s has exist", function.String())
	}
	store.Set(key, key)
	return nil
}

// deleteFunction delete function by role
func (k Keeper) deleteFunction(ctx sdk.Context,
	role core.Role,
	protocol core.Protocol,
	denom string,
	function core.Function,
) error {
	store := k.funcStore(ctx, role, protocol, denom)
	key := funKey(function)
	if !store.Has(key) {
		return sdkerrors.Wrapf(auth.ErrFunctionNotExist, "function: %s not exist", function.String())
	}
	store.Delete(key)
	return nil
}

func (k Keeper) getFunction(ctx sdk.Context,
	role core.Role,
	protocol core.Protocol,
	denom string,
) (fun []core.Function) {
	iterator := sdk.KVStorePrefixIterator(k.prefixStore(ctx), prefixRoleAndFunBindKey(role, protocol, denom))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		fun = append(fun, core.Function(sdk.BigEndianToUint64(iterator.Value())))
	}
	return
}

func (k Keeper) hasFunction(ctx sdk.Context,
	role core.Role,
	protocol core.Protocol,
	denom string,
	function core.Function,
) bool {
	store := k.funcStore(ctx, role, protocol, denom)
	key := funKey(function)
	return store.Has(key)
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L422
func (k Keeper) hasFunctionPermission(ctx sdk.Context,
	address string,
	protocol core.Protocol,
	denom string,
	function core.Function,
) error {
	account, err := k.GetAccount(ctx, address)
	if err != nil {
		return err
	}
	if k.isActive(account) {
		return sdkerrors.Wrapf(auth.ErrAccountNotActive, "account: %s is not active", address)
	}

	if k.hasFunction(ctx, account.Role, protocol, denom, function) {
		return sdkerrors.Wrapf(auth.ErrFunctionNotExist, "function: %s not exist", function.String())
	}
	return nil
}
