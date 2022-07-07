package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper of the auth store
type Keeper struct {
	cdc        codec.Codec
	storeKey   sdk.StoreKey
	paramSpace paramstypes.Subspace
}

// NewKeeper creates a new nft Keeper instance
func NewKeeper(cdc codec.Codec, key sdk.StoreKey, paramSpace paramstypes.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   key,
		paramSpace: paramSpace,
	}
}

func (k Keeper) prefixStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, []byte(SubModule))
}

func (k Keeper) ControlByDDC(ctx sdk.Context, denomID string) bool {
	store := k.prefixStore(ctx)
	return store.Has(ddcKey(denomID))
}
