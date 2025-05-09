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

func (k Keeper) CourseContentAll(ctx context.Context, req *types.QueryAllCourseContentRequest) (*types.QueryAllCourseContentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var courseContents []types.CourseContent

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	courseContentStore := prefix.NewStore(store, types.KeyPrefix(types.CourseContentKeyPrefix))

	pageRes, err := query.Paginate(courseContentStore, req.Pagination, func(key []byte, value []byte) error {
		var courseContent types.CourseContent
		if err := k.cdc.Unmarshal(value, &courseContent); err != nil {
			return err
		}

		courseContents = append(courseContents, courseContent)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCourseContentResponse{CourseContent: courseContents, Pagination: pageRes}, nil
}

func (k Keeper) CourseContent(ctx context.Context, req *types.QueryGetCourseContentRequest) (*types.QueryGetCourseContentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetCourseContent(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCourseContentResponse{CourseContent: val}, nil
}
