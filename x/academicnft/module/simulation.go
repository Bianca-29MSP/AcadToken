package academicnft

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/Bianca-29MSP/AcademicToken/testutil/sample"
	academicnftsimulation "github.com/Bianca-29MSP/AcademicToken/x/academicnft/simulation"
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
)

// avoid unused import issue
var (
	_ = academicnftsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgMintCourseNft = "op_weight_msg_mint_course_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintCourseNft int = 100

	opWeightMsgTransferCourseNft = "op_weight_msg_transfer_course_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferCourseNft int = 100

	opWeightMsgApproveEquivalence = "op_weight_msg_approve_equivalence"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveEquivalence int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	academicnftGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&academicnftGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMintCourseNft int
	simState.AppParams.GetOrGenerate(opWeightMsgMintCourseNft, &weightMsgMintCourseNft, nil,
		func(_ *rand.Rand) {
			weightMsgMintCourseNft = defaultWeightMsgMintCourseNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintCourseNft,
		academicnftsimulation.SimulateMsgMintCourseNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferCourseNft int
	simState.AppParams.GetOrGenerate(opWeightMsgTransferCourseNft, &weightMsgTransferCourseNft, nil,
		func(_ *rand.Rand) {
			weightMsgTransferCourseNft = defaultWeightMsgTransferCourseNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferCourseNft,
		academicnftsimulation.SimulateMsgTransferCourseNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveEquivalence int
	simState.AppParams.GetOrGenerate(opWeightMsgApproveEquivalence, &weightMsgApproveEquivalence, nil,
		func(_ *rand.Rand) {
			weightMsgApproveEquivalence = defaultWeightMsgApproveEquivalence
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveEquivalence,
		academicnftsimulation.SimulateMsgApproveEquivalence(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgMintCourseNft,
			defaultWeightMsgMintCourseNft,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				academicnftsimulation.SimulateMsgMintCourseNft(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgTransferCourseNft,
			defaultWeightMsgTransferCourseNft,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				academicnftsimulation.SimulateMsgTransferCourseNft(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApproveEquivalence,
			defaultWeightMsgApproveEquivalence,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				academicnftsimulation.SimulateMsgApproveEquivalence(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
