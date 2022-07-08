package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/core"
)

var (
	AccountKey        = []byte{0x01}
	RoleAndFunBindKey = []byte{0x02}
	CrossPlatformKey  = []byte{0x03}
	PlatformDIDKey    = []byte{0x04}
	DDCKey            = []byte{0x05}
	PlatformSwitcher  = []byte{0x06}

	Delimiter   = []byte{0x00}
	Placeholder = []byte{0x01}
)

const (
	SubModule = "auth"
)

// accountKey returns the byte representation of the AccountInfo
func accountKey(address string) []byte {
	key := make([]byte, len(AccountKey)+len(address))
	copy(key, AccountKey)
	copy(key[len(AccountKey):], []byte(address))
	return key
}

// roleAndFunBindKey returns the byte representation of the function
func prefixRoleAndFunBindKey(role core.Role, protocol core.Protocol, denom string) []byte {
	rbz := sdk.Uint64ToBigEndian(uint64(role))
	pbz := sdk.Uint64ToBigEndian(uint64(protocol))
	dbz := []byte(denom)

	len := len(RoleAndFunBindKey) + len(dbz) + len(rbz) + len(pbz)

	b := bytes.NewBuffer(make([]byte, 0, len))
	b.Write(RoleAndFunBindKey)
	b.Write(rbz)
	b.Write(pbz)
	b.Write(dbz)
	return b.Bytes()
}

func funKey(function core.Function) []byte {
	return sdk.Uint64ToBigEndian(uint64(function))
}

func crossPlatformKey(fromDID, toDID string) []byte {
	fromBz := []byte(fromDID)
	toBz := []byte(toDID)

	len := len(CrossPlatformKey) + len(fromBz) + len(Delimiter) + len(toBz)

	b := bytes.NewBuffer(make([]byte, 0, len))
	b.Write(CrossPlatformKey)
	b.Write(fromBz)
	b.Write(Delimiter)
	b.Write(toBz)
	return b.Bytes()
}

// accountKey returns the byte representation of the AccountInfo
func platformDIDKey(accountDID string) []byte {
	key := make([]byte, len(PlatformDIDKey)+len(accountDID))
	copy(key, PlatformDIDKey)
	copy(key[len(PlatformDIDKey):], []byte(accountDID))
	return key
}

// ddcKey returns the byte representation of the ddc
func ddcKey(denomID string) []byte {
	key := make([]byte, len(DDCKey)+len(denomID))
	copy(key, DDCKey)
	copy(key[len(DDCKey):], denomID)
	return key
}

// platformSwitcher returns the byte representation of the ddc
func platformSwitcher() []byte {
	return PlatformSwitcher
}
