package auth

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDID             = sdkerrors.Register("ddc", 2, "invalid DID")
	ErrInvalidRole            = sdkerrors.Register("ddc", 3, "invalid role")
	ErrInvalidName            = sdkerrors.Register("ddc", 4, "invalid name")
	ErrAccountNotExist        = sdkerrors.Register("ddc", 5, "account not exist")
	ErrAccountHasExist        = sdkerrors.Register("ddc", 6, "account has exist")
	ErrPlatformNotExist       = sdkerrors.Register("ddc", 7, "platform not exist")
	ErrPlatformSwitcherClosed = sdkerrors.Register("ddc", 7, "platform switcher closed")
)
