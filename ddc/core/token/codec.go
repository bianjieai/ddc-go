package token

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
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
	cdc.RegisterConcrete(&MsgApprove{}, "bianjieai.ddc.token.Msg/MsgApprove", nil)
	cdc.RegisterConcrete(&MsgApproveForAll{}, "bianjieai.ddc.token.Msg/MsgApproveForAll", nil)
	cdc.RegisterConcrete(&MsgFreeze{}, "bianjieai.ddc.token.Msg/MsgFreeze", nil)
	cdc.RegisterConcrete(&MsgUnfreeze{}, "bianjieai.ddc.token.Msg/MsgUnfreeze", nil)
	cdc.RegisterConcrete(&MsgBatchBurn{}, "bianjieai.ddc.token.Msg/MsgBatchBurn", nil)
	cdc.RegisterConcrete(&MsgBatchTransfer{}, "bianjieai.ddc.token.Msg/MsgBatchTransfer", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApprove{},
		&MsgApproveForAll{},
		&MsgFreeze{},
		&MsgUnfreeze{},
		&MsgBatchBurn{},
		&MsgBatchTransfer{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
