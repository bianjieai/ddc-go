package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/core"
)

var (
	BalanceKey   = []byte{0x01}
	FeeRuleKey   = []byte{0x02}
	DenomAuthKey = []byte{0x03}
	SupplyKey    = []byte{0x04}

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
