package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCourseContent{}

func NewMsgCreateCourseContent(creator string, courseId string, institution string, title string, code string, workloadHours uint64, credits uint64, description string, objectives []string, topicUnits []string, methodologies []string, evaluationMethods []string, bibliographyBasic []string, bibliographyComplementary []string, keywords []string, contentHash string) *MsgCreateCourseContent {
	return &MsgCreateCourseContent{
		Creator:                   creator,
		CourseId:                  courseId,
		Institution:               institution,
		Title:                     title,
		Code:                      code,
		WorkloadHours:             workloadHours,
		Credits:                   credits,
		Description:               description,
		Objectives:                objectives,
		TopicUnits:                topicUnits,
		Methodologies:             methodologies,
		EvaluationMethods:         evaluationMethods,
		BibliographyBasic:         bibliographyBasic,
		BibliographyComplementary: bibliographyComplementary,
		Keywords:                  keywords,
		ContentHash:               contentHash,
	}
}

func (msg *MsgCreateCourseContent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
