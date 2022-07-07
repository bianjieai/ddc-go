package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

// GetAccount return the account information by address
func (k Keeper) GetAccount(ctx sdk.Context, address string) (account *core.AccountInfo, err error) {
	store := k.prefixStore(ctx)
	bz := store.Get(accountKey(address))
	if bz == nil || len(bz) == 0 {
		return account, sdkerrors.Wrapf(auth.ErrAccountNotExist, "Account: %s not exist", address)
	}
	err = k.cdc.Unmarshal(bz, account)
	return
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L58
func (k Keeper) addOperator(ctx sdk.Context,
	account,
	accountName,
	accountDID string,
) error {
	if !k.requireNotExist(ctx, account) {
		return sdkerrors.Wrapf(auth.ErrAccountHasExist, "Account: %s has exist", account)
	}
	return k.addAccount(ctx, account, accountDID, accountDID, accountName, core.Role_OPERATOR)
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L81
func (k Keeper) addAccountByPlatform(ctx sdk.Context,
	account,
	accountName,
	accountDID string,
	sender *core.AccountInfo,
) error {
	if !k.requireOpenedSwitcherOfPlatform(ctx) {
		return sdkerrors.Wrapf(auth.ErrPlatformSwitcherClosed, "Account: %s no access", account)
	}
	if !k.requireNotExist(ctx, account) {
		return sdkerrors.Wrapf(auth.ErrAccountHasExist, "Account: %s has exist", account)
	}
	return k.addAccount(ctx, account, accountDID, sender.DID, accountName, core.Role_CONSUMER)
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L103
func (k Keeper) addBatchAccountByPlatform(ctx sdk.Context,
	accounts []string,
	accountNames []string,
	accountDIDs []string,
	sender *core.AccountInfo,
) error {
	if !k.requireOpenedSwitcherOfPlatform(ctx) {
		return sdkerrors.Wrapf(auth.ErrPlatformSwitcherClosed, "Account: %s no access", sender.Address)
	}

	for i := range accounts {
		if !k.requireNotExist(ctx, accounts[i]) {
			return sdkerrors.Wrapf(auth.ErrAccountHasExist, "Account: %s has exist", accounts[i])
		}
		if err := k.addAccount(ctx, accounts[i],
			accountDIDs[i], sender.DID, accountNames[i], core.Role_CONSUMER); err != nil {
			return err
		}
	}
	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L158
func (k Keeper) addAccountByOperator(ctx sdk.Context,
	account,
	accountName,
	accountDID,
	leaderDID string,
	sender *core.AccountInfo,
) error {
	if !k.requireNotExist(ctx, account) {
		return sdkerrors.Wrapf(auth.ErrAccountHasExist, "Account: %s has exist", account)
	}
	role := core.Role_CONSUMER
	if len(leaderDID) == 0 {
		// PlatformManager
		leaderDID = sender.DID
		role = core.Role_PLATFORM_MANAGER
		k.savePlatformDID(ctx, accountDID)
		return k.addAccount(ctx, account, accountDID, leaderDID, accountName, role)
	}
	// CONSUMER
	if !k.requireExistPlatformDID(ctx, leaderDID) {
		return sdkerrors.Wrapf(auth.ErrPlatformNotExist, "leaderDID: %s not exist", leaderDID)
	}
	return k.addAccount(ctx, account, accountDID, leaderDID, accountName, role)
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L172
func (k Keeper) addBatchAccountByOperator(ctx sdk.Context,
	accounts []string,
	accountNames []string,
	accountDIDs []string,
	leaderDIDs []string,
	sender *core.AccountInfo,
) error {
	for i := range accounts {
		if err := k.addAccountByOperator(ctx, accounts[i],
			accountNames[i], accountDIDs[i], leaderDIDs[i], sender); err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) addAccount(ctx sdk.Context,
	address string,
	did string,
	leaderDID string,
	name string,
	role core.Role,
) error {
	account := &core.AccountInfo{
		Address:       address,
		DID:           did,
		Name:          name,
		Role:          role,
		LeaderDID:     leaderDID,
		PlatformState: core.State_ACTIVE,
		OperatorState: core.State_ACTIVE,
	}
	bz, err := k.cdc.Marshal(account)
	if err != nil {
		return err
	}

	store := k.prefixStore(ctx)
	store.Set(accountKey(address), bz)
	return nil
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L690
func (k Keeper) requireNotExist(ctx sdk.Context, address string) bool {
	store := k.prefixStore(ctx)
	return !store.Has(accountKey(address))
}

func (k Keeper) requireExistPlatformDID(ctx sdk.Context, did string) bool {
	store := k.prefixStore(ctx)
	return store.Has(platformDIDKey(did))
}

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L722
func (k Keeper) requireOpenedSwitcherOfPlatform(ctx sdk.Context) bool {
	store := k.prefixStore(ctx)
	return store.Has(platformSwitcher())
}

func (k Keeper) isRoot(ctx sdk.Context, address string) bool {
	return k.GetRoot(ctx) == address
}

func (k Keeper) savePlatformDID(ctx sdk.Context, did string) {
	store := k.prefixStore(ctx)
	store.Set(platformDIDKey(did), []byte(did))
}
