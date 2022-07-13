package fee

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDID         = sdkerrors.Register("ddc", 21, "invalid DID")
	ErrFeeRuleUnavailable = sdkerrors.Register("ddc", 22, "denom fee rule unavailable")
)
