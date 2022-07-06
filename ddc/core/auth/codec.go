package auth

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
	cdc.RegisterConcrete(&MsgAddAccount{}, "bianjieai.ddc.auth.Msg/AddAccount", nil)
	cdc.RegisterConcrete(&MsgAddBatchAccount{}, "bianjieai.ddc.auth.Msg/MsgAddBatchAccount", nil)
	cdc.RegisterConcrete(&MsgUpdateAccountState{}, "bianjieai.ddc.auth.Msg/MsgUpdateAccountState", nil)
	cdc.RegisterConcrete(&MsgDeleteAccount{}, "bianjieai.ddc.auth.Msg/MsgDeleteAccount", nil)
	cdc.RegisterConcrete(&MsgAddFunction{}, "bianjieai.ddc.auth.Msg/MsgAddFunction", nil)
	cdc.RegisterConcrete(&MsgDeleteFunction{}, "bianjieai.ddc.auth.Msg/MsgDeleteFunction", nil)
	cdc.RegisterConcrete(&MsgApproveCrossPlatform{}, "bianjieai.ddc.auth.Msg/MsgApproveCrossPlatform", nil)
	cdc.RegisterConcrete(&MsgSyncPlatformDID{}, "bianjieai.ddc.auth.Msg/MsgSyncPlatformDID", nil)
	cdc.RegisterConcrete(&MsgUpgradeToDDC{}, "bianjieai.ddc.auth.Msg/MsgUpgradeToDDC", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAccount{},
		&MsgAddBatchAccount{},
		&MsgUpdateAccountState{},
		&MsgDeleteAccount{},
		&MsgAddFunction{},
		&MsgDeleteFunction{},
		&MsgApproveCrossPlatform{},
		&MsgSyncPlatformDID{},
		&MsgUpgradeToDDC{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
