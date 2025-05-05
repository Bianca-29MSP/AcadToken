package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProposeEquivalence{}

func NewMsgProposeEquivalence(creator string, sourceTokenId string, targetInstitution string, targetCourseCode string) *MsgProposeEquivalence {
	return &MsgProposeEquivalence{
		Creator:           creator,
		SourceTokenId:     sourceTokenId,
		TargetInstitution: targetInstitution,
		TargetCourseCode:  targetCourseCode,
	}
}

func (msg *MsgProposeEquivalence) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
