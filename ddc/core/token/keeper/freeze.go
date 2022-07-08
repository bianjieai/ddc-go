package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L405
func (k Keeper) freezeDDC721(ctx sdk.Context, denomID, tokenID, operator string, protocol core.Protocol) error {
	// TODO
	// if !requireSenderHasFuncPermission()
	// if !requireOperator()

	_, err := k.nftKeeper.GetNFT(ctx, denomID, tokenID)
	if err != nil {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "ddc is not existent")
	}

	if k.isInBlocklist(ctx, appendProtocolPrefix(denomID, protocol), tokenID) {
		return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is already in blocklist")
	}

	k.setTokenBlocklistKey(ctx, appendProtocolPrefix(denomID, protocol), tokenID)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L281
func (k Keeper) freezeDDC1155(ctx sdk.Context, denomID, tokenID, operator string, protocol core.Protocol) error {
	// TODO
	// if !requireSenderHasFuncPermission()
	// if !requireOperator()

	_, err := k.mtKeeper.GetMT(ctx, denomID, tokenID)
	if err != nil {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "ddc is not existent")
	}

	if k.isInBlocklist(ctx, denomID, tokenID) {
		return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is already in blocklist")
	}

	k.setTokenBlocklistKey(ctx, appendProtocolPrefix(denomID, protocol), tokenID)

	return nil
}

// https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L420
func (k Keeper) unfreezeDDC721(ctx sdk.Context, denomID, tokenID, operator string, protocol core.Protocol) error {
	// TODO
	// if !requireSenderHasFuncPermission()
	// if !requireOperator()

	_, err := k.nftKeeper.GetNFT(ctx, denomID, tokenID)
	if err != nil {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "ddc is not existent")
	}

	if !k.isInBlocklist(ctx, appendProtocolPrefix(denomID, protocol), tokenID) {
		return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is not in blocklist")
	}

	k.unsetTokenBlocklistKey(ctx, appendProtocolPrefix(denomID, protocol), tokenID)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L293
func (k Keeper) unfreezeDDC1155(ctx sdk.Context, denomID, tokenID, operator string, protocol core.Protocol) error {
	// TODO
	// if !requireSenderHasFuncPermission()
	// if !requireOperator()

	_, err := k.mtKeeper.GetMT(ctx, denomID, tokenID)
	if err != nil {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "ddc is not existent")
	}

	if !k.isInBlocklist(ctx, appendProtocolPrefix(denomID, protocol), tokenID) {
		return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is not in blocklist")
	}

	k.unsetTokenBlocklistKey(ctx, appendProtocolPrefix(denomID, protocol), tokenID)

	return nil
}

func (k Keeper) isInBlocklist(ctx sdk.Context, denom, tokenID string) bool {
	store := k.prefixStore(ctx)
	return store.Has(tokenBlocklistKey(denom, tokenID))
}

// setTokenBlocklistKey add a blocklist key which hasn't been set
func (k Keeper) setTokenBlocklistKey(ctx sdk.Context, denom, tokenID string) {
	store := k.prefixStore(ctx)
	if store.Has(tokenBlocklistKey(denom, tokenID)) {
		return
	}
	store.Set(tokenBlocklistKey(denom, tokenID), []byte{0x01})
}

// unsetTokenBlocklistKey remove a blocklist key which has already been set
func (k Keeper) unsetTokenBlocklistKey(ctx sdk.Context, denom, tokenID string) {
	store := k.prefixStore(ctx)
	if !store.Has(tokenBlocklistKey(denom, tokenID)) {
		return
	}
	store.Delete(tokenBlocklistKey(denom, tokenID))
}
