package academictoken_test

import (
	"testing"

	keepertest "github.com/Bianca-29MSP/AcademicToken/testutil/keeper"
	"github.com/Bianca-29MSP/AcademicToken/testutil/nullify"
	academictoken "github.com/Bianca-29MSP/AcademicToken/x/academictoken/module"
	"github.com/Bianca-29MSP/AcademicToken/x/academictoken/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AcademictokenKeeper(t)
	academictoken.InitGenesis(ctx, k, genesisState)
	got := academictoken.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
