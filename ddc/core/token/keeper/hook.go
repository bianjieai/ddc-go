package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/export"
)

var _ export.Hook = Keeper{}

// BeforeDenomTransfer implements export.Hook
func (Keeper) BeforeDenomTransfer(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress) error {
	panic("unimplemented")
}

// BeforeTokenBurn implements export.Hook
func (Keeper) BeforeTokenBurn(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	panic("unimplemented")
}

// BeforeTokenEdit implements export.Hook
func (Keeper) BeforeTokenEdit(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	panic("unimplemented")
}

// BeforeTokenMint implements export.Hook
func (Keeper) BeforeTokenMint(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	panic("unimplemented")
}

// BeforeTokenTransfer implements export.Hook
func (Keeper) BeforeTokenTransfer(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	panic("unimplemented")
}
