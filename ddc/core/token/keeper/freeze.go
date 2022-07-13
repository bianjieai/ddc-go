package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L405
func (k Keeper) freezeDDC721(ctx sdk.Context,
	denomID string,
	tokenID string,
	operator string,
	protocol core.Protocol,
) error {
	// TODO
	// if !requireSenderHasFuncPermission()
	// if !requireOperator()

	if err := k.requireDisabledDDC(ctx, core.Protocol_name[int32(protocol)], denomID, tokenID); err != nil {
		return err
	}

	k.setTokenBlocklist(ctx, protocol, denomID, tokenID)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L281
func (k Keeper) freezeDDC1155(ctx sdk.Context,
	denomID string,
	tokenID string,
	operator string,
	protocol core.Protocol,
) error {
	// TODO
	// if !requireSenderHasFuncPermission()
	// if !requireOperator()

	if err := k.requireAvailableDDC(ctx, core.Protocol_name[int32(protocol)], denomID, tokenID); err != nil {
		return err
	}

	k.setTokenBlocklist(ctx, protocol, denomID, tokenID)

	return nil
}

// https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L420
func (k Keeper) unfreezeDDC721(ctx sdk.Context,
	denomID string,
	tokenID string,
	operator string,
	protocol core.Protocol,
) error {
	// TODO
	// if !requireSenderHasFuncPermission()
	// if !requireOperator()
	if err := k.requireDisabledDDC(ctx, core.Protocol_name[int32(protocol)], denomID, tokenID); err != nil {
		return err
	}

	if !k.isInBlocklist(ctx, protocol, denomID, tokenID) {
		return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is not in blocklist")
	}

	k.unsetTokenBlocklist(ctx, protocol, denomID, tokenID)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L293
func (k Keeper) unfreezeDDC1155(ctx sdk.Context,
	denomID string,
	tokenID string,
	operator string,
	protocol core.Protocol,
) error {
	// TODO
	// if !requireSenderHasFuncPermission()
	// if !requireOperator()

	if err := k.requireDisabledDDC(ctx, core.Protocol_name[int32(protocol)], denomID, tokenID); err != nil {
		return err
	}

	k.unsetTokenBlocklist(ctx, protocol, denomID, tokenID)

	return nil
}

func (k Keeper) isInBlocklist(ctx sdk.Context,
	protocol core.Protocol,
	denomID string,
	tokenID string,
) bool {
	store := k.prefixStore(ctx)
	return store.Has(tokenBlocklistKey(protocol, denomID, tokenID))
}

func (k Keeper) setTokenBlocklist(ctx sdk.Context,
	protocol core.Protocol,
	denomID string,
	tokenID string,
) {
	store := k.prefixStore(ctx)
	key := tokenBlocklistKey(protocol, denomID, tokenID)
	if !store.Has(key) {
		store.Set(key, Placeholder)
	}
}

func (k Keeper) unsetTokenBlocklist(ctx sdk.Context,
	protocol core.Protocol,
	denomID string,
	tokenID string,
) {
	store := k.prefixStore(ctx)
	key := tokenBlocklistKey(protocol, denomID, tokenID)
	if store.Has(key) {
		store.Delete(key)
	}
}
