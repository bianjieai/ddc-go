package auth

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keys for parameter access
// nolint
var (
	KeyRoot = []byte("Root")
)

// ParamSetPairs implements paramstypes.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyRoot, &p.Root, func(value interface{}) error {
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("invalid parameter type: %T", value)
			}

			_, err := sdk.AccAddressFromBech32(v)
			return err
		}),
	}
}
