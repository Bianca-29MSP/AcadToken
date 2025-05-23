package keeper

import (
    "context"
    "cosmossdk.io/store/prefix"
    "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
    "github.com/cosmos/cosmos-sdk/runtime"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/query"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (k Keeper) CourseNftAll(ctx context.Context, req *types.QueryAllCourseNftRequest) (*types.QueryAllCourseNftResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
    
    var courseNfts []types.CourseNft
    sdkCtx := sdk.UnwrapSDKContext(ctx)
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(sdkCtx))
    courseNftStore := prefix.NewStore(store, types.CourseNftKeyPrefix)
    
    pageRes, err := query.Paginate(courseNftStore, req.Pagination, func(key []byte, value []byte) error {
        var courseNft types.CourseNft
        if err := k.cdc.Unmarshal(value, &courseNft); err != nil {
            return err
        }
        courseNfts = append(courseNfts, courseNft)
        return nil
    })
    
    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }
    
    return &types.QueryAllCourseNftResponse{CourseNft: courseNfts, Pagination: pageRes}, nil
}

func (k Keeper) CourseNft(ctx context.Context, req *types.QueryGetCourseNftRequest) (*types.QueryGetCourseNftResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
    
    sdkCtx := sdk.UnwrapSDKContext(ctx)
    val, found := k.GetCourseNft(sdkCtx, req.Index)
    
    if !found {
        return nil, status.Error(codes.NotFound, "not found")
    }
    
    return &types.QueryGetCourseNftResponse{CourseNft: val}, nil
}