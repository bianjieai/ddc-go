package keeper

var (
	DDCApprovalKey     = []byte{0x01}
	AccountApprovalKey = []byte{0x02}
	TokenBlackListKey  = []byte{0x03}
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

func tokenBlacklistKey(denom, tokenID string) []byte {
	str := denom + tokenID
	key := make([]byte, len(TokenBlackListKey)+len(str))
	copy(key, TokenBlackListKey)
	copy(key[len(TokenBlackListKey):], str)
	return key
}
