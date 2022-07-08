package token

import (
	"github.com/bianjieai/ddc-go/ddc/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

var (
	_ sdk.Msg = &MsgApprove{}
	_ sdk.Msg = &MsgApproveForAll{}
	_ sdk.Msg = &MsgFreeze{}
	_ sdk.Msg = &MsgUnfreeze{}
	_ sdk.Msg = &MsgBatchBurn{}
	_ sdk.Msg = &MsgBatchTransfer{}
)

// ValidateBasic implements Msg.
func (m MsgApprove) ValidateBasic() error {
	if len(strings.TrimSpace(m.Denom)) == 0 {
		return sdkerrors.Wrap(ErrInvalidDenom, "Denom cannot be empty!")
	}

	if len(strings.TrimSpace(m.TokenID)) == 0 {
		return sdkerrors.Wrap(ErrInvalidTokenID, "TokenID cannot be empty!")
	}

	_, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.To)
	if err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgApprove) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgApproveForAll) ValidateBasic() error {
	if !validateProtocol(m.Protocol) {
		return sdkerrors.Wrapf(ErrInvalidProtocol, "Protocol is not NFT nor MT!")
	}

	if len(strings.TrimSpace(m.Denom)) == 0 {
		return sdkerrors.Wrap(ErrInvalidDenom, "Denom cannot be empty!")
	}

	_, err := sdk.AccAddressFromBech32(m.Operator)
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
func (m MsgApproveForAll) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgFreeze) ValidateBasic() error {
	if !validateProtocol(m.Protocol) {
		return sdkerrors.Wrapf(ErrInvalidProtocol, "Protocol is not NFT nor MT!")
	}

	if len(strings.TrimSpace(m.Denom)) == 0 {
		return sdkerrors.Wrap(ErrInvalidDenom, "Denom cannot be empty!")
	}

	if len(strings.TrimSpace(m.TokenID)) == 0 {
		return sdkerrors.Wrap(ErrInvalidTokenID, "TokenID cannot be empty!")
	}

	_, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgFreeze) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgUnfreeze) ValidateBasic() error {
	if !validateProtocol(m.Protocol) {
		return sdkerrors.Wrapf(ErrInvalidProtocol, "Protocol is not NFT nor MT!")
	}

	if len(strings.TrimSpace(m.Denom)) == 0 {
		return sdkerrors.Wrap(ErrInvalidDenom, "Denom cannot be empty!")
	}

	if len(strings.TrimSpace(m.TokenID)) == 0 {
		return sdkerrors.Wrap(ErrInvalidTokenID, "TokenID cannot be empty!")
	}

	_, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgUnfreeze) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements Msg.
func (m MsgBatchBurn) ValidateBasic() error {
	// TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgBatchBurn) GetSigners() []sdk.AccAddress {
	// TODO
	return nil
}

// ValidateBasic implements Msg.
func (m MsgBatchTransfer) ValidateBasic() error {
	// TODO
	return nil
}

// GetSigners implements Msg.
func (m MsgBatchTransfer) GetSigners() []sdk.AccAddress {
	// TODO
	return nil
}

// validateProtocol validates protocol
func validateProtocol(protocol core.Protocol) bool {
	switch protocol {
	case core.Protocol_NFT:
	case core.Protocol_MT:
	default:
		return false
	}
	return true
}
