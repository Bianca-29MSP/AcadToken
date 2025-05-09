package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CriticalCourses(goCtx context.Context, req *types.QueryCriticalCoursesRequest) (*types.QueryCriticalCoursesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verificar se a instituição existe
	_, found := k.GetInstitution(ctx, req.Institution)
	if !found {
		return nil, status.Error(codes.NotFound, "instituição não encontrada")
	}

	// Mapa para contar quantas vezes cada disciplina aparece como pré-requisito
	criticalityMap := make(map[string]uint64) // Alterado para uint64

	// Percorrer todos os tokens de disciplina
	allCourseTokens := k.GetAllCourseToken(ctx)
	for _, token := range allCourseTokens {
		// Considerar apenas disciplinas da instituição solicitada
		if token.Institution == req.Institution {
			// Incrementar contador para cada pré-requisito
			for _, prereq := range token.Prerequisites {
				criticalityMap[prereq]++
			}
		}
	}

	// Filtrar apenas disciplinas que atendem ao threshold de criticidade
	criticalCourses := make([]types.CourseToken, 0)
	for tokenId, criticality := range criticalityMap {
		if criticality >= req.Threshold { // Comparação agora é compatível com uint64
			// Buscar os detalhes da disciplina
			courseToken, found := k.GetCourseToken(ctx, tokenId)
			if found {
				criticalCourses = append(criticalCourses, courseToken)
			}
		}
	}

	// Retornar resposta vazia, pois o tipo não tem o campo CriticalCourses
	return &types.QueryCriticalCoursesResponse{}, nil
}
