package keeper

import (
    "context"
    "github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
    // sdk "github.com/cosmos/cosmos-sdk/types"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (k Keeper) AvailableCourses(goCtx context.Context, req *types.QueryAvailableCoursesRequest) (*types.QueryAvailableCoursesResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
    }
    
    // Buscar a árvore acadêmica para o estudante usando o método renomeado
    academicTree, found := k.GetAcademicTreeByStudent(goCtx, req.Student)
    if !found {
        return nil, status.Error(codes.NotFound, "estudante não encontrado")
    }
    
    // Se já tivermos a lista de disciplinas disponíveis, podemos retorná-la diretamente
    if len(academicTree.AvailableTokens) > 0 {
        availableCourses := make([]types.CourseToken, 0)
        
        // Buscar os detalhes completos de cada token disponível
        for _, tokenId := range academicTree.AvailableTokens {
            courseToken, found := k.GetCourseToken(goCtx, tokenId)
            if found {
                availableCourses = append(availableCourses, courseToken)
            }
        }
        
        // Retornar apenas uma resposta vazia pois o tipo não tem o campo AvailableCourses
        return &types.QueryAvailableCoursesResponse{}, nil
    }
    
    // Se não tivermos disciplinas disponíveis na árvore, retornar apenas resposta vazia
    return &types.QueryAvailableCoursesResponse{}, nil
}
