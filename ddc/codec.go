package ddc

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/bianjieai/ddc-go/ddc/core/auth"
	"github.com/bianjieai/ddc-go/ddc/core/fee"
	"github.com/bianjieai/ddc-go/ddc/core/token"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

// RegisterLegacyAminoCodec registers the necessary interfaces and concrete types
// on the provided Amino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	auth.RegisterLegacyAminoCodec(cdc)
	fee.RegisterLegacyAminoCodec(cdc)
	token.RegisterLegacyAminoCodec(cdc)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	auth.RegisterInterfaces(registry)
	fee.RegisterInterfaces(registry)
	token.RegisterInterfaces(registry)
}
