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
	// TODO: auth
	// requireSenderHasFuncPermission

	for i := 0; i <= len(tokenIDs); i++ {
		if err := k.requireExistsAndApproved(ctx, core.Protocol_name[int32(protocol)], denomID, tokenIDs[i], operator); err != nil {
			return err
		}

		nft, err := k.nftKeeper.GetNFT(ctx, denomID, tokenIDs[i])
		if err != nil {
			return err
		}

		owner := nft.GetOwner().String()
		if operator != owner {
			return sdkerrors.Wrapf(token.ErrInvalidOwner, "operator is not the owner")
		}

		ownerAddr, err := sdk.AccAddressFromBech32(owner)
		if err != nil {
			return err
		}

		err = k.nftKeeper.BurnNFT(ctx, denomID, tokenIDs[i], ownerAddr)
		if err != nil {
			return err
		}

		// TODO: clear approvals
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
	// TODO: auth
	// requireSenderHasFuncPermission()

	if err := k.requireApprovedOrOwner(ctx, core.Protocol_name[int32(protocol)], denomID, "", operator); err != nil {
		return err
	}

	// TODO:
	// var owners []string
	// getOwners => owners
	// for each tokenIDs
	//   for each owner
	//      do burn ...
	var owner string
	panic("owner is empty")
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

		// TODO: deleted in blocklist
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

	// TODO: auth
	// _requireTransferConstraints_FistStep

	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return nil
	}
	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return nil
	}

	for i := 0; i <= len(tokenIDs); i++ {
		if err := k.requireTransferConstraintsSecondStep(ctx, core.Protocol_name[int32(protocol)], denomID, tokenIDs[i], sender); err != nil {
			return err
		}

		nft, err := k.nftKeeper.GetNFT(ctx, denomID, tokenIDs[i])
		if err != nil {
			return err
		}

		err = k.nftKeeper.TransferOwner(ctx, denomID, tokenIDs[i], nft.GetName(), nft.GetURI(), nft.GetURIHash(), nft.GetData(), fromAddr, toAddr)
		if err != nil {
			return err
		}

		// TODO: clearApprovals
		// TODO: pay
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

	if err := k.requireApprovedOrOwner(ctx, core.Protocol_name[int32(protocol)], denomID, "", sender); err != nil {
		return err
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
