package keeper

var (
	AccountKey        = []byte{0x01}
	RoleAndFunBindKey = []byte{0x02}
	CrossPlatformKey  = []byte{0x03}
	PlatformDIDKey    = []byte{0x04}
	DDCKey            = []byte{0x05}
)

const (
	SubModule = "auth"
)

// ddcKey returns the byte representation of the ddc
func ddcKey(denomID string) []byte {
	key := make([]byte, len(DDCKey)+len(denomID))
	copy(key, DDCKey)
	copy(key[len(DDCKey):], denomID)
	return key
}
