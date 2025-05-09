package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferCourseNft{}

func NewMsgTransferCourseNft(creator string, nftId string, recipient string) *MsgTransferCourseNft {
	return &MsgTransferCourseNft{
		Creator:   creator,
		NftId:     nftId,
		Recipient: recipient,
	}
}

func (msg *MsgTransferCourseNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
