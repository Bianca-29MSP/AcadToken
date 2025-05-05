package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CheckEquivalence(goCtx context.Context, req *types.QueryCheckEquivalenceRequest) (*types.QueryCheckEquivalenceResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
    }
    
    ctx := sdk.UnwrapSDKContext(goCtx)
    
    // Verificar se o token existe
    sourceToken, found := k.GetCourseToken(ctx, req.TokenID)
    if !found {
        return nil, status.Error(codes.NotFound, "token não encontrado")
    }
    
    // Buscar todas as equivalências para este token
    // Isso seria mais eficiente com um índice, mas vamos fazer uma busca simples por enquanto
    equivalences := make([]types.CourseEquivalence, 0)
    
    // Percorrer todas as equivalências (ineficiente, mas funciona para demonstração)
    allEquivalences := k.GetAllCourseEquivalence(ctx)
    for _, equivalence := range allEquivalences {
        if equivalence.SourceTokenID == req.TokenID && equivalence.TargetInstitution == req.Institution {
            equivalences = append(equivalences, equivalence)
        }
    }
    
    return &types.QueryCheckEquivalenceResponse{
        Equivalences: equivalences,
    }, nil
}
