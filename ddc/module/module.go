package module

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/bianjieai/ddc-go/ddc"
	"github.com/bianjieai/ddc-go/ddc/client/cli"
	"github.com/bianjieai/ddc-go/ddc/core"
	"github.com/bianjieai/ddc-go/ddc/core/auth"
	"github.com/bianjieai/ddc-go/ddc/core/fee"
	"github.com/bianjieai/ddc-go/ddc/core/token"
	"github.com/bianjieai/ddc-go/ddc/keeper"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the identity module.
type AppModuleBasic struct{}

// Name returns the identity module's name.
func (AppModuleBasic) Name() string {
	return ddc.ModuleName
}

// RegisterLegacyAminoCodec registers the identity module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// DefaultGenesis returns default genesis state as raw bytes for the identity module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(ddc.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the identity module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data core.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", ddc.ModuleName, err)
	}

	return ddc.ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the identity module.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the identity module.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	auth.RegisterQueryHandlerClient(context.Background(), mux, auth.NewQueryClient(clientCtx))
	fee.RegisterQueryHandlerClient(context.Background(), mux, fee.NewQueryClient(clientCtx))
	token.RegisterQueryHandlerClient(context.Background(), mux, token.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the identity module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for the identity module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers interfaces and implementations of the identity module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	ddc.RegisterInterfaces(registry)
}

// ____________________________________________________________________________

// AppModule implements an application module for the identity module.
type AppModule struct {
	AppModuleBasic
	k keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		k:              k,
	}
}

// Name returns the identity module's name.
func (AppModule) Name() string {
	return ddc.ModuleName
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	auth.RegisterMsgServer(cfg.MsgServer(), am.k.AuthKeeper)
	fee.RegisterMsgServer(cfg.MsgServer(), am.k.FeeKeeper)
	token.RegisterMsgServer(cfg.MsgServer(), am.k.TokenKeeper)

	auth.RegisterQueryServer(cfg.QueryServer(), am.k.AuthKeeper)
	fee.RegisterQueryServer(cfg.QueryServer(), am.k.FeeKeeper)
	token.RegisterQueryServer(cfg.QueryServer(), am.k.TokenKeeper)
}

// RegisterInvariants registers the identity module invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// Route returns the message routing key for the identity module.
// Deprecated: Route returns the message routing key for the bank module.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute returns the identity module's querier route name.
func (AppModule) QuerierRoute() string { return ddc.RouterKey }

// LegacyQuerierHandler returns the identity module sdk.Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// InitGenesis performs genesis initialization for the identity module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState core.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	ddc.InitGenesis(ctx, am.k, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the identity module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := ddc.ExportGenesis(ctx, am.k)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock returns the begin blocker for the identity module.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the identity module. It returns no validator updates.
func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}
