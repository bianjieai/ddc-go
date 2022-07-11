package auth

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

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
		&MsgSetSwitcherStateOfPlatform{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
