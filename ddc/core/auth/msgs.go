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
	if len(m.Addresses) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the length of addresses can not be zero")
	}

	if len(m.Addresses) != len(m.Names) || len(m.Dids) != len(m.Names) {
		return sdkerrors.Wrap(ErrLengthMismatch, "the length of addresses,names,dids mismatch")
	}

	for i := range m.Addresses {
		if len(strings.TrimSpace(m.Dids[i])) == 0 {
			return sdkerrors.Wrap(ErrInvalidDID, "DID cannot be empty!")
		}

		if len(strings.TrimSpace(m.Names[i])) == 0 {
			return sdkerrors.Wrap(ErrInvalidName, "Name cannot be empty!")
		}

		_, err := sdk.AccAddressFromBech32(m.Addresses[i])
		if err != nil {
			return err
		}
	}
	_, err := sdk.AccAddressFromBech32(m.Sender)
	return err
}

// GetSigners implements Msg.
func (m MsgAddBatchAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgUpdateAccountState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return err
	}

	if _, ok := core.State_value[m.State.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidState, "state not exist")
	}

	_, err = sdk.AccAddressFromBech32(m.Sender)
	return err
}

// GetSigners implements Msg.
func (m MsgUpdateAccountState) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgDeleteAccount) ValidateBasic() error {
	return sdkerrors.Wrap(ErrInvalidOperator, "not implement")
}

// GetSigners implements Msg.
func (m MsgDeleteAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgAddFunction) ValidateBasic() error {
	if _, ok := core.Role_value[m.Role.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidRole, "role not exist")
	}

	if _, ok := core.Protocol_value[m.Protocol.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidProtocol, "protocol not exist")
	}

	if _, ok := core.Function_value[m.Function.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidFunction, "function not exist")
	}

	if len(strings.TrimSpace(m.Denom)) == 0 {
		return sdkerrors.Wrap(ErrInvalidDenom, "denom cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(m.Operator)
	return err
}

// GetSigners implements Msg.
func (m MsgAddFunction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgDeleteFunction) ValidateBasic() error {
	if _, ok := core.Role_value[m.Role.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidRole, "role not exist")
	}

	if _, ok := core.Protocol_value[m.Protocol.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidProtocol, "protocol not exist")
	}

	if _, ok := core.Function_value[m.Function.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidFunction, "function not exist")
	}

	if len(strings.TrimSpace(m.Denom)) == 0 {
		return sdkerrors.Wrap(ErrInvalidDenom, "denom cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(m.Operator)
	return err
}

// GetSigners implements Msg.
func (m MsgDeleteFunction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgApproveCrossPlatform) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.To)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.Operator)
	return err
}

// GetSigners implements Msg.
func (m MsgApproveCrossPlatform) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgSyncPlatformDID) ValidateBasic() error {
	if len(m.DIDs) == 0 {
		return sdkerrors.Wrap(ErrInvalidOperator, "dids can not be empty")
	}
	return nil
}

// GetSigners implements Msg.
func (m MsgSyncPlatformDID) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgUpgradeToDDC) ValidateBasic() error {
	if _, ok := core.Protocol_value[m.Protocol.String()]; !ok {
		return sdkerrors.Wrap(ErrInvalidProtocol, "protocol not exist")
	}

	if len(strings.TrimSpace(m.Denom)) == 0 {
		return sdkerrors.Wrap(ErrInvalidDenom, "denom cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(m.Operator)
	return err
}

// GetSigners implements Msg.
func (m MsgUpgradeToDDC) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgSetSwitcherStateOfPlatform) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Operator)
	return err
}

// GetSigners implements Msg.
func (m MsgSetSwitcherStateOfPlatform) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}
