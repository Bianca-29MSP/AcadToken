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

func (k Keeper) InstitutionAll(ctx context.Context, req *types.QueryAllInstitutionRequest) (*types.QueryAllInstitutionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var institutions []types.Institution

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	institutionStore := prefix.NewStore(store, types.KeyPrefix(types.InstitutionKeyPrefix))

	pageRes, err := query.Paginate(institutionStore, req.Pagination, func(key []byte, value []byte) error {
		var institution types.Institution
		if err := k.cdc.Unmarshal(value, &institution); err != nil {
			return err
		}

		institutions = append(institutions, institution)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInstitutionResponse{Institution: institutions, Pagination: pageRes}, nil
}

func (k Keeper) Institution(ctx context.Context, req *types.QueryGetInstitutionRequest) (*types.QueryGetInstitutionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetInstitution(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInstitutionResponse{Institution: val}, nil
}
