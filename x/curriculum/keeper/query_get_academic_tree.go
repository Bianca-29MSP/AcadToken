package keeper

import (
    "context"
    "github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (k Keeper) GetAcademicTree(goCtx context.Context, req *types.QueryGetAcademicTreeRequest) (*types.QueryGetAcademicTreeResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
    
    // Buscar a árvore acadêmica usando o método renomeado
    academicTree, found := k.GetAcademicTreeByStudent(goCtx, req.Student)
    if !found {
        return nil, status.Error(codes.NotFound, "árvore acadêmica não encontrada para este estudante")
    }
    
    return &types.QueryGetAcademicTreeResponse{
        AcademicTree: &academicTree,
    }, nil
}