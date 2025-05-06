package keeper

import (
    "context"
    "fmt"
    "github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) IssueCourseToken(goCtx context.Context, msg *types.MsgIssueCourseToken) (*types.MsgIssueCourseTokenResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    
    // Verificar se o emissor é uma instituição autorizada
    institution, found := k.Keeper.GetInstitution(goCtx, msg.Creator)
    if !found || !institution.IsAuthorized {
        return nil, fmt.Errorf("apenas instituições autorizadas podem emitir tokens de disciplina: %w", sdkerrors.ErrUnauthorized)
    }
    
    // Verificar se o ID do token já existe
    _, found = k.Keeper.GetCourseToken(goCtx, msg.TokenId)
    if found {
        return nil, fmt.Errorf("token já existe com este ID: %w", sdkerrors.ErrInvalidRequest)
    }
    
    // Criar o token da disciplina
    courseToken := types.CourseToken{
        TokenId:         msg.TokenId,
        Name:            msg.Name,
        Code:            msg.Code,
        ContentHash:     msg.ContentHash,
        Institution:     msg.Institution,
        CompletionDate:  msg.CompletionDate,
        Grade:           msg.Grade,
        Prerequisites:   msg.Prerequisites,
        Owner:           msg.Recipient,
    }
    
    // Salvar o token no estado
    k.Keeper.SetCourseToken(goCtx, courseToken)
    
    // Obter a árvore acadêmica do estudante
    academicTree, found := k.Keeper.GetAcademicTreeByStudent(goCtx, msg.Recipient)
    
    // Se a árvore não existir, criar uma nova
    if !found {
        academicTree = types.AcademicTree{
            Student:         msg.Recipient,
            Institution:     msg.Institution,
            CompletedTokens: []string{msg.TokenId},
            AvailableTokens: []string{},
        }
    } else {
        // Adicionar à árvore existente
        academicTree.CompletedTokens = append(academicTree.CompletedTokens, msg.TokenId)
    }
    
    // Salvar a árvore atualizada
    k.Keeper.SetAcademicTree(goCtx, academicTree)
    
    // Emitir evento
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(
            "course_token_issued",
            sdk.NewAttribute("token_id", msg.TokenId),
            sdk.NewAttribute("course_code", msg.Code),
            sdk.NewAttribute("student", msg.Recipient),
            sdk.NewAttribute("institution", msg.Institution),
        ),
    )
    
    return &types.MsgIssueCourseTokenResponse{}, nil
}