package keeper

import (
	"context"
	"fmt"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProposeEquivalence(goCtx context.Context, msg *types.MsgProposeEquivalence) (*types.MsgProposeEquivalenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verificar se o emissor é uma instituição autorizada
	institution, found := k.Keeper.GetInstitution(ctx, msg.Creator)
	if !found || !institution.IsAuthorized {
		return nil, fmt.Errorf("apenas instituições autorizadas podem propor equivalências")
	}

	// Buscar o token de origem
	sourceToken, found := k.Keeper.GetCourseToken(ctx, msg.SourceTokenId)
	if !found {
		return nil, fmt.Errorf("token de origem %s não encontrado", msg.SourceTokenId)
	}

	// Verificar se a instituição alvo existe e está autorizada
	targetInstitution, found := k.Keeper.GetInstitution(ctx, msg.TargetInstitution)
	if !found {
		return nil, fmt.Errorf("instituição alvo %s não encontrada", msg.TargetInstitution)
	}

	if !targetInstitution.IsAuthorized {
		return nil, fmt.Errorf("instituição alvo %s não está autorizada", msg.TargetInstitution)
	}

	// Criar ID único para a equivalência
	equivalenceId := msg.SourceTokenId + "-" + msg.TargetInstitution + "-" + msg.TargetCourseCode

	// Verificar se já existe uma proposta para esta combinação
	_, found = k.Keeper.GetCourseEquivalence(ctx, equivalenceId)
	if found {
		return nil, fmt.Errorf("equivalência já proposta")
	}

	// Inicializar variáveis de similaridade e status
	var similarityScore uint64 = 0
	var status uint64 = 1 // Por padrão, aguardando revisão

	// Se existir contentHash no token, buscar os conteúdos correspondentes
	if sourceToken.ContentHash != "" {
		// Implementar lógica de similaridade aqui quando tiver o tipo CourseContent implementado
		// Por enquanto, definimos um status padrão
	}

	// Criar nova proposta de equivalência com o índice definido
	equivalence := types.CourseEquivalence{
		Index:             equivalenceId, // Definir o índice aqui
		SourceTokenId:     msg.SourceTokenId,
		TargetInstitution: msg.TargetInstitution,
		TargetCourseCode:  msg.TargetCourseCode,
		EquivalenceStatus: status,
		ApprovalCount:     0,
	}

	// Salvar a equivalência no estado
	k.Keeper.SetCourseEquivalence(ctx, equivalence)

	// Emitir evento
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"propose_equivalence",
			sdk.NewAttribute("source_token", msg.SourceTokenId),
			sdk.NewAttribute("target_institution", msg.TargetInstitution),
			sdk.NewAttribute("target_course", msg.TargetCourseCode),
			sdk.NewAttribute("status", fmt.Sprintf("%d", status)),
			sdk.NewAttribute("similarity_score", fmt.Sprintf("%d", similarityScore)),
		),
	)

	return &types.MsgProposeEquivalenceResponse{}, nil
}
