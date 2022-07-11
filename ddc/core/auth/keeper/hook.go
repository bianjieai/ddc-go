package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/exported"
)

var _ exported.Hook = Keeper{}

// BeforeDenomTransfer implements exported.Hook
func (k Keeper) BeforeDenomTransfer(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress) error {
	if !k.ControlByDDC(ctx, protocol, denomID) {
		return nil
	}
	//TODO
	return nil
}

// BeforeTokenBurn implements exported.Hook
func (k Keeper) BeforeTokenBurn(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	if !k.ControlByDDC(ctx, protocol, denomID) {
		return nil
	}
	//TODO
	return nil
}

// BeforeTokenEdit implements exported.Hook
func (k Keeper) BeforeTokenEdit(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	if !k.ControlByDDC(ctx, protocol, denomID) {
		return nil
	}
	//TODO
	return nil
}

// BeforeTokenMint implements exported.Hook
func (k Keeper) BeforeTokenMint(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	if !k.ControlByDDC(ctx, protocol, denomID) {
		return nil
	}
	//TODO
	return nil
}

// BeforeTokenTransfer implements exported.Hook
func (k Keeper) BeforeTokenTransfer(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	if !k.ControlByDDC(ctx, protocol, denomID) {
		return nil
	}
	//TODO
	return nil
}
