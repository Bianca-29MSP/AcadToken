package academicnft_test

import (
	"testing"

	keepertest "github.com/Bianca-29MSP/AcademicToken/testutil/keeper"
	"github.com/Bianca-29MSP/AcademicToken/testutil/nullify"
	academicnft "github.com/Bianca-29MSP/AcademicToken/x/academicnft/module"
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CourseNftList: []types.CourseNft{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AcademicnftKeeper(t)
	academicnft.InitGenesis(ctx, k, genesisState)
	got := academicnft.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CourseNftList, got.CourseNftList)
	// this line is used by starport scaffolding # genesis/test/assert
}
