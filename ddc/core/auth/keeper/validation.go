package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

func (k Keeper) requireSenderHasFuncPermission(ctx sdk.Context,
	address string,
	protocol string,
	denom string,
	function core.Function) error {
	proto := core.Protocol_value[protocol]
	if err := k.hasFunctionPermission(ctx, address, core.Protocol(proto), denom, core.Function_MINT); err != nil {
		return err
	}
	return nil
}

func (k Keeper) requireAvailableDDCAccount(ctx sdk.Context, address string) error {
	_, err := k.requireAccountActive(ctx, address)
	return err
}

func (k Keeper) requireOnePlatform(ctx sdk.Context, sender, receiver string) error {
	return k.isOnePlatform(ctx, sender, receiver)
}

func (k Keeper) requireOnePlatformOrCrossPlatformApproval(ctx sdk.Context, sender, receiver string) error {
	if err := k.isOnePlatform(ctx, sender, receiver); err == nil {
		return nil
	}

	if err := k.isCrossPlatformApproval(ctx, sender, receiver); err == nil {
		return nil
	}

	return sdkerrors.Wrap(auth.ErrInvalidOperator, "only one platform or cross-platform approval")
}

func (k Keeper) isOnePlatform(ctx sdk.Context, sender, receiver string) error {
	senderAcc, err := k.requireAccountActive(ctx, sender)
	if err != nil {
		return err
	}

	receiverAcc, err := k.requireAccountActive(ctx, receiver)
	if err != nil {
		return err
	}

	// a. All are platform roles
	if senderAcc.Role == core.Role_PLATFORM_MANAGER && receiverAcc.Role == core.Role_PLATFORM_MANAGER {
		if senderAcc.LeaderDID == receiverAcc.LeaderDID && senderAcc.DID == receiverAcc.DID {
			return nil
		}
	}

	// b. `from` is the platform, `to` is the consumer
	if senderAcc.Role == core.Role_PLATFORM_MANAGER && receiverAcc.Role == core.Role_CONSUMER {
		if senderAcc.DID == receiverAcc.LeaderDID {
			return nil
		}
	}

	// c. `to` is the platform, `from` is the consumer
	if senderAcc.Role == core.Role_CONSUMER && receiverAcc.Role == core.Role_PLATFORM_MANAGER {
		if senderAcc.LeaderDID == receiverAcc.DID {
			return nil
		}
	}

	// d. Both are consumers
	if senderAcc.Role == core.Role_CONSUMER && receiverAcc.Role == core.Role_CONSUMER {
		if senderAcc.LeaderDID == receiverAcc.LeaderDID {
			return nil
		}
	}
	return sdkerrors.Wrap(auth.ErrInvalidOperator, "only on the same platform")
}

func (k Keeper) isCrossPlatformApproval(ctx sdk.Context, sender, receiver string) error {
	senderAcc, err := k.requireAccountActive(ctx, sender)
	if err != nil {
		return err
	}

	receiverAcc, err := k.requireAccountActive(ctx, receiver)
	if err != nil {
		return err
	}

	// a. All are platform roles
	if senderAcc.Role == core.Role_PLATFORM_MANAGER && receiverAcc.Role == core.Role_PLATFORM_MANAGER {
		if (senderAcc.LeaderDID == receiverAcc.LeaderDID && senderAcc.DID == receiverAcc.DID) ||
			(k.crossPlatformApproval(ctx, senderAcc.DID, receiverAcc.DID)) {
			return nil
		}
	}

	// b. `from` is the platform, `to` is the consumer
	if senderAcc.Role == core.Role_PLATFORM_MANAGER && receiverAcc.Role == core.Role_CONSUMER {
		if k.crossPlatformApproval(ctx, senderAcc.DID, receiverAcc.LeaderDID) {
			return nil
		}
	}

	// c. `to` is the platform, `from` is the consumer
	if senderAcc.Role == core.Role_CONSUMER && receiverAcc.Role == core.Role_PLATFORM_MANAGER {
		if k.crossPlatformApproval(ctx, senderAcc.LeaderDID, receiverAcc.DID) {
			return nil
		}
	}

	// d. Both are consumers
	if senderAcc.Role == core.Role_CONSUMER && receiverAcc.Role == core.Role_CONSUMER {
		if k.crossPlatformApproval(ctx, senderAcc.LeaderDID, receiverAcc.LeaderDID) {
			return nil
		}
	}
	return nil
}

func (k Keeper) requireTransferConstraints_FistStep(ctx sdk.Context,
	protocol,
	denomID,
	tokenID,
	sender, receiver string,
) error {
	if err := k.requireSenderHasFuncPermission(ctx, sender, protocol, denomID, core.Function_TRANSFER); err != nil {
		return err
	}

	if err := k.requireAvailableDDCAccount(ctx, sender); err != nil {
		return err
	}

	if err := k.requireAvailableDDCAccount(ctx, receiver); err != nil {
		return err
	}

	return k.requireOnePlatformOrCrossPlatformApproval(ctx, sender, receiver)
}
