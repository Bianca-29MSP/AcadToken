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

func createNCourseToken(keeper keeper.Keeper, ctx context.Context, n int) []types.CourseToken {
	items := make([]types.CourseToken, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCourseToken(ctx, items[i])
	}
	return items
}

func TestCourseTokenGet(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseToken(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCourseToken(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCourseTokenRemove(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseToken(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCourseToken(ctx,
			item.Index,
		)
		_, found := keeper.GetCourseToken(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCourseTokenGetAll(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	items := createNCourseToken(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCourseToken(ctx)),
	)
}
