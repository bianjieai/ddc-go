package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/core/fee"
)

// Keeper of the fee store
type Keeper struct {
	cdc        codec.Codec
	storeKey   sdk.StoreKey
	authKeeper fee.AuthKeeper
}

// NewKeeper creates a new nft Keeper instance
func NewKeeper(cdc codec.Codec, key sdk.StoreKey, authKeeper fee.AuthKeeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   key,
		authKeeper: authKeeper,
	}
}

func (k Keeper) prefixStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, []byte(SubModule))
}
