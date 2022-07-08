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
	ErrPlatformSwitcherClosed = sdkerrors.Register("ddc", 8, "platform switcher closed")
	ErrLengthMismatch         = sdkerrors.Register("ddc", 9, "length mismatch")
	ErrInvalidProtocol        = sdkerrors.Register("ddc", 10, "invalid protocol")
	ErrInvalidFunction        = sdkerrors.Register("ddc", 11, "invalid function")
	ErrInvalidDenom           = sdkerrors.Register("ddc", 12, "invalid denom")
	ErrFunctionHasExist       = sdkerrors.Register("ddc", 13, "function already exists")
	ErrFunctionNotExist       = sdkerrors.Register("ddc", 14, "function not exists")
	ErrAccountNotActive       = sdkerrors.Register("ddc", 15, "account has been frozen")
	ErrInvalidOperator        = sdkerrors.Register("ddc", 16, "invalid operate")
)
