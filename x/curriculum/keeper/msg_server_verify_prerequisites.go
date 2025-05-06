package keeper

import (
    "context"
    "fmt"
    "github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
    // sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) VerifyPrerequisites(goCtx context.Context, msg *types.MsgVerifyPrerequisites) (*types.MsgVerifyPrerequisitesResponse, error) {
    // Obter a árvore acadêmica do estudante
    academicTree, found := k.Keeper.GetAcademicTreeByStudent(goCtx, msg.Student)
    if !found {
        return nil, fmt.Errorf("estudante não encontrado: %w", sdkerrors.ErrNotFound)
    }
    
    var targetCourse types.CourseToken
    var courseFound bool
    
    // Buscar disciplina percorrendo todos os tokens
    allTokens := k.Keeper.GetAllCourseToken(goCtx)
    for _, token := range allTokens {
        if token.Code == msg.CourseCode && token.Institution == academicTree.Institution {
            targetCourse = token
            courseFound = true
            break
        }
    }
    
    if !courseFound {
        return nil, fmt.Errorf("disciplina não encontrada: %w", sdkerrors.ErrNotFound)
    }
    
    // Verificar se todos os pré-requisitos foram cumpridos
    // Removido allPrerequisitesMet que não estava sendo usado
    missingPrerequisites := []string{}
    
    for _, prereqId := range targetCourse.Prerequisites {
        prereqMet := false
        for _, completedId := range academicTree.CompletedTokens {
            if completedId == prereqId {
                prereqMet = true
                break
            }
        }
        
        if !prereqMet {
            // Encontrar o nome do pré-requisito para incluir na mensagem
            prereqToken, found := k.Keeper.GetCourseToken(goCtx, prereqId)
            if found {
                missingPrerequisites = append(missingPrerequisites, prereqToken.Name)
            } else {
                missingPrerequisites = append(missingPrerequisites, prereqId)
            }
        }
    }
    
    // Retornar a resposta sem os campos que não existem na definição do tipo
    return &types.MsgVerifyPrerequisitesResponse{}, nil
}
