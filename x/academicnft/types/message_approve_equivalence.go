package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApproveEquivalence{}

func NewMsgApproveEquivalence(creator string, nftId1 string, nftId2 string, score string, justification string) *MsgApproveEquivalence {
	return &MsgApproveEquivalence{
		Creator:       creator,
		NftId1:        nftId1,
		NftId2:        nftId2,
		Score:         score,
		Justification: justification,
	}
}

func (msg *MsgApproveEquivalence) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
