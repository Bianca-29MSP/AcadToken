package curriculum

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/Bianca-29MSP/AcademicToken/testutil/sample"
	curriculumsimulation "github.com/Bianca-29MSP/AcademicToken/x/curriculum/simulation"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
)

// avoid unused import issue
var (
	_ = curriculumsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterInstitution = "op_weight_msg_register_institution"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterInstitution int = 100

	opWeightMsgIssueCourseToken = "op_weight_msg_issue_course_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIssueCourseToken int = 100

	opWeightMsgVerifyPrerequisites = "op_weight_msg_verify_prerequisites"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyPrerequisites int = 100

	opWeightMsgProposeEquivalence = "op_weight_msg_propose_equivalence"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProposeEquivalence int = 100

	opWeightMsgCreateCourseContent = "op_weight_msg_create_course_content"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCourseContent int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	curriculumGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&curriculumGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterInstitution int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterInstitution, &weightMsgRegisterInstitution, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterInstitution = defaultWeightMsgRegisterInstitution
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterInstitution,
		curriculumsimulation.SimulateMsgRegisterInstitution(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgIssueCourseToken int
	simState.AppParams.GetOrGenerate(opWeightMsgIssueCourseToken, &weightMsgIssueCourseToken, nil,
		func(_ *rand.Rand) {
			weightMsgIssueCourseToken = defaultWeightMsgIssueCourseToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIssueCourseToken,
		curriculumsimulation.SimulateMsgIssueCourseToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyPrerequisites int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyPrerequisites, &weightMsgVerifyPrerequisites, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyPrerequisites = defaultWeightMsgVerifyPrerequisites
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyPrerequisites,
		curriculumsimulation.SimulateMsgVerifyPrerequisites(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProposeEquivalence int
	simState.AppParams.GetOrGenerate(opWeightMsgProposeEquivalence, &weightMsgProposeEquivalence, nil,
		func(_ *rand.Rand) {
			weightMsgProposeEquivalence = defaultWeightMsgProposeEquivalence
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeEquivalence,
		curriculumsimulation.SimulateMsgProposeEquivalence(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateCourseContent int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateCourseContent, &weightMsgCreateCourseContent, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCourseContent = defaultWeightMsgCreateCourseContent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCourseContent,
		curriculumsimulation.SimulateMsgCreateCourseContent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterInstitution,
			defaultWeightMsgRegisterInstitution,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				curriculumsimulation.SimulateMsgRegisterInstitution(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgIssueCourseToken,
			defaultWeightMsgIssueCourseToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				curriculumsimulation.SimulateMsgIssueCourseToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyPrerequisites,
			defaultWeightMsgVerifyPrerequisites,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				curriculumsimulation.SimulateMsgVerifyPrerequisites(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgProposeEquivalence,
			defaultWeightMsgProposeEquivalence,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				curriculumsimulation.SimulateMsgProposeEquivalence(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateCourseContent,
			defaultWeightMsgCreateCourseContent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				curriculumsimulation.SimulateMsgCreateCourseContent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
