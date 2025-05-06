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

func (k Keeper) CourseTokenAll(ctx context.Context, req *types.QueryAllCourseTokenRequest) (*types.QueryAllCourseTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var courseTokens []types.CourseToken

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	courseTokenStore := prefix.NewStore(store, types.KeyPrefix(types.CourseTokenKeyPrefix))

	pageRes, err := query.Paginate(courseTokenStore, req.Pagination, func(key []byte, value []byte) error {
		var courseToken types.CourseToken
		if err := k.cdc.Unmarshal(value, &courseToken); err != nil {
			return err
		}

		courseTokens = append(courseTokens, courseToken)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCourseTokenResponse{CourseToken: courseTokens, Pagination: pageRes}, nil
}

func (k Keeper) CourseToken(ctx context.Context, req *types.QueryGetCourseTokenRequest) (*types.QueryGetCourseTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetCourseToken(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCourseTokenResponse{CourseToken: val}, nil
}
