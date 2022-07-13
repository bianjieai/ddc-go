package keeper

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/core"
)

var (
	BalanceKey = []byte{0x01}
	FeeRuleKey = []byte{0x02}
	DDCAuthKey = []byte{0x03}
	SupplyKey  = []byte{0x04}

	Delimiter   = []byte{0x00}
	Placeholder = []byte{0x01}
)

const (
	SubModule = "fee"
)

func balanceKey(address string) []byte {
	addressBz := []byte(address)
	key := make([]byte, 0, len(BalanceKey)+len(addressBz))
	copy(key, BalanceKey)
	copy(key[len(BalanceKey):], addressBz)
	return key
}

// feeRuleKey returns the byte representation of the FeeRule
func feeRuleKey(protocol core.Protocol, denom string, function core.Function) []byte {
	pbz := sdk.Uint64ToBigEndian(uint64(protocol))
	fbz := sdk.Uint64ToBigEndian(uint64(function))
	dbz := []byte(denom)

	len := len(FeeRuleKey) + len(pbz) + len(Delimiter) + len(fbz) + len(Delimiter) + len(dbz)

	b := bytes.NewBuffer(make([]byte, 0, len))
	b.Write(FeeRuleKey)
	b.Write(pbz)
	b.Write(Delimiter)
	b.Write(fbz)
	b.Write(Delimiter)
	b.Write(dbz)
	return b.Bytes()
}

func ddcAuthKey(protocol core.Protocol, denom string) []byte {
	pbz := sdk.Uint64ToBigEndian(uint64(protocol))
	dbz := []byte(denom)

	len := len(DDCAuthKey) + len(pbz) + len(Delimiter) + len(dbz)
	b := bytes.NewBuffer(make([]byte, 0, len))
	b.Write(DDCAuthKey)
	b.Write(pbz)
	b.Write(Delimiter)
	b.Write(dbz)
	return b.Bytes()
}

func GetDDCEscrowAddress(protocol core.Protocol, denom string) sdk.AccAddress {
	contents := fmt.Sprintf("%s/%s", protocol, denom)
	hash := sha256.Sum256([]byte(contents))
	return hash[:20]
}
