package token

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDenom    = sdkerrors.Register("ddc", 40, "invalid Denom")
	ErrInvalidTokenID  = sdkerrors.Register("ddc", 41, "invalid TokenID")
	ErrBlackListedDDC  = sdkerrors.Register("ddc", 42, "blacklisted DDC")
	ErrInvalidApprovee = sdkerrors.Register("ddc", 43, "invalid Approvee")
	ErrInvalidOperator = sdkerrors.Register("ddc", 44, "invalid Operator")
)
