package auth

import (
	"strings"

	"github.com/bianjieai/ddc-go/ddc/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	if len(strings.TrimSpace(m.Did)) == 0 {
		return sdkerrors.Wrap(ErrInvalidDID, "DID cannot be empty!")
	}

	if len(strings.TrimSpace(m.Name)) == 0 {
		return sdkerrors.Wrap(ErrInvalidName, "Name cannot be empty!")
	}

	if _, ok := core.Role_value[m.Role.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidRole, "DID cannot be empty!")
	}

	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

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
