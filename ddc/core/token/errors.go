package token

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDenom       = sdkerrors.Register("ddc", 40, "invalid Denom")
	ErrInvalidTokenID     = sdkerrors.Register("ddc", 41, "invalid TokenID")
	ErrInvalidApprovee    = sdkerrors.Register("ddc", 42, "invalid Approvee")
	ErrInvalidOperator    = sdkerrors.Register("ddc", 43, "invalid Operator")
	ErrInvalidOwner       = sdkerrors.Register("ddc", 44, "invalid Owner")
	ErrInvalidProtocol    = sdkerrors.Register("ddc", 45, "invalid Protocol")
	ErrInvalidAmount      = sdkerrors.Register("ddc", 46, "invalid Amount")
	ErrInconsistentLength = sdkerrors.Register("ddc", 47, "inconsistent length")
	ErrNonExistentDDC     = sdkerrors.Register("ddc", 48, "DDC is not existent")
	ErrDDCBlockList       = sdkerrors.Register("ddc", 49, "DDC blocklist error")
	ErrRequireNotMet      = sdkerrors.Register("ddc", 50, "Requirement is noe met")
)
