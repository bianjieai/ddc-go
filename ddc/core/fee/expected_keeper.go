package fee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/ddc-go/ddc/core"
)

type AuthKeeper interface {
	CheckAvailableAndRole(ctx sdk.Context, sender string, role core.Role) error
	GetAccount(ctx sdk.Context, address string) (account *core.AccountInfo, err error)
}
