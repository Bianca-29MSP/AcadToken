package academicnft

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/Bianca-29MSP/AcademicToken/api/academictoken/academicnft"
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
					RpcMethod: "CourseNftAll",
					Use:       "list-course-nft",
					Short:     "List all CourseNFT",
				},
				{
					RpcMethod:      "CourseNft",
					Use:            "show-course-nft [id]",
					Short:          "Shows a CourseNFT",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
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
					RpcMethod:      "MintCourseNft",
					Use:            "mint-course-nft [course-id] [institution] [title] [code] [workload-hours] [credits] [description] [objectives] [topic-units] [methodologies] [evaluation-methods] [bibliography-basic] [bibliography-complementary] [keywords] [content-hash]",
					Short:          "Send a mint-course-nft tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "courseId"}, {ProtoField: "institution"}, {ProtoField: "title"}, {ProtoField: "code"}, {ProtoField: "workloadHours"}, {ProtoField: "credits"}, {ProtoField: "description"}, {ProtoField: "objectives"}, {ProtoField: "topicUnits"}, {ProtoField: "methodologies"}, {ProtoField: "evaluationMethods"}, {ProtoField: "bibliographyBasic"}, {ProtoField: "bibliographyComplementary"}, {ProtoField: "keywords"}, {ProtoField: "contentHash"}},
				},
				{
					RpcMethod:      "TransferCourseNft",
					Use:            "transfer-course-nft [nft-id] [recipient]",
					Short:          "Send a transfer-course-nft tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nftId"}, {ProtoField: "recipient"}},
				},
				{
					RpcMethod:      "ApproveEquivalence",
					Use:            "approve-equivalence [nft-id-1] [nft-id-2] [score] [justification]",
					Short:          "Send a approve-equivalence tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nftId1"}, {ProtoField: "nftId2"}, {ProtoField: "score"}, {ProtoField: "justification"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
