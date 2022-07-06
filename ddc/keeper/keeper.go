package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	authkeeper "github.com/bianjieai/ddc-go/ddc/core/auth/keeper"
	feekeeper "github.com/bianjieai/ddc-go/ddc/core/fee/keeper"
	tokenkeeper "github.com/bianjieai/ddc-go/ddc/core/token/keeper"
	"github.com/bianjieai/ddc-go/ddc/export"
)

var _ export.Hook = Keeper{}

// Keeper of the auth store
type Keeper struct {
	AuthKeeper  authkeeper.Keeper
	FeeKeeper   feekeeper.Keeper
	TokenKeeper tokenkeeper.Keeper
}

// NewKeeper creates a new nft Keeper instance
func NewKeeper(cdc codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		AuthKeeper:  authkeeper.NewKeeper(cdc, key),
		FeeKeeper:   feekeeper.NewKeeper(cdc, key),
		TokenKeeper: tokenkeeper.NewKeeper(cdc, key),
	}
}

// BeforeDenomTransfer implements export.Hook
func (k Keeper) BeforeDenomTransfer(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress) error {
	if err := k.AuthKeeper.BeforeDenomTransfer(ctx, protocol, denomID, sender); err != nil {
		return err
	}

	if err := k.TokenKeeper.BeforeDenomTransfer(ctx, protocol, denomID, sender); err != nil {
		return err
	}

	if err := k.FeeKeeper.BeforeDenomTransfer(ctx, protocol, denomID, sender); err != nil {
		return err
	}
	return nil
}

// BeforeTokenBurn implements export.Hook
func (k Keeper) BeforeTokenBurn(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	if err := k.AuthKeeper.BeforeTokenBurn(ctx, protocol, denomID, tokenID, sender); err != nil {
		return err
	}

	if err := k.TokenKeeper.BeforeTokenBurn(ctx, protocol, denomID, tokenID, sender); err != nil {
		return err
	}

	if err := k.FeeKeeper.BeforeTokenBurn(ctx, protocol, denomID, tokenID, sender); err != nil {
		return err
	}
	return nil
}

// BeforeTokenEdit implements export.Hook
func (k Keeper) BeforeTokenEdit(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	if err := k.AuthKeeper.BeforeTokenEdit(ctx, protocol, denomID, tokenID, sender); err != nil {
		return err
	}

	if err := k.TokenKeeper.BeforeTokenEdit(ctx, protocol, denomID, tokenID, sender); err != nil {
		return err
	}

	if err := k.FeeKeeper.BeforeTokenEdit(ctx, protocol, denomID, tokenID, sender); err != nil {
		return err
	}
	return nil
}

// BeforeTokenMint implements export.Hook
func (k Keeper) BeforeTokenMint(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	if err := k.AuthKeeper.BeforeTokenMint(ctx, protocol, denomID, sender, receiver); err != nil {
		return err
	}

	if err := k.TokenKeeper.BeforeTokenMint(ctx, protocol, denomID, sender, receiver); err != nil {
		return err
	}

	if err := k.FeeKeeper.BeforeTokenMint(ctx, protocol, denomID, sender, receiver); err != nil {
		return err
	}
	return nil
}

// BeforeTokenTransfer implements export.Hook
func (k Keeper) BeforeTokenTransfer(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	if err := k.AuthKeeper.BeforeTokenTransfer(ctx, protocol, denomID, tokenID, sender, receiver); err != nil {
		return err
	}

	if err := k.TokenKeeper.BeforeTokenTransfer(ctx, protocol, denomID, tokenID, sender, receiver); err != nil {
		return err
	}

	if err := k.FeeKeeper.BeforeTokenTransfer(ctx, protocol, denomID, tokenID, sender, receiver); err != nil {
		return err
	}
	return nil
}
