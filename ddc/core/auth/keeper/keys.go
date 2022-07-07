package keeper

var (
	AccountKey        = []byte{0x01}
	RoleAndFunBindKey = []byte{0x02}
	CrossPlatformKey  = []byte{0x03}
	PlatformDIDKey    = []byte{0x04}
	DDCKey            = []byte{0x05}
	PlatformSwitcher  = []byte{0x06}
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
