package token

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDenom    = sdkerrors.Register("ddc", 40, "invalid Denom")
	ErrInvalidTokenID  = sdkerrors.Register("ddc", 41, "invalid TokenID")
	ErrBlackListedDDC  = sdkerrors.Register("ddc", 42, "blacklisted DDC")
	ErrNonExistentDDC  = sdkerrors.Register("ddc", 43, "DDC is not existent")
	ErrInvalidApprovee = sdkerrors.Register("ddc", 44, "invalid Approvee")
	ErrInvalidOperator = sdkerrors.Register("ddc", 45, "invalid Operator")
	ErrInvalidOwner    = sdkerrors.Register("ddc", 46, "invalid Owner")
	ErrInvalidProtocol = sdkerrors.Register("ddc", 47, "invalid Protocol")
)
