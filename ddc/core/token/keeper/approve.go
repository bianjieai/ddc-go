package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implements: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L224
func (k Keeper) approve(ctx sdk.Context, denomID, tokenID, operator, to string) error {

	if !k.requireApprovalConstraintsDDC721(ctx, operator, to) {
		// TODO
	}

	// require nft exists
	nft, err := k.nftKeeper.GetNFT(ctx, denomID, tokenID)
	if err != nil {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "ddc is not existent")
	}

	// require nft not in blocklist
	if k.isInBlocklist(ctx, denomID, tokenID) {
		return sdkerrors.Wrapf(token.ErrBlockListedDDC, "ddc is in blocklist")
	}

	// require not approving to owner
	owner := nft.GetOwner().String()
	if owner == to {
		return sdkerrors.Wrapf(token.ErrInvalidApprovee, "cannot approve to owner")
	}

	// require operator is owner or is approved for all
	if operator != owner && !k.isApprovedForAll(ctx, denomID, owner, operator) {
		return sdkerrors.Wrapf(token.ErrInvalidOperator, "approve operator is not owner nor approved for all")
	}

	k.setDDCApprovals(ctx, appendProtocolPrefix(denomID, core.Protocol_NFT), tokenID, to)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L274
func (k Keeper) setApproveForAllDDC721(ctx sdk.Context, denomID, sender, operator string, protocol core.Protocol) error {

	if !k.requireApprovalConstraintsDDC721(ctx, operator, operator) {
		// TODO
	}

	denom, exist := k.nftKeeper.GetDenom(ctx, denomID)
	if !exist {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "denom is not existent")
	}

	// NOTE: is Creator the Owner?
	if denom.Creator != sender {
		return sdkerrors.Wrapf(token.ErrInvalidOwner, "sender is not the owner")
	}

	k.setAccountApprovalKey(ctx, appendProtocolPrefix(denomID, protocol), denom.Creator, operator)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L191
func (k Keeper) setApproveForAllDDC1155(ctx sdk.Context, denomID, sender, operator string, protocol core.Protocol) error {

	if !k.requireApprovalConstraintsDDC1155(ctx, operator) {
		// TODO
	}

	// require mt denom exists
	denom, exist := k.mtKeeper.GetDenom(ctx, denomID)
	if !exist {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "denom is not existent")
	}

	// require sender is not owner
	if denom.Owner != sender {
		return sdkerrors.Wrapf(token.ErrInvalidOwner, "sender is not the owner")
	}

	k.setAccountApprovalKey(ctx, appendProtocolPrefix(denomID, protocol), denom.Owner, operator)

	return nil
}

// implement:  https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L245
func (k Keeper) requireApprovalConstraintsDDC721(ctx sdk.Context, operator, to string) bool {
	// TODO
	// if !k.requireSenderHasFuncPermission() {}
	// if !k.requireAvailableDDCAccount() {}
	// if !k.requireOnePlatform() {}

	return true
}

// implement:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L569
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L586
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L621
func (k Keeper) requireApprovalConstraintsDDC1155(ctx sdk.Context, operator string) bool {
	// TODO
	// if !k.requireSenderHasFuncPermission() {}
	// if !k.requireAvailableDDCAccount() {}
	// if !k.requireOnePlatform() {}

	return true
}

// implement:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L288
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L206
func (k Keeper) isApprovedForAll(ctx sdk.Context, denom, owner, operator string) bool {
	store := k.prefixStore(ctx)
	return store.Has(accountApprovalKey(denom, owner, operator))
}

func (k Keeper) isInBlocklist(ctx sdk.Context, denom, tokenID string) bool {
	store := k.prefixStore(ctx)
	return store.Has(tokenBlocklistKey(denom, tokenID))
}

func (k Keeper) setDDCApprovals(ctx sdk.Context, denom, tokenId, to string) {
	store := k.prefixStore(ctx)
	store.Set(ddcApprovalKey(denom, tokenId), []byte(to))
}

func (k Keeper) setAccountApprovalKey(ctx sdk.Context, denom, owner, operator string) {
	store := k.prefixStore(ctx)
	if store.Has(accountApprovalKey(denom, owner, operator)) {
		return
	}
	store.Set(accountApprovalKey(denom, owner, operator), []byte{0x01})
}
