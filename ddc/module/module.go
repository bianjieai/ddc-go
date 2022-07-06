package module

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/bianjieai/ddc-go/ddc"
	"github.com/bianjieai/ddc-go/ddc/client/cli"
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
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the identity module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	//return cdc.MustMarshalJSON(DefaultGenesisState())
	return nil
}

// ValidateGenesis performs genesis state validation for the identity module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	// var data GenesisState
	// if err := cdc.UnmarshalJSON(bz, &data); err != nil {
	// 	return fmt.Errorf("failed to unmarshal %s genesis state: %w", ModuleName, err)
	// }

	// return ValidateGenesis(data)
	return nil
}

// RegisterRESTRoutes registers the REST routes for the identity module.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the identity module.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	//_ = types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
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
	//types.RegisterInterfaces(registry)
}

// ____________________________________________________________________________

// AppModule implements an application module for the identity module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

// Name returns the identity module's name.
func (AppModule) Name() string {
	return ddc.ModuleName
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	// types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
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
	// var genesisState GenesisState
	// cdc.MustUnmarshalJSON(data, &genesisState)

	// InitGenesis(ctx, am.keeper, genesisState)
	// return []abci.ValidatorUpdate{}
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the identity module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	// gs := ExportGenesis(ctx, am.keeper)
	// return cdc.MustMarshalJSON(gs)
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock returns the begin blocker for the identity module.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the identity module. It returns no validator updates.
func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}
