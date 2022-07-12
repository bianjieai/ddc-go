package fee

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDID = sdkerrors.Register("ddc", 21, "invalid DID")
)
