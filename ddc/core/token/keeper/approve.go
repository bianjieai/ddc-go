package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implements: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L224
func (k Keeper) approve(ctx sdk.Context, denom, tokenID, operator, to string) error {

	// NOTE: is necessary?
	if !k.requireApprovalConstraints(ctx, to) {
		// TODO
	}

	// require nft exists
	nft, err := k.nftKeeper.GetNFT(ctx, denom, tokenID)
	if err != nil {
		return nil
	}

	// require nft not blacklisted
	if !k.isInBlacklist(ctx, denom, tokenID) {
		return sdkerrors.Wrapf(token.ErrBlackListedDDC, "ddc is blacklisted")
	}

	// require not approving to owner
	owner := nft.GetOwner().String()
	if owner == to {
		return sdkerrors.Wrapf(token.ErrInvalidApprovee, "cannot approve to owner")
	}

	// require operator is owner or is approved for all
	if operator != owner && !k.isApprovedForAll(ctx, denom, owner, operator) {
		return sdkerrors.Wrapf(token.ErrInvalidOperator, "approve operator is not owner nor approved for all")
	}

	k.saveDDCApprovals(ctx, denom, tokenID, to)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L288
func (k Keeper) isApprovedForAll(ctx sdk.Context, denom, owner, operator string) bool {
	store := k.prefixStore(ctx)
	return store.Has(accountApprovalKey(denom, owner, operator))
}

func (k Keeper) isInBlacklist(ctx sdk.Context, denom, tokenID string) bool {
	store := k.prefixStore(ctx)
	return store.Has(tokenBlacklistKey(denom, tokenID))
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L245
func (k Keeper) requireApprovalConstraints(ctx sdk.Context, to string) bool {
	// TODO
	if !k.requireSenderHasFuncPermission() {
		// TODO
	}

	if !k.requireAvailableDDCAccount() {
		// TODO
	}

	if !k.requireOnePlatform() {
		// TODO
	}

	return true
}

func (k Keeper) saveDDCApprovals(ctx sdk.Context, denom, tokenId, to string) {
	store := k.prefixStore(ctx)
	store.Set(ddcApprovalKey(denom, tokenId), []byte(to))
}
