package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

// implements: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L224
func (k Keeper) approve(ctx sdk.Context,
	denomID string,
	tokenID string,
	operator string,
	to string,
) error {
	// requireApprovalConstraints
	if !k.requireApprovalConstraintsDDC721(ctx, operator, to) {
		// TODO
	}

	nft, err := k.nftKeeper.GetNFT(ctx, denomID, tokenID)
	if err != nil {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "ddc is not existent")
	}

	if k.isInBlocklist(ctx, core.Protocol_NFT, denomID, tokenID) {
		return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is already in blocklist")
	}

	owner := nft.GetOwner().String()
	if owner == to {
		return sdkerrors.Wrapf(token.ErrInvalidApprovee, "cannot approve to owner")
	}

	if operator != owner && !k.isApprovedForAll(ctx, core.Protocol_NFT, denomID, owner, operator) {
		return sdkerrors.Wrapf(token.ErrInvalidOperator, "approve operator is not owner nor approved for all")
	}

	k.setDDCApproval(ctx, core.Protocol_NFT, denomID, tokenID, to)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L274
func (k Keeper) setApproveForAllDDC721(ctx sdk.Context,
	denomID string,
	sender string,
	operator string,
	protocol core.Protocol,
) error {
	if !k.requireApprovalConstraintsDDC721(ctx, operator, operator) {
		// TODO
	}

	denom, exist := k.nftKeeper.GetDenom(ctx, denomID)
	if !exist {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "denom is not existent")
	}

	// NOTE: necessary?
	if denom.Creator != sender {
		return sdkerrors.Wrapf(token.ErrInvalidOwner, "sender is not the owner")
	}

	if operator != sender {
		return sdkerrors.Wrapf(token.ErrInvalidOperator, "operator should not the sender")
	}

	k.setAccountApproval(ctx, protocol, denomID, sender, operator)

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L191
func (k Keeper) setApproveForAllDDC1155(ctx sdk.Context,
	denomID string,
	sender string,
	operator string,
	protocol core.Protocol,
) error {

	if !k.requireApprovalConstraintsDDC1155(ctx, operator) {
		// TODO
	}

	denom, exist := k.mtKeeper.GetDenom(ctx, denomID)
	if !exist {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "denom is not existent")
	}

	// NOTE: necessary?
	if denom.Owner != sender {
		return sdkerrors.Wrapf(token.ErrInvalidOwner, "sender is not the owner")
	}

	if operator != sender {
		return sdkerrors.Wrapf(token.ErrInvalidOperator, "operator should not be the sender")
	}

	k.setAccountApproval(ctx, protocol, denomID, denom.Owner, operator)

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
func (k Keeper) isApprovedForAll(ctx sdk.Context,
	protocol core.Protocol,
	denomID string,
	owner string,
	operator string,
) bool {
	store := k.prefixStore(ctx)
	return store.Has(accountApprovalKey(protocol, denomID, owner, operator))
}

func (k Keeper) setDDCApproval(ctx sdk.Context,
	protocol core.Protocol,
	denomID string,
	tokenID string,
	to string,
) {
	store := k.prefixStore(ctx)
	store.Set(ddcApprovalKey(protocol, denomID, tokenID), []byte(to))
}

func (k Keeper) getDDCApproval(ctx sdk.Context,
	protocol core.Protocol,
	denomID string,
	tokenID string,
) string {
	store := k.prefixStore(ctx)
	to := store.Get(ddcApprovalKey(protocol, denomID, tokenID))
	return string(to[:])
}

func (k Keeper) setAccountApproval(ctx sdk.Context,
	protocol core.Protocol,
	denomID string,
	owner string,
	operator string,
) {
	store := k.prefixStore(ctx)
	key := accountApprovalKey(protocol, denomID, owner, operator)
	if store.Has(key) {
		return
	}
	store.Set(key, Placeholder)
}

func (k Keeper) getAccountsApproval(ctx sdk.Context,
	protocol core.Protocol,
	denomID string,
	owner string,
) []string {
	store := k.prefixStore(ctx)
	prefix := string(accountApprovalKey(protocol, denomID, owner, "")[:])

	iterator := sdk.KVStorePrefixIterator(store, AccountApprovalKey)
	defer iterator.Close()

	var operators []string
	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key()[:])
		if strings.HasPrefix(key, prefix) {
			operator := strings.TrimPrefix(key, prefix)
			operators = append(operators, operator)
		}
	}
	return operators
}
