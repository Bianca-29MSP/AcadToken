package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) IssueCourseToken(goCtx context.Context, msg *types.MsgIssueCourseToken) (*types.MsgIssueCourseTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verificar se o emissor é uma instituição autorizada
	institution, found := k.GetInstitution(ctx, msg.Creator)
	if !found || !institution.IsAuthorized {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "apenas instituições autorizadas podem emitir tokens de disciplina")
	}

	// Verificar se o ID do token já existe
	_, found = k.GetCourseToken(ctx, msg.TokenID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token já existe com este ID")
	}

	// Criar o token da disciplina
	courseToken := types.CourseToken{
		TokenID:        msg.TokenID,
		Name:           msg.Name,
		Code:           msg.Code,
		ContentHash:    msg.ContentHash,
		Institution:    msg.Institution,
		CompletionDate: msg.CompletionDate,
		Grade:          msg.Grade,
		Prerequisites:  msg.Prerequisites,
		Owner:          msg.Recipient,
	}

	// Salvar o token no estado
	k.SetCourseToken(ctx, courseToken)

	// Atualizar a árvore acadêmica do estudante
	academicTree, found := k.GetAcademicTree(ctx, msg.Recipient)
	if !found {
		// Criar nova árvore se não existir
		academicTree = types.AcademicTree{
			Student:         msg.Recipient,
			Institution:     msg.Institution,
			CompletedTokens: []string{msg.TokenID},
			AvailableTokens: []string{},
		}
	} else {
		// Adicionar à árvore existente
		academicTree.CompletedTokens = append(academicTree.CompletedTokens, msg.TokenID)
	}

	// Salvar a árvore atualizada
	k.SetAcademicTree(ctx, academicTree)

	// Emitir evento
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"course_token_issued",
			sdk.NewAttribute("token_id", msg.TokenID),
			sdk.NewAttribute("course_code", msg.Code),
			sdk.NewAttribute("student", msg.Recipient),
			sdk.NewAttribute("institution", msg.Institution),
		),
	)

	return &types.MsgIssueCourseTokenResponse{}, nil
}
