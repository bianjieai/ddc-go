package keeper_test

import (
	context "context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bianjieai/ddc-go/ddc/core/auth"
)

func TestKeeperSuite1(t *testing.T) {
	suite.Run(t, new(KeeperSuite))
}

func (suite *KeeperSuite) TestKeeper_AddAccount(t *testing.T) {
	type args struct {
		goctx context.Context
		msg   *auth.MsgAddAccount
	}
	tests := []struct {
		name    string
		args    args
		wantRes *auth.MsgAddAccountResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := suite.k.AddAccount(tt.args.goctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keeper.AddAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Keeper.AddAccount() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
