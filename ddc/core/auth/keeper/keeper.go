package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the auth store
type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey
}

// NewKeeper creates a new nft Keeper instance
func NewKeeper(cdc codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: key,
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
