package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCourseContent(goCtx context.Context, msg *types.MsgCreateCourseContent) (*types.MsgCreateCourseContentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verificar se já existe um conteúdo com este ID
	_, found := k.GetCourseContent(ctx, msg.CourseId)
	if found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "conteúdo de curso já existe para o ID: %s", msg.CourseId)
	}

	// Criar o objeto CourseContent com os campos da mensagem
	courseContent := types.CourseContent{
		CourseId:                  msg.CourseId,
		Institution:               msg.Institution,
		Title:                     msg.Title,
		Code:                      msg.Code,
		WorkloadHours:             msg.WorkloadHours,
		Credits:                   msg.Credits,
		Description:               msg.Description,
		Objectives:                msg.Objectives,
		TopicUnits:                msg.TopicUnits,
		Methodologies:             msg.Methodologies,
		EvaluationMethods:         msg.EvaluationMethods,
		BibliographyBasic:         msg.BibliographyBasic,
		BibliographyComplementary: msg.BibliographyComplementary,
		Keywords:                  msg.Keywords,
		ContentHash:               msg.ContentHash,
	}

	// Salvar o conteúdo no armazenamento
	k.SetCourseContent(ctx, courseContent)

	// Emitir evento
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute("action", "create_course_content"),
			sdk.NewAttribute("course_id", msg.CourseId),
			sdk.NewAttribute("content_hash", msg.ContentHash),
		),
	)

	// Retornar a resposta vazia
	return &types.MsgCreateCourseContentResponse{}, nil
}
