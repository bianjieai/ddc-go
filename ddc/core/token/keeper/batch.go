package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L444
func (k Keeper) batchBurnDDC721(ctx sdk.Context, denomID string, tokenIDs []string, operator string, protocol core.Protocol) error {
	// TODO
	// if !k.requireSenderHasFuncPermission() {}

	// NOTE: what if err arises in iteration
	for i := 0; i <= len(tokenIDs); i++ {
		// require ddc exist
		nft, err := k.nftKeeper.GetNFT(ctx, denomID, tokenIDs[i])
		if err != nil {
			return sdkerrors.Wrapf(token.ErrNonExistentDDC, "ddc is not existent")
		}

		// require approved or real owner
		prefixDenomID := appendProtocolPrefix(denomID, protocol)
		owner := nft.GetOwner().String()
		approvee := k.getAccountApproval(ctx, prefixDenomID, tokenIDs[i])
		approved := k.isApprovedForAll(ctx, prefixDenomID, owner, operator)
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
func (k Keeper) batchBurnDDC1155(ctx sdk.Context, denomID string, tokenIDs []string, operator string, protocol core.Protocol) error {
	// TODO
	//if !k.requireSenderHasFuncPermission() {}

	denom, exist := k.mtKeeper.GetDenom(ctx, denomID)
	if !exist {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "denom is not existent")
	}

	owner := denom.Owner
	if !k.isApprovedForAll(ctx, appendProtocolPrefix(denomID, protocol), owner, operator) && operator != owner {
		return sdkerrors.Wrapf(token.ErrInvalidOperator, "operator is not owner nor approved")
	}

	ownerAddr, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return err
	}

	for i := 0; i <= len(tokenIDs); i++ {
		// burn all
		amount := k.mtKeeper.GetBalance(ctx, denomID, tokenIDs[i], ownerAddr)
		err := k.mtKeeper.BurnMT(ctx, denomID, tokenIDs[i], amount, ownerAddr)
		if err != nil {
			return err
		}
	}

	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L377
func (k Keeper) batchTransferDDC721(ctx sdk.Context, denomID string, tokenIDs []string, from, to, sender string, protocol core.Protocol) error {

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

		owner := nft.GetOwner().String()
		if !k.isApprovedForAll(ctx, appendProtocolPrefix(denomID, protocol), owner, sender) && sender != owner {
			return sdkerrors.Wrapf(token.ErrInvalidOperator, "operator is not owner nor approved")
		}

		err = k.nftKeeper.TransferOwner(ctx, denomID, tokenIDs[i], nft.GetName(), nft.GetURI(), nft.GetURIHash(), nft.GetData(), fromAddr, toAddr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) batchTransferDDC1155(ctx sdk.Context, denomID string, tokenIDs []string, amounts []uint64, from, to, sender string, protocol core.Protocol) error {
	// TODO
	// requireSenderHasFuncPermission()
	// requireAvailableDDCAccount(from) & (to)
	// requireOnePlatformOrCrossPlatformApproval()

	denom, exist := k.mtKeeper.GetDenom(ctx, denomID)
	if !exist {
		return sdkerrors.Wrapf(token.ErrNonExistentDDC, "denom is not existent")
	}
	owner := denom.Owner
	if !k.isApprovedForAll(ctx, appendProtocolPrefix(denomID, protocol), owner, sender) && sender != owner {
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
		err := k.mtKeeper.Transfer(ctx, denomID, tokenIDs[i], amounts[i], fromAddr, toAddr)
		if err != nil {
			return err
		}
	}

	return nil
}
