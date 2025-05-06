package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CourseEquivalenceAll(ctx context.Context, req *types.QueryAllCourseEquivalenceRequest) (*types.QueryAllCourseEquivalenceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var courseEquivalences []types.CourseEquivalence

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	courseEquivalenceStore := prefix.NewStore(store, types.KeyPrefix(types.CourseEquivalenceKeyPrefix))

	pageRes, err := query.Paginate(courseEquivalenceStore, req.Pagination, func(key []byte, value []byte) error {
		var courseEquivalence types.CourseEquivalence
		if err := k.cdc.Unmarshal(value, &courseEquivalence); err != nil {
			return err
		}

		courseEquivalences = append(courseEquivalences, courseEquivalence)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCourseEquivalenceResponse{CourseEquivalence: courseEquivalences, Pagination: pageRes}, nil
}

func (k Keeper) CourseEquivalence(ctx context.Context, req *types.QueryGetCourseEquivalenceRequest) (*types.QueryGetCourseEquivalenceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetCourseEquivalence(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCourseEquivalenceResponse{CourseEquivalence: val}, nil
}
