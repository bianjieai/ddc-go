package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/ddc-go/ddc/core"
)

func (k Keeper) checkRechargeAuth(ctx sdk.Context, from, to string, amount uint64) error {
	fromAcc, err := k.authKeeper.GetAccount(ctx, from)
	if err != nil {
		return err
	}

	if fromAcc.PlatformState != core.State_ACTIVE || fromAcc.OperatorState != core.State_ACTIVE {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "account: %s is frozen", from)
	}

	toAcc, err := k.authKeeper.GetAccount(ctx, from)
	if err != nil {
		return err
	}

	if toAcc.PlatformState != core.State_ACTIVE || toAcc.OperatorState != core.State_ACTIVE {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "account: %s is frozen", to)
	}

	pass := fromAcc.Role == core.Role_OPERATOR ||
		fromAcc.DID == toAcc.LeaderDID ||
		(fromAcc.LeaderDID == toAcc.LeaderDID &&
			fromAcc.DID == toAcc.DID &&
			toAcc.Role != core.Role_CONSUMER)
	if !pass {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "unsupported recharge operation")
	}
	return nil
}
