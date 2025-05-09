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

func createNCourseContent(keeper keeper.Keeper, ctx context.Context, n int) []types.CourseContent {
	items := make([]types.CourseContent, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCourseContent(ctx, items[i])
	}
	return items
}

func TestCourseContentGet(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseContent(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCourseContent(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCourseContentRemove(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseContent(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCourseContent(ctx,
			item.Index,
		)
		_, found := keeper.GetCourseContent(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCourseContentGetAll(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseContent(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCourseContent(ctx)),
	)
}
