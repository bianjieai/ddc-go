package keeper

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L327
func (k Keeper) requireTransferConstraintsSecondStep(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender string,
) error {
	if err := k.requireAvailableDDC(ctx, protocol, denomID, tokenID); err != nil {
		return err
	}

	if err := k.requireApprovedOrOwner(ctx, protocol, denomID, tokenID, sender); err != nil {
		return err
	}

	return nil
}

// implement:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L786
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L601
func (k Keeper) requireAvailableDDC(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
) error {
	// TODO: requireExists checked by Auth

	proto := core.Protocol_value[protocol]
	if k.isInBlocklist(ctx, core.Protocol(proto), denomID, tokenID) {
		return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is already in blocklist")
	}

	return nil
}

// implement:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L798
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L613
func (k Keeper) requireDisabledDDC(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
) error {
	// TODO: requireExists checked by Auth

	proto := core.Protocol_value[protocol]
	if !k.isInBlocklist(ctx, core.Protocol(proto), denomID, tokenID) {
		return sdkerrors.Wrapf(token.ErrDDCBlockList, "ddc is not in blocklist")
	}

	return nil
}

// implement:
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L837
// - https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L690
func (k Keeper) requireApprovedOrOwner(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender string,
) error {
	var err error
	switch protocol {
	case core.Protocol_name[0]:
		err = k.requireApprovedOrOwnerDDC721(ctx, protocol, denomID, tokenID, sender)
	case core.Protocol_name[1]:
		err = k.requireApprovedOrOwnerDDC1155(ctx, protocol, denomID, sender)
	}
	return err
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L837
func (k Keeper) requireApprovedOrOwnerDDC721(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender string,
) error {
	// get owner
	nft, err := k.nftKeeper.GetNFT(ctx, denomID, tokenID)
	if err != nil {
		return err
	}

	owner := nft.GetOwner().String()
	if sender == owner {
		return nil
	}

	proto := core.Protocol_value[protocol]
	if k.getDDCApproval(ctx, core.Protocol(proto), denomID, tokenID) == sender {
		return nil
	}

	if k.isApprovedForAll(ctx, core.Protocol(proto), denomID, owner, sender) {
		return nil
	}

	return sdkerrors.Wrapf(token.ErrRequireNotMet, "operator is not approved nor owner")
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC1155/DDC1155.sol#L690
func (k Keeper) requireApprovedOrOwnerDDC1155(ctx sdk.Context,
	protocol string,
	denomID string,
	sender string,
) error {
	// TODO: get owner
	var owner string
	panic("owner is empty")

	if sender == owner {
		return nil
	}

	proto := core.Protocol_value[protocol]
	if k.isApprovedForAll(ctx, core.Protocol(proto), denomID, owner, sender) {
		return nil
	}

	return sdkerrors.Wrapf(token.ErrRequireNotMet, "operator is not approved nor owner")
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L436
func (k Keeper) requireExistsAndApproved(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender string,
) error {
	// TODO: requireExists checked by Auth

	if err := k.requireApprovedOrOwner(ctx, protocol, denomID, tokenID, sender); err != nil {
		return err
	}
	return nil
}

// implement:  https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L245
func (k Keeper) requireApprovalConstraints(ctx sdk.Context, operator, to string) error {
	// TODO: auth
	// k.requireSenderHasFuncPermission() {}
	// k.requireAvailableDDCAccount() {}
	// k.requireOnePlatform() {}

	return nil
}
