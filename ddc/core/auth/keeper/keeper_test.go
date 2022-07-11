package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/bianjieai/ddc-go/ddc/core/auth/keeper"
	"github.com/bianjieai/ddc-go/simapp"
)

type KeeperSuite struct {
	suite.Suite

	ctx sdk.Context
	k   keeper.Keeper
}

func (suite *KeeperSuite) SetupTest() {
	app := simapp.Setup(false)
	suite.k = app.DDCKeeper.AuthKeeper
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperSuite))
}
