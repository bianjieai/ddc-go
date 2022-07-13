package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

func (k Keeper) CheckAvailableAndRole(ctx sdk.Context, sender string, role core.Role) error {
	account, err := k.requireAccountActive(ctx, sender)
	if err != nil {
		return err
	}

	if account.Role != role {
		return sdkerrors.Wrap(auth.ErrInvalidOperator, "not a operator role or disabled")
	}
	return nil
}
