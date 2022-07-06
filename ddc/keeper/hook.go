package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/export"
	sdk "github.com/cosmos/cosmos-sdk/types"

	authkeeper "github.com/bianjieai/ddc-go/ddc/core/auth/keeper"
)

var _ export.Hook = Hooks{}

type Hooks struct {
	hs         []export.Hook
	authKeeper authkeeper.Keeper
}

// BeforeDenomTransfer implements export.Hook
func (hs Hooks) BeforeDenomTransfer(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress) error {
	if !hs.authKeeper.ControlByDDC(ctx, denomID) {
		return nil
	}
	for _, h := range hs.hs {
		if err := h.BeforeDenomTransfer(ctx, protocol, denomID, sender); err != nil {
			return err
		}
	}
	return nil
}

// BeforeTokenBurn implements export.Hook
func (hs Hooks) BeforeTokenBurn(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	if !hs.authKeeper.ControlByDDC(ctx, denomID) {
		return nil
	}
	for _, h := range hs.hs {
		if err := h.BeforeTokenBurn(ctx, protocol, denomID, tokenID, sender); err != nil {
			return err
		}
	}
	return nil
}

// BeforeTokenEdit implements export.Hook
func (hs Hooks) BeforeTokenEdit(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress) error {
	if !hs.authKeeper.ControlByDDC(ctx, denomID) {
		return nil
	}
	for _, h := range hs.hs {
		if err := h.BeforeTokenEdit(ctx, protocol, denomID, tokenID, sender); err != nil {
			return err
		}
	}
	return nil
}

// BeforeTokenMint implements export.Hook
func (hs Hooks) BeforeTokenMint(ctx sdk.Context, protocol string, denomID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	if !hs.authKeeper.ControlByDDC(ctx, denomID) {
		return nil
	}
	for _, h := range hs.hs {
		if err := h.BeforeTokenMint(ctx, protocol, denomID, sender, receiver); err != nil {
			return err
		}
	}
	return nil
}

// BeforeTokenTransfer implements export.Hook
func (hs Hooks) BeforeTokenTransfer(ctx sdk.Context, protocol string, denomID string, tokenID string, sender sdk.AccAddress, receiver sdk.AccAddress) error {
	if !hs.authKeeper.ControlByDDC(ctx, denomID) {
		return nil
	}
	for _, h := range hs.hs {
		if err := h.BeforeTokenTransfer(ctx, protocol, denomID, tokenID, sender, receiver); err != nil {
			return err
		}
	}
	return nil
}
