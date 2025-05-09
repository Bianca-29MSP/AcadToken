package curriculum_test

import (
	"testing"

	keepertest "github.com/Bianca-29MSP/AcademicToken/testutil/keeper"
	"github.com/Bianca-29MSP/AcademicToken/testutil/nullify"
	curriculum "github.com/Bianca-29MSP/AcademicToken/x/curriculum/module"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		InstitutionList: []types.Institution{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CourseTokenList: []types.CourseToken{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CourseEquivalenceList: []types.CourseEquivalence{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CourseContentList: []types.CourseContent{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CurriculumKeeper(t)
	curriculum.InitGenesis(ctx, k, genesisState)
	got := curriculum.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.InstitutionList, got.InstitutionList)
	require.ElementsMatch(t, genesisState.CourseTokenList, got.CourseTokenList)
	require.ElementsMatch(t, genesisState.CourseEquivalenceList, got.CourseEquivalenceList)
	require.ElementsMatch(t, genesisState.CourseContentList, got.CourseContentList)
	// this line is used by starport scaffolding # genesis/test/assert
}
