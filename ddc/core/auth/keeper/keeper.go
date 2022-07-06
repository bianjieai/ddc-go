package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
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
