package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/Bianca-29MSP/AcademicToken/testutil/keeper"
	"github.com/Bianca-29MSP/AcademicToken/testutil/nullify"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/keeper"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNInstitution(keeper keeper.Keeper, ctx context.Context, n int) []types.Institution {
	items := make([]types.Institution, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetInstitution(ctx, items[i])
	}
	return items
}

func TestInstitutionGet(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNInstitution(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetInstitution(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestInstitutionRemove(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNInstitution(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInstitution(ctx,
			item.Index,
		)
		_, found := keeper.GetInstitution(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestInstitutionGetAll(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNInstitution(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllInstitution(ctx)),
	)
}
