package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/exported"
)

var _ exported.Hook = Keeper{}

// BeforeDenomTransfer implements exported.Hook
func (Keeper) BeforeDenomTransfer(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress) error {
	panic("unimplemented")
}

// BeforeTokenBurn implements exported.Hook
func (Keeper) BeforeTokenBurn(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	panic("unimplemented")
}

// BeforeTokenEdit implements exported.Hook
func (Keeper) BeforeTokenEdit(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	panic("unimplemented")
}

// BeforeTokenMint implements exported.Hook
func (Keeper) BeforeTokenMint(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	panic("unimplemented")
}

// BeforeTokenTransfer implements exported.Hook
func (Keeper) BeforeTokenTransfer(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	panic("unimplemented")
}
