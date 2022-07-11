package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

// ParamKeyTable for service module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&auth.Params{})
}

// GetParams gets all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) auth.Params {
	var root string
	k.paramSpace.Get(ctx, auth.KeyRoot, &root)
	return auth.Params{
		Root: root,
	}
}

// GetRoot return the root
func (k Keeper) GetRoot(ctx sdk.Context) string {
	var root string
	k.paramSpace.Get(ctx, auth.KeyRoot, &root)
	return root
}

// SetParams sets the params to the store
func (k Keeper) SetParams(ctx sdk.Context, params auth.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
