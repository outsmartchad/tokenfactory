package tokenfactory

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "tokenfactory/api/tokenfactory/tokenfactory"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "DenomAll",
					Use:       "list-denom",
					Short:     "List all Denom",
				},
				{
					RpcMethod:      "Denom",
					Use:            "show-denom [id]",
					Short:          "Shows a Denom",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateDenom",
					Use:            "create-denom [denom] [description] [ticker] [precision] [url] [maxSupply] [supply] [canChangeMaxSupply]",
					Short:          "Create a new Denom",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "description"}, {ProtoField: "ticker"}, {ProtoField: "precision"}, {ProtoField: "url"}, {ProtoField: "maxSupply"}, {ProtoField: "supply"}, {ProtoField: "canChangeMaxSupply"}},
				},
				{
					RpcMethod:      "UpdateDenom",
					Use:            "update-denom [denom] [description] [ticker] [precision] [url] [maxSupply] [supply] [canChangeMaxSupply]",
					Short:          "Update Denom",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "description"}, {ProtoField: "ticker"}, {ProtoField: "precision"}, {ProtoField: "url"}, {ProtoField: "maxSupply"}, {ProtoField: "supply"}, {ProtoField: "canChangeMaxSupply"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
