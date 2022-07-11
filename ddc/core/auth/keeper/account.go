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

func (k Keeper) updateAccountState(ctx sdk.Context,
	address string,
	state core.State,
	changePlatformState bool,
	sender string,
) error {
	account, err := k.GetAccount(ctx, address)
	if err != nil {
		return err
	}
	senderAcc, err := k.GetAccount(ctx, sender)
	if err != nil {
		return err
	}

	if !k.isActive(senderAcc) {
		return sdkerrors.Wrapf(auth.ErrAccountNotActive, "account: %s is not active", address)
	}

	if !(account.LeaderDID == senderAcc.DID || senderAcc.Role == core.Role_OPERATOR) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "account's role does not match")
	}

	if senderAcc.Role == core.Role_CONSUMER {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "wrong role")
	}

	switch senderAcc.Role {
	case core.Role_CONSUMER:
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "wrong role")
	case core.Role_OPERATOR:
		if changePlatformState {
			if account.PlatformState == state {
				return sdkerrors.Wrap(auth.ErrInvalidOperator, "PlatformState doesn't need to change")
			}
			account.PlatformState = state
		} else {
			if account.OperatorState == state {
				return sdkerrors.Wrap(auth.ErrInvalidOperator, "operatorState doesn't need to change")
			}
			account.OperatorState = state
		}
	case core.Role_PLATFORM_MANAGER:
		if account.PlatformState == state {
			return sdkerrors.Wrap(auth.ErrInvalidOperator, "platformState doesn't need to change")
		}
		account.PlatformState = state
	}
	return k.setAccount(ctx, account)
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
	return k.setAccount(ctx, account)
}

func (k Keeper) approveCrossPlatform(ctx sdk.Context, from, to string) error {
	fromInfo, err := k.requireAccountActive(ctx, from)
	if err != nil {
		return err
	}

	toInfo, err := k.requireAccountActive(ctx, to)
	if err != nil {
		return err
	}

	if !(fromInfo.Role == core.Role_PLATFORM_MANAGER &&
		toInfo.Role == core.Role_PLATFORM_MANAGER) {
		return sdkerrors.Wrap(auth.ErrInvalidOperator, "both should be `platform` roles")
	}

	if fromInfo.DID == toInfo.DID {
		return sdkerrors.Wrap(auth.ErrInvalidOperator, "both should not be the same platform account")
	}

	store := k.prefixStore(ctx)
	store.Set(crossPlatformKey(fromInfo.DID, toInfo.DID), Placeholder)
	return nil
}

func (k Keeper) crossPlatformApproval(ctx sdk.Context, fromDID, toDID string) bool {
	store := k.prefixStore(ctx)
	return store.Has(crossPlatformKey(fromDID, toDID))
}

func (k Keeper) getPlatformSwitcher(ctx sdk.Context) bool {
	store := k.prefixStore(ctx)
	return store.Has(platformSwitcher())
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

// implement: https://github.com/bianjieai/tibc-ddc/blob/master/contracts/logic/Authority/Authority.sol#L676
func (k Keeper) requireAccountActive(ctx sdk.Context, address string) (*core.AccountInfo, error) {
	account, err := k.GetAccount(ctx, address)
	if err != nil {
		return nil, err
	}

	if !k.isActive(account) {
		return nil, sdkerrors.Wrapf(auth.ErrAccountNotActive, "account: %s is not active", address)
	}
	return account, nil
}

func (k Keeper) isActive(account *core.AccountInfo) bool {
	return account.OperatorState == core.State_ACTIVE && account.PlatformState == core.State_ACTIVE
}

func (k Keeper) setAccount(ctx sdk.Context, account *core.AccountInfo) error {
	bz, err := k.cdc.Marshal(account)
	if err != nil {
		return err
	}

	store := k.prefixStore(ctx)
	store.Set(accountKey(account.Address), bz)
	return nil
}

func (k Keeper) isRoot(ctx sdk.Context, address string) bool {
	return k.GetRoot(ctx) == address
}

func (k Keeper) savePlatformDID(ctx sdk.Context, did string) {
	store := k.prefixStore(ctx)
	store.Set(platformDIDKey(did), []byte(did))
}
