package keeper

import (
	"bytes"
	"github.com/bianjieai/ddc-go/ddc/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DDCApprovalKey     = []byte{0x01}
	AccountApprovalKey = []byte{0x02}
	TokenBlockListKey  = []byte{0x03}

	Placeholder = []byte{0x01}
)

const (
	SubModule = "token"
)

// DDCApprovalKey/Protocol/DenomID/TokenID
func ddcApprovalKey(protocol core.Protocol,
	denomID string,
	tokenID string,
) []byte {
	pbz := sdk.Uint64ToBigEndian(uint64(protocol))
	dbz := []byte(denomID)
	tbz := []byte(tokenID)

	length := len(DDCApprovalKey) + len(pbz) + len(dbz) + len(tbz)

	b := bytes.NewBuffer(make([]byte, 0, length))
	b.Write(DDCApprovalKey)
	b.Write(pbz)
	b.Write(dbz)
	b.Write(tbz)
	return b.Bytes()
}

// AccountApprovalKey/Protocol/DenomID/Owner/Operator
func accountApprovalKey(protocol core.Protocol,
	denomID string,
	owner string,
	operator string,
) []byte {
	pbz := sdk.Uint64ToBigEndian(uint64(protocol))
	dbz := []byte(denomID)
	owbz := []byte(owner)
	opbz := []byte(operator)

	length := len(AccountApprovalKey) + len(pbz) + len(dbz) + len(owbz) + len(opbz)

	b := bytes.NewBuffer(make([]byte, 0, length))
	b.Write(AccountApprovalKey)
	b.Write(pbz)
	b.Write(dbz)
	b.Write(owbz)
	b.Write(opbz)
	return b.Bytes()
}

// TokenBlockListKey/Protocol/DenomID/TokenID
func tokenBlocklistKey(protocol core.Protocol,
	denomID string,
	tokenID string,
) []byte {
	pbz := sdk.Uint64ToBigEndian(uint64(protocol))
	dbz := []byte(denomID)
	tbz := []byte(tokenID)

	length := len(TokenBlockListKey) + len(pbz) + len(dbz) + len(tbz)
	b := bytes.NewBuffer(make([]byte, 0, length))
	b.Write(TokenBlockListKey)
	b.Write(pbz)
	b.Write(dbz)
	b.Write(tbz)
	return b.Bytes()
}
