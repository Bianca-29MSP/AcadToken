package keeper

import (
    "context"
    "fmt"
    "github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ProposeEquivalence(goCtx context.Context, msg *types.MsgProposeEquivalence) (*types.MsgProposeEquivalenceResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    
    // Verificar se o emissor é uma instituição autorizada
    institution, found := k.Keeper.GetInstitution(goCtx, msg.Creator)
    if !found || !institution.IsAuthorized {
        return nil, fmt.Errorf("apenas instituições autorizadas podem propor equivalências: %w", sdkerrors.ErrUnauthorized)
    }
    
    // Verificar se o token de origem existe
    // Usando _ para evitar a variável não utilizada
    _, found = k.Keeper.GetCourseToken(goCtx, msg.SourceTokenId)
    if !found {
        return nil, fmt.Errorf("token de origem não encontrado: %w", sdkerrors.ErrNotFound)
    }
    
    // Verificar se a instituição alvo existe
    // Usando _ para evitar a variável não utilizada
    _, found = k.Keeper.GetInstitution(goCtx, msg.TargetInstitution)
    if !found {
        return nil, fmt.Errorf("instituição alvo não encontrada: %w", sdkerrors.ErrNotFound)
    }
    
    // Verificar se já existe uma proposta de equivalência para esta combinação
    // Primeiro, criar uma chave única para esta equivalência
    equivalenceId := msg.SourceTokenId + "-" + msg.TargetInstitution + "-" + msg.TargetCourseCode
    _, found = k.Keeper.GetCourseEquivalence(goCtx, equivalenceId)
    if found {
        return nil, fmt.Errorf("equivalência já proposta: %w", sdkerrors.ErrInvalidRequest)
    }
    
    // Criar nova proposta de equivalência
    equivalence := types.CourseEquivalence{
        Index:            equivalenceId, // Usar ID composto como índice
        SourceTokenId:    msg.SourceTokenId,
        TargetInstitution: msg.TargetInstitution,
        TargetCourseCode: msg.TargetCourseCode,
        EquivalenceStatus: 1, // 1 = Proposto, 2 = Aprovado, 3 = Rejeitado
        ApprovalCount:    1,  // A instituição que propõe automaticamente aprova
    }
    
    // Salvar a equivalência no estado
    k.Keeper.SetCourseEquivalence(goCtx, equivalence)
    
    // Emitir evento
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(
            "equivalence_proposed",
            sdk.NewAttribute("source_token", msg.SourceTokenId),
            sdk.NewAttribute("target_institution", msg.TargetInstitution),
            sdk.NewAttribute("target_course", msg.TargetCourseCode),
            sdk.NewAttribute("proposer", msg.Creator),
        ),
    )
    
    return &types.MsgProposeEquivalenceResponse{}, nil
}