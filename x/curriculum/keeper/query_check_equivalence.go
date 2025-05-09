package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CheckEquivalence(goCtx context.Context, req *types.QueryCheckEquivalenceRequest) (*types.QueryCheckEquivalenceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Verificar se o token existe, mas descartar a variável com _
	_, found := k.GetCourseToken(goCtx, req.TokenId)
	if !found {
		return nil, status.Error(codes.NotFound, "token não encontrado")
	}

	// Buscar todas as equivalências para este token
	equivalences := make([]types.CourseEquivalence, 0)

	// Percorrer todas as equivalências
	allEquivalences := k.GetAllCourseEquivalence(goCtx)
	for _, equivalence := range allEquivalences {
		if equivalence.SourceTokenId == req.TokenId && equivalence.TargetInstitution == req.Institution {
			equivalences = append(equivalences, equivalence)
		}
	}

	return &types.QueryCheckEquivalenceResponse{}, nil
}
