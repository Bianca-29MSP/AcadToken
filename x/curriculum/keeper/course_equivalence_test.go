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

func createNCourseEquivalence(keeper keeper.Keeper, ctx context.Context, n int) []types.CourseEquivalence {
	items := make([]types.CourseEquivalence, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCourseEquivalence(ctx, items[i])
	}
	return items
}

func TestCourseEquivalenceGet(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseEquivalence(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCourseEquivalence(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCourseEquivalenceRemove(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseEquivalence(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCourseEquivalence(ctx,
			item.Index,
		)
		_, found := keeper.GetCourseEquivalence(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCourseEquivalenceGetAll(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseEquivalence(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCourseEquivalence(ctx)),
	)
}
