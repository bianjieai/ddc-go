package keeper

import (
	context "context"

	"github.com/bianjieai/ddc-go/ddc/core/token"
)

var _ token.QueryServer = Keeper{}

// ApprovedForAll implements token.QueryServer
func (Keeper) ApprovedForAll(context.Context, *token.QueryApprovedForAllRequest) (*token.QueryApprovedForAllResponse, error) {
	panic("unimplemented")
}

// TokenApproved implements token.QueryServer
func (Keeper) TokenApproved(context.Context, *token.QueryTokenApprovedRequest) (*token.QueryTokenApprovedResponse, error) {
	panic("unimplemented")
}
