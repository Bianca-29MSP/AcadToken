package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAcademicTree(goCtx context.Context, req *types.QueryGetAcademicTreeRequest) (*types.QueryGetAcademicTreeResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
    
    ctx := sdk.UnwrapSDKContext(goCtx)
    
    // Buscar a árvore acadêmica para o estudante especificado
    academicTree, found := k.GetAcademicTree(ctx, req.Student)
    if !found {
        return nil, status.Error(codes.NotFound, "árvore acadêmica não encontrada para este estudante")
    }
    
    return &types.QueryGetAcademicTreeResponse{
        AcademicTree: academicTree,
    }, nil
}
