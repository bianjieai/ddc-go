package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/exported"
)

var _ exported.Hook = Keeper{}

// BeforeDenomTransfer implements exported.Hook
func (k Keeper) BeforeDenomTransfer(ctx sdk.Context,
	protocol string,
	denomID string,
	sender sdk.AccAddress,
) error {
	return nil
}

// BeforeTokenBurn implements exported.Hook
// implement:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L430
func (k Keeper) BeforeTokenBurn(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
) error {
	if err := k.requireExistsAndApproved(ctx, protocol, denomID, tokenID, sender.String()); err != nil {
		return nil
	}

	return nil
}

// BeforeTokenEdit implements exported.Hook
// implement:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L209-L210
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L176-L177
func (k Keeper) BeforeTokenEdit(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
) error {
	if err := k.requireAvailableDDC(ctx, protocol, denomID, tokenID); err != nil {
		return err
	}

	if err := k.requireApprovedOrOwner(ctx, protocol, denomID, tokenID, sender.String()); err != nil {
		return err
	}

	return nil
}

// BeforeTokenMint implements exported.Hook
func (Keeper) BeforeTokenMint(ctx sdk.Context,
	protocol string,
	denomID string,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
) error {
	return nil
}

// BeforeTokenTransfer implements exported.Hook
// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L368
func (k Keeper) BeforeTokenTransfer(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
	receiver sdk.AccAddress) error {
	if err := k.requireTransferConstraintsSecondStep(ctx, protocol, denomID, tokenID, sender.String(), receiver.String()); err != nil {
		return err
	}
	return nil
}
