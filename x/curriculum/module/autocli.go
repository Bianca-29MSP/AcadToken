package curriculum

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/Bianca-29MSP/AcademicToken/api/academictoken/curriculum"
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
					RpcMethod:      "GetAcademicTree",
					Use:            "get-academic-tree [student]",
					Short:          "Query get-academic-tree",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "student"}},
				},

				{
					RpcMethod:      "AvailableCourses",
					Use:            "available-courses [student]",
					Short:          "Query available-courses",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "student"}},
				},

				{
					RpcMethod:      "CheckEquivalence",
					Use:            "check-equivalence [token-id] [institution]",
					Short:          "Query check-equivalence",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "institution"}},
				},

				{
					RpcMethod:      "CriticalCourses",
					Use:            "critical-courses [institution] [threshold]",
					Short:          "Query critical-courses",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "institution"}, {ProtoField: "threshold"}},
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
					RpcMethod:      "RegisterInstitution",
					Use:            "register-institution [address] [name]",
					Short:          "Send a register-institution tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "name"}},
				},
				{
					RpcMethod:      "IssueCourseToken",
					Use:            "issue-course-token [token-id] [name] [code] [content-hash] [institution] [completion-date] [grade] [prerequisites] [recipient]",
					Short:          "Send a issue-course-token tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "name"}, {ProtoField: "code"}, {ProtoField: "contentHash"}, {ProtoField: "institution"}, {ProtoField: "completionDate"}, {ProtoField: "grade"}, {ProtoField: "prerequisites"}, {ProtoField: "recipient"}},
				},
				{
					RpcMethod:      "VerifyPrerequisites",
					Use:            "verify-prerequisites [student] [course-code]",
					Short:          "Send a verify-prerequisites tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "student"}, {ProtoField: "courseCode"}},
				},
				{
					RpcMethod:      "ProposeEquivalence",
					Use:            "propose-equivalence [source-token-id] [target-institution] [target-course-code]",
					Short:          "Send a propose-equivalence tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sourceTokenId"}, {ProtoField: "targetInstitution"}, {ProtoField: "targetCourseCode"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
