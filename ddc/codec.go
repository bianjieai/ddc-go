package ddc

import (
	"github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/bianjieai/ddc-go/ddc/core/auth"
	"github.com/bianjieai/ddc-go/ddc/core/fee"
	"github.com/bianjieai/ddc-go/ddc/core/token"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	auth.RegisterInterfaces(registry)
	fee.RegisterInterfaces(registry)
	token.RegisterInterfaces(registry)
}
