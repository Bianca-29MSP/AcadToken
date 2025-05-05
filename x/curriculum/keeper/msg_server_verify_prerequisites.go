package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) VerifyPrerequisites(goCtx context.Context, msg *types.MsgVerifyPrerequisites) (*types.MsgVerifyPrerequisitesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Obter a árvore acadêmica do estudante
	academicTree, found := k.GetAcademicTree(ctx, msg.Student)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "estudante não encontrado")
	}

	// Buscar informações sobre a disciplina
	// Aqui você precisaria de uma maneira de encontrar a disciplina pelo código
	// Por exemplo, percorrendo todos os tokens para encontrar a correspondência
	var targetCourse types.CourseToken
	var courseFound bool

	// Implementação simplificada: percorrendo tokens para encontrar a disciplina alvo
	// Em um sistema real, você teria um índice para buscar disciplinas por código
	allTokens := k.GetAllCourseToken(ctx)
	for _, token := range allTokens {
		if token.Code == msg.CourseCode && token.Institution == academicTree.Institution {
			targetCourse = token
			courseFound = true
			break
		}
	}

	if !courseFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "disciplina não encontrada")
	}

	// Verificar se todos os pré-requisitos foram cumpridos
	allPrerequisitesMet := true
	missingPrerequisites := []string{}

	for _, prereqID := range targetCourse.Prerequisites {
		prereqMet := false
		for _, completedID := range academicTree.CompletedTokens {
			if completedID == prereqID {
				prereqMet = true
				break
			}
		}

		if !prereqMet {
			allPrerequisitesMet = false
			// Encontrar o nome do pré-requisito para incluir na mensagem
			prereqToken, found := k.GetCourseToken(ctx, prereqID)
			if found {
				missingPrerequisites = append(missingPrerequisites, prereqToken.Name)
			} else {
				missingPrerequisites = append(missingPrerequisites, prereqID)
			}
		}
	}

	// Retornar a resposta com o resultado da verificação
	return &types.MsgVerifyPrerequisitesResponse{
		IsMet:                allPrerequisitesMet,
		MissingPrerequisites: missingPrerequisites,
	}, nil
}
