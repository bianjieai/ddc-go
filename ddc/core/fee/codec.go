package fee

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
	cdc.RegisterConcrete(&MsgRecharge{}, "bianjieai.ddc.fee.Msg/MsgRecharge", nil)
	cdc.RegisterConcrete(&MsgRechargeBatch{}, "bianjieai.ddc.fee.Msg/MsgRechargeBatch", nil)
	cdc.RegisterConcrete(&MsgSettlement{}, "bianjieai.ddc.fee.Msg/MsgSettlement", nil)
	cdc.RegisterConcrete(&MsgSetFeeRule{}, "bianjieai.ddc.fee.Msg/MsgSetFeeRule", nil)
	cdc.RegisterConcrete(&MsgDeleteFeeRule{}, "bianjieai.ddc.fee.Msg/MsgDeleteFeeRule", nil)
	cdc.RegisterConcrete(&MsgRevokeDDC{}, "bianjieai.ddc.fee.Msg/MsgRevokeDDC", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRecharge{},
		&MsgRechargeBatch{},
		&MsgRechargeBatch{},
		&MsgSettlement{},
		&MsgSetFeeRule{},
		&MsgDeleteFeeRule{},
		&MsgRevokeDDC{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
