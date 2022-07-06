package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	authkeeper "github.com/bianjieai/ddc-go/ddc/core/auth/keeper"
	feekeeper "github.com/bianjieai/ddc-go/ddc/core/fee/keeper"
	tokenkeeper "github.com/bianjieai/ddc-go/ddc/core/token/keeper"
)

// Keeper of the auth store
type Keeper struct {
	AuthKeeper  authkeeper.Keeper
	FeeKeeper   feekeeper.Keeper
	TokenKeeper tokenkeeper.Keeper
}

// NewKeeper creates a new nft Keeper instance
func NewKeeper(cdc codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		AuthKeeper:  authkeeper.NewKeeper(cdc, key),
		FeeKeeper:   feekeeper.NewKeeper(cdc, key),
		TokenKeeper: tokenkeeper.NewKeeper(cdc, key),
	}
}
