package auth

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgAddAccount{}
	_ sdk.Msg = &MsgAddBatchAccount{}
	_ sdk.Msg = &MsgUpdateAccountState{}
	_ sdk.Msg = &MsgDeleteAccount{}
	_ sdk.Msg = &MsgAddFunction{}
	_ sdk.Msg = &MsgDeleteFunction{}
	_ sdk.Msg = &MsgApproveCrossPlatform{}
	_ sdk.Msg = &MsgSyncPlatformDID{}
	_ sdk.Msg = &MsgUpgradeToDDC{}
)

// ValidateBasic implements Msg.
func (m MsgAddAccount) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgAddAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgAddBatchAccount) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgAddBatchAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgUpdateAccountState) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgUpdateAccountState) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgDeleteAccount) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgDeleteAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgAddFunction) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgAddFunction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgDeleteFunction) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgDeleteFunction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgApproveCrossPlatform) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgApproveCrossPlatform) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgSyncPlatformDID) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgSyncPlatformDID) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgUpgradeToDDC) ValidateBasic() error {
	//TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgUpgradeToDDC) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}
