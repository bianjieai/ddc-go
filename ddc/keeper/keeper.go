package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// Keeper of the nft store
type Keeper struct {
	cdc codec.Codec
}
