package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/exported"
)

var _ exported.Hook = Keeper{}

// BeforeDenomTransfer implements exported.Hook
func (Keeper) BeforeDenomTransfer(ctx sdk.Context,
	protocol string,
	denomID string,
	sender sdk.AccAddress,
) error {
	//do nothing
	return nil
}

// BeforeTokenBurn implements exported.Hook
func (k Keeper) BeforeTokenBurn(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
) error {
	proto := core.Protocol_value[protocol]
	return k.pay(ctx, sender.String(), core.Protocol(proto), denomID, core.Function_EDIT)
}

// BeforeTokenEdit implements exported.Hook
func (k Keeper) BeforeTokenEdit(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
) error {
	//do nothing
	return nil
}

// BeforeTokenMint implements exported.Hook
func (k Keeper) BeforeTokenMint(ctx sdk.Context,
	protocol string,
	denomID string,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
) error {
	proto := core.Protocol_value[protocol]
	return k.pay(ctx, sender.String(), core.Protocol(proto), denomID, core.Function_MINT)
}

// BeforeTokenTransfer implements exported.Hook
func (k Keeper) BeforeTokenTransfer(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
) error {
	proto := core.Protocol_value[protocol]
	return k.pay(ctx, sender.String(), core.Protocol(proto), denomID, core.Function_TRANSFER)
}
