package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	mtkeeper "github.com/irisnet/irismod/modules/mt/keeper"
	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"

	authkeeper "github.com/bianjieai/ddc-go/ddc/core/auth/keeper"
	feekeeper "github.com/bianjieai/ddc-go/ddc/core/fee/keeper"
	tokenkeeper "github.com/bianjieai/ddc-go/ddc/core/token/keeper"
	"github.com/bianjieai/ddc-go/ddc/exported"
)

// Keeper of the auth store
type Keeper struct {
	AuthKeeper  authkeeper.Keeper
	FeeKeeper   feekeeper.Keeper
	TokenKeeper tokenkeeper.Keeper
}

// NewKeeper creates a new nft Keeper instance
func NewKeeper(cdc codec.Codec, key sdk.StoreKey, nftKeeper nftkeeper.Keeper, mtKeeper mtkeeper.Keeper, paramSpace paramstypes.Subspace) Keeper {
	authKeeper := authkeeper.NewKeeper(cdc, key, paramSpace)
	return Keeper{
		AuthKeeper:  authKeeper,
		FeeKeeper:   feekeeper.NewKeeper(cdc, key, authKeeper),
		TokenKeeper: tokenkeeper.NewKeeper(cdc, key, nftKeeper, mtKeeper),
	}
}

// Hooks implements export.Hook
func (k Keeper) Hooks() Hooks {
	return Hooks{
		hs: []exported.Hook{
			k.AuthKeeper,
			k.TokenKeeper,
			k.FeeKeeper,
		},
	}
}
