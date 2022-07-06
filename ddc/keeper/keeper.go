package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	authkeeper "github.com/bianjieai/ddc-go/ddc/core/auth/keeper"
	feekeeper "github.com/bianjieai/ddc-go/ddc/core/fee/keeper"
	tokenkeeper "github.com/bianjieai/ddc-go/ddc/core/token/keeper"
	"github.com/bianjieai/ddc-go/ddc/export"
)

// Keeper of the auth store
type Keeper struct {
	AuthKeeper  authkeeper.Keeper
	FeeKeeper   feekeeper.Keeper
	TokenKeeper tokenkeeper.Keeper
}

// NewKeeper creates a new nft Keeper instance
func NewKeeper(cdc codec.Codec, key sdk.StoreKey, paramSpace paramstypes.Subspace) Keeper {
	return Keeper{
		AuthKeeper:  authkeeper.NewKeeper(cdc, key, paramSpace),
		FeeKeeper:   feekeeper.NewKeeper(cdc, key),
		TokenKeeper: tokenkeeper.NewKeeper(cdc, key),
	}
}

// Hooks implements export.Hook
func (k Keeper) Hooks() Hooks {
	return Hooks{
		hs: []export.Hook{
			k.AuthKeeper,
			k.TokenKeeper,
			k.FeeKeeper,
		},
		authKeeper: k.AuthKeeper,
	}
}
