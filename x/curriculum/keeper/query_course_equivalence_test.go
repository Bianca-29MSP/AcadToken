package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/Bianca-29MSP/AcademicToken/testutil/keeper"
	"github.com/Bianca-29MSP/AcademicToken/testutil/nullify"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCourseEquivalenceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	msgs := createNCourseEquivalence(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetCourseEquivalenceRequest
		response *types.QueryGetCourseEquivalenceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCourseEquivalenceRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetCourseEquivalenceResponse{CourseEquivalence: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCourseEquivalenceRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetCourseEquivalenceResponse{CourseEquivalence: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCourseEquivalenceRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CourseEquivalence(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestCourseEquivalenceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CurriculumKeeper(t)
	msgs := createNCourseEquivalence(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCourseEquivalenceRequest {
		return &types.QueryAllCourseEquivalenceRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CourseEquivalenceAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CourseEquivalence), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CourseEquivalence),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CourseEquivalenceAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CourseEquivalence), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CourseEquivalence),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CourseEquivalenceAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.CourseEquivalence),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CourseEquivalenceAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
