package keeper

import (
    "context"
    
    "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/query"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// NFTsByOwner implementa o handler para consulta de NFTs por proprietário
func (k Keeper) NFTsByOwner(ctx context.Context, req *types.QueryNFTsByOwnerRequest) (*types.QueryNFTsByOwnerResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
    }

    sdkCtx := sdk.UnwrapSDKContext(ctx)
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(sdkCtx))
    
    var nfts []types.CourseNft
    
    // Usando o prefixo correto para CourseNft
    courseNftStore := prefix.NewStore(store, []byte(types.CourseNftKeyPrefix))
    
    pageRes, err := query.Paginate(courseNftStore, req.Pagination, func(key []byte, value []byte) error {
        var courseNft types.CourseNft
        k.cdc.MustUnmarshal(value, &courseNft)
        
        // Verifica se o proprietário corresponde
        if courseNft.Owner == req.Owner {
            nfts = append(nfts, courseNft)
        }
        
        return nil
    })
    
    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryNFTsByOwnerResponse{
        Nfts:       nfts,
        Pagination: pageRes,
    }, nil
}