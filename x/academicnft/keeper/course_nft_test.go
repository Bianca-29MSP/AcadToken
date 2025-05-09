package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/Bianca-29MSP/AcademicToken/testutil/keeper"
	"github.com/Bianca-29MSP/AcademicToken/testutil/nullify"
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/keeper"
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCourseNft(keeper keeper.Keeper, ctx context.Context, n int) []types.CourseNft {
	items := make([]types.CourseNft, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCourseNft(ctx, items[i])
	}
	return items
}

func TestCourseNftGet(t *testing.T) {
	keeper, ctx := keepertest.AcademicnftKeeper(t)
	items := createNCourseNft(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCourseNft(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCourseNftRemove(t *testing.T) {
	keeper, ctx := keepertest.AcademicnftKeeper(t)
	items := createNCourseNft(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCourseNft(ctx,
			item.Index,
		)
		_, found := keeper.GetCourseNft(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCourseNftGetAll(t *testing.T) {
	keeper, ctx := keepertest.AcademicnftKeeper(t)
	items := createNCourseNft(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCourseNft(ctx)),
	)
}
