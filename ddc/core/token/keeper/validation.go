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

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/DDC721/DDC721.sol#L837
// NOTE: ddc721 and ddc1155 have different implementations
func (k Keeper) requireApprovedOrOwner(ctx sdk.Context,
	protocol string,
	denomID string,
	tokenID string,
	sender string,
) error {
	var owner string
	// TODO: require getting owner by Auth

	if owner == sender {
		return nil
	}

	proto := core.Protocol_value[protocol]
	// getDDCApproval is called only when protocol is NFT
	if core.Protocol(proto) == core.Protocol_NFT && k.getDDCApproval(ctx, core.Protocol(proto), denomID, tokenID) == sender {
		return nil
	}
	if k.isApprovedForAll(ctx, core.Protocol(proto), denomID, owner, sender) {
		return nil
	}

	return sdkerrors.Wrapf(token.ErrRequireNotMet, "not owner nor approved")
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
