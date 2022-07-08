package keeper

var (
	DDCApprovalKey     = []byte{0x01}
	AccountApprovalKey = []byte{0x02}
	TokenBlockListKey  = []byte{0x03}
)

const (
	SubModule = "token"
)

func ddcApprovalKey(denom, tokenID string) []byte {
	str := denom + tokenID
	key := make([]byte, len(DDCApprovalKey)+len(str))
	copy(key, DDCApprovalKey)
	copy(key[len(DDCApprovalKey):], str)
	return key
}

func accountApprovalKey(denom, owner, operator string) []byte {
	str := denom + owner + operator
	key := make([]byte, len(AccountApprovalKey)+len(str))
	copy(key, AccountApprovalKey)
	copy(key[len(AccountApprovalKey):], str)
	return key
}

func tokenBlocklistKey(denom, tokenID string) []byte {
	str := denom + tokenID
	key := make([]byte, len(TokenBlockListKey)+len(str))
	copy(key, TokenBlockListKey)
	copy(key[len(TokenBlockListKey):], str)
	return key
}
