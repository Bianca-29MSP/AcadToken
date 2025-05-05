package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIssueCourseToken{}

func NewMsgIssueCourseToken(creator string, tokenId string, name string, code string, contentHash string, institution string, completionDate string, grade uint64, prerequisites []string, recipient string) *MsgIssueCourseToken {
	return &MsgIssueCourseToken{
		Creator:        creator,
		TokenId:        tokenId,
		Name:           name,
		Code:           code,
		ContentHash:    contentHash,
		Institution:    institution,
		CompletionDate: completionDate,
		Grade:          grade,
		Prerequisites:  prerequisites,
		Recipient:      recipient,
	}
}

func (msg *MsgIssueCourseToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
