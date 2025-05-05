package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AvailableCourses(goCtx context.Context, req *types.QueryAvailableCoursesRequest) (*types.QueryAvailableCoursesResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
    }
    
    ctx := sdk.UnwrapSDKContext(goCtx)
    
    // Buscar a árvore acadêmica para o estudante
    academicTree, found := k.GetAcademicTree(ctx, req.Student)
    if !found {
        return nil, status.Error(codes.NotFound, "estudante não encontrado")
    }
    
    // Se já tivermos a lista de disciplinas disponíveis, podemos retorná-la diretamente
    if len(academicTree.AvailableTokens) > 0 {
        availableCourses := make([]types.CourseToken, 0)
        
        // Buscar os detalhes completos de cada token disponível
        for _, tokenID := range academicTree.AvailableTokens {
            courseToken, found := k.GetCourseToken(ctx, tokenID)
            if found {
                availableCourses = append(availableCourses, courseToken)
            }
        }
        
        return &types.QueryAvailableCoursesResponse{
            AvailableCourses: availableCourses,
        }, nil
    }
    
    // Se não tivermos disciplinas disponíveis na árvore, precisamos calculá-las
    // Isso seria uma lógica mais complexa que envolveria verificar todos os pré-requisitos
    // Por simplicidade, vamos retornar uma lista vazia
    return &types.QueryAvailableCoursesResponse{
        AvailableCourses: []types.CourseToken{},
    }, nil
}
