package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/exported"
)

var _ exported.Hook = Keeper{}

// BeforeDenomTransfer implements exported.Hook
func (k Keeper) BeforeDenomTransfer(ctx sdk.Context,
	protocol string,
	denomID string,
	sender sdk.AccAddress,
) error {
	if !k.requireDDC(ctx, protocol, denomID) {
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrNotSupported, "ddc denom not support to transfer")
}

// BeforeTokenBurn implements exported.Hook
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L429
func (k Keeper) BeforeTokenBurn(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
) error {
	if !k.requireDDC(ctx, protocol, denomID) {
		return nil
	}
	if err := k.requireSenderHasFuncPermission(ctx,
		sender.String(), protocol, denomID, core.Function_BURN); err != nil {
		return err
	}
	// TODO
	// _requireExistsAndApproved check by token Hook
	return nil
}

// BeforeTokenEdit implements exported.Hook
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L208-L209
func (k Keeper) BeforeTokenEdit(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
) error {
	if !k.requireDDC(ctx, protocol, denomID) {
		return nil
	}
	if err := k.requireSenderHasFuncPermission(ctx,
		sender.String(), protocol, denomID, core.Function_EDIT); err != nil {
		return err
	}
	// TODO
	// _requireAvailableDDC check by token Hook
	return nil
}

// BeforeTokenMint implements exported.Hook
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L143-L145
func (k Keeper) BeforeTokenMint(ctx sdk.Context,
	protocol string,
	denomID string,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
) error {
	if !k.requireDDC(ctx, protocol, denomID) {
		return nil
	}

	if err := k.requireSenderHasFuncPermission(ctx, sender.String(),
		protocol, denomID, core.Function_MINT); err != nil {
		return err
	}

	if err := k.requireAvailableDDCAccount(ctx, receiver.String()); err != nil {
		return err
	}

	if err := k.requireOnePlatform(ctx, sender.String(), receiver.String()); err != nil {
		return err
	}
	return nil
}

// BeforeTokenTransfer implements exported.Hook
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L366-L368
func (k Keeper) BeforeTokenTransfer(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
) error {
	if !k.requireDDC(ctx, protocol, denomID) {
		return nil
	}

	if err := k.requireTransferConstraints_FistStep(ctx,
		protocol, denomID, tokenID, sender.String(), receiver.String()); err != nil {
		return err
	}
	// TODO
	// _requireTransferConstraints_SecondStep check by token Hook
	return nil
}
