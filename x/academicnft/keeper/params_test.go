package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/Bianca-29MSP/AcademicToken/testutil/keeper"
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AcademicnftKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
