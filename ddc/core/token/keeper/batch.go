package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L444
func (k Keeper) batchBurnDDC721(ctx sdk.Context,
	denomID string,
	tokenIDs []string,
	operator string,
	protocol core.Protocol,
) error {
	// TODO
	// if !k.requireSenderHasFuncPermission() {}

	// NOTE: what if err arises in iteration
	for i := 0; i <= len(tokenIDs); i++ {
		nft, err := k.nftKeeper.GetNFT(ctx, denomID, tokenIDs[i])
		if err != nil {
			return sdkerrors.Wrapf(token.ErrNonExistentDDC, "ddc is not existent")
		}

		owner := nft.GetOwner().String()
		approvee := k.getDDCApproval(ctx, protocol, denomID, tokenIDs[i])
		approved := k.isApprovedForAll(ctx, protocol, denomID, owner, operator)
		if operator != owner && operator != approvee && !approved {
			return sdkerrors.Wrapf(token.ErrInvalidOperator, "operator has no access to burning ddc")
		}

		err = k.nftKeeper.BurnNFT(ctx, denomID, tokenIDs[i], nft.GetOwner())
		if err != nil {
			return err
		}
	}

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L315
func (k Keeper) batchBurnDDC1155(ctx sdk.Context,
	denomID string,
	tokenIDs []string,
	operator string,
	protocol core.Protocol,
) error {
	// TODO
	//if !k.requireSenderHasFuncPermission() {}

	denom, exist := k.mtKeeper.GetDenom(ctx, denomID)
	if !exist {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "denom is not existent")
	}

	owner := denom.Owner
	if !k.isApprovedForAll(ctx, protocol, denomID, owner, operator) && operator != owner {
		return sdkerrors.Wrapf(token.ErrInvalidOperator, "operator is not owner nor approved")
	}

	ownerAddr, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return err
	}

	for i := 0; i <= len(tokenIDs); i++ {
		// NOTE: burn all
		amount := k.mtKeeper.GetBalance(ctx, denomID, tokenIDs[i], ownerAddr)
		err := k.mtKeeper.BurnMT(ctx, denomID, tokenIDs[i], amount, ownerAddr)
		if err != nil {
			return err
		}
	}

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L377
func (k Keeper) batchTransferDDC721(ctx sdk.Context,
	denomID string,
	tokenIDs []string,
	from string,
	to string,
	sender string,
	protocol core.Protocol,
) error {

	// TODO
	//_requireSenderHasFuncPermission();
	//_requireAvailableDDCAccount(from);
	//_requireAvailableDDCAccount(to);
	//_requireOnePlatformOrCrossPlatformApproval(from, to);

	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return nil
	}
	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return nil
	}

	for i := 0; i <= len(tokenIDs); i++ {

		nft, err := k.nftKeeper.GetNFT(ctx, denomID, tokenIDs[i])
		if err != nil {
			return err
		}

		if k.isInBlocklist(ctx, protocol, denomID, tokenIDs[i]) {
			return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is in blocklist")
		}

		owner := nft.GetOwner().String()
		if !k.isApprovedForAll(ctx, protocol, denomID, owner, sender) && sender != owner {
			return sdkerrors.Wrapf(token.ErrInvalidOperator, "operator is not owner nor approved")
		}

		// NOTE: transfer?
		err = k.nftKeeper.TransferOwner(ctx, denomID, tokenIDs[i], nft.GetName(), nft.GetURI(), nft.GetURIHash(), nft.GetData(), fromAddr, toAddr)
		if err != nil {
			return err
		}
	}
	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L250
func (k Keeper) batchTransferDDC1155(ctx sdk.Context,
	denomID string,
	tokenIDs []string,
	amounts []uint64,
	from string,
	to string,
	sender string,
	protocol core.Protocol,
) error {
	// TODO
	// requireSenderHasFuncPermission()
	// requireAvailableDDCAccount(from) & (to)
	// requireOnePlatformOrCrossPlatformApproval()

	denom, exist := k.mtKeeper.GetDenom(ctx, denomID)
	if !exist {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "denom is not existent")
	}
	owner := denom.Owner
	if !k.isApprovedForAll(ctx, protocol, denomID, owner, sender) && sender != owner {
		return sdkerrors.Wrapf(token.ErrInvalidOperator, "operator is not owner nor approved")
	}

	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return err
	}
	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return err
	}

	for i := 0; i <= len(tokenIDs); i++ {
		err := k.mtKeeper.TransferOwner(ctx, denomID, tokenIDs[i], amounts[i], fromAddr, toAddr)
		if err != nil {
			return err
		}
	}

	return nil
}
