package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mtkeeper "github.com/irisnet/irismod/modules/mt/keeper"
	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"
)

// Keeper of the token store
type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey

	nftKeeper nftkeeper.Keeper
	mtKeeper  mtkeeper.Keeper
}

// NewKeeper creates a new nft Keeper instance
func NewKeeper(cdc codec.Codec,
	key sdk.StoreKey,
	nftKeeper nftkeeper.Keeper,
	mtKeeper mtkeeper.Keeper,
) Keeper {
	return Keeper{
		cdc:       cdc,
		storeKey:  key,
		nftKeeper: nftKeeper,
		mtKeeper:  mtKeeper,
	}
}

func (k Keeper) prefixStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, []byte(SubModule))
}
